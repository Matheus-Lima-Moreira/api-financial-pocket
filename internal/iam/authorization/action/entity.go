package action

import "time"

type ActionEntity struct {
	ID        uint      `json:"id"`
	Resource  string    `json:"resource"`
	Action    string    `json:"action"`
	Label     string    `json:"label"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
