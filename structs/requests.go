package structs

// Some request types have optional parameters all around. Use these structs to structure such paramters

type CreateSubAccountRequest struct {
	FirstName       string
	LastName        string
	Email           string
	Password        string
	ContentLanguage string
}

func (csar CreateSubAccountRequest) IsEmpty() bool {
	if csar.FirstName == "" && csar.LastName == "" && csar.Email == "" && csar.Password == "" && csar.ContentLanguage == "" {
		return true
	}
	return false
}
