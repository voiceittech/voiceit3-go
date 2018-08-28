package voiceit2

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/voiceittech/VoiceIt2-Go/structs"
)

func getMessage(arg string) string {
	var dat map[string]interface{}
	json.Unmarshal([]byte(arg), &dat)
	return dat["status"].(string)
}

func getStatus(arg string) int {
	var dat map[string]interface{}
	json.Unmarshal([]byte(arg), &dat)
	return int(dat["status"].(float64))
}

func getResponseCode(arg string) string {
	var dat map[string]interface{}
	json.Unmarshal([]byte(arg), &dat)
	return dat["responseCode"].(string)
}

func getUserId(arg string) string {
	var dat map[string]interface{}
	json.Unmarshal([]byte(arg), &dat)
	return dat["userId"].(string)
}

func getGroupId(arg string) string {
	var dat map[string]interface{}
	json.Unmarshal([]byte(arg), &dat)
	return dat["groupId"].(string)
}

func getEnrollmentId(arg string) int {
	var dat map[string]interface{}
	json.Unmarshal([]byte(arg), &dat)
	return int(dat["id"].(float64))
}

func getFaceEnrollmentId(arg string) int {
	var dat map[string]interface{}
	json.Unmarshal([]byte(arg), &dat)
	return int(dat["faceEnrollmentId"].(float64))
}

func TestIO(t *testing.T) {
	assert := assert.New(t)
	myVoiceIt := NewClient(os.Getenv("VIAPIKEY"), os.Getenv("VIAPITOKEN"))
	_, err := myVoiceIt.CreateVoiceEnrollment("", "", "", "not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to CreateVoiceEnrollmentFunction (should return real error)")
	_, err = myVoiceIt.CreateVideoEnrollment("", "", "", "not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to CreateVideoEnrollment (should return real error)")
	_, err = myVoiceIt.CreateFaceEnrollment("", "not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to CreateFaceEnrollment (should return real error)")
	_, err = myVoiceIt.VoiceVerification("", "", "", "not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to VoiceVerification(should return real error)")
	_, err = myVoiceIt.VideoVerification("", "", "", "not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to VideoVerification(should return real error)")
	_, err = myVoiceIt.FaceVerification("", "not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to VideoVerification(should return real error)")
	_, err = myVoiceIt.VoiceIdentification("", "en-US", "", "not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to VoiceIdentification(should return real error)")
	_, err = myVoiceIt.VideoIdentification("", "", "", "not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to VideoIdentification(should return real error)")
	_, err = myVoiceIt.FaceIdentification("", "not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to FaceIdentification(should return real error)")
}

func TestBasics(t *testing.T) {
	assert := assert.New(t)
	myVoiceIt := NewClient(os.Getenv("VIAPIKEY"), os.Getenv("VIAPITOKEN"))
	ret := myVoiceIt.CreateUser()
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)
	userId := getUserId(ret)

	ret = myVoiceIt.GetAllUsers()
	var gau structs.GetAllUsersReturn
	json.Unmarshal([]byte(ret), &gau)
	assert.Equal(200, gau.Status, ret)
	assert.Equal("SUCC", gau.ResponseCode, ret)
	assert.True(0 < len(gau.Users), ret)

	ret = myVoiceIt.CheckUserExists(userId)
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.CreateGroup("Sample Group Description")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)
	groupId := getGroupId(ret)

	ret = myVoiceIt.CheckGroupExists(groupId)
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.GetGroup(groupId)
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.AddUserToGroup(groupId, userId)
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.GetGroupsForUser(userId)
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)
	var ggfu structs.GetGroupsForUserReturn
	json.Unmarshal([]byte(ret), &ggfu)
	assert.Equal(1, len(ggfu.Groups), ret)

	ret = myVoiceIt.RemoveUserFromGroup(groupId, userId)
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.DeleteUser(userId)
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.DeleteGroup(groupId)
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.GetPhrases("en-US")
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)
}

