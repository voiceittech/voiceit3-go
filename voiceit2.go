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

const PlatformVersion string = "v2.2.2"

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
func (vi VoiceIt2) GetAllUsers() ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/users"+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetAllUsers():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetAllUsers():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetAllUsers():", err)
		return []byte{}, err
	}
	return reply, nil
}

// CreateUser creates a new user profile and returns a unique userId
// that is used for all future calls related to the user profile
// For more details see https://api.voiceit.io/#create-a-user
func (vi VoiceIt2) CreateUser() ([]byte, error) {
	req, err := http.NewRequest("POST", vi.BaseUrl+"/users"+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateUser():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateUser():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateUser():", err)
		return []byte{}, err
	}
	return reply, nil
}

// CheckUserExists takes the userId generated during a createUser and returns
// an object which contains the boolean "exists" which shows whether a given user exists
// For more details see https://api.voiceit.io/#check-if-a-specific-user-exists
func (vi VoiceIt2) CheckUserExists(userId string) ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/users/"+userId+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CheckUserExists():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CheckUserExists():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CheckUserExists():", err)
		return []byte{}, err
	}
	return reply, nil
}

// DeleteUser takes the userId generated during a createUser and deletes
// the user profile and all associated face and voice enrollments
// For more details see https://api.voiceit.io/#delete-a-specific-user
func (vi VoiceIt2) DeleteUser(userId string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", vi.BaseUrl+"/users/"+userId+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.DeleteUser():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.DeleteUser():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.DeleteUser():", err)
		return []byte{}, err
	}
	return reply, nil
}

// GetGroupsForUser takes the userId generated during a createUser and returns
// a list of all groups that the user belongs to
// For more details see https://api.voiceit.io/#get-groups-for-user
func (vi VoiceIt2) GetGroupsForUser(userId string) ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/users/"+userId+"/groups"+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetGroupsForUser():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetGroupsForUser():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetGroupsForUser():", err)
		return []byte{}, err
	}
	return reply, nil
}

// GetAllGroups returns a list of all groups associated with the API Key
// For more details see https://api.voiceit.io/#get-all-groups
func (vi VoiceIt2) GetAllGroups() ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/groups"+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetAllGroups():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetAllGroups():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetAllGroups():", err)
		return []byte{}, err
	}
	return reply, nil
}

// GetGroup takes the groupId generated during a createGroup
// and returns the group along with a list of associated users in the group
// For more details see https://api.voiceit.io/#get-a-specific-group
func (vi VoiceIt2) GetGroup(groupId string) ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/groups/"+groupId+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetGroup():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetGroup():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetGroup():", err)
		return []byte{}, err
	}
	return reply, nil
}

// CheckGroupExists takes the groupId generated during a createGroup
// and returns whether the group exists for the given groupId
// For more details see https://api.voiceit.io/#check-if-group-exists
func (vi VoiceIt2) CheckGroupExists(groupId string) ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/groups/"+groupId+"/exists"+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CheckGroupExists():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CheckGroupExists():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CheckGroupExists():", err)
		return []byte{}, err
	}
	return reply, nil
}

// CreateGroup creates a new group profile and returns a unique groupId
// that is used for all future calls related to the group
// For more details see https://api.voiceit.io/#create-a-group
func (vi VoiceIt2) CreateGroup(description string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("description", description)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateGroup():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/groups"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateGroup():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateGroup():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateGroup():", err)
		return []byte{}, err
	}
	return reply, nil
}

// AddUserToGroup takes the groupId generated during a createGroup
// and the userId generated during createUser and adds the user to the group
// For more details see https://api.voiceit.io/#add-user-to-group
func (vi VoiceIt2) AddUserToGroup(groupId string, userId string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("groupId", groupId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.AddUserToGroup():", err)
		return []byte{}, err
	}

	err = writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.AddUserToGroup():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("PUT", vi.BaseUrl+"/groups/addUser"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.AddUserToGroup():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.AddUserToGroup():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.AddUserToGroup():", err)
		return []byte{}, err
	}
	return reply, nil
}

