package createdto

type CreateResponse struct {
	ProjectId uint32 `json:"project_id" binding:"required"`
}
