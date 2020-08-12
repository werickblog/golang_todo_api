package forms

// BookmarkPayload defines the payload sent by the user
type BookmarkPayload struct {
	Name string `json:"name" binding:"required"`
	Link string `json:"link" binding:"required"`
}