// RemoveUserFromGroup takes the groupId generated during a createGroup
// and the userId generated during createUser and removes the user from the group
// For more details see https://api.voiceit.io/#remove-user-from-group
func (vi VoiceIt2) RemoveUserFromGroup(groupId string, userId string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("groupId", groupId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.RemoveUserFromGroup():", err)
		return []byte{}, err
	}

	err = writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.RemoveUserFromGroup():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("PUT", vi.BaseUrl+"/groups/removeUser"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.RemoveUserFromGroup():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.RemoveUserFromGroup():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.RemoveUserFromGroup():", err)
		return []byte{}, err
	}
	return reply, nil
}

// DeleteGroup takes the groupId generated during a createGroup and deletes
// the group profile disassociates all users associated with it
// For more details see https://api.voiceit.io/#delete-a-specific-group
func (vi VoiceIt2) DeleteGroup(groupId string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req, err := http.NewRequest("DELETE", vi.BaseUrl+"/groups/"+groupId+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.DeleteGroup():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.DeleteGroup():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.DeleteGroup():", err)
		return []byte{}, err
	}
	return reply, nil
}

// GetAllVoiceEnrollments takes the userId generated during a createUser
// and returns a list of all voice enrollments for the user
// For more details see https://api.voiceit.io/#get-voice-enrollments
func (vi VoiceIt2) GetAllVoiceEnrollments(userId string) ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/enrollments/voice/"+userId+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetAllVoiceEnrollments():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetAllVoiceEnrollments():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetAllVoiceEnrollments():", err)
		return []byte{}, err
	}
	return reply, nil
}

// GetAllVideoEnrollments takes the userId generated during a createUser
// and returns a list of all video enrollments for the user
// For more details see https://api.voiceit.io/#get-video-enrollments
func (vi VoiceIt2) GetAllVideoEnrollments(userId string) ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/enrollments/video/"+userId+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetAllVideoEnrollments():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetAllVideoEnrollments():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetAllVideoEnrollments():", err)
		return []byte{}, err
	}
	return reply, nil
}

// GetAllFaceEnrollments takes the userId generated during a createUser
// and returns a list of all face enrollments for the user
// For more details see https://api.voiceit.io/#get-face-enrollments
func (vi VoiceIt2) GetAllFaceEnrollments(userId string) ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/enrollments/face/"+userId+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetAllFaceEnrollments():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetAllFaceEnrollments():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetAllFaceEnrollments():", err)
		return []byte{}, err
	}
	return reply, nil
}

// CreateVoiceEnrollment takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a audio recording to create a voice enrollment for the user
// For more details see https://api.voiceit.io/#create-voice-enrollment
func (vi VoiceIt2) CreateVoiceEnrollment(userId string, contentLanguage string, phrase string, filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVoiceEnrollment():", err)
		return []byte{}, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVoiceEnrollment():", err)
		return []byte{}, err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("recording", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVoiceEnrollment():", err)
		return []byte{}, err
	}
	part.Write(fileContents)

	err = writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVoiceEnrollment():", err)
		return []byte{}, err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVoiceEnrollment():", err)
		return []byte{}, err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVoiceEnrollment():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/voice"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVoiceEnrollment():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVoiceEnrollment():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVoiceEnrollment():", err)
		return []byte{}, err
	}
	return reply, nil
}

// CreateVoiceEnrollmentByUrl takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a audio recording to create a voice enrollment for the user
// For more details see https://api.voiceit.io/#create-voice-enrollment-by-url
func (vi VoiceIt2) CreateVoiceEnrollmentByUrl(userId string, contentLanguage string, phrase string, fileUrl string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVoiceEnrollmentByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVoiceEnrollmentByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVoiceEnrollmentByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVoiceEnrollmentByUrl():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/voice/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVoiceEnrollmentByUrl():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVoiceEnrollmentByUrl():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVoiceEnrollmentByUrl():", err)
		return []byte{}, err
	}
	return reply, nil
}

