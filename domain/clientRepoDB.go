package domain

import (
	"database/sql"
	"time"

	"github.com/EricOgie/egobank/konstants"
	"github.com/EricOgie/egobank/mylogger"
	"github.com/EricOgie/egobank/reserrs"
	"github.com/jmoiron/sqlx"

	// github.com/go-sql-driver/mysql for mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// RepoDBMysql that will implement the ClientRepo interface
type RepoDBMysql struct {
	sqlDB *sqlx.DB
}

// FindAllClient implemtion on RepoDBMysql struct
func (dBConnection RepoDBMysql) FindAllClient(status string) ([]*Client, *reserrs.MyError) {

	var sqlQuery string
	var err error
	var clientList = make([]*Client, 0)
	if status == "" {
		sqlQuery = "SELECT * FROM customers"
		err = dBConnection.sqlDB.Select(&clientList, sqlQuery)
	} else {
		sqlQuery = "SELECT * FROM customers WHERE status = ?"
		err = dBConnection.sqlDB.Select(&clientList, sqlQuery, status)
	}

	if err != nil {
		mylogger.Error("Query Erro! error Msg: " + err.Error())
		return nil, reserrs.UnexpectedError(konstants.ServerError)
	}

	return clientList, nil
}

// ClientByID interface func implementation
func (dBConnection RepoDBMysql) ClientByID(id string) (*Client, *reserrs.MyError) {
	var c Client
	var err error
	sqlQuery := "SELECT * FROM customers WHERE customer_id = ?"
	err = dBConnection.sqlDB.Get(&c, sqlQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			mylogger.Error(konstants.QueryError + err.Error())
			return nil, reserrs.NotFoundError(konstants.NoClient)
		}
		mylogger.Error(konstants.ServerError + "Error Msg: " + err.Error())
		return nil, reserrs.UnexpectedError(konstants.ServerError)
	}

	return &c, nil
}

// CreateNewRepoDBMysql serving both as a helper function for RepoDBMysql and
// a Databse connection function
func CreateNewRepoDBMysql() RepoDBMysql {
	mysqlDatabase, err := sqlx.Open("mysql", konstants.DBCredentials)
	if err != nil {
		panic(err.Error())
	}

	mysqlDatabase.SetConnMaxLifetime(time.Minute * 3)
	mysqlDatabase.SetMaxOpenConns(10)
	mysqlDatabase.SetMaxIdleConns(10)

	return RepoDBMysql{mysqlDatabase}

}
