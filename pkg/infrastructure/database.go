package infrastructure

import (
	"errors"
	"go-api/pkg/shared/wraperror"
)

const (
	// MySQLDatabase instance database
	MySQLDatabase = "mysql"

	// ErrRecordNotFound Error
	ErrRecordNotFound = "record not found"

	// ErrFromToDateTime Error
	ErrFromToDateTime = "Field.From great than Field.To"

	// ErrInvalidID invalid ID
	ErrInvalidID = "Invalid ID"

	// ErrPeriod Error
	ErrPeriod = "Invalid period"

	// ErrName not include symbol and number
	ErrName = "Field.Name not include symbol and number"

	// ErrBirthday is invalid
	ErrBirthday = "Field.Birthday is invalid"

	// ErrEmailIsInvalid email is invalid
	ErrEmailIsInvalid = "Email is invalid"

	// ErrEmailAlreadyExist already exist email
	ErrEmailAlreadyExist = "Email already exist"

	// ErrEmailAuthentication has sent verification link in email
	ErrEmailAuthentication = "Has sent verification link in email"

	// ErrVerifyTokenFail verify token fail
	ErrVerifyTokenFail = "Verify Token Fail"

	// ErrVerifyTokenIsExpried token is expired
	ErrVerifyTokenIsExpried = "Token is expired"

	// ErrTokenIsInvalid token is invalid
	ErrTokenIsInvalid = "Token is invalid"

	// ErrLoginFail login fail
	ErrLoginFail = "Email or password is invalid"

	// ErrPasswordInvalid new password has same old password
	ErrPasswordInvalid = "password is duplicated"

	// ErrCompanyNotFound not found company record in database
	ErrCompanyNotFound = "company not found"
)

// Database interface
type Database interface {
	Create(value interface{}) error
	Save(value interface{}) error
	Find(condition interface{}, value interface{}) error
	FindAll(out interface{}, where ...interface{}) error
	Query(result interface{}, query string, args ...interface{}) error
	Update(deleteFilter bool, model interface{}, oldVal interface{}, newVal interface{}) error
	IsRecordNotFoundError(err error) bool
	CloseDB() error
	Begin() interface{}
	Commit(interface{}) error
	Rollback(interface{}) error
	CreateWithTransaction(tx, value interface{}) error
	UpdateWithTransaction(tx interface{}, deleteFilter bool, model interface{}, oldVal interface{}, newVal interface{}) error
	MigrationDB() error
}

// NewDatabase returns an instance of database
func NewDatabase(dbInstance string) (Database, error) {
	switch dbInstance {
	case MySQLDatabase:
		db, err := newMySQLDatabase()
		if err != nil {
			return nil, err
		}
		return db, nil

	default:
		return nil, wraperror.WithTrace(errors.New("Invalid database instance"), wraperror.Fields{"dbInstance": dbInstance}, nil)
	}
}

// CloseDB close DB
func CloseDB(db Database) error {
	return db.CloseDB()
}

// MigrationDB func migration DB
func MigrationDB(db Database) error {
	return db.MigrationDB()
}