// Helper function to download files to disk
func downloadFromUrl(url string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]

	log.Println("Downloading " + url + "...")
	output, err := os.Create(fileName)
	if err != nil {
		log.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		return
	}

}

func TestVideo(t *testing.T) {
	assert := assert.New(t)
	apikey := os.Getenv("VIAPIKEY")
	apitoken := os.Getenv("VIAPITOKEN")
	myVoiceIt := NewClient(apikey, apitoken)
	ret := myVoiceIt.CreateUser()
	userId1 := getUserId(ret)
	ret = myVoiceIt.CreateUser()
	userId2 := getUserId(ret)
	ret = myVoiceIt.CreateGroup("Sample Group Description")
	groupId := getGroupId(ret)
	myVoiceIt.AddUserToGroup(groupId, userId1)
	myVoiceIt.AddUserToGroup(groupId, userId2)

	// Video Enrollments
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentArmaan1.mov")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentArmaan2.mov")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentArmaan3.mov")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoVerificationArmaan1.mov")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentStephen1.mov")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentStephen2.mov")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentStephen3.mov")

	ret, err := myVoiceIt.CreateVideoEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentArmaan1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)
	enrollmentId := getEnrollmentId(ret)

	ret, err = myVoiceIt.CreateVideoEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentArmaan2.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret, err = myVoiceIt.CreateVideoEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentArmaan3.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret, err = myVoiceIt.CreateVideoEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentStephen1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret, err = myVoiceIt.CreateVideoEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentStephen2.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret, err = myVoiceIt.CreateVideoEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentStephen3.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	// Get Video Enrollment
	ret = myVoiceIt.GetVideoEnrollments(userId1)
	var gve1 structs.GetVideoEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve1)
	assert.Equal(200, gve1.Status, ret)
	assert.Equal("SUCC", gve1.ResponseCode, ret)
	assert.Equal(3, len(gve1.VideoEnrollments), ret)

	// Video Verification
	ret, err = myVoiceIt.VideoVerification(userId1, "en-US", "never forget tomorrow is a new day", "./videoVerificationArmaan1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	// Video Identification
	ret, err = myVoiceIt.VideoIdentification(groupId, "en-US", "never forget tomorrow is a new day", "./videoVerificationArmaan1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)
	assert.Equal(userId1, getUserId(ret), ret)

	// Delete Video Enrollment
	ret = myVoiceIt.DeleteVideoEnrollment(userId1, enrollmentId)
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.GetVideoEnrollments(userId1)
	var gve2 structs.GetVideoEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve2)
	for _, videoEnrollment := range gve2.VideoEnrollments {
		assert.NotEqual(enrollmentId, videoEnrollment.VideoEnrollmentId, ret)
	}

	// Delete All Video Enrollments
	ret = myVoiceIt.DeleteAllVideoEnrollments(userId1)
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	var gve3 structs.GetVideoEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve3)
	assert.Equal(0, len(gve3.VideoEnrollments), ret)

	// Delete All Enrollments
	ret = myVoiceIt.DeleteAllEnrollmentsForUser(userId2)
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.GetVideoEnrollments(userId2)
	var gve4 structs.GetVideoEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve4)
	assert.Equal(0, len(gve4.VideoEnrollments), ret)

	// Reset for ...ByUrl calls
	myVoiceIt.DeleteUser(userId1)
	myVoiceIt.DeleteUser(userId2)
	myVoiceIt.DeleteGroup(groupId)
	ret = myVoiceIt.CreateUser()
	userId1 = getUserId(ret)
	ret = myVoiceIt.CreateUser()
	userId2 = getUserId(ret)
	ret = myVoiceIt.CreateGroup("Sample Group Description")
	groupId = getGroupId(ret)
	myVoiceIt.AddUserToGroup(groupId, userId1)
	myVoiceIt.AddUserToGroup(groupId, userId2)

	// Video Enrollments By Url
	ret = myVoiceIt.CreateVideoEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentArmaan1.mov")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.CreateVideoEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentArmaan2.mov")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.CreateVideoEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentArmaan3.mov")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.CreateVideoEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentStephen1.mov")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.CreateVideoEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentStephen2.mov")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.CreateVideoEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentStephen3.mov")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	// Video Verification
	ret = myVoiceIt.VideoVerificationByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoVerificationArmaan1.mov")
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	// Video Identification
	ret = myVoiceIt.VideoIdentificationByUrl(groupId, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoVerificationArmaan1.mov")
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)
	assert.Equal(userId1, getUserId(ret), ret)

	myVoiceIt.DeleteAllEnrollmentsForUser(userId1)
	myVoiceIt.DeleteAllEnrollmentsForUser(userId2)
	myVoiceIt.DeleteUser(userId1)
	myVoiceIt.DeleteUser(userId2)
	myVoiceIt.DeleteGroup(groupId)

	os.Remove("./videoEnrollmentArmaan1.mov")
	os.Remove("./videoEnrollmentArmaan2.mov")
	os.Remove("./videoEnrollmentArmaan3.mov")
	os.Remove("./videoVerificationArmaan1.mov")
	os.Remove("./videoEnrollmentStephen1.mov")
	os.Remove("./videoEnrollmentStephen2.mov")
	os.Remove("./videoEnrollmentStephen3.mov")
}

