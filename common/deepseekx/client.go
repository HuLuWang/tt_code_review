package deepseekx

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

type Client struct {
	client *fasthttp.Client
	conf   *Config
	url    string
}

func NewClient(conf *Config) (*Client, error) {
	if conf == nil || conf.ApiKey == "" {
		return nil, fmt.Errorf("API key is not set")
	}
	return &Client{
		client: &fasthttp.Client{
			ReadTimeout: time.Second * 30, // 修改: 增加读取超时时间为 30 秒
		},
		conf: conf,
	}, nil
}

// CreateChatCompletion - 创建对话补全请求
func (c *Client) CreateChatCompletion(request *CompletionsRequest) (*CompletionsResponse, error) {
	// 确保 response_format.type 字段有有效的值
	if request.ResponseFormat.Type == "" {
		request.ResponseFormat.Type = "text" // 设置默认值为 "text"
	}

	// 序列化请求参数
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	// 构建 HTTP 请求
	url := "https://api.deepseek.com/chat/completions"
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.conf.ApiKey))
	req.SetBody(reqBody)

	// 发送请求
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if err := c.client.Do(req, resp); err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	// 检查响应状态码
	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.Body())
	}

	// 解析响应
	var completionsResponse CompletionsResponse
	if err := json.Unmarshal(resp.Body(), &completionsResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &completionsResponse, nil
}