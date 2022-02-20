package factory

import (
	"github.com/Express-24/courier-location-tracker/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dbInstance DB

const (
	ErrConnHandlerNil = DBError("Cannot call non provided connHandler property of DBWrapper")
	ErrDBConfigNil    = DBError("DBWrapper's conf property is nil")
)

type DB interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) error
	Connect() error
}

func GetDBInstance() DB {
	if dbInstance == nil {
		Lock.Lock()
		defer Lock.Unlock()

		if dbInstance == nil {
			conf := config.GetConfig().Database
			db := DBWrapper{conf: &conf, connProvide: sqlx.Open}

			err := HandleConnection(&db)
			if err != nil {
				panic(err)
			}

			dbInstance = &db
		}
	}

	return dbInstance
}

func HandleConnection(db DB) error {
	err := db.Connect()
	if err != nil {
		return err
	}

	return nil
}

type DBWrapper struct {
	conf        *config.Database
	conn        *sqlx.DB
	connProvide func(driverName, dataSourceName string) (*sqlx.DB, error)
}

func (wrapper *DBWrapper) Connect() error {
	var err error
	switch {
	case wrapper.conf == nil:
		err = ErrDBConfigNil
	case wrapper.connProvide == nil:
		err = ErrConnHandlerNil
	default:
		err = nil
	}

	if err != nil {
		return err
	}

	wrapper.conn, err = wrapper.connProvide(wrapper.conf.Dialect, wrapper.conf.GetDsn())
	wrapper.conn.SetMaxOpenConns(wrapper.conf.MaxConnections)

	err = wrapper.conn.Ping()

	return err
}

func (wrapper *DBWrapper) Select(dest interface{}, query string, args ...interface{}) error {
	return wrapper.conn.Select(dest, query, args...)
}

func (wrapper *DBWrapper) Get(dest interface{}, query string, args ...interface{}) error {
	return wrapper.conn.Get(dest, query, args...)
}

func (wrapper *DBWrapper) Exec(query string, args ...interface{}) error {
	_, err := wrapper.conn.Exec(query, args...)
	return err
}

type DBError string

func (e DBError) Error() string {
	return string(e)
}
