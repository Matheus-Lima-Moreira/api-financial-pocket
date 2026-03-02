package dtos

import shared_errors "github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/errors"

type ReplyDTO struct {
	Code       StatusCode                  `json:"code,omitempty"`
	Message    string                      `json:"message,omitempty"`
	Data       interface{}                 `json:"data,omitempty"`
	Pagination *PaginationDTO              `json:"pagination,omitempty"`
	Errors     []shared_errors.ErrorDetail `json:"errors,omitempty"`
}
