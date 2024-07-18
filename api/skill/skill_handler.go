package skill

import (
	"net/http"
	"skill-api-kafka/response"

	"github.com/gin-gonic/gin"
)

type SkillRepo interface {
	GetSkills() ([]Skill, error)
	GetSkillByKey(string) (Skill, error)
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

func (h *SkillHandler) GetSkillByKey(ctx *gin.Context) {
	key := ctx.Param("key")

	skill, err := h.skillrepo.GetSkillByKey(key)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, http.StatusOK, skill)
}
