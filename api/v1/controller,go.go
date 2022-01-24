package v1

import "C"
import (
	"fmt"
	"net/http"

	"github.com/tamim1715/mysql-app/database/mysql"
	"github.com/tamim1715/mysql-app/dto"
	"github.com/tamim1715/mysql-app/helper"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
)

// var Db *sql.DB
// var Err error

type CacheControllerInf interface {
	Get(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type CacheControllerInstance struct {
}

func CacheController() CacheControllerInf {
	return new(CacheControllerInstance)
}
func (CacheControllerInstance) Create(c echo.Context) error {
	var value dto.Info
	err := c.Bind(&value)
	if err != nil {
		return err
	}

	err = helper.ValidateInput(value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
		
	}
	a, err := mysql.GetMysqlMaster().Query("select id from users where id="+value.ID+";")
	if err != nil{
		fmt.Println(err.Error())
	}

	var ID string
	for a.Next(){
		var id string
		err2 := a.Scan(&id)
		if err2 !=nil{
			fmt.Println(err2.Error())
		}else{
			ID = id
		}
	}

	if(ID==value.ID){
		c.JSON(http.StatusBadRequest, "Data Already Stored")
		return nil
	}
	defer a.Close()

	insert , err := mysql.GetMysqlMaster().Query("insert into users values('"+value.ID+"','"+value.Name+"','"+value.Designation+"','"+value.Branch+"');")
	if err !=nil{
		fmt.Println(err.Error())
	}
	defer insert.Close()

	fmt.Println("Successfully insert table value")
	
	c.JSON(http.StatusOK, "value set")
	return nil
}

func (CacheControllerInstance) Update(c echo.Context) error {
	var value dto.Info
	err := c.Bind(&value)
	if err != nil {
		return err
	}

	err = helper.ValidateInput(value)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}
	a, err :=mysql.GetMysqlMaster().Query("select id from users where id="+value.ID+";")
	if err != nil{
		fmt.Println(err.Error())
	}

	var ID string
	for a.Next(){
		var id string
		err2 := a.Scan(&id)
		if err2 !=nil{
			fmt.Println(err2.Error())
		}else{
			ID = id
		}
	}

	if(ID!=value.ID){
		c.JSON(http.StatusBadRequest, "Data Not Found")
		return nil
	}
	defer a.Close()

	fmt.Println("Successfully found select query")

	update, err := mysql.GetMysqlMaster().Query("update users set name='"+value.Name+"' , designation = '"+value.Designation+"', branch = '"+value.Branch+"'  where id = "+value.ID+";")

	if err != nil { 
		c.JSON(http.StatusBadRequest, "error when update value")
		return nil
	}
	defer update.Close()
	c.JSON(http.StatusOK, "Successfully update vlaue")
	return nil
}

func (CacheControllerInstance) Delete(c echo.Context) error {
	ID := c.Param("id")
	
	a, err := mysql.GetMysqlMaster().Query("select id from users where id="+ID+";")
	if err != nil{
		fmt.Println(err.Error())
	}

	var ID1 string
	for a.Next(){
		var id string
		err2 := a.Scan(&id)
		if err2 !=nil{
			fmt.Println(err2.Error())
		}else{
			ID1 = id
		}
	}

	if(ID!=ID1){
		c.JSON(http.StatusBadRequest, "Data Not Found")
		return nil
	}
	defer a.Close()

	delete, err := mysql.GetMysqlMaster().Query("delete from users where id="+ID+";")
	defer delete.Close()
	c.JSON(http.StatusOK, "Successfully Delete Your value")
	return nil
}
func (CacheControllerInstance) Get(c echo.Context) error {

	id := c.Param("id")
	a, err := mysql.GetMysqlSlave().Query("select* from users where id="+id+";")
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println(a)
	value := dto.Info{}

	for a.Next(){
		var id string
		var name string
		var designation string
		var branch string 
		err2 := a.Scan(&id, &name, &designation, &branch)
		if err2 !=nil{
			fmt.Println(err2.Error())
		}else{
			value = dto.Info{id, name, designation, branch}
		}
	}

	if value.ID != ""{
		c.JSON(http.StatusOK, value)
	}else{
		c.JSON(http.StatusBadRequest, "Data Not Found")
	}
	return nil
}
