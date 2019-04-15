package voiceit2

import (
	"bytes"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const PlatformVersion string = "v2.0.0"

type VoiceIt2 struct {
	ApiKey          string
	ApiToken        string
	BaseUrl         string
	NotificationUrl string
}

// NewClient returns a new VoiceIt2 client
func NewClient(key string, tok string) *VoiceIt2 {
	return &VoiceIt2{
		ApiKey:          key,
		ApiToken:        tok,
		BaseUrl:         "https://api.voiceit.io",
		NotificationUrl: "",
	}
}

// AddNotificationUrl adds a notification URL field in the VoiceIt2 object.
// If one is already specified, it will be overwritten
// For more details, see https://api.voiceit.io/#webhook-notification
func (vi *VoiceIt2) AddNotificationUrl(notificationUrl string) {
	vi.NotificationUrl = "?notificationURL=" + url.QueryEscape(notificationUrl)
}

// RemoveNotificationUrl removes the notification URL field from the VoiceIt2 struct
func (vi *VoiceIt2) RemoveNotificationUrl() {
	vi.NotificationUrl = ""
}

// GetAllUsers returns a list of all users associated with the API Key
// For more details see https://api.voiceit.io/#get-all-users
func (vi VoiceIt2) GetAllUsers() (string, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/users"+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in GetAllUsers():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in GetAllUsers():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in GetAllUsers():", err)
		return "", err
	}
	return string(reply), nil
}

// CreateUser creates a new user profile and returns a unique userId
// that is used for all future calls related to the user profile
// For more details see https://api.voiceit.io/#create-a-user
func (vi VoiceIt2) CreateUser() (string, error) {
	req, err := http.NewRequest("POST", vi.BaseUrl+"/users"+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in CreateUser():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in CreateUser():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in CreateUser():", err)
		return "", err
	}
	return string(reply), nil
}

// CheckUserExists takes the userId generated during a createUser and returns
// a JSON object which contains the boolean "exists" which shows whether a given user exists
// For more details see https://api.voiceit.io/#check-if-a-specific-user-exists
func (vi VoiceIt2) CheckUserExists(userId string) (string, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/users/"+userId+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in CheckUserExists():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in CheckUserExists():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in CheckUserExists():", err)
		return "", err
	}
	return string(reply), nil
}

// DeleteUser takes the userId generated during a createUser and deletes
// the user profile and all associated face and voice enrollments
// For more details see https://api.voiceit.io/#delete-a-specific-user
func (vi VoiceIt2) DeleteUser(userId string) (string, error) {
	req, err := http.NewRequest("DELETE", vi.BaseUrl+"/users/"+userId+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in DeleteUser():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in DeleteUser():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in DeleteUser():", err)
		return "", err
	}
	return string(reply), nil
}

// GetGroupsForUser takes the userId generated during a createUser and returns
// a list of all groups that the user belongs to
// For more details see https://api.voiceit.io/#get-groups-for-user
func (vi VoiceIt2) GetGroupsForUser(userId string) (string, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/users/"+userId+"/groups"+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in GetGroupsForUser():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in GetGroupsForUser():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in GetGroupsForUser():", err)
		return "", err
	}
	return string(reply), nil
}

// GetAllGroups returns a list of all groups associated with the API Key
// For more details see https://api.voiceit.io/#get-all-groups
func (vi VoiceIt2) GetAllGroups() (string, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/groups"+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in GetAllGroups():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in GetAllGroups():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in GetAllGroups():", err)
		return "", err
	}
	return string(reply), nil
}

// GetGroup takes the groupId generated during a createGroup
// and returns the group along with a list of associated users in the group
// For more details see https://api.voiceit.io/#get-a-specific-group
func (vi VoiceIt2) GetGroup(groupId string) (string, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/groups/"+groupId+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in GetGroup():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in GetGroup():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in GetGroup():", err)
		return "", err
	}
	return string(reply), nil
}

// CheckGroupExists takes the groupId generated during a createGroup
// and returns whether the group exists for the given groupId
// For more details see https://api.voiceit.io/#check-if-group-exists
func (vi VoiceIt2) CheckGroupExists(groupId string) (string, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/groups/"+groupId+"/exists"+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in CheckGroupExists():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in CheckGroupExists():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in CheckGroupExists():", err)
		return "", err
	}
	return string(reply), nil
}

