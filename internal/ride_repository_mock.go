// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package internal

import (
	"sync"
)

var (
	lockRideRepositoryMockGet  sync.RWMutex
	lockRideRepositoryMockSave sync.RWMutex
)

// Ensure, that RideRepositoryMock does implement RideRepository.
// If this is not the case, regenerate this file with moq.
var _ RideRepository = &RideRepositoryMock{}

// RideRepositoryMock is a mock implementation of RideRepository.
//
//     func TestSomethingThatUsesRideRepository(t *testing.T) {
//
//         // make and configure a mocked RideRepository
//         mockedRideRepository := &RideRepositoryMock{
//             GetFunc: func(id int) (Ride, error) {
// 	               panic("mock out the Get method")
//             },
//             SaveFunc: func(in1 Ride) error {
// 	               panic("mock out the Save method")
//             },
//         }
//
//         // use mockedRideRepository in code that requires RideRepository
//         // and then make assertions.
//
//     }
type RideRepositoryMock struct {
	// GetFunc mocks the Get method.
	GetFunc func(id int) (Ride, error)

	// SaveFunc mocks the Save method.
	SaveFunc func(in1 Ride) error

	// calls tracks calls to the methods.
	calls struct {
		// Get holds details about calls to the Get method.
		Get []struct {
			// ID is the id argument value.
			ID int
		}
		// Save holds details about calls to the Save method.
		Save []struct {
			// In1 is the in1 argument value.
			In1 Ride
		}
	}
}

// Get calls GetFunc.
func (mock *RideRepositoryMock) Get(id int) (Ride, error) {
	if mock.GetFunc == nil {
		panic("RideRepositoryMock.GetFunc: method is nil but RideRepository.Get was just called")
	}
	callInfo := struct {
		ID int
	}{
		ID: id,
	}
	lockRideRepositoryMockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	lockRideRepositoryMockGet.Unlock()
	return mock.GetFunc(id)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedRideRepository.GetCalls())
func (mock *RideRepositoryMock) GetCalls() []struct {
	ID int
} {
	var calls []struct {
		ID int
	}
	lockRideRepositoryMockGet.RLock()
	calls = mock.calls.Get
	lockRideRepositoryMockGet.RUnlock()
	return calls
}

// Save calls SaveFunc.
func (mock *RideRepositoryMock) Save(in1 Ride) error {
	if mock.SaveFunc == nil {
		panic("RideRepositoryMock.SaveFunc: method is nil but RideRepository.Save was just called")
	}
	callInfo := struct {
		In1 Ride
	}{
		In1: in1,
	}
	lockRideRepositoryMockSave.Lock()
	mock.calls.Save = append(mock.calls.Save, callInfo)
	lockRideRepositoryMockSave.Unlock()
	return mock.SaveFunc(in1)
}

// SaveCalls gets all the calls that were made to Save.
// Check the length with:
//     len(mockedRideRepository.SaveCalls())
func (mock *RideRepositoryMock) SaveCalls() []struct {
	In1 Ride
} {
	var calls []struct {
		In1 Ride
	}
	lockRideRepositoryMockSave.RLock()
	calls = mock.calls.Save
	lockRideRepositoryMockSave.RUnlock()
	return calls
}
