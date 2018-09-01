package structs

type VoiceEnrollment struct {
	CreatedAt         int    `json:"createdAt"`
	ContentLanguage   string `json:"contentLanguage"`
	VoiceEnrollmentId int    `json:"voiceEnrollmentId"`
	Text              string `json:"text"`
}

type GetAllVoiceEnrollmentsReturn struct {
	Message          string            `json:"message"`
	Count            int               `json:"count"`
	Status           int               `json:"status"`
	TimeTaken        string            `json:"timeTaken"`
	VoiceEnrollments []VoiceEnrollment `json:"voiceEnrollments"`
	ResponseCode     string            `json:"responseCode"`
}

type FaceEnrollment struct {
	CreatedAt        int `json:"createdAt"`
	FaceEnrollmentId int `json:"faceEnrollmentId"`
}

type GetAllFaceEnrollmentsReturn struct {
	Message         string           `json:"message"`
	Count           int              `json:"count"`
	Status          int              `json:"status"`
	TimeTaken       string           `json:"timeTaken"`
	FaceEnrollments []FaceEnrollment `json:"faceEnrollments"`
	ResponseCode    string           `json:"responseCode"`
}

type VideoEnrollment struct {
	CreatedAt         int    `json:"createdAt"`
	ContentLanguage   string `json:"contentLanguage"`
	VideoEnrollmentId int    `json:"videoEnrollmentId"`
	Text              string `json:"text"`
}

type GetAllVideoEnrollmentsReturn struct {
	Message          string            `json:"message"`
	Count            int               `json:"count"`
	Status           int               `json:"status"`
	TimeTaken        string            `json:"timeTaken"`
	VideoEnrollments []VideoEnrollment `json:"videoEnrollments"`
	ResponseCode     string            `json:"responseCode"`
}

type CreateVoiceEnrollmentReturn struct {
	Message         string  `json:"message"`
	ContentLanguage string  `json:"contentLanguage"`
	Id              int     `json:"id"`
	Status          int     `json:"status"`
	Text            string  `json:"text"`
	TextConfidence  float32 `json:"textConfidence"`
	CreatedAt       int     `json:"createdAt"`
	TimeTaken       string  `json:"timeTaken"`
	ResponseCode    string  `json:"responseCode"`
}

type CreateVoiceEnrollmentByUrlReturn struct {
	Message         string  `json:"message"`
	ContentLanguage string  `json:"contentLanguage"`
	Id              int     `json:"id"`
	Status          int     `json:"status"`
	Text            string  `json:"text"`
	TextConfidence  float32 `json:"textConfidence"`
	CreatedAt       int     `json:"createdAt"`
	TimeTaken       string  `json:"timeTaken"`
	ResponseCode    string  `json:"responseCode"`
}

type CreateFaceEnrollmentReturn struct {
	Message          string `json:"message"`
	Status           int    `json:"status"`
	BlinksCount      int    `json:"blinksCount"`
	TimeTaken        string `json:"timeTaken"`
	FaceEnrollmentId int    `json:"faceEnrollmentId"`
	CreatedAt        int    `json:"createdAt"`
	ResponseCode     string `json:"responseCode"`
}

type CreateFaceEnrollmentByUrlReturn struct {
	Message          string `json:"message"`
	Status           int    `json:"status"`
	BlinksCount      int    `json:"blinksCount"`
	TimeTaken        string `json:"timeTaken"`
	FaceEnrollmentId int    `json:"faceEnrollmentId"`
	CreatedAt        int    `json:"createdAt"`
	ResponseCode     string `json:"responseCode"`
}

type CreateVideoEnrollmentReturn struct {
	Message         string  `json:"message"`
	ContentLanguage string  `json:"contentLanguage"`
	Id              int     `json:"id"`
	Status          int     `json:"status"`
	Text            string  `json:"text"`
	TextConfidence  float32 `json:"textConfidence"`
	CreatedAt       int     `json:"createdAt"`
	TimeTaken       string  `json:"timeTaken"`
	ResponseCode    string  `json:"responseCode"`
}

type CreateVideoEnrollmentByUrlReturn struct {
	Message         string  `json:"message"`
	ContentLanguage string  `json:"contentLanguage"`
	Id              int     `json:"id"`
	Status          int     `json:"status"`
	Text            string  `json:"text"`
	TextConfidence  float32 `json:"textConfidence"`
	CreatedAt       int     `json:"createdAt"`
	TimeTaken       string  `json:"timeTaken"`
	ResponseCode    string  `json:"responseCode"`
}

type DeleteVoiceEnrollmentReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
}

type DeleteFaceEnrollmentReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
}

type DeleteVideoEnrollmentReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
}

type DeleteAllVoiceEnrollmentsReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
}

type DeleteAllFaceEnrollmentsReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
}

type DeleteAllVideoEnrollmentsReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
}

type DeleteAllEnrollmentsReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
}
