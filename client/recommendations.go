package client

import (
        "bytes"
        "crypto/tls"
        "encoding/json"
        "fmt"
        "io/ioutil"
        "net/http"
        "net/url"
        "strconv"
)

type RecommendationService struct {
        client *http.Client
        key string
}

func NewRecommendationService(apiKey string) *RecommendationService {
        tr := &http.Transport{
                TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }
        return &RecommendationService {
                client: &http.Client{Transport: tr},
                key: apiKey,
        }
}

func (service *RecommendationService) GetVisualRecommendations(prodId string, num int, filters ...string) {
        prodId = url.QueryEscape(prodId)
        number := strconv.Itoa(num)
        number = url.QueryEscape(number)
        url := fmt.Sprintf(RecommendationsAPI + "similar/" +"%s" + "/" +"%s", prodId, number)
        var reqBody *bytes.Buffer
        if len(filters) > 0 {
                data, _ := json.Marshal(filters)
                reqBody = bytes.NewBuffer(data)
        } else {
                var byt []byte
                reqBody = bytes.NewBuffer(byt)
        }
        request := NewRequest("POST", url, reqBody)
        request = request.AddHeaders(service.key)
        resp, err := service.client.Do(request)
        if err != nil {
                panic(err)
        }
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)

        fmt.Println(string(body))
}

func (service *RecommendationService) GetCprRecommendations(prodId string, num int) {
        prodId = url.QueryEscape(prodId)
        number := strconv.Itoa(num)
        number = url.QueryEscape(number)
        url := fmt.Sprintf(RecommendationsAPI + "collections/" +"%s" + "/" + "%s", prodId, number)
        request := NewRequest("GET", url, nil)
        request = request.AddHeaders(service.key)
        resp, err := service.client.Do(request)
        if err != nil {
                panic(err)
        }
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)

        fmt.Println(string(body))
}

func (service *RecommendationService) GetPersonalizedRecommendations(userId int, num int) {
        user := strconv.Itoa(userId)
        user = url.QueryEscape(user)
        number := strconv.Itoa(num)
        number = url.QueryEscape(number)
        url := fmt.Sprintf(RecommendationsAPI + "user/" + "%s" + "/" + "%s", user, number)
        request := NewRequest("GET", url, nil)
        request = request.AddHeaders(service.key)
        resp, err := service.client.Do(request)
        if err != nil {
                panic(err)
        }
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)

        fmt.Println(string(body))
}