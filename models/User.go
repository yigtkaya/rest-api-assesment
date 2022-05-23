package models

// swagger:model
type User struct {
	//the id for this User
	//required : true
	//in: string
	ID string `json:"id" bson:"id"`
	//the email for this User
	//required : true
	//in: string
	Email string `json:"email" bson:"email"`
	//the password for this User
	//required: true
	//in: string
	Password string `json:"password" bson:"password"`
	//the name for this User
	//required: true
	//in: string
	Name string `json:"name" bson:"name"`
	//the Group for this User
	//required: true
	//max item: 1
	//in: object
	Group Group `json:"membership" bson:"membership"`
}
