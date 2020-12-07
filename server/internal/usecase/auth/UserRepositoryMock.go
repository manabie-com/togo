// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package auth

import (
	"context"
	"github.com/HoangVyDuong/togo/internal/storages/user"
	"sync"
)

var (
	lockRepositoryMockGetUserByName sync.RWMutex
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
//     func TestSomethingThatUsesRepository(t *testing.T) {
//
//         // make and configure a mocked Repository
//         mockedRepository := &RepositoryMock{
//             GetUserByNameFunc: func(ctx context.Context, name string) (user.User, error) {
// 	               panic("mock out the GetUserByName method")
//             },
//         }
//
//         // use mockedRepository in code that requires Repository
//         // and then make assertions.
//
//     }
type RepositoryMock struct {
	// GetUserByNameFunc mocks the GetUserByName method.
	GetUserByNameFunc func(ctx context.Context, name string) (user.User, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetUserByName holds details about calls to the GetUserByName method.
		GetUserByName []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
		}
	}
}

// GetUserByName calls GetUserByNameFunc.
func (mock *RepositoryMock) GetUserByName(ctx context.Context, name string) (user.User, error) {
	if mock.GetUserByNameFunc == nil {
		panic("RepositoryMock.GetUserByNameFunc: method is nil but Repository.GetUserByName was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Name string
	}{
		Ctx:  ctx,
		Name: name,
	}
	lockRepositoryMockGetUserByName.Lock()
	mock.calls.GetUserByName = append(mock.calls.GetUserByName, callInfo)
	lockRepositoryMockGetUserByName.Unlock()
	return mock.GetUserByNameFunc(ctx, name)
}

// GetUserByNameCalls gets all the calls that were made to GetUserByName.
// Check the length with:
//     len(mockedRepository.GetUserByNameCalls())
func (mock *RepositoryMock) GetUserByNameCalls() []struct {
	Ctx  context.Context
	Name string
} {
	var calls []struct {
		Ctx  context.Context
		Name string
	}
	lockRepositoryMockGetUserByName.RLock()
	calls = mock.calls.GetUserByName
	lockRepositoryMockGetUserByName.RUnlock()
	return calls
}
