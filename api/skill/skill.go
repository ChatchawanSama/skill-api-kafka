package skill

type Skill struct {
	Key         string   `json:"key"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Logo        string   `json:"logo"`
	Tags        []string `json:"tags"`
}

type PatchSkillNameRequest struct {
	Name string `json:"name"`
}

type PatchSkillDescriptionRequest struct {
	Description string `json:"description"`
}

type PatchSkillLogoRequest struct {
	Logo string `json:"logo"`
}

type PatchSkillTagsRequest struct {
	Tags []string `json:"tags"`
}
