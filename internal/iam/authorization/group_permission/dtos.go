package group_permission

type ListRequest struct {
	Page int `form:"page" binding:"required,min=1"`
}

type DetailsRequest struct {
	ID uint `uri:"id" binding:"required,min=1"`
}

type CreateRequest struct {
	GroupPermission GroupPermissionEntity `json:"group_permission" binding:"required"`
}

type UpdateRequest struct {
	GroupPermission GroupPermissionEntity `json:"group_permission" binding:"required"`
}

type DeleteRequest struct {
	ID uint `uri:"id" binding:"required,min=1"`
}
