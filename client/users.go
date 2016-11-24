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

type UserItems struct {
        Items []int `json:"items"`
}

type UserHistoryService struct {
        client *http.Client
        key string
}

func NewUserHistoryService(apiKey string) *UserHistoryService {
        tr := &http.Transport{
                TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }
        return &UserHistoryService {
                client: &http.Client{Transport: tr},
                key: apiKey,
        }
}

func (service *UserHistoryService) UploadUserViewedItems(userId int, req UserItems) {
        data, _ := json.Marshal(req)
        user := strconv.Itoa(userId)
        user = url.QueryEscape(user)
        url := fmt.Sprintf(UsersAPI + "%s" + "/viewed_items", user)
        request := NewRequest("POST", url, bytes.NewBuffer(data))
        request = request.AddHeaders(service.key)
        resp, err := service.client.Do(request)
        if err != nil {
                panic(err)
        }
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)

        fmt.Println(string(body))
}

func (service *UserHistoryService) UploadUserPurchasedItems(userId int, req UserItems) {
        data, _ := json.Marshal(req)
        user := strconv.Itoa(userId)
        user = url.QueryEscape(user)
        url := fmt.Sprintf(UsersAPI + "%s" + "/purchased_items", user)
        request := NewRequest("POST", url, bytes.NewBuffer(data))
        request = request.AddHeaders(service.key)
        resp, err := service.client.Do(request)
        if err != nil {
                panic(err)
        }
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)

        fmt.Println(string(body))
}