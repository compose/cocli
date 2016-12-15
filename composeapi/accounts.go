package composeapi

// Account structure
type Account struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// AccountResponse holding structure
type AccountResponse struct {
	Embedded struct {
		Accounts []Account `json:"accounts"`
	} `json:"_embedded"`
}