func TestVoice(t *testing.T) {
	assert := assert.New(t)
	apikey := os.Getenv("VIAPIKEY")
	apitoken := os.Getenv("VIAPITOKEN")
	myVoiceIt := NewClient(apikey, apitoken)
	ret := myVoiceIt.CreateUser()
	userId1 := getUserId(ret)
	ret = myVoiceIt.CreateUser()
	userId2 := getUserId(ret)
	ret = myVoiceIt.CreateGroup("Sample Group Description")
	groupId := getGroupId(ret)
	myVoiceIt.AddUserToGroup(groupId, userId1)
	myVoiceIt.AddUserToGroup(groupId, userId2)

	// Voice Enrollments
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentArmaan1.wav")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentArmaan2.wav")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentArmaan3.wav")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/verificationArmaan1.wav")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentStephen1.wav")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentStephen2.wav")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentStephen3.wav")

	ret, err := myVoiceIt.CreateVoiceEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./enrollmentArmaan1.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)
	enrollmentId := getEnrollmentId(ret)

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./enrollmentArmaan2.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./enrollmentArmaan3.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./enrollmentStephen1.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./enrollmentStephen2.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./enrollmentStephen3.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	// Get Voice Enrollment
	ret = myVoiceIt.GetVoiceEnrollments(userId1)
	var gve1 structs.GetVoiceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve1)
	assert.Equal(200, gve1.Status, ret)
	assert.Equal("SUCC", gve1.ResponseCode, ret)
	assert.Equal(3, len(gve1.VoiceEnrollments), ret)

	// Voice Verification
	ret, err = myVoiceIt.VoiceVerification(userId1, "en-US", "never forget tomorrow is a new day", "./verificationArmaan1.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	// Voice Identification
	ret, err = myVoiceIt.VoiceIdentification(groupId, "en-US", "never forget tomorrow is a new day", "./verificationArmaan1.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)
	assert.Equal(userId1, getUserId(ret), ret)

	// Delete Voice Enrollment
	ret = myVoiceIt.DeleteVoiceEnrollment(userId1, enrollmentId)
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.GetVoiceEnrollments(userId1)
	var gve2 structs.GetVoiceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve2)
	for _, videoEnrollment := range gve2.VoiceEnrollments {
		assert.NotEqual(enrollmentId, videoEnrollment.VoiceEnrollmentId, ret)
	}

	// Delete All Voice Enrollments
	ret = myVoiceIt.DeleteAllVoiceEnrollments(userId1)
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	var gve3 structs.GetVoiceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve3)
	assert.Equal(0, len(gve3.VoiceEnrollments), ret)

	// Delete All Enrollments
	ret = myVoiceIt.DeleteAllEnrollmentsForUser(userId2)
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	// Reset for ...ByUrl calls
	myVoiceIt.DeleteUser(userId1)
	myVoiceIt.DeleteUser(userId2)
	myVoiceIt.DeleteGroup(groupId)
	ret = myVoiceIt.CreateUser()
	userId1 = getUserId(ret)
	ret = myVoiceIt.CreateUser()
	userId2 = getUserId(ret)
	ret = myVoiceIt.CreateGroup("Sample Group Description")
	groupId = getGroupId(ret)
	myVoiceIt.AddUserToGroup(groupId, userId1)
	myVoiceIt.AddUserToGroup(groupId, userId2)

	// Voice Enrollments By Url
	ret = myVoiceIt.CreateVoiceEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentArmaan1.wav")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.CreateVoiceEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentArmaan2.wav")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.CreateVoiceEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentArmaan3.wav")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.CreateVoiceEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentStephen1.wav")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.CreateVoiceEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentStephen2.wav")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.CreateVoiceEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentStephen3.wav")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	// Voice Verification
	ret = myVoiceIt.VoiceVerificationByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/verificationArmaan1.wav")
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	// Voice Identification
	ret = myVoiceIt.VoiceIdentificationByUrl(groupId, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/verificationArmaan1.wav")
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)
	assert.Equal(userId1, getUserId(ret), ret)

	myVoiceIt.DeleteAllEnrollmentsForUser(userId1)
	myVoiceIt.DeleteAllEnrollmentsForUser(userId2)
	myVoiceIt.DeleteUser(userId1)
	myVoiceIt.DeleteUser(userId2)
	myVoiceIt.DeleteGroup(groupId)

	os.Remove("./enrollmentArmaan1.wav")
	os.Remove("./enrollmentArmaan2.wav")
	os.Remove("./enrollmentArmaan3.wav")
	os.Remove("./verificationArmaan1.wav")
	os.Remove("./enrollmentStephen1.wav")
	os.Remove("./enrollmentStephen2.wav")
	os.Remove("./enrollmentStephen3.wav")
}

