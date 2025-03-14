package deepseekx

// CompletionsRequest - 对话补全请求参数
type CompletionsRequest struct {
    Messages         []CompletionsMessage      `json:"messages"`                    // 对话消息列表，必填
    Model            string                    `json:"model"`                       // 模型名称，必填 [deepseek-01-chat, deepseek-reasoner]
    FrequencyPenalty float64                   `json:"frequency_penalty,omitempty"` // 降低模型重复相同内容的可能性，范围 [-2.0, 2.0]，可选
    MaxTokens        int                       `json:"max_tokens,omitempty"`        // 限制一次请求中生成的最大 token 数，默认 4096，范围 [1, 8192]，可选
    PresencePenalty  float64                   `json:"presence_penalty,omitempty"`  // 增加模型谈论新主题的可能性，范围 [-2.0, 2.0]，可选
    ResponseFormat   CompletionsResponseFormat `json:"response_format,omitempty"`   // 输出结果格式，默认 text，[text, json_object]，可选
    Stop             []string                  `json:"stop,omitempty"`              // 停止生成 token 的条件，最多包含 16 个字符串，可选
    Stream           bool                      `json:"stream,omitempty"`            // 是否以流式方式返回结果，可选
    StreamOptions    interface{}               `json:"stream_options,omitempty"`    // 流式选项，可选
    Temperature      float64                   `json:"temperature,omitempty"`       // 采样温度，范围 [0, 2]，可选
    TopP             float64                   `json:"top_p,omitempty"`             // 调节采样概率，范围 (0, 1]，可选
    Tools            interface{}               `json:"tools,omitempty"`             // 工具配置，可选
    ToolChoice       string                    `json:"tool_choice,omitempty"`       // 工具选择策略，可选
    LogProbS         bool                      `json:"logprobs,omitempty"`          // 是否返回输出 token 的对数概率，可选
    TopLogProbS      interface{}               `json:"top_logprobs,omitempty"`      // 返回每个位置 top N 的 token 概率，范围 [0, 20]，可选
}

// CompletionsMessage - 对话消息
type CompletionsMessage struct {
	Content string `json:"content"`        // 消息内容，必填
	Role    string `json:"role"`           // 消息角色，必填
	Name    string `json:"name,omitempty"` // 消息名称，可选
}

// CompletionsResponseFormat - 响应格式
type CompletionsResponseFormat struct {
	Type string `json:"type"` // 响应格式类型，[text, json_object]，必填
}

// CompletionsResponse - 对话补全响应
type CompletionsResponse struct {
	Id                string              `json:"id"`                 // 请求 ID，必填
	Choices           []CompletionsChoice `json:"choices"`            // 响应选择列表，必填
	Created           int                 `json:"created"`            // 创建时间戳，必填
	Model             string              `json:"model"`              // 使用的模型名称，必填
	SystemFingerprint string              `json:"system_fingerprint"` // 系统指纹，必填
	Object            string              `json:"object"`             // 对象类型，必填
	Usage             Usage               `json:"usage"`              // 使用情况统计，必填
}

// CompletionsChoice - 响应选择项
type CompletionsChoice struct {
	FinishReason string    `json:"finish_reason"`      // 完成原因，必填
	Index        int       `json:"index"`              // 索引，必填
	Message      Message   `json:"message"`            // 消息内容，必填
	LogProbS     *LogProbS `json:"logprobs,omitempty"` // 对数概率信息，可选
}

// Message - 消息内容
type Message struct {
	Content          string     `json:"content"`                     // 消息内容，必填
	ReasoningContent string     `json:"reasoning_content,omitempty"` // 推理内容，可选
	ToolCalls        []ToolCall `json:"tool_calls,omitempty"`        // 工具调用列表，可选
	Role             string     `json:"role"`                        // 消息角色，必填
}

// ToolCall - 工具调用
type ToolCall struct {
	Id       string   `json:"id"`       // 工具调用 ID，必填
	Type     string   `json:"type"`     // 工具类型，必填
	Function Function `json:"function"` // 工具函数信息，必填
}

// Function - 工具函数
type Function struct {
	Name      string `json:"name"`      // 函数名称，必填
	Arguments string `json:"arguments"` // 函数参数，必填
}

// LogProbS - 对数概率信息
type LogProbS struct {
	Content []LogprobContent `json:"content"` // 对数概率内容列表，必填
}

// LogprobContent - 对数概率内容
type LogprobContent struct {
	Token       string       `json:"token"`        // Token，必填
	LogProb     int          `json:"logprob"`      // 对数概率值，必填
	Bytes       []int        `json:"bytes"`        // 字节数据，必填
	TopLogProbS []TopLogProb `json:"top_logprobs"` // 顶部对数概率列表，必填
}

// TopLogProb - 顶部对数概率
type TopLogProb struct {
	Token   string `json:"token"`   // Token，必填
	LogProb int    `json:"logprob"` // 对数概率值，必填
	Bytes   []int  `json:"bytes"`   // 字节数据，必填
}

// Usage - 使用情况统计
type Usage struct {
	CompletionTokens        int                      `json:"completion_tokens"`                   // 完成 token 数，必填
	PromptTokens            int                      `json:"prompt_tokens"`                       // 提示 token 数，必填
	PromptCacheHitTokens    int                      `json:"prompt_cache_hit_tokens,omitempty"`   // 提示缓存命中 token 数，可选
	PromptCacheMissTokens   int                      `json:"prompt_cache_miss_tokens,omitempty"`  // 提示缓存未命中 token 数，可选
	TotalTokens             int                      `json:"total_tokens"`                        // 总 token 数，必填
	CompletionTokensDetails *CompletionTokensDetails `json:"completion_tokens_details,omitempty"` // 完成 token 详情，可选
}

// CompletionTokensDetails - 完成 token 详情
type CompletionTokensDetails struct {
	ReasoningTokens int `json:"reasoning_tokens"` // 推理 token 数，必填
}
