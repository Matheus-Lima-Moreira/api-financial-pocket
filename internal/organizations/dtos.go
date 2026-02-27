package organizations

type ListRequest struct {
	Page int `form:"page" binding:"required,min=1"`
}

type DetailsRequest struct {
	ID string `uri:"id" binding:"required,min=1"`
}

type CreateRequest struct {
	Organization OrganizationEntity `json:"organization" binding:"required"`
}

type UpdateRequest struct {
	Organization OrganizationEntity `json:"organization" binding:"required"`
}

type DeleteRequest struct {
	ID string `uri:"id" binding:"required,min=1"`
}
