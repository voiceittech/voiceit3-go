package voiceit2

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type VoiceIt2 struct {
	ApiKey   string
	ApiToken string
	BaseUrl  string
}

func NewClient(key string, tok string) *VoiceIt2 {
	return &VoiceIt2{
		ApiKey:   key,
		ApiToken: tok,
		BaseUrl:  "https://api.voiceit.io",
	}
}

// GetAllUsers returns a list of all users associated with the API Key
// For more details see https://api.voiceit.io/#get-all-users
func (vi *VoiceIt2) GetAllUsers() string {
	req, _ := http.NewRequest("GET", vi.BaseUrl+"/users", nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// CreateUser creates a new user profile and returns a unique userId
// that is used for all future calls related to the user profile
// For more details see https://api.voiceit.io/#create-a-user
func (vi *VoiceIt2) CreateUser() string {
	req, _ := http.NewRequest("POST", vi.BaseUrl+"/users", nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// CheckUserExists takes the userId generated during a createUser and returns
// a JSON object which contains the boolean "exists" which shows whether a given user exists
// For more details see https://api.voiceit.io/#check-if-a-specific-user-exists
func (vi *VoiceIt2) CheckUserExists(userId string) string {
	req, _ := http.NewRequest("GET", vi.BaseUrl+"/users/"+userId, nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// DeleteUser takes the userId generated during a createUser and deletes
// the user profile and all associated face and voice enrollments
// For more details see https://api.voiceit.io/#delete-a-specific-user
func (vi *VoiceIt2) DeleteUser(userId string) string {
	req, _ := http.NewRequest("DELETE", vi.BaseUrl+"/users/"+userId, nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// GetGroupsForUser takes the userId generated during a createUser and returns
// a list of all groups that the user belongs to
// For more details see https://api.voiceit.io/#get-groups-for-user
func (vi *VoiceIt2) GetGroupsForUser(userId string) string {
	req, _ := http.NewRequest("GET", vi.BaseUrl+"/users/"+userId+"/groups", nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// GetAllGroups returns a list of all groups associated with the API Key
// For more details see https://api.voiceit.io/#get-all-groups
func (vi *VoiceIt2) GetAllGroups() string {
	req, _ := http.NewRequest("GET", vi.BaseUrl+"/groups", nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// GetGroup takes the groupId generated during a createGroup
// and returns the group along with a list of associated users in the group
// For more details see https://api.voiceit.io/#get-a-specific-group
func (vi *VoiceIt2) GetGroup(groupId string) string {
	req, _ := http.NewRequest("GET", vi.BaseUrl+"/groups/"+groupId, nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// CheckGroupExists takes the groupId generated during a createGroup
// and returns whether the group exists for the given groupId
// For more details see https://api.voiceit.io/#check-if-group-exists
func (vi *VoiceIt2) CheckGroupExists(groupId string) string {
	req, _ := http.NewRequest("GET", vi.BaseUrl+"/groups/"+groupId+"/exists", nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// CreateGroup creates a new group profile and returns a unique groupId
// that is used for all future calls related to the group
// For more details see https://api.voiceit.io/#create-a-group
func (vi *VoiceIt2) CreateGroup(description string) string {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("description", description)

	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/groups", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// AddUserToGroup takes the groupId generated during a createGroup
// and the userId generated during createUser and adds the user to the group
// For more details see https://api.voiceit.io/#add-user-to-group
func (vi *VoiceIt2) AddUserToGroup(groupId string, userId string) string {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("groupId", groupId)
	writer.WriteField("userId", userId)

	writer.Close()

	req, _ := http.NewRequest("PUT", vi.BaseUrl+"/groups/addUser", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// RemoveUserFromGroup takes the groupId generated during a createGroup
// and the userId generated during createUser and removes the user from the group
// For more details see https://api.voiceit.io/#remove-user-from-group
func (vi *VoiceIt2) RemoveUserFromGroup(groupId string, userId string) string {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("groupId", groupId)
	writer.WriteField("userId", userId)

	writer.Close()

	req, _ := http.NewRequest("PUT", vi.BaseUrl+"/groups/removeUser", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// DeleteGroup takes the groupId generated during a createGroup and deletes
// the group profile disassociates all users associated with it
// For more details see https://api.voiceit.io/#delete-a-specific-group
func (vi *VoiceIt2) DeleteGroup(groupId string) string {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("DELETE", vi.BaseUrl+"/groups/"+groupId, body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// GetAllVoiceEnrollments takes the userId generated during a createUser
// and returns a list of all voice enrollments for the user
// For more details see https://api.voiceit.io/#get-voice-enrollments
func (vi *VoiceIt2) GetAllVoiceEnrollments(userId string) string {
	req, _ := http.NewRequest("GET", vi.BaseUrl+"/enrollments/voice/"+userId, nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// GetAllVideoEnrollments takes the userId generated during a createUser
// and returns a list of all video enrollments for the user
// For more details see https://api.voiceit.io/#get-video-enrollments
func (vi *VoiceIt2) GetAllVideoEnrollments(userId string) string {
	req, _ := http.NewRequest("GET", vi.BaseUrl+"/enrollments/video/"+userId, nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// GetAllFaceEnrollments takes the userId generated during a createUser
// and returns a list of all face enrollments for the user
// For more details see https://api.voiceit.io/#get-face-enrollments
func (vi *VoiceIt2) GetAllFaceEnrollments(userId string) string {
	req, _ := http.NewRequest("GET", vi.BaseUrl+"/enrollments/face/"+userId, nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// CreateVoiceEnrollment takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a audio recording to create a voice enrollment for the user
// For more details see https://api.voiceit.io/#create-voice-enrollment
func (vi *VoiceIt2) CreateVoiceEnrollment(userId string, contentLanguage string, phrase string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	fileContents, _ := ioutil.ReadAll(file)
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("recording", filepath.Base(file.Name()))
	part.Write(fileContents)
	writer.WriteField("userId", userId)
	writer.WriteField("contentLanguage", contentLanguage)
	writer.WriteField("phrase", phrase)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/enrollments/voice", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply), nil
}

// CreateVoiceEnrollmentByUrl takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a audio recording to create a voice enrollment for the user
// For more details see https://api.voiceit.io/#create-voice-enrollment-by-url
func (vi *VoiceIt2) CreateVoiceEnrollmentByUrl(userId string, contentLanguage string, phrase string, fileUrl string) string {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("userId", userId)
	writer.WriteField("contentLanguage", contentLanguage)
	writer.WriteField("fileUrl", fileUrl)
	writer.WriteField("phrase", phrase)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/enrollments/voice/byUrl", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// CreateFaceEnrollment takes the userId generated during a createUser and
// absolute file path for a video recording to create a face enrollment for the user
// For more details see https://api.voiceit.io/#create-face-enrollment
func (vi *VoiceIt2) CreateFaceEnrollment(userId string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	fileContents, _ := ioutil.ReadAll(file)
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("video", filepath.Base(file.Name()))
	part.Write(fileContents)
	writer.WriteField("userId", userId)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/enrollments/face", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply), nil
}

// CreateFaceEnrollmentByUrl takes the userId generated during a createUser
// and a fully qualified URL to a video recording to verify the user's face
// For more details see https://api.voiceit.io/#create-face-enrollment-by-url
func (vi *VoiceIt2) CreateFaceEnrollmentByUrl(userId string, fileUrl string) string {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("userId", userId)
	writer.WriteField("fileUrl", fileUrl)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/enrollments/face/byUrl", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// CreateVideoEnrollment takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a video recording to create a video enrollment for the user
// For more details see https://api.voiceit.io/#create-video-enrollment
func (vi *VoiceIt2) CreateVideoEnrollment(userId string, contentLanguage string, phrase string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	fileContents, _ := ioutil.ReadAll(file)
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("video", filepath.Base(file.Name()))
	part.Write(fileContents)
	writer.WriteField("userId", userId)
	writer.WriteField("contentLanguage", contentLanguage)
	writer.WriteField("phrase", phrase)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/enrollments/video", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply), nil
}

// CreateVideoEnrollment takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a video recording to create a video enrollment for the user
// For more details see https://api.voiceit.io/#create-video-enrollment-by-url
func (vi *VoiceIt2) CreateVideoEnrollmentByUrl(userId string, contentLanguage string, phrase string, fileUrl string) string {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("userId", userId)
	writer.WriteField("contentLanguage", contentLanguage)
	writer.WriteField("fileUrl", fileUrl)
	writer.WriteField("phrase", phrase)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/enrollments/video/byUrl", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// DeleteFaceEnrollment takes the userId generated during a createUser and
// a faceEnrollmentId returned during a faceEnrollment and deletes the specific
// faceEnrollment for the user
// For more details see https://api.voiceit.io/#delete-face-enrollment
func (vi *VoiceIt2) DeleteFaceEnrollment(userId string, faceEnrollmentId int) string {
	req, _ := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/face/"+userId+"/"+strconv.Itoa(faceEnrollmentId), nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// DeleteAllFaceEnrollments takes the userId generated during a createUser
// For more details see https://api.voiceit.io/#delete-all-face-enrollments
func (vi *VoiceIt2) DeleteAllFaceEnrollments(userId string) string {
	req, _ := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/"+userId+"/face", nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// DeleteVoiceEnrollment takes the userId generated during a createUser and
// an enrollmentId returned during a voiceEnrollment/videoEnrollment and deletes
// the voice enrollment for the user
// For more details see https://api.voiceit.io/#delete-voice-enrollment
func (vi *VoiceIt2) DeleteVoiceEnrollment(userId string, id int) string {
	req, _ := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/voice/"+userId+"/"+strconv.Itoa(id), nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// DeleteAllVoiceEnrollments takes the userId generated during a createUser
// For more details https://api.voiceit.io/#delete-all-voice-enrollments
func (vi *VoiceIt2) DeleteAllVoiceEnrollments(userId string) string {
	req, _ := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/"+userId+"/voice", nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// DeleteVideoEnrollment takes the userId generated during a createUser and
// an enrollmentId returned during a voiceEnrollment/videoEnrollment and deletes
// the voice/video enrollment for the user
// For more details see https://api.voiceit.io/#delete-video-enrollment
func (vi *VoiceIt2) DeleteVideoEnrollment(userId string, id int) string {
	req, _ := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/video/"+userId+"/"+strconv.Itoa(id), nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// DeleteAllVideoEnrollments takes the userId generated during a createUser
// For more details see https://api.voiceit.io/#delete-all-video-enrollments
func (vi *VoiceIt2) DeleteAllVideoEnrollments(userId string) string {
	req, _ := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/"+userId+"/video", nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// DeleteAllEnrollments takes the userId generated during a createUser
// and deletes all video/voice enrollments for the user
// For more details see https://api.voiceit.io/#delete-all-enrollments-for-user
func (vi *VoiceIt2) DeleteAllEnrollments(userId string) string {
	req, _ := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/"+userId+"/all", nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// VoiceVerification takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a audio recording to verify the user's voice
// For more details see https://api.voiceit.io/#verify-a-user-39-s-voice
func (vi *VoiceIt2) VoiceVerification(userId string, contentLanguage string, phrase string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	fileContents, _ := ioutil.ReadAll(file)
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("recording", filepath.Base(file.Name()))
	part.Write(fileContents)
	writer.WriteField("userId", userId)
	writer.WriteField("contentLanguage", contentLanguage)
	writer.WriteField("phrase", phrase)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/verification/voice", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply), nil
}

// VoiceVerificationByUrl takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a audio recording to verify the user's voice
// For more details see https://api.voiceit.io/#verify-a-user-39-s-voice-by-url
func (vi *VoiceIt2) VoiceVerificationByUrl(userId string, contentLanguage string, phrase string, fileUrl string) string {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("userId", userId)
	writer.WriteField("contentLanguage", contentLanguage)
	writer.WriteField("fileUrl", fileUrl)
	writer.WriteField("phrase", phrase)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/verification/voice/byUrl", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// FaceVerification takes the userId generated during a createUser and a
// absolute file path for a video recording to verify the user's face
// For more details see https://api.voiceit.io/#verify-a-user-39-s-face
func (vi *VoiceIt2) FaceVerification(userId string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	fileContents, _ := ioutil.ReadAll(file)
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("video", filepath.Base(file.Name()))
	part.Write(fileContents)
	writer.WriteField("userId", userId)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/verification/face", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply), nil
}

// FaceVerificationByUrl takes the userId generated during a createUser
// and a fully qualified URL to a video recording to verify the user's face
// For more details see https://api.voiceit.io/#verify-a-user-39-s-face-by-url
func (vi *VoiceIt2) FaceVerificationByUrl(userId string, fileUrl string) string {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("fileUrl", fileUrl)
	writer.WriteField("userId", userId)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/verification/face/byUrl", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// VideoVerification takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a video recording to verify the user's face and voice
// For more details see https://api.voiceit.io/#video-verification
func (vi *VoiceIt2) VideoVerification(userId string, contentLanguage string, phrase string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	fileContents, _ := ioutil.ReadAll(file)
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("video", filepath.Base(file.Name()))
	part.Write(fileContents)
	writer.WriteField("userId", userId)
	writer.WriteField("contentLanguage", contentLanguage)
	writer.WriteField("phrase", phrase)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/verification/video", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply), nil
}

// VideoVerificationByUrl takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a video recording to verify the user's face and voice
// For more details see https://api.voiceit.io/#video-verification-by-url
func (vi *VoiceIt2) VideoVerificationByUrl(userId string, contentLanguage string, phrase string, fileUrl string) string {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("userId", userId)
	writer.WriteField("contentLanguage", contentLanguage)
	writer.WriteField("fileUrl", fileUrl)
	writer.WriteField("phrase", phrase)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/verification/video/byUrl", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// VoiceIdentification takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a audio recording to idetify the user's voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-39-s-voice
func (vi *VoiceIt2) VoiceIdentification(groupId string, contentLanguage string, phrase string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	fileContents, _ := ioutil.ReadAll(file)
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("recording", filepath.Base(file.Name()))
	part.Write(fileContents)
	writer.WriteField("groupId", groupId)
	writer.WriteField("contentLanguage", contentLanguage)
	writer.WriteField("phrase", phrase)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/identification/voice", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply), nil
}

// VoiceIdentificationByUrl takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a audio recording to idetify the user's voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-39-s-voice-by-url
func (vi *VoiceIt2) VoiceIdentificationByUrl(groupId string, contentLanguage string, phrase string, fileUrl string) string {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("fileUrl", fileUrl)
	writer.WriteField("groupId", groupId)
	writer.WriteField("contentLanguage", contentLanguage)
	writer.WriteField("phrase", phrase)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/identification/voice/byUrl", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// VideoIdentification takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a video recording to idetify the user's face and voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-39-s-voice-amp-face
func (vi *VoiceIt2) VideoIdentification(groupId string, contentLanguage string, phrase string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	fileContents, _ := ioutil.ReadAll(file)
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("video", filepath.Base(file.Name()))
	part.Write(fileContents)
	writer.WriteField("groupId", groupId)
	writer.WriteField("contentLanguage", contentLanguage)
	writer.WriteField("phrase", phrase)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/identification/video", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply), nil
}

// VideoIdentificationByUrl takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a video recording to idetify the user's face and voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-39-s-voice-amp-face-by-url
func (vi *VoiceIt2) VideoIdentificationByUrl(groupId string, contentLanguage string, phrase string, fileUrl string) string {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("fileUrl", fileUrl)
	writer.WriteField("groupId", groupId)
	writer.WriteField("contentLanguage", contentLanguage)
	writer.WriteField("phrase", phrase)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/identification/video/byUrl", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// FaceIdentification takes the groupId generated during a createGroup,
// and absolute file path for a face recording to idetify the user's face
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-39-s-face
func (vi *VoiceIt2) FaceIdentification(groupId string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	fileContents, _ := ioutil.ReadAll(file)
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("video", filepath.Base(file.Name()))
	part.Write(fileContents)
	writer.WriteField("groupId", groupId)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/identification/face", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply), nil
}

// FaceIdentificationByUrl takes the groupId generated during a createGroup,
// and a fully qualified URL to a face recording to idetify the user's face
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-39-s-face-by-url
func (vi *VoiceIt2) FaceIdentificationByUrl(groupId string, fileUrl string) string {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("fileUrl", fileUrl)
	writer.WriteField("groupId", groupId)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/identification/face/byUrl", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// GetPhrases takes the contentLanguage
// For more details see https://api.voiceit.io/#get-phrases
func (vi *VoiceIt2) GetPhrases(contentLanguage string) string {
	req, _ := http.NewRequest("GET", vi.BaseUrl+"/phrases/"+contentLanguage, nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// CreateUserToken takes the userId
// For more details see https://api.voiceit.io/?go#user-token-generation
func (vi *VoiceIt2) CreateUserToken(userId string) string {
	req, _ := http.NewRequest("POST", vi.BaseUrl+"/users/"+userId+"/token", nil)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}
