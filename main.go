package main

import (
	"database/sql"
	"fmt"
	"log"

	"go_error/dao"
	"go_error/service"

	"github.com/pkg/errors"
)

const dsn string = ""

func main() {

	//新建Dao层
	db, err := dao.NewSqlServerDB(dsn)
	if err != nil {
		panic(err)

	}

	//新建Service层
	srv := service.NewService(db)

	//查找指定用户
	_, err = srv.GetUser("test")
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("This user was not found")
		log.Fatalf("%+v\n", err)

		return

	}
	if err != nil {
		fmt.Println(err)
		log.Fatalf("%+v\n", err)

		return

	}

	//........................................................................
}
