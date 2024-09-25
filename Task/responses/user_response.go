package responses

type UserResponses struct {
	Message string                 `json:"message"`
	Status  int                    `json:"status"`
	Data    map[string]interface{} `json:"data"`
}
