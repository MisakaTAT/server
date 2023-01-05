// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/bangumi/server/domain"
	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/internal/model"

	subject "github.com/bangumi/server/internal/subject"
)

// SubjectRepo is an autogenerated mock type for the Repo type
type SubjectRepo struct {
	mock.Mock
}

type SubjectRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *SubjectRepo) EXPECT() *SubjectRepo_Expecter {
	return &SubjectRepo_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, id, filter
func (_m *SubjectRepo) Get(ctx context.Context, id model.SubjectID, filter subject.Filter) (model.Subject, error) {
	ret := _m.Called(ctx, id, filter)

	var r0 model.Subject
	if rf, ok := ret.Get(0).(func(context.Context, model.SubjectID, subject.Filter) model.Subject); ok {
		r0 = rf(ctx, id, filter)
	} else {
		r0 = ret.Get(0).(model.Subject)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.SubjectID, subject.Filter) error); ok {
		r1 = rf(ctx, id, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type SubjectRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id model.SubjectID
//   - filter subject.Filter
func (_e *SubjectRepo_Expecter) Get(ctx interface{}, id interface{}, filter interface{}) *SubjectRepo_Get_Call {
	return &SubjectRepo_Get_Call{Call: _e.mock.On("Get", ctx, id, filter)}
}

func (_c *SubjectRepo_Get_Call) Run(run func(ctx context.Context, id model.SubjectID, filter subject.Filter)) *SubjectRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SubjectID), args[2].(subject.Filter))
	})
	return _c
}

func (_c *SubjectRepo_Get_Call) Return(_a0 model.Subject, _a1 error) *SubjectRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetActors provides a mock function with given fields: ctx, subjectID, characterIDs
func (_m *SubjectRepo) GetActors(ctx context.Context, subjectID model.SubjectID, characterIDs []model.CharacterID) (map[model.CharacterID][]model.PersonID, error) {
	ret := _m.Called(ctx, subjectID, characterIDs)

	var r0 map[model.CharacterID][]model.PersonID
	if rf, ok := ret.Get(0).(func(context.Context, model.SubjectID, []model.CharacterID) map[model.CharacterID][]model.PersonID); ok {
		r0 = rf(ctx, subjectID, characterIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[model.CharacterID][]model.PersonID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.SubjectID, []model.CharacterID) error); ok {
		r1 = rf(ctx, subjectID, characterIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectRepo_GetActors_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getActors'
type SubjectRepo_GetActors_Call struct {
	*mock.Call
}

// GetActors is a helper method to define mock.On call
//   - ctx context.Context
//   - subjectID model.SubjectID
//   - characterIDs []model.CharacterID
func (_e *SubjectRepo_Expecter) GetActors(ctx interface{}, subjectID interface{}, characterIDs interface{}) *SubjectRepo_GetActors_Call {
	return &SubjectRepo_GetActors_Call{Call: _e.mock.On("getActors", ctx, subjectID, characterIDs)}
}

func (_c *SubjectRepo_GetActors_Call) Run(run func(ctx context.Context, subjectID model.SubjectID, characterIDs []model.CharacterID)) *SubjectRepo_GetActors_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SubjectID), args[2].([]model.CharacterID))
	})
	return _c
}

func (_c *SubjectRepo_GetActors_Call) Return(_a0 map[model.CharacterID][]model.PersonID, _a1 error) *SubjectRepo_GetActors_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetByIDs provides a mock function with given fields: ctx, ids, filter
func (_m *SubjectRepo) GetByIDs(ctx context.Context, ids []model.SubjectID, filter subject.Filter) (map[model.SubjectID]model.Subject, error) {
	ret := _m.Called(ctx, ids, filter)

	var r0 map[model.SubjectID]model.Subject
	if rf, ok := ret.Get(0).(func(context.Context, []model.SubjectID, subject.Filter) map[model.SubjectID]model.Subject); ok {
		r0 = rf(ctx, ids, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[model.SubjectID]model.Subject)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []model.SubjectID, subject.Filter) error); ok {
		r1 = rf(ctx, ids, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectRepo_GetByIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByIDs'
type SubjectRepo_GetByIDs_Call struct {
	*mock.Call
}

// GetByIDs is a helper method to define mock.On call
//   - ctx context.Context
//   - ids []model.SubjectID
//   - filter subject.Filter
func (_e *SubjectRepo_Expecter) GetByIDs(ctx interface{}, ids interface{}, filter interface{}) *SubjectRepo_GetByIDs_Call {
	return &SubjectRepo_GetByIDs_Call{Call: _e.mock.On("GetByIDs", ctx, ids, filter)}
}

func (_c *SubjectRepo_GetByIDs_Call) Run(run func(ctx context.Context, ids []model.SubjectID, filter subject.Filter)) *SubjectRepo_GetByIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]model.SubjectID), args[2].(subject.Filter))
	})
	return _c
}

