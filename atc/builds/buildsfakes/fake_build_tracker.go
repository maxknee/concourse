// Code generated by counterfeiter. DO NOT EDIT.
package buildsfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/builds"
)

type FakeBuildTracker struct {
	ReleaseStub        func()
	releaseMutex       sync.RWMutex
	releaseArgsForCall []struct {
	}
	TrackStub        func()
	trackMutex       sync.RWMutex
	trackArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuildTracker) Release() {
	fake.releaseMutex.Lock()
	fake.releaseArgsForCall = append(fake.releaseArgsForCall, struct {
	}{})
	fake.recordInvocation("Release", []interface{}{})
	fake.releaseMutex.Unlock()
	if fake.ReleaseStub != nil {
		fake.ReleaseStub()
	}
}

func (fake *FakeBuildTracker) ReleaseCallCount() int {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return len(fake.releaseArgsForCall)
}

func (fake *FakeBuildTracker) ReleaseCalls(stub func()) {
	fake.releaseMutex.Lock()
	defer fake.releaseMutex.Unlock()
	fake.ReleaseStub = stub
}

func (fake *FakeBuildTracker) Track() {
	fake.trackMutex.Lock()
	fake.trackArgsForCall = append(fake.trackArgsForCall, struct {
	}{})
	fake.recordInvocation("Track", []interface{}{})
	fake.trackMutex.Unlock()
	if fake.TrackStub != nil {
		fake.TrackStub()
	}
}

func (fake *FakeBuildTracker) TrackCallCount() int {
	fake.trackMutex.RLock()
	defer fake.trackMutex.RUnlock()
	return len(fake.trackArgsForCall)
}

func (fake *FakeBuildTracker) TrackCalls(stub func()) {
	fake.trackMutex.Lock()
	defer fake.trackMutex.Unlock()
	fake.TrackStub = stub
}

func (fake *FakeBuildTracker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	fake.trackMutex.RLock()
	defer fake.trackMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBuildTracker) recordInvocation(key string, args []interface{}) {
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

var _ builds.BuildTracker = new(FakeBuildTracker)