// CreateGroup creates a new group profile and returns a unique groupId
// that is used for all future calls related to the group
// For more details see https://api.voiceit.io/#create-a-group
func (vi VoiceIt2) CreateGroup(description string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("description", description)
	if err != nil {
		log.Println("Error in CreateGroup():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/groups"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in CreateGroup():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in CreateGroup():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in CreateGroup():", err)
		return "", err
	}
	return string(reply), nil
}

// AddUserToGroup takes the groupId generated during a createGroup
// and the userId generated during createUser and adds the user to the group
// For more details see https://api.voiceit.io/#add-user-to-group
func (vi VoiceIt2) AddUserToGroup(groupId string, userId string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("groupId", groupId)
	if err != nil {
		log.Println("Error in AddUserToGroup():", err)
		return "", err
	}

	err = writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in AddUserToGroup():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("PUT", vi.BaseUrl+"/groups/addUser"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in AddUserToGroup():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in AddUserToGroup():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in AddUserToGroup():", err)
		return "", err
	}
	return string(reply), nil
}

// RemoveUserFromGroup takes the groupId generated during a createGroup
// and the userId generated during createUser and removes the user from the group
// For more details see https://api.voiceit.io/#remove-user-from-group
func (vi VoiceIt2) RemoveUserFromGroup(groupId string, userId string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("groupId", groupId)
	if err != nil {
		log.Println("Error in RemoveUserFromGroup():", err)
		return "", err
	}

	err = writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in RemoveUserFromGroup():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("PUT", vi.BaseUrl+"/groups/removeUser"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in RemoveUserFromGroup():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in RemoveUserFromGroup():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in RemoveUserFromGroup():", err)
		return "", err
	}
	return string(reply), nil
}

// DeleteGroup takes the groupId generated during a createGroup and deletes
// the group profile disassociates all users associated with it
// For more details see https://api.voiceit.io/#delete-a-specific-group
func (vi VoiceIt2) DeleteGroup(groupId string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, err := http.NewRequest("DELETE", vi.BaseUrl+"/groups/"+groupId+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in DeleteGroup():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in DeleteGroup():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in DeleteGroup():", err)
		return "", err
	}
	return string(reply), nil
}

// GetAllVoiceEnrollments takes the userId generated during a createUser
// and returns a list of all voice enrollments for the user
// For more details see https://api.voiceit.io/#get-voice-enrollments
func (vi VoiceIt2) GetAllVoiceEnrollments(userId string) (string, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/enrollments/voice/"+userId+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in GetAllVoiceEnrollments():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in GetAllVoiceEnrollments():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in GetAllVoiceEnrollments():", err)
		return "", err
	}
	return string(reply), nil
}

// GetAllVideoEnrollments takes the userId generated during a createUser
// and returns a list of all video enrollments for the user
// For more details see https://api.voiceit.io/#get-video-enrollments
func (vi VoiceIt2) GetAllVideoEnrollments(userId string) (string, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/enrollments/video/"+userId+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in GetAllVideoEnrollments():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in GetAllVideoEnrollments():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in GetAllVideoEnrollments():", err)
		return "", err
	}
	return string(reply), nil
}

// GetAllFaceEnrollments takes the userId generated during a createUser
// and returns a list of all face enrollments for the user
// For more details see https://api.voiceit.io/#get-face-enrollments
func (vi VoiceIt2) GetAllFaceEnrollments(userId string) (string, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/enrollments/face/"+userId+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in GetAllFaceEnrollments():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in GetAllFaceEnrollments():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in GetAllFaceEnrollments():", err)
		return "", err
	}
	return string(reply), nil
}

// CreateVoiceEnrollment takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a audio recording to create a voice enrollment for the user
// For more details see https://api.voiceit.io/#create-voice-enrollment
func (vi VoiceIt2) CreateVoiceEnrollment(userId string, contentLanguage string, phrase string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in CreateVoiceEnrollment():", err)
		return "", err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in CreateVoiceEnrollment():", err)
		return "", err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("recording", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error in CreateVoiceEnrollment():", err)
		return "", err
	}
	part.Write(fileContents)

	err = writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in CreateVoiceEnrollment():", err)
		return "", err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in CreateVoiceEnrollment():", err)
		return "", err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in CreateVoiceEnrollment():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/voice"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in CreateVoiceEnrollment():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in CreateVoiceEnrollment():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in CreateVoiceEnrollment():", err)
		return "", err
	}
	return string(reply), nil
}

