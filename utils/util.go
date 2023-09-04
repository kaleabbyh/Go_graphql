package utils

import (
	"fmt"
	"os"

	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

func GenerateJWT() (string, error) {
    // Create a new token object
    token := jwt.New(jwt.SigningMethodHS256)

    // Set the claims (payload) of the token
    claims := token.Claims.(jwt.MapClaims)
    claims["user_id"] = 123
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

    // Generate the JWT token string
    tokenString, err := token.SignedString([]byte("your-secret-key"))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}


func GenerateToken(email string) (string, error) {
	err := godotenv.Load()
	CheckErr(err)
	secret := os.Getenv("secret")

	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("Failed to sign the token:", err)
		return "", err
	}

	fmt.Println("Generated token:", tokenString)
	return tokenString,nil
}



func ValidateToken(tokenString string) (string, error)  {
	
	err := godotenv.Load()
	CheckErr(err)
	secret := os.Getenv("secret")

	// Parse and verify the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	// Check for errors during token parsing or verification
	if err != nil {
		return "", fmt.Errorf("token is invalid")
	}

	// Check if the token is valid
	if token.Valid {
		// Access the claims
		claims := token.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		//expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)

		fmt.Println("email:", email)
		return email, nil
	} else {
		fmt.Println("Token is invalid")
		return "", fmt.Errorf("token is invalid")
	}
}