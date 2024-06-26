package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/go-playground/validator/v10"
)

// PaymentMethod represents the method of payment for an order
type PaymentMethod string

const (
	// Digital payment method
	Digital PaymentMethod = "Digital"

	// Cash on Delivery payment method
	COD     PaymentMethod = "COD"
)

// OrderStatus represents the current status of an order
type OrderStatus string

const (
	// Order is pending
	Pending   OrderStatus = "Pending"

	// Order has been shipped
	Shipped   OrderStatus = "Shipped"

	// Order has been delivered
	Delivered OrderStatus = "Delivered"

	// Order has been canceled
	Canceled  OrderStatus = "Canceled"
)

// User represents a user in the system
type User struct {
	// Unique identifier for the user
	ID        primitive.ObjectID `bson:"_id,omitempty"`

	// First name of the user
	FirstName string             `bson:"first_name" validate:"required"`

	// Last name of the user
	LastName  string             `bson:"last_name" validate:"required"`

	// Username of the user
	Username  string             `bson:"username" validate:"required,min=3,max=20"`

	// Password of the user
	Password  string             `bson:"password" validate:"required,min=8"`

	// Email address of the user
	Email     string             `bson:"email" validate:"required,email"`

	// Phone number of the user
	Phone     string             `bson:"phone" validate:"required"`

	// Address of the user
	Address   Address            `bson:"address" validate:"required"`

	// Authentication token for the user
	Token     string             `bson:"token"`

	// Refresh token for the user
	RefreshToken string          `bson:"refresh_token"`

	// Timestamp when the user was created
	CreatedAt time.Time          `bson:"created_at"`

	// Timestamp when the user was last updated
	UpdatedAt time.Time          `bson:"updated_at"`
}

// Product represents a product in the system
type Product struct {
	// Unique identifier for the product
	ID          primitive.ObjectID `bson:"_id,omitempty"`

	// Name of the product
	Name        string             `bson:"name" validate:"required"`

	// Description of the product
	Description string             `bson:"description" validate:"required"`

	// Price of the product
	Price       float64            `bson:"price" validate:"required,gt=0"`

	// Rating of the product
	Rating      float64            `bson:"rating" validate:"min=0,max=5"`

	// Image URL of the product
	Image       string             `bson:"image" validate:"required,url"`

	// Timestamp when the product was created
	CreatedAt   time.Time          `bson:"created_at"`

	// Timestamp when the product was last updated
	UpdatedAt   time.Time          `bson:"updated_at"`
}

// Order represents an order in the system
type Order struct {
	// Unique identifier for the order
	ID            primitive.ObjectID   `bson:"_id,omitempty"`

	// Unique identifier of the user who placed the order
	UserID        primitive.ObjectID   `bson:"user_id" validate:"required"`

	// List of unique identifiers of the products in the order
	ProductIDs    []primitive.ObjectID `bson:"product_ids" validate:"required,min=1"`

	// Total price of the order
	TotalPrice    float64              `bson:"total_price" validate:"required,gt=0"`

	// Method of payment for the order (Digital or COD)
	PaymentMethod PaymentMethod        `bson:"payment_method" validate:"required,oneof=Digital COD"`

	// Current status of the order (Pending, Shipped, Delivered, Canceled)
	OrderStatus   OrderStatus          `bson:"order_status" validate:"required,oneof=Pending Shipped Delivered Canceled"`

	// Timestamp when the order was created
	CreatedAt     time.Time            `bson:"created_at"`

	// Timestamp when the order was last updated
	UpdatedAt     time.Time            `bson:"updated_at"`
}

// Cart represents a shopping cart in the system
type Cart struct {
	// Unique identifier for the cart
	ID          primitive.ObjectID   `bson:"_id,omitempty"`

	// Unique identifier of the user who owns the cart
	UserID      primitive.ObjectID   `bson:"user_id" validate:"required"`

	// List of unique identifiers of the products in the cart
	ProductIDs  []primitive.ObjectID `bson:"product_ids"`

	// Total price of the items in the cart
	TotalPrice  float64              `bson:"total_price" validate:"min=0"`

	// Timestamp when the cart was created
	CreatedAt   time.Time            `bson:"created_at"`

	// Timestamp when the cart was last updated
	UpdatedAt   time.Time            `bson:"updated_at"`
}

// Address represents an address in the system
type Address struct {
	// Unique identifier for the address
	ID        primitive.ObjectID `bson:"_id,omitempty"`

	// Unique identifier of the user who owns the address
	UserID    primitive.ObjectID `bson:"user_id" validate:"required"`

	// Street address
	Street    string             `bson:"street" validate:"required"`

	// City of the address
	City      string             `bson:"city" validate:"required"`

	// State of the address
	State     string             `bson:"state" validate:"required"`

	// Zip code of the address
	ZipCode   string             `bson:"zip_code" validate:"required"`

	// Country of the address
	Country   string             `bson:"country" validate:"required"`

	// Timestamp when the address was created
	CreatedAt time.Time          `bson:"created_at"`

	// Timestamp when the address was last updated
	UpdatedAt time.Time          `bson:"updated_at"`
}

// Validate checks if the struct is valid
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// Validate checks if the struct is valid
func (p *Product) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

// Validate checks if the struct is valid
func (o *Order) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}

// Validate checks if the struct is valid
func (c *Cart) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

// Validate checks if the struct is valid
func (a *Address) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}