// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
)

type CpcAPI struct {
	CreateAdapterStub        func(string, *zhmcclient.Adapter) (*zhmcclient.Adapter, error)
	createAdapterMutex       sync.RWMutex
	createAdapterArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.Adapter
	}
	createAdapterReturns struct {
		result1 *zhmcclient.Adapter
		result2 error
	}
	createAdapterReturnsOnCall map[int]struct {
		result1 *zhmcclient.Adapter
		result2 error
	}
	DeleteAdapterStub        func(string) error
	deleteAdapterMutex       sync.RWMutex
	deleteAdapterArgsForCall []struct {
		arg1 string
	}
	deleteAdapterReturns struct {
		result1 error
	}
	deleteAdapterReturnsOnCall map[int]struct {
		result1 error
	}
	ListAdaptersStub        func(string) ([]zhmcclient.Adapter, error)
	listAdaptersMutex       sync.RWMutex
	listAdaptersArgsForCall []struct {
		arg1 string
	}
	listAdaptersReturns struct {
		result1 []zhmcclient.Adapter
		result2 error
	}
	listAdaptersReturnsOnCall map[int]struct {
		result1 []zhmcclient.Adapter
		result2 error
	}
	ListCPCsStub        func() ([]zhmcclient.CPC, error)
	listCPCsMutex       sync.RWMutex
	listCPCsArgsForCall []struct {
	}
	listCPCsReturns struct {
		result1 []zhmcclient.CPC
		result2 error
	}
	listCPCsReturnsOnCall map[int]struct {
		result1 []zhmcclient.CPC
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CpcAPI) CreateAdapter(arg1 string, arg2 *zhmcclient.Adapter) (*zhmcclient.Adapter, error) {
	fake.createAdapterMutex.Lock()
	ret, specificReturn := fake.createAdapterReturnsOnCall[len(fake.createAdapterArgsForCall)]
	fake.createAdapterArgsForCall = append(fake.createAdapterArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.Adapter
	}{arg1, arg2})
	stub := fake.CreateAdapterStub
	fakeReturns := fake.createAdapterReturns
	fake.recordInvocation("CreateAdapter", []interface{}{arg1, arg2})
	fake.createAdapterMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CpcAPI) CreateAdapterCallCount() int {
	fake.createAdapterMutex.RLock()
	defer fake.createAdapterMutex.RUnlock()
	return len(fake.createAdapterArgsForCall)
}

func (fake *CpcAPI) CreateAdapterCalls(stub func(string, *zhmcclient.Adapter) (*zhmcclient.Adapter, error)) {
	fake.createAdapterMutex.Lock()
	defer fake.createAdapterMutex.Unlock()
	fake.CreateAdapterStub = stub
}

func (fake *CpcAPI) CreateAdapterArgsForCall(i int) (string, *zhmcclient.Adapter) {
	fake.createAdapterMutex.RLock()
	defer fake.createAdapterMutex.RUnlock()
	argsForCall := fake.createAdapterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CpcAPI) CreateAdapterReturns(result1 *zhmcclient.Adapter, result2 error) {
	fake.createAdapterMutex.Lock()
	defer fake.createAdapterMutex.Unlock()
	fake.CreateAdapterStub = nil
	fake.createAdapterReturns = struct {
		result1 *zhmcclient.Adapter
		result2 error
	}{result1, result2}
}

func (fake *CpcAPI) CreateAdapterReturnsOnCall(i int, result1 *zhmcclient.Adapter, result2 error) {
	fake.createAdapterMutex.Lock()
	defer fake.createAdapterMutex.Unlock()
	fake.CreateAdapterStub = nil
	if fake.createAdapterReturnsOnCall == nil {
		fake.createAdapterReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.Adapter
			result2 error
		})
	}
	fake.createAdapterReturnsOnCall[i] = struct {
		result1 *zhmcclient.Adapter
		result2 error
	}{result1, result2}
}

