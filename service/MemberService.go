package service

import (
	"CloudRestaurant/tool"
	"encoding/json"
	"fmt"
	dysmsapi "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
	"math/rand"
	"time"
)

type MemberService struct {
}

func (ms *MemberService) Sendcode(phone string) bool {
	//1.产生验证码
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	//2.调用阿里云sdk 完成发送
	config := tool.GetConfig().Sms
	client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AppKey, config.AppSecret)
	if err != nil {
		log.Fatalln(err.Error())
		return false
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = config.SignName
	request.TemplateCode = config.TemplateCode
	request.PhoneNumbers = phone
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)

	response, err := client.SendSms(request)
	fmt.Println(response)
	if err != nil {
		log.Fatalln(err.Error())
		return false
	}
	//3.接受返回结果冰判断发送状态
	if response.Code == "OK" {
		return true
	}

	return false
}
