package jwt

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/prometheus/common/log"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"strings"
	"github.com/gorilla/context"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct{
	Message string `json:"message"`
}

func CreateTokenEndPoint(w http.ResponseWriter,req *http.Request){
	fmt.Println("in create token end point")
	var user User
	_=json.NewDecoder(req.Body).Decode(&user)
	token:= jwt.NewWithClaims(jwt.SigningMethodES256,jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})
	tokenString,errors := token.SignedString([]byte("secret"))
	if errors != nil{
		fmt.Println(errors)
	}
	json.NewEncoder(w).Encode(JwtToken{Token:tokenString})
}

func ProtectedEndpoint(w http.ResponseWriter, req *http.Request){
	params:=req.URL.Query()
	token,_:=jwt.Parse(params["token"][0], func(token *jwt.Token) (interface{}, error) {
		if _,ok:=token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil,fmt.Errorf("There was an error")
		}
		return []byte("secrey"),nil
	})
	if claims,ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		var user User
		mapstructure.Decode(claims,&user)
		json.NewEncoder(w).Encode(user)
	}else {
		json.NewEncoder(w).Encode(Exception{Message:"Invalid authorization token"})
	}
}

func TestEndpoint(w http.ResponseWriter, req *http.Request) {
	decoded := context.Get(req, "decoded")
	var user User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	json.NewEncoder(w).Encode(user)
}

func ValidateMiddleware(next http.HandlerFunc)http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter,req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != ""{
			bearerToken := strings.Split(authorizationHeader," ")
			if len(bearerToken)==2{
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(Exception{Message: error.Error()})
					return
				}
				if token.Valid{
					context.Set(req,"decoded",token.Claims)
					next(w,req)
				}else {
					json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				}
			}
		}else{
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
		}
	})
}

func JwtMain(){
	router := mux.NewRouter()
	fmt.Println("Starting the application")
	router.HandleFunc("/authenticate",CreateTokenEndPoint).Methods("POST")
	router.HandleFunc("/protected",ProtectedEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":1234",router))
}