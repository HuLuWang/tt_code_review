package deepseekx

type DeepSeekModel string

const (
	defaultHost               = "https://api.deepseek.com"
	DeepSeekV3  DeepSeekModel = "deepseek-01-chat"
	DeepSeekR1  DeepSeekModel = "deepseek-reasoner"
)

type Config struct {
	ApiKey string
	Host   string
	Model  DeepSeekModel
}
