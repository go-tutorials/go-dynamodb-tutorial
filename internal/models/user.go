package models

import "time"

type UserKey struct {
	Id  string    `json:"id"`
}

type UserUpdate struct {
	Username    string     		`json:":u"`
	Email       string     		`json:":e"`
	Phone       string     		`json:":p"`
	DateOfBirth *time.Time 		`json:":d"`
}

type User struct {
	Id          string     `json:"id" gorm:"column:id;primary_key" bson:"_id" dynamodbav:"id" firestore:"-" validate:"required,max=40"`
	Username    string     `json:"username" gorm:"column:username" bson:"username" dynamodbav:"username" firestore:"username" validate:"required,username,max=100"`
	Email       string     `json:"email" gorm:"column:email" bson:"email" dynamodbav:"email" firestore:"email" validate:"email,max=100"`
	Phone       string     `json:"phone" gorm:"column:phone" bson:"phone" dynamodbav:"phone" firestore:"phone" validate:"required,phone,max=18"`
	DateOfBirth *time.Time `json:"dateOfBirth" gorm:"column:date_of_birth" bson:"dateOfBirth" dynamodbav:"dateOfBirth" firestore:"dateOfBirth2"`
}

func (u User)GetTableName() string {
	return "users"
}