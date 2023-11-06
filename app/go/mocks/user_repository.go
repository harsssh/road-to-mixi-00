// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"problem1/models"
	"problem1/services"
	"sync"
)

// Ensure, that IUserRepositoryMock does implement services.IUserRepository.
// If this is not the case, regenerate this file with moq.
var _ services.IUserRepository = &IUserRepositoryMock{}

// IUserRepositoryMock is a mock implementation of services.IUserRepository.
//
//	func TestSomethingThatUsesIUserRepository(t *testing.T) {
//
//		// make and configure a mocked services.IUserRepository
//		mockedIUserRepository := &IUserRepositoryMock{
//			FindFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
//				panic("mock out the FindFriendsByUserID method")
//			},
//			FindFriendsOfFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
//				panic("mock out the FindFriendsOfFriendsByUserID method")
//			},
//			FindFriendsOfFriendsPagingByUserIDFunc: func(userID int, page int, limit int) ([]*models.User, error) {
//				panic("mock out the FindFriendsOfFriendsPagingByUserID method")
//			},
//		}
//
//		// use mockedIUserRepository in code that requires services.IUserRepository
//		// and then make assertions.
//
//	}
type IUserRepositoryMock struct {
	// FindFriendsByUserIDFunc mocks the FindFriendsByUserID method.
	FindFriendsByUserIDFunc func(userID int) ([]*models.User, error)

	// FindFriendsOfFriendsByUserIDFunc mocks the FindFriendsOfFriendsByUserID method.
	FindFriendsOfFriendsByUserIDFunc func(userID int) ([]*models.User, error)

	// FindFriendsOfFriendsPagingByUserIDFunc mocks the FindFriendsOfFriendsPagingByUserID method.
	FindFriendsOfFriendsPagingByUserIDFunc func(userID int, page int, limit int) ([]*models.User, error)

	// calls tracks calls to the methods.
	calls struct {
		// FindFriendsByUserID holds details about calls to the FindFriendsByUserID method.
		FindFriendsByUserID []struct {
			// UserID is the userID argument value.
			UserID int
		}
		// FindFriendsOfFriendsByUserID holds details about calls to the FindFriendsOfFriendsByUserID method.
		FindFriendsOfFriendsByUserID []struct {
			// UserID is the userID argument value.
			UserID int
		}
		// FindFriendsOfFriendsPagingByUserID holds details about calls to the FindFriendsOfFriendsPagingByUserID method.
		FindFriendsOfFriendsPagingByUserID []struct {
			// UserID is the userID argument value.
			UserID int
			// Page is the page argument value.
			Page int
			// Limit is the limit argument value.
			Limit int
		}
	}
	lockFindFriendsByUserID                sync.RWMutex
	lockFindFriendsOfFriendsByUserID       sync.RWMutex
	lockFindFriendsOfFriendsPagingByUserID sync.RWMutex
}

// FindFriendsByUserID calls FindFriendsByUserIDFunc.
func (mock *IUserRepositoryMock) FindFriendsByUserID(userID int) ([]*models.User, error) {
	if mock.FindFriendsByUserIDFunc == nil {
		panic("IUserRepositoryMock.FindFriendsByUserIDFunc: method is nil but IUserRepository.FindFriendsByUserID was just called")
	}
	callInfo := struct {
		UserID int
	}{
		UserID: userID,
	}
	mock.lockFindFriendsByUserID.Lock()
	mock.calls.FindFriendsByUserID = append(mock.calls.FindFriendsByUserID, callInfo)
	mock.lockFindFriendsByUserID.Unlock()
	return mock.FindFriendsByUserIDFunc(userID)
}

// FindFriendsByUserIDCalls gets all the calls that were made to FindFriendsByUserID.
// Check the length with:
//
//	len(mockedIUserRepository.FindFriendsByUserIDCalls())
func (mock *IUserRepositoryMock) FindFriendsByUserIDCalls() []struct {
	UserID int
} {
	var calls []struct {
		UserID int
	}
	mock.lockFindFriendsByUserID.RLock()
	calls = mock.calls.FindFriendsByUserID
	mock.lockFindFriendsByUserID.RUnlock()
	return calls
}

// FindFriendsOfFriendsByUserID calls FindFriendsOfFriendsByUserIDFunc.
func (mock *IUserRepositoryMock) FindFriendsOfFriendsByUserID(userID int) ([]*models.User, error) {
	if mock.FindFriendsOfFriendsByUserIDFunc == nil {
		panic("IUserRepositoryMock.FindFriendsOfFriendsByUserIDFunc: method is nil but IUserRepository.FindFriendsOfFriendsByUserID was just called")
	}
	callInfo := struct {
		UserID int
	}{
		UserID: userID,
	}
	mock.lockFindFriendsOfFriendsByUserID.Lock()
	mock.calls.FindFriendsOfFriendsByUserID = append(mock.calls.FindFriendsOfFriendsByUserID, callInfo)
	mock.lockFindFriendsOfFriendsByUserID.Unlock()
	return mock.FindFriendsOfFriendsByUserIDFunc(userID)
}

// FindFriendsOfFriendsByUserIDCalls gets all the calls that were made to FindFriendsOfFriendsByUserID.
// Check the length with:
//
//	len(mockedIUserRepository.FindFriendsOfFriendsByUserIDCalls())
func (mock *IUserRepositoryMock) FindFriendsOfFriendsByUserIDCalls() []struct {
	UserID int
} {
	var calls []struct {
		UserID int
	}
	mock.lockFindFriendsOfFriendsByUserID.RLock()
	calls = mock.calls.FindFriendsOfFriendsByUserID
	mock.lockFindFriendsOfFriendsByUserID.RUnlock()
	return calls
}

// FindFriendsOfFriendsPagingByUserID calls FindFriendsOfFriendsPagingByUserIDFunc.
func (mock *IUserRepositoryMock) FindFriendsOfFriendsPagingByUserID(userID int, page int, limit int) ([]*models.User, error) {
	if mock.FindFriendsOfFriendsPagingByUserIDFunc == nil {
		panic("IUserRepositoryMock.FindFriendsOfFriendsPagingByUserIDFunc: method is nil but IUserRepository.FindFriendsOfFriendsPagingByUserID was just called")
	}
	callInfo := struct {
		UserID int
		Page   int
		Limit  int
	}{
		UserID: userID,
		Page:   page,
		Limit:  limit,
	}
	mock.lockFindFriendsOfFriendsPagingByUserID.Lock()
	mock.calls.FindFriendsOfFriendsPagingByUserID = append(mock.calls.FindFriendsOfFriendsPagingByUserID, callInfo)
	mock.lockFindFriendsOfFriendsPagingByUserID.Unlock()
	return mock.FindFriendsOfFriendsPagingByUserIDFunc(userID, page, limit)
}

// FindFriendsOfFriendsPagingByUserIDCalls gets all the calls that were made to FindFriendsOfFriendsPagingByUserID.
// Check the length with:
//
//	len(mockedIUserRepository.FindFriendsOfFriendsPagingByUserIDCalls())
func (mock *IUserRepositoryMock) FindFriendsOfFriendsPagingByUserIDCalls() []struct {
	UserID int
	Page   int
	Limit  int
} {
	var calls []struct {
		UserID int
		Page   int
		Limit  int
	}
	mock.lockFindFriendsOfFriendsPagingByUserID.RLock()
	calls = mock.calls.FindFriendsOfFriendsPagingByUserID
	mock.lockFindFriendsOfFriendsPagingByUserID.RUnlock()
	return calls
}
