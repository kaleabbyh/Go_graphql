package middleware

import (
	"context"
	"net/http"
	"strings"

	"fmt"

	"github.com/gin-gonic/gin"
	conn "github.com/kaleabbyh/Food_Recipie/config"
	models "github.com/kaleabbyh/Food_Recipie/model"
	"github.com/kaleabbyh/Food_Recipie/utils"

	"github.com/dgrijalva/jwt-go"
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
// Middleware function to handle JWT authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			secretKey := []byte("Kaleabbyh")
			return secretKey, nil
		})
		
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		// Extract the email value from the token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		email, ok := claims["email"].(string)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		
		// Set the email value to a request context variable
		ctx := context.WithValue(r.Context(), "email", email)
		// fmt.Println("my email is: ",ctx)
		// Create a new request with the updated context
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}