package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

//结构体调用环境配置
type Config struct {
	AppName string    `json:"app_name"`
	AppNode string    `json:"app_node"`
	AppHost string    `json:"app_host"`
	AppPort string    `json:"app_port"`
	Sms     SmsConfig `json:"sms"`
}

//结构体调用阿里云短信服务相关信息
type SmsConfig struct {
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	RegionId     string `json:"region_id"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
}

var _cfg *Config = nil

func GetConfig() *Config {
	return _cfg
}

//读取环境配置
func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, err

}
