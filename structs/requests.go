package structs

// Some request types have optional parameters all around. Use these structs to structure such paramters

type CreateSubAccountRequest struct {
	FirstName       string
	LastName        string
	Email           string
	Password        string
	ContentLanguage string
	APICallId       string `json:"apiCallId"`
}