// CreateVoiceEnrollmentByUrl takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a audio recording to create a voice enrollment for the user
// For more details see https://api.voiceit.io/#create-voice-enrollment-by-url
func (vi VoiceIt2) CreateVoiceEnrollmentByUrl(userId string, contentLanguage string, phrase string, fileUrl string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in CreateVoiceEnrollmentByUrl():", err)
		return "", err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in CreateVoiceEnrollmentByUrl():", err)
		return "", err
	}

	err = writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		log.Println("Error in CreateVoiceEnrollmentByUrl():", err)
		return "", err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in CreateVoiceEnrollmentByUrl():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/voice/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in CreateVoiceEnrollmentByUrl():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in CreateVoiceEnrollmentByUrl():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in CreateVoiceEnrollmentByUrl():", err)
		return "", err
	}
	return string(reply), nil
}

// CreateFaceEnrollment takes the userId generated during a createUser and
// absolute file path for a video recording to create a face enrollment for the user
// For more details see https://api.voiceit.io/#create-face-enrollment
func (vi VoiceIt2) CreateFaceEnrollment(userId string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in CreateFaceEnrollment():", err)
		return "", err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in CreateFaceEnrollment():", err)
		return "", err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error in CreateFaceEnrollment():", err)
		return "", err
	}
	part.Write(fileContents)
	err = writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in CreateFaceEnrollment():", err)
		return "", err
	}
	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/face"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in CreateFaceEnrollment():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in CreateFaceEnrollment():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in CreateFaceEnrollment():", err)
		return "", err
	}
	return string(reply), nil
}

// CreateFaceEnrollmentByUrl takes the userId generated during a createUser
// and a fully qualified URL to a video recording to verify the user's face
// For more details see https://api.voiceit.io/#create-face-enrollment-by-url
func (vi VoiceIt2) CreateFaceEnrollmentByUrl(userId string, fileUrl string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in CreateFaceEnrollmentByUrl():", err)
		return "", err
	}

	err = writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		log.Println("Error in CreateFaceEnrollmentByUrl():", err)
		return "", err
	}
	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/face/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in CreateFaceEnrollmentByUrl():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in CreateFaceEnrollmentByUrl():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in CreateFaceEnrollmentByUrl():", err)
		return "", err
	}
	return string(reply), nil
}

// CreateVideoEnrollment takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a video recording to create a video enrollment for the user
// For more details see https://api.voiceit.io/#create-video-enrollment
func (vi VoiceIt2) CreateVideoEnrollment(userId string, contentLanguage string, phrase string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in CreateVideoEnrollment():", err)
		return "", err
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in CreateVideoEnrollment():", err)
		return "", err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error in CreateVideoEnrollment():", err)
		return "", err
	}
	part.Write(fileContents)
	err = writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in CreateVideoEnrollment():", err)
		return "", err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in CreateVideoEnrollment():", err)
		return "", err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in CreateVideoEnrollment():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/video"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in CreateVideoEnrollment():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in CreateVideoEnrollment():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in CreateVideoEnrollment():", err)
		return "", err
	}
	return string(reply), nil
}

// CreateVideoEnrollment takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a video recording to create a video enrollment for the user
// For more details see https://api.voiceit.io/#create-video-enrollment-by-url
func (vi VoiceIt2) CreateVideoEnrollmentByUrl(userId string, contentLanguage string, phrase string, fileUrl string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in CreateVideoEnrollmentByUrl():", err)
		return "", err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in CreateVideoEnrollmentByUrl():", err)
		return "", err
	}

	err = writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		log.Println("Error in CreateVideoEnrollmentByUrl():", err)
		return "", err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in CreateVideoEnrollmentByUrl():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/video/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in CreateVideoEnrollmentByUrl():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in CreateVideoEnrollmentByUrl():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in CreateVideoEnrollmentByUrl():", err)
		return "", err
	}
	return string(reply), nil
}

// DeleteFaceEnrollment takes the userId generated during a createUser and
// a faceEnrollmentId returned during a faceEnrollment and deletes the specific
// faceEnrollment for the user
// For more details see https://api.voiceit.io/#delete-face-enrollment
func (vi VoiceIt2) DeleteFaceEnrollment(userId string, faceEnrollmentId int) (string, error) {
	req, err := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/face/"+userId+"/"+strconv.Itoa(faceEnrollmentId)+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in DeleteFaceEnrollment():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in DeleteFaceEnrollment():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in DeleteFaceEnrollment():", err)
		return "", err
	}
	return string(reply), nil
}

