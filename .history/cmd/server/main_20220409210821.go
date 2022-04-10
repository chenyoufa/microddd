package main

import (
	"context"
	"fmt"
	"gomicroddd/infrastructure/db/dbcore"
	infrasrepo "gomicroddd/infrastructure/repository"
	"log"

	"github.com/google/uuid"
)

func main() {

	err := dbcore.Connect(&dbcore.DBConfig{
		DbType:      "mysql",
		DSN:         "fage:Fage501526~@(127.0.0.1:3306)/mytest",
		AutoMigrate: true,
	})
	if err != nil {
		fmt.Printf("err fail:%v", err)
	}
	// p := d.Customer{}
	// p.Paaaff = 222
	roduct := infrasrepo.MemberDomain{}
	list, err := roduct.Get(context.Background(), uuid.New())
	if err != nil {
		log.Printf("err:%v", err)
	}
	fmt.Printf("item:%v", list)

}
