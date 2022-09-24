package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/swiggy-private/grill/pkg/grillmysql"
)

func TestMain(t *testing.T) {
	mysql := grillmysql.Mysql{
		Version: "8.0",
	}
	ctx := context.Background()
	err := mysql.Start(ctx)
	handleErr(err)
	stb := mysql.CreateTable("create table grill_test ( id int(11), name varchar(255))")
	err = stb.Stub()
	handleErr(err)
	conn, err := mysql.Client().Conn(ctx)
	handleErr(err)
	res, err := conn.ExecContext(ctx, "insert into grill_test values (1, 'shriprasad')")
	handleErr(err)
	cnt, err := res.RowsAffected()
	handleErr(err)
	fmt.Printf("inserted %d records\n", cnt)
	stmt, err := conn.PrepareContext(ctx, "select id as Id , name as Name from grill_test")
	handleErr(err)
	rows, err := stmt.Query()
	handleErr(err)
	for rows.Next() {
		obj := DbObject{}
		err = rows.Scan(&obj.Id, &obj.Name)
		handleErr(err)
		fmt.Println(obj)
	}
	conn.Close()
	mysql.Stop()
	main()
}

type DbObject struct {
	Id   int
	Name string
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