func (fake *CpcAPI) DeleteAdapter(arg1 string) error {
	fake.deleteAdapterMutex.Lock()
	ret, specificReturn := fake.deleteAdapterReturnsOnCall[len(fake.deleteAdapterArgsForCall)]
	fake.deleteAdapterArgsForCall = append(fake.deleteAdapterArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteAdapterStub
	fakeReturns := fake.deleteAdapterReturns
	fake.recordInvocation("DeleteAdapter", []interface{}{arg1})
	fake.deleteAdapterMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *CpcAPI) DeleteAdapterCallCount() int {
	fake.deleteAdapterMutex.RLock()
	defer fake.deleteAdapterMutex.RUnlock()
	return len(fake.deleteAdapterArgsForCall)
}

func (fake *CpcAPI) DeleteAdapterCalls(stub func(string) error) {
	fake.deleteAdapterMutex.Lock()
	defer fake.deleteAdapterMutex.Unlock()
	fake.DeleteAdapterStub = stub
}

func (fake *CpcAPI) DeleteAdapterArgsForCall(i int) string {
	fake.deleteAdapterMutex.RLock()
	defer fake.deleteAdapterMutex.RUnlock()
	argsForCall := fake.deleteAdapterArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CpcAPI) DeleteAdapterReturns(result1 error) {
	fake.deleteAdapterMutex.Lock()
	defer fake.deleteAdapterMutex.Unlock()
	fake.DeleteAdapterStub = nil
	fake.deleteAdapterReturns = struct {
		result1 error
	}{result1}
}

func (fake *CpcAPI) DeleteAdapterReturnsOnCall(i int, result1 error) {
	fake.deleteAdapterMutex.Lock()
	defer fake.deleteAdapterMutex.Unlock()
	fake.DeleteAdapterStub = nil
	if fake.deleteAdapterReturnsOnCall == nil {
		fake.deleteAdapterReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteAdapterReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *CpcAPI) ListAdapters(arg1 string) ([]zhmcclient.Adapter, error) {
	fake.listAdaptersMutex.Lock()
	ret, specificReturn := fake.listAdaptersReturnsOnCall[len(fake.listAdaptersArgsForCall)]
	fake.listAdaptersArgsForCall = append(fake.listAdaptersArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ListAdaptersStub
	fakeReturns := fake.listAdaptersReturns
	fake.recordInvocation("ListAdapters", []interface{}{arg1})
	fake.listAdaptersMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CpcAPI) ListAdaptersCallCount() int {
	fake.listAdaptersMutex.RLock()
	defer fake.listAdaptersMutex.RUnlock()
	return len(fake.listAdaptersArgsForCall)
}

func (fake *CpcAPI) ListAdaptersCalls(stub func(string) ([]zhmcclient.Adapter, error)) {
	fake.listAdaptersMutex.Lock()
	defer fake.listAdaptersMutex.Unlock()
	fake.ListAdaptersStub = stub
}

func (fake *CpcAPI) ListAdaptersArgsForCall(i int) string {
	fake.listAdaptersMutex.RLock()
	defer fake.listAdaptersMutex.RUnlock()
	argsForCall := fake.listAdaptersArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CpcAPI) ListAdaptersReturns(result1 []zhmcclient.Adapter, result2 error) {
	fake.listAdaptersMutex.Lock()
	defer fake.listAdaptersMutex.Unlock()
	fake.ListAdaptersStub = nil
	fake.listAdaptersReturns = struct {
		result1 []zhmcclient.Adapter
		result2 error
	}{result1, result2}
}

func (fake *CpcAPI) ListAdaptersReturnsOnCall(i int, result1 []zhmcclient.Adapter, result2 error) {
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

func (fake *CpcAPI) ListCPCs() ([]zhmcclient.CPC, error) {
	fake.listCPCsMutex.Lock()
	ret, specificReturn := fake.listCPCsReturnsOnCall[len(fake.listCPCsArgsForCall)]
	fake.listCPCsArgsForCall = append(fake.listCPCsArgsForCall, struct {
	}{})
	stub := fake.ListCPCsStub
	fakeReturns := fake.listCPCsReturns
	fake.recordInvocation("ListCPCs", []interface{}{})
	fake.listCPCsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CpcAPI) ListCPCsCallCount() int {
	fake.listCPCsMutex.RLock()
	defer fake.listCPCsMutex.RUnlock()
	return len(fake.listCPCsArgsForCall)
}

func (fake *CpcAPI) ListCPCsCalls(stub func() ([]zhmcclient.CPC, error)) {
	fake.listCPCsMutex.Lock()
	defer fake.listCPCsMutex.Unlock()
	fake.ListCPCsStub = stub
}

func (fake *CpcAPI) ListCPCsReturns(result1 []zhmcclient.CPC, result2 error) {
	fake.listCPCsMutex.Lock()
	defer fake.listCPCsMutex.Unlock()
	fake.ListCPCsStub = nil
	fake.listCPCsReturns = struct {
		result1 []zhmcclient.CPC
		result2 error
	}{result1, result2}
}

func (fake *CpcAPI) ListCPCsReturnsOnCall(i int, result1 []zhmcclient.CPC, result2 error) {
	fake.listCPCsMutex.Lock()
	defer fake.listCPCsMutex.Unlock()
	fake.ListCPCsStub = nil
	if fake.listCPCsReturnsOnCall == nil {
		fake.listCPCsReturnsOnCall = make(map[int]struct {
			result1 []zhmcclient.CPC
			result2 error
		})
	}
	fake.listCPCsReturnsOnCall[i] = struct {
		result1 []zhmcclient.CPC
		result2 error
	}{result1, result2}
}

func (fake *CpcAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createAdapterMutex.RLock()
	defer fake.createAdapterMutex.RUnlock()
	fake.deleteAdapterMutex.RLock()
	defer fake.deleteAdapterMutex.RUnlock()
	fake.listAdaptersMutex.RLock()
	defer fake.listAdaptersMutex.RUnlock()
	fake.listCPCsMutex.RLock()
	defer fake.listCPCsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CpcAPI) recordInvocation(key string, args []interface{}) {
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

var _ zhmcclient.CpcAPI = new(CpcAPI)
