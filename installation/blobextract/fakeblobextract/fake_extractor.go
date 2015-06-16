// This file was generated by counterfeiter
package fakeblobextract

import (
	"sync"

	"github.com/cloudfoundry/bosh-init/installation/blobextract"
)

type FakeExtractor struct {
	ExtractPkgStub        func(blobID, blobSHA1, targetDir string) error
	extractPkgMutex       sync.RWMutex
	extractPkgArgsForCall []struct {
		blobID    string
		blobSHA1  string
		targetDir string
	}
	extractPkgReturns struct {
		result1 error
	}
	ExtractStub        func(renderedJobRef struct{ Name, Version, Path string }, targetDir string) error
	extractMutex       sync.RWMutex
	extractArgsForCall []struct {
		renderedJobRef struct{ Name, Version, Path string }
		targetDir      string
	}
	extractReturns struct {
		result1 error
	}
	CleanupStub        func(extractedBlobPath string) error
	cleanupMutex       sync.RWMutex
	cleanupArgsForCall []struct {
		extractedBlobPath string
	}
	cleanupReturns struct {
		result1 error
	}
	ChmodExecutablesStub        func(binPath string) error
	chmodExecutablesMutex       sync.RWMutex
	chmodExecutablesArgsForCall []struct {
		binPath string
	}
	chmodExecutablesReturns struct {
		result1 error
	}
}

func (fake *FakeExtractor) ExtractPkg(blobID string, blobSHA1 string, targetDir string) error {
	fake.extractPkgMutex.Lock()
	fake.extractPkgArgsForCall = append(fake.extractPkgArgsForCall, struct {
		blobID    string
		blobSHA1  string
		targetDir string
	}{blobID, blobSHA1, targetDir})
	fake.extractPkgMutex.Unlock()
	if fake.ExtractPkgStub != nil {
		return fake.ExtractPkgStub(blobID, blobSHA1, targetDir)
	} else {
		return fake.extractPkgReturns.result1
	}
}

func (fake *FakeExtractor) ExtractPkgCallCount() int {
	fake.extractPkgMutex.RLock()
	defer fake.extractPkgMutex.RUnlock()
	return len(fake.extractPkgArgsForCall)
}

func (fake *FakeExtractor) ExtractPkgArgsForCall(i int) (string, string, string) {
	fake.extractPkgMutex.RLock()
	defer fake.extractPkgMutex.RUnlock()
	return fake.extractPkgArgsForCall[i].blobID, fake.extractPkgArgsForCall[i].blobSHA1, fake.extractPkgArgsForCall[i].targetDir
}

func (fake *FakeExtractor) ExtractPkgReturns(result1 error) {
	fake.ExtractPkgStub = nil
	fake.extractPkgReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExtractor) Extract(renderedJobRef struct{ Name, Version, Path string }, targetDir string) error {
	fake.extractMutex.Lock()
	fake.extractArgsForCall = append(fake.extractArgsForCall, struct {
		renderedJobRef struct{ Name, Version, Path string }
		targetDir      string
	}{renderedJobRef, targetDir})
	fake.extractMutex.Unlock()
	if fake.ExtractStub != nil {
		return fake.ExtractStub(renderedJobRef, targetDir)
	} else {
		return fake.extractReturns.result1
	}
}

func (fake *FakeExtractor) ExtractCallCount() int {
	fake.extractMutex.RLock()
	defer fake.extractMutex.RUnlock()
	return len(fake.extractArgsForCall)
}

func (fake *FakeExtractor) ExtractArgsForCall(i int) (struct{ Name, Version, Path string }, string) {
	fake.extractMutex.RLock()
	defer fake.extractMutex.RUnlock()
	return fake.extractArgsForCall[i].renderedJobRef, fake.extractArgsForCall[i].targetDir
}

func (fake *FakeExtractor) ExtractReturns(result1 error) {
	fake.ExtractStub = nil
	fake.extractReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExtractor) Cleanup(extractedBlobPath string) error {
	fake.cleanupMutex.Lock()
	fake.cleanupArgsForCall = append(fake.cleanupArgsForCall, struct {
		extractedBlobPath string
	}{extractedBlobPath})
	fake.cleanupMutex.Unlock()
	if fake.CleanupStub != nil {
		return fake.CleanupStub(extractedBlobPath)
	} else {
		return fake.cleanupReturns.result1
	}
}

func (fake *FakeExtractor) CleanupCallCount() int {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return len(fake.cleanupArgsForCall)
}

func (fake *FakeExtractor) CleanupArgsForCall(i int) string {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return fake.cleanupArgsForCall[i].extractedBlobPath
}

func (fake *FakeExtractor) CleanupReturns(result1 error) {
	fake.CleanupStub = nil
	fake.cleanupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExtractor) ChmodExecutables(binPath string) error {
	fake.chmodExecutablesMutex.Lock()
	fake.chmodExecutablesArgsForCall = append(fake.chmodExecutablesArgsForCall, struct {
		binPath string
	}{binPath})
	fake.chmodExecutablesMutex.Unlock()
	if fake.ChmodExecutablesStub != nil {
		return fake.ChmodExecutablesStub(binPath)
	} else {
		return fake.chmodExecutablesReturns.result1
	}
}

func (fake *FakeExtractor) ChmodExecutablesCallCount() int {
	fake.chmodExecutablesMutex.RLock()
	defer fake.chmodExecutablesMutex.RUnlock()
	return len(fake.chmodExecutablesArgsForCall)
}

func (fake *FakeExtractor) ChmodExecutablesArgsForCall(i int) string {
	fake.chmodExecutablesMutex.RLock()
	defer fake.chmodExecutablesMutex.RUnlock()
	return fake.chmodExecutablesArgsForCall[i].binPath
}

func (fake *FakeExtractor) ChmodExecutablesReturns(result1 error) {
	fake.ChmodExecutablesStub = nil
	fake.chmodExecutablesReturns = struct {
		result1 error
	}{result1}
}

var _ blobextract.Extractor = new(FakeExtractor)
