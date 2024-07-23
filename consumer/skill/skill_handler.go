package skill

import (
	"encoding/json"
	"fmt"
)

type SkillRepo interface {
	PostSkillRepo(Skill) error
	PutSkillByKeyRepo(Skill) error
}

type SkillHandler struct {
	skillrepo SkillRepo
}

func NewSkillHandler(skillrepo SkillRepo) *SkillHandler {
	return &SkillHandler{skillrepo: skillrepo}
}

func (h *SkillHandler) PostSkillHandler(value string) {
	var skill Skill
	
	// Unmarshal the JSON string into a Skill struct
	err := json.Unmarshal([]byte(value), &skill)
	if err != nil {
		// Handle unmarshaling error
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Pass the Skill struct to the repository method
	err = h.skillrepo.PostSkillRepo(skill)
	if err != nil {
		fmt.Println("Error saving skill to database:", err)
		return
	}
}


func (h *SkillHandler) PutSkillByKeyHandler(value string) {
	var skill Skill
	
	// Unmarshal the JSON string into a Skill struct
	err := json.Unmarshal([]byte(value), &skill)
	if err != nil {
		// Handle unmarshaling error
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Pass the Skill struct to the repository method
	err = h.skillrepo.PutSkillByKeyRepo(skill)
	if err != nil {
		fmt.Println("Error updating skill to database:", err)
		return
	}
}