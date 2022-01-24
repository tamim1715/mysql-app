package main

import (
	"fmt"
	"log"

	"github.com/tamim1715/mysql-app/config"
	"github.com/tamim1715/mysql-app/database/mysql"
	"github.com/tamim1715/mysql-app/router"
	"github.com/tamim1715/mysql-app/server"
)

func main() {

	//init envVars
	//init db connections
	//init certificates
	srv := server.New()
	router.Routes(srv)
	
	// err1 := godotenv.Load("local.env")
	// if err1 != nil {
	// 	log.Fatalf("Some error occured. Err: %s", err1)
	// }

	err := config.InitEnvironmentVariables()
	if err != nil {
		log.Fatal("envVars error: " + err.Error())
	}
	err = mysql.InitMysqlMaster()
	if err != nil {
		log.Fatal("master endpoint error: " + err.Error())
	}
	err = mysql.InitMysqlSlave()
	if err != nil {
		log.Fatal("slave endpoint error: " + err.Error())
	}
	fmt.Println("Successfully connect with mysql database")	
	
	srv.Logger.Fatal(srv.Start(":" + config.ServerPort))
}
