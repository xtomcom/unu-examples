package unu

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiURL = "https://u.nu/api.php"
)

type Request struct {
	URL     string
	Keyword string
	Title   string
	User    *url.Userinfo
}

type Response struct {
	Code       string      `json:"code"`
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	URL        ResponseURL `json:"url"`
	Title      string      `json:"title"`
	ShortURL   string      `json:"shorturl"`
	Message    string      `json:"message"`
}

func (r *Response) String() string {
	return r.Status + ": " + r.Message + "\n" + r.ShortURL
}

type ResponseURL struct {
	IP      string `json:"ip"`
	URL     string `json:"url"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Clicks  int64  `json:"clicks,string"`
	Keyword string `json:"keyword"`
}

func Submit(request *Request) (*Response, error) {
	return SubmitWithClient(request, http.DefaultClient)
}

func SubmitWithClient(request *Request, client *http.Client) (body *Response, err error) {
	requestBody, err := buildRequestBody(request)
	if err != nil {
		return
	}
	response, err := client.Post(
		apiURL,
		"application/x-www-form-urlencoded",
		strings.NewReader(requestBody.Encode()),
	)
	if err != nil {
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	body = new(Response)
	if err = json.Unmarshal(data, body); err != nil {
		return
	}
	if body.Status == "error" {
		err = errors.New(body.Message)
	}
	return
}

func buildRequestBody(request *Request) (values *url.Values, err error) {
	if request == nil {
		err = errors.New("invalid Request")
		return
	}
	values = &url.Values{}
	if _, err = url.Parse(request.URL); err != nil {
		err = errors.New("invalid URL")
		return
	}
	values.Set("url", request.URL)
	if request.Keyword != "" {
		values.Set("keyword", request.Keyword)
	}
	if request.Title != "" {
		values.Set("title", request.Title)
	}
	values.Set("format", "json")
	values.Set("action", "shorturl")
	if request.User != nil {
		username := request.User.Username()
		if password, ok := request.User.Password(); ok {
			values.Set("username", username)
			values.Set("password", password)
		}
	}
	return
}
