package structs

type VoiceIdentificationReturn struct {
	Message        string  `json:"message"`
	UserId         string  `json:"userId"`
	GroupId        string  `json:"groupId"`
	Confidence     float64 `json:"confidence"`
	Status         int     `json:"status"`
	Text           string  `json:"text"`
	TextConfidence float64 `json:"textConfidence"`
	TimeTaken      string  `json:"timeTaken"`
	ResponseCode   string  `json:"responseCode"`
}

type VoiceIdentificationByUrlReturn struct {
	Message        string  `json:"message"`
	UserId         string  `json:"userId"`
	GroupId        string  `json:"groupId"`
	Confidence     float64 `json:"confidence"`
	Status         int     `json:"status"`
	Text           string  `json:"text"`
	TextConfidence float64 `json:"textConfidence"`
	TimeTaken      string  `json:"timeTaken"`
	ResponseCode   string  `json:"responseCode"`
}

type FaceIdentificationReturn struct {
	Message        string  `json:"message"`
	UserId         string  `json:"userId"`
	GroupId        string  `json:"groupId"`
	Status         int     `json:"status"`
	FaceConfidence float64 `json:"faceConfidence"`
	TimeTaken      string  `json:"timeTaken"`
	ResponseCode   string  `json:"responseCode"`
}

type FaceIdentificationByUrlReturn struct {
	Message        string  `json:"message"`
	UserId         string  `json:"userId"`
	GroupId        string  `json:"groupId"`
	Status         int     `json:"status"`
	FaceConfidence float64 `json:"faceConfidence"`
	TimeTaken      string  `json:"timeTaken"`
	ResponseCode   string  `json:"responseCode"`
}

type VideoIdentificationReturn struct {
	Message         string  `json:"message"`
	UserId          string  `json:"userId"`
	GroupId         string  `json:"groupId"`
	Status          int     `json:"status"`
	VoiceConfidence float64 `json:"voiceConfidence"`
	FaceConfidence  float64 `json:"faceConfidence"`
	Text            string  `json:"text"`
	TextConfidence  float64 `json:"textConfidence"`
	TimeTaken       string  `json:"timeTaken"`
	ResponseCode    string  `json:"responseCode"`
}

type VideoIdentificationByUrlReturn struct {
	Message         string  `json:"message"`
	UserId          string  `json:"userId"`
	GroupId         string  `json:"groupId"`
	Status          int     `json:"status"`
	VoiceConfidence float64 `json:"voiceConfidence"`
	FaceConfidence  float64 `json:"faceConfidence"`
	Text            string  `json:"text"`
	TextConfidence  float64 `json:"textConfidence"`
	TimeTaken       string  `json:"timeTaken"`
	ResponseCode    string  `json:"responseCode"`
}
