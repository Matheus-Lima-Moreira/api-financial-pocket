package group_permission

import "time"

type GroupPermissionType string

const (
	GroupPermissionSystem GroupPermissionType = "SYSTEM"
	GroupPermissionCustom GroupPermissionType = "CUSTOM"
)

type GroupPermissionEntity struct {
	ID        uint                `json:"id"`
	Name      string              `json:"name"`
	Type      GroupPermissionType `json:"type"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}
