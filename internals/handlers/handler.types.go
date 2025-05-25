package handlers

type ApiRegistry struct{}

type Routine struct {
	ID int		`json:"id"`
	Name string	`json:"name"`
}

type IndexResponse struct {
	Status string `json:"status"`
}