package mysql

import(
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)
var Db *sqlx.DB
func init(){
	sqlx.Open("", "")
}