package skill

import (
	"api/response"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SkillRepo interface {
	GetSkills() ([]Skill, error)
	GetSkillByKey(string) (Skill, error)
	PostSkillByKey(Skill) error
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

func (h *SkillHandler) PostSkill(ctx *gin.Context) {
	var skill Skill

	if err := ctx.BindJSON(&skill); err != nil {
		fmt.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Marshal skill struct to JSON
	skillJSON, err := json.Marshal(skill)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error marshaling skill data"})
		return
	}

	// Convert JSON byte array to string
	skillJSONString := string(skillJSON)

	fmt.Println("Skill -> ", skillJSON)
	fmt.Println("Skill JSON -> ", skillJSONString)

	produceMessage(skillJSONString, "post")
	ctx.JSON(http.StatusOK, gin.H{"message": "Skill creation request queued"})
}

func (h *SkillHandler) PutSkillByKey(ctx *gin.Context) {
	key := ctx.Param("key")

	var skill Skill
	skill.Key = key

	if err := ctx.BindJSON(&skill); err != nil {
		fmt.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Marshal skill struct to JSON
	skillJSON, err := json.Marshal(skill)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error marshaling skill data"})
		return
	}

	// Convert JSON byte array to string
	skillJSONString := string(skillJSON)

	fmt.Println("Skill -> ", skillJSON)
	fmt.Println("Skill JSON -> ", skillJSONString)

	fmt.Println("SKill Joker ----------------------------------------------------------------> ", skillJSONString)
	produceMessage(skillJSONString, "put")
	ctx.JSON(http.StatusOK, gin.H{"message": "Skill updation request queued"})
}
