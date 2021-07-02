package test

import (
	"github.com/stretchr/testify/mock"
)

// DatabaseMock struct
type DatabaseMock struct {
	mock.Mock
}

// Create func
func (db *DatabaseMock) Create(value interface{}) error {
	args := db.Called(value)
	return args.Error(0)
}

// Save func
func (db *DatabaseMock) Save(value interface{}) error {
	args := db.Called(value)
	return args.Error(0)
}

// FindAll func
func (db *DatabaseMock) FindAll(out interface{}, where ...interface{}) error {
	args := db.Called(out, where)
	return args.Error(0)
}

// Find func
func (db *DatabaseMock) Find(condition interface{}, value interface{}) error {
	args := db.Called(condition, value)
	return args.Error(0)
}

// Query func
func (db *DatabaseMock) Query(result interface{}, query string, args ...interface{}) error {
	a := db.Called(result, query, args)
	return a.Error(0)
}

// Update func
func (db *DatabaseMock) Update(deleteFilter bool, model interface{}, oldVal interface{}, newVal interface{}) error {
	args := db.Called(deleteFilter, model, oldVal, newVal)
	return args.Error(0)
}

// Count func
func (db *DatabaseMock) Count(deleteFilter bool, model interface{}, condition1 interface{}, condition2 interface{}, conditionVal2 interface{}) int {
	args := db.Called(deleteFilter, model, condition1, condition2, conditionVal2)
	return args.Int(0)
}

// IsRecordNotFoundError func
func (db *DatabaseMock) IsRecordNotFoundError(err error) bool {
	args := db.Called(err)
	return args.Bool(0)
}

// CloseDB func
func (db *DatabaseMock) CloseDB() error {
	args := db.Called()
	return args.Error(0)
}

// Begin func
func (db *DatabaseMock) Begin() interface{} {
	args := db.Called()
	return args.Get(0)
}

// Commit func
func (db *DatabaseMock) Commit(tx interface{}) error {
	args := db.Called(tx)
	return args.Error(0)
}

// Rollback func
func (db *DatabaseMock) Rollback(interface{}) error {
	args := db.Called()
	return args.Error(0)
}

// CreateWithTransaction func
func (db *DatabaseMock) CreateWithTransaction(tx, value interface{}) error {
	args := db.Called(tx, value)
	return args.Error(0)
}

// UpdateWithTransaction func
func (db *DatabaseMock) UpdateWithTransaction(tx interface{}, deleteFilter bool, model interface{}, oldVal interface{}, newVal interface{}) error {
	args := db.Called(tx, deleteFilter, model, oldVal, newVal)
	return args.Error(0)
}

// MigrationDB func
func (db *DatabaseMock) MigrationDB() error {
	args := db.Called()
	return args.Error(0)
}
