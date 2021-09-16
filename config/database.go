package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	MysqlUsersUsername string
	MysqlUsersPassword string
	MysqlUsersHost     string
	MysqlUsersPort     uint64
	MysqlUsersSchema   string

	//ConnectionString string to connect with the database
	UserConnectionString string
)

func ChargeDatabase() {
	var err error
	if godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	MysqlUsersUsername = os.Getenv("MYSQL_USERS_USERNAME")
	MysqlUsersPassword = os.Getenv("MYSQL_USERS_PASSWORD")
	MysqlUsersHost = os.Getenv("MYSQL_USERS_HOST")
	MysqlUsersSchema = os.Getenv("MYSQL_USERS_SCHEMA")

	MysqlUsersPort, err = strconv.ParseUint(os.Getenv("MYSQL_USERS_PORT"), 10, 64)
	if err != nil {
		MysqlUsersPort = 3306
	}

	UserConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		MysqlUsersUsername, MysqlUsersPassword, MysqlUsersHost, MysqlUsersPort, MysqlUsersSchema,
	)
}
