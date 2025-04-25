package models

// HTTPError describes an error response.
// swagger:model HTTPError
type HTTPError struct {
	// the HTTP status code
	// example: 400
	Code int `json:"code" example:"400"`

	// a descriptive error message
	// example: bad request
	Message string `json:"message" example:"bad request"`
}
