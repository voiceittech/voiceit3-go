package voiceit2

import (
	"bytes"
	"io"
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

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("GET", vi.BaseUrl+"/users", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

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

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/users", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// GetUser takes the userId generated during a createUser and returns
// a user exists for the given userId
// For more details see https://api.voiceit.io/#check-if-a-specific-user-exists
func (vi *VoiceIt2) CheckUserExists(userId string) string {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("GET", vi.BaseUrl+"/users/"+userId, body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

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

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("DELETE", vi.BaseUrl+"/users/"+userId, body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

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

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("GET", vi.BaseUrl+"/users/"+userId+"/groups", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// GetAllGroups returns a list of all groups associated with the API Key
// For more details see https://api.voiceit.io/#get-all-groups
func (vi *VoiceIt2) GetAllGroups() string {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("GET", vi.BaseUrl+"/groups", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

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

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("GET", vi.BaseUrl+"/groups/"+groupId, body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

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

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("GET", vi.BaseUrl+"/groups/"+groupId+"/exists", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

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

	_ = writer.WriteField("description", description)

	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/groups", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
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

	_ = writer.WriteField("groupId", groupId)
	_ = writer.WriteField("userId", userId)

	writer.Close()

	req, _ := http.NewRequest("PUT", vi.BaseUrl+"/groups/addUser", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
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

	_ = writer.WriteField("groupId", groupId)
	_ = writer.WriteField("userId", userId)

	writer.Close()

	req, _ := http.NewRequest("PUT", vi.BaseUrl+"/groups/removeUser", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// DeleteUser takes the userId generated during a createUser and deletes
// the user profile and all associated face and voice enrollments
// For more details see https://api.voiceit.io/#delete-a-specific-group
func (vi *VoiceIt2) DeleteGroup(groupId string) string {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("DELETE", vi.BaseUrl+"/groups/"+groupId, body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// GetAllEnrollmentsForUser takes the userId generated during a createUser
// and returns a list of all video/voice enrollments for the user
// For more details see https://api.voiceit.io/#get-all-enrollments-for-user
func (vi *VoiceIt2) GetAllEnrollmentsForUser(userId string) string {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("GET", vi.BaseUrl+"/enrollments/"+userId, body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// GetFaceEnrollmentsForUser takes the userId generated during a createUser
// and returns a list of all face enrollments for the user
// For more details see https://api.voiceit.io/#get-user-39-s-face-enrollments
func (vi *VoiceIt2) GetFaceEnrollmentsForUser(userId string) string {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("GET", vi.BaseUrl+"/enrollments/face/"+userId, body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// CreateVoiceEnrollment takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// and absolute file path for a audio recording to create a voice enrollment for the user
// For more details see https://api.voiceit.io/#create-voice-enrollment
func (vi *VoiceIt2) CreateVoiceEnrollment(userId string, contentLanguage string, filePath string) string {

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("recording", filepath.Base(file.Name()))
	io.Copy(part, file)
	_ = writer.WriteField("userId", userId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/enrollments", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// CreateVoiceEnrollmentByUrl takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// and a fully qualified URL to a audio recording to create a voice enrollment for the user
// For more details see https://api.voiceit.io/#create-voice-enrollment-by-url
func (vi *VoiceIt2) CreateVoiceEnrollmentByUrl(userId string, contentLanguage string, fileUrl string) string {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("userId", userId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	_ = writer.WriteField("fileUrl", fileUrl)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/enrollments/byUrl", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
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
func (vi *VoiceIt2) CreateFaceEnrollment(userId string, contentLanguage string, filePath string, doBlinkDetection ...bool) string {
	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("video", filepath.Base(file.Name()))
	io.Copy(part, file)
	_ = writer.WriteField("userId", userId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	if len(doBlinkDetection) > 0 {
		_ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
	}
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/enrollments/face", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// CreateVideoEnrollment takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// and absolute file path for a video recording to create a video enrollment for the user
// For more details see https://api.voiceit.io/#create-video-enrollment
func (vi *VoiceIt2) CreateVideoEnrollment(userId string, contentLanguage string, filePath string, doBlinkDetection ...bool) string {

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("video", filepath.Base(file.Name()))
	io.Copy(part, file)
	_ = writer.WriteField("userId", userId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	if len(doBlinkDetection) > 0 {
		_ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
	}
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/enrollments/video", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// CreateVideoEnrollment takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// and a fully qualified URL to a video recording to create a video enrollment for the user
// For more details see https://api.voiceit.io/#create-video-enrollment-by-url
func (vi *VoiceIt2) CreateVideoEnrollmentByUrl(userId string, contentLanguage string, fileUrl string, doBlinkDetection ...bool) string {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("userId", userId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	_ = writer.WriteField("fileUrl", fileUrl)
	if len(doBlinkDetection) > 0 {
		_ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
	}
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/enrollments/video/byUrl", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
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
func (vi *VoiceIt2) DeleteFaceEnrollment(userId string, faceEnrollmentId string) string {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/face/"+userId+"/"+faceEnrollmentId, body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// DeleteEnrollmentForUser takes the userId generated during a createUser and
// an enrollmentId returned during a voiceEnrollment/videoEnrollment and deletes
// the voice/video enrollment for the user
// For more details see https://api.voiceit.io/#delete-enrollment-for-user
func (vi *VoiceIt2) DeleteEnrollmentForUser(userId string, faceEnrollmentId string) string {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/"+userId+"/"+faceEnrollmentId, body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// DeleteAllEnrollmentsForUser takes the userId generated during a createUser
// and deletes all video/voice enrollments for the user
func (vi *VoiceIt2) DeleteAllEnrollmentsForUser(userId string) string {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/"+userId+"/all", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// VoiceVerification takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// and absolute file path for a audio recording to verify the user's voice
// For more details see https://api.voiceit.io/#verify-a-user-39-s-voice
func (vi *VoiceIt2) VoiceVerification(userId string, contentLanguage string, filePath string) string {

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("recording", filepath.Base(file.Name()))
	io.Copy(part, file)
	_ = writer.WriteField("userId", userId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/verification", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// VoiceVerification takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// and a fully qualified URL to a audio recording to verify the user's voice
// For more details see https://api.voiceit.io/#verify-a-user-39-s-voice
func (vi *VoiceIt2) VoiceVerificationByUrl(userId string, contentLanguage string, fileUrl string) string {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("userId", userId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	_ = writer.WriteField("fileUrl", fileUrl)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/verification/byUrl", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
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
func (vi *VoiceIt2) FaceVerification(userId string, filePath string, doBlinkDetection ...bool) string {

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("video", filepath.Base(file.Name()))
	io.Copy(part, file)
	_ = writer.WriteField("userId", userId)
	if len(doBlinkDetection) > 0 {
		_ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
	}
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/verification/face", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// VideoVerification takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// and absolute file path for a video recording to verify the user's face and voice
// For more details see https://api.voiceit.io/#video-verification
func (vi *VoiceIt2) VideoVerification(userId string, contentLanguage string, filePath string, doBlinkDetection ...bool) string {

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("video", filepath.Base(file.Name()))
	io.Copy(part, file)
	_ = writer.WriteField("userId", userId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	if len(doBlinkDetection) > 0 {
		_ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
	}
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/verification/video", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// VideoVerification takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// and a fully qualified URL to a video recording to verify the user's face and voice
// For more details see https://api.voiceit.io/#video-verification-by-url
func (vi *VoiceIt2) VideoVerificationByUrl(userId string, contentLanguage string, fileUrl string, doBlinkDetection ...bool) string {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("userId", userId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	_ = writer.WriteField("fileUrl", fileUrl)
	if len(doBlinkDetection) > 0 {
		_ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
	}
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/verification/video/byUrl", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// VoiceIdentification takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// and absolute file path for a audio recording to idetify the user's voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-39-s-voice
func (vi *VoiceIt2) VoiceIdentification(groupId string, contentLanguage string, filePath string) string {

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("recording", filepath.Base(file.Name()))
	io.Copy(part, file)
	_ = writer.WriteField("groupId", groupId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/identification", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// VoiceIdentification takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// and a fully qualified URL to a audio recording to idetify the user's voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-39-s-voice-by-url
func (vi *VoiceIt2) VoiceIdentificationByUrl(groupId string, contentLanguage string, fileUrl string) string {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("fileUrl", fileUrl)
	_ = writer.WriteField("groupId", groupId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/identification/byUrl", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// VideoIdentification takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// and absolute file path for a video recording to idetify the user's face and voice
// amongst others in the group
// For more details see https://api.voiceit.io/#video-identification
func (vi *VoiceIt2) VideoIdentification(groupId string, contentLanguage string, filePath string, doBlinkDetection ...bool) string {

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("video", filepath.Base(file.Name()))
	io.Copy(part, file)
	_ = writer.WriteField("groupId", groupId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	if len(doBlinkDetection) > 0 {
		_ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
	}
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/identification/video", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}

// VideoIdentification takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// and a fully qualified URL to a video recording to idetify the user's face and voice
// amongst others in the group
// For more details see https://api.voiceit.io/#video-identification
func (vi *VoiceIt2) VideoIdentificationByUrl(groupId string, contentLanguage string, fileUrl string, doBlinkDetection ...bool) string {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("fileUrl", fileUrl)
	_ = writer.WriteField("groupId", groupId)
	_ = writer.WriteField("contentLanguage", contentLanguage)
	if len(doBlinkDetection) > 0 {
		_ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
	}
	writer.Close()

	req, _ := http.NewRequest("POST", vi.BaseUrl+"/identification/video/byUrl", body)
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	reply, _ := ioutil.ReadAll(resp.Body)
	return string(reply)
}
