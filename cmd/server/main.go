package main

import (
	"context"
	"fmt"
	"gomicroddd/infrastructure/db/dbcore"
	infrasrepo "gomicroddd/infrastructure/repository"
	"log"
)

func main() {

	err := dbcore.Connect(&dbcore.DBConfig{
		DbType:      "mysql",
		DSN:         "root:501526@(127.0.0.1:3306)/go-demo",
		AutoMigrate: true,
	})
	if err != nil {
		fmt.Printf("err fail:%v", err)
	}
	roduct := infrasrepo.MemoryProductRepository()
	list, err := roduct.GetAll(context.Background())
	if err != nil {
		log.Printf("err:%v", err)
	}
	for item := range list {
		fmt.Printf("item:%v", item)
	}

}
