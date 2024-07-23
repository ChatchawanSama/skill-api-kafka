package skill

import (
	"api/response"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SkillRepo interface {
	GetSkillsRepo() ([]Skill, error)
	GetSkillByKeyRepo(string) (Skill, error)
}

type SkillHandler struct {
	skillrepo SkillRepo
}

func NewSkillHandler(skillrepo SkillRepo) *SkillHandler {
	return &SkillHandler{skillrepo: skillrepo}
}

func (h *SkillHandler) GetSkillsHandler(ctx *gin.Context) {
	skills, err := h.skillrepo.GetSkillsRepo()
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, http.StatusOK, skills)
}

func (h *SkillHandler) GetSkillByKeyHandler(ctx *gin.Context) {
	key := ctx.Param("key")

	skill, err := h.skillrepo.GetSkillByKeyRepo(key)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, http.StatusOK, skill)
}

func (h *SkillHandler) PostSkillHandler(ctx *gin.Context) {
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

func (h *SkillHandler) PutSkillByKeyHandler(ctx *gin.Context) {
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

func (h *SkillHandler) DeleteSkillByKeyHandler(ctx *gin.Context) {
	key := ctx.Param("key")

	skill, err := h.skillrepo.GetSkillByKeyRepo(key)
	if err != nil {
		response.Error(ctx, err)
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
	produceMessage(skillJSONString, "delete")
	ctx.JSON(http.StatusOK, gin.H{"message": "Skill deletation request queued"})
}

func (h *SkillHandler) PatchSkillNameHandler(ctx *gin.Context) {
	key := ctx.Param("key")
	var name PatchSkillNameRequest

	if err := ctx.BindJSON(&name); err != nil {
		fmt.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	skill, err := h.skillrepo.GetSkillByKeyRepo(key)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	skill.Name = name.Name

	// Marshal skill struct to JSON
	skillJSON, err := json.Marshal(skill)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error marshaling skill data"})
		return
	}

	// Convert JSON byte array to string
	skillJSONString := string(skillJSON)

	fmt.Println("SKill Joker ----------------------------------------------------------------> ", skillJSONString)
	produceMessage(skillJSONString, "patch-name")
	ctx.JSON(http.StatusOK, gin.H{"message": "Skill patcing name request queued"})
}

func (h *SkillHandler) PatchSkillDescriptionHandler(ctx *gin.Context) {
	key := ctx.Param("key")
	var description PatchSkillDescriptionRequest

	if err := ctx.BindJSON(&description); err != nil {
		fmt.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	skill, err := h.skillrepo.GetSkillByKeyRepo(key)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	skill.Description = description.Description

	// Marshal skill struct to JSON
	skillJSON, err := json.Marshal(skill)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error marshaling skill data"})
		return
	}

	// Convert JSON byte array to string
	skillJSONString := string(skillJSON)

	fmt.Println("SKill Joker ----------------------------------------------------------------> ", skillJSONString)
	produceMessage(skillJSONString, "patch-description")
	ctx.JSON(http.StatusOK, gin.H{"message": "Skill patcing description request queued"})
}

func (h *SkillHandler) PatchSkillLogoHandler(ctx *gin.Context) {
	key := ctx.Param("key")
	var logo PatchSkillLogoRequest

	if err := ctx.BindJSON(&logo); err != nil {
		fmt.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	skill, err := h.skillrepo.GetSkillByKeyRepo(key)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	skill.Logo = logo.Logo

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
	produceMessage(skillJSONString, "patch-logo")
	ctx.JSON(http.StatusOK, gin.H{"message": "Skill patcing logo request queued"})
}
