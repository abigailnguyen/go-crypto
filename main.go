package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/abigailnguyen/go-crypto/crypto"
)

func main() {
	crypto.Encrypted4()
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("Message 1", msg1)
	case msg2 := <-c2:
		fmt.Println("Message 2", msg2)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	default:
		fmt.Println("nothing ready")
	}

	var input string
	fmt.Scanln(&input)
}

// func main() {
// 	println("GET CANONICAL REQUEST")
// 	getCanonicalRequest()
// 	getHeaders()
// }

type Request struct {
	Method   string   `json:"method"`
	Hostname string   `json:"hostname"`
	Query    struct{} `json:"query,omitempty"`
	Headers  Headers  `json:"headers"`
	Body     string   `json:"body"`
	Protocol string   `json:"protocol"`
	Path     string   `json:"path"`
}

type Headers struct {
	ContentType       string `json:"content-type"`
	ContentLength     string `json:"content-length"`
	Host              string `json:"host"`
	XAmzUserAgent     string `json:"x-amz-user-agent"`
	XAmzDate          string `json:"x-amz-date"`
	XAmzSecurityToken string `json:"x-amz-security-token"`
	XAmzContentSha256 string `json:"x-amz-content-sha256"`
	Authorization     string `json:"authorization"`
	UserAgent         string `json:"user-agent"`
}

func getSignedRequestExpected() {
	jData, err := json.Marshal(&Request{
		Method:   "POST",
		Hostname: "sts.us-east-1.amazonaws.com",
		Headers: Headers{
			ContentType:       "application/x-www-form-urlencoded",
			ContentLength:     "43",
			Host:              "sts.us-east-1.amazonaws.com",
			XAmzUserAgent:     "aws-sdk-js/3.470.0",
			UserAgent:         "aws-sdk-js/3.470.0 ua/2.0 os/darwin#23.1.0 lang/js md/nodejs#18.2.0 api/sts#3.470.0",
			XAmzDate:          "20240105T132551Z",
			XAmzSecurityToken: "FwoGZXIvYXdzEAcaDN3pP3kUH7++erdbxyL4AWEXpMny+fo6WfA69LOn4xB5vHsyt/Vps4D1FuBFJE5AT0IgIks4WcWXa+z0QouGPR8Pc3jqwvWMlt9AyrqYum/hnuZYF7IE5Kfc11xWo3op/dJj7yiJy7cVNb7al+gJBhsO/CmdL3ahXoNJXYlqqBa3Yn4Tk3WJza0cYf1jduovCUwCXT5NiWmjmKQTzqtfZ6VChDaYq74VdKWleUQADrjwqq+1NhU66ROR9kheD30WArKQpYF6v2Hd/Unw3/6N2XTruYh8ZqDtTmMi4Zugz2+uvvDKCfgPkvd9n4sJ1BZBO+awYc+oSkiiGfBuYnaco8oindiyOJvEKMGG4KwGMiv2JBWN8QjdSv9LMvUYofqfdBIWnWeAhKf1ehBtMEv/pP72xoNIo+SkF4xd",
			XAmzContentSha256: "ab821ae955788b0e33ebd34c208442ccfc2d406e2edc5e7a39bd6458fbb4f843",
			Authorization:     "AWS4-HMAC-SHA256 Credential=ASIA2EP5HGA3LCNV6JTX/20240105/us-east-1/sts/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-content-sha256;x-amz-date;x-amz-security-token;x-amz-user-agent, Signature=f21232794197118177ee34acbfc784013adf187bf4e20bdad6695da7260bfc24",
		},
		Body:     "Action=GetCallerIdentity&Version=2011-06-15",
		Protocol: "https:",
		Path:     "/",
	})
	if err != nil {
		println("Something wrong", err)
		return
	}
	println(jData)
}

func getHeaders() {
	a := map[string]any{
		"password1": "abc123",
		"password2": "correct horses battery staple",
	}
	data, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	log.Default().Print(string(data))

}

// 1.Canonical Request
func getCanonicalRequest() {
	stsRequest := &Request{
		Method:   "POST",
		Hostname: "sts.us-east-1.amazonaws.com",
		Headers: Headers{
			ContentType:   "application/x-www-form-urlencoded",
			ContentLength: "43",
			Host:          "sts.us-east-1.amazonaws.com",
			XAmzUserAgent: "aws-sdk-js/3.470.0",
			UserAgent:     "aws-sdk-js/3.470.0 ua/2.0 os/darwin#23.1.0 lang/js md/nodejs#18.2.0 api/sts#3.470.0",
		},
		Body:     "Action=GetCallerIdentity&Version=2011-06-15",
		Protocol: "https:",
		Path:     "/",
	}
	fmt.Print(stsRequest)
}
