package client

import (
        "io"
        "net/http"
)

func NewRequest(method string, url string, body io.Reader) *http.Request {
        request, err := http.NewRequest(method, url, body)
        if err != nil {
            panic(err)
            return nil
        }
        return request
}

func (req *http.Request) AddHeaders(key string) *http.Request {
        req.Header.Add("Content-Type", "application/json")
        req.Header.Add("api_key", key)
        return req
}