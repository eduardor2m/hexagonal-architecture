package models

type Project struct {
	Name     string `json:"name"`
	Tech     string `json:"tech"`
	Repo     string `json:"repo"`
	Deployed string `json:"deployed"`
}
