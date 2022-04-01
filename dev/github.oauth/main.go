package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	clientSecret = flag.String("cs", "", "github oauth client secret")
	clientID     = flag.String("ci", "", "github oauth client id")
)

type Conf struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

type Token struct {
	AccessToken string `json:"access_token"`
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

func RefreshToken(token string) {
	rtReq := RefreshTokenReq{
		RefreshToken: token,
		ClientId:     *clientID,
		ClientSecret: *clientSecret,
		GrantType:    "refresh_token",
	}
	jsonStr, _ := json.Marshal(rtReq)
	log.Printf("RefreshToken jsonStr: %s", jsonStr)

	var userInfoUrl = "https://github.com/login/oauth/access_token"
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, strings.NewReader(string(jsonStr))); err != nil {
		log.Printf(err.Error())
		return
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	// 发送请求并获取响应
	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		log.Printf("client.Do err: %s", err.Error())
		return
	}
	defer res.Body.Close()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("RefreshToken ReadAll err: %s", err.Error())
	}
	log.Printf("RefreshToken respBody: %s", respBody)
}

// 认证并获取用户信息
func OAuth(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	// 获取 code
	code := r.URL.Query().Get("code")

	// 通过 code, 获取 token
	var tokenAuthUrl = GetTokenAuthURL(code)
	var token *Token
	if token, err = GetToken(tokenAuthUrl); err != nil {
		fmt.Println(err)
		return
	}

	//RefreshToken(token.AccessToken)

	// 通过token，获取用户信息
	var userInfo map[string]interface{}
	if userInfo, err = GetUserInfo(token); err != nil {
		fmt.Println("获取用户信息失败，错误信息为:", err)
		return
	}

	//  将用户信息返回前端
	var userInfoBytes []byte
	if userInfoBytes, err = json.Marshal(userInfo); err != nil {
		fmt.Println("在将用户信息(map)转为用户信息([]byte)时发生错误，错误信息为:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(userInfoBytes); err != nil {
		fmt.Println("在将用户信息([]byte)返回前端时发生错误，错误信息为:", err)
		return
	}

}

// 通过code获取token认证url
func GetTokenAuthURL(code string) string {
	return fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		*clientID, *clientSecret, code,
	)
}

// 获取 token
func GetToken(url string) (*Token, error) {
	// 形成请求
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var (
		httpClient = http.Client{}
		res        *http.Response
		respBody   = make([]byte, 0)
		token      Token
	)

	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}

	respBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("token: %s", string(respBody))

	// 将响应体解析为 token，并返回
	err = json.Unmarshal(respBody, &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

// 获取用户信息
func GetUserInfo(token *Token) (map[string]interface{}, error) {
	// 形成请求
	var userInfoUrl = "https://api.github.com/user" // github用户信息获取接口
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))

	// 发送请求并获取响应
	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, err
	}

	// 将响应的数据写入 userInfo 中，并返回
	var userInfo = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}

func Html(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	var (
		temp *template.Template
		err  error
	)

	dir, _ := os.Getwd()

	if temp, err = template.ParseFiles(dir + "/oauth.html"); err != nil {
		fmt.Println("读取文件失败，错误信息为:", err)
		return
	}

	// 利用给定数据渲染模板(html页面)，并将结果写入w，返回给前端
	if err = temp.Execute(w, Conf{
		ClientId:     *clientID,
		ClientSecret: *clientSecret,
		RedirectUrl:  "http://127.0.0.1:9000/oauth/callback",
	}); err != nil {
		fmt.Println("读取渲染html页面失败，错误信息为:", err)
		return
	}
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	log.Printf("UserInfo token: %s", token)
	var (
		err      error
		userInfo map[string]interface{}
	)
	if userInfo, err = GetUserInfo(&Token{AccessToken: token}); err != nil {
		fmt.Println("获取用户信息失败，错误信息为:", err)
		return
	}

	//  将用户信息返回前端
	var userInfoBytes []byte
	if userInfoBytes, err = json.Marshal(userInfo); err != nil {
		fmt.Println("在将用户信息(map)转为用户信息([]byte)时发生错误，错误信息为:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(userInfoBytes); err != nil {
		fmt.Println("在将用户信息([]byte)返回前端时发生错误，错误信息为:", err)
		return
	}
}

func main() {
	flag.Parse()
	log.Printf("clientSecrets: %s,clientID: %s", *clientSecret, *clientID)

	if *clientSecret == "" || *clientID == "" {
		log.Fatal("clientSecrets or clientID is required")
	}

	http.HandleFunc("/", Html)
	http.HandleFunc("/oauth/callback", OAuth)
	http.HandleFunc("/getUserInfo", UserInfo)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Println("监听失败，错误信息为:", err)
		return
	}

}
