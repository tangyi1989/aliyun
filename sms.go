package aliyun

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gopub/log"
	"github.com/gopub/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

type SMSSendRequest struct {
	PhoneNumbers   []string          //Required
	TemplateCode   string            //Required
	TemplateParams map[string]string //Optional
	SignName       string            //Required
}

type SMSSendResponse struct {
	RequestID string `json:"RequestId"`
	Code      string
	Message   string
	BizID     string `json:"BizId"`
}

func (c *Client) SMSSend(req *SMSSendRequest) (*SMSSendResponse, error) {
	params := map[string]string{
		"Action":        "SendSms",
		"Version":       "2017-05-25",
		"RegionId":      "cn-hangzhou",
		"PhoneNumbers":  strings.Join(req.PhoneNumbers, ","),
		"TemplateCode":  req.TemplateCode,
		"TemplateParam": utils.JSONMarshalStr(req.TemplateParams),
		"SignName":      req.SignName,
	}

	params = c.makeRequestParams("POST", params)
	buf := new(bytes.Buffer)
	for k, v := range params {
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(v)
		buf.WriteString("&")
	}
	resp, err := http.Post("http://dysmsapi.aliyuncs.com", utils.MIMEPOSTForm, buf)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer resp.Body.Close()
	var result *SMSSendResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Error(err)
	}
	return result, err
}

type SMSQueryRequest struct {
	PhoneNumber string //Required
	BizID       string //Optional
	SendDate    string //Required yyyyMMdd, e.g. 20170709
	PageSize    int    //Required
	CurrentPage int    //Required
}

type SMSQueryResponse struct {
	RequestID      string `json:"RequestId"`
	Code           string
	Message        string
	TotalCount     int
	TotalPage      int
	SendDetailDTOs map[string]interface{}
}

func (c *Client) SMSQuerySendDetails(req *SMSQueryRequest) (*SMSQueryResponse, error) {
	params := map[string]string{
		"Action":      "QuerySendDetails",
		"Version":     "2017-05-25",
		"RegionId":    "cn-hangzhou",
		"PhoneNumber": req.PhoneNumber,
		"BizId":       req.BizID,
		"SendDate":    req.SendDate,
		"PageSize":    fmt.Sprint(req.PageSize),
		"CurrentPage": fmt.Sprint(req.CurrentPage),
	}

	params = c.makeRequestParams("POST", params)
	buf := new(bytes.Buffer)
	for k, v := range params {
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(v)
		buf.WriteString("&")
	}
	resp, err := http.Post("http://dysmsapi.aliyuncs.com", utils.MIMEPOSTForm, buf)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer resp.Body.Close()
	var result *SMSQueryResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Error(err)
	}
	return result, err
}
