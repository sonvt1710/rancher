// Code generated by MockGen. DO NOT EDIT.
// Source: git.go

// Package git is a generated GoMock package.
package git

import (
	reflect "reflect"

	plumbing "github.com/go-git/go-git/v5/plumbing"
	transport "github.com/go-git/go-git/v5/plumbing/transport"
	gomock "github.com/golang/mock/gomock"
)

// Mockrepository is a mock of repository interface.
type Mockrepository struct {
	ctrl     *gomock.Controller
	recorder *MockrepositoryMockRecorder
}

// MockrepositoryMockRecorder is the mock recorder for Mockrepository.
type MockrepositoryMockRecorder struct {
	mock *Mockrepository
}

// NewMockrepository creates a new mock instance.
func NewMockrepository(ctrl *gomock.Controller) *Mockrepository {
	mock := &Mockrepository{ctrl: ctrl}
	mock.recorder = &MockrepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockrepository) EXPECT() *MockrepositoryMockRecorder {
	return m.recorder
}

// GetAuth mocks base method.
func (m *Mockrepository) GetAuth() transport.AuthMethod {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuth")
	ret0, _ := ret[0].(transport.AuthMethod)
	return ret0
}

// GetAuth indicates an expected call of GetAuth.
func (mr *MockrepositoryMockRecorder) GetAuth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuth", reflect.TypeOf((*Mockrepository)(nil).GetAuth))
}

// GetConfig mocks base method.
func (m *Mockrepository) GetConfig() *git {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*git)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockrepositoryMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*Mockrepository)(nil).GetConfig))
}

// checkout mocks base method.
func (m *Mockrepository) checkout(branch plumbing.ReferenceName) (plumbing.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "checkout", branch)
	ret0, _ := ret[0].(plumbing.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// checkout indicates an expected call of checkout.
func (mr *MockrepositoryMockRecorder) checkout(branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "checkout", reflect.TypeOf((*Mockrepository)(nil).checkout), branch)
}

// cloneOrOpen mocks base method.
func (m *Mockrepository) cloneOrOpen(branch string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "cloneOrOpen", branch)
	ret0, _ := ret[0].(error)
	return ret0
}

// cloneOrOpen indicates an expected call of cloneOrOpen.
func (mr *MockrepositoryMockRecorder) cloneOrOpen(branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "cloneOrOpen", reflect.TypeOf((*Mockrepository)(nil).cloneOrOpen), branch)
}

// fetch mocks base method.
func (m *Mockrepository) fetch(branch string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "fetch", branch)
	ret0, _ := ret[0].(error)
	return ret0
}

// fetch indicates an expected call of fetch.
func (mr *MockrepositoryMockRecorder) fetch(branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "fetch", reflect.TypeOf((*Mockrepository)(nil).fetch), branch)
}

// fetchAndReset mocks base method.
func (m *Mockrepository) fetchAndReset(branch string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "fetchAndReset", branch)
	ret0, _ := ret[0].(error)
	return ret0
}

// fetchAndReset indicates an expected call of fetchAndReset.
func (mr *MockrepositoryMockRecorder) fetchAndReset(branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "fetchAndReset", reflect.TypeOf((*Mockrepository)(nil).fetchAndReset), branch)
}

// getCurrentCommit mocks base method.
func (m *Mockrepository) getCurrentCommit() (plumbing.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getCurrentCommit")
	ret0, _ := ret[0].(plumbing.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getCurrentCommit indicates an expected call of getCurrentCommit.
func (mr *MockrepositoryMockRecorder) getCurrentCommit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getCurrentCommit", reflect.TypeOf((*Mockrepository)(nil).getCurrentCommit))
}

// getLastCommitHash mocks base method.
func (m *Mockrepository) getLastCommitHash(branch string, commitHASH plumbing.Hash) (plumbing.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getLastCommitHash", branch, commitHASH)
	ret0, _ := ret[0].(plumbing.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getLastCommitHash indicates an expected call of getLastCommitHash.
func (mr *MockrepositoryMockRecorder) getLastCommitHash(branch, commitHASH interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getLastCommitHash", reflect.TypeOf((*Mockrepository)(nil).getLastCommitHash), branch, commitHASH)
}

// hardReset mocks base method.
func (m *Mockrepository) hardReset(reference string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "hardReset", reference)
	ret0, _ := ret[0].(error)
	return ret0
}

// hardReset indicates an expected call of hardReset.
func (mr *MockrepositoryMockRecorder) hardReset(reference interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "hardReset", reflect.TypeOf((*Mockrepository)(nil).hardReset), reference)
}

// plainOpen mocks base method.
func (m *Mockrepository) plainOpen() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "plainOpen")
	ret0, _ := ret[0].(error)
	return ret0
}

// plainOpen indicates an expected call of plainOpen.
func (mr *MockrepositoryMockRecorder) plainOpen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "plainOpen", reflect.TypeOf((*Mockrepository)(nil).plainOpen))
}

// updateRefSpec mocks base method.
func (m *Mockrepository) updateRefSpec(branch string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "updateRefSpec", branch)
}

// updateRefSpec indicates an expected call of updateRefSpec.
func (mr *MockrepositoryMockRecorder) updateRefSpec(branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "updateRefSpec", reflect.TypeOf((*Mockrepository)(nil).updateRefSpec), branch)
}
