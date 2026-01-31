package getalldto

import projectdomain "projectservice/internal/domain/project"

type GetAllResponse struct {
	Projects []*projectdomain.ProjectDomain `json:"projects" binding:"required"`
}
