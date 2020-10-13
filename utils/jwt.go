package utils

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/nurrizkyimani/golang-mux-backend/model"
)

func jwtgenerate(model *model.Article) (string, error) {

	secret := "setering"

	jwt := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			"email": "nurrizkyimani",
			"iss" : "course",
		})

	tokenString , err := jwt.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}
		
	spew.Dump(jwt)

	return tokenString, nil

	 

}

