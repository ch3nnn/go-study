package main

import (
	"context"
	"ent/crud"
	"ent/ent"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// 运行自动迁移工具。
	//if err := client.Schema.Create(context.Background()); err != nil {
	//	log.Fatalf("failed creating schema resources: %v", err)
	//}

	ctx := context.Background()

	if _, err := crud.CreateUser(ctx, client); err != nil {
		log.Fatal(err)
	}

	if _, err := crud.QueryUser(ctx, client); err != nil {
		log.Fatal(err)

	}

	if _, err := crud.UpdateUser(ctx, client); err != nil {
		log.Fatal(err)

	}

	if _, err := crud.DeleteUser(ctx, client); err != nil {
		log.Fatal(err)

	}

}
