package config

type Config struct {
	CurrentStage string   `json:"currentStage" validate:"required"`
	Developers   []string `json:"developers" validate:"required"`
	Url          string   `json:"url" validate:"required"`
	Port         string   `json:"port" validate:"required"`
	MongoUri     string   `json:"mongoUri" validate:"required"`
}

