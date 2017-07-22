package aliyun

import (
	"bytes"
	//"encoding/json"
	"encoding/json"
	"github.com/natande/gox"
	"io/ioutil"
	"net/http"
	"strings"
)

type SMSRequest struct {
	PhoneNumbers   []string
	TemplateCode   string
	TemplateParams map[string]string
	SignName       string
}

type SMSResponse struct {
	RequestID string
	Code      string
	Message   string
	BizID     string
}

type SMSSendDetails struct {
}

func (c *Client) SendSMS(req *SMSRequest) (*SMSResponse, error) {
	params := map[string]string{
		"Action":        "SendSms",
		"Version":       "2017-05-25",
		"RegionId":      "cn-hangzhou",
		"PhoneNumbers":  strings.Join(req.PhoneNumbers, ","),
		"TemplateCode":  req.TemplateCode,
		"TemplateParam": gox.JSONMarshalStr(req.TemplateParams),
		"SignName":      req.SignName,
	}

	params = c.MakeRequestParams("POST", params)
	buf := new(bytes.Buffer)
	for k, v := range params {
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(v)
		buf.WriteString("&")
	}
	resp, err := http.Post("http://dysmsapi.aliyuncs.com", gox.MIMEPOSTForm, buf)
	if err != nil {
		gox.LogError(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gox.LogError(err)
		return nil, err
	}
	defer resp.Body.Close()
	var result *SMSResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		gox.LogError(err)
	}
	return result, err
}

func (c *Client) QuerySMSSendDetails() *SMSSendDetails {
	return nil
}
