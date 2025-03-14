package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tt_code_review/common/deepseekx"
)

func main() {
	// 创建一个 DeepSeek 客户端实例
	client, err := deepseekx.NewClient(&deepseekx.Config{
		ApiKey: "your api key", // 替换为你的 API 密钥
	})
	if err != nil {
		fmt.Printf("创建客户端失败: %v\n", err)
		return
	}
	// 创建一个消息列表
	messages := []deepseekx.CompletionsMessage{}

	// 创建一个扫描器来读取用户输入
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("欢迎使用 DeepSeek 聊天机器人！输入 'exit' 退出。")

	for {
		fmt.Print("你: ")
		scanner.Scan()
		userInput := scanner.Text()

		if strings.ToLower(userInput) == "exit" {
			break
		}

		// 将用户输入添加到消息列表中
		messages = append(messages, deepseekx.CompletionsMessage{
			Role:    "user",
			Content: userInput,
		})

		// 创建一个对话补全请求
		request := &deepseekx.CompletionsRequest{
			Messages: messages,
			Model:    "deepseek-chat",
		}

		// 调用 CreateChatCompletion 方法获取回复
		response, err := client.CreateChatCompletion(request)
		if err != nil {
			fmt.Printf("请求失败: %v\n", err)
			continue
		}

		// 获取助手的回复
		if len(response.Choices) > 0 {
			assistantMessage := response.Choices[0].Message.Content
			fmt.Printf("助手: %s\n", assistantMessage)

			// 将助手的回复添加到消息列表中
			messages = append(messages, deepseekx.CompletionsMessage{
				Role:    "assistant",
				Content: assistantMessage,
			})
		}
	}
}
