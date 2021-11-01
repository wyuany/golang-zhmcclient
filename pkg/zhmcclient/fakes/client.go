// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"io"
	"net/url"
	"sync"

	"github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
)

type ClientAPI struct {
	ExecuteRequestStub        func(string, *url.URL, interface{}) (int, []byte, error)
	executeRequestMutex       sync.RWMutex
	executeRequestArgsForCall []struct {
		arg1 string
		arg2 *url.URL
		arg3 interface{}
	}
	executeRequestReturns struct {
		result1 int
		result2 []byte
		result3 error
	}
	executeRequestReturnsOnCall map[int]struct {
		result1 int
		result2 []byte
		result3 error
	}
	GetEndpointURLStub        func() *url.URL
	getEndpointURLMutex       sync.RWMutex
	getEndpointURLArgsForCall []struct {
	}
	getEndpointURLReturns struct {
		result1 *url.URL
	}
	getEndpointURLReturnsOnCall map[int]struct {
		result1 *url.URL
	}
	IsLogonStub        func(bool) bool
	isLogonMutex       sync.RWMutex
	isLogonArgsForCall []struct {
		arg1 bool
	}
	isLogonReturns struct {
		result1 bool
	}
	isLogonReturnsOnCall map[int]struct {
		result1 bool
	}
	LogoffStub        func() error
	logoffMutex       sync.RWMutex
	logoffArgsForCall []struct {
	}
	logoffReturns struct {
		result1 error
	}
	logoffReturnsOnCall map[int]struct {
		result1 error
	}
	LogonStub        func() error
	logonMutex       sync.RWMutex
	logonArgsForCall []struct {
	}
	logonReturns struct {
		result1 error
	}
	logonReturnsOnCall map[int]struct {
		result1 error
	}
	SetCertVerifyStub        func(bool)
	setCertVerifyMutex       sync.RWMutex
	setCertVerifyArgsForCall []struct {
		arg1 bool
	}
	TraceOffStub        func()
	traceOffMutex       sync.RWMutex
	traceOffArgsForCall []struct {
	}
	TraceOnStub        func(io.Writer)
	traceOnMutex       sync.RWMutex
	traceOnArgsForCall []struct {
		arg1 io.Writer
	}
	UploadRequestStub        func(string, *url.URL, []byte) (int, []byte, error)
	uploadRequestMutex       sync.RWMutex
	uploadRequestArgsForCall []struct {
		arg1 string
		arg2 *url.URL
		arg3 []byte
	}
	uploadRequestReturns struct {
		result1 int
		result2 []byte
		result3 error
	}
	uploadRequestReturnsOnCall map[int]struct {
		result1 int
		result2 []byte
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ClientAPI) ExecuteRequest(arg1 string, arg2 *url.URL, arg3 interface{}) (int, []byte, error) {
	fake.executeRequestMutex.Lock()
	ret, specificReturn := fake.executeRequestReturnsOnCall[len(fake.executeRequestArgsForCall)]
	fake.executeRequestArgsForCall = append(fake.executeRequestArgsForCall, struct {
		arg1 string
		arg2 *url.URL
		arg3 interface{}
	}{arg1, arg2, arg3})
	stub := fake.ExecuteRequestStub
	fakeReturns := fake.executeRequestReturns
	fake.recordInvocation("ExecuteRequest", []interface{}{arg1, arg2, arg3})
	fake.executeRequestMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ClientAPI) ExecuteRequestCallCount() int {
	fake.executeRequestMutex.RLock()
	defer fake.executeRequestMutex.RUnlock()
	return len(fake.executeRequestArgsForCall)
}

func (fake *ClientAPI) ExecuteRequestCalls(stub func(string, *url.URL, interface{}) (int, []byte, error)) {
	fake.executeRequestMutex.Lock()
	defer fake.executeRequestMutex.Unlock()
	fake.ExecuteRequestStub = stub
}

func (fake *ClientAPI) ExecuteRequestArgsForCall(i int) (string, *url.URL, interface{}) {
	fake.executeRequestMutex.RLock()
	defer fake.executeRequestMutex.RUnlock()
	argsForCall := fake.executeRequestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ClientAPI) ExecuteRequestReturns(result1 int, result2 []byte, result3 error) {
	fake.executeRequestMutex.Lock()
	defer fake.executeRequestMutex.Unlock()
	fake.ExecuteRequestStub = nil
	fake.executeRequestReturns = struct {
		result1 int
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *ClientAPI) ExecuteRequestReturnsOnCall(i int, result1 int, result2 []byte, result3 error) {
	fake.executeRequestMutex.Lock()
	defer fake.executeRequestMutex.Unlock()
	fake.ExecuteRequestStub = nil
	if fake.executeRequestReturnsOnCall == nil {
		fake.executeRequestReturnsOnCall = make(map[int]struct {
			result1 int
			result2 []byte
			result3 error
		})
	}
	fake.executeRequestReturnsOnCall[i] = struct {
		result1 int
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *ClientAPI) GetEndpointURL() *url.URL {
	fake.getEndpointURLMutex.Lock()
	ret, specificReturn := fake.getEndpointURLReturnsOnCall[len(fake.getEndpointURLArgsForCall)]
	fake.getEndpointURLArgsForCall = append(fake.getEndpointURLArgsForCall, struct {
	}{})
	stub := fake.GetEndpointURLStub
	fakeReturns := fake.getEndpointURLReturns
	fake.recordInvocation("GetEndpointURL", []interface{}{})
	fake.getEndpointURLMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ClientAPI) GetEndpointURLCallCount() int {
	fake.getEndpointURLMutex.RLock()
	defer fake.getEndpointURLMutex.RUnlock()
	return len(fake.getEndpointURLArgsForCall)
}

func (fake *ClientAPI) GetEndpointURLCalls(stub func() *url.URL) {
	fake.getEndpointURLMutex.Lock()
	defer fake.getEndpointURLMutex.Unlock()
	fake.GetEndpointURLStub = stub
}

func (fake *ClientAPI) GetEndpointURLReturns(result1 *url.URL) {
	fake.getEndpointURLMutex.Lock()
	defer fake.getEndpointURLMutex.Unlock()
	fake.GetEndpointURLStub = nil
	fake.getEndpointURLReturns = struct {
		result1 *url.URL
	}{result1}
}

func (fake *ClientAPI) GetEndpointURLReturnsOnCall(i int, result1 *url.URL) {
	fake.getEndpointURLMutex.Lock()
	defer fake.getEndpointURLMutex.Unlock()
	fake.GetEndpointURLStub = nil
	if fake.getEndpointURLReturnsOnCall == nil {
		fake.getEndpointURLReturnsOnCall = make(map[int]struct {
			result1 *url.URL
		})
	}
	fake.getEndpointURLReturnsOnCall[i] = struct {
		result1 *url.URL
	}{result1}
}

func (fake *ClientAPI) IsLogon(arg1 bool) bool {
	fake.isLogonMutex.Lock()
	ret, specificReturn := fake.isLogonReturnsOnCall[len(fake.isLogonArgsForCall)]
	fake.isLogonArgsForCall = append(fake.isLogonArgsForCall, struct {
		arg1 bool
	}{arg1})
	stub := fake.IsLogonStub
	fakeReturns := fake.isLogonReturns
	fake.recordInvocation("IsLogon", []interface{}{arg1})
	fake.isLogonMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ClientAPI) IsLogonCallCount() int {
	fake.isLogonMutex.RLock()
	defer fake.isLogonMutex.RUnlock()
	return len(fake.isLogonArgsForCall)
}

func (fake *ClientAPI) IsLogonCalls(stub func(bool) bool) {
	fake.isLogonMutex.Lock()
	defer fake.isLogonMutex.Unlock()
	fake.IsLogonStub = stub
}

func (fake *ClientAPI) IsLogonArgsForCall(i int) bool {
	fake.isLogonMutex.RLock()
	defer fake.isLogonMutex.RUnlock()
	argsForCall := fake.isLogonArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ClientAPI) IsLogonReturns(result1 bool) {
	fake.isLogonMutex.Lock()
	defer fake.isLogonMutex.Unlock()
	fake.IsLogonStub = nil
	fake.isLogonReturns = struct {
		result1 bool
	}{result1}
}

func (fake *ClientAPI) IsLogonReturnsOnCall(i int, result1 bool) {
	fake.isLogonMutex.Lock()
	defer fake.isLogonMutex.Unlock()
	fake.IsLogonStub = nil
	if fake.isLogonReturnsOnCall == nil {
		fake.isLogonReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isLogonReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *ClientAPI) Logoff() error {
	fake.logoffMutex.Lock()
	ret, specificReturn := fake.logoffReturnsOnCall[len(fake.logoffArgsForCall)]
	fake.logoffArgsForCall = append(fake.logoffArgsForCall, struct {
	}{})
	stub := fake.LogoffStub
	fakeReturns := fake.logoffReturns
	fake.recordInvocation("Logoff", []interface{}{})
	fake.logoffMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ClientAPI) LogoffCallCount() int {
	fake.logoffMutex.RLock()
	defer fake.logoffMutex.RUnlock()
	return len(fake.logoffArgsForCall)
}

func (fake *ClientAPI) LogoffCalls(stub func() error) {
	fake.logoffMutex.Lock()
	defer fake.logoffMutex.Unlock()
	fake.LogoffStub = stub
}

func (fake *ClientAPI) LogoffReturns(result1 error) {
	fake.logoffMutex.Lock()
	defer fake.logoffMutex.Unlock()
	fake.LogoffStub = nil
	fake.logoffReturns = struct {
		result1 error
	}{result1}
}

func (fake *ClientAPI) LogoffReturnsOnCall(i int, result1 error) {
	fake.logoffMutex.Lock()
	defer fake.logoffMutex.Unlock()
	fake.LogoffStub = nil
	if fake.logoffReturnsOnCall == nil {
		fake.logoffReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.logoffReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ClientAPI) Logon() error {
	fake.logonMutex.Lock()
	ret, specificReturn := fake.logonReturnsOnCall[len(fake.logonArgsForCall)]
	fake.logonArgsForCall = append(fake.logonArgsForCall, struct {
	}{})
	stub := fake.LogonStub
	fakeReturns := fake.logonReturns
	fake.recordInvocation("Logon", []interface{}{})
	fake.logonMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ClientAPI) LogonCallCount() int {
	fake.logonMutex.RLock()
	defer fake.logonMutex.RUnlock()
	return len(fake.logonArgsForCall)
}

func (fake *ClientAPI) LogonCalls(stub func() error) {
	fake.logonMutex.Lock()
	defer fake.logonMutex.Unlock()
	fake.LogonStub = stub
}

func (fake *ClientAPI) LogonReturns(result1 error) {
	fake.logonMutex.Lock()
	defer fake.logonMutex.Unlock()
	fake.LogonStub = nil
	fake.logonReturns = struct {
		result1 error
	}{result1}
}

func (fake *ClientAPI) LogonReturnsOnCall(i int, result1 error) {
	fake.logonMutex.Lock()
	defer fake.logonMutex.Unlock()
	fake.LogonStub = nil
	if fake.logonReturnsOnCall == nil {
		fake.logonReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.logonReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ClientAPI) SetCertVerify(arg1 bool) {
	fake.setCertVerifyMutex.Lock()
	fake.setCertVerifyArgsForCall = append(fake.setCertVerifyArgsForCall, struct {
		arg1 bool
	}{arg1})
	stub := fake.SetCertVerifyStub
	fake.recordInvocation("SetCertVerify", []interface{}{arg1})
	fake.setCertVerifyMutex.Unlock()
	if stub != nil {
		fake.SetCertVerifyStub(arg1)
	}
}

func (fake *ClientAPI) SetCertVerifyCallCount() int {
	fake.setCertVerifyMutex.RLock()
	defer fake.setCertVerifyMutex.RUnlock()
	return len(fake.setCertVerifyArgsForCall)
}

func (fake *ClientAPI) SetCertVerifyCalls(stub func(bool)) {
	fake.setCertVerifyMutex.Lock()
	defer fake.setCertVerifyMutex.Unlock()
	fake.SetCertVerifyStub = stub
}

func (fake *ClientAPI) SetCertVerifyArgsForCall(i int) bool {
	fake.setCertVerifyMutex.RLock()
	defer fake.setCertVerifyMutex.RUnlock()
	argsForCall := fake.setCertVerifyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ClientAPI) TraceOff() {
	fake.traceOffMutex.Lock()
	fake.traceOffArgsForCall = append(fake.traceOffArgsForCall, struct {
	}{})
	stub := fake.TraceOffStub
	fake.recordInvocation("TraceOff", []interface{}{})
	fake.traceOffMutex.Unlock()
	if stub != nil {
		fake.TraceOffStub()
	}
}

func (fake *ClientAPI) TraceOffCallCount() int {
	fake.traceOffMutex.RLock()
	defer fake.traceOffMutex.RUnlock()
	return len(fake.traceOffArgsForCall)
}

func (fake *ClientAPI) TraceOffCalls(stub func()) {
	fake.traceOffMutex.Lock()
	defer fake.traceOffMutex.Unlock()
	fake.TraceOffStub = stub
}

func (fake *ClientAPI) TraceOn(arg1 io.Writer) {
	fake.traceOnMutex.Lock()
	fake.traceOnArgsForCall = append(fake.traceOnArgsForCall, struct {
		arg1 io.Writer
	}{arg1})
	stub := fake.TraceOnStub
	fake.recordInvocation("TraceOn", []interface{}{arg1})
	fake.traceOnMutex.Unlock()
	if stub != nil {
		fake.TraceOnStub(arg1)
	}
}

func (fake *ClientAPI) TraceOnCallCount() int {
	fake.traceOnMutex.RLock()
	defer fake.traceOnMutex.RUnlock()
	return len(fake.traceOnArgsForCall)
}

func (fake *ClientAPI) TraceOnCalls(stub func(io.Writer)) {
	fake.traceOnMutex.Lock()
	defer fake.traceOnMutex.Unlock()
	fake.TraceOnStub = stub
}

func (fake *ClientAPI) TraceOnArgsForCall(i int) io.Writer {
	fake.traceOnMutex.RLock()
	defer fake.traceOnMutex.RUnlock()
	argsForCall := fake.traceOnArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ClientAPI) UploadRequest(arg1 string, arg2 *url.URL, arg3 []byte) (int, []byte, error) {
	var arg3Copy []byte
	if arg3 != nil {
		arg3Copy = make([]byte, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.uploadRequestMutex.Lock()
	ret, specificReturn := fake.uploadRequestReturnsOnCall[len(fake.uploadRequestArgsForCall)]
	fake.uploadRequestArgsForCall = append(fake.uploadRequestArgsForCall, struct {
		arg1 string
		arg2 *url.URL
		arg3 []byte
	}{arg1, arg2, arg3Copy})
	stub := fake.UploadRequestStub
	fakeReturns := fake.uploadRequestReturns
	fake.recordInvocation("UploadRequest", []interface{}{arg1, arg2, arg3Copy})
	fake.uploadRequestMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ClientAPI) UploadRequestCallCount() int {
	fake.uploadRequestMutex.RLock()
	defer fake.uploadRequestMutex.RUnlock()
	return len(fake.uploadRequestArgsForCall)
}

func (fake *ClientAPI) UploadRequestCalls(stub func(string, *url.URL, []byte) (int, []byte, error)) {
	fake.uploadRequestMutex.Lock()
	defer fake.uploadRequestMutex.Unlock()
	fake.UploadRequestStub = stub
}

func (fake *ClientAPI) UploadRequestArgsForCall(i int) (string, *url.URL, []byte) {
	fake.uploadRequestMutex.RLock()
	defer fake.uploadRequestMutex.RUnlock()
	argsForCall := fake.uploadRequestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ClientAPI) UploadRequestReturns(result1 int, result2 []byte, result3 error) {
	fake.uploadRequestMutex.Lock()
	defer fake.uploadRequestMutex.Unlock()
	fake.UploadRequestStub = nil
	fake.uploadRequestReturns = struct {
		result1 int
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *ClientAPI) UploadRequestReturnsOnCall(i int, result1 int, result2 []byte, result3 error) {
	fake.uploadRequestMutex.Lock()
	defer fake.uploadRequestMutex.Unlock()
	fake.UploadRequestStub = nil
	if fake.uploadRequestReturnsOnCall == nil {
		fake.uploadRequestReturnsOnCall = make(map[int]struct {
			result1 int
			result2 []byte
			result3 error
		})
	}
	fake.uploadRequestReturnsOnCall[i] = struct {
		result1 int
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *ClientAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeRequestMutex.RLock()
	defer fake.executeRequestMutex.RUnlock()
	fake.getEndpointURLMutex.RLock()
	defer fake.getEndpointURLMutex.RUnlock()
	fake.isLogonMutex.RLock()
	defer fake.isLogonMutex.RUnlock()
	fake.logoffMutex.RLock()
	defer fake.logoffMutex.RUnlock()
	fake.logonMutex.RLock()
	defer fake.logonMutex.RUnlock()
	fake.setCertVerifyMutex.RLock()
	defer fake.setCertVerifyMutex.RUnlock()
	fake.traceOffMutex.RLock()
	defer fake.traceOffMutex.RUnlock()
	fake.traceOnMutex.RLock()
	defer fake.traceOnMutex.RUnlock()
	fake.uploadRequestMutex.RLock()
	defer fake.uploadRequestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ClientAPI) recordInvocation(key string, args []interface{}) {
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

var _ zhmcclient.ClientAPI = new(ClientAPI)
