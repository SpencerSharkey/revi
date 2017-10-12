package api

type ClientConfig struct {
	token     string
	tokenType string
	userAgent string
}

type Client struct {
	config *ApiConfig
}

// NewAPI - Creates a new Api given an ApiConfig
func NewAPI(config *ApiConfig) *Api {
	api := &ClientConfig{}
	api.config = config

	return api
}

// NewBotAPI - Creates new Api with default config given a Bot Token
func NewBotAPI(botToken string) *Api {
	config := &ApiConfig{}
	config.tokenType = "Bearer"
	config.token = botToken
	config.userAgent = "DiscordBot https://github.com/SpencerSharkey/revi 0.0.1"

	api := NewAPI(config)
	return api
}

// Api Request
type GetGatewayResponse struct {
	Url string
}

// RestRequest - Request
func (api *Api) RestRequest(method, url string, request interface{}, response interface{}) interface{} {

	httpClient := &http.Client{
		Timeout: time.Second * 5
	}

	request, err := http.NewRequest(method, url)

}
