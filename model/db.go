package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func DBConnect()(db *sql.DB) {

	if os.Getenv("CLEARDB_DATABASE_URL") != "" {
		dbDriver:="mysql"
		dbUser:=os.Getenv("DB_USERNAME")
		dbPass:=os.Getenv("DB_PASSWORD")
		dbName:=os.Getenv("DB_NAME")
		Dbhostname:=os.Getenv("DB_HOSTNAME")
		dboption:="?parseTime=true"
		db,err:=sql.Open(dbDriver,dbUser+":"+dbPass+"@tcp("+Dbhostname+":3306)/"+dbName+dboption)
		if err != nil {
			log.Fatal(err)
		}
		if err= db.Ping();err!=nil{
			log.Fatal(err)
		}
	}else {
		//local
		godotenv.Load(".env")
		dbDriver:="mysql"
		dbUser:=os.Getenv("DB_USERNAME")
		dbPass:=os.Getenv("DB_PASSWORD")
		dbName:=os.Getenv("DB_NAME")
		dboption:="?parseTime=true&loc=Asia%2FTokyo"
		dataSource:=dbUser+":"+dbPass+"@tcp(us-cdbr-east-04.cleardb.com:3306)/"+dbName+dboption
		db,err:=sql.Open(dbDriver,dataSource)
		if err != nil {
			log.Fatal(err)
		}
		if err= db.Ping();err!=nil{
			log.Fatal(err)
		}
	}
	return db
}