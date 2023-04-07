package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                    primitive.ObjectID `bson:"_id"`
	First_name            *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name             *string            `json:"last_name" validate:"required,min=2,max=100"`
	Email                 *string            `json:"email" validate:"email,required"`
	Password              *string            `json:"Password" validate:"required,min=8"`
	Phone                 *string            `json:"phone" validate:"required"`
	Token                 *string            `json:"token"`
	User_type             *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token         *string            `json:"refresh_token"`
	Created_at            time.Time          `json:"created_at"`
	Updated_at            time.Time          `json:"updated_at"`
	User_id               string             `json:"user_id"`
	Property_ID           *int               `json:"property_id"`
	Property_Name         *string            `json:"property_name"`
	Contact_No            *string            `json:"contact_no"`
	Property_Type         *string            `json:"property_type"`
	Property_Address      *string            `json:"property_address"`
	City                  *string            `json:"city"`
	Pin_Code              *string            `json:"pin_code"`
	Landmark              *string            `json:"landmark"`
	Property_Image        *string            `json:"property_image"`
	Ammeneties            *string            `json:"ammeneties"`
	Price_Month           *int               `json:"price_month"`
	Price_Day             *int               `json:"price_day"`
	Advance_Deposit_Month *int               `json:"advance_deposit_month"`
	Advance_Deposit_Day   *int               `json:"advance_deposit_day"`
}
