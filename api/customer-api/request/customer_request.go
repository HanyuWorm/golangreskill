package customer_request

type CreateCustomerRequest struct {
	Role           int32  `json:"role" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required"`
	PhoneNumber    string `json:"phoneNumber" binding:"required"`
	DateOfBith     string `json:"dateOfBith" binding:"required"`
	IdentityCard   string `json:"identityCard" binding:"required"`
	Address        string `json:"address" binding:"max=256,min=6"`
	MembershipCard string `json:"membershipCard"`
	Password       string `json:"password"`
	Status         int32  `json:"status" binding:"required"`
}

type UpdateCustomerRequest struct {
	Id             string `json:"id" binding:"required"`
	Role           int32  `json:"role"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phoneNumber"`
	DateOfBith     string `json:"dateOfBith"`
	IdentityCard   string `json:"identityCard"`
	Address        string `json:"address"`
	MembershipCard string `json:"membershipCard"`
	Password       string `json:"password"`
	Status         int32  `json:"status"`
}

type ChangePasswordRequest struct {
	Id              string `json:"id" binding:"required"`
	OldPassword     string `json:"oldPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}
