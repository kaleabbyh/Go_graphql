package resolvers

import (
	"database/sql"
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
				if recipe, ok := p.Source.(*models.Recipe); ok {
					return recipe.Recipe_id, nil
				}

				return nil, nil
			},
		},
		"recipe_title": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The title of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.Recipe); ok {
					return recipe.Recipe_title, nil
				}

				return nil, nil
			},
		},

		"instructions": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The Instructions of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.Recipe); ok {
					return recipe.Instructions, nil
				}

				return nil, nil
			},
		},
		"preparation_time": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Float),
			Description: "The preparation_time of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.Recipe); ok {
					return recipe.Preparation_time, nil
				}

				return nil, nil
			},
		},

		"cooking_time": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Float),
			Description: "The Cooking_time of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.Recipe); ok {
					return recipe.Cooking_time, nil
				}

				return nil, nil
			},
		},

		"user_id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The User_id of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.Recipe); ok {
					return recipe.User_id, nil
				}

				return nil, nil
			},
		},

		"category_id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The Category_id of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.Recipe); ok {
					return recipe.Category_id, nil
				}

				return nil, nil
			},
		},
		

		"created_at": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The created_at date of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.Recipe); ok {
					return recipe.Created_at, nil
				}

				return nil, nil
			},
		},

		"updated_at": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The updated_at date of the recipe.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if recipe, ok := p.Source.(*models.Recipe); ok {
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
			


			newRecipe := &models.Recipe{
				Recipe_id		:lastInsertId,
				Recipe_title	:recipe_title,
				Instructions	:instructions,
				Preparation_time:preparation_time,
				Cooking_time	:cooking_time,
				User_id			:user_id,
				Category_id		:category_id,
				Created_at		:created_at,
				Updated_at		:updated_at,
			}

			return newRecipe, nil
		},
	}
}




//get users
// func GetRecipes(db *sql.DB) *graphql.Field {

// 	return  &graphql.Field{
// 		Type:        graphql.NewList(recipeType),
// 		Description: "List of ingredients.",
// 		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

// 			rows, err := db.Query("SELECT * FROM ingredients")
// 			utils.CheckErr(err)
			
// 			var ingredients []*models.Ingredient

// 			for rows.Next() {
// 				Ingredient := &models.Ingredient{}
// 				err = rows.Scan(&Ingredient.Ingredient_id, &Ingredient.Ingredient_name,  &Ingredient.Created_at, &Ingredient.Updated_at)
// 				utils.CheckErr(err)
// 				ingredients = append(ingredients, Ingredient)
// 			}

// 			return ingredients, nil
// 		},
// 	}
// }

