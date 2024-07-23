package skill

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
)

func ConsumeMessage(msg *sarama.ConsumerMessage, handler SkillHandler) {
	topic, key, value := msg.Topic, string(msg.Key), msg.Value
	valueStr := string(value)

	var skill Skill

	// Unmarshal the JSON string into a Skill struct
	err := json.Unmarshal([]byte(valueStr), &skill)
	if err != nil {
		// Handle unmarshaling error
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// splitKey := strings.Split(key, "-")

	// var action string
	// if len(splitKey) > 2 {
	// 	action = splitKey[0] + "-" + splitKey[1]
	// } else {
	// 	action = splitKey[0]
	// }

	// skillKey := splitKey[len(splitKey)-1]

	if topic == "skills" {
		if key == "post" {
			handler.PostSkillHandler(skill)
		} else if key == "put" {
			handler.PutSkillByKeyHandler(skill)
		} else if key == "delete" {
			handler.DeleteSkillByKeyHandler(skill)
		} else if key == "patch-name" {
			handler.PatchSkillNameHandler(skill)
		} else if key == "patch-description" {
			handler.PatchSkillDescriptionHandler(skill)
		} else if key == "patch-logo" {
			handler.PatchSkillLogoHandler(skill)
		} else if key == "patch-tags" {
			handler.PatchSkillTagsHandler(skill)
		}
	}
}
