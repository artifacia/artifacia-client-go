package main

import (
        "encoding/json"
        "github.com/artifacia/artifacia-client-go/client"
)

func main() {
        uploadData := `{
           "features":["vr"],
           "details":[
                  {
                         "category":"Dress",
                         "color":"NA",
                         "url":"http://img6a.flixcart.com/image/dress/k/7/x/foj6000204-grey-folklore-xxl-400x400-imaeg7gmznjjcjgf.jpeg",
                         "length":"Maxi/Full Length",
                         "material":"100% Polyester",
                         "pattern":"Solid",
                         "prodid":1
                  }
                 ]
                }`
      deleteData := `{
         "prodids":[
              1
         ]
      }`
        byt1 := []byte(uploadData)
        byt2 := []byte(deleteData)
        var dat1 client.UploadRequest
        var dat2 client.DeleteRequest
        if err1 := json.Unmarshal(byt1, &dat1); err1 != nil {
            panic(err1)
        }
        if err2 := json.Unmarshal(byt2, &dat2); err2 != nil {
            panic(err2)
        }
        apiKey := "your_api_key"
        client := client.NewClient(apiKey)
        client.Items.UploadItemData(dat1)
        client.Items.DeleteItemData(dat2)
}
