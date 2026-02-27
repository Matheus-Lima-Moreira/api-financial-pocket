package action

type ListRequest struct {
	Page int `form:"page" binding:"required,min=1"`
}
