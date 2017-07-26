### 阿里云服务的Go SDK
1. 发送短信接口
    
        client := aliyun.NewClient("your accessKeyID", "your accessKeySecret")
        request := &SMSSendRequest{
                    PhoneNumbers:[]string{"13800000001","13800000002"},
                    TemplateCode:"your template code",
                    TemplateParams:map[string]string{"code":"123456"},
                    SignName:"测试应用"
        }
        client.SMSSend(request)
2. 查询短信发送状态接口

        SMSQuerySendDetails