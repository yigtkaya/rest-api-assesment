package controller

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yigtkaya/rest-api-assesment/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	usercollection *mongo.Collection
	usercontroller UserController
	userservice    services.UserService
	mongoclientu   *mongo.Client
	ctxx           context.Context
	erro           error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclientu, erro = mongo.Connect(ctx, mongoconn)

	if erro != nil {
		log.Fatal(erro)
	}
	erro = mongoclient.Ping(ctx, readpref.Primary())
	if erro != nil {
		log.Fatal(erro)
	}

	usercollection = mongoclient.Database("denemedb").Collection("users")
	userservice = services.NewUserService(usercollection, ctx)
	usercontroller = NewUser(userservice)

}
func TestCreateUser(t *testing.T) {

	gin.SetMode(gin.TestMode)
	var json = []byte(`
	{"id" : "0",
    "email":"hasan@gmail.com",
    "password": "123456",
    "name" : "hasan",
    "membership":{
        "id": "0",
        "group_name": "Fire"
    }
    }`)

	req, _ := http.NewRequest("POST", "/v1/user/createUser", bytes.NewBuffer(json))

	req.Header.Set("Content-Type", "application/json")
	router := gin.Default()
	router.POST("/v1/user/createUser", usercontroller.CreateUser)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	fmt.Println(w.Body)
}

func TestGetuser(t *testing.T) {

	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest("GET", "/v1/user/getUser/0", nil)
	req.Header.Set("Content-Type", "application/json")
	router := gin.Default()
	router.GET("/v1/user/getUser/:id", usercontroller.GetUser)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	fmt.Println(w.Body)
}

func TestGetAll(t *testing.T) {

	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest("GET", "/v1/user/getAllUsers", nil)
	req.Header.Set("Content-Type", "application/json")
	router := gin.Default()
	router.GET("/v1/user/getAllUsers", usercontroller.GetAll)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	fmt.Println(w.Body)

}

func TestUpdateuser(t *testing.T) {

	gin.SetMode(gin.TestMode)
	var json = []byte(`{
	"id" : "0",
    "email":"hasankaya@gmail.com",
    "password": "123456789",
    "name" : "hasan yiğit",
    "membership":{
        "id": "1",
        "group_name": "Water"
    }
    }`)

	req, _ := http.NewRequest("PATCH", "/v1/user/updateUser", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	router := gin.Default()
	router.PATCH("/v1/user/updateUser", usercontroller.UpdateUser)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	fmt.Println(w.Body)

}

func TestDeleteuser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest("DELETE", "/v1/user/deleteUser/0", nil)
	req.Header.Set("Content-Type", "application/json")
	router := gin.Default()
	router.DELETE("/v1/user/deleteUser/:id", usercontroller.DeleteUser)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	fmt.Println(w.Body)

	t.Cleanup(func() {

	})
}
