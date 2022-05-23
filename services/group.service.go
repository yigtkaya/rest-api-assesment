package services

import (
	"github.com/yigtkaya/rest-api-assesment/models"
)

type GroupService interface {
	CreateGroup(*models.Group) error
	GetGroup(*string) (*models.Group, error)
	GetAllG() ([]*models.Group, error)
	UpdateGroup(*models.Group) error
	DeleteGroup(*string) error
}
