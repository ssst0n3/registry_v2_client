package response

type Catalog struct {
	Repositories []string `json:"repositories"`
	Next         string   `json:"next"`
}
