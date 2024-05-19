package models

type Subscriber struct {
	Email string `json:"email" form:"email" gorm:"primaryKey;email"`
}
