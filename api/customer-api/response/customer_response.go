package customer_response

type CustomerResponse struct {
	Id             string `json:"id"`
	Role           int32  `json:"role"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phoneNumber"`
	DateOfBith     string `json:"dateOfBith"`
	IdentityCard   string `json:"identityCard"`
	Address        string `json:"address"`
	MembershipCard string `json:"membershipCard"`
	Status         int32  `json:"status"`
}

type ChangePasswordResponse struct {
	Id      string `json:"id" binding:"required"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
