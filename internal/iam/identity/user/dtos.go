package user

type ListRequest struct {
	Page int `form:"page" binding:"required,min=1"`
}

type DetailsRequest struct {
	ID uint `uri:"id" binding:"required,min=1"`
}
