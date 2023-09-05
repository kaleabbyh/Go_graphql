package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	conn "github.com/kaleabbyh/Food_Recipie/config"
	models "github.com/kaleabbyh/Food_Recipie/model"
	"github.com/kaleabbyh/Food_Recipie/utils"
)

func MiddlewareUser() gin.HandlerFunc {
	db,_:=conn.ConnectDB()
	defer db.Close()
	return func(ctx *gin.Context) {
		var token string
		cookie, err := ctx.Cookie("token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		} else if err == nil {
			token = cookie
		}

		if token != "" {
		email, err := utils.ValidateToken(token)
		utils.CheckErr(err)

		user := &models.User{}
		err = db.QueryRow("select id, name, email from users where email = $1", email).Scan(&user.ID, &user.Name, &user.Email)
		utils.CheckErr(err)
		// fmt.Println(user)
		ctx.Set("currentUser", user)
		ctx.Next()
		}

		
	}
}


// Middleware function to handle JWT authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		email, err :=utils.ValidateToken(tokenString)
		
		if err == nil {
			// Set the email value to a request context variable
		ctx := context.WithValue(r.Context(), "email", email)
		
		// Create a new request with the updated context
		r = r.WithContext(ctx)
		}
		

		next.ServeHTTP(w, r)
	})
}