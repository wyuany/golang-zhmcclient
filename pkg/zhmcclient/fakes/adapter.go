// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
)

type AdapterAPI struct {
	CreateHipersocketStub        func(string, *zhmcclient.HypersocketPayload) (string, error)
	createHipersocketMutex       sync.RWMutex
	createHipersocketArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.HypersocketPayload
	}
	createHipersocketReturns struct {
		result1 string
		result2 error
	}
	createHipersocketReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	DeleteHipersocketStub        func(string) error
	deleteHipersocketMutex       sync.RWMutex
	deleteHipersocketArgsForCall []struct {
		arg1 string
	}
	deleteHipersocketReturns struct {
		result1 error
	}
	deleteHipersocketReturnsOnCall map[int]struct {
		result1 error
	}
	ListAdaptersStub        func(string, map[string]string) ([]zhmcclient.Adapter, error)
	listAdaptersMutex       sync.RWMutex
	listAdaptersArgsForCall []struct {
		arg1 string
		arg2 map[string]string
	}
	listAdaptersReturns struct {
		result1 []zhmcclient.Adapter
		result2 error
	}
	listAdaptersReturnsOnCall map[int]struct {
		result1 []zhmcclient.Adapter
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *AdapterAPI) CreateHipersocket(arg1 string, arg2 *zhmcclient.HypersocketPayload) (string, error) {
	fake.createHipersocketMutex.Lock()
	ret, specificReturn := fake.createHipersocketReturnsOnCall[len(fake.createHipersocketArgsForCall)]
	fake.createHipersocketArgsForCall = append(fake.createHipersocketArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.HypersocketPayload
	}{arg1, arg2})
	stub := fake.CreateHipersocketStub
	fakeReturns := fake.createHipersocketReturns
	fake.recordInvocation("CreateHipersocket", []interface{}{arg1, arg2})
	fake.createHipersocketMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AdapterAPI) CreateHipersocketCallCount() int {
	fake.createHipersocketMutex.RLock()
	defer fake.createHipersocketMutex.RUnlock()
	return len(fake.createHipersocketArgsForCall)
}

func (fake *AdapterAPI) CreateHipersocketCalls(stub func(string, *zhmcclient.HypersocketPayload) (string, error)) {
	fake.createHipersocketMutex.Lock()
	defer fake.createHipersocketMutex.Unlock()
	fake.CreateHipersocketStub = stub
}

