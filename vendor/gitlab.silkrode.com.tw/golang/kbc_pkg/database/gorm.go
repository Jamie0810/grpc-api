package database

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	// database driver for gorm
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
)

var supported = map[string]func(port uint, host, dbname, user, password, instanceName, connectTimeout, readTimeout, writeTimeout string, sslmode bool) gorm.Dialector{
	"cloudsql": func(port uint, host, dbname, user, password, instanceName, connectTimeout, readTimeout, writeTimeout string, sslmode bool) gorm.Dialector {
		return mysql.Open(fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?charset=utf8mb4&parseTime=true&loc=UTC&time_zone=UTC&timeout=%s&readTimeout=%s&writeTimeout=%s", user, password, instanceName, dbname, connectTimeout, readTimeout, writeTimeout))
	},
	"mysql": func(port uint, host, dbname, user, password, instanceName, connectTimeout, readTimeout, writeTimeout string, sslmode bool) gorm.Dialector {
		return mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=UTC&time_zone=UTC&timeout=%s&readTimeout=%s&writeTimeout=%s", user, password, host, port, dbname, connectTimeout, readTimeout, writeTimeout))
	},
	"postgres": func(port uint, host, dbname, user, password, instanceName, connectTimeout, readTimeout, writeTimeout string, sslmode bool) gorm.Dialector {
		ssl := "disable"
		if sslmode {
			ssl = "allow"
		}
		// TODO: 增加read,write,conn timeout
		return postgres.Open(fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s timezone=UTC", host, port, user, dbname, password, ssl))
	},
	"sqlite": func(port uint, host, dbname, user, password, instanceName, connectTimeout, readTimeout, writeTimeout string, sslmode bool) gorm.Dialector {
		return sqlite.Open(fmt.Sprintf("%s.db", dbname))
	},
}

// NewDB initialize a gorm.DB for further usage
func NewDB(
	driver, host string,
	port uint, dbname, instanceName string,
	user, password string, sslmode bool,
	connectTimeout, readTimeout, writeTimeout string,
	options ...Option,
) (*gorm.DB, error) {
	connectionFunc := supported[driver]
	if connectionFunc == nil {
		return nil, fmt.Errorf("not supported driver [%s]", driver)
	}

	sqlDriver := connectionFunc(
		port, host, dbname,
		user, password, instanceName,
		connectTimeout, readTimeout, writeTimeout, sslmode)

	engine, err := gorm.Open(sqlDriver, &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, err
	}
	for _, opt := range options {
		opt.Apply(engine)
	}
	return engine, nil
}

// An Option configures a gorm.DB
type Option interface {
	Apply(*gorm.DB)
}

// OptionFunc is a function that configures a gorm.DB
type OptionFunc func(*gorm.DB)

// Apply is a function that set value to gorm.DB
func (f OptionFunc) Apply(engine *gorm.DB) {
	f(engine)
}

func SetConnMaxLifetime(maxlifetime time.Duration) Option {
	return OptionFunc(func(engine *gorm.DB) {
		sql, err := engine.DB()
		if err != nil {
			return
		}
		sql.SetConnMaxLifetime(maxlifetime)
	})
}

func SetMaxIdleConns(maxIdleConns int) Option {
	return OptionFunc(func(engine *gorm.DB) {
		sql, err := engine.DB()
		if err != nil {
			return
		}
		sql.SetMaxIdleConns(maxIdleConns)
	})
}

func SetMaxOpenConns(maxOpenConns int) Option {
	return OptionFunc(func(engine *gorm.DB) {
		sql, err := engine.DB()
		if err != nil {
			return
		}
		sql.SetMaxOpenConns(maxOpenConns)
	})
}

func SetLogger(logger logger.Interface) Option {
	return OptionFunc(func(engine *gorm.DB) {
		engine.Logger = logger
	})
}
