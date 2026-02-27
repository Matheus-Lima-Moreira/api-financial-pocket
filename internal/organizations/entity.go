package organizations

import "time"

type OrganizationEntity struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Cellphone string    `json:"cellphone"`
	Logo      string    `json:"logo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
