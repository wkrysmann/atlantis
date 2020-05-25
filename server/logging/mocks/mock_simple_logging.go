// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/logging (interfaces: SimpleLogging)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	logging "github.com/runatlantis/atlantis/server/logging"
	log "log"
	"reflect"
	"time"
)

type MockSimpleLogging struct {
	fail func(message string, callerSkip ...int)
}

func NewMockSimpleLogging(options ...pegomock.Option) *MockSimpleLogging {
	mock := &MockSimpleLogging{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockSimpleLogging) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockSimpleLogging) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockSimpleLogging) Debug(format string, a ...interface{}) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSimpleLogging().")
	}
	params := []pegomock.Param{format}
	for _, param := range a {
		params = append(params, param)
	}
	pegomock.GetGenericMockFrom(mock).Invoke("Debug", params, []reflect.Type{})
}

func (mock *MockSimpleLogging) Info(format string, a ...interface{}) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSimpleLogging().")
	}
	params := []pegomock.Param{format}
	for _, param := range a {
		params = append(params, param)
	}
	pegomock.GetGenericMockFrom(mock).Invoke("Info", params, []reflect.Type{})
}

func (mock *MockSimpleLogging) Warn(format string, a ...interface{}) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSimpleLogging().")
	}
	params := []pegomock.Param{format}
	for _, param := range a {
		params = append(params, param)
	}
	pegomock.GetGenericMockFrom(mock).Invoke("Warn", params, []reflect.Type{})
}

func (mock *MockSimpleLogging) Err(format string, a ...interface{}) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSimpleLogging().")
	}
	params := []pegomock.Param{format}
	for _, param := range a {
		params = append(params, param)
	}
	pegomock.GetGenericMockFrom(mock).Invoke("Err", params, []reflect.Type{})
}

func (mock *MockSimpleLogging) Log(level logging.LogLevel, format string, a ...interface{}) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSimpleLogging().")
	}
	params := []pegomock.Param{level, format}
	for _, param := range a {
		params = append(params, param)
	}
	pegomock.GetGenericMockFrom(mock).Invoke("Log", params, []reflect.Type{})
}

func (mock *MockSimpleLogging) Underlying() *log.Logger {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSimpleLogging().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Underlying", params, []reflect.Type{reflect.TypeOf((**log.Logger)(nil)).Elem()})
	var ret0 *log.Logger
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*log.Logger)
		}
	}
	return ret0
}

func (mock *MockSimpleLogging) GetLevel() logging.LogLevel {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSimpleLogging().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetLevel", params, []reflect.Type{reflect.TypeOf((*logging.LogLevel)(nil)).Elem()})
	var ret0 logging.LogLevel
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(logging.LogLevel)
		}
	}
	return ret0
}

func (mock *MockSimpleLogging) NewLogger(_param0 string, _param1 bool, _param2 logging.LogLevel) *logging.SimpleLogger {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSimpleLogging().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	result := pegomock.GetGenericMockFrom(mock).Invoke("NewLogger", params, []reflect.Type{reflect.TypeOf((**logging.SimpleLogger)(nil)).Elem()})
	var ret0 *logging.SimpleLogger
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*logging.SimpleLogger)
		}
	}
	return ret0
}

