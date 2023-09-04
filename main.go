package main

import (
	"fmt"
	"net/http"
	conn "sample/config"
	"sample/resolvers"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/lib/pq"
)



func main() {
	//DB connection

	db,_:=conn.ConnectDB()
	defer db.Close()

	//root query
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{

			// User query
			"user": resolvers.GetUser(db),
			"users":resolvers.GetUsers(db),
			
		},
	})

	//root mutation
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{

			// User mutation
			"createUser":resolvers.CreateUser(db),
			"updateUser":resolvers.UpdateUser(db),
			"deleteUser": resolvers.DeleteUser(db),
		},
	})


	//schema
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// serve HTTP
	fmt.Println("connected successfully")
	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}