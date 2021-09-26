package voiceit2

import (
	"bytes"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/voiceittech/VoiceIt2-Go/v2/structs"
)

var (
	PlatformVersion string = "v2.6.1"
	PlatformId      string = "39"
)

type VoiceIt2 struct {
	APIKey          string
	APIToken        string
	BaseUrl         string
	NotificationUrl string
}

// NewClient returns a new VoiceIt2 client
func NewClient(key, tok string) VoiceIt2 {
	return VoiceIt2{
		APIKey:          key,
		APIToken:        tok,
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
		return []byte{}, errors.New("GetAllUsers Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("GetAllUsers Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("GetAllUsers Exception: " + err.Error())
	}
	return reply, nil
}

// CreateUser creates a new user profile and returns a unique userId
// that is used for all future calls related to the user profile
// For more details see https://api.voiceit.io/#create-a-user
func (vi VoiceIt2) CreateUser() ([]byte, error) {
	req, err := http.NewRequest("POST", vi.BaseUrl+"/users"+vi.NotificationUrl, nil)
	if err != nil {
		return []byte{}, errors.New("CreateUser Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("CreateUser Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("CreateUser Exception: " + err.Error())
	}
	return reply, nil
}

// CheckUserExists takes the userId generated during a createUser and returns
// an object which contains the boolean "exists" which shows whether a given user exists
// For more details see https://api.voiceit.io/#check-if-a-specific-user-exists
func (vi VoiceIt2) CheckUserExists(userId string) ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/users/"+userId+vi.NotificationUrl, nil)
	if err != nil {
		return []byte{}, errors.New("CheckUserExists Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("CheckUserExists Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("CheckUserExists Exception: " + err.Error())
	}
	return reply, nil
}

// DeleteUser takes the userId generated during a createUser and deletes
// the user profile and all associated face and voice enrollments
// For more details see https://api.voiceit.io/#delete-a-specific-user
func (vi VoiceIt2) DeleteUser(userId string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", vi.BaseUrl+"/users/"+userId+vi.NotificationUrl, nil)
	if err != nil {
		return []byte{}, errors.New("DeleteUser Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("DeleteUser Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("DeleteUser Exception: " + err.Error())
	}
	return reply, nil
}

// GetGroupsForUser takes the userId generated during a createUser and returns
// a list of all groups that the user belongs to
// For more details see https://api.voiceit.io/#get-groups-for-user
func (vi VoiceIt2) GetGroupsForUser(userId string) ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/users/"+userId+"/groups"+vi.NotificationUrl, nil)
	if err != nil {
		return []byte{}, errors.New("GetGroupsForUser Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("GetGroupsForUser Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("GetGroupsForUser Exception: " + err.Error())
	}
	return reply, nil
}

// GetAllGroups returns a list of all groups associated with the API Key
// For more details see https://api.voiceit.io/#get-all-groups
func (vi VoiceIt2) GetAllGroups() ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/groups"+vi.NotificationUrl, nil)
	if err != nil {
		return []byte{}, errors.New("GetAllGroups Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("GetAllGroups Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("GetAllGroups Exception: " + err.Error())
	}
	return reply, nil
}

// GetGroup takes the groupId generated during a createGroup
// and returns the group along with a list of associated users in the group
// For more details see https://api.voiceit.io/#get-a-specific-group
func (vi VoiceIt2) GetGroup(groupId string) ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/groups/"+groupId+vi.NotificationUrl, nil)
	if err != nil {
		return []byte{}, errors.New("GetGroup Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("GetGroup Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("GetGroup Exception: " + err.Error())
	}
	return reply, nil
}

// CheckGroupExists takes the groupId generated during a createGroup
// and returns whether the group exists for the given groupId
// For more details see https://api.voiceit.io/#check-if-group-exists
func (vi VoiceIt2) CheckGroupExists(groupId string) ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/groups/"+groupId+"/exists"+vi.NotificationUrl, nil)
	if err != nil {
		return []byte{}, errors.New("CheckGroupExists Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("CheckGroupExists Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("CheckGroupExists Exception: " + err.Error())
	}
	return reply, nil
}

// CreateGroup creates a new group profile and returns a unique groupId
// that is used for all future calls related to the group
// For more details see https://api.voiceit.io/#create-a-group
func (vi VoiceIt2) CreateGroup(description string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("description", description); err != nil {
		return []byte{}, errors.New("CreateGroup Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/groups"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("CreateGroup Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("CreateGroup Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("CreateGroup Exception: " + err.Error())
	}
	return reply, nil
}

// AddUserToGroup takes the groupId generated during a createGroup
// and the userId generated during createUser and adds the user to the group
// For more details see https://api.voiceit.io/#add-user-to-group
func (vi VoiceIt2) AddUserToGroup(groupId, userId string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("groupId", groupId); err != nil {
		return []byte{}, errors.New("AddUserToGroup Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("AddUserToGroup Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("PUT", vi.BaseUrl+"/groups/addUser"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("AddUserToGroup Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("AddUserToGroup Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("AddUserToGroup Exception: " + err.Error())
	}
	return reply, nil
}

// RemoveUserFromGroup takes the groupId generated during a createGroup
// and the userId generated during createUser and removes the user from the group
// For more details see https://api.voiceit.io/#remove-user-from-group
func (vi VoiceIt2) RemoveUserFromGroup(groupId, userId string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("groupId", groupId); err != nil {
		return []byte{}, errors.New("RemoveUserFromGroup Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("RemoveUserFromGroup Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("PUT", vi.BaseUrl+"/groups/removeUser"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("RemoveUserFromGroup Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("RemoveUserFromGroup Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("RemoveUserFromGroup Exception: " + err.Error())
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
		return []byte{}, errors.New("DeleteGroup Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("DeleteGroup Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("DeleteGroup Exception: " + err.Error())
	}
	return reply, nil
}

// GetAllVoiceEnrollments takes the userId generated during a createUser
// and returns a list of all voice enrollments for the user
// For more details see https://api.voiceit.io/#get-voice-enrollments
func (vi VoiceIt2) GetAllVoiceEnrollments(userId string) ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/enrollments/voice/"+userId+vi.NotificationUrl, nil)
	if err != nil {
		return []byte{}, errors.New("GetAllVoiceEnrollments Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("GetAllVoiceEnrollments Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("GetAllVoiceEnrollments Exception: " + err.Error())
	}
	return reply, nil
}

// GetAllVideoEnrollments takes the userId generated during a createUser
// and returns a list of all video enrollments for the user
// For more details see https://api.voiceit.io/#get-video-enrollments
func (vi VoiceIt2) GetAllVideoEnrollments(userId string) ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/enrollments/video/"+userId+vi.NotificationUrl, nil)
	if err != nil {
		return []byte{}, errors.New("GetAllVideoEnrollments Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("GetAllVideoEnrollments Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("GetAllVideoEnrollments Exception: " + err.Error())
	}
	return reply, nil
}

// GetAllFaceEnrollments takes the userId generated during a createUser
// and returns a list of all face enrollments for the user
// For more details see https://api.voiceit.io/#get-face-enrollments
func (vi VoiceIt2) GetAllFaceEnrollments(userId string) ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/enrollments/face/"+userId+vi.NotificationUrl, nil)
	if err != nil {
		return []byte{}, errors.New("GetAllFaceEnrollments Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("GetAllFaceEnrollments Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("GetAllFaceEnrollments Exception: " + err.Error())
	}
	return reply, nil
}

// CreateVoiceEnrollment takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for an audio recording to create a voice enrollment for the user
// For more details see https://api.voiceit.io/#create-voice-enrollment
func (vi VoiceIt2) CreateVoiceEnrollment(userId, contentLanguage, phrase, filePath string) ([]byte, error) {

	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollment Exception: " + err.Error())
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("recording", path.Base(filePath))
	if err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollment Exception: " + err.Error())
	}

	if _, err := part.Write(fileContents); err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollment Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollment Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollment Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollment Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/voice"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollment Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollment Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollment Exception: " + err.Error())
	}
	return reply, nil
}

// CreateVoiceEnrollmentByByteSlice takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// file name for an audio recording to create a voice enrollment for the user
// file data in []byte form for an audio recording to create a voice enrollment for the user
func (vi VoiceIt2) CreateVoiceEnrollmentByByteSlice(userId, contentLanguage, phrase, filename string, fileData []byte) ([]byte, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("recording", filename)
	if err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollmentByByteSlice Exception: " + err.Error())
	}

	if _, err := part.Write(fileData); err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollmentByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollmentByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollmentByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollmentByByteSlice Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/voice"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollmentByByteSlice Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollmentByByteSlice Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollmentByByteSlice Exception: " + err.Error())
	}
	return reply, nil
}

// CreateVoiceEnrollmentByUrl takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to an audio recording to create a voice enrollment for the user
// For more details see https://api.voiceit.io/#create-voice-enrollment-by-url
func (vi VoiceIt2) CreateVoiceEnrollmentByUrl(userId, contentLanguage, phrase, fileUrl string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollmentByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollmentByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("fileUrl", fileUrl); err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollmentByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollmentByUrl Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/voice/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollmentByUrl Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollmentByUrl Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("CreateVoiceEnrollmentByUrl Exception: " + err.Error())
	}
	return reply, nil
}

// CreateFaceEnrollment takes the userId generated during a createUser and
// absolute file path for a video recording to create a face enrollment for the user
func (vi VoiceIt2) CreateFaceEnrollment(userId, filePath string, isPhoto ...bool) ([]byte, error) {

	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte{}, errors.New("CreateFaceEnrollment Exception: " + err.Error())
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	var fileFieldKey string
	if len(isPhoto) < 1 || !isPhoto[0] {
		fileFieldKey = "video"
	} else {
		fileFieldKey = "photo"
	}

	part, err := writer.CreateFormFile(fileFieldKey, path.Base(filePath))
	if err != nil {
		return []byte{}, errors.New("CreateFaceEnrollment Exception: " + err.Error())
	}

	if _, err := part.Write(fileContents); err != nil {
		return []byte{}, errors.New("CreateFaceEnrollment Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("CreateFaceEnrollment Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/face"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("CreateFaceEnrollment Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("CreateFaceEnrollment Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("CreateFaceEnrollment Exception: " + err.Error())
	}
	return reply, nil
}

// CreateFaceEnrollmentByByteSlice takes the userId generated during a CreateUser and
// filename for a video recording to create a face enrollment for the user
// fileData in []byte form for a video recording to create a face enrollment for the user
func (vi VoiceIt2) CreateFaceEnrollmentByByteSlice(userId, filename string, fileData []byte, isPhoto ...bool) ([]byte, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	var fileFieldKey string
	if len(isPhoto) < 1 || !isPhoto[0] {
		fileFieldKey = "video"
	} else {
		fileFieldKey = "photo"
	}

	part, err := writer.CreateFormFile(fileFieldKey, filename)
	if err != nil {
		return []byte{}, errors.New("CreateFaceEnrollmentByByteSlice Exception: " + err.Error())
	}

	if _, err := part.Write(fileData); err != nil {
		return []byte{}, errors.New("CreateFaceEnrollmentByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("CreateFaceEnrollmentByByteSlice Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/face"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("CreateFaceEnrollmentByByteSlice Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("CreateFaceEnrollment Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("CreateFaceEnrollment Exception: " + err.Error())
	}
	return reply, nil
}

// CreateFaceEnrollmentByUrl takes the userId generated during a createUser
// and a fully qualified URL to a video recording to verify the user's face
// For more details see https://api.voiceit.io/#create-face-enrollment-by-url
func (vi VoiceIt2) CreateFaceEnrollmentByUrl(userId, fileUrl string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("CreateFaceEnrollmentByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("fileUrl", fileUrl); err != nil {
		return []byte{}, errors.New("CreateFaceEnrollmentByUrl Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/face/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("CreateFaceEnrollmentByUrl Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("CreateFaceEnrollmentByUrl Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("CreateFaceEnrollmentByUrl Exception: " + err.Error())
	}
	return reply, nil
}

// CreateVideoEnrollment takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a video recording to create a video enrollment for the user
// For more details see https://api.voiceit.io/#create-video-enrollment
func (vi VoiceIt2) CreateVideoEnrollment(userId, contentLanguage, phrase, filePath string) ([]byte, error) {

	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte{}, errors.New("CreateVideoEnrollment Exception: " + err.Error())
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", path.Base(filePath))
	if err != nil {
		return []byte{}, errors.New("CreateVideoEnrollment Exception: " + err.Error())
	}

	if _, err := part.Write(fileContents); err != nil {
		return []byte{}, errors.New("CreateVideoEnrollment Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("CreateVideoEnrollment Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("CreateVideoEnrollment Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("CreateVideoEnrollment Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/video"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("CreateVideoEnrollment Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("CreateVideoEnrollment Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("CreateVideoEnrollment Exception: " + err.Error())
	}
	return reply, nil
}

// CreateVideoEnrollmentByByteSlice takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// filename for a video recording to create a video enrollment for the user
// and file data in []byte form for a video recording to create a video enrollment for the user
func (vi VoiceIt2) CreateVideoEnrollmentByByteSlice(userId, contentLanguage, phrase, filename string, fileData []byte) ([]byte, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", filename)
	if err != nil {
		return []byte{}, errors.New("CreateVideoEnrollmentByByteSlice Exception: " + err.Error())
	}

	if _, err := part.Write(fileData); err != nil {
		return []byte{}, errors.New("CreateVideoEnrollmentByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("CreateVideoEnrollmentByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("CreateVideoEnrollmentByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("CreateVideoEnrollmentByByteSlice Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/video"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("CreateVideoEnrollmentByByteSlice Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("CreateVideoEnrollmentByByteSlice Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("CreateVideoEnrollmentByByteSlice Exception: " + err.Error())
	}
	return reply, nil
}

// CreateSplitVideoEnrollment takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file paths for a photo and audio recording
// Written for VoiceIt internal projects
func (vi VoiceIt2) CreateSplitVideoEnrollment(userId, contentLanguage, phrase, audioFilePath, photoFilePath string) ([]byte, error) {

	audioFileContents, err := ioutil.ReadFile(audioFilePath)
	if err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollment Exception: " + err.Error())
	}

	photoFileContents, err := ioutil.ReadFile(photoFilePath)
	if err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollment Exception: " + err.Error())
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	audioPart, err := writer.CreateFormFile("audio", path.Base(audioFilePath))
	if err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollment Exception: " + err.Error())
	}

	if _, err := audioPart.Write(audioFileContents); err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollment Exception: " + err.Error())
	}

	photoPart, err := writer.CreateFormFile("photo", path.Base(photoFilePath))
	if err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollment Exception: " + err.Error())
	}

	if _, err := photoPart.Write(photoFileContents); err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollment Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollment Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollment Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollment Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/video"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollment Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollment Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollment Exception: " + err.Error())
	}
	return reply, nil
}

// CreateSplitVideoEnrollmentByByteSlice takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// filename for a photo and audio recording
// and file data in []byte form for a photo and audio recording
// Written for VoiceIt internal projects
func (vi VoiceIt2) CreateSplitVideoEnrollmentByByteSlice(userId, contentLanguage, phrase, audioFilename, photoFilename string, audioFileData, photoFileData []byte) ([]byte, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	audioPart, err := writer.CreateFormFile("audio", audioFilename)
	if err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollmentByByteSlice Exception: " + err.Error())
	}

	if _, err := audioPart.Write(audioFileData); err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollmentByByteSlice Exception: " + err.Error())
	}

	photoPart, err := writer.CreateFormFile("photo", photoFilename)
	if err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollmentByByteSlice Exception: " + err.Error())
	}

	if _, err := photoPart.Write(photoFileData); err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollmentByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollmentByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollmentByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollmentByByteSlice Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/video"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollmentByByteSlice Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollmentByByteSlice Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("CreateSplitVideoEnrollmentByByteSlice Exception: " + err.Error())
	}
	return reply, nil
}

// CreateVideoEnrollmentByUrl takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a video recording to create a video enrollment for the user
// For more details see https://api.voiceit.io/#create-video-enrollment-by-url
func (vi VoiceIt2) CreateVideoEnrollmentByUrl(userId, contentLanguage, phrase, fileUrl string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("CreateVideoEnrollmentByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("CreateVideoEnrollmentByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("fileUrl", fileUrl); err != nil {
		return []byte{}, errors.New("CreateVideoEnrollmentByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("CreateVideoEnrollmentByUrl Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/enrollments/video/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("CreateVideoEnrollmentByUrl Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("CreateVideoEnrollmentByUrl Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("CreateVideoEnrollmentByUrl Exception: " + err.Error())
	}
	return reply, nil
}

// DeleteAllEnrollments takes the userId generated during a createUser
// and deletes all video/voice enrollments for the user
// For more details see https://api.voiceit.io/#delete-all-enrollments-for-user
func (vi VoiceIt2) DeleteAllEnrollments(userId string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", vi.BaseUrl+"/enrollments/"+userId+"/all"+vi.NotificationUrl, nil)
	if err != nil {
		return []byte{}, errors.New("DeleteAllEnrollments Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("DeleteAllEnrollments Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("DeleteAllEnrollments Exception: " + err.Error())
	}
	return reply, nil
}

// VoiceVerification takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for an audio recording to verify the user's voice
// For more details see https://api.voiceit.io/#verify-a-user-s-voice
func (vi VoiceIt2) VoiceVerification(userId, contentLanguage, phrase, filePath string) ([]byte, error) {

	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte{}, errors.New("VoiceVerification Exception: " + err.Error())
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("recording", path.Base(filePath))
	if err != nil {
		return []byte{}, errors.New("VoiceVerification Exception: " + err.Error())
	}

	if _, err := part.Write(fileContents); err != nil {
		return []byte{}, errors.New("VoiceVerification Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("VoiceVerification Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("VoiceVerification Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("VoiceVerification Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/voice"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("VoiceVerification Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("VoiceVerification Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("VoiceVerification Exception: " + err.Error())
	}
	return reply, nil
}

// VoiceVerificationByByteSlice takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// filename for an audio recording to verify the user's voice
// and file data in []byte form for an audio recording to verify the user's voice
func (vi VoiceIt2) VoiceVerificationByByteSlice(userId, contentLanguage, phrase, filename string, fileData []byte) ([]byte, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("recording", filename)
	if err != nil {
		return []byte{}, errors.New("VoiceVerificationByByteSlice Exception: " + err.Error())
	}

	if _, err := part.Write(fileData); err != nil {
		return []byte{}, errors.New("VoiceVerificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("VoiceVerificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("VoiceVerificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("VoiceVerificationByByteSlice Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/voice"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("VoiceVerificationByByteSlice Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("VoiceVerificationByByteSlice Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("VoiceVerificationByByteSlice Exception: " + err.Error())
	}
	return reply, nil
}

// VoiceVerificationByUrl takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to an audio recording to verify the user's voice
// For more details see https://api.voiceit.io/#verify-a-user-s-voice-by-url
func (vi VoiceIt2) VoiceVerificationByUrl(userId, contentLanguage, phrase, fileUrl string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("VoiceVerificationByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("VoiceVerificationByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("fileUrl", fileUrl); err != nil {
		return []byte{}, errors.New("VoiceVerificationByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("VoiceVerificationByUrl Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/voice/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("VoiceVerificationByUrl Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("VoiceVerificationByUrl Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("VoiceVerificationByUrl Exception: " + err.Error())
	}
	return reply, nil
}

// FaceVerification takes the userId generated during a createUser and a
// absolute file path for a video recording to verify the user's face
// For more details see https://api.voiceit.io/#verify-a-user-s-face
func (vi VoiceIt2) FaceVerification(userId, filePath string, isPhoto ...bool) ([]byte, error) {

	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte{}, errors.New("FaceVerification Exception: " + err.Error())
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	var fileFieldKey string
	if len(isPhoto) < 1 || !isPhoto[0] {
		fileFieldKey = "video"
	} else {
		fileFieldKey = "photo"
	}

	part, err := writer.CreateFormFile(fileFieldKey, path.Base(filePath))
	if err != nil {
		return []byte{}, errors.New("FaceVerification Exception: " + err.Error())
	}

	if _, err := part.Write(fileContents); err != nil {
		return []byte{}, errors.New("FaceVerification Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("FaceVerification Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/face"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("FaceVerification Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("FaceVerification Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("FaceVerification Exception: " + err.Error())
	}
	return reply, nil
}

// FaceVerificationByByteSlice takes the userId generated during a createUser and a
// filename for a video recording to verify the user's face
// and file data in []byte form for a video recording to verify the user's face
func (vi VoiceIt2) FaceVerificationByByteSlice(userId, filename string, fileData []byte, isPhoto ...bool) ([]byte, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	var fileFieldKey string
	if len(isPhoto) < 1 || !isPhoto[0] {
		fileFieldKey = "video"
	} else {
		fileFieldKey = "photo"
	}

	part, err := writer.CreateFormFile(fileFieldKey, filename)
	if err != nil {
		return []byte{}, errors.New("FaceVerificationByByteSlice Exception: " + err.Error())
	}

	if _, err := part.Write(fileData); err != nil {
		return []byte{}, errors.New("FaceVerificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("FaceVerificationByteSlice Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/face"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("FaceVerificationByteSlice Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("FaceVerificationByteSlice Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("FaceVerificationByteSlice Exception: " + err.Error())
	}
	return reply, nil
}

// FaceVerificationByUrl takes the userId generated during a createUser
// and a fully qualified URL to a video recording to verify the user's face
// For more details see https://api.voiceit.io/#verify-a-user-s-face-by-url
func (vi VoiceIt2) FaceVerificationByUrl(userId, fileUrl string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("fileUrl", fileUrl); err != nil {
		return []byte{}, errors.New("FaceVerificationByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("FaceVerificationByUrl Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/face/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("FaceVerificationByUrl Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("FaceVerificationByUrl Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("FaceVerificationByUrl Exception: " + err.Error())
	}
	return reply, nil
}

// VideoVerification takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a video recording to verify the user's face and voice
// For more details see https://api.voiceit.io/#video-verification
func (vi VoiceIt2) VideoVerification(userId, contentLanguage, phrase, filePath string) ([]byte, error) {

	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte{}, errors.New("VideoVerification Exception: " + err.Error())
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", path.Base(filePath))
	if err != nil {
		return []byte{}, errors.New("VideoVerification Exception: " + err.Error())
	}

	if _, err := part.Write(fileContents); err != nil {
		return []byte{}, errors.New("VideoVerification Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("VideoVerification Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("VideoVerification Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("VideoVerification Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/video"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("VideoVerification Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("VideoVerification Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("VideoVerification Exception: " + err.Error())
	}
	return reply, nil
}

// VideoVerificationByByteSlice takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and filename for a video recording to verify the user's face and voice
// and file data in []byte form for a video recording to verify the user's face and voice
func (vi VoiceIt2) VideoVerificationByByteSlice(userId, contentLanguage, phrase, filename string, fileData []byte) ([]byte, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", filename)
	if err != nil {
		return []byte{}, errors.New("VideoVerificationByByteSlice Exception: " + err.Error())
	}

	if _, err := part.Write(fileData); err != nil {
		return []byte{}, errors.New("VideoVerificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("VideoVerificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("VideoVerificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("VideoVerificationByByteSlice Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/video"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("VideoVerificationByByteSlice Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("VideoVerificationByByteSlice Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("VideoVerificationByByteSlice Exception: " + err.Error())
	}
	return reply, nil
}

// SplitVideoVerification takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file paths for a photo and audio recording to verify the user's face and voice
// Written for VoiceIt internal projects
func (vi VoiceIt2) SplitVideoVerification(userId, contentLanguage, phrase, audioFilePath, photoFilePath string) ([]byte, error) {

	audioContents, err := ioutil.ReadFile(audioFilePath)
	if err != nil {
		return []byte{}, errors.New("SplitVideoVerification Exception: " + err.Error())
	}

	photoContents, err := ioutil.ReadFile(photoFilePath)
	if err != nil {
		return []byte{}, errors.New("SplitVideoVerification Exception: " + err.Error())
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	audioPart, err := writer.CreateFormFile("audio", path.Base(audioFilePath))
	if err != nil {
		return []byte{}, errors.New("SplitVideoVerification Exception: " + err.Error())
	}

	if _, err := audioPart.Write(audioContents); err != nil {
		return []byte{}, errors.New("SplitVideoVerification Exception: " + err.Error())
	}

	photoPart, err := writer.CreateFormFile("photo", path.Base(photoFilePath))
	if err != nil {
		return []byte{}, errors.New("SplitVideoVerification Exception: " + err.Error())
	}

	if _, err := photoPart.Write(photoContents); err != nil {
		return []byte{}, errors.New("SplitVideoVerification Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("SplitVideoVerification Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("SplitVideoVerification Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("SplitVideoVerification Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/video"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("SplitVideoVerification Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("SplitVideoVerification Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("SplitVideoVerification Exception: " + err.Error())
	}
	return reply, nil
}

// SplitVideoVerificationByByteSlice takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// file names for a photo and audio recording to verify the user's face and voice
// and file data in []byte form for a photo and audio recording to verify the user's face and voice
func (vi VoiceIt2) SplitVideoVerificationByByteSlice(userId, contentLanguage, phrase, audioFilename, photoFilename string, audioFileData, photoFileData []byte) ([]byte, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	audioPart, err := writer.CreateFormFile("audio", audioFilename)
	if err != nil {
		return []byte{}, errors.New("SplitVideoVerificationByByteSlice Exception: " + err.Error())
	}

	if _, err := audioPart.Write(audioFileData); err != nil {
		return []byte{}, errors.New("SplitVideoVerificationByByteSlice Exception: " + err.Error())
	}

	photoPart, err := writer.CreateFormFile("photo", photoFilename)
	if err != nil {
		return []byte{}, errors.New("SplitVideoVerificationByByteSlice Exception: " + err.Error())
	}

	if _, err := photoPart.Write(photoFileData); err != nil {
		return []byte{}, errors.New("SplitVideoVerificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("SplitVideoVerificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("SplitVideoVerificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("SplitVideoVerificationByByteSlice Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/video"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("SplitVideoVerificationByByteSlice Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("SplitVideoVerificationByByteSlice Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("SplitVideoVerificationByByteSlice Exception: " + err.Error())
	}
	return reply, nil
}

// VideoVerificationByUrl takes the userId generated during a createUser,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a video recording to verify the user's face and voice
// For more details see https://api.voiceit.io/#video-verification-by-url
func (vi VoiceIt2) VideoVerificationByUrl(userId, contentLanguage, phrase, fileUrl string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("userId", userId); err != nil {
		return []byte{}, errors.New("VideoVerificationByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("VideoVerificationByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("fileUrl", fileUrl); err != nil {
		return []byte{}, errors.New("VideoVerificationByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("VideoVerificationByUrl Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/verification/video/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("VideoVerificationByUrl Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("VideoVerificationByUrl Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("VideoVerificationByUrl Exception: " + err.Error())
	}
	return reply, nil
}

// VoiceIdentification takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for an audio recording to idetify the user's voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-voice
func (vi VoiceIt2) VoiceIdentification(groupId, contentLanguage, phrase, filePath string) ([]byte, error) {

	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte{}, errors.New("VoiceIdentification Exception: " + err.Error())
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("recording", path.Base(filePath))
	if err != nil {
		return []byte{}, errors.New("VoiceIdentification Exception: " + err.Error())
	}

	if _, err := part.Write(fileContents); err != nil {
		return []byte{}, errors.New("VoiceIdentification Exception: " + err.Error())
	}

	if err := writer.WriteField("groupId", groupId); err != nil {
		return []byte{}, errors.New("VoiceIdentification Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("VoiceIdentification Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("VoiceIdentification Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/voice"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("VoiceIdentification Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("VoiceIdentification Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("VoiceIdentification Exception: " + err.Error())
	}
	return reply, nil
}

// VoiceIdentificationByByteSlice takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// file name for an audio recording to idetify the user's voice
// and file data in []byte form for an audio recording to idetify the user's voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-voice
func (vi VoiceIt2) VoiceIdentificationByByteSlice(groupId, contentLanguage, phrase, filename string, fileData []byte) ([]byte, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("recording", filename)
	if err != nil {
		return []byte{}, errors.New("VoiceIdentificationByByteSlice Exception: " + err.Error())
	}

	if _, err := part.Write(fileData); err != nil {
		return []byte{}, errors.New("VoiceIdentificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("groupId", groupId); err != nil {
		return []byte{}, errors.New("VoiceIdentificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("VoiceIdentificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("VoiceIdentificationByByteSlice Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/voice"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("VoiceIdentificationByByteSlice Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("VoiceIdentificationByByteSlice Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("VoiceIdentificationByByteSlice Exception: " + err.Error())
	}
	return reply, nil
}

// VoiceIdentificationByUrl takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to an audio recording to idetify the user's voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-voice-by-url
func (vi VoiceIt2) VoiceIdentificationByUrl(groupId, contentLanguage, phrase, fileUrl string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("fileUrl", fileUrl); err != nil {
		return []byte{}, errors.New("VoiceIdentificationByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("groupId", groupId); err != nil {
		return []byte{}, errors.New("VoiceIdentificationByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("VoiceIdentificationByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("VoiceIdentificationByUrl Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/voice/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("VoiceIdentificationByUrl Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("VoiceIdentificationByUrl Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("VoiceIdentificationByUrl Exception: " + err.Error())
	}
	return reply, nil
}

// VideoIdentification takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a video recording to idetify the user's face and voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-voice-amp-face
func (vi VoiceIt2) VideoIdentification(groupId, contentLanguage, phrase, filePath string) ([]byte, error) {

	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte{}, errors.New("VideoIdentification Exception: " + err.Error())
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", path.Base(filePath))
	if err != nil {
		return []byte{}, errors.New("VideoIdentification Exception: " + err.Error())
	}

	if _, err := part.Write(fileContents); err != nil {
		return []byte{}, errors.New("VideoIdentification Exception: " + err.Error())
	}

	if err := writer.WriteField("groupId", groupId); err != nil {
		return []byte{}, errors.New("VideoIdentification Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("VideoIdentification Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("VideoIdentification Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/video"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("VideoIdentification Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("VideoIdentification Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("VideoIdentification Exception: " + err.Error())
	}
	return reply, nil
}

// VideoIdentificationByByteSlice takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// file name for a video recording to idetify the user's face and voice
// and file data in []byte form for a video recording to idetify the user's face and voice
// amongst others in the group
func (vi VoiceIt2) VideoIdentificationByByteSlice(groupId, contentLanguage, phrase, filename string, fileData []byte) ([]byte, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", filename)
	if err != nil {
		return []byte{}, errors.New("VideoIdentificationByByteSlice Exception: " + err.Error())
	}

	if _, err := part.Write(fileData); err != nil {
		return []byte{}, errors.New("VideoIdentificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("groupId", groupId); err != nil {
		return []byte{}, errors.New("VideoIdentificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("VideoIdentificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("VideoIdentificationByByteSlice Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/video"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("VideoIdentificationByByteSlice Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("VideoIdentificationByByteSlice Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("VideoIdentificationByByteSlice Exception: " + err.Error())
	}
	return reply, nil
}

// SplitVideoIdentification takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and absolute file path for a video recording to idetify the user's face and voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-voice-amp-face
func (vi VoiceIt2) SplitVideoIdentification(groupId, contentLanguage, phrase, audioFilePath, photoFilePath string) ([]byte, error) {

	audioContents, err := ioutil.ReadFile(audioFilePath)
	if err != nil {
		return []byte{}, errors.New("SplitVideoIdentification Exception: " + err.Error())
	}

	photoContents, err := ioutil.ReadFile(photoFilePath)
	if err != nil {
		return []byte{}, errors.New("SplitVideoIdentification Exception: " + err.Error())
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	audioPart, err := writer.CreateFormFile("audio", path.Base(audioFilePath))
	if err != nil {
		return []byte{}, errors.New("SplitVideoIdentification Exception: " + err.Error())
	}

	if _, err := audioPart.Write(audioContents); err != nil {
		return []byte{}, errors.New("SplitVideoIdentification Exception: " + err.Error())
	}

	photoPart, err := writer.CreateFormFile("photo", path.Base(photoFilePath))
	if err != nil {
		return []byte{}, errors.New("SplitVideoIdentification Exception: " + err.Error())
	}

	if _, err := photoPart.Write(photoContents); err != nil {
		return []byte{}, errors.New("SplitVideoIdentification Exception: " + err.Error())
	}

	if err := writer.WriteField("groupId", groupId); err != nil {
		return []byte{}, errors.New("SplitVideoIdentification Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("SplitVideoIdentification Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("SplitVideoIdentification Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/video"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("SplitVideoIdentification Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("SplitVideoIdentification Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("SplitVideoIdentification Exception: " + err.Error())
	}
	return reply, nil
}

// SplitVideoIdentificationByByteSlice takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// file name for a video recording to idetify the user's face and voice
// and file data in []byte form for a video recording to idetify the user's face and voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-voice-amp-face
func (vi VoiceIt2) SplitVideoIdentificationByByteSlice(groupId, contentLanguage, phrase, audioFilename, photoFilename string, audioFileData, photoFileData []byte) ([]byte, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	audioPart, err := writer.CreateFormFile("audio", audioFilename)
	if err != nil {
		return []byte{}, errors.New("SplitVideoIdentificationByByteSlice Exception: " + err.Error())
	}

	if _, err := audioPart.Write(audioFileData); err != nil {
		return []byte{}, errors.New("SplitVideoIdentificationByByteSlice Exception: " + err.Error())
	}

	photoPart, err := writer.CreateFormFile("photo", photoFilename)
	if err != nil {
		return []byte{}, errors.New("SplitVideoIdentificationByByteSlice Exception: " + err.Error())
	}

	if _, err := photoPart.Write(photoFileData); err != nil {
		return []byte{}, errors.New("SplitVideoIdentificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("groupId", groupId); err != nil {
		return []byte{}, errors.New("SplitVideoIdentificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("SplitVideoIdentificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("SplitVideoIdentificationByByteSlice Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/video"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("SplitVideoIdentificationByByteSlice Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("SplitVideoIdentificationByByteSlice Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("SplitVideoIdentificationByByteSlice Exception: " + err.Error())
	}
	return reply, nil
}

// VideoIdentificationByUrl takes the groupId generated during a createGroup,
// the contentLanguage(https://api.voiceit.io/#content-languages) for the phrase,
// the text of a valid phrase for the developer account,
// and a fully qualified URL to a video recording to idetify the user's face and voice
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-voice-amp-face-by-url
func (vi VoiceIt2) VideoIdentificationByUrl(groupId, contentLanguage, phrase, fileUrl string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("fileUrl", fileUrl)
	if err != nil {
		return []byte{}, errors.New("VideoIdentificationByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("groupId", groupId); err != nil {
		return []byte{}, errors.New("VideoIdentificationByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", contentLanguage); err != nil {
		return []byte{}, errors.New("VideoIdentificationByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("phrase", phrase); err != nil {
		return []byte{}, errors.New("VideoIdentificationByUrl Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/video/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("VideoIdentificationByUrl Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("VideoIdentificationByUrl Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("VideoIdentificationByUrl Exception: " + err.Error())
	}
	return reply, nil
}

// FaceIdentification takes the groupId generated during a createGroup,
// and absolute file path for a face recording to idetify the user's face
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-face
func (vi VoiceIt2) FaceIdentification(groupId, filePath string, isPhoto ...bool) ([]byte, error) {

	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte{}, errors.New("FaceIdentification Exception: " + err.Error())
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	var fileFieldKey string
	if len(isPhoto) < 1 || !isPhoto[0] {
		fileFieldKey = "video"
	} else {
		fileFieldKey = "photo"
	}

	part, err := writer.CreateFormFile(fileFieldKey, path.Base(filePath))
	if err != nil {
		return []byte{}, errors.New("FaceIdentification Exception: " + err.Error())
	}

	if _, err := part.Write(fileContents); err != nil {
		return []byte{}, errors.New("FaceIdentification Exception: " + err.Error())
	}

	if err := writer.WriteField("groupId", groupId); err != nil {
		return []byte{}, errors.New("FaceIdentification Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/face"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("FaceIdentification Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("FaceIdentification Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("FaceIdentification Exception: " + err.Error())
	}
	return reply, nil
}

// FaceIdentificationByByteSlice takes the groupId generated during a createGroup,
// file name for a face recording to idetify the user's face
// and file data in []byte form for a face recording to idetify the user's face
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-face
func (vi VoiceIt2) FaceIdentificationByByteSlice(groupId, filename string, fileData []byte, isPhoto ...bool) ([]byte, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	var fileFieldKey string
	if len(isPhoto) < 1 || !isPhoto[0] {
		fileFieldKey = "video"
	} else {
		fileFieldKey = "photo"
	}

	part, err := writer.CreateFormFile(fileFieldKey, filename)
	if err != nil {
		return []byte{}, errors.New("FaceIdentificationByByteSlice Exception: " + err.Error())
	}

	if _, err := part.Write(fileData); err != nil {
		return []byte{}, errors.New("FaceIdentificationByByteSlice Exception: " + err.Error())
	}

	if err := writer.WriteField("groupId", groupId); err != nil {
		return []byte{}, errors.New("FaceIdentificationByByteSlice Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/face"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("FaceIdentificationByByteSlice Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("FaceIdentificationByByteSlice Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("FaceIdentificationByByteSlice Exception: " + err.Error())
	}
	return reply, nil
}

// FaceIdentificationByUrl takes the groupId generated during a createGroup,
// and a fully qualified URL to a face recording to idetify the user's face
// amongst others in the group
// For more details see https://api.voiceit.io/#identify-a-user-s-face-by-url
func (vi VoiceIt2) FaceIdentificationByUrl(groupId, fileUrl string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("fileUrl", fileUrl); err != nil {
		return []byte{}, errors.New("FaceIdentificationByUrl Exception: " + err.Error())
	}

	if err := writer.WriteField("groupId", groupId); err != nil {
		return []byte{}, errors.New("FaceIdentificationByUrl Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/identification/face/byUrl"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("FaceIdentificationByUrl Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("FaceIdentificationByUrl Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("FaceIdentificationByUrl Exception: " + err.Error())
	}
	return reply, nil
}

// GetPhrases takes the contentLanguage
// For more details see https://api.voiceit.io/#get-phrases
func (vi VoiceIt2) GetPhrases(contentLanguage string) ([]byte, error) {
	req, err := http.NewRequest("GET", vi.BaseUrl+"/phrases/"+contentLanguage+vi.NotificationUrl, nil)
	if err != nil {
		return []byte{}, errors.New("GetPhrases Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("GetPhrases Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("GetPhrases Exception: " + err.Error())
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
		return []byte{}, errors.New("CreateUserToken Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("CreateUserToken Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("CreateUserToken Exception: " + err.Error())
	}
	return reply, nil
}

// ExpireUserTokens takes a userId (string).
// For more details see https://api.voiceit.io/?go#user-token-expiration
func (vi VoiceIt2) ExpireUserTokens(userId string) ([]byte, error) {
	req, err := http.NewRequest("POST", vi.BaseUrl+"/users/"+userId+"/expireTokens"+vi.NotificationUrl, nil)
	if err != nil {
		return []byte{}, errors.New("ExpireUserTokens Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("ExpireUserTokens Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("ExpireUserTokens Exception: " + err.Error())
	}
	return reply, nil
}

// CreateManagedSubAccount creates a managed sub-account.
func (vi VoiceIt2) CreateManagedSubAccount(params structs.CreateSubAccountRequest) ([]byte, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("firstName", params.FirstName); err != nil {
		return []byte{}, errors.New("CreateManagedSubAccount Exception: " + err.Error())
	}

	if err := writer.WriteField("lastName", params.LastName); err != nil {
		return []byte{}, errors.New("CreateManagedSubAccount Exception: " + err.Error())
	}

	if err := writer.WriteField("email", params.Email); err != nil {
		return []byte{}, errors.New("CreateManagedSubAccount Exception: " + err.Error())
	}

	if err := writer.WriteField("password", params.Password); err != nil {
		return []byte{}, errors.New("CreateManagedSubAccount Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", params.ContentLanguage); err != nil {
		return []byte{}, errors.New("CreateManagedSubAccount Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/subaccount/managed"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("CreateManagedSubAccount Exception: " + err.Error())
	}

	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("CreateManagedSubAccount Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("CreateManagedSubAccount Exception: " + err.Error())
	}
	return reply, nil
}

// CreateUnmanagedSubAccount creates an unmanaged sub-account.
func (vi VoiceIt2) CreateUnmanagedSubAccount(params structs.CreateSubAccountRequest) ([]byte, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("firstName", params.FirstName); err != nil {
		return []byte{}, errors.New("CreateUnmanagedSubAccount Exception: " + err.Error())
	}

	if err := writer.WriteField("lastName", params.LastName); err != nil {
		return []byte{}, errors.New("CreateUnmanagedSubAccount Exception: " + err.Error())
	}

	if err := writer.WriteField("email", params.Email); err != nil {
		return []byte{}, errors.New("CreateUnmanagedSubAccount Exception: " + err.Error())
	}

	if err := writer.WriteField("password", params.Password); err != nil {
		return []byte{}, errors.New("CreateUnmanagedSubAccount Exception: " + err.Error())
	}

	if err := writer.WriteField("contentLanguage", params.ContentLanguage); err != nil {
		return []byte{}, errors.New("CreateUnmanagedSubAccount Exception: " + err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", vi.BaseUrl+"/subaccount/unmanaged"+vi.NotificationUrl, body)
	if err != nil {
		return []byte{}, errors.New("CreateUnmanagedSubAccount Exception: " + err.Error())
	}

	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("CreateUnmanagedSubAccount Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("CreateUnmanagedSubAccount Exception: " + err.Error())
	}
	return reply, nil
}

// RegenerateSubAccountAPIToken takes a subAccountAPIKey (string).
func (vi VoiceIt2) RegenerateSubAccountAPIToken(subAccountAPIKey string) ([]byte, error) {
	req, err := http.NewRequest("POST", vi.BaseUrl+"/subaccount/"+subAccountAPIKey, nil)
	if err != nil {
		return []byte{}, errors.New("RegenerateSubAccountAPIToken Exception: " + err.Error())
	}

	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("RegenerateSubAccountAPIToken Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("RegenerateSubAccountAPIToken Exception: " + err.Error())
	}
	return reply, nil
}

// DeleteSubAccount takes a subAccountAPIKey (string).
func (vi VoiceIt2) DeleteSubAccount(subAccountAPIKey string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", vi.BaseUrl+"/subaccount/"+subAccountAPIKey, nil)
	if err != nil {
		return []byte{}, errors.New("DeleteSubAccount Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("DeleteSubAccount Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("DeleteSubAccount Exception: " + err.Error())
	}
	return reply, nil
}

// SwitchSubAccountType takes a subAccountAPIKey (string)  (
func (vi VoiceIt2) SwitchSubAccountType(subAccountAPIKey string) ([]byte, error) {
	req, err := http.NewRequest("POST", vi.BaseUrl+"/subaccount/"+subAccountAPIKey+"/switchType", nil)
	if err != nil {
		return []byte{}, errors.New("SwitchSubAccountType Exception: " + err.Error())
	}
	req.SetBasicAuth(vi.APIKey, vi.APIToken)
	req.Header.Add("platformId", PlatformId)
	req.Header.Add("platformVersion", PlatformVersion)

	client := &http.Client{
		// Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("SwitchSubAccountType Exception: " + err.Error())
	}
	defer resp.Body.Close()
	reply, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("SwitchSubAccountType Exception: " + err.Error())
	}
	return reply, nil
}
