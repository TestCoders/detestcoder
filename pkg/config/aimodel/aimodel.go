package aimodel

type AIModel struct {
	ApiKey       string `yaml:"apikey"`
	AiModel      string `yaml:"aimodel"`
	ModelVersion string `yaml:"modelversion"`
}

func NewAiModel() *AIModel {
	return &AIModel{
		ApiKey:       "",
		AiModel:      "",
		ModelVersion: "",
	}
}

func (a *AIModel) SetApiKey(apiKey string) {
	a.ApiKey = apiKey
}

func (a *AIModel) SetModel(model string) {
	a.AiModel = model
}

func (a *AIModel) SetModelVersion(modelVersion string) {
	a.ModelVersion = modelVersion
}
