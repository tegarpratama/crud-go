package dto

import "github.com/google/uuid"

type ProductRequest struct {
	Name        string 		`json:"name" binding:"required"`
	Description string 		`json:"description" binding:"required"`
	CategoryID  uuid.UUID 	`json:"category_id" binding:"required"`
}

type ProductResponse struct {
	ID          uuid.UUID   	`json:"id"`
	Name 		string			`json:"name"`
	Description string 			`json:"description"`	
	CategoryID	uuid.UUID		`json:"-"`
	Category	ProductCategory	`gorm:"foreignKey:CategoryID" json:"category"`
	CreatedAt  string 			`json:"created_at"`
}
