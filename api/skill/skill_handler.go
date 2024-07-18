package skill

import (
	"net/http"
	"skill-api-kafka/response"

	"github.com/gin-gonic/gin"
)

type SkillRepo interface {
	GetSkills() ([]Skill, error)
}

type SkillHandler struct {
	skillrepo SkillRepo
}

func NewSkillHandler(skillrepo SkillRepo) *SkillHandler {
	return &SkillHandler{skillrepo: skillrepo}
}

func (h *SkillHandler) GetSkills(ctx *gin.Context) {
	skills, err := h.skillrepo.GetSkills()
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, http.StatusOK, skills)
}
