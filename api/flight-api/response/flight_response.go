package flight_response

type CreateFlightResponse struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	From          string `json:"from"`
	To            string `json:"to"`
	DepartDate    string `json:"departDate"`
	DepartTime    string `json:"departTime"`
	Status        string `json:"status"`
	AvailableSlot int32  `json:"slot"`
}

type FlightResponse struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	From          string `json:"from"`
	To            string `json:"to"`
	Status        string `json:"status"`
	AvailableSlot int32  `json:"slot"`
	DepatureDate  string `json:"depature_date"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
