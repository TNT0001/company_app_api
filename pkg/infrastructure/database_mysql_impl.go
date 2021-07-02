package infrastructure

import (
	"errors"
	"fmt"
	"go-api/pkg/shared/gorm/database"
	"go-api/pkg/shared/utils"
	"go-api/pkg/shared/wraperror"

	"github.com/jinzhu/gorm"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	// import source file
	_ "github.com/golang-migrate/migrate/v4/source/file"

	// import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	// DBMaster set master database string.
	DBMaster = "master"
	// DBRead set read replica database string.
	DBRead = "read"
	// DBMS mysql type
	DBMS = "mysql"
	// LogModeKey stores logmode of gorm.DB
	LogModeKey = "logMode"
	// SourceKey stores source Master or Read
	SourceKey = "source"
	// NotDelete soft delete
	NotDelete = "deleted_datetime is NULL"

	// TransactionInvalidMsg transaction invalid message
	TransactionInvalidMsg = "Transaction Invalid"
)

// GormSQL struct
type GormSQL struct {
	// Master connections master database.
	Master *gorm.DB
	// Read connections read replica database.
	Read *gorm.DB
}

type dbInfo struct {
	host    string
	user    string
	pass    string
	name    string
	port    string
	charset string
	logmode bool
}

func newMySQLDatabase() (Database, error) {
	info := map[string]dbInfo{}
	info[DBMaster] = dbInfo{
		host:    utils.GetStringFlag(GetConfigString("flags.database.host")),
		user:    utils.GetStringFlag(GetConfigString("flags.database.user")),
		pass:    utils.GetStringFlag(GetConfigString("flags.database.pass")),
		name:    utils.GetStringFlag(GetConfigString("flags.database.name")),
		port:    utils.GetStringFlag(GetConfigString("flags.database.port")),
		charset: "utf8mb4",
		logmode: true,
	}
	info[DBRead] = dbInfo{
		host:    utils.GetStringFlag(GetConfigString("flags.database.host")),
		user:    utils.GetStringFlag(GetConfigString("flags.database.user")),
		pass:    utils.GetStringFlag(GetConfigString("flags.database.pass")),
		name:    utils.GetStringFlag(GetConfigString("flags.database.name")),
		port:    utils.GetStringFlag(GetConfigString("flags.database.port")),
		charset: "utf8mb4",
		logmode: true,
	}
	var master, read *gorm.DB

	for i, v := range info {
		// user:password@(localhost:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
		connect := v.user + ":" + v.pass + "@(" + v.host + ":" + v.port + ")/" + v.name + "?charset=" + v.charset + "&parseTime=True&loc=Local"
		db, err := gorm.Open(DBMS, connect)
		if err != nil {
			return nil, wraperror.WithTrace(err, wraperror.Fields{"host": v.host, "port": v.port, "dbname": v.name}, nil)
		}
		// db.SetLogger()
		db.LogMode(v.logmode)
		db.InstantSet(LogModeKey, v.logmode)
		db.InstantSet(SourceKey, i)
		// Disable table name's pluralization globally
		// if set this to true, `User`'s default table name will be `user`, table name setted with `TableName` won't be affected
		db.SingularTable(true)

		if i == DBMaster {
			master = db
		} else if i == DBRead {
			// can't create/update/delete read replica database.
			db.Callback().Create().Before("gorm:create").Register("read_create", database.CreateErrorCallback)
			db.Callback().Update().Before("gorm:update").Register("read_update", database.UpdateErrorCallback)
			db.Callback().Delete().Before("gorm:delete").Register("read_delete", database.DeleteErrorCallback)
			read = db
		}
	}
	return &GormSQL{Master: master, Read: read}, nil
}

// Create func
func (db *GormSQL) Create(value interface{}) error {
	err := db.Master.Create(value).Error
	return err
}

// Save func
func (db *GormSQL) Save(value interface{}) error {
	err := db.Master.Save(value).Error
	return err
}

// Find func
func (db *GormSQL) Find(condition interface{}, value interface{}) error {
	return db.Read.First(value, condition).Error
}

// FindAll func
func (db *GormSQL) FindAll(out interface{}, where ...interface{}) error {
	return db.Read.Find(out, where...).Error
}

// Update func
func (db *GormSQL) Update(deleteFilter bool, model interface{}, oldVal interface{}, newVal interface{}) error {
	if deleteFilter {
		return db.Master.Model(model).Where(NotDelete).Where(oldVal).Updates(newVal).Error
	}

	return db.Master.Model(model).Where(oldVal).Updates(newVal).Error
}

//Query func
func (db *GormSQL) Query(result interface{}, query string, args ...interface{}) error {
	err := db.Read.Raw(query, args...).Scan(result).Error
	return err
}

// CloseDB func
func (db *GormSQL) CloseDB() error {
	err := db.Master.Close()
	if err != nil {
		return err
	}

	err = db.Read.Close()
	if err != nil {
		return err
	}

	return nil
}

// IsRecordNotFoundError func
func (db *GormSQL) IsRecordNotFoundError(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}

// Begin func
func (db *GormSQL) Begin() interface{} {
	tx := db.Master.Begin()
	return tx
}

// Commit func
func (db *GormSQL) Commit(tx interface{}) error {
	mdb, ok := tx.(*gorm.DB)
	if !ok {
		return errors.New(TransactionInvalidMsg)
	}
	return mdb.Commit().Error
}

// Rollback func
func (db *GormSQL) Rollback(tx interface{}) error {
	mdb, ok := tx.(*gorm.DB)
	if !ok {
		return errors.New(TransactionInvalidMsg)
	}
	return mdb.Rollback().Error
}

// CreateWithTransaction func
func (db *GormSQL) CreateWithTransaction(tx, value interface{}) error {
	mdb, ok := tx.(*gorm.DB)
	if !ok {
		return errors.New(TransactionInvalidMsg)
	}
	return mdb.Create(value).Error
}

// UpdateWithTransaction func
func (db *GormSQL) UpdateWithTransaction(tx interface{}, deleteFilter bool, model interface{}, oldVal interface{}, newVal interface{}) error {
	mdb, ok := tx.(*gorm.DB)
	if !ok {
		return errors.New(TransactionInvalidMsg)
	}
	if deleteFilter {
		return mdb.Model(model).Where(NotDelete).Where(oldVal).Updates(newVal).Error
	}
	return mdb.Model(model).Where(oldVal).Updates(newVal).Error
}

// MigrationDB func
func (db *GormSQL) MigrationDB() error {
	driver, err := mysql.WithInstance(db.Master.DB(), &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///"+utils.GetStringFlag("dirMigration"),
		"mysql", driver)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// return m.Force(-1)
	err = m.Up()
	if err != nil {
		currentVer, dirty, _ := m.Version()
		if dirty {
			if currentVer == 1 {
				_ = m.Force(-1)
			} else {
				_ = m.Force(int(currentVer) - 1)
			}
		}
		return err
	}

	return nil
}
