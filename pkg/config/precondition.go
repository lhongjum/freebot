package config

type MatchLabel struct {
	BasePrefix   string `json:"base_prefix"`
	TargetPrefix string `json:"target_prefix"`
}

type Precondition struct {
	IsAuthor            bool         `json:"is_author"`
	RequiredRoles       []string     `json:"required_roles"`
	RequiredLabels      []string     `json:"required_labels"`
	RequiredLabelPrefix []string     `json:"required_label_prefix"`
	MatchLabels         []MatchLabel `json:"match_labels"`
}
