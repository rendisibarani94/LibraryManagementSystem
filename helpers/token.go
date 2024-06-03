package helpers

import (
	"first-jwt/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
// this mechanism is a lil fckd up :3

var signatureKey = []byte("ytta") // this is the key of every interaction with token

type CustomClaims struct { // this is the struct in collecting the token (based on docs :3)
	ID   int    `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(user *models.User) (string, error){ // function create token
	claims := CustomClaims { // mapping the user data into CustomClaims struct
		user.ID,
		user.Name,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)), // expired at 60 minutes
			IssuedAt: jwt.NewNumericDate(time.Now()), // based on docs
			NotBefore: jwt.NewNumericDate(time.Now()), // based on docs
		},
	}
	// NewWithClaims creates a new [Token] with the specified signing method and claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) 
	// SignedString creates and returns a complete, signed JWT. The token is signed using the Signing Method specified in the token.
	ss, err := token.SignedString(signatureKey) 
	// and addition also returned err
	return ss, err
}

// function that everything depends on docs
func ValidateToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error){
		return []byte(signatureKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("Unauthorized")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Unauthorized")
	}
	
	return claims, nil
}