package services

import (
	"context"
	"errors"

	"github.com/yigtkaya/rest-api-assesment/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GroupServiceImpl struct {
	groupcollection *mongo.Collection
	ctx             context.Context
}

func NewGroupService(groupcollection *mongo.Collection, ctx context.Context) GroupService {
	return &GroupServiceImpl{
		groupcollection: groupcollection,
		ctx:             ctx,
	}
}

func (g *GroupServiceImpl) CreateGroup(group *models.Group) error {

	_, err := g.groupcollection.InsertOne(g.ctx, group)
	return err
}

func (g *GroupServiceImpl) GetGroup(id *string) (*models.Group, error) {
	var group *models.Group
	query := bson.D{bson.E{Key: "id", Value: id}}
	err := g.groupcollection.FindOne(g.ctx, query).Decode(&group)
	return group, err
}

func (g *GroupServiceImpl) GetAllG() ([]*models.Group, error) {
	var groups []*models.Group
	cursor, err := g.groupcollection.Find(g.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(g.ctx) {
		var group models.Group
		err := cursor.Decode(&group)
		if err != nil {
			return nil, err
		}
		groups = append(groups, &group)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(g.ctx)

	if len(groups) == 0 {
		return nil, errors.New("Not found any User")
	}
	return groups, nil
}

func (g *GroupServiceImpl) UpdateGroup(group *models.Group) error {
	filter := bson.D{bson.E{Key: "id", Value: group.ID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "id", Value: group.ID}, bson.E{Key: "group_name", Value: group.Name}}}}
	result, _ := g.groupcollection.UpdateOne(g.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("No matched User found for update")
	}
	return nil
}

func (g *GroupServiceImpl) DeleteGroup(id *string) error {
	filter := bson.D{bson.E{Key: "id", Value: id}}
	result, _ := g.groupcollection.DeleteOne(g.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("No matched User found for delete")
	}
	return nil
}
