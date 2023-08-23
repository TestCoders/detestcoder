package jira

type IssueDescriptionResponse struct {
	Expand string `json:"expand"`
	Id     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields struct {
		Description struct {
			Version int    `json:"version"`
			Type    string `json:"type"`
			Content []struct {
				Type    string `json:"type"`
				Content []struct {
					Type  string `json:"type"`
					Text  string `json:"text"`
					Marks []struct {
						Type string `json:"type"`
					} `json:"marks,omitempty"`
				} `json:"content"`
			} `json:"content"`
		} `json:"description"`
	} `json:"fields"`
}
