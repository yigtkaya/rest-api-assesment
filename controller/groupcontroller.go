package controller

import (
	"net/http"

	"github.com/yigtkaya/rest-api-assesment/models"
	"github.com/yigtkaya/rest-api-assesment/services"

	"github.com/gin-gonic/gin"
)

type GroupController struct {
	GroupService services.GroupService
}

func NewGroup(groupservice services.GroupService) GroupController {
	return GroupController{
		GroupService: groupservice,
	}
}

func (gc *GroupController) CreateGroup(ctx *gin.Context) {
	//swagger:route POST /createGroup Group CreateGroup
	// Create a group.
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
	//       + name: Group
	//         in: body
	//         required: true
	//		   schema:
	//		   	$ref: '#/definitions/Group'
	//     Responses:
	//       200:
	//       	description: OK
	var group models.Group
	if err := ctx.ShouldBindJSON(&group); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return

	}
	err := gc.GroupService.CreateGroup(&group)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "succes"})
}

func (gc *GroupController) GetGroup(ctx *gin.Context) {
	//swagger:route GET /getGroup/{id} Group GetGroup
	// get a group.
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
	groupname := ctx.Param("id")
	group, err := gc.GroupService.GetGroup(&groupname)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, group)
}

func (gc *GroupController) GetAllG(ctx *gin.Context) {
	//swagger:route GET /getAllGroups Group getAllG
	// get list of groups.
	//     Consumes:
	//     - application/json
	//     Produces:
	//     - application/json
	//     Deprecated: true
	// responses:
	//   200:
	//		description: OK
	groups, err := gc.GroupService.GetAllG()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, groups)
}

func (gc *GroupController) UpdateGroup(ctx *gin.Context) {
	//swagger:route PATCH /updateGroup/ Group UpdateGroup
	// Update group with body.
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
	//       + name: Group
	//         in: body
	//         required: true
	//		   schema:
	//		   	$ref: '#/definitions/Group'
	//     Responses:
	//       200:
	//       	description: OK

	var group models.Group
	if err := ctx.ShouldBindJSON(&group); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := gc.GroupService.UpdateGroup(&group)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "succes"})
}

func (gc *GroupController) DeleteGroup(ctx *gin.Context) {
	//swagger:route DELETE /deleteGroup/{id} Group DeleteGroup
	// Delete a group by name.
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
	//       + name: group_name
	//         in: path
	//         required: true
	//         type: string
	//
	//     Responses:
	//       200:
	//        description: OK
	groupname := ctx.Param("id")
	err := gc.GroupService.DeleteGroup(&groupname)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "succes"})
}

func (gc *GroupController) RegisterGroupRoutes(rg *gin.RouterGroup) {
	grouproute := rg.Group("/group")

	grouproute.POST("/createGroup", gc.CreateGroup)
	grouproute.GET("/getGroup/:id", gc.GetGroup)
	grouproute.GET("getAllGroups", gc.GetAllG)
	grouproute.PATCH("/updateGroup/", gc.UpdateGroup)
	grouproute.DELETE("/deleteGroup/:id", gc.DeleteGroup)
}
