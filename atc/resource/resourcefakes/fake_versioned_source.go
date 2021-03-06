// Code generated by counterfeiter. DO NOT EDIT.
package resourcefakes

import (
	io "io"
	sync "sync"

	atc "github.com/concourse/concourse/atc"
	resource "github.com/concourse/concourse/atc/resource"
	worker "github.com/concourse/concourse/atc/worker"
)

type FakeVersionedSource struct {
	MetadataStub        func() []atc.MetadataField
	metadataMutex       sync.RWMutex
	metadataArgsForCall []struct {
	}
	metadataReturns struct {
		result1 []atc.MetadataField
	}
	metadataReturnsOnCall map[int]struct {
		result1 []atc.MetadataField
	}
	StreamInStub        func(string, io.Reader) error
	streamInMutex       sync.RWMutex
	streamInArgsForCall []struct {
		arg1 string
		arg2 io.Reader
	}
	streamInReturns struct {
		result1 error
	}
	streamInReturnsOnCall map[int]struct {
		result1 error
	}
	StreamOutStub        func(string) (io.ReadCloser, error)
	streamOutMutex       sync.RWMutex
	streamOutArgsForCall []struct {
		arg1 string
	}
	streamOutReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	streamOutReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	VersionStub        func() atc.Version
	versionMutex       sync.RWMutex
	versionArgsForCall []struct {
	}
	versionReturns struct {
		result1 atc.Version
	}
	versionReturnsOnCall map[int]struct {
		result1 atc.Version
	}
	VolumeStub        func() worker.Volume
	volumeMutex       sync.RWMutex
	volumeArgsForCall []struct {
	}
	volumeReturns struct {
		result1 worker.Volume
	}
	volumeReturnsOnCall map[int]struct {
		result1 worker.Volume
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVersionedSource) Metadata() []atc.MetadataField {
	fake.metadataMutex.Lock()
	ret, specificReturn := fake.metadataReturnsOnCall[len(fake.metadataArgsForCall)]
	fake.metadataArgsForCall = append(fake.metadataArgsForCall, struct {
	}{})
	fake.recordInvocation("Metadata", []interface{}{})
	fake.metadataMutex.Unlock()
	if fake.MetadataStub != nil {
		return fake.MetadataStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.metadataReturns
	return fakeReturns.result1
}

func (fake *FakeVersionedSource) MetadataCallCount() int {
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	return len(fake.metadataArgsForCall)
}

func (fake *FakeVersionedSource) MetadataCalls(stub func() []atc.MetadataField) {
	fake.metadataMutex.Lock()
	defer fake.metadataMutex.Unlock()
	fake.MetadataStub = stub
}

func (fake *FakeVersionedSource) MetadataReturns(result1 []atc.MetadataField) {
	fake.metadataMutex.Lock()
	defer fake.metadataMutex.Unlock()
	fake.MetadataStub = nil
	fake.metadataReturns = struct {
		result1 []atc.MetadataField
	}{result1}
}

func (fake *FakeVersionedSource) MetadataReturnsOnCall(i int, result1 []atc.MetadataField) {
	fake.metadataMutex.Lock()
	defer fake.metadataMutex.Unlock()
	fake.MetadataStub = nil
	if fake.metadataReturnsOnCall == nil {
		fake.metadataReturnsOnCall = make(map[int]struct {
			result1 []atc.MetadataField
		})
	}
	fake.metadataReturnsOnCall[i] = struct {
		result1 []atc.MetadataField
	}{result1}
}

func (fake *FakeVersionedSource) StreamIn(arg1 string, arg2 io.Reader) error {
	fake.streamInMutex.Lock()
	ret, specificReturn := fake.streamInReturnsOnCall[len(fake.streamInArgsForCall)]
	fake.streamInArgsForCall = append(fake.streamInArgsForCall, struct {
		arg1 string
		arg2 io.Reader
	}{arg1, arg2})
	fake.recordInvocation("StreamIn", []interface{}{arg1, arg2})
	fake.streamInMutex.Unlock()
	if fake.StreamInStub != nil {
		return fake.StreamInStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.streamInReturns
	return fakeReturns.result1
}

func (fake *FakeVersionedSource) StreamInCallCount() int {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return len(fake.streamInArgsForCall)
}

func (fake *FakeVersionedSource) StreamInCalls(stub func(string, io.Reader) error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = stub
}

func (fake *FakeVersionedSource) StreamInArgsForCall(i int) (string, io.Reader) {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	argsForCall := fake.streamInArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVersionedSource) StreamInReturns(result1 error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = nil
	fake.streamInReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVersionedSource) StreamInReturnsOnCall(i int, result1 error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = nil
	if fake.streamInReturnsOnCall == nil {
		fake.streamInReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.streamInReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVersionedSource) StreamOut(arg1 string) (io.ReadCloser, error) {
	fake.streamOutMutex.Lock()
	ret, specificReturn := fake.streamOutReturnsOnCall[len(fake.streamOutArgsForCall)]
	fake.streamOutArgsForCall = append(fake.streamOutArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("StreamOut", []interface{}{arg1})
	fake.streamOutMutex.Unlock()
	if fake.StreamOutStub != nil {
		return fake.StreamOutStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.streamOutReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVersionedSource) StreamOutCallCount() int {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return len(fake.streamOutArgsForCall)
}

func (fake *FakeVersionedSource) StreamOutCalls(stub func(string) (io.ReadCloser, error)) {
	fake.streamOutMutex.Lock()
	defer fake.streamOutMutex.Unlock()
	fake.StreamOutStub = stub
}

func (fake *FakeVersionedSource) StreamOutArgsForCall(i int) string {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	argsForCall := fake.streamOutArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVersionedSource) StreamOutReturns(result1 io.ReadCloser, result2 error) {
	fake.streamOutMutex.Lock()
	defer fake.streamOutMutex.Unlock()
	fake.StreamOutStub = nil
	fake.streamOutReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeVersionedSource) StreamOutReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.streamOutMutex.Lock()
	defer fake.streamOutMutex.Unlock()
	fake.StreamOutStub = nil
	if fake.streamOutReturnsOnCall == nil {
		fake.streamOutReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.streamOutReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeVersionedSource) Version() atc.Version {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct {
	}{})
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.versionReturns
	return fakeReturns.result1
}

func (fake *FakeVersionedSource) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeVersionedSource) VersionCalls(stub func() atc.Version) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = stub
}

func (fake *FakeVersionedSource) VersionReturns(result1 atc.Version) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeVersionedSource) VersionReturnsOnCall(i int, result1 atc.Version) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 atc.Version
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeVersionedSource) Volume() worker.Volume {
	fake.volumeMutex.Lock()
	ret, specificReturn := fake.volumeReturnsOnCall[len(fake.volumeArgsForCall)]
	fake.volumeArgsForCall = append(fake.volumeArgsForCall, struct {
	}{})
	fake.recordInvocation("Volume", []interface{}{})
	fake.volumeMutex.Unlock()
	if fake.VolumeStub != nil {
		return fake.VolumeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.volumeReturns
	return fakeReturns.result1
}

func (fake *FakeVersionedSource) VolumeCallCount() int {
	fake.volumeMutex.RLock()
	defer fake.volumeMutex.RUnlock()
	return len(fake.volumeArgsForCall)
}

func (fake *FakeVersionedSource) VolumeCalls(stub func() worker.Volume) {
	fake.volumeMutex.Lock()
	defer fake.volumeMutex.Unlock()
	fake.VolumeStub = stub
}

func (fake *FakeVersionedSource) VolumeReturns(result1 worker.Volume) {
	fake.volumeMutex.Lock()
	defer fake.volumeMutex.Unlock()
	fake.VolumeStub = nil
	fake.volumeReturns = struct {
		result1 worker.Volume
	}{result1}
}

func (fake *FakeVersionedSource) VolumeReturnsOnCall(i int, result1 worker.Volume) {
	fake.volumeMutex.Lock()
	defer fake.volumeMutex.Unlock()
	fake.VolumeStub = nil
	if fake.volumeReturnsOnCall == nil {
		fake.volumeReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
		})
	}
	fake.volumeReturnsOnCall[i] = struct {
		result1 worker.Volume
	}{result1}
}

func (fake *FakeVersionedSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	fake.volumeMutex.RLock()
	defer fake.volumeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVersionedSource) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ resource.VersionedSource = new(FakeVersionedSource)
