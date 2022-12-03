package customer_response

import "github.com/google/uuid"

type CustomerResponse struct {
	Id             uuid.UUID
	Role           int32
	Name           string
	Email          string
	PhoneNumber    string
	DateOfBith     string
	IdentityCard   string
	Address        string
	MembershipCard string
	Password       string
	Status         int32
}
