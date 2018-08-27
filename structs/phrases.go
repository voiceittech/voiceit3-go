package structs

type Phrase struct {
	Text            string `json:"text"`
	ContentLanguage string `json:"contentLanguage"`
}

type GetPhrasesReturn struct {
	Message      string   `json:"message"`
	Count        int      `json:"count"`
	Status       int      `json:"status"`
	TimeTaken    string   `json:"timeTaken"`
	Phrases      []Phrase `json:"phrases"`
	ResponseCode string   `json:"responseCode"`
}