// CreateFaceEnrollment takes the userId generated during a createUser and
// absolute file path for a video recording to create a face enrollment for the user
// For more details see https://api.voiceit.io/#create-face-enrollment
func (vi VoiceIt2) CreateFaceEnrollment(userId string, filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateFaceEnrollment():", err)
		return []byte{}, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateFaceEnrollment():", err)
		return []byte{}, err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateFaceEnrollment():", err)
		return []byte{}, err
	}
	part.Write(fileContents)
	err = writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateFaceEnrollment():", err)
		return []byte{}, err
	}
	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/face"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateFaceEnrollment():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateFaceEnrollment():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateFaceEnrollment():", err)
		return []byte{}, err
	}
	return reply, nil
}

// CreateFaceEnrollmentByUrl takes the userId generated during a createUser
// and a fully qualified URL to a video recording to verify the user's face
// For more details see https://api.voiceit.io/#create-face-enrollment-by-url
func (vi VoiceIt2) CreateFaceEnrollmentByUrl(userId string, fileUrl string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateFaceEnrollmentByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateFaceEnrollmentByUrl():", err)
		return []byte{}, err
	}
	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/face/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateFaceEnrollmentByUrl():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateFaceEnrollmentByUrl():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateFaceEnrollmentByUrl():", err)
		return []byte{}, err
	}
	return reply, nil
}

// CreateVideoEnrollment takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a video recording to create a video enrollment for the user
// For more details see https://api.voiceit.io/#create-video-enrollment
func (vi VoiceIt2) CreateVideoEnrollment(userId string, contentLanguage string, phrase string, filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVideoEnrollment():", err)
		return []byte{}, err
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVideoEnrollment():", err)
		return []byte{}, err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVideoEnrollment():", err)
		return []byte{}, err
	}
	part.Write(fileContents)
	err = writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVideoEnrollment():", err)
		return []byte{}, err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVideoEnrollment():", err)
		return []byte{}, err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVideoEnrollment():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/video"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVideoEnrollment():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVideoEnrollment():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVideoEnrollment():", err)
		return []byte{}, err
	}
	return reply, nil
}

// CreateVideoEnrollment takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a video recording to create a video enrollment for the user
// For more details see https://api.voiceit.io/#create-video-enrollment-by-url
func (vi VoiceIt2) CreateVideoEnrollmentByUrl(userId string, contentLanguage string, phrase string, fileUrl string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVideoEnrollmentByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVideoEnrollmentByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVideoEnrollmentByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVideoEnrollmentByUrl():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/video/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVideoEnrollmentByUrl():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVideoEnrollmentByUrl():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateVideoEnrollmentByUrl():", err)
		return []byte{}, err
	}
	return reply, nil
}

// DeleteAllEnrollments takes the userId generated during a createUser
// and deletes all video/voice enrollments for the user
// For more details see https://api.voiceit.io/#delete-all-enrollments-for-user
func (vi VoiceIt2) DeleteAllEnrollments(userId string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/"+userId+"/all"+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.DeleteAllEnrollments():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.DeleteAllEnrollments():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.DeleteAllEnrollments():", err)
		return []byte{}, err
	}
	return reply, nil
}

// VoiceVerification takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a audio recording to verify the user's voice
// For more details see https://api.voiceit.io/#verify-a-user-s-voice
func (vi VoiceIt2) VoiceVerification(userId string, contentLanguage string, phrase string, filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceVerification():", err)
		return []byte{}, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceVerification():", err)
		return []byte{}, err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("recording", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceVerification():", err)
		return []byte{}, err
	}
	part.Write(fileContents)

	err = writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceVerification():", err)
		return []byte{}, err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceVerification():", err)
		return []byte{}, err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceVerification():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/voice"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceVerification():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceVerification():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceVerification():", err)
		return []byte{}, err
	}
	return reply, nil
}

