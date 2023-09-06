package resolvers

import (
	"database/sql"
	"time"

	"github.com/graphql-go/graphql"
	models "github.com/kaleabbyh/Food_Recipie/model"
	"github.com/kaleabbyh/Food_Recipie/utils"
	_ "github.com/lib/pq"
)



var categoryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "category",
	Description: "a categorey",
	Fields: graphql.Fields{
		"category_id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The identifier of the category.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if category, ok := p.Source.(*models.Category); ok {
					return category.Category_id, nil
				}

				return nil, nil
			},
		},
		"category_name": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The name of the category.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if category, ok := p.Source.(*models.Category); ok {
					return category.Category_name, nil
				}

				return nil, nil
			},
		},
		

		"created_at": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The created_at date of the category.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if categorey, ok := p.Source.(*models.Category); ok {
					return categorey.Created_at, nil
				}

				return nil, nil
			},
		},

		"updated_at": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The updated_at date of the category.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if category, ok := p.Source.(*models.Category); ok {
					return category.Updated_at, nil
				}

				return nil, nil
			},
		},
		

		
	},
})


//create category
func CreateCategory(db *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        categoryType,
		Description: "Create new category",
		Args: graphql.FieldConfigArgument{
			"category_name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			category_name, _ := params.Args["category_name"].(string)
			createdAt := time.Now()
			updatedAt := time.Now()
			
			
			var lastInsertId int
			err := db.QueryRow("INSERT INTO categories(category_name, created_at,updated_at) VALUES($1, $2, $3) returning category_id;", category_name,  createdAt,updatedAt).Scan(&lastInsertId)
			utils.CheckErr(err)
			


			newCategory := &models.Category{
				Category_id:    lastInsertId,
				Category_name:  category_name,
				Created_at:     createdAt,
				Updated_at:     updatedAt,
			}

			return newCategory, nil
		},
	}
}




//get users
func GetCategories(db *sql.DB) *graphql.Field {

	return  &graphql.Field{
		Type:        graphql.NewList(categoryType),
		Description: "List of categories.",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			rows, err := db.Query("SELECT * FROM categories")
			utils.CheckErr(err)
			
			var categories []*models.Category

			for rows.Next() {
				category := &models.Category{}
				err = rows.Scan(&category.Category_id, &category.Category_name,  &category.Created_at, &category.Updated_at)
				utils.CheckErr(err)
				categories = append(categories, category)
			}

			return categories, nil
		},
	}
}

