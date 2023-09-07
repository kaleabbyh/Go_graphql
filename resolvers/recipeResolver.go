package resolvers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
	models "github.com/kaleabbyh/Food_Recipie/model"
	"github.com/kaleabbyh/Food_Recipie/utils"
	_ "github.com/lib/pq"
)



var recipeType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "recipe",
	Description: "a recipe",
	Fields: graphql.Fields{
		"recipe_id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The identifier of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.CreateRecipeType); ok {
					return recipe.Recipe_id, nil
				}

				return nil, nil
			},
		},
		"recipe_title": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The title of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.CreateRecipeType); ok {
					return recipe.Recipe_title, nil
				}

				return nil, nil
			},
		},

		"instructions": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The Instructions of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.CreateRecipeType); ok {
					return recipe.Instructions, nil
				}

				return nil, nil
			},
		},
		"preparation_time": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Float),
			Description: "The preparation_time of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.CreateRecipeType); ok {
					return recipe.Preparation_time, nil
				}

				return nil, nil
			},
		},

		"cooking_time": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Float),
			Description: "The Cooking_time of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.CreateRecipeType); ok {
					return recipe.Cooking_time, nil
				}

				return nil, nil
			},
		},

		"user_id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The User_id of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.CreateRecipeType); ok {
					return recipe.User_id, nil
				}

				return nil, nil
			},
		},

		"category_id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The Category_id of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.CreateRecipeType); ok {
					return recipe.Category_id, nil
				}

				return nil, nil
			},
		},
		
		"user": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The User of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.CreateRecipeType); ok {
					return recipe.User, nil
				}

				return nil, nil
			},
		},

		"category": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The Category of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.CreateRecipeType); ok {
					return recipe.Category, nil
				}

				return nil, nil
			},
		},
		

		"created_at": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The created_at date of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.CreateRecipeType); ok {
					return recipe.Created_at, nil
				}

				return nil, nil
			},
		},

		"updated_at": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The updated_at date of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.CreateRecipeType); ok {
					return recipe.Updated_at, nil
				}

				return nil, nil
			},
		},
		

		
	},
})


//create recipe
func CreateRecipe(db *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        recipeType,
		Description: "Create new recipe",
		Args: graphql.FieldConfigArgument{
	
			"recipe_title" : &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"instructions" : &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"preparation_time" : &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
			"cooking_time" : &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
			"user_id"      : &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"category_id"   : &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			recipe_title, _ 	:= params.Args["recipe_title"].(string)
			instructions, _ 	:= params.Args["instructions"].(string)
			preparation_time, _ := params.Args["preparation_time"].(float64)
			cooking_time,_ 		:= params.Args["cooking_time"].(float64)
			user_id,_ 			:= params.Args["user_id"].(int)
			category_id,_ 		:= params.Args["category_id"].(int)
			created_at 			:= time.Now()
			updated_at 			:= time.Now()
			
			

			var lastInsertId int
			err := db.QueryRow(
					`INSERT INTO recipes(recipe_title,instructions,preparation_time,cooking_time, user_id,
					category_id, created_at,updated_at) 
					VALUES($1, $2, $3,$4,$5,$6,$7,$8) returning recipe_id;`,
			recipe_title,instructions,preparation_time, cooking_time,user_id,category_id, created_at,updated_at).
			Scan(&lastInsertId)
			utils.CheckErr(err)
			

			user := &models.RegisterUser{}
			err = db.QueryRow("select name from users where id = $1", user_id).Scan(&user.Name)
			utils.CheckErr(err)

			category := &models.Category{}
			err = db.QueryRow("select category_name from categories where category_id = $1", category_id).Scan(&category.Category_name)
			utils.CheckErr(err)
			

			


			newRecipe := &models.CreateRecipeType{
				Recipe_id		:lastInsertId,
				Recipe_title	:recipe_title,
				Instructions	:instructions,
				Preparation_time:preparation_time,
				Cooking_time	:cooking_time,
				User_id			:user_id,
				Category_id		:category_id,
				User			:user.Name,
				Category		:category.Category_name,
				Created_at		:created_at,
				Updated_at		:updated_at,
			}



			return newRecipe, nil
		},
	}
}



func GetRecipes(db *sql.DB) *graphql.Field {

	return  &graphql.Field{
		Type:        graphql.NewList(recipeType),
		Description: "List of recipes.",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			rows, err := db.Query(`SELECT recipe_id, recipe_title,instructions,preparation_time,cooking_time, user_id,
								    category_id, created_at,updated_at FROM recipes`)
			utils.CheckErr(err)
			var Recipes []*models.CreateRecipeType

			for rows.Next() {
				Recipe := &models.CreateRecipeType{}
				err = rows.Scan(&Recipe.Recipe_id, &Recipe.Recipe_title, &Recipe.Instructions,&Recipe.Preparation_time, 
								&Recipe.Cooking_time,&Recipe.User_id,&Recipe.Category_id,
								&Recipe.Created_at, &Recipe.Updated_at)
					 
				utils.CheckErr(err)
				

				Recipes = append(Recipes, Recipe)
			}


			var recipesWithUsername []*models.CreateRecipeType;
			for _, recipe := range Recipes {
				Recipe := &models.CreateRecipeType{}

				user := &models.RegisterUser{}
				err := db.QueryRow("select name from users where id = $1", recipe.User_id).Scan(&user.Name)
				utils.CheckErr(err)
				Recipe.User=user.Name
				

				category := &models.Category{}
				err = db.QueryRow("select category_name from categories where category_id = $1", recipe.Category_id).Scan(&category.Category_name)
				utils.CheckErr(err)
				Recipe.Category=category.Category_name

				recipesWithUsername = append(recipesWithUsername, Recipe)
				fmt.Println(Recipe.User)
			}

			return recipesWithUsername, nil
		},
	}
}



