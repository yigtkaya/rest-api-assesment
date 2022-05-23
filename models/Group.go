package models

// swagger:model
type Group struct {
	//the id for this group
	//required: true
	//max: 1
	//in: string
	ID string `json:"id" bson:"id"`
	//the name for this group
	//required: true
	//max: 1
	//in: string
	Name string `json:"group_name" bson:"group_name"`
}