// DeleteAllFaceEnrollments takes the userId generated during a createUser
// For more details see https://api.voiceit.io/#delete-all-face-enrollments
func (vi VoiceIt2) DeleteAllFaceEnrollments(userId string) (string, error) {
	req, err := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/"+userId+"/face"+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in DeleteAllFaceEnrollments():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in DeleteAllFaceEnrollments():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in DeleteAllFaceEnrollments():", err)
		return "", err
	}
	return string(reply), nil
}

// DeleteVoiceEnrollment takes the userId generated during a createUser and
// an enrollmentId returned during a voiceEnrollment/videoEnrollment and deletes
// the voice enrollment for the user
// For more details see https://api.voiceit.io/#delete-voice-enrollment
func (vi VoiceIt2) DeleteVoiceEnrollment(userId string, id int) (string, error) {
	req, err := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/voice/"+userId+"/"+strconv.Itoa(id)+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in DeleteVoiceEnrollment():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in DeleteVoiceEnrollment():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in DeleteVoiceEnrollment():", err)
		return "", err
	}
	return string(reply), nil
}

// DeleteAllVoiceEnrollments takes the userId generated during a createUser
// For more details https://api.voiceit.io/#delete-all-voice-enrollments
func (vi VoiceIt2) DeleteAllVoiceEnrollments(userId string) (string, error) {
	req, err := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/"+userId+"/voice"+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in DeleteAllVoiceEnrollments():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in DeleteAllVoiceEnrollments():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in DeleteAllVoiceEnrollments():", err)
		return "", err
	}
	return string(reply), nil
}

// DeleteVideoEnrollment takes the userId generated during a createUser and
// an enrollmentId returned during a voiceEnrollment/videoEnrollment and deletes
// the voice/video enrollment for the user
// For more details see https://api.voiceit.io/#delete-video-enrollment
func (vi VoiceIt2) DeleteVideoEnrollment(userId string, id int) (string, error) {
	req, err := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/video/"+userId+"/"+strconv.Itoa(id)+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in DeleteVideoEnrollment():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in DeleteVideoEnrollment():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in DeleteVideoEnrollment():", err)
		return "", err
	}
	return string(reply), nil
}

// DeleteAllVideoEnrollments takes the userId generated during a createUser
// For more details see https://api.voiceit.io/#delete-all-video-enrollments
func (vi VoiceIt2) DeleteAllVideoEnrollments(userId string) (string, error) {
	req, err := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/"+userId+"/video"+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in DeleteAllVideoEnrollments():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in DeleteAllVideoEnrollments():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in DeleteAllVideoEnrollments():", err)
		return "", err
	}
	return string(reply), nil
}

// DeleteAllEnrollments takes the userId generated during a createUser
// and deletes all video/voice enrollments for the user
// For more details see https://api.voiceit.io/#delete-all-enrollments-for-user
func (vi VoiceIt2) DeleteAllEnrollments(userId string) (string, error) {
	req, err := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/"+userId+"/all"+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in DeleteAllEnrollments():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in DeleteAllEnrollments():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in DeleteAllEnrollments():", err)
		return "", err
	}
	return string(reply), nil
}

// VoiceVerification takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a audio recording to verify the user's voice
// For more details see https://api.voiceit.io/#verify-a-user-s-voice
func (vi VoiceIt2) VoiceVerification(userId string, contentLanguage string, phrase string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in VoiceVerification():", err)
		return "", err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in VoiceVerification():", err)
		return "", err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("recording", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error in VoiceVerification():", err)
		return "", err
	}
	part.Write(fileContents)

	err = writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in VoiceVerification():", err)
		return "", err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VoiceVerification():", err)
		return "", err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VoiceVerification():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/voice"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceVerification():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceVerification():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceVerification():", err)
		return "", err
	}
	return string(reply), nil
}

