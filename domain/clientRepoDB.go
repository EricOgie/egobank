package domain

import (
	"database/sql"
	"log"
	"time"

	"github.com/EricOgie/egobank/konstants"
	"github.com/EricOgie/egobank/reserrs"

	// github.com/go-sql-driver/mysql for mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// RepoDBMysql that will implement the ClientRepo interface
type RepoDBMysql struct {
	sqlDB *sql.DB
}

// FindAllClient implemtion on RepoDBMysql struct
func (dBConnection RepoDBMysql) FindAllClient(status string) ([]*Client, *reserrs.MyError) {

	var sqlQuery string
	if status == "" {
		sqlQuery = "SELECT * FROM customers"
	} else {
		sqlQuery = "SELECT * FROM customers WHERE status = " + status
	}

	rows, err := dBConnection.sqlDB.Query(sqlQuery)

	if err != nil {
		log.Println("Query Erro! error Msg: " + err.Error())
		return nil, reserrs.UnexpectedError(konstants.ServerError)
	}

	var clientList = make([]*Client, 0)

	for rows.Next() {
		var c Client
		scanErr := rows.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)
		if scanErr != nil {
			log.Println("Scanning Error! Error MSg:  " + scanErr.Error())
			return nil, reserrs.NotFoundError("Could not Find any resource")
		}

		clientList = append(clientList, &c)
	}
	return clientList, nil
}

// ClientByID interface func implementation
func (dBConnection RepoDBMysql) ClientByID(id string) (*Client, *reserrs.MyError) {
	sqlQuery := "SELECT * FROM customers WHERE customer_id = ?"
	row := dBConnection.sqlDB.QueryRow(sqlQuery, id)
	var c Client
	err := row.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Scanning Error Noticed! Error Msg: " + err.Error())
			return nil, reserrs.NotFoundError("Client with id, " + id + " does not exit in our Record")
		}
		log.Println("Unexpected Server Error! Error Msg: " + err.Error())
		return nil, reserrs.UnexpectedError("Unexpected Server Error")
	}

	return &c, nil
}

// CreateNewRepoDBMysql serving both as a helper function for RepoDBMysql and
// a Databse connection function
func CreateNewRepoDBMysql() RepoDBMysql {
	mysqlDatabase, err := sql.Open("mysql", "root@tcp(localhost)/banking")
	if err != nil {
		panic(err.Error())
	}

	mysqlDatabase.SetConnMaxLifetime(time.Minute * 3)
	mysqlDatabase.SetMaxOpenConns(10)
	mysqlDatabase.SetMaxIdleConns(10)

	return RepoDBMysql{mysqlDatabase}

}
