package entity

// Error is struct of error object
type Error struct {
	Code    int   `json:"code"`
	Message any   `json:"message"`
	Error   error `json:"-"`
}
