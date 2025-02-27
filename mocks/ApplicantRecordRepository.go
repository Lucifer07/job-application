// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/dto"
	mock "github.com/stretchr/testify/mock"
)

// ApplicantRecordRepository is an autogenerated mock type for the ApplicantRecordRepository type
type ApplicantRecordRepository struct {
	mock.Mock
}

// AppliedJob provides a mock function with given fields: ctx, record
func (_m *ApplicantRecordRepository) AppliedJob(ctx context.Context, record dto.AppliedReq) (*int, error) {
	ret := _m.Called(ctx, record)

	var r0 *int
	if rf, ok := ret.Get(0).(func(context.Context, dto.AppliedReq) *int); ok {
		r0 = rf(ctx, record)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dto.AppliedReq) error); ok {
		r1 = rf(ctx, record)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAppliedJob provides a mock function with given fields: ctx, jobId, userId
func (_m *ApplicantRecordRepository) FindAppliedJob(ctx context.Context, jobId int, userId int) (*dto.ApplicantRecord, error) {
	ret := _m.Called(ctx, jobId, userId)

	var r0 *dto.ApplicantRecord
	if rf, ok := ret.Get(0).(func(context.Context, int, int) *dto.ApplicantRecord); ok {
		r0 = rf(ctx, jobId, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ApplicantRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, jobId, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