// VoiceVerificationByUrl takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a audio recording to verify the user's voice
// For more details see https://api.voiceit.io/#verify-a-user-s-voice-by-url
func (vi VoiceIt2) VoiceVerificationByUrl(userId string, contentLanguage string, phrase string, fileUrl string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in VoiceVerificationByUrl():", err)
		return "", err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VoiceVerificationByUrl():", err)
		return "", err
	}

	err = writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		log.Println("Error in VoiceVerificationByUrl():", err)
		return "", err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VoiceVerificationByUrl():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/voice/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceVerificationByUrl():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceVerificationByUrl():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceVerificationByUrl():", err)
		return "", err
	}
	return string(reply), nil
}

// FaceVerification takes the userId generated during a createUser and a
// absolute file path for a video recording to verify the user's face
// For more details see https://api.voiceit.io/#verify-a-user-s-face
func (vi VoiceIt2) FaceVerification(userId string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in FaceVerification():", err)
		return "", err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in FaceVerification():", err)
		return "", err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error in FaceVerification():", err)
		return "", err
	}
	part.Write(fileContents)

	err = writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in FaceVerification():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/face"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in FaceVerification():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in FaceVerification():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in FaceVerification():", err)
		return "", err
	}
	return string(reply), nil
}

// FaceVerificationByUrl takes the userId generated during a createUser
// and a fully qualified URL to a video recording to verify the user's face
// For more details see https://api.voiceit.io/#verify-a-user-s-face-by-url
func (vi VoiceIt2) FaceVerificationByUrl(userId string, fileUrl string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	err := writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		log.Println("Error in FaceVerificationByUrl():", err)
		return "", err
	}

	err = writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in FaceVerificationByUrl():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/face/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in FaceVerificationByUrl():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in FaceVerificationByUrl():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in FaceVerificationByUrl():", err)
		return "", err
	}
	return string(reply), nil
}

// VideoVerification takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a video recording to verify the user's face and voice
// For more details see https://api.voiceit.io/#video-verification
func (vi VoiceIt2) VideoVerification(userId string, contentLanguage string, phrase string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in VideoVerification():", err)
		return "", err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in VideoVerification():", err)
		return "", err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error in VideoVerification():", err)
		return "", err
	}
	part.Write(fileContents)

	err = writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in VideoVerification():", err)
		return "", err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VideoVerification():", err)
		return "", err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VideoVerification():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/video"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VideoVerification():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VideoVerification():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VideoVerification():", err)
		return "", err
	}
	return string(reply), nil
}

// VideoVerificationByUrl takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a video recording to verify the user's face and voice
// For more details see https://api.voiceit.io/#video-verification-by-url
func (vi VoiceIt2) VideoVerificationByUrl(userId string, contentLanguage string, phrase string, fileUrl string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in VideoVerificationByUrl():", err)
		return "", err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VideoVerificationByUrl():", err)
		return "", err
	}

	err = writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		log.Println("Error in VideoVerificationByUrl():", err)
		return "", err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VideoVerificationByUrl():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/video/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VideoVerificationByUrl():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VideoVerificationByUrl():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VideoVerificationByUrl():", err)
		return "", err
	}
	return string(reply), nil
}

// VoiceIdentification takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a audio recording to idetify the user's voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-voice
func (vi VoiceIt2) VoiceIdentification(groupId string, contentLanguage string, phrase string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in VoiceIdentification():", err)
		return "", err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in VoiceIdentification():", err)
		return "", err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("recording", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error in VoiceIdentification():", err)
		return "", err
	}
	part.Write(fileContents)

	err = writer.WriteField("groupId", groupId)
	if err != nil {
		log.Println("Error in VoiceIdentification():", err)
		return "", err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VoiceIdentification():", err)
		return "", err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VoiceIdentification():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/voice"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIdentification():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIdentification():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIdentification():", err)
		return "", err
	}
	return string(reply), nil
}

// VoiceIdentificationByUrl takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a audio recording to idetify the user's voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-voice-by-url
func (vi VoiceIt2) VoiceIdentificationByUrl(groupId string, contentLanguage string, phrase string, fileUrl string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		log.Println("Error in VoiceIdentificationByUrl():", err)
		return "", err
	}

	err = writer.WriteField("groupId", groupId)
	if err != nil {
		log.Println("Error in VoiceIdentificationByUrl():", err)
		return "", err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VoiceIdentificationByUrl():", err)
		return "", err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VoiceIdentificationByUrl():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/voice/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIdentificationByUrl():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIdentificationByUrl():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIdentificationByUrl():", err)
		return "", err
	}
	return string(reply), nil
}

