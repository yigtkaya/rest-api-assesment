package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yigtkaya/rest-api-assesment/models"
	"github.com/yigtkaya/rest-api-assesment/services"
)

type UserController struct {
	UserService services.UserService
}

func NewUser(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	//swagger:route POST /createUser User CreateUser
	//Create a user.
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https, ws, wss
	//
	//     Deprecated: false
	//     Parameters:
	//       + name: User
	//         in: body
	//         required: true
	//		   schema:
	//		   	$ref: '#/definitions/User'
	//     Responses:
	//       200:
	//       	description: OK
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "succes"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	//swagger:route GET /getUser/{id} User GetUser
	//Get user by name.
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https, ws, wss
	//
	//     Deprecated: false
	//     Parameters:
	//       + name: name
	//         in: path
	//         required: true
	//         type: string
	//
	//     Responses:
	//       200:
	//       	description: OK
	username := ctx.Param("id")
	user, err := uc.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	//swagger:route GET /getAllUsers User GetAll
	// get list of users.
	//     Consumes:
	//     - application/json
	//     Produces:
	//     - application/json
	//     Deprecated: false
	// responses:
	//   200:
	//		description: OK
	users, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	//swagger:route PATCH /updateUser/ User UpdateUser
	// Update the user.
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https, ws, wss
	//
	//     Deprecated: false
	//     Parameters:
	//       + name: User
	//         in: body
	//         required: true
	//		   schema:
	//		   	$ref: '#/definitions/User'
	//     Responses:
	//       200:
	//       	description: OK
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "succes"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	//swagger:route DELETE /deleteUser/{id} User DeleteUser
	//Delete a user.
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https, ws, wss
	//
	//     Deprecated: false
	//     Parameters:
	//       + name: id
	//         in: path
	//         required: true
	//         type: string
	//
	//     Responses:
	//       200:
	//       	description: OK
	username := ctx.Param("id")
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "succes"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")

	userroute.POST("/createUser", uc.CreateUser)
	userroute.GET("/getUser/:id", uc.GetUser)
	userroute.GET("getAllUsers", uc.GetAll)
	userroute.PATCH("/updateUser", uc.UpdateUser)
	userroute.DELETE("/deleteUser/:id", uc.DeleteUser)

}
