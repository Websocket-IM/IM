package external

import (
	"encoding/json"
	"fmt"
	"ginchat/model"
	"net/http"
	"time"
)

// 获取用户登录的请求
func GetGithubApiRequest(url string, token *model.GithubToken) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))
	return req, nil
}

// 获取用户信息
func GetUserInfo(token *model.GithubToken) (model.User, error) {

	// 使用github提供的接口
	var userInfoUrl = "https://api.github.com/user"
	req, err := GetGithubApiRequest(userInfoUrl, token)
	if err != nil {
		return model.User{}, err
	}

	// 发送请求并获取响应
	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return model.User{}, err
	}
	fmt.Println("huanjing")
	// 将响应的数据写入userInfo中，并返回
	var userInfo = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		fmt.Println(err)
		return model.User{}, err
	}
	fmt.Println(userInfo["name"], "  这里是name")
	fmt.Println(userInfo["id"].(float64), "  这里是id")
	fmt.Sprintln(userInfo["name"])
	user := model.User{
		Nickname:      fmt.Sprintln(userInfo["name"]),
		GithubID:      fmt.Sprintln(userInfo["id"]),
		Email:         fmt.Sprintln(userInfo["email"]),
		LoginTime:     time.Now(),
		LoginOutTime:  time.Now(),
		HeartbeatTime: time.Now(),
	}
	fmt.Println("结构体user", user)
	return user, nil
}

// 获取 github token
func GetGithubToken(url string) (*model.GithubToken, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var token model.GithubToken
	if err := json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}
