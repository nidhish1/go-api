package dbCon

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

var (

	Db *sql.DB
	Err error
) 
