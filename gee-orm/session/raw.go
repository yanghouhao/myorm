package session

import (
	"database/sql"
	"geeorm/log"
	"strings"
)

//Session is a struct that contract with db
type Session struct {
	db      *sql.DB
	sql     strings.Builder
	sqlVars []interface{}
}

//New create a nre session
func New(db *sql.DB) *Session {
	return &Session{db: db}
}

//Clear clear the information on session
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

//DB returns a db object
func (s *Session) DB() *sql.DB {
	return s.db
}

//Raw changes the value of variable
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

//Exec raw sql with sqlVars
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

//QueryRow gets a record from db
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

//QueryRows get a list of records from db
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
