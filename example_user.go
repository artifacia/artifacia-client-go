package main

import (
        "encoding/json"
        "github.com/artifacia/go-artifacia/src/client"
)

func main() {
        userData := `{
           "items":[
              2,
              3,
              7
           ]
        }`
        byt := []byte(userData)
        var dat client.UserItems
        if err := json.Unmarshal(byt, &dat); err != nil {
                panic(err)
            }
        apiKey := "your_api_key"
        client := client.NewClient(apiKey)
        userId := 12
        client.Users.UploadUserViewedItems(userId, dat)
        client.Users.UploadUserPurchasedItems(userId, dat)
}