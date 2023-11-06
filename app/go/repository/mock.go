// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package repository

import (
	"problem1/models"
	"sync"
)

// IUserRepositoryMock is a mock implementation of services.IUserRepository.
//
//	func TestSomethingThatUsesIUserRepository(t *testing.T) {
//
//		// make and configure a mocked services.IUserRepository
//		mockedIUserRepository := &IUserRepositoryMock{
//			FindBlockedUsersByUserIDFunc: func(userID int) ([]*models.User, error) {
//				panic("mock out the FindBlockedUsersByUserID method")
//			},
//			FindFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
//				panic("mock out the FindFriendsByUserID method")
//			},
//			FindFriendsOfFriendsByUserIDFunc: func(userID int) ([]*models.User, error) {
//				panic("mock out the FindFriendsOfFriendsByUserID method")
//			},
//			FindUsersByIDsFunc: func(userIDs []int) ([]*models.User, error) {
//				panic("mock out the FindUsersByIDs method")
//			},
//			FindUsersByIDsPagingFunc: func(userIDs []int, page int, limit int) ([]*models.User, error) {
//				panic("mock out the FindUsersByIDsPaging method")
//			},
//		}
//
//		// use mockedIUserRepository in code that requires services.IUserRepository
//		// and then make assertions.
//
//	}
type IUserRepositoryMock struct {
	// FindBlockedUsersByUserIDFunc mocks the FindBlockedUsersByUserID method.
	FindBlockedUsersByUserIDFunc func(userID int) ([]*models.User, error)

	// FindFriendsByUserIDFunc mocks the FindFriendsByUserID method.
	FindFriendsByUserIDFunc func(userID int) ([]*models.User, error)

	// FindFriendsOfFriendsByUserIDFunc mocks the FindFriendsOfFriendsByUserID method.
	FindFriendsOfFriendsByUserIDFunc func(userID int) ([]*models.User, error)

	// FindUsersByIDsFunc mocks the FindUsersByIDs method.
	FindUsersByIDsFunc func(userIDs []int) ([]*models.User, error)

	// FindUsersByIDsPagingFunc mocks the FindUsersByIDsPaging method.
	FindUsersByIDsPagingFunc func(userIDs []int, page int, limit int) ([]*models.User, error)

	// calls tracks calls to the methods.
	calls struct {
		// FindBlockedUsersByUserID holds details about calls to the FindBlockedUsersByUserID method.
		FindBlockedUsersByUserID []struct {
			// UserID is the userID argument value.
			UserID int
		}
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
		// FindUsersByIDs holds details about calls to the FindUsersByIDs method.
		FindUsersByIDs []struct {
			// UserIDs is the userIDs argument value.
			UserIDs []int
		}
		// FindUsersByIDsPaging holds details about calls to the FindUsersByIDsPaging method.
		FindUsersByIDsPaging []struct {
			// UserIDs is the userIDs argument value.
			UserIDs []int
			// Page is the page argument value.
			Page int
			// Limit is the limit argument value.
			Limit int
		}
	}
	lockFindBlockedUsersByUserID     sync.RWMutex
	lockFindFriendsByUserID          sync.RWMutex
	lockFindFriendsOfFriendsByUserID sync.RWMutex
	lockFindUsersByIDs               sync.RWMutex
	lockFindUsersByIDsPaging         sync.RWMutex
}

// FindBlockedUsersByUserID calls FindBlockedUsersByUserIDFunc.
func (mock *IUserRepositoryMock) FindBlockedUsersByUserID(userID int) ([]*models.User, error) {
	if mock.FindBlockedUsersByUserIDFunc == nil {
		panic("IUserRepositoryMock.FindBlockedUsersByUserIDFunc: method is nil but IUserRepository.FindBlockedUsersByUserID was just called")
	}
	callInfo := struct {
		UserID int
	}{
		UserID: userID,
	}
	mock.lockFindBlockedUsersByUserID.Lock()
	mock.calls.FindBlockedUsersByUserID = append(mock.calls.FindBlockedUsersByUserID, callInfo)
	mock.lockFindBlockedUsersByUserID.Unlock()
	return mock.FindBlockedUsersByUserIDFunc(userID)
}

// FindBlockedUsersByUserIDCalls gets all the calls that were made to FindBlockedUsersByUserID.
// Check the length with:
//
//	len(mockedIUserRepository.FindBlockedUsersByUserIDCalls())
func (mock *IUserRepositoryMock) FindBlockedUsersByUserIDCalls() []struct {
	UserID int
} {
	var calls []struct {
		UserID int
	}
	mock.lockFindBlockedUsersByUserID.RLock()
	calls = mock.calls.FindBlockedUsersByUserID
	mock.lockFindBlockedUsersByUserID.RUnlock()
	return calls
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

// FindUsersByIDs calls FindUsersByIDsFunc.
func (mock *IUserRepositoryMock) FindUsersByIDs(userIDs []int) ([]*models.User, error) {
	if mock.FindUsersByIDsFunc == nil {
		panic("IUserRepositoryMock.FindUsersByIDsFunc: method is nil but IUserRepository.FindUsersByIDs was just called")
	}
	callInfo := struct {
		UserIDs []int
	}{
		UserIDs: userIDs,
	}
	mock.lockFindUsersByIDs.Lock()
	mock.calls.FindUsersByIDs = append(mock.calls.FindUsersByIDs, callInfo)
	mock.lockFindUsersByIDs.Unlock()
	return mock.FindUsersByIDsFunc(userIDs)
}

// FindUsersByIDsCalls gets all the calls that were made to FindUsersByIDs.
// Check the length with:
//
//	len(mockedIUserRepository.FindUsersByIDsCalls())
func (mock *IUserRepositoryMock) FindUsersByIDsCalls() []struct {
	UserIDs []int
} {
	var calls []struct {
		UserIDs []int
	}
	mock.lockFindUsersByIDs.RLock()
	calls = mock.calls.FindUsersByIDs
	mock.lockFindUsersByIDs.RUnlock()
	return calls
}

// FindUsersByIDsPaging calls FindUsersByIDsPagingFunc.
func (mock *IUserRepositoryMock) FindUsersByIDsPaging(userIDs []int, page int, limit int) ([]*models.User, error) {
	if mock.FindUsersByIDsPagingFunc == nil {
		panic("IUserRepositoryMock.FindUsersByIDsPagingFunc: method is nil but IUserRepository.FindUsersByIDsPaging was just called")
	}
	callInfo := struct {
		UserIDs []int
		Page    int
		Limit   int
	}{
		UserIDs: userIDs,
		Page:    page,
		Limit:   limit,
	}
	mock.lockFindUsersByIDsPaging.Lock()
	mock.calls.FindUsersByIDsPaging = append(mock.calls.FindUsersByIDsPaging, callInfo)
	mock.lockFindUsersByIDsPaging.Unlock()
	return mock.FindUsersByIDsPagingFunc(userIDs, page, limit)
}

// FindUsersByIDsPagingCalls gets all the calls that were made to FindUsersByIDsPaging.
// Check the length with:
//
//	len(mockedIUserRepository.FindUsersByIDsPagingCalls())
func (mock *IUserRepositoryMock) FindUsersByIDsPagingCalls() []struct {
	UserIDs []int
	Page    int
	Limit   int
} {
	var calls []struct {
		UserIDs []int
		Page    int
		Limit   int
	}
	mock.lockFindUsersByIDsPaging.RLock()
	calls = mock.calls.FindUsersByIDsPaging
	mock.lockFindUsersByIDsPaging.RUnlock()
	return calls
}