// VoiceVerificationByUrl takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a audio recording to verify the user's voice
// For more details see https://api.voiceit.io/#verify-a-user-s-voice-by-url
func (vi VoiceIt2) VoiceVerificationByUrl(userId string, contentLanguage string, phrase string, fileUrl string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceVerificationByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceVerificationByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceVerificationByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceVerificationByUrl():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/voice/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceVerificationByUrl():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceVerificationByUrl():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceVerificationByUrl():", err)
		return []byte{}, err
	}
	return reply, nil
}

// FaceVerification takes the userId generated during a createUser and a
// absolute file path for a video recording to verify the user's face
// For more details see https://api.voiceit.io/#verify-a-user-s-face
func (vi VoiceIt2) FaceVerification(userId string, filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceVerification():", err)
		return []byte{}, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceVerification():", err)
		return []byte{}, err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceVerification():", err)
		return []byte{}, err
	}
	part.Write(fileContents)

	err = writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceVerification():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/face"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceVerification():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceVerification():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceVerification():", err)
		return []byte{}, err
	}
	return reply, nil
}

// FaceVerificationByUrl takes the userId generated during a createUser
// and a fully qualified URL to a video recording to verify the user's face
// For more details see https://api.voiceit.io/#verify-a-user-s-face-by-url
func (vi VoiceIt2) FaceVerificationByUrl(userId string, fileUrl string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	err := writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceVerificationByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceVerificationByUrl():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/face/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceVerificationByUrl():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceVerificationByUrl():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceVerificationByUrl():", err)
		return []byte{}, err
	}
	return reply, nil
}

// VideoVerification takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a video recording to verify the user's face and voice
// For more details see https://api.voiceit.io/#video-verification
func (vi VoiceIt2) VideoVerification(userId string, contentLanguage string, phrase string, filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoVerification():", err)
		return []byte{}, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoVerification():", err)
		return []byte{}, err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoVerification():", err)
		return []byte{}, err
	}
	part.Write(fileContents)

	err = writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoVerification():", err)
		return []byte{}, err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoVerification():", err)
		return []byte{}, err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoVerification():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/video"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoVerification():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoVerification():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoVerification():", err)
		return []byte{}, err
	}
	return reply, nil
}

// VideoVerificationByUrl takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a video recording to verify the user's face and voice
// For more details see https://api.voiceit.io/#video-verification-by-url
func (vi VoiceIt2) VideoVerificationByUrl(userId string, contentLanguage string, phrase string, fileUrl string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("userId", userId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoVerificationByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoVerificationByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoVerificationByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoVerificationByUrl():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/video/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoVerificationByUrl():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoVerificationByUrl():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoVerificationByUrl():", err)
		return []byte{}, err
	}
	return reply, nil
}

// VoiceIdentification takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a audio recording to idetify the user's voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-voice
func (vi VoiceIt2) VoiceIdentification(groupId string, contentLanguage string, phrase string, filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceIdentification():", err)
		return []byte{}, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceIdentification():", err)
		return []byte{}, err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("recording", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceIdentification():", err)
		return []byte{}, err
	}
	part.Write(fileContents)

	err = writer.WriteField("groupId", groupId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceIdentification():", err)
		return []byte{}, err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceIdentification():", err)
		return []byte{}, err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceIdentification():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/voice"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceIdentification():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceIdentification():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceIdentification():", err)
		return []byte{}, err
	}
	return reply, nil
}

// VoiceIdentificationByUrl takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a audio recording to idetify the user's voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-voice-by-url
func (vi VoiceIt2) VoiceIdentificationByUrl(groupId string, contentLanguage string, phrase string, fileUrl string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceIdentificationByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("groupId", groupId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceIdentificationByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceIdentificationByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceIdentificationByUrl():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/voice/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceIdentificationByUrl():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceIdentificationByUrl():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VoiceIdentificationByUrl():", err)
		return []byte{}, err
	}
	return reply, nil
}