func TestFace(t *testing.T) {
	assert := assert.New(t)
	apikey := os.Getenv("VIAPIKEY")
	apitoken := os.Getenv("VIAPITOKEN")
	myVoiceIt := NewClient(apikey, apitoken)
	ret := myVoiceIt.CreateUser()
	userId1 := getUserId(ret)
	ret = myVoiceIt.CreateUser()
	userId2 := getUserId(ret)
	ret = myVoiceIt.CreateGroup("Sample Group Description")
	groupId := getGroupId(ret)
	myVoiceIt.AddUserToGroup(groupId, userId1)
	myVoiceIt.AddUserToGroup(groupId, userId2)

	// Face Enrollments
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceEnrollmentArmaan1.mp4")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceEnrollmentArmaan2.mp4")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceEnrollmentArmaan3.mp4")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceVerificationArmaan1.mp4")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentStephen1.mov")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentStephen2.mov")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentStephen3.mov")

	ret, err := myVoiceIt.CreateFaceEnrollment(userId1, "./faceEnrollmentArmaan1.mp4")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)
	faceEnrollmentId := getFaceEnrollmentId(ret)

	ret, err = myVoiceIt.CreateFaceEnrollment(userId1, "./faceEnrollmentArmaan2.mp4")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret, err = myVoiceIt.CreateFaceEnrollment(userId1, "./faceEnrollmentArmaan3.mp4")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret, err = myVoiceIt.CreateFaceEnrollment(userId2, "./videoEnrollmentStephen1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret, err = myVoiceIt.CreateFaceEnrollment(userId2, "./videoEnrollmentStephen1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret, err = myVoiceIt.CreateFaceEnrollment(userId2, "./videoEnrollmentStephen1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	// Get Face Enrollment
	ret = myVoiceIt.GetFaceEnrollments(userId1)
	var fve1 structs.GetFaceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &fve1)
	assert.Equal(200, fve1.Status, ret)
	assert.Equal("SUCC", fve1.ResponseCode, ret)
	assert.Equal(3, len(fve1.FaceEnrollments), ret)

	// Face Verification
	ret, err = myVoiceIt.FaceVerification(userId1, "./faceVerificationArmaan1.mp4")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	// Face Identification
	ret, err = myVoiceIt.FaceIdentification(groupId, "./faceVerificationArmaan1.mp4")
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)
	assert.Equal(userId1, getUserId(ret), ret)

	ret = myVoiceIt.GetFaceEnrollments(userId1)

	// Delete Face Enrollment
	ret = myVoiceIt.DeleteFaceEnrollment(userId1, faceEnrollmentId)
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.GetFaceEnrollments(userId1)
	var fve2 structs.GetFaceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &fve2)

	for _, faceEnrollment := range fve2.FaceEnrollments {
		assert.NotEqual(faceEnrollmentId, faceEnrollment.FaceEnrollmentId, ret)
	}

	// Delete All Face Enrollments
	ret = myVoiceIt.DeleteAllFaceEnrollments(userId1)
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	var fve3 structs.GetFaceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &fve3)
	assert.Equal(0, len(fve3.FaceEnrollments), ret)

	// Delete All Enrollments
	ret = myVoiceIt.DeleteAllEnrollmentsForUser(userId2)
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	// Reset for ...ByUrl calls
	myVoiceIt.DeleteUser(userId1)
	myVoiceIt.DeleteUser(userId2)
	myVoiceIt.DeleteGroup(groupId)
	ret = myVoiceIt.CreateUser()
	userId1 = getUserId(ret)
	ret = myVoiceIt.CreateUser()
	userId2 = getUserId(ret)
	ret = myVoiceIt.CreateGroup("Sample Group Description")
	groupId = getGroupId(ret)
	myVoiceIt.AddUserToGroup(groupId, userId1)
	myVoiceIt.AddUserToGroup(groupId, userId2)

	// Face Enrollments By Url
	ret = myVoiceIt.CreateFaceEnrollmentByUrl(userId1, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceEnrollmentArmaan1.mp4")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.CreateFaceEnrollmentByUrl(userId1, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceEnrollmentArmaan2.mp4")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.CreateFaceEnrollmentByUrl(userId1, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceEnrollmentArmaan3.mp4")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.CreateFaceEnrollmentByUrl(userId2, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentStephen1.mov")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.CreateFaceEnrollmentByUrl(userId2, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentStephen2.mov")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	ret = myVoiceIt.CreateFaceEnrollmentByUrl(userId2, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentStephen3.mov")
	assert.Equal(201, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	// Face Verification
	ret = myVoiceIt.FaceVerificationByUrl(userId1, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceVerificationArmaan1.mp4")
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)

	// Face Identification
	ret = myVoiceIt.FaceIdentificationByUrl(groupId, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceVerificationArmaan1.mp4")
	assert.Equal(200, getStatus(ret), ret)
	assert.Equal("SUCC", getResponseCode(ret), ret)
	assert.Equal(userId1, getUserId(ret), ret)

	myVoiceIt.DeleteAllEnrollmentsForUser(userId1)
	myVoiceIt.DeleteAllEnrollmentsForUser(userId2)
	myVoiceIt.DeleteUser(userId1)
	myVoiceIt.DeleteUser(userId2)
	myVoiceIt.DeleteGroup(groupId)

	os.Remove("./faceEnrollmentArmaan1.mp4")
	os.Remove("./faceEnrollmentArmaan2.mp4")
	os.Remove("./faceEnrollmentArmaan3.mp4")
	os.Remove("./faceVerificationArmaan1.mp4")
	os.Remove("./videoEnrollmentStephen1.mov")
	os.Remove("./videoEnrollmentStephen2.mov")
	os.Remove("./videoEnrollmentStephen3.mov")
}
