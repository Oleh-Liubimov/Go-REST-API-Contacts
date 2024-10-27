package models

// import "go.mongodb.org/mongo-driver/bson/primitive"

type Contact struct {
	Name        string `bson:"name" json:"name" validate:"required"`
	PhoneNumber string `bson:"phoneNumber" json:"phoneNumber" validate:"required"`
	Email       string `bson:"email,omitempty" json:"email" validate:"email"`
	IsFavorite  bool   `bson:"isFavorite" json:"isFavorite,omitempty"`
	ContactType string `bson:"contactType" json:"contactType" binding:"required,oneof=work home personal"`
	// UserId      primitive.ObjectID `bson:"userId" json:"userId,omitempty"`
}

var ContactTypeEnum = []string{"work", "home", "personal"}
