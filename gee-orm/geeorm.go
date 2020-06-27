package geeorm

import (
	"database/sql"
	"geeorm/log"
	"geeorm/session"
)

//Engine prepares db visiting
type Engine struct {
	db *sql.DB
}

//NewEngine opens the db
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}

	e = &Engine{db: db}
	log.Info("Connect database success")
	return
}

//Close close the database
func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

//NewSession create a session object
func (e *Engine) NewSession() *session.Session {
	return session.New(e.db)
}
