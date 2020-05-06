package structs

type VoiceVerificationReturn struct {
	Message        string  `json:"message"`
	Status         int     `json:"status"`
	Text           string  `json:"text"`
	TextConfidence float64 `json:"textConfidence"`
	TimeTaken      string  `json:"timeTaken"`
	ResponseCode   string  `json:"responseCode"`
}

type VoiceVerificationByUrlReturn struct {
	Message        string  `json:"message"`
	Status         int     `json:"status"`
	Text           string  `json:"text"`
	TextConfidence float64 `json:"textConfidence"`
	TimeTaken      string  `json:"timeTaken"`
	ResponseCode   string  `json:"responseCode"`
}

type FaceVerificationReturn struct {
	Message        string  `json:"message"`
	Status         int     `json:"status"`
	FaceConfidence float64 `json:"faceConfidence"`
	TimeTaken      string  `json:"timeTaken"`
	ResponseCode   string  `json:"responseCode"`
}

type FaceVerificationByUrlReturn struct {
	Message        string  `json:"message"`
	Status         int     `json:"status"`
	FaceConfidence float64 `json:"faceConfidence"`
	TimeTaken      string  `json:"timeTaken"`
	ResponseCode   string  `json:"responseCode"`
}

type VideoVerificationReturn struct {
	Message         string  `json:"message"`
	Status          int     `json:"status"`
	VoiceConfidence float64 `json:"voiceConfidence"`
	FaceConfidence  float64 `json:"faceConfidence"`
	Text            string  `json:"text"`
	TextConfidence  float64 `json:"textConfidence"`
	TimeTaken       string  `json:"timeTaken"`
	ResponseCode    string  `json:"responseCode"`
}

type VideoVerificationByUrlReturn struct {
	Message         string  `json:"message"`
	Status          int     `json:"status"`
	VoiceConfidence float64 `json:"voiceConfidence"`
	FaceConfidence  float64 `json:"faceConfidence"`
	Text            string  `json:"text"`
	TextConfidence  float64 `json:"textConfidence"`
	TimeTaken       string  `json:"timeTaken"`
	ResponseCode    string  `json:"responseCode"`
}
