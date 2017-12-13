package mocks

import author "github.com/bxcodec/go-clean-arch/author"
import mock "github.com/stretchr/testify/mock"

// AuthorRepository is an autogenerated mock type for the AuthorRepository type
type AuthorRepository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: id
func (_m *AuthorRepository) GetByID(id int64) (*author.Author, error) {
	ret := _m.Called(id)

	var r0 *author.Author
	if rf, ok := ret.Get(0).(func(int64) *author.Author); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*author.Author)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
