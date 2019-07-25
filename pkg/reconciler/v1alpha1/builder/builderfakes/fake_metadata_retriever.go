// Code generated by counterfeiter. DO NOT EDIT.
package builderfakes

import (
	"sync"

	"github.com/pivotal/build-service-beam/pkg/cnb"
	"github.com/pivotal/build-service-beam/pkg/reconciler/v1alpha1/builder"
	"github.com/pivotal/build-service-beam/pkg/registry"
)

type FakeMetadataRetriever struct {
	GetBuilderBuildpacksStub        func(registry.ImageRef) (cnb.BuilderMetadata, error)
	getBuilderBuildpacksMutex       sync.RWMutex
	getBuilderBuildpacksArgsForCall []struct {
		arg1 registry.ImageRef
	}
	getBuilderBuildpacksReturns struct {
		result1 cnb.BuilderMetadata
		result2 error
	}
	getBuilderBuildpacksReturnsOnCall map[int]struct {
		result1 cnb.BuilderMetadata
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMetadataRetriever) GetBuilderBuildpacks(arg1 registry.ImageRef) (cnb.BuilderMetadata, error) {
	fake.getBuilderBuildpacksMutex.Lock()
	ret, specificReturn := fake.getBuilderBuildpacksReturnsOnCall[len(fake.getBuilderBuildpacksArgsForCall)]
	fake.getBuilderBuildpacksArgsForCall = append(fake.getBuilderBuildpacksArgsForCall, struct {
		arg1 registry.ImageRef
	}{arg1})
	fake.recordInvocation("GetBuilderBuildpacks", []interface{}{arg1})
	fake.getBuilderBuildpacksMutex.Unlock()
	if fake.GetBuilderBuildpacksStub != nil {
		return fake.GetBuilderBuildpacksStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getBuilderBuildpacksReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeMetadataRetriever) GetBuilderBuildpacksCallCount() int {
	fake.getBuilderBuildpacksMutex.RLock()
	defer fake.getBuilderBuildpacksMutex.RUnlock()
	return len(fake.getBuilderBuildpacksArgsForCall)
}

func (fake *FakeMetadataRetriever) GetBuilderBuildpacksCalls(stub func(registry.ImageRef) (cnb.BuilderMetadata, error)) {
	fake.getBuilderBuildpacksMutex.Lock()
	defer fake.getBuilderBuildpacksMutex.Unlock()
	fake.GetBuilderBuildpacksStub = stub
}

func (fake *FakeMetadataRetriever) GetBuilderBuildpacksArgsForCall(i int) registry.ImageRef {
	fake.getBuilderBuildpacksMutex.RLock()
	defer fake.getBuilderBuildpacksMutex.RUnlock()
	argsForCall := fake.getBuilderBuildpacksArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMetadataRetriever) GetBuilderBuildpacksReturns(result1 cnb.BuilderMetadata, result2 error) {
	fake.getBuilderBuildpacksMutex.Lock()
	defer fake.getBuilderBuildpacksMutex.Unlock()
	fake.GetBuilderBuildpacksStub = nil
	fake.getBuilderBuildpacksReturns = struct {
		result1 cnb.BuilderMetadata
		result2 error
	}{result1, result2}
}

func (fake *FakeMetadataRetriever) GetBuilderBuildpacksReturnsOnCall(i int, result1 cnb.BuilderMetadata, result2 error) {
	fake.getBuilderBuildpacksMutex.Lock()
	defer fake.getBuilderBuildpacksMutex.Unlock()
	fake.GetBuilderBuildpacksStub = nil
	if fake.getBuilderBuildpacksReturnsOnCall == nil {
		fake.getBuilderBuildpacksReturnsOnCall = make(map[int]struct {
			result1 cnb.BuilderMetadata
			result2 error
		})
	}
	fake.getBuilderBuildpacksReturnsOnCall[i] = struct {
		result1 cnb.BuilderMetadata
		result2 error
	}{result1, result2}
}

func (fake *FakeMetadataRetriever) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getBuilderBuildpacksMutex.RLock()
	defer fake.getBuilderBuildpacksMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMetadataRetriever) recordInvocation(key string, args []interface{}) {
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

var _ builder.MetadataRetriever = new(FakeMetadataRetriever)
