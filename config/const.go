package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const ServerPort = "4040"

//
const MysqlPort = "3306"


//evn
var MysqlPassword string
var MysqlMasterEndpoint string
var MysqlSlaveEndpoints [50]string
var MysqlSlaveCount int

//temp variable
var boolVal bool
var slaveCountTemp string

func InitEnvironmentVariables() error {

	//DB CREDENTIALS + Cluster Endpoint + Others
	MysqlPassword, boolVal = os.LookupEnv("MYSQL_PASSWORD")
	if boolVal == false {
		return errors.New("MYSQL_PASSWORD not found in envVars")
	}
	MysqlMasterEndpoint, boolVal = os.LookupEnv("MASTER_ENDPOINT")
	if boolVal == false {
		return errors.New("MASTER_ENDPOINT not found in envVars")
	}
	err := initSlaveEndpoints()
	if err != nil {
		return err
	}
	fmt.Println("environment vars loaded")
	return nil
}

func initSlaveEndpoints() error {
	slaveCountTemp, boolVal = os.LookupEnv("SLAVE_COUNT")
	if boolVal == true {
		var err error
		MysqlSlaveCount, err = strconv.Atoi(slaveCountTemp)
		if err != nil {
			return err
		}
		if MysqlSlaveCount < 0 || MysqlSlaveCount > 50 {
			return errors.New("invalid slave number: " + slaveCountTemp)
		}
	} else {
		MysqlSlaveCount= 0
		return nil
	}
	for i := 0; i < MysqlSlaveCount; i++ {
		MysqlSlaveEndpoints[i], boolVal = os.LookupEnv("SLAVE_ENDPOINT_" + strconv.Itoa(i))
		if boolVal == false {
			return errors.New("SLAVE_ENDPOINT_" + strconv.Itoa(i) + " not found in envVars")
		}
	}
	return nil
}
