package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB
func DBConnect() *sql.DB{

	/*if os.Getenv("CLEARDB_DATABASE_URL") != "" {
		dbDriver:="mysql"
		dbUser:=os.Getenv("DB_USERNAME")
		dbPass:=os.Getenv("DB_PASSWORD")
		dbName:=os.Getenv("DB_NAME")
		Dbhostname:=os.Getenv("DB_HOSTNAME")
		dboption:="?parseTime=true"
		_db,err:=sql.Open(dbDriver,dbUser+":"+dbPass+"@tcp("+Dbhostname+":3306)/"+dbName+dboption)
		if err != nil {
			log.Fatal(err)
		}
		if err= db.Ping();err==nil{
			log.Println("1success")
		}else{
			log.Println("fail")
		}
		db=_db
	}*/
		//local

		dbDriver:="mysql"
		dbUser:=os.Getenv("DB_LOCALUSERNAME")
		dbPass:=os.Getenv("DB_LOCALPASSWORD")
		dbName:=os.Getenv("DB_LOCALNAME")
		dbOption:="?parseTime=true&loc=Asia%2FTokyo"
		dataSource:=dbUser+":"+dbPass+"@tcp(us-cdbr-east-04.cleardb.com:3306)/"+dbName+dbOption
		_db,err:=sql.Open(dbDriver,dataSource)
		if err != nil {
			log.Fatal(err)
		}
		if err= db.Ping();err==nil{
			log.Println("2success")
		}else{
			log.Println("fail")
		}
		db=_db
	if err != nil {
		log.Fatal(err)
	}
	db,err=sql.Open(dbDriver,os.Getenv("DB_LOCALUSERNAME")+":"+os.Getenv("DB_LOCALPASSWORD")+"@/"+os.Getenv("DB_LOCALNAME")+dbOption)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")
	return db

	return db
}