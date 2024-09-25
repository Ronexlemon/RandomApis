package responses

type UserResponses struct{
	Message string `json:"username"`
	Status   int `json:"status"`
	Data  map[string] interface{}  `json:"data"`
}