func (mock *MockSimpleLogging) VerifyWasCalledOnce() *VerifierMockSimpleLogging {
	return &VerifierMockSimpleLogging{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockSimpleLogging) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockSimpleLogging {
	return &VerifierMockSimpleLogging{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockSimpleLogging) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockSimpleLogging {
	return &VerifierMockSimpleLogging{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockSimpleLogging) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockSimpleLogging {
	return &VerifierMockSimpleLogging{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockSimpleLogging struct {
	mock                   *MockSimpleLogging
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockSimpleLogging) Debug(format string, a ...interface{}) *MockSimpleLogging_Debug_OngoingVerification {
	params := []pegomock.Param{format}
	for _, param := range a {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Debug", params, verifier.timeout)
	return &MockSimpleLogging_Debug_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSimpleLogging_Debug_OngoingVerification struct {
	mock              *MockSimpleLogging
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSimpleLogging_Debug_OngoingVerification) GetCapturedArguments() (string, []interface{}) {
	format, a := c.GetAllCapturedArguments()
	return format[len(format)-1], a[len(a)-1]
}

func (c *MockSimpleLogging_Debug_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 [][]interface{}) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([][]interface{}, len(c.methodInvocations))
		for u := 0; u < len(c.methodInvocations); u++ {
			_param1[u] = make([]interface{}, len(params)-1)
			for x := 1; x < len(params); x++ {
				if params[x][u] != nil {
					_param1[u][x-1] = params[x][u].(interface{})
				}
			}
		}
	}
	return
}

func (verifier *VerifierMockSimpleLogging) Info(format string, a ...interface{}) *MockSimpleLogging_Info_OngoingVerification {
	params := []pegomock.Param{format}
	for _, param := range a {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Info", params, verifier.timeout)
	return &MockSimpleLogging_Info_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSimpleLogging_Info_OngoingVerification struct {
	mock              *MockSimpleLogging
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSimpleLogging_Info_OngoingVerification) GetCapturedArguments() (string, []interface{}) {
	format, a := c.GetAllCapturedArguments()
	return format[len(format)-1], a[len(a)-1]
}

func (c *MockSimpleLogging_Info_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 [][]interface{}) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([][]interface{}, len(c.methodInvocations))
		for u := 0; u < len(c.methodInvocations); u++ {
			_param1[u] = make([]interface{}, len(params)-1)
			for x := 1; x < len(params); x++ {
				if params[x][u] != nil {
					_param1[u][x-1] = params[x][u].(interface{})
				}
			}
		}
	}
	return
}

func (verifier *VerifierMockSimpleLogging) Warn(format string, a ...interface{}) *MockSimpleLogging_Warn_OngoingVerification {
	params := []pegomock.Param{format}
	for _, param := range a {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Warn", params, verifier.timeout)
	return &MockSimpleLogging_Warn_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSimpleLogging_Warn_OngoingVerification struct {
	mock              *MockSimpleLogging
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSimpleLogging_Warn_OngoingVerification) GetCapturedArguments() (string, []interface{}) {
	format, a := c.GetAllCapturedArguments()
	return format[len(format)-1], a[len(a)-1]
}

func (c *MockSimpleLogging_Warn_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 [][]interface{}) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([][]interface{}, len(c.methodInvocations))
		for u := 0; u < len(c.methodInvocations); u++ {
			_param1[u] = make([]interface{}, len(params)-1)
			for x := 1; x < len(params); x++ {
				if params[x][u] != nil {
					_param1[u][x-1] = params[x][u].(interface{})
				}
			}
		}
	}
	return
}

func (verifier *VerifierMockSimpleLogging) Err(format string, a ...interface{}) *MockSimpleLogging_Err_OngoingVerification {
	params := []pegomock.Param{format}
	for _, param := range a {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Err", params, verifier.timeout)
	return &MockSimpleLogging_Err_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSimpleLogging_Err_OngoingVerification struct {
	mock              *MockSimpleLogging
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSimpleLogging_Err_OngoingVerification) GetCapturedArguments() (string, []interface{}) {
	format, a := c.GetAllCapturedArguments()
	return format[len(format)-1], a[len(a)-1]
}

func (c *MockSimpleLogging_Err_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 [][]interface{}) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([][]interface{}, len(c.methodInvocations))
		for u := 0; u < len(c.methodInvocations); u++ {
			_param1[u] = make([]interface{}, len(params)-1)
			for x := 1; x < len(params); x++ {
				if params[x][u] != nil {
					_param1[u][x-1] = params[x][u].(interface{})
				}
			}
		}
	}
	return
}

func (verifier *VerifierMockSimpleLogging) Log(level logging.LogLevel, format string, a ...interface{}) *MockSimpleLogging_Log_OngoingVerification {
	params := []pegomock.Param{level, format}
	for _, param := range a {
		params = append(params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Log", params, verifier.timeout)
	return &MockSimpleLogging_Log_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSimpleLogging_Log_OngoingVerification struct {
	mock              *MockSimpleLogging
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSimpleLogging_Log_OngoingVerification) GetCapturedArguments() (logging.LogLevel, string, []interface{}) {
	level, format, a := c.GetAllCapturedArguments()
	return level[len(level)-1], format[len(format)-1], a[len(a)-1]
}

func (c *MockSimpleLogging_Log_OngoingVerification) GetAllCapturedArguments() (_param0 []logging.LogLevel, _param1 []string, _param2 [][]interface{}) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]logging.LogLevel, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(logging.LogLevel)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([][]interface{}, len(c.methodInvocations))
		for u := 0; u < len(c.methodInvocations); u++ {
			_param2[u] = make([]interface{}, len(params)-2)
			for x := 2; x < len(params); x++ {
				if params[x][u] != nil {
					_param2[u][x-2] = params[x][u].(interface{})
				}
			}
		}
	}
	return
}

func (verifier *VerifierMockSimpleLogging) Underlying() *MockSimpleLogging_Underlying_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Underlying", params, verifier.timeout)
	return &MockSimpleLogging_Underlying_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSimpleLogging_Underlying_OngoingVerification struct {
	mock              *MockSimpleLogging
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSimpleLogging_Underlying_OngoingVerification) GetCapturedArguments() {
}

func (c *MockSimpleLogging_Underlying_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockSimpleLogging) GetLevel() *MockSimpleLogging_GetLevel_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetLevel", params, verifier.timeout)
	return &MockSimpleLogging_GetLevel_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSimpleLogging_GetLevel_OngoingVerification struct {
	mock              *MockSimpleLogging
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSimpleLogging_GetLevel_OngoingVerification) GetCapturedArguments() {
}

func (c *MockSimpleLogging_GetLevel_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockSimpleLogging) NewLogger(_param0 string, _param1 bool, _param2 logging.LogLevel) *MockSimpleLogging_NewLogger_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "NewLogger", params, verifier.timeout)
	return &MockSimpleLogging_NewLogger_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSimpleLogging_NewLogger_OngoingVerification struct {
	mock              *MockSimpleLogging
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSimpleLogging_NewLogger_OngoingVerification) GetCapturedArguments() (string, bool, logging.LogLevel) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *MockSimpleLogging_NewLogger_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []bool, _param2 []logging.LogLevel) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]bool, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(bool)
		}
		_param2 = make([]logging.LogLevel, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(logging.LogLevel)
		}
	}
	return
}
