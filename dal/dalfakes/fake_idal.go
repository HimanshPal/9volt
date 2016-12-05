// This file was generated by counterfeiter
package dalfakes

import (
	"sync"

	"github.com/9corp/9volt/dal"
	"github.com/coreos/etcd/client"
)

type FakeIDal struct {
	GetStub        func(string, *dal.GetOptions) (map[string]string, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
		arg2 *dal.GetOptions
	}
	getReturns struct {
		result1 map[string]string
		result2 error
	}
	SetStub        func(string, string, bool, int, string) error
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 bool
		arg4 int
		arg5 string
	}
	setReturns struct {
		result1 error
	}
	DeleteStub        func(string, bool) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
		arg2 bool
	}
	deleteReturns struct {
		result1 error
	}
	RefreshStub        func(string, int) error
	refreshMutex       sync.RWMutex
	refreshArgsForCall []struct {
		arg1 string
		arg2 int
	}
	refreshReturns struct {
		result1 error
	}
	KeyExistsStub        func(string) (bool, bool, error)
	keyExistsMutex       sync.RWMutex
	keyExistsArgsForCall []struct {
		arg1 string
	}
	keyExistsReturns struct {
		result1 bool
		result2 bool
		result3 error
	}
	IsKeyNotFoundStub        func(error) bool
	isKeyNotFoundMutex       sync.RWMutex
	isKeyNotFoundArgsForCall []struct {
		arg1 error
	}
	isKeyNotFoundReturns struct {
		result1 bool
	}
	CreateDirectorStateStub        func(string) error
	createDirectorStateMutex       sync.RWMutex
	createDirectorStateArgsForCall []struct {
		arg1 string
	}
	createDirectorStateReturns struct {
		result1 error
	}
	UpdateDirectorStateStub        func(string, string, bool) error
	updateDirectorStateMutex       sync.RWMutex
	updateDirectorStateArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 bool
	}
	updateDirectorStateReturns struct {
		result1 error
	}
	NewWatcherStub        func(string, bool) client.Watcher
	newWatcherMutex       sync.RWMutex
	newWatcherArgsForCall []struct {
		arg1 string
		arg2 bool
	}
	newWatcherReturns struct {
		result1 client.Watcher
	}
	GetClusterMembersStub        func() ([]string, error)
	getClusterMembersMutex       sync.RWMutex
	getClusterMembersArgsForCall []struct{}
	getClusterMembersReturns     struct {
		result1 []string
		result2 error
	}
	GetCheckKeysStub        func() ([]string, error)
	getCheckKeysMutex       sync.RWMutex
	getCheckKeysArgsForCall []struct{}
	getCheckKeysReturns     struct {
		result1 []string
		result2 error
	}
	CreateCheckReferenceStub        func(string, string) error
	createCheckReferenceMutex       sync.RWMutex
	createCheckReferenceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	createCheckReferenceReturns struct {
		result1 error
	}
	ClearCheckReferenceStub        func(string, string) error
	clearCheckReferenceMutex       sync.RWMutex
	clearCheckReferenceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	clearCheckReferenceReturns struct {
		result1 error
	}
	ClearCheckReferencesStub        func(string) error
	clearCheckReferencesMutex       sync.RWMutex
	clearCheckReferencesArgsForCall []struct {
		arg1 string
	}
	clearCheckReferencesReturns struct {
		result1 error
	}
	FetchAllMemberRefsStub        func() (map[string]string, error)
	fetchAllMemberRefsMutex       sync.RWMutex
	fetchAllMemberRefsArgsForCall []struct{}
	fetchAllMemberRefsReturns     struct {
		result1 map[string]string
		result2 error
	}
	FetchCheckStatsStub        func() (map[string]int, error)
	fetchCheckStatsMutex       sync.RWMutex
	fetchCheckStatsArgsForCall []struct{}
	fetchCheckStatsReturns     struct {
		result1 map[string]int
		result2 error
	}
	FetchAlerterConfigStub        func(string) (string, error)
	fetchAlerterConfigMutex       sync.RWMutex
	fetchAlerterConfigArgsForCall []struct {
		arg1 string
	}
	fetchAlerterConfigReturns struct {
		result1 string
		result2 error
	}
	invocations map[string][][]interface{}
}

