
package infrastructure


import (
	// "errors"
	// "fmt"
	"net/http"
	"net/url"
	// "strconv"
	"strings"

	"io/ioutil"
)


type Http struct {}


func NewHttp() *Http {
	return new(Http)
}


func (p *Http) Get(URL string, params map[string]string) (body string, err error) {

	request, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		return "", err
	}

	var q url.Values
	q = request.URL.Query()

	for key, param := range params {
		q.Add(key, param)
	}

	request.URL.RawQuery = q.Encode()
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(b), nil
}


func (p *Http) Post(URL string, params map[string]string) (body string, err error) {

	request, err := http.NewRequest("POST", URL, strings.NewReader(values.Encode()))

	if err != nil {
		return "", err
	}

	values := url.Values{}

	for key, param := range params {
		values.Add(key, param)
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
