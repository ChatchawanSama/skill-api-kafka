package skill

import (
	"github.com/IBM/sarama"
)

func ConsumeMessage(msg *sarama.ConsumerMessage, handler SkillHandler) {
	topic, key, value := msg.Topic, string(msg.Key), msg.Value
	valueStr := string(value)

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
			handler.PostSkillHandler(valueStr)
		}else if key == "put" {
			handler.PutSkillByKeyHandler(valueStr)
		}
		// else if action == "update-name" {
		// 	handler.updateNameByKeyHandler(value, skillKey)
		// } else if action == "update-description" {
		// 	handler.updateDescriptionByKeyHandler(value, skillKey)
		// } else if action == "update-logo" {
		// 	handler.updateLogoByKeyHandler(value, skillKey)
		// } else if action == "update-tags" {
		// 	handler.updateTagsByKeyHandler(value, skillKey)
		// } else if action == "delete" {
		// 	handler.deleteByKeyHandler(skillKey)
		// }
	}
}
