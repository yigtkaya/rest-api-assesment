package main

import (
	"context"
	"fmt"
	"log"

	"github.com/yigtkaya/rest-api-assesment/configs"
	"github.com/yigtkaya/rest-api-assesment/controller"
	"github.com/yigtkaya/rest-api-assesment/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server          *gin.Engine
	userservice     services.UserService
	groupservice    services.GroupService
	usercontroller  controller.UserController
	groupcontroller controller.GroupController
	ctx             context.Context
	usercollection  *mongo.Collection
	groupcollection *mongo.Collection
	mongoclient     *mongo.Client
	err             error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI(configs.EnvMongoURI())
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connection established")

	usercollection = mongoclient.Database("Iaestedb").Collection("Users")
	groupcollection = mongoclient.Database("Iaestedb").Collection("Groups")

	userservice = services.NewUserService(usercollection, ctx)
	groupservice = services.NewGroupService(groupcollection, ctx)

	usercontroller = controller.NewUser(userservice)
	groupcontroller = controller.NewGroup(groupservice)

	server = gin.Default()

}

func main() {

	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("/v1")
	server.Static("/swaggerui", "slovenia/swaggerui")
	usercontroller.RegisterUserRoutes(basepath)
	groupcontroller.RegisterGroupRoutes(basepath)
	log.Fatal(server.Run(":9090"))

}
