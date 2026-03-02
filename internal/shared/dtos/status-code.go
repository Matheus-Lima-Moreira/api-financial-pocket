package dtos

type StatusCode int

const (
	SUCCESS               StatusCode = 1
	INVALID_FIELDS        StatusCode = 2
	NOT_FOUND             StatusCode = 3
	VALIDATION_ERROR      StatusCode = 4
	INTERNAL_SERVER_ERROR StatusCode = 5
	UNAUTHORIZED          StatusCode = 6
	SERVICE_UNAVAILABLE   StatusCode = 7
	FORBIDDEN             StatusCode = 8
)
