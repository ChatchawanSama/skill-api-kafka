package skill

import (
	"encoding/json"
	"fmt"
)

type SkillRepo interface {
	PostSkillByKey(Skill) error
}

type SkillHandler struct {
	skillrepo SkillRepo
}

func NewSkillHandler(skillrepo SkillRepo) *SkillHandler {
	return &SkillHandler{skillrepo: skillrepo}
}

func (h *SkillHandler) PostSkillByKey(value string) {
	var skill Skill
	
	// Unmarshal the JSON string into a Skill struct
	err := json.Unmarshal([]byte(value), &skill)
	if err != nil {
		// Handle unmarshaling error
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Pass the Skill struct to the repository method
	err = h.skillrepo.PostSkillByKey(skill)
	if err != nil {
		fmt.Println("Error saving skill to database:", err)
		return
	}

	// Optionally, you could add further response handling here
}