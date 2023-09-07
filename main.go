package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	conn "github.com/kaleabbyh/Food_Recipie/config"
	"github.com/kaleabbyh/Food_Recipie/middleware"
	"github.com/kaleabbyh/Food_Recipie/resolvers"
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

			"getUser": resolvers.GetUser(db),
			"getUsers":resolvers.GetUsers(db),
			"getCategories":resolvers.GetCategories(db),
			"getIngredients":resolvers.GetIngredients(db),
			"getRecipes":resolvers.GetRecipes(db),
			"getRecipe":resolvers.GetRecipe(db),
			
		},
	})

	//root mutation
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{

			"createUser":resolvers.CreateUser(db),
			"login"		:resolvers.Login(db),
			"updateUser":resolvers.UpdateUser(db),
			"deleteUser": resolvers.DeleteUser(db),
			"createCategory":resolvers.CreateCategory(db),
			"createIngredient":resolvers.CreateIngredient(db),
			"createRecipe":resolvers.CreateRecipe(db),
			"deleteRecipe":resolvers.DeleteRecipe(db), 
			"updateRecipe":resolvers.UpdateRecipe(db),  
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
	

	authHandler := middleware.AuthMiddleware(h)

	http.Handle("/graphql", authHandler)
	http.ListenAndServe(":8080", nil)


}