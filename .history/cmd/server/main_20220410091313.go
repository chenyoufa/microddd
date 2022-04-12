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
		Debug:       false,
	})
	if err != nil {
		fmt.Printf("err fail:%v", err)
	}
	// p := d.Customer{}
	// p.Paaaff = 222
	roduct := infrasrepo.MemberDomain{}
	uuid, _ := uuid.Parse("b4fcb682-f023-419f-9c5b-f67124da7256")
	list, err := roduct.Get(context.Background(), uuid)
	if err != nil {
		log.Printf("err:%v", err)
	}
	fmt.Printf("item:%v", list)

}
