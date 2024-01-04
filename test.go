package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/shinhagunn/fastfood/postgresql"
)

func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	querier := postgresql.New(db)

	resultGetUsers, err := querier.GetUsers(ctx)
	if err != nil {
		panic(err)
	}
	log.Println("Result get users: ", resultGetUsers)

	resultGetUserByID, err := querier.GetUserByID(ctx, 1)
	if err != nil {
		panic(err)
	}
	log.Println("Result get user: ", resultGetUserByID)

	err = querier.CreateUser(ctx, "test")
	if err != nil {
		panic(err)
	}

	err = querier.UpdateUser(ctx, postgresql.UpdateUserParams{
		ID:   1,
		Name: "test2",
	})
	if err != nil {
		panic(err)
	}

	// err = querier.DeleteUser(ctx, 1)
	// if err != nil {
	// 	panic(err)
	// }

	log.Println(db)
}
