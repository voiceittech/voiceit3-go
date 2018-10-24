package structs

type VoiceVerificationReturn struct {
	Message        string  `json:"message"`
	Status         int     `json:"status"`
	Confidence     float32 `json:"confidence"`
	Text           string  `json:"text"`
	TextConfidence float32 `json:"textConfidence"`
	TimeTaken      string  `json:"timeTaken"`
	ResponseCode   string  `json:"responseCode"`
	// UserId         string  `json:"userId"`
}

type VoiceVerificationByUrlReturn struct {
	Message        string  `json:"message"`
	Status         int     `json:"status"`
	Confidence     float32 `json:"confidence"`
	Text           string  `json:"text"`
	TextConfidence float32 `json:"textConfidence"`
	TimeTaken      string  `json:"timeTaken"`
	ResponseCode   string  `json:"responseCode"`
	// UserId         string  `json:"userId"`
}

type FaceVerificationReturn struct {
	Message        string  `json:"message"`
	Status         int     `json:"status"`
	FaceConfidence float32 `json:"faceConfidence"`
	TimeTaken      string  `json:"timeTaken"`
	ResponseCode   string  `json:"responseCode"`
	// UserId         string  `json:"userId"`
}

type FaceVerificationByUrlReturn struct {
	Message        string  `json:"message"`
	Status         int     `json:"status"`
	FaceConfidence float32 `json:"faceConfidence"`
	TimeTaken      string  `json:"timeTaken"`
	ResponseCode   string  `json:"responseCode"`
	// UserId         string  `json:"userId"`
}

type VideoVerificationReturn struct {
	Message         string  `json:"message"`
	Status          int     `json:"status"`
	VoiceConfidence float32 `json:"voiceConfidence"`
	FaceConfidence  float32 `json:"faceConfidence"`
	Text            string  `json:"text"`
	TextConfidence  float32 `json:"textConfidence"`
	TimeTaken       string  `json:"timeTaken"`
	ResponseCode    string  `json:"responseCode"`
	// UserId          string  `json:"userId"`
}

type VideoVerificationByUrlReturn struct {
	Message         string  `json:"message"`
	Status          int     `json:"status"`
	VoiceConfidence float32 `json:"voiceConfidence"`
	FaceConfidence  float32 `json:"faceConfidence"`
	Text            string  `json:"text"`
	TextConfidence  float32 `json:"textConfidence"`
	TimeTaken       string  `json:"timeTaken"`
	ResponseCode    string  `json:"responseCode"`
	// UserId          string  `json:"userId"`
}
