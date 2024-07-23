package skill

import (
	"fmt"
)

type SkillRepo interface {
	PostSkillRepo(Skill) error
	PutSkillByKeyRepo(Skill) error
	DeleteSkillByKeyRepo(string) error
	PatchSkillNameRepo(string, string) error
	PatchSkillDescriptionRepo(string, string) error
	PatchSkillLogoRepo(string, string) error
	PatchSkillTagsRepo(string, []string) error
}

type SkillHandler struct {
	skillrepo SkillRepo
}

func NewSkillHandler(skillrepo SkillRepo) *SkillHandler {
	return &SkillHandler{skillrepo: skillrepo}
}

func (h *SkillHandler) PostSkillHandler(skill Skill) {
	err := h.skillrepo.PostSkillRepo(skill)
	if err != nil {
		fmt.Println("Error saving skill to database:", err)
		return
	}
}

func (h *SkillHandler) PutSkillByKeyHandler(skill Skill) {
	err := h.skillrepo.PutSkillByKeyRepo(skill)
	if err != nil {
		fmt.Println("Error updating skill to database:", err)
		return
	}
}

func (h *SkillHandler) DeleteSkillByKeyHandler(skill Skill) {
	err := h.skillrepo.DeleteSkillByKeyRepo(skill.Key)
	if err != nil {
		fmt.Println("Error updating skill to database:", err)
		return
	}
}

func (h *SkillHandler) PatchSkillNameHandler(skill Skill) {
	err := h.skillrepo.PatchSkillNameRepo(skill.Key, skill.Name)
	if err != nil {
		fmt.Println("Error patching skill name to database:", err)
		return
	}
}

func (h *SkillHandler) PatchSkillDescriptionHandler(skill Skill) {
	err := h.skillrepo.PatchSkillDescriptionRepo(skill.Key, skill.Name)
	if err != nil {
		fmt.Println("Error patching skill name to database:", err)
		return
	}
}

func (h *SkillHandler) PatchSkillLogoHandler(skill Skill) {
	err := h.skillrepo.PatchSkillLogoRepo(skill.Key, skill.Logo)
	if err != nil {
		fmt.Println("Error patching skill logo to database:", err)
		return
	}
}

func (h *SkillHandler) PatchSkillTagsHandler(skill Skill) {
	err := h.skillrepo.PatchSkillTagsRepo(skill.Key, skill.Tags)
	if err != nil {
		fmt.Println("Error patching skill tags to database:", err)
		return
	}
}
