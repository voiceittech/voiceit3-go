package structs

type VoiceIdentificationReturn struct {
	Message        string  `json:"message"`
	UserId         string  `json:"userId"`
	GroupId        string  `json:"groupId"`
	Confidence     float32 `json:"confidence"`
	Status         int     `json:"status"`
	Text           string  `json:"text"`
	TextConfidence float32 `json:"textConfidence"`
	TimeTaken      string  `json:"timeTaken"`
	ResponseCode   string  `json:"responseCode"`
}

type VoiceIdentificationByUrlReturn struct {
	Message        string  `json:"message"`
	UserId         string  `json:"userId"`
	GroupId        string  `json:"groupId"`
	Confidence     float32 `json:"confidence"`
	Status         int     `json:"status"`
	Text           string  `json:"text"`
	TextConfidence float32 `json:"textConfidence"`
	TimeTaken      string  `json:"timeTaken"`
	ResponseCode   string  `json:"responseCode"`
}

type FaceIdentificationReturn struct {
	Message        string  `json:"message"`
	UserId         string  `json:"userId"`
	GroupId        string  `json:"groupId"`
	Status         int     `json:"status"`
	FaceConfidence float32 `json:"faceConfidence"`
	BlinksCount    int     `json:"blinksCount"`
	TimeTaken      string  `json:"timeTaken"`
	ResponseCode   string  `json:"responseCode"`
}

type FaceIdentificationByUrlReturn struct {
	Message        string  `json:"message"`
	UserId         string  `json:"userId"`
	GroupId        string  `json:"groupId"`
	Status         int     `json:"status"`
	FaceConfidence float32 `json:"faceConfidence"`
	BlinksCount    int     `json:"blinksCount"`
	TimeTaken      string  `json:"timeTaken"`
	ResponseCode   string  `json:"responseCode"`
}

type VideoIdentificationReturn struct {
	Message         string  `json:"message"`
	UserId          string  `json:"userId"`
	GroupId         string  `json:"groupId"`
	Status          int     `json:"status"`
	VoiceConfidence float32 `json:"voiceConfidence"`
	FaceConfidence  float32 `json:"faceConfidence"`
	Text            string  `json:"text"`
	TextConfidence  float32 `json:"textConfidence"`
	BlinksCount     int     `json:"blinksCount"`
	TimeTaken       string  `json:"timeTaken"`
	ResponseCode    string  `json:"responseCode"`
}

type VideoIdentificationByUrlReturn struct {
	Message         string  `json:"message"`
	UserId          string  `json:"userId"`
	GroupId         string  `json:"groupId"`
	Status          int     `json:"status"`
	VoiceConfidence float32 `json:"voiceConfidence"`
	FaceConfidence  float32 `json:"faceConfidence"`
	Text            string  `json:"text"`
	TextConfidence  float32 `json:"textConfidence"`
	BlinksCount     int     `json:"blinksCount"`
	TimeTaken       string  `json:"timeTaken"`
	ResponseCode    string  `json:"responseCode"`
}
