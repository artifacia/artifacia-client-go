package client

const (
        ItemsAPI             = "https://api.artifacia.com/v1/items"
        UsersAPI             = "https://api.artifacia.com/v1/users/"
        RecommendationsAPI   = "https://api.artifacia.com/v1/recommendation/"
)

type Client struct {
        Items *ItemService
        Users *UserHistoryService
        Recommendations *RecommendationService
}

func NewClient(apiKey string) *Client {
        return &Client{
            Items: NewItemService(apiKey),
            Users: NewUserHistoryService(apiKey),
            Recommendations: NewRecommendationService(apiKey),
        }
}