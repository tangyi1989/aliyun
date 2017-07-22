package aliyun

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/natande/gox"
	"go4.org/sort"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	accessKeyID     string
	accessKeySecret string
}

func NewClient(accessKeyID, accessKeySecret string) *Client {
	return &Client{
		accessKeyID:     accessKeyID,
		accessKeySecret: accessKeySecret,
	}
}

// MakeRequestParams makes request parameters according to https://help.aliyun.com/document_detail/56189.html
func (c *Client) MakeRequestParams(httpMethod string, bizParams map[string]string) map[string]string {
	params := map[string]string{}
	for k, v := range bizParams {
		params[k] = v
	}
	delete(params, "Signature")
	params["AccessKeyId"] = c.accessKeyID
	params["SignatureMethod"] = "HMAC-SHA1"
	params["SignatureVersion"] = "1.0"
	params["SignatureNonce"] = gox.NewUUID()
	params["Format"] = "JSON"
	params["Timestamp"] = time.Now().UTC().Format("2006-01-02T15:04:05Z")

	// Sort by key
	keys := make([]string, 0, len(params))
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	buf := new(bytes.Buffer)
	for _, k := range keys {
		buf.WriteString(specialURLEncode(k))
		buf.WriteString("=")
		buf.WriteString(specialURLEncode(params[k]))
		buf.WriteString("&")
	}

	queryStr := buf.String()
	queryStr = queryStr[:len(queryStr)-1]
	gox.LogInfo(queryStr)
	buf.Reset()
	buf.WriteString(httpMethod)
	buf.WriteString("&")
	buf.WriteString(specialURLEncode("/"))
	buf.WriteString("&")
	buf.WriteString(specialURLEncode(queryStr))

	h := hmac.New(sha1.New, []byte(c.accessKeySecret+"&"))
	h.Write(buf.Bytes())
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	params["Signature"] = specialURLEncode(signature)

	return params
}

func specialURLEncode(s string) string {
	str := url.QueryEscape(s)
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)
	return str
}
