package customer_model

import (
	"time"

	"mock-golang/protobuf"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Customer struct {
	Id             uuid.UUID `gorm:"type:uuid;primaryKey"`
	Role           int32     `gorm:"column:role"`
	Name           string    `gorm:"column:customer_name"`
	Email          string    `gorm:"column:email"`
	PhoneNumber    string    `gorm:"column:phone_number"`
	DateOfBith     string    `gorm:"column:date_of_bith"`
	IdentityCard   string    `gorm:"column:identity_card"`
	Address        string    `gorm:"column:address"`
	MembershipCard string    `gorm:"column:membership_card"`
	Password       string    `gorm:"column:password"`
	Status         int32     `gorm:"column:status"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

func (in *Customer) ToResponse() *protobuf.Customer {
	customerRes := &protobuf.Customer{
		Id:             in.Id.String(),
		Role:           in.Role,
		Name:           in.Name,
		Email:          in.Email,
		PhoneNumber:    in.PhoneNumber,
		DateOfBith:     in.DateOfBith,
		IdentityCard:   in.IdentityCard,
		Address:        in.Address,
		MembershipCard: in.MembershipCard,
		Password:       in.Password,
		Status:         in.Status,
		CreatedAt:      timestamppb.New(in.CreatedAt),
		UpdatedAt:      timestamppb.New(in.UpdatedAt),
	}

	return customerRes
}
