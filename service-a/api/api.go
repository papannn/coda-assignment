package api

type ExampleRequest struct {
	Game    string `json:"game"`
	GamerID string `json:"gamerID"`
	Points  int64  `json:"points"`
}

type ExampleResponse struct {
	Game    string `json:"game"`
	GamerID string `json:"gamerID"`
	Points  int64  `json:"points"`
}
