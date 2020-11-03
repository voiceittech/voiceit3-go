package structs

type VoiceEnrollment struct {
	CreatedAt         int    `json:"createdAt"`
	ContentLanguage   string `json:"contentLanguage"`
	VoiceEnrollmentId int    `json:"voiceEnrollmentId"`
	Text              string `json:"text"`
	APICallId         string `json:"apiCallId"`
}

type GetAllVoiceEnrollmentsReturn struct {
	Message          string            `json:"message"`
	Count            int               `json:"count"`
	Status           int               `json:"status"`
	TimeTaken        string            `json:"timeTaken"`
	VoiceEnrollments []VoiceEnrollment `json:"voiceEnrollments"`
	ResponseCode     string            `json:"responseCode"`
	APICallId        string            `json:"apiCallId"`
}

type FaceEnrollment struct {
	CreatedAt        int    `json:"createdAt"`
	FaceEnrollmentId int    `json:"faceEnrollmentId"`
	APICallId        string `json:"apiCallId"`
}

type GetAllFaceEnrollmentsReturn struct {
	Message         string           `json:"message"`
	Count           int              `json:"count"`
	Status          int              `json:"status"`
	TimeTaken       string           `json:"timeTaken"`
	FaceEnrollments []FaceEnrollment `json:"faceEnrollments"`
	ResponseCode    string           `json:"responseCode"`
	APICallId       string           `json:"apiCallId"`
}

type VideoEnrollment struct {
	CreatedAt         int    `json:"createdAt"`
	ContentLanguage   string `json:"contentLanguage"`
	VideoEnrollmentId int    `json:"videoEnrollmentId"`
	Text              string `json:"text"`
	APICallId         string `json:"apiCallId"`
}

type GetAllVideoEnrollmentsReturn struct {
	Message          string            `json:"message"`
	Count            int               `json:"count"`
	Status           int               `json:"status"`
	TimeTaken        string            `json:"timeTaken"`
	VideoEnrollments []VideoEnrollment `json:"videoEnrollments"`
	ResponseCode     string            `json:"responseCode"`
	APICallId        string            `json:"apiCallId"`
}

type CreateVoiceEnrollmentReturn struct {
	Message         string  `json:"message"`
	ContentLanguage string  `json:"contentLanguage"`
	Id              int     `json:"id"`
	Status          int     `json:"status"`
	Text            string  `json:"text"`
	TextConfidence  float64 `json:"textConfidence"`
	CreatedAt       int     `json:"createdAt"`
	TimeTaken       string  `json:"timeTaken"`
	ResponseCode    string  `json:"responseCode"`
	APICallId       string  `json:"apiCallId"`
}

type CreateVoiceEnrollmentByUrlReturn struct {
	Message         string  `json:"message"`
	ContentLanguage string  `json:"contentLanguage"`
	Id              int     `json:"id"`
	Status          int     `json:"status"`
	Text            string  `json:"text"`
	TextConfidence  float64 `json:"textConfidence"`
	CreatedAt       int     `json:"createdAt"`
	TimeTaken       string  `json:"timeTaken"`
	ResponseCode    string  `json:"responseCode"`
	APICallId       string  `json:"apiCallId"`
}

type CreateFaceEnrollmentReturn struct {
	Message          string `json:"message"`
	Status           int    `json:"status"`
	TimeTaken        string `json:"timeTaken"`
	FaceEnrollmentId int    `json:"faceEnrollmentId"`
	CreatedAt        int    `json:"createdAt"`
	ResponseCode     string `json:"responseCode"`
	APICallId        string `json:"apiCallId"`
}

type CreateFaceEnrollmentByUrlReturn struct {
	Message          string `json:"message"`
	Status           int    `json:"status"`
	TimeTaken        string `json:"timeTaken"`
	FaceEnrollmentId int    `json:"faceEnrollmentId"`
	CreatedAt        int    `json:"createdAt"`
	ResponseCode     string `json:"responseCode"`
	APICallId        string `json:"apiCallId"`
}

type CreateVideoEnrollmentReturn struct {
	Message         string  `json:"message"`
	ContentLanguage string  `json:"contentLanguage"`
	Id              int     `json:"id"`
	Status          int     `json:"status"`
	Text            string  `json:"text"`
	TextConfidence  float64 `json:"textConfidence"`
	CreatedAt       int     `json:"createdAt"`
	TimeTaken       string  `json:"timeTaken"`
	ResponseCode    string  `json:"responseCode"`
	APICallId       string  `json:"apiCallId"`
}

type CreateVideoEnrollmentByUrlReturn struct {
	Message         string  `json:"message"`
	ContentLanguage string  `json:"contentLanguage"`
	Id              int     `json:"id"`
	Status          int     `json:"status"`
	Text            string  `json:"text"`
	TextConfidence  float64 `json:"textConfidence"`
	CreatedAt       int     `json:"createdAt"`
	TimeTaken       string  `json:"timeTaken"`
	ResponseCode    string  `json:"responseCode"`
	APICallId       string  `json:"apiCallId"`
}

type DeleteVoiceEnrollmentReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
	APICallId    string `json:"apiCallId"`
}

type DeleteFaceEnrollmentReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
	APICallId    string `json:"apiCallId"`
}

type DeleteVideoEnrollmentReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
	APICallId    string `json:"apiCallId"`
}

type DeleteAllVoiceEnrollmentsReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
	APICallId    string `json:"apiCallId"`
}

type DeleteAllFaceEnrollmentsReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
	APICallId    string `json:"apiCallId"`
}

type DeleteAllVideoEnrollmentsReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
	APICallId    string `json:"apiCallId"`
}

type DeleteAllEnrollmentsReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
	APICallId    string `json:"apiCallId"`
}
