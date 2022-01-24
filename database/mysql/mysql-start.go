package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/tamim1715/mysql-app/config"
)

var master *sql.DB
var slave [50]*sql.DB
var Err error
var count = 0
var mtx sync.Mutex

func InitMysqlMaster() error {
	master, Err = sql.Open("mysql", "root:"+config.MysqlPassword+"@tcp("+config.MysqlMasterEndpoint+":"+config.MysqlPort+")/klovercloud")
	if Err != nil {
		return errors.New("error connecting master")
	}
	fmt.Println("master instance loaded")
	return nil
}


func InitMysqlSlave() error {
	for i := 0; i < config.MysqlSlaveCount; i++ {
		
		slave[i], Err = sql.Open("mysql", "root:"+config.MysqlPassword+"@tcp("+config.MysqlSlaveEndpoints[i]+":"+config.MysqlPort+")/klovercloud")
		if Err != nil {
			return errors.New("error connecting slave: " + strconv.Itoa(i))
		}
	}
	fmt.Println("slave instance loaded: " + strconv.Itoa(config.MysqlSlaveCount))
	return nil
}

func GetMysqlMaster() *sql.DB{
	return master
}
func GetMysqlSlave() *sql.DB{
	fmt.Println("Count ==> ", count)
	if config.MysqlSlaveCount == 0 {
		return master
	}
	instance := slave[count]
	mtx.Lock()
	count++
	count = count % config.MysqlSlaveCount
	mtx.Unlock()
	return instance
}
