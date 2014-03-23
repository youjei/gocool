package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int64
	Birthday  time.Time
	Age       int64
	Name      string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	Emails            []Email       // Embedded structs
	BillingAddress    Address       // Embedded struct
	BillingAddressId  sql.NullInt64 // BillingAddress's foreign key
	ShippingAddress   Address       // Another Embedded struct with same type
	ShippingAddressId int64         // ShippingAddress's foreign key
	IgnoreMe          int64         `sql:"-"` // Ignore this field
}

type Email struct {
	Id         int64
	UserId     int64  // Foreign key for User
	Email      string `sql:"type:varchar(100);"` // Set this field's type
	Subscribed bool
}

type Address struct {
	Id       int64
	Address1 string         `sql:"not null;unique"` // Set this field as not nullable and unique in database
	Address2 string         `sql:"type:varchar(100);unique"`
	Post     sql.NullString `sql:not null`
	// FYI, "NOT NULL" will only work well with NullXXX Scanner, because golang will initalize a default value for most type...
}
