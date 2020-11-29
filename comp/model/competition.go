package model

type Competition struct {
	ID            int64          `json:"id"`
	Status        status         `json:"status"`
	ClimbingTypes []ClimbingType `json:"climbingTypes"`
	Year          int64          `json:"year"`
	Name          string         `json:"name"`
	Mnemonic      string         `json:"mnemonic"`
	Event         int64          `json:"event"`
	Notes         string         `json:"notes"`
}

type status string

const (
	Active status = "ACTIVE"
	Closed status = "CLOSED"
)
