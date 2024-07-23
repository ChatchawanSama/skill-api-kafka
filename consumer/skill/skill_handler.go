package skill

import (
	"fmt"
)

type SkillRepo interface {
	PostSkillRepo(Skill) error
	PutSkillByKeyRepo(Skill) error
	DeleteSkillByKeyRepo(string) error
}

type SkillHandler struct {
	skillrepo SkillRepo
}

func NewSkillHandler(skillrepo SkillRepo) *SkillHandler {
	return &SkillHandler{skillrepo: skillrepo}
}

func (h *SkillHandler) PostSkillHandler(skill Skill) {
	// Pass the Skill struct to the repository method
	err := h.skillrepo.PostSkillRepo(skill)
	if err != nil {
		fmt.Println("Error saving skill to database:", err)
		return
	}
}

func (h *SkillHandler) PutSkillByKeyHandler(skill Skill) {
	// Pass the Skill struct to the repository method
	err := h.skillrepo.PutSkillByKeyRepo(skill)
	if err != nil {
		fmt.Println("Error updating skill to database:", err)
		return
	}
}

func (h *SkillHandler) DeleteSkillByKeyHandler(skill Skill) {
	// Pass the Skill struct to the repository method
	err := h.skillrepo.DeleteSkillByKeyRepo(skill.Key)
	if err != nil {
		fmt.Println("Error updating skill to database:", err)
		return
	}
}
