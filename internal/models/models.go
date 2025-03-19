package models

type Reservation struct {
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	People      int    `json:"people"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	RoomChoice  string `json:"room_choice"`

}