package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func HTTP(reqURL string, header map[string]string,
	timeout time.Duration, body, method string) (respBody string, err error) {
	timeBegin := time.Now()

	defer func() {
		logrus.Infof("url:%s, cost:%d ms, body:%s", reqURL, time.Since(timeBegin).Milliseconds(), body)
	}()

	logrus.Infof("req:%s", reqURL)

	newReq, err := http.NewRequest(method, reqURL, strings.NewReader(body))
	if err != nil {
		err = errors.Wrapf(err, "NewRequest error:%s", reqURL)
		return "", err
	}

	if header != nil {
		for k, v := range header {
			if strings.EqualFold(k, "host") {
				newReq.Host = v
			}
			newReq.Header.Set(k, v)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	newReq = newReq.WithContext(ctx)
	logrus.Tracef("newReq:%+v", newReq)

	// 忽略对证书的校验
	tr := &http.Transport{
		DisableKeepAlives: true,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
	}

	newResp, err := (&http.Client{
		Transport: tr,
	}).Do(newReq)

	if err != nil {
		err = errors.Wrapf(err, "request error:%s", reqURL)
		return "", err
	}
	defer newResp.Body.Close()

	if newResp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("statusCode :%d,respBody: %s ", newResp.StatusCode, respBody)
	}

	newBody, err := ioutil.ReadAll(newResp.Body)
	return string(newBody), err
}

func main() {
	finish_api := fmt.Sprintf(`https://cat-match.easygame2021.com/sheep/v1/game/game_over?rank_score=1&rank_state=%d&rank_time=%d&rank_role=1&skin=1`, 1, 30)

	//for i:=0;i<1000000;i++ {
	respBody, err := HTTP(finish_api, map[string]string{
		"Host":       "cat-match.easygame2021.com",
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36 Edg/105.0.1343.33",
		"t":          "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTQ0NDQ2NzIsIm5iZiI6MTY2MzM0MjQ3MiwiaWF0IjoxNjYzMzQwNjcyLCJqdGkiOiJDTTpjYXRfbWF0Y2g6bHQxMjM0NTYiLCJvcGVuX2lkIjoiIiwidWlkIjo1MDgxNTE2MSwiZGVidWciOiIiLCJsYW5nIjoiIn0.WR7LfP1o0wnEPAl0joiLhxV8P1io2z8b6sZyLOjzDkw",
	}, 30*time.Second, "", http.MethodGet)
	if err != nil {
		err = errors.Wrapf(err, "request error")
		logrus.Error(err.Error())
		return
	}

	fmt.Println("success", respBody)
	//}

}
