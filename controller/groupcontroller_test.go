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
	groupcollection *mongo.Collection
	groupcontroller GroupController
	groupservice    services.GroupService
	mongoclient     *mongo.Client
	ctx             context.Context
	err             error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)

	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	groupcollection = mongoclient.Database("denemedb").Collection("groups")
	groupservice = services.NewGroupService(groupcollection, ctx)
	groupcontroller = NewGroup(groupservice)

}
func TestCreateGroup(t *testing.T) {

	gin.SetMode(gin.TestMode)
	var json = []byte(`{"id": "0","group_name" : "TEST NAME"}`)

	req, _ := http.NewRequest("POST", "/v1/group/createGroup", bytes.NewBuffer(json))

	req.Header.Set("Content-Type", "application/json")
	router := gin.Default()
	router.POST("/v1/group/createGroup", groupcontroller.CreateGroup)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	fmt.Println(w.Body)
}

func TestGetGroup(t *testing.T) {

	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest("GET", "/v1/group/getGroup/0", nil)
	req.Header.Set("Content-Type", "application/json")
	router := gin.Default()
	router.GET("/v1/group/getGroup/:id", groupcontroller.GetGroup)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	fmt.Println(w.Body)
}

func TestGetAllG(t *testing.T) {

	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest("GET", "/v1/group/getAllGroup", nil)
	req.Header.Set("Content-Type", "application/json")
	router := gin.Default()
	router.GET("/v1/group/getAllGroup", groupcontroller.GetAllG)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	fmt.Println(w.Body)

}

func TestUpdateGroup(t *testing.T) {

	gin.SetMode(gin.TestMode)
	var json = []byte(`{"id": "0","group_name" : "AAA BBB"}`)

	req, _ := http.NewRequest("PATCH", "/v1/group/updateGroup", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	router := gin.Default()
	router.PATCH("/v1/group/updateGroup", groupcontroller.UpdateGroup)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	fmt.Println(w.Body)

}

func TestDeleteGroup(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest("DELETE", "/v1/group/deleteGroup/0", nil)
	req.Header.Set("Content-Type", "application/json")
	router := gin.Default()
	router.DELETE("/v1/group/deleteGroup/:id", groupcontroller.DeleteGroup)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	fmt.Println(w.Body)

}