func (fake *AdapterAPI) CreateHipersocketArgsForCall(i int) (string, *zhmcclient.HypersocketPayload) {
	fake.createHipersocketMutex.RLock()
	defer fake.createHipersocketMutex.RUnlock()
	argsForCall := fake.createHipersocketArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *AdapterAPI) CreateHipersocketReturns(result1 string, result2 error) {
	fake.createHipersocketMutex.Lock()
	defer fake.createHipersocketMutex.Unlock()
	fake.CreateHipersocketStub = nil
	fake.createHipersocketReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *AdapterAPI) CreateHipersocketReturnsOnCall(i int, result1 string, result2 error) {
	fake.createHipersocketMutex.Lock()
	defer fake.createHipersocketMutex.Unlock()
	fake.CreateHipersocketStub = nil
	if fake.createHipersocketReturnsOnCall == nil {
		fake.createHipersocketReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.createHipersocketReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *AdapterAPI) DeleteHipersocket(arg1 string) error {
	fake.deleteHipersocketMutex.Lock()
	ret, specificReturn := fake.deleteHipersocketReturnsOnCall[len(fake.deleteHipersocketArgsForCall)]
	fake.deleteHipersocketArgsForCall = append(fake.deleteHipersocketArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteHipersocketStub
	fakeReturns := fake.deleteHipersocketReturns
	fake.recordInvocation("DeleteHipersocket", []interface{}{arg1})
	fake.deleteHipersocketMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *AdapterAPI) DeleteHipersocketCallCount() int {
	fake.deleteHipersocketMutex.RLock()
	defer fake.deleteHipersocketMutex.RUnlock()
	return len(fake.deleteHipersocketArgsForCall)
}

func (fake *AdapterAPI) DeleteHipersocketCalls(stub func(string) error) {
	fake.deleteHipersocketMutex.Lock()
	defer fake.deleteHipersocketMutex.Unlock()
	fake.DeleteHipersocketStub = stub
}

func (fake *AdapterAPI) DeleteHipersocketArgsForCall(i int) string {
	fake.deleteHipersocketMutex.RLock()
	defer fake.deleteHipersocketMutex.RUnlock()
	argsForCall := fake.deleteHipersocketArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AdapterAPI) DeleteHipersocketReturns(result1 error) {
	fake.deleteHipersocketMutex.Lock()
	defer fake.deleteHipersocketMutex.Unlock()
	fake.DeleteHipersocketStub = nil
	fake.deleteHipersocketReturns = struct {
		result1 error
	}{result1}
}

func (fake *AdapterAPI) DeleteHipersocketReturnsOnCall(i int, result1 error) {
	fake.deleteHipersocketMutex.Lock()
	defer fake.deleteHipersocketMutex.Unlock()
	fake.DeleteHipersocketStub = nil
	if fake.deleteHipersocketReturnsOnCall == nil {
		fake.deleteHipersocketReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteHipersocketReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *AdapterAPI) ListAdapters(arg1 string, arg2 map[string]string) ([]zhmcclient.Adapter, error) {
	fake.listAdaptersMutex.Lock()
	ret, specificReturn := fake.listAdaptersReturnsOnCall[len(fake.listAdaptersArgsForCall)]
	fake.listAdaptersArgsForCall = append(fake.listAdaptersArgsForCall, struct {
		arg1 string
		arg2 map[string]string
	}{arg1, arg2})
	stub := fake.ListAdaptersStub
	fakeReturns := fake.listAdaptersReturns
	fake.recordInvocation("ListAdapters", []interface{}{arg1, arg2})
	fake.listAdaptersMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AdapterAPI) ListAdaptersCallCount() int {
	fake.listAdaptersMutex.RLock()
	defer fake.listAdaptersMutex.RUnlock()
	return len(fake.listAdaptersArgsForCall)
}

func (fake *AdapterAPI) ListAdaptersCalls(stub func(string, map[string]string) ([]zhmcclient.Adapter, error)) {
	fake.listAdaptersMutex.Lock()
	defer fake.listAdaptersMutex.Unlock()
	fake.ListAdaptersStub = stub
}

func (fake *AdapterAPI) ListAdaptersArgsForCall(i int) (string, map[string]string) {
	fake.listAdaptersMutex.RLock()
	defer fake.listAdaptersMutex.RUnlock()
	argsForCall := fake.listAdaptersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *AdapterAPI) ListAdaptersReturns(result1 []zhmcclient.Adapter, result2 error) {
	fake.listAdaptersMutex.Lock()
	defer fake.listAdaptersMutex.Unlock()
	fake.ListAdaptersStub = nil
	fake.listAdaptersReturns = struct {
		result1 []zhmcclient.Adapter
		result2 error
	}{result1, result2}
}

func (fake *AdapterAPI) ListAdaptersReturnsOnCall(i int, result1 []zhmcclient.Adapter, result2 error) {
	fake.listAdaptersMutex.Lock()
	defer fake.listAdaptersMutex.Unlock()
	fake.ListAdaptersStub = nil
	if fake.listAdaptersReturnsOnCall == nil {
		fake.listAdaptersReturnsOnCall = make(map[int]struct {
			result1 []zhmcclient.Adapter
			result2 error
		})
	}
	fake.listAdaptersReturnsOnCall[i] = struct {
		result1 []zhmcclient.Adapter
		result2 error
	}{result1, result2}
}

func (fake *AdapterAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createHipersocketMutex.RLock()
	defer fake.createHipersocketMutex.RUnlock()
	fake.deleteHipersocketMutex.RLock()
	defer fake.deleteHipersocketMutex.RUnlock()
	fake.listAdaptersMutex.RLock()
	defer fake.listAdaptersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *AdapterAPI) recordInvocation(key string, args []interface{}) {
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

var _ zhmcclient.AdapterAPI = new(AdapterAPI)