func (fake *FakeIDal) Get(arg1 string, arg2 *dal.GetOptions) (map[string]string, error) {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
		arg2 *dal.GetOptions
	}{arg1, arg2})
	fake.guard("Get")
	fake.invocations["Get"] = append(fake.invocations["Get"], []interface{}{arg1, arg2})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2)
	} else {
		return fake.getReturns.result1, fake.getReturns.result2
	}
}

func (fake *FakeIDal) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeIDal) GetArgsForCall(i int) (string, *dal.GetOptions) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].arg1, fake.getArgsForCall[i].arg2
}

func (fake *FakeIDal) GetReturns(result1 map[string]string, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeIDal) Set(arg1 string, arg2 string, arg3 bool, arg4 int, arg5 string) error {
	fake.setMutex.Lock()
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 bool
		arg4 int
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.guard("Set")
	fake.invocations["Set"] = append(fake.invocations["Set"], []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.setMutex.Unlock()
	if fake.SetStub != nil {
		return fake.SetStub(arg1, arg2, arg3, arg4, arg5)
	} else {
		return fake.setReturns.result1
	}
}

func (fake *FakeIDal) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *FakeIDal) SetArgsForCall(i int) (string, string, bool, int, string) {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return fake.setArgsForCall[i].arg1, fake.setArgsForCall[i].arg2, fake.setArgsForCall[i].arg3, fake.setArgsForCall[i].arg4, fake.setArgsForCall[i].arg5
}

func (fake *FakeIDal) SetReturns(result1 error) {
	fake.SetStub = nil
	fake.setReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIDal) Delete(arg1 string, arg2 bool) error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
		arg2 bool
	}{arg1, arg2})
	fake.guard("Delete")
	fake.invocations["Delete"] = append(fake.invocations["Delete"], []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1, arg2)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeIDal) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeIDal) DeleteArgsForCall(i int) (string, bool) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].arg1, fake.deleteArgsForCall[i].arg2
}

