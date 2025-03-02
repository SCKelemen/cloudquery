// Code generated by MockGen. DO NOT EDIT.
// Source: accessanalyzer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	accessanalyzer "github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	gomock "github.com/golang/mock/gomock"
)

// MockAccessanalyzerClient is a mock of AccessanalyzerClient interface.
type MockAccessanalyzerClient struct {
	ctrl     *gomock.Controller
	recorder *MockAccessanalyzerClientMockRecorder
}

// MockAccessanalyzerClientMockRecorder is the mock recorder for MockAccessanalyzerClient.
type MockAccessanalyzerClientMockRecorder struct {
	mock *MockAccessanalyzerClient
}

// NewMockAccessanalyzerClient creates a new mock instance.
func NewMockAccessanalyzerClient(ctrl *gomock.Controller) *MockAccessanalyzerClient {
	mock := &MockAccessanalyzerClient{ctrl: ctrl}
	mock.recorder = &MockAccessanalyzerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessanalyzerClient) EXPECT() *MockAccessanalyzerClientMockRecorder {
	return m.recorder
}

// GetAccessPreview mocks base method.
func (m *MockAccessanalyzerClient) GetAccessPreview(arg0 context.Context, arg1 *accessanalyzer.GetAccessPreviewInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.GetAccessPreviewOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccessPreview", varargs...)
	ret0, _ := ret[0].(*accessanalyzer.GetAccessPreviewOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessPreview indicates an expected call of GetAccessPreview.
func (mr *MockAccessanalyzerClientMockRecorder) GetAccessPreview(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessPreview", reflect.TypeOf((*MockAccessanalyzerClient)(nil).GetAccessPreview), varargs...)
}

// GetAnalyzedResource mocks base method.
func (m *MockAccessanalyzerClient) GetAnalyzedResource(arg0 context.Context, arg1 *accessanalyzer.GetAnalyzedResourceInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.GetAnalyzedResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAnalyzedResource", varargs...)
	ret0, _ := ret[0].(*accessanalyzer.GetAnalyzedResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnalyzedResource indicates an expected call of GetAnalyzedResource.
func (mr *MockAccessanalyzerClientMockRecorder) GetAnalyzedResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnalyzedResource", reflect.TypeOf((*MockAccessanalyzerClient)(nil).GetAnalyzedResource), varargs...)
}

// GetAnalyzer mocks base method.
func (m *MockAccessanalyzerClient) GetAnalyzer(arg0 context.Context, arg1 *accessanalyzer.GetAnalyzerInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.GetAnalyzerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAnalyzer", varargs...)
	ret0, _ := ret[0].(*accessanalyzer.GetAnalyzerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnalyzer indicates an expected call of GetAnalyzer.
func (mr *MockAccessanalyzerClientMockRecorder) GetAnalyzer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnalyzer", reflect.TypeOf((*MockAccessanalyzerClient)(nil).GetAnalyzer), varargs...)
}

// GetArchiveRule mocks base method.
func (m *MockAccessanalyzerClient) GetArchiveRule(arg0 context.Context, arg1 *accessanalyzer.GetArchiveRuleInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.GetArchiveRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetArchiveRule", varargs...)
	ret0, _ := ret[0].(*accessanalyzer.GetArchiveRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArchiveRule indicates an expected call of GetArchiveRule.
func (mr *MockAccessanalyzerClientMockRecorder) GetArchiveRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArchiveRule", reflect.TypeOf((*MockAccessanalyzerClient)(nil).GetArchiveRule), varargs...)
}

// GetFinding mocks base method.
func (m *MockAccessanalyzerClient) GetFinding(arg0 context.Context, arg1 *accessanalyzer.GetFindingInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.GetFindingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFinding", varargs...)
	ret0, _ := ret[0].(*accessanalyzer.GetFindingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFinding indicates an expected call of GetFinding.
func (mr *MockAccessanalyzerClientMockRecorder) GetFinding(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFinding", reflect.TypeOf((*MockAccessanalyzerClient)(nil).GetFinding), varargs...)
}

// GetGeneratedPolicy mocks base method.
func (m *MockAccessanalyzerClient) GetGeneratedPolicy(arg0 context.Context, arg1 *accessanalyzer.GetGeneratedPolicyInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.GetGeneratedPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGeneratedPolicy", varargs...)
	ret0, _ := ret[0].(*accessanalyzer.GetGeneratedPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGeneratedPolicy indicates an expected call of GetGeneratedPolicy.
func (mr *MockAccessanalyzerClientMockRecorder) GetGeneratedPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGeneratedPolicy", reflect.TypeOf((*MockAccessanalyzerClient)(nil).GetGeneratedPolicy), varargs...)
}

// ListAccessPreviewFindings mocks base method.
func (m *MockAccessanalyzerClient) ListAccessPreviewFindings(arg0 context.Context, arg1 *accessanalyzer.ListAccessPreviewFindingsInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAccessPreviewFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccessPreviewFindings", varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListAccessPreviewFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessPreviewFindings indicates an expected call of ListAccessPreviewFindings.
func (mr *MockAccessanalyzerClientMockRecorder) ListAccessPreviewFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessPreviewFindings", reflect.TypeOf((*MockAccessanalyzerClient)(nil).ListAccessPreviewFindings), varargs...)
}

// ListAccessPreviews mocks base method.
func (m *MockAccessanalyzerClient) ListAccessPreviews(arg0 context.Context, arg1 *accessanalyzer.ListAccessPreviewsInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAccessPreviewsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccessPreviews", varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListAccessPreviewsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessPreviews indicates an expected call of ListAccessPreviews.
func (mr *MockAccessanalyzerClientMockRecorder) ListAccessPreviews(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessPreviews", reflect.TypeOf((*MockAccessanalyzerClient)(nil).ListAccessPreviews), varargs...)
}

// ListAnalyzedResources mocks base method.
func (m *MockAccessanalyzerClient) ListAnalyzedResources(arg0 context.Context, arg1 *accessanalyzer.ListAnalyzedResourcesInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAnalyzedResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAnalyzedResources", varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListAnalyzedResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAnalyzedResources indicates an expected call of ListAnalyzedResources.
func (mr *MockAccessanalyzerClientMockRecorder) ListAnalyzedResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAnalyzedResources", reflect.TypeOf((*MockAccessanalyzerClient)(nil).ListAnalyzedResources), varargs...)
}

// ListAnalyzers mocks base method.
func (m *MockAccessanalyzerClient) ListAnalyzers(arg0 context.Context, arg1 *accessanalyzer.ListAnalyzersInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAnalyzersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAnalyzers", varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListAnalyzersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAnalyzers indicates an expected call of ListAnalyzers.
func (mr *MockAccessanalyzerClientMockRecorder) ListAnalyzers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAnalyzers", reflect.TypeOf((*MockAccessanalyzerClient)(nil).ListAnalyzers), varargs...)
}

// ListArchiveRules mocks base method.
func (m *MockAccessanalyzerClient) ListArchiveRules(arg0 context.Context, arg1 *accessanalyzer.ListArchiveRulesInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListArchiveRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListArchiveRules", varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListArchiveRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListArchiveRules indicates an expected call of ListArchiveRules.
func (mr *MockAccessanalyzerClientMockRecorder) ListArchiveRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListArchiveRules", reflect.TypeOf((*MockAccessanalyzerClient)(nil).ListArchiveRules), varargs...)
}

// ListFindings mocks base method.
func (m *MockAccessanalyzerClient) ListFindings(arg0 context.Context, arg1 *accessanalyzer.ListFindingsInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFindings", varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFindings indicates an expected call of ListFindings.
func (mr *MockAccessanalyzerClientMockRecorder) ListFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFindings", reflect.TypeOf((*MockAccessanalyzerClient)(nil).ListFindings), varargs...)
}

// ListPolicyGenerations mocks base method.
func (m *MockAccessanalyzerClient) ListPolicyGenerations(arg0 context.Context, arg1 *accessanalyzer.ListPolicyGenerationsInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListPolicyGenerationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPolicyGenerations", varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListPolicyGenerationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPolicyGenerations indicates an expected call of ListPolicyGenerations.
func (mr *MockAccessanalyzerClientMockRecorder) ListPolicyGenerations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPolicyGenerations", reflect.TypeOf((*MockAccessanalyzerClient)(nil).ListPolicyGenerations), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockAccessanalyzerClient) ListTagsForResource(arg0 context.Context, arg1 *accessanalyzer.ListTagsForResourceInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockAccessanalyzerClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockAccessanalyzerClient)(nil).ListTagsForResource), varargs...)
}
