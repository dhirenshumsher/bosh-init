package installation

import (
	biinstallmanifest "github.com/cloudfoundry/bosh-init/installation/manifest"
	bireljob "github.com/cloudfoundry/bosh-init/release/job"
	bitemplate "github.com/cloudfoundry/bosh-init/templatescompiler"
	biui "github.com/cloudfoundry/bosh-init/ui"
	boshblob "github.com/cloudfoundry/bosh-utils/blobstore"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshcmd "github.com/cloudfoundry/bosh-utils/fileutil"
	biproperty "github.com/cloudfoundry/bosh-utils/property"
)

type JobRenderer interface {
	RenderAndUploadFrom(biinstallmanifest.Manifest, []bireljob.Job, biui.Stage) ([]RenderedJobRef, error)
}

type jobRenderer struct {
	jobListRenderer bitemplate.JobListRenderer
	compressor      boshcmd.Compressor
	blobstore       boshblob.Blobstore
}

type RenderedJobRef struct {
	Name    string
	Version string
	// BlobstoreID string
	// SHA1        string
	Path string
}

// func NewRenderedJobRef(name, version, blobstoreID, sha1 string) RenderedJobRef {
func NewRenderedJobRef(name, version, path string) RenderedJobRef {
	return RenderedJobRef{
		Name:    name,
		Version: version,
		// BlobstoreID: blobstoreID,
		// SHA1:        sha1,
		Path: path,
	}
}

func NewJobRenderer(
	jobListRenderer bitemplate.JobListRenderer,
	compressor boshcmd.Compressor,
	blobstore boshblob.Blobstore,
) JobRenderer {
	return &jobRenderer{
		jobListRenderer: jobListRenderer,
		compressor:      compressor,
		blobstore:       blobstore,
	}
}

func (b *jobRenderer) RenderAndUploadFrom(installationManifest biinstallmanifest.Manifest, jobs []bireljob.Job, stage biui.Stage) ([]RenderedJobRef, error) {
	// installation jobs do not get rendered with global deployment properties, only the cloud_provider properties
	globalProperties := biproperty.Map{}
	jobProperties := installationManifest.Properties

	renderedJobRefs, err := b.renderJobTemplates(jobs, jobProperties, globalProperties, installationManifest.Name, stage)
	if err != nil {
		return nil, bosherr.WrapError(err, "Rendering job templates for installation")
	}

	if len(renderedJobRefs) != 1 {
		return nil, bosherr.Error("Too many jobs rendered... oops?")
	}

	return renderedJobRefs, nil
}

// renderJobTemplates renders all the release job templates for multiple release jobs specified
// by a deployment job and randomly uploads them to blobstore
func (b *jobRenderer) renderJobTemplates(
	releaseJobs []bireljob.Job,
	jobProperties biproperty.Map,
	globalProperties biproperty.Map,
	deploymentName string,
	stage biui.Stage,
) ([]RenderedJobRef, error) {
	renderedJobRefs := make([]RenderedJobRef, 0, len(releaseJobs))
	err := stage.Perform("Rendering job templates", func() error {
		renderedJobList, err := b.jobListRenderer.Render(releaseJobs, jobProperties, globalProperties, deploymentName)
		if err != nil {
			return err
		}
		defer renderedJobList.DeleteSilently()

		for _, renderedJob := range renderedJobList.All() {
			renderedJobRef, err := b.compressAndUpload(renderedJob)
			if err != nil {
				return err
			}

			renderedJobRefs = append(renderedJobRefs, renderedJobRef)
		}

		return nil
	})

	return renderedJobRefs, err
}

func (b *jobRenderer) compressAndUpload(renderedJob bitemplate.RenderedJob) (RenderedJobRef, error) {
	tarballPath, err := b.compressor.CompressFilesInDir(renderedJob.Path())
	if err != nil {
		return RenderedJobRef{}, bosherr.WrapError(err, "Compressing rendered job templates")
	}
	// defer b.compressor.CleanUp(tarballPath)

	// blobID, blobSHA1, err := b.blobstore.Create(tarballPath)
	// if err != nil {
	// 	return RenderedJobRef{}, bosherr.WrapError(err, "Creating blob")
	// }

	releaseJob := renderedJob.Job()

	return RenderedJobRef{
		Name:    releaseJob.Name,
		Version: releaseJob.Fingerprint,
		// BlobstoreID: blobID,
		// SHA1:        blobSHA1,
		Path: tarballPath,
	}, nil
}
