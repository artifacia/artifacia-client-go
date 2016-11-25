package client

import (
        "bytes"
        "crypto/tls"
        "encoding/json"
        "fmt"
        "io/ioutil"
        "net/http"
)

type UploadRequest struct {
        Features []string `json:"features"`
        Details []struct {
                Category string `json:"category"`
                Color string `json:"color"`
                URL string `json:"url"`
                Length string `json:"length"`
                Material string `json:"material"`
                Pattern string `json:"pattern"`
                Prodid int `json:"prodid"`
                ItemHighLevelCat string `json:"item_high_level_cat,omitempty"`
        } `json:"details"`
}

type DeleteRequest struct {
        Prodids []int `json:"prodids"`
}

type ItemService struct {
        client *http.Client
        key string
}

func NewItemService(apiKey string) *ItemService {
        tr := &http.Transport{
                TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }
        return &ItemService {
                client: &http.Client{Transport: tr},
                key: apiKey,
        }
}

func (service *ItemService) UploadItemData(req UploadRequest) {
        data, _ := json.Marshal(req)
        request := NewRequest("POST", ItemsAPI, bytes.NewBuffer(data))
        request = request.AddHeaders(service.key)
        resp, err := service.client.Do(request)
        if err != nil {
                panic(err)
        }
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)

        fmt.Println(string(body))
}

func (service *ItemService) DeleteItemData(req DeleteRequest){
        data, _ := json.Marshal(req)
        request := NewRequest("POST", ItemsAPI, bytes.NewBuffer(data))
        request = request.AddHeaders(service.key)
        resp, err := service.client.Do(request)
        if err != nil {
                panic(err)
        }
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)

        fmt.Println(string(body))
}
