package resolvers

import (
	"database/sql"
	"time"

	"github.com/graphql-go/graphql"
	models "github.com/kaleabbyh/Food_Recipie/model"
	"github.com/kaleabbyh/Food_Recipie/utils"
	_ "github.com/lib/pq"
)



var ingredientType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ingredient",
	Description: "a ingredient",
	Fields: graphql.Fields{
		"ingredient_id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The identifier of the ingredient.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if ingredient, ok := p.Source.(*models.Ingredient); ok {
					return ingredient.Ingredient_id, nil
				}

				return nil, nil
			},
		},
		"ingredient_name": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The name of the ingredient.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if ingredient, ok := p.Source.(*models.Ingredient); ok {
					return ingredient.Ingredient_name, nil
				}

				return nil, nil
			},
		},
		

		"created_at": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The created_at date of the ingredient.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if ingredient, ok := p.Source.(*models.Ingredient); ok {
					return ingredient.Created_at, nil
				}

				return nil, nil
			},
		},

		"updated_at": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The updated_at date of the ingredient.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if ingredient, ok := p.Source.(*models.Ingredient); ok {
					return ingredient.Updated_at, nil
				}

				return nil, nil
			},
		},
		

		
	},
})


//create Ingredient
func CreateIngredient(db *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        ingredientType,
		Description: "Create new ingredient",
		Args: graphql.FieldConfigArgument{
			"ingredient_name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			ingredient_name, _ := params.Args["ingredient_name"].(string)
			createdAt := time.Now()
			updatedAt := time.Now()
			
			
			var lastInsertId int
			err := db.QueryRow("INSERT INTO ingredients(ingredient_name, created_at,updated_at) VALUES($1, $2, $3) returning ingredient_id;", ingredient_name,  createdAt,updatedAt).Scan(&lastInsertId)
			utils.CheckErr(err)
			


			newIngredient := &models.Ingredient{
				Ingredient_id:    lastInsertId,
				Ingredient_name:  ingredient_name,
				Created_at:     createdAt,
				Updated_at:     updatedAt,
			}

			return newIngredient, nil
		},
	}
}




//get users
func GetIngredients(db *sql.DB) *graphql.Field {

	return  &graphql.Field{
		Type:        graphql.NewList(ingredientType),
		Description: "List of ingredients.",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			rows, err := db.Query("SELECT * FROM ingredients")
			utils.CheckErr(err)
			
			var ingredients []*models.Ingredient

			for rows.Next() {
				Ingredient := &models.Ingredient{}
				err = rows.Scan(&Ingredient.Ingredient_id, &Ingredient.Ingredient_name,  &Ingredient.Created_at, &Ingredient.Updated_at)
				utils.CheckErr(err)
				ingredients = append(ingredients, Ingredient)
			}

			return ingredients, nil
		},
	}
}

