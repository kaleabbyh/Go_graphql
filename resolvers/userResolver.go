package resolvers

import (
	"database/sql"
	"fmt"
	"time"

	models "github.com/kaleabbyh/Food_Recipie/model"
	"github.com/kaleabbyh/Food_Recipie/utils"

	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
)




var userType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "user",
	Description: "An user",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The identifier of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.ID, nil
				}

				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The name of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.Name, nil
				}

				return nil, nil
			},
		},
		"email": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The email address of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.Email, nil
				}

				return nil, nil
			},
		},
		"password": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The password of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.Password, nil
				}

				return nil, nil
			},
		},

		"created_at": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The created_at date of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.Created_at, nil
				}

				return nil, nil
			},
		},

		"updated_at": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The updated_at date of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*models.User); ok {
					return user.Updated_at, nil
				}

				return nil, nil
			},
		},
	},
})


//create user
func CreateUser(db *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        userType,
		Description: "Create new user",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			name, _ := params.Args["name"].(string)
			email, _ := params.Args["email"].(string)
			password, _ := utils.HashPassword(params.Args["password"].(string))
			createdAt := time.Now()
			updatedAt := time.Now()

			var lastInsertId int
			err := db.QueryRow("INSERT INTO users(name, email,password, created_at,updated_at) VALUES($1, $2, $3, $4, $5) returning id;", name, email,password, createdAt,updatedAt).Scan(&lastInsertId)
			utils.CheckErr(err)
			
			newUser := &models.User{
				ID:        lastInsertId,
				Name:      name,
				Email:     email,
				Password:  password,
				Created_at: createdAt,
				Updated_at: updatedAt,
			}

			return newUser, nil
		},
	}
}


//update user
func UpdateUser(db *sql.DB) *graphql.Field {
	return  &graphql.Field{
		Type:        userType,
		Description: "Update an user",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(int)
			name, _ := params.Args["name"].(string)
			email, _ := params.Args["email"].(string)
			updatedAt := time.Now()

			stmt, err := db.Prepare("UPDATE users SET name = $1, email = $2, updated_at = $3 WHERE id = $4")
			utils.CheckErr(err)
			_, err2 := stmt.Exec(name, email,updatedAt, id)
			utils.CheckErr(err2)

			newUser := &models.User{
				ID:    id,
				Name:  name,
				Email: email,
				Updated_at: updatedAt,
			}

			return newUser, nil
		},
	}
}

//delete user
func DeleteUser(db *sql.DB) *graphql.Field {
return &graphql.Field{
	Type:        userType,
	Description: "Delete an user",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, _ := params.Args["id"].(int)

		stmt, err := db.Prepare("DELETE FROM users WHERE id = $1")
		utils.CheckErr(err)

		_, err2 := stmt.Exec(id)
		utils.CheckErr(err2)

		return nil, nil
	},
}
}


//get users
func GetUsers(db *sql.DB) *graphql.Field {

	
	return  &graphql.Field{
		Type:        graphql.NewList(userType),
		Description: "List of users.",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			email, ok := p.Context.Value("email").(string)
			if ok {
				fmt.Println(email)
			}

			rows, err := db.Query("SELECT * FROM users")
			utils.CheckErr(err)
			
			var users []*models.User

			for rows.Next() {
				user := &models.User{}
				err = rows.Scan(&user.ID, &user.Name, &user.Email,&user.Password, &user.Created_at, &user.Updated_at)
				utils.CheckErr(err)
				users = append(users, user)
			}

			return users, nil
		},
	}
}


//get user
func GetUser(db *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        userType,
		Description: "Get an user.",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(int)

			user := &models.User{}
			err := db.QueryRow("select id, name, email from users where id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
			utils.CheckErr(err)

			return user, nil
		},
	}
}