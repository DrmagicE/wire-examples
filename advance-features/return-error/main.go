package main

import (
	"errors"
	"fmt"
)

// Config 配置
type Config struct {
	// RemoteAddr 连接的远程地址
	RemoteAddr string
}
// APIClient API客户端
type APIClient struct {
	c Config
}
// NewAPIClient  APIClient构造函数，如果入参校验失败，返回错误原因
func NewAPIClient(c Config) (*APIClient,error) { //<-- 第二个参数设置成error
	if c.RemoteAddr == "" {
		return nil, errors.New("没有设置远程地址")
	}
	return &APIClient{
		c:c,
	},nil
}
// Service
type Service struct {
	client *APIClient
}
// NewService Service构造函数
func NewService(client *APIClient) *Service{
	return &Service{
		client:client,
	}
}

// go run .
func main() {
	fmt.Println(InitializeClient(Config{""}))
}