// VideoIdentification takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a video recording to idetify the user's face and voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-voice-amp-face
func (vi VoiceIt2) VideoIdentification(groupId string, contentLanguage string, phrase string, filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoIdentification():", err)
		return []byte{}, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoIdentification():", err)
		return []byte{}, err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoIdentification():", err)
		return []byte{}, err
	}
	part.Write(fileContents)

	err = writer.WriteField("groupId", groupId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoIdentification():", err)
		return []byte{}, err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoIdentification():", err)
		return []byte{}, err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoIdentification():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/video"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoIdentification():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoIdentification():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoIdentification():", err)
		return []byte{}, err
	}
	return reply, nil
}

// VideoIdentificationByUrl takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a video recording to idetify the user's face and voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-voice-amp-face-by-url
func (vi VoiceIt2) VideoIdentificationByUrl(groupId string, contentLanguage string, phrase string, fileUrl string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoIdentificationByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("groupId", groupId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoIdentificationByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("contentLanguage", contentLanguage)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoIdentificationByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("phrase", phrase)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoIdentificationByUrl():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/video/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoIdentificationByUrl():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoIdentificationByUrl():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.VideoIdentificationByUrl():", err)
		return []byte{}, err
	}
	return reply, nil
}

// FaceIdentification takes the groupId generated during a createGroup,
// and absolute file path for a face recording to idetify the user's face
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-face
func (vi VoiceIt2) FaceIdentification(groupId string, filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceIdentification():", err)
		return []byte{}, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceIdentification():", err)
		return []byte{}, err
	}
	file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", filepath.Base(file.Name()))
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceIdentification():", err)
		return []byte{}, err
	}
	part.Write(fileContents)

	err = writer.WriteField("groupId", groupId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceIdentification():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/face"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceIdentification():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceIdentification():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceIdentification():", err)
		return []byte{}, err
	}
	return reply, nil
}

// FaceIdentificationByUrl takes the groupId generated during a createGroup,
// and a fully qualified URL to a face recording to idetify the user's face
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-face-by-url
func (vi VoiceIt2) FaceIdentificationByUrl(groupId string, fileUrl string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceIdentificationByUrl():", err)
		return []byte{}, err
	}

	err = writer.WriteField("groupId", groupId)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceIdentificationByUrl():", err)
		return []byte{}, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/face/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceIdentificationByUrl():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceIdentificationByUrl():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.FaceIdentificationByUrl():", err)
		return []byte{}, err
	}
	return reply, nil
}

// GetPhrases takes the contentLanguage
// For more details see https://api.voiceit.io/#get-phrases
func (vi VoiceIt2) GetPhrases(contentLanguage string) ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/phrases/"+contentLanguage+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetPhrases():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetPhrases():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.GetPhrases():", err)
		return []byte{}, err
	}
	return reply, nil
}

// CreateUserToken takes the userId (string) and a timeout (time.Duration).
// The returned user token can be used to construct a new VoiceIt2 instance which has user level rights for the given user.
// The timeout controls the expiration of the user token.
// For more details see https://api.voiceit.io/?go#user-token-generation
func (vi VoiceIt2) CreateUserToken(userId string, timeout time.Duration) ([]byte, error) {

	var req *http.Request
	req, err := http.NewRequest("POST", vi.BaseUrl+"/users/"+userId+"/token"+"?timeOut="+strconv.Itoa(int(timeout.Seconds())), nil)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateUserToken():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateUserToken():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.CreateUserToken():", err)
		return []byte{}, err
	}
	return reply, nil
}

// ExpireUserTokens takes a userId (string).
// For more details see https://api.voiceit.io/?go#user-token-expiration
func (vi VoiceIt2) ExpireUserTokens(userId string) ([]byte, error) {
	req, err := http.NewRequest("POST", vi.BaseUrl+"/users/"+userId+"/expireTokens"+vi.NotificationUrl, nil)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.ExpireUserTokens():", err)
		return []byte{}, err
	}
	req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
	req.Header.Add("platformId", "39")
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.ExpireUserTokens():", err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in VoiceIt2-Go.ExpireUserTokens():", err)
		return []byte{}, err
	}
	return reply, nil
}