//get recipe
func GetRecipe(db *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        recipeType,
		Description: "Get a recipe.",
		Args: graphql.FieldConfigArgument{
			"recipe_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			recipe_id, _ := params.Args["recipe_id"].(int)

			Recipe := &models.CreateRecipeType{}
			err := db.QueryRow(`SELECT recipe_id, recipe_title,instructions,preparation_time,cooking_time, user_id,
								category_id, created_at,updated_at FROM recipes where recipe_id = $1`, recipe_id).
								Scan(&Recipe.Recipe_id, &Recipe.Recipe_title, &Recipe.Instructions,&Recipe.Preparation_time, 
									&Recipe.Cooking_time,&Recipe.User_id,&Recipe.Category_id,
									&Recipe.Created_at, &Recipe.Updated_at)
			utils.CheckErr(err)
			
			//middleware for authentication and authorization
			email, ok := params.Context.Value("email").(string)
			if !ok {
				fmt.Println(email)
				return nil, fmt.Errorf(" not authenticated")
			}
		    
			
			user := &models.RegisterUser{}
			err = db.QueryRow("select email from users where id = $1", Recipe.User_id).Scan(&user.Email)
			utils.CheckErr(err)
				

			if email != user.Email  {
				return nil, fmt.Errorf(" not authorized")
			}
			return Recipe, nil
		},
	}
}



//delete user
func DeleteRecipe(db *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        recipeType,
		Description: "Delete an recipe",
		Args: graphql.FieldConfigArgument{
			"recipe_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			recipe_id, _ := params.Args["recipe_id"].(int)

			Recipe := &models.CreateRecipeType{}
			err := db.QueryRow(`SELECT recipe_id,  user_id FROM recipes where recipe_id = $1`, recipe_id).
								Scan(&Recipe.Recipe_id, &Recipe.User_id)
			utils.CheckErr(err)


			user := &models.RegisterUser{}
			err = db.QueryRow("select id,  email from users where id = $1", Recipe.User_id).Scan(&user.ID, &user.Email)
			utils.CheckErr(err)
	

			//middlware
			email, ok := params.Context.Value("email").(string)
			if !ok {
				fmt.Println(email)
				return nil, fmt.Errorf(" not authenticated")
			}
			
			
			if email != user.Email  {
				return nil, fmt.Errorf(" not authorized")
			}

			stmt, err := db.Prepare("DELETE FROM recipes WHERE recipe_id = $1")
			utils.CheckErr(err)
	
			_, err2 := stmt.Exec(Recipe.Recipe_id)
			utils.CheckErr(err2)
	
			return nil, nil
		},
	}
	}





//create recipe
func UpdateRecipe(db *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        recipeType,
		Description: "Update new recipe",
		Args: graphql.FieldConfigArgument{
			"recipe_id" : &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"recipe_title" : &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"instructions" : &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"preparation_time" : &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
			"cooking_time" : &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
			
			"category_id"   : &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			recipe_id,_ 		:= params.Args["recipe_id"].(int)
			recipe_title, _ 	:= params.Args["recipe_title"].(string)
			instructions, _ 	:= params.Args["instructions"].(string)
			preparation_time, _ := params.Args["preparation_time"].(float64)
			cooking_time,_ 		:= params.Args["cooking_time"].(float64)
			category_id,_ 		:= params.Args["category_id"].(int)
			updated_at 			:= time.Now()
			
			

			user := &models.RegisterUser{}
			err := db.QueryRow("select user_id from recipes where recipe_id = $1", recipe_id).Scan(&user.ID)
			utils.CheckErr(err)

			err = db.QueryRow("select id,email from users where id = $1", user.ID).Scan(&user.ID,&user.Email)
			utils.CheckErr(err)

			category := &models.Category{}
			err = db.QueryRow("select category_name from categories where category_id = $1", category_id).Scan(&category.Category_name)
			utils.CheckErr(err)
			
			//middlware
			email, ok := params.Context.Value("email").(string)
			if !ok {
				fmt.Println(email)
				return nil, fmt.Errorf(" not authenticated")
			}


			if email != user.Email  {
				return nil, fmt.Errorf(" not authorized")
			}




			stmt, err := db.Prepare(`UPDATE recipes SET recipe_title = $1, instructions = $2, preparation_time = $3,
									cooking_time = $4, category_id = $5, updated_at = $6 WHERE recipe_id = $7`)
			utils.CheckErr(err)

			_, err2 := stmt.Exec(recipe_title,instructions,preparation_time, cooking_time,category_id,updated_at,recipe_id)
			utils.CheckErr(err2)
			


			newRecipe := &models.CreateRecipeType{
				Recipe_id		:recipe_id,
				Recipe_title	:recipe_title,
				Instructions	:instructions,
				Preparation_time:preparation_time,
				Cooking_time	:cooking_time,
				Category_id		:category_id,
				User			:user.Name,
				Category		:category.Category_name,
				Updated_at		:updated_at,
			}



			return newRecipe, nil
		},
	}
}