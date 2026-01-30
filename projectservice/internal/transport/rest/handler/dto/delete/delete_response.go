package deletedto

type DeleteResponse struct {
	IsDeleted bool `json:"is_deleted" binding:"required"`
}