func (_c *SubjectRepo_GetByIDs_Call) Return(_a0 map[model.SubjectID]model.Subject, _a1 error) *SubjectRepo_GetByIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetCharacterRelated provides a mock function with given fields: ctx, characterID
func (_m *SubjectRepo) GetCharacterRelated(ctx context.Context, characterID model.CharacterID) ([]domain.SubjectCharacterRelation, error) {
	ret := _m.Called(ctx, characterID)

	var r0 []domain.SubjectCharacterRelation
	if rf, ok := ret.Get(0).(func(context.Context, model.CharacterID) []domain.SubjectCharacterRelation); ok {
		r0 = rf(ctx, characterID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.SubjectCharacterRelation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.CharacterID) error); ok {
		r1 = rf(ctx, characterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectRepo_GetCharacterRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCharacterRelated'
type SubjectRepo_GetCharacterRelated_Call struct {
	*mock.Call
}

// GetCharacterRelated is a helper method to define mock.On call
//   - ctx context.Context
//   - characterID model.CharacterID
func (_e *SubjectRepo_Expecter) GetCharacterRelated(ctx interface{}, characterID interface{}) *SubjectRepo_GetCharacterRelated_Call {
	return &SubjectRepo_GetCharacterRelated_Call{Call: _e.mock.On("GetCharacterRelated", ctx, characterID)}
}

func (_c *SubjectRepo_GetCharacterRelated_Call) Run(run func(ctx context.Context, characterID model.CharacterID)) *SubjectRepo_GetCharacterRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.CharacterID))
	})
	return _c
}

func (_c *SubjectRepo_GetCharacterRelated_Call) Return(_a0 []domain.SubjectCharacterRelation, _a1 error) *SubjectRepo_GetCharacterRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPersonRelated provides a mock function with given fields: ctx, personID
func (_m *SubjectRepo) GetPersonRelated(ctx context.Context, personID model.PersonID) ([]domain.SubjectPersonRelation, error) {
	ret := _m.Called(ctx, personID)

	var r0 []domain.SubjectPersonRelation
	if rf, ok := ret.Get(0).(func(context.Context, model.PersonID) []domain.SubjectPersonRelation); ok {
		r0 = rf(ctx, personID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.SubjectPersonRelation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.PersonID) error); ok {
		r1 = rf(ctx, personID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectRepo_GetPersonRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPersonRelated'
type SubjectRepo_GetPersonRelated_Call struct {
	*mock.Call
}

// GetPersonRelated is a helper method to define mock.On call
//   - ctx context.Context
//   - personID model.PersonID
func (_e *SubjectRepo_Expecter) GetPersonRelated(ctx interface{}, personID interface{}) *SubjectRepo_GetPersonRelated_Call {
	return &SubjectRepo_GetPersonRelated_Call{Call: _e.mock.On("GetPersonRelated", ctx, personID)}
}

func (_c *SubjectRepo_GetPersonRelated_Call) Run(run func(ctx context.Context, personID model.PersonID)) *SubjectRepo_GetPersonRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.PersonID))
	})
	return _c
}

func (_c *SubjectRepo_GetPersonRelated_Call) Return(_a0 []domain.SubjectPersonRelation, _a1 error) *SubjectRepo_GetPersonRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetSubjectRelated provides a mock function with given fields: ctx, subjectID
func (_m *SubjectRepo) GetSubjectRelated(ctx context.Context, subjectID model.SubjectID) ([]domain.SubjectInternalRelation, error) {
	ret := _m.Called(ctx, subjectID)

	var r0 []domain.SubjectInternalRelation
	if rf, ok := ret.Get(0).(func(context.Context, model.SubjectID) []domain.SubjectInternalRelation); ok {
		r0 = rf(ctx, subjectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.SubjectInternalRelation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.SubjectID) error); ok {
		r1 = rf(ctx, subjectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectRepo_GetSubjectRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubjectRelated'
type SubjectRepo_GetSubjectRelated_Call struct {
	*mock.Call
}

// GetSubjectRelated is a helper method to define mock.On call
//   - ctx context.Context
//   - subjectID model.SubjectID
func (_e *SubjectRepo_Expecter) GetSubjectRelated(ctx interface{}, subjectID interface{}) *SubjectRepo_GetSubjectRelated_Call {
	return &SubjectRepo_GetSubjectRelated_Call{Call: _e.mock.On("GetSubjectRelated", ctx, subjectID)}
}

func (_c *SubjectRepo_GetSubjectRelated_Call) Run(run func(ctx context.Context, subjectID model.SubjectID)) *SubjectRepo_GetSubjectRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SubjectID))
	})
	return _c
}

func (_c *SubjectRepo_GetSubjectRelated_Call) Return(_a0 []domain.SubjectInternalRelation, _a1 error) *SubjectRepo_GetSubjectRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewSubjectRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewSubjectRepo creates a new instance of SubjectRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSubjectRepo(t mockConstructorTestingTNewSubjectRepo) *SubjectRepo {
	mock := &SubjectRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
