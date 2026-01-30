package deletedto

type DeleteRequest struct {
	Name string `json:"name" binding:"required"`
}