func (fake *FakeIDal) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIDal) Refresh(arg1 string, arg2 int) error {
	fake.refreshMutex.Lock()
	fake.refreshArgsForCall = append(fake.refreshArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.guard("Refresh")
	fake.invocations["Refresh"] = append(fake.invocations["Refresh"], []interface{}{arg1, arg2})
	fake.refreshMutex.Unlock()
	if fake.RefreshStub != nil {
		return fake.RefreshStub(arg1, arg2)
	} else {
		return fake.refreshReturns.result1
	}
}

func (fake *FakeIDal) RefreshCallCount() int {
	fake.refreshMutex.RLock()
	defer fake.refreshMutex.RUnlock()
	return len(fake.refreshArgsForCall)
}

func (fake *FakeIDal) RefreshArgsForCall(i int) (string, int) {
	fake.refreshMutex.RLock()
	defer fake.refreshMutex.RUnlock()
	return fake.refreshArgsForCall[i].arg1, fake.refreshArgsForCall[i].arg2
}

func (fake *FakeIDal) RefreshReturns(result1 error) {
	fake.RefreshStub = nil
	fake.refreshReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIDal) KeyExists(arg1 string) (bool, bool, error) {
	fake.keyExistsMutex.Lock()
	fake.keyExistsArgsForCall = append(fake.keyExistsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.guard("KeyExists")
	fake.invocations["KeyExists"] = append(fake.invocations["KeyExists"], []interface{}{arg1})
	fake.keyExistsMutex.Unlock()
	if fake.KeyExistsStub != nil {
		return fake.KeyExistsStub(arg1)
	} else {
		return fake.keyExistsReturns.result1, fake.keyExistsReturns.result2, fake.keyExistsReturns.result3
	}
}

func (fake *FakeIDal) KeyExistsCallCount() int {
	fake.keyExistsMutex.RLock()
	defer fake.keyExistsMutex.RUnlock()
	return len(fake.keyExistsArgsForCall)
}

func (fake *FakeIDal) KeyExistsArgsForCall(i int) string {
	fake.keyExistsMutex.RLock()
	defer fake.keyExistsMutex.RUnlock()
	return fake.keyExistsArgsForCall[i].arg1
}

func (fake *FakeIDal) KeyExistsReturns(result1 bool, result2 bool, result3 error) {
	fake.KeyExistsStub = nil
	fake.keyExistsReturns = struct {
		result1 bool
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeIDal) IsKeyNotFound(arg1 error) bool {
	fake.isKeyNotFoundMutex.Lock()
	fake.isKeyNotFoundArgsForCall = append(fake.isKeyNotFoundArgsForCall, struct {
		arg1 error
	}{arg1})
	fake.guard("IsKeyNotFound")
	fake.invocations["IsKeyNotFound"] = append(fake.invocations["IsKeyNotFound"], []interface{}{arg1})
	fake.isKeyNotFoundMutex.Unlock()
	if fake.IsKeyNotFoundStub != nil {
		return fake.IsKeyNotFoundStub(arg1)
	} else {
		return fake.isKeyNotFoundReturns.result1
	}
}

func (fake *FakeIDal) IsKeyNotFoundCallCount() int {
	fake.isKeyNotFoundMutex.RLock()
	defer fake.isKeyNotFoundMutex.RUnlock()
	return len(fake.isKeyNotFoundArgsForCall)
}

func (fake *FakeIDal) IsKeyNotFoundArgsForCall(i int) error {
	fake.isKeyNotFoundMutex.RLock()
	defer fake.isKeyNotFoundMutex.RUnlock()
	return fake.isKeyNotFoundArgsForCall[i].arg1
}

func (fake *FakeIDal) IsKeyNotFoundReturns(result1 bool) {
	fake.IsKeyNotFoundStub = nil
	fake.isKeyNotFoundReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeIDal) CreateDirectorState(arg1 string) error {
	fake.createDirectorStateMutex.Lock()
	fake.createDirectorStateArgsForCall = append(fake.createDirectorStateArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.guard("CreateDirectorState")
	fake.invocations["CreateDirectorState"] = append(fake.invocations["CreateDirectorState"], []interface{}{arg1})
	fake.createDirectorStateMutex.Unlock()
	if fake.CreateDirectorStateStub != nil {
		return fake.CreateDirectorStateStub(arg1)
	} else {
		return fake.createDirectorStateReturns.result1
	}
}

func (fake *FakeIDal) CreateDirectorStateCallCount() int {
	fake.createDirectorStateMutex.RLock()
	defer fake.createDirectorStateMutex.RUnlock()
	return len(fake.createDirectorStateArgsForCall)
}

func (fake *FakeIDal) CreateDirectorStateArgsForCall(i int) string {
	fake.createDirectorStateMutex.RLock()
	defer fake.createDirectorStateMutex.RUnlock()
	return fake.createDirectorStateArgsForCall[i].arg1
}

func (fake *FakeIDal) CreateDirectorStateReturns(result1 error) {
	fake.CreateDirectorStateStub = nil
	fake.createDirectorStateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIDal) UpdateDirectorState(arg1 string, arg2 string, arg3 bool) error {
	fake.updateDirectorStateMutex.Lock()
	fake.updateDirectorStateArgsForCall = append(fake.updateDirectorStateArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 bool
	}{arg1, arg2, arg3})
	fake.guard("UpdateDirectorState")
	fake.invocations["UpdateDirectorState"] = append(fake.invocations["UpdateDirectorState"], []interface{}{arg1, arg2, arg3})
	fake.updateDirectorStateMutex.Unlock()
	if fake.UpdateDirectorStateStub != nil {
		return fake.UpdateDirectorStateStub(arg1, arg2, arg3)
	} else {
		return fake.updateDirectorStateReturns.result1
	}
}

func (fake *FakeIDal) UpdateDirectorStateCallCount() int {
	fake.updateDirectorStateMutex.RLock()
	defer fake.updateDirectorStateMutex.RUnlock()
	return len(fake.updateDirectorStateArgsForCall)
}

func (fake *FakeIDal) UpdateDirectorStateArgsForCall(i int) (string, string, bool) {
	fake.updateDirectorStateMutex.RLock()
	defer fake.updateDirectorStateMutex.RUnlock()
	return fake.updateDirectorStateArgsForCall[i].arg1, fake.updateDirectorStateArgsForCall[i].arg2, fake.updateDirectorStateArgsForCall[i].arg3
}

func (fake *FakeIDal) UpdateDirectorStateReturns(result1 error) {
	fake.UpdateDirectorStateStub = nil
	fake.updateDirectorStateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIDal) NewWatcher(arg1 string, arg2 bool) client.Watcher {
	fake.newWatcherMutex.Lock()
	fake.newWatcherArgsForCall = append(fake.newWatcherArgsForCall, struct {
		arg1 string
		arg2 bool
	}{arg1, arg2})
	fake.guard("NewWatcher")
	fake.invocations["NewWatcher"] = append(fake.invocations["NewWatcher"], []interface{}{arg1, arg2})
	fake.newWatcherMutex.Unlock()
	if fake.NewWatcherStub != nil {
		return fake.NewWatcherStub(arg1, arg2)
	} else {
		return fake.newWatcherReturns.result1
	}
}

func (fake *FakeIDal) NewWatcherCallCount() int {
	fake.newWatcherMutex.RLock()
	defer fake.newWatcherMutex.RUnlock()
	return len(fake.newWatcherArgsForCall)
}

func (fake *FakeIDal) NewWatcherArgsForCall(i int) (string, bool) {
	fake.newWatcherMutex.RLock()
	defer fake.newWatcherMutex.RUnlock()
	return fake.newWatcherArgsForCall[i].arg1, fake.newWatcherArgsForCall[i].arg2
}

func (fake *FakeIDal) NewWatcherReturns(result1 client.Watcher) {
	fake.NewWatcherStub = nil
	fake.newWatcherReturns = struct {
		result1 client.Watcher
	}{result1}
}

func (fake *FakeIDal) GetClusterMembers() ([]string, error) {
	fake.getClusterMembersMutex.Lock()
	fake.getClusterMembersArgsForCall = append(fake.getClusterMembersArgsForCall, struct{}{})
	fake.guard("GetClusterMembers")
	fake.invocations["GetClusterMembers"] = append(fake.invocations["GetClusterMembers"], []interface{}{})
	fake.getClusterMembersMutex.Unlock()
	if fake.GetClusterMembersStub != nil {
		return fake.GetClusterMembersStub()
	} else {
		return fake.getClusterMembersReturns.result1, fake.getClusterMembersReturns.result2
	}
}

func (fake *FakeIDal) GetClusterMembersCallCount() int {
	fake.getClusterMembersMutex.RLock()
	defer fake.getClusterMembersMutex.RUnlock()
	return len(fake.getClusterMembersArgsForCall)
}

func (fake *FakeIDal) GetClusterMembersReturns(result1 []string, result2 error) {
	fake.GetClusterMembersStub = nil
	fake.getClusterMembersReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeIDal) GetCheckKeys() ([]string, error) {
	fake.getCheckKeysMutex.Lock()
	fake.getCheckKeysArgsForCall = append(fake.getCheckKeysArgsForCall, struct{}{})
	fake.guard("GetCheckKeys")
	fake.invocations["GetCheckKeys"] = append(fake.invocations["GetCheckKeys"], []interface{}{})
	fake.getCheckKeysMutex.Unlock()
	if fake.GetCheckKeysStub != nil {
		return fake.GetCheckKeysStub()
	} else {
		return fake.getCheckKeysReturns.result1, fake.getCheckKeysReturns.result2
	}
}

func (fake *FakeIDal) GetCheckKeysCallCount() int {
	fake.getCheckKeysMutex.RLock()
	defer fake.getCheckKeysMutex.RUnlock()
	return len(fake.getCheckKeysArgsForCall)
}

func (fake *FakeIDal) GetCheckKeysReturns(result1 []string, result2 error) {
	fake.GetCheckKeysStub = nil
	fake.getCheckKeysReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeIDal) CreateCheckReference(arg1 string, arg2 string) error {
	fake.createCheckReferenceMutex.Lock()
	fake.createCheckReferenceArgsForCall = append(fake.createCheckReferenceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.guard("CreateCheckReference")
	fake.invocations["CreateCheckReference"] = append(fake.invocations["CreateCheckReference"], []interface{}{arg1, arg2})
	fake.createCheckReferenceMutex.Unlock()
	if fake.CreateCheckReferenceStub != nil {
		return fake.CreateCheckReferenceStub(arg1, arg2)
	} else {
		return fake.createCheckReferenceReturns.result1
	}
}

func (fake *FakeIDal) CreateCheckReferenceCallCount() int {
	fake.createCheckReferenceMutex.RLock()
	defer fake.createCheckReferenceMutex.RUnlock()
	return len(fake.createCheckReferenceArgsForCall)
}

func (fake *FakeIDal) CreateCheckReferenceArgsForCall(i int) (string, string) {
	fake.createCheckReferenceMutex.RLock()
	defer fake.createCheckReferenceMutex.RUnlock()
	return fake.createCheckReferenceArgsForCall[i].arg1, fake.createCheckReferenceArgsForCall[i].arg2
}

func (fake *FakeIDal) CreateCheckReferenceReturns(result1 error) {
	fake.CreateCheckReferenceStub = nil
	fake.createCheckReferenceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIDal) ClearCheckReference(arg1 string, arg2 string) error {
	fake.clearCheckReferenceMutex.Lock()
	fake.clearCheckReferenceArgsForCall = append(fake.clearCheckReferenceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.guard("ClearCheckReference")
	fake.invocations["ClearCheckReference"] = append(fake.invocations["ClearCheckReference"], []interface{}{arg1, arg2})
	fake.clearCheckReferenceMutex.Unlock()
	if fake.ClearCheckReferenceStub != nil {
		return fake.ClearCheckReferenceStub(arg1, arg2)
	} else {
		return fake.clearCheckReferenceReturns.result1
	}
}

func (fake *FakeIDal) ClearCheckReferenceCallCount() int {
	fake.clearCheckReferenceMutex.RLock()
	defer fake.clearCheckReferenceMutex.RUnlock()
	return len(fake.clearCheckReferenceArgsForCall)
}

func (fake *FakeIDal) ClearCheckReferenceArgsForCall(i int) (string, string) {
	fake.clearCheckReferenceMutex.RLock()
	defer fake.clearCheckReferenceMutex.RUnlock()
	return fake.clearCheckReferenceArgsForCall[i].arg1, fake.clearCheckReferenceArgsForCall[i].arg2
}

func (fake *FakeIDal) ClearCheckReferenceReturns(result1 error) {
	fake.ClearCheckReferenceStub = nil
	fake.clearCheckReferenceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIDal) ClearCheckReferences(arg1 string) error {
	fake.clearCheckReferencesMutex.Lock()
	fake.clearCheckReferencesArgsForCall = append(fake.clearCheckReferencesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.guard("ClearCheckReferences")
	fake.invocations["ClearCheckReferences"] = append(fake.invocations["ClearCheckReferences"], []interface{}{arg1})
	fake.clearCheckReferencesMutex.Unlock()
	if fake.ClearCheckReferencesStub != nil {
		return fake.ClearCheckReferencesStub(arg1)
	} else {
		return fake.clearCheckReferencesReturns.result1
	}
}

func (fake *FakeIDal) ClearCheckReferencesCallCount() int {
	fake.clearCheckReferencesMutex.RLock()
	defer fake.clearCheckReferencesMutex.RUnlock()
	return len(fake.clearCheckReferencesArgsForCall)
}

func (fake *FakeIDal) ClearCheckReferencesArgsForCall(i int) string {
	fake.clearCheckReferencesMutex.RLock()
	defer fake.clearCheckReferencesMutex.RUnlock()
	return fake.clearCheckReferencesArgsForCall[i].arg1
}

func (fake *FakeIDal) ClearCheckReferencesReturns(result1 error) {
	fake.ClearCheckReferencesStub = nil
	fake.clearCheckReferencesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIDal) FetchAllMemberRefs() (map[string]string, error) {
	fake.fetchAllMemberRefsMutex.Lock()
	fake.fetchAllMemberRefsArgsForCall = append(fake.fetchAllMemberRefsArgsForCall, struct{}{})
	fake.guard("FetchAllMemberRefs")
	fake.invocations["FetchAllMemberRefs"] = append(fake.invocations["FetchAllMemberRefs"], []interface{}{})
	fake.fetchAllMemberRefsMutex.Unlock()
	if fake.FetchAllMemberRefsStub != nil {
		return fake.FetchAllMemberRefsStub()
	} else {
		return fake.fetchAllMemberRefsReturns.result1, fake.fetchAllMemberRefsReturns.result2
	}
}

func (fake *FakeIDal) FetchAllMemberRefsCallCount() int {
	fake.fetchAllMemberRefsMutex.RLock()
	defer fake.fetchAllMemberRefsMutex.RUnlock()
	return len(fake.fetchAllMemberRefsArgsForCall)
}

func (fake *FakeIDal) FetchAllMemberRefsReturns(result1 map[string]string, result2 error) {
	fake.FetchAllMemberRefsStub = nil
	fake.fetchAllMemberRefsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeIDal) FetchCheckStats() (map[string]int, error) {
	fake.fetchCheckStatsMutex.Lock()
	fake.fetchCheckStatsArgsForCall = append(fake.fetchCheckStatsArgsForCall, struct{}{})
	fake.guard("FetchCheckStats")
	fake.invocations["FetchCheckStats"] = append(fake.invocations["FetchCheckStats"], []interface{}{})
	fake.fetchCheckStatsMutex.Unlock()
	if fake.FetchCheckStatsStub != nil {
		return fake.FetchCheckStatsStub()
	} else {
		return fake.fetchCheckStatsReturns.result1, fake.fetchCheckStatsReturns.result2
	}
}

func (fake *FakeIDal) FetchCheckStatsCallCount() int {
	fake.fetchCheckStatsMutex.RLock()
	defer fake.fetchCheckStatsMutex.RUnlock()
	return len(fake.fetchCheckStatsArgsForCall)
}

func (fake *FakeIDal) FetchCheckStatsReturns(result1 map[string]int, result2 error) {
	fake.FetchCheckStatsStub = nil
	fake.fetchCheckStatsReturns = struct {
		result1 map[string]int
		result2 error
	}{result1, result2}
}

func (fake *FakeIDal) FetchAlerterConfig(arg1 string) (string, error) {
	fake.fetchAlerterConfigMutex.Lock()
	fake.fetchAlerterConfigArgsForCall = append(fake.fetchAlerterConfigArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.guard("FetchAlerterConfig")
	fake.invocations["FetchAlerterConfig"] = append(fake.invocations["FetchAlerterConfig"], []interface{}{arg1})
	fake.fetchAlerterConfigMutex.Unlock()
	if fake.FetchAlerterConfigStub != nil {
		return fake.FetchAlerterConfigStub(arg1)
	} else {
		return fake.fetchAlerterConfigReturns.result1, fake.fetchAlerterConfigReturns.result2
	}
}

func (fake *FakeIDal) FetchAlerterConfigCallCount() int {
	fake.fetchAlerterConfigMutex.RLock()
	defer fake.fetchAlerterConfigMutex.RUnlock()
	return len(fake.fetchAlerterConfigArgsForCall)
}

func (fake *FakeIDal) FetchAlerterConfigArgsForCall(i int) string {
	fake.fetchAlerterConfigMutex.RLock()
	defer fake.fetchAlerterConfigMutex.RUnlock()
	return fake.fetchAlerterConfigArgsForCall[i].arg1
}

func (fake *FakeIDal) FetchAlerterConfigReturns(result1 string, result2 error) {
	fake.FetchAlerterConfigStub = nil
	fake.fetchAlerterConfigReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeIDal) Invocations() map[string][][]interface{} {
	return fake.invocations
}

func (fake *FakeIDal) guard(key string) {
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
}

var _ dal.IDal = new(FakeIDal)
