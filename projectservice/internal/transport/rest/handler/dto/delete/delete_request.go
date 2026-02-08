package deletedto

type DeleteRequest struct {
	ProjectId uint32 `json:"project_id" binding:"required"`
}