// VideoIdentification takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a video recording to idetify the user's face and voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-voice-amp-face
func (vi VoiceIt2) VideoIdentification(groupId string, contentLanguage string, phrase string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in VideoIdentification():", err)
		return "", err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in VideoIdentification():", err)
		return "", err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error in VideoIdentification():", err)
		return "", err
	}
	part.Write(fileContents)

	err = writer.WriteField("groupId", groupId)
	if err != nil {
		log.Println("Error in VideoIdentification():", err)
		return "", err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VideoIdentification():", err)
		return "", err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VideoIdentification():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/video"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VideoIdentification():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VideoIdentification():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VideoIdentification():", err)
		return "", err
	}
	return string(reply), nil
}

// VideoIdentificationByUrl takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a video recording to idetify the user's face and voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-voice-amp-face-by-url
func (vi VoiceIt2) VideoIdentificationByUrl(groupId string, contentLanguage string, phrase string, fileUrl string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		log.Println("Error in VideoIdentificationByUrl():", err)
		return "", err
	}

	err = writer.WriteField("groupId", groupId)
	if err != nil {
		log.Println("Error in VideoIdentificationByUrl():", err)
		return "", err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VideoIdentificationByUrl():", err)
		return "", err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VideoIdentificationByUrl():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/video/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VideoIdentificationByUrl():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VideoIdentificationByUrl():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VideoIdentificationByUrl():", err)
		return "", err
	}
	return string(reply), nil
}

// FaceIdentification takes the groupId generated during a createGroup,
// and absolute file path for a face recording to idetify the user's face
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-face
func (vi VoiceIt2) FaceIdentification(groupId string, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in FaceIdentification():", err)
		return "", err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in FaceIdentification():", err)
		return "", err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error in FaceIdentification():", err)
		return "", err
	}
	part.Write(fileContents)

	err = writer.WriteField("groupId", groupId)
	if err != nil {
		log.Println("Error in FaceIdentification():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/face"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in FaceIdentification():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in FaceIdentification():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in FaceIdentification():", err)
		return "", err
	}
	return string(reply), nil
}

// FaceIdentificationByUrl takes the groupId generated during a createGroup,
// and a fully qualified URL to a face recording to idetify the user's face
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-face-by-url
func (vi VoiceIt2) FaceIdentificationByUrl(groupId string, fileUrl string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		log.Println("Error in FaceIdentificationByUrl():", err)
		return "", err
	}

	err = writer.WriteField("groupId", groupId)
	if err != nil {
		log.Println("Error in FaceIdentificationByUrl():", err)
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/face/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in FaceIdentificationByUrl():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in FaceIdentificationByUrl():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in FaceIdentificationByUrl():", err)
		return "", err
	}
	return string(reply), nil
}

// GetPhrases takes the contentLanguage
// For more details see https://api.voiceit.io/#get-phrases
func (vi VoiceIt2) GetPhrases(contentLanguage string) (string, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/phrases/"+contentLanguage+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in GetPhrases():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in GetPhrases():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in GetPhrases():", err)
		return "", err
	}
	return string(reply), nil
}

// CreateUserToken takes the userId (string) and a timeout (time.Duration).
// The returned user token can be used to construct a new VoiceIt2 instance which has user level rights for the given user.
// The timeout controls the expiration of the user token.
// For more details see https://api.voiceit.io/?go#user-token-generation
func (vi VoiceIt2) CreateUserToken(userId string, timeout time.Duration) (string, error) {

	var req *http.Request
	req, err := http.NewRequest("POST", vi.BaseUrl+"/users/"+userId+"/token"+"?timeOut="+strconv.Itoa(int(timeout.Seconds())), nil)
	if err != nil {
		log.Println("Error in CreateUserToken():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in CreateUserToken():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in CreateUserToken():", err)
		return "", err
	}
	return string(reply), nil
}

// ExpireUserTokens takes a userId (string).
// For more details see https://api.voiceit.io/?go#user-token-expiration
func (vi VoiceIt2) ExpireUserTokens(userId string) (string, error) {
	req, err := http.NewRequest("POST", vi.BaseUrl+"/users/"+userId+"/expireTokens"+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in ExpireUserTokens():", err)
		return "", err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in ExpireUserTokens():", err)
		return "", err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in ExpireUserTokens():", err)
		return "", err
	}
	return string(reply), nil
}
