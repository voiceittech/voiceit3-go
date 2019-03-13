package voiceit2

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/voiceittech/VoiceIt2-Go/structs"
)

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

func TestIO(t *testing.T) {
	if os.Getenv("BOXFUSE_ENV") == "voiceittest" {
		writefileerr := ioutil.WriteFile(os.Getenv("HOME")+"/platformVersion", []byte(PlatformVersion), 0644)
		if writefileerr != nil {
			panic(writefileerr)
		}
	}

	assert := assert.New(t)
	myVoiceIt := &VoiceIt2{ApiKey: os.Getenv("VIAPIKEY"), ApiToken: os.Getenv("VIAPITOKEN"), BaseUrl: "https://api.voiceit.io"}
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
	myVoiceIt := &VoiceIt2{ApiKey: os.Getenv("VIAPIKEY"), ApiToken: os.Getenv("VIAPITOKEN"), BaseUrl: "https://api.voiceit.io"}

	ret := myVoiceIt.CreateUser()
	var cu structs.CreateUserReturn
	json.Unmarshal([]byte(ret), &cu)
	assert.Equal(201, cu.Status, "CreateUser() message: "+cu.Message)
	assert.Equal("SUCC", cu.ResponseCode, "CreateUser() message: "+cu.Message)
	userId := getUserId(ret)

	ret = myVoiceIt.GetAllUsers()
	var gau structs.GetAllUsersReturn
	json.Unmarshal([]byte(ret), &gau)
	assert.Equal(200, gau.Status, "GetAllUsers() message: "+ret)
	assert.Equal("SUCC", gau.ResponseCode, "GetAllUsers() message: "+ret)
	assert.True(0 < len(gau.Users), "GetAllUsers() message: "+ret)

	ret = myVoiceIt.CheckUserExists(userId)
	var cue structs.CheckUserExistsReturn
	json.Unmarshal([]byte(ret), &cue)
	assert.Equal(200, cue.Status, "CheckUserExists() message: "+cue.Message)
	assert.Equal("SUCC", cue.ResponseCode, "CheckUserExists() message: "+cue.Message)

	ret = myVoiceIt.CreateGroup("Sample Group Description")
	var cg structs.CreateGroupReturn
	json.Unmarshal([]byte(ret), &cg)
	assert.Equal(201, cg.Status, "CreateGroup() message: "+cg.Message)
	assert.Equal("SUCC", cg.ResponseCode, "CreateGroup() message: "+cg.Message)
	groupId := cg.GroupId

	ret = myVoiceIt.CheckGroupExists(groupId)
	var cge structs.CheckGroupExistsReturn
	json.Unmarshal([]byte(ret), &cge)
	assert.Equal(200, cge.Status, "CheckGroupExists() message: "+cge.Message)
	assert.Equal("SUCC", cge.ResponseCode, "CheckGroupExists() message: "+cge.Message)

	ret = myVoiceIt.AddUserToGroup(groupId, userId)
	var autg structs.AddUserToGroupReturn
	json.Unmarshal([]byte(ret), &autg)
	assert.Equal(200, autg.Status, "AddUserToGroup() message: "+autg.Message)
	assert.Equal("SUCC", autg.ResponseCode, "AddUserToGroup() message: "+autg.Message)

	ret = myVoiceIt.GetGroup(groupId)
	var gg structs.GetGroupsForUserReturn
	json.Unmarshal([]byte(ret), &gg)
	assert.Equal(200, gg.Status, "GetGroup() message: "+gg.Message)
	assert.Equal("SUCC", gg.ResponseCode, "GetGroup() message: "+gg.Message)

	ret = myVoiceIt.GetAllGroups()
	var gag structs.GetAllGroupsReturn
	json.Unmarshal([]byte(ret), &gag)
	assert.Equal(200, gag.Status, "GetAllGroups() message: "+gag.Message)
	assert.Equal("SUCC", gag.ResponseCode, "GetAllGroups() message: "+gag.Message)

	ret = myVoiceIt.GetGroupsForUser(userId)
	var ggfu structs.GetGroupsForUserReturn
	json.Unmarshal([]byte(ret), &ggfu)
	assert.Equal(200, ggfu.Status, "GetGroupsForUser() message: "+ggfu.Message)
	assert.Equal("SUCC", ggfu.ResponseCode, "GetGroupsForUser() message: "+ggfu.Message)
	assert.Equal(1, len(ggfu.Groups), "GetGroupsForUser() message: "+ggfu.Message)
	assert.Equal(200, ggfu.Status, "GetGroupsForUser() message: "+ggfu.Message)

	ret = myVoiceIt.RemoveUserFromGroup(groupId, userId)
	var rufg structs.RemoveUserFromGroupReturn
	json.Unmarshal([]byte(ret), &rufg)
	assert.Equal(200, rufg.Status, "RemoveUserFromGroup() message: "+rufg.Message)
	assert.Equal("SUCC", rufg.ResponseCode, "RemoveUserFromGroup() message: "+rufg.Message)

	ret = myVoiceIt.CreateUserToken(userId, 3*time.Second)
	var cut structs.CreateUserTokenReturn
	json.Unmarshal([]byte(ret), &cut)
	assert.Equal(201, cut.Status, "CreateUserToken() message: "+cut.Message)
	assert.Equal("SUCC", cut.ResponseCode, "CreateUserToken() message: "+cut.Message)

	ret = myVoiceIt.DeleteUser(userId)
	var du structs.DeleteUserReturn
	json.Unmarshal([]byte(ret), &du)
	assert.Equal(200, du.Status, "DeleteUser() message: "+du.Message)
	assert.Equal("SUCC", du.ResponseCode, "DeleteUser() message: "+du.Message)

	ret = myVoiceIt.DeleteGroup(groupId)
	var dg structs.DeleteGroupReturn
	json.Unmarshal([]byte(ret), &dg)
	assert.Equal(200, dg.Status, "DeleteGroup() message: "+dg.Message)
	assert.Equal("SUCC", dg.ResponseCode, "DeleteGroup() message: "+dg.Message)

	ret = myVoiceIt.GetPhrases("en-US")
	var gp structs.GetPhrasesReturn
	json.Unmarshal([]byte(ret), &gp)
	assert.Equal(200, gp.Status, "GetPhrases() message: "+gp.Message)
	assert.Equal("SUCC", gp.ResponseCode, "GetPhrases() message: "+gp.Message)

	myVoiceIt.AddNotificationUrl("https://voiceit.io")
	assert.Equal("?notificationURL=https%3A%2F%2Fvoiceit.io", myVoiceIt.NotificationUrl, nil)
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
	myVoiceIt := &VoiceIt2{ApiKey: os.Getenv("VIAPIKEY"), ApiToken: os.Getenv("VIAPITOKEN"), BaseUrl: "https://api.voiceit.io"}
	ret := myVoiceIt.CreateUser()
	userId1 := getUserId(ret)
	ret = myVoiceIt.CreateUser()
	userId2 := getUserId(ret)
	ret = myVoiceIt.CreateGroup("Sample Group Description")
	groupId := getGroupId(ret)
	myVoiceIt.AddUserToGroup(groupId, userId1)
	myVoiceIt.AddUserToGroup(groupId, userId2)

	// Video Enrollments
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentB1.mov")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentB2.mov")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentB3.mov")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoVerificationB1.mov")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentC1.mov")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentC2.mov")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentC3.mov")

	ret, err := myVoiceIt.CreateVideoEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentB1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve1 structs.CreateVideoEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve1)
	assert.Equal(201, cve1.Status, "CreateVideoEnrollment() message: "+cve1.Message)
	assert.Equal("SUCC", cve1.ResponseCode, "CreateVideoEnrollment() message: "+cve1.Message)
	enrollmentId := cve1.Id

	ret, err = myVoiceIt.CreateVideoEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentB2.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve2 structs.CreateVideoEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve2)
	assert.Equal(201, cve2.Status, "CreateVideoEnrollment() message: "+cve2.Message)
	assert.Equal("SUCC", cve2.ResponseCode, "CreateVideoEnrollment() message: "+cve2.Message)

	ret, err = myVoiceIt.CreateVideoEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentB3.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve3 structs.CreateVideoEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve3)
	assert.Equal(201, cve3.Status, "CreateVideoEnrollment() message: "+cve3.Message)
	assert.Equal("SUCC", cve3.ResponseCode, "CreateVideoEnrollment() message: "+cve3.Message)

	ret, err = myVoiceIt.CreateVideoEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentC1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve4 structs.CreateVideoEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve4)
	assert.Equal(201, cve4.Status, "CreateVideoEnrollment() message: "+cve4.Message)
	assert.Equal("SUCC", cve4.ResponseCode, "CreateVideoEnrollment() message: "+cve4.Message)

	ret, err = myVoiceIt.CreateVideoEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentC2.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve5 structs.CreateVideoEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve5)
	assert.Equal(201, cve5.Status, "CreateVideoEnrollment() message: "+cve5.Message)
	assert.Equal("SUCC", cve5.ResponseCode, "CreateVideoEnrollment() message: "+cve5.Message)

	ret, err = myVoiceIt.CreateVideoEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentC3.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve6 structs.CreateVideoEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve6)
	assert.Equal(201, cve6.Status, "CreateVideoEnrollment() message: "+cve6.Message)
	assert.Equal("SUCC", cve6.ResponseCode, "CreateVideoEnrollment() message: "+cve6.Message)

	// Get Video Enrollment
	ret = myVoiceIt.GetAllVideoEnrollments(userId1)
	var gve1 structs.GetAllVideoEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve1)
	assert.Equal(200, gve1.Status, "GetAllVideoEnrollments() message: "+ret)
	assert.Equal("SUCC", gve1.ResponseCode, "GetAllVideoEnrollments() message: "+ret)
	assert.Equal(3, len(gve1.VideoEnrollments), "GetAllVideoEnrollments() message: "+ret)

	// Video Verification
	ret, err = myVoiceIt.VideoVerification(userId1, "en-US", "never forget tomorrow is a new day", "./videoVerificationB1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var vv1 structs.VideoVerificationReturn
	json.Unmarshal([]byte(ret), &vv1)
	assert.Equal(200, vv1.Status, "VideoVerification() message: "+vv1.Message)
	assert.Equal("SUCC", vv1.ResponseCode, "VideoVerification() message: "+vv1.Message)

	// Video Identification
	ret, err = myVoiceIt.VideoIdentification(groupId, "en-US", "never forget tomorrow is a new day", "./videoVerificationB1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var vi1 structs.VideoIdentificationReturn
	json.Unmarshal([]byte(ret), &vi1)
	assert.Equal(200, vi1.Status, "VideoIdentification() message: "+vi1.Message)
	assert.Equal("SUCC", vi1.ResponseCode, "VideoIdentification() message: "+vi1.Message)
	assert.Equal(userId1, vi1.UserId, "VideoIdentification() message: "+vi1.Message)

	// Delete Video Enrollment
	ret = myVoiceIt.DeleteVideoEnrollment(userId1, enrollmentId)
	var dve structs.DeleteVideoEnrollmentReturn
	json.Unmarshal([]byte(ret), &dve)
	assert.Equal(200, dve.Status, "DeleteVideoEnrollment() message: "+dve.Message)
	assert.Equal("SUCC", dve.ResponseCode, "DeleteVideoEnrollment() message: "+dve.Message)

	ret = myVoiceIt.GetAllVideoEnrollments(userId1)
	var gve2 structs.GetAllVideoEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve2)
	for _, videoEnrollment := range gve2.VideoEnrollments {
		assert.NotEqual(enrollmentId, videoEnrollment.VideoEnrollmentId, "GetAllVideoEnrollments() message: "+gve2.Message)
	}

	// Delete All Video Enrollments
	ret = myVoiceIt.DeleteAllVideoEnrollments(userId1)
	var dave structs.DeleteAllVideoEnrollmentsReturn
	json.Unmarshal([]byte(ret), &dave)
	assert.Equal(200, dave.Status, "DeleteAllVideoEnrollments() message: "+dave.Message)
	assert.Equal("SUCC", dave.ResponseCode, "DeleteAllVideoEnrollments() message: "+dave.Message)

	var gve3 structs.GetAllVideoEnrollmentsReturn
	ret = myVoiceIt.GetAllVideoEnrollments(userId1)
	json.Unmarshal([]byte(ret), &gve3)
	assert.Equal(200, gve3.Status, "GetAllVideoEnrollments() message: "+gve3.Message)
	assert.Equal("SUCC", gve3.ResponseCode, "GetAllVideoEnrollments() message: "+gve3.Message)
	assert.Equal(0, len(gve3.VideoEnrollments), "GetAllVideoEnrollments() message: "+gve3.Message)

	// Delete All Enrollments
	ret = myVoiceIt.DeleteAllEnrollments(userId2)
	var dae structs.DeleteAllEnrollmentsReturn
	json.Unmarshal([]byte(ret), &dae)
	assert.Equal(200, dae.Status, "DeleteAllEnrollments() message: "+dae.Message)
	assert.Equal("SUCC", dae.ResponseCode, "DeleteAllEnrollments() message: "+dae.Message)

	ret = myVoiceIt.GetAllVideoEnrollments(userId2)
	var gve4 structs.GetAllVideoEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve4)
	assert.Equal(200, gve4.Status, "GetAllVideoEnrollments() message: "+gve4.Message)
	assert.Equal("SUCC", gve4.ResponseCode, "GetAllVideoEnrollments() message: "+gve4.Message)
	assert.Equal(0, len(gve4.VideoEnrollments), "GetAllVideoEnrollments() message: "+gve4.Message)

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
	ret = myVoiceIt.CreateVideoEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentB1.mov")
	var cvebu1 structs.CreateVideoEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu1)
	assert.Equal(201, cvebu1.Status, "CreateVideoEnrollmentByUrl() message: "+cvebu1.Message)
	assert.Equal("SUCC", cvebu1.ResponseCode, "CreateVideoEnrollmentByUrl() message: "+cvebu1.Message)

	ret = myVoiceIt.CreateVideoEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentB2.mov")
	var cvebu2 structs.CreateVideoEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu2)
	assert.Equal(201, cvebu2.Status, "CreateVideoEnrollmentByUrl() message: "+cvebu2.Message)
	assert.Equal("SUCC", cvebu2.ResponseCode, "CreateVideoEnrollmentByUrl() message: "+cvebu2.Message)

	ret = myVoiceIt.CreateVideoEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentB3.mov")
	var cvebu3 structs.CreateVideoEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu3)
	assert.Equal(201, cvebu3.Status, "CreateVideoEnrollmentByUrl() message: "+cvebu3.Message)
	assert.Equal("SUCC", cvebu3.ResponseCode, "CreateVideoEnrollmentByUrl() message: "+cvebu3.Message)

	ret = myVoiceIt.CreateVideoEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentC1.mov")
	var cvebu4 structs.CreateVideoEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu4)
	assert.Equal(201, cvebu4.Status, "CreateVideoEnrollmentByUrl() message: "+cvebu4.Message)
	assert.Equal("SUCC", cvebu4.ResponseCode, "CreateVideoEnrollmentByUrl() message: "+cvebu4.Message)

	ret = myVoiceIt.CreateVideoEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentC2.mov")
	var cvebu5 structs.CreateVideoEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu5)
	assert.Equal(201, cvebu5.Status, "CreateVideoEnrollmentByUrl() message: "+cvebu5.Message)
	assert.Equal("SUCC", cvebu5.ResponseCode, "CreateVideoEnrollmentByUrl() message: "+cvebu5.Message)

	ret = myVoiceIt.CreateVideoEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentC3.mov")
	var cvebu6 structs.CreateVideoEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu6)
	assert.Equal(201, cvebu6.Status, "CreateVideoEnrollmentByUrl() message: "+cvebu6.Message)
	assert.Equal("SUCC", cvebu6.ResponseCode, "CreateVideoEnrollmentByUrl() message: "+cvebu6.Message)

	// Video Verification
	ret = myVoiceIt.VideoVerificationByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoVerificationB1.mov")
	var vvbu structs.VideoVerificationByUrlReturn
	json.Unmarshal([]byte(ret), &vvbu)
	assert.Equal(200, vvbu.Status, "VideoVerificationByUrl() message: "+vvbu.Message)
	assert.Equal("SUCC", vvbu.ResponseCode, "VideoVerificationByUrl() message: "+vvbu.Message)

	// Video Identification
	ret = myVoiceIt.VideoIdentificationByUrl(groupId, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoVerificationB1.mov")
	var vibu structs.VideoIdentificationByUrlReturn
	json.Unmarshal([]byte(ret), &vibu)
	assert.Equal(200, vibu.Status, "VideoIdentificationByUrl() message: "+vibu.Message)
	assert.Equal("SUCC", vibu.ResponseCode, "VideoIdentificationByUrl() message: "+vibu.Message)
	assert.Equal(userId1, vibu.UserId, "VideoIdentificationByUrl() message: "+vibu.Message)

	myVoiceIt.DeleteAllEnrollments(userId1)
	myVoiceIt.DeleteAllEnrollments(userId2)
	myVoiceIt.DeleteUser(userId1)
	myVoiceIt.DeleteUser(userId2)
	myVoiceIt.DeleteGroup(groupId)

	os.Remove("./videoEnrollmentB1.mov")
	os.Remove("./videoEnrollmentB2.mov")
	os.Remove("./videoEnrollmentB3.mov")
	os.Remove("./videoVerificationB1.mov")
	os.Remove("./videoEnrollmentC1.mov")
	os.Remove("./videoEnrollmentC2.mov")
	os.Remove("./videoEnrollmentC3.mov")
}

func TestVoice(t *testing.T) {
	assert := assert.New(t)
	myVoiceIt := &VoiceIt2{ApiKey: os.Getenv("VIAPIKEY"), ApiToken: os.Getenv("VIAPITOKEN"), BaseUrl: "https://api.voiceit.io"}
	ret := myVoiceIt.CreateUser()
	userId1 := getUserId(ret)
	ret = myVoiceIt.CreateUser()
	userId2 := getUserId(ret)
	ret = myVoiceIt.CreateGroup("Sample Group Description")
	groupId := getGroupId(ret)
	myVoiceIt.AddUserToGroup(groupId, userId1)
	myVoiceIt.AddUserToGroup(groupId, userId2)

	// Voice Enrollments
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentA1.wav")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentA2.wav")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentA3.wav")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/verificationA1.wav")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentC1.wav")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentC2.wav")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentC3.wav")

	ret, err := myVoiceIt.CreateVoiceEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./enrollmentA1.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve1 structs.CreateVoiceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve1)
	assert.Equal(201, cve1.Status, "CreateVoiceEnrollment() message: "+cve1.Message)
	assert.Equal("SUCC", cve1.ResponseCode, "CreateVoiceEnrollment() message: "+cve1.Message)
	enrollmentId := cve1.Id

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./enrollmentA2.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve2 structs.CreateVoiceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve2)
	assert.Equal(201, cve2.Status, "CreateVoiceEnrollment() message: "+cve2.Message)
	assert.Equal("SUCC", cve2.ResponseCode, "CreateVoiceEnrollment() message: "+cve2.Message)

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./enrollmentA3.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve3 structs.CreateVoiceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve3)
	assert.Equal(201, cve3.Status, "CreateVoiceEnrollment() message: "+cve3.Message)
	assert.Equal("SUCC", cve3.ResponseCode, "CreateVoiceEnrollment() message: "+cve3.Message)

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./enrollmentC1.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve4 structs.CreateVoiceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve4)
	assert.Equal(201, cve4.Status, "CreateVoiceEnrollment() message: "+cve4.Message)
	assert.Equal("SUCC", cve4.ResponseCode, "CreateVoiceEnrollment() message: "+cve4.Message)

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./enrollmentC2.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve5 structs.CreateVoiceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve5)
	assert.Equal(201, cve5.Status, "CreateVoiceEnrollment() message: "+cve5.Message)
	assert.Equal("SUCC", cve5.ResponseCode, "CreateVoiceEnrollment() message: "+cve5.Message)

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./enrollmentC3.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve6 structs.CreateVoiceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve6)
	assert.Equal(201, cve6.Status, "CreateVoiceEnrollment() message: "+cve6.Message)
	assert.Equal("SUCC", cve6.ResponseCode, "CreateVoiceEnrollment() message: "+cve6.Message)

	// Get Voice Enrollment
	ret = myVoiceIt.GetAllVoiceEnrollments(userId1)
	var gve1 structs.GetAllVoiceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve1)
	assert.Equal(200, gve1.Status, "GetAllVoiceEnrollments() message: "+gve1.Message)
	assert.Equal("SUCC", gve1.ResponseCode, "GetAllVoiceEnrollments() message: "+gve1.Message)
	assert.Equal(3, len(gve1.VoiceEnrollments), "GetAllVoiceEnrollments() message: "+gve1.Message)

	// Voice Verification
	ret, err = myVoiceIt.VoiceVerification(userId1, "en-US", "never forget tomorrow is a new day", "./verificationA1.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var vv structs.VoiceVerificationReturn
	json.Unmarshal([]byte(ret), &vv)
	assert.Equal(200, vv.Status, "VoiceVerification() message: "+vv.Message)
	assert.Equal("SUCC", vv.ResponseCode, "VoiceVerification() message: "+vv.Message)

	// Voice Identification
	ret, err = myVoiceIt.VoiceIdentification(groupId, "en-US", "never forget tomorrow is a new day", "./verificationA1.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var vi structs.VoiceIdentificationReturn
	json.Unmarshal([]byte(ret), &vi)
	assert.Equal(200, vi.Status, "VoiceIdentification() message: "+vi.Message)
	assert.Equal("SUCC", vi.ResponseCode, "VoiceIdentification() message: "+vi.Message)
	assert.Equal(userId1, vi.UserId, "VoiceIdentification() message: "+vi.Message)

	// Delete Voice Enrollment
	ret = myVoiceIt.DeleteVoiceEnrollment(userId1, enrollmentId)
	var dve structs.DeleteVoiceEnrollmentReturn
	json.Unmarshal([]byte(ret), &dve)
	assert.Equal(200, dve.Status, "DeleteVoiceEnrollment() message: "+dve.Message)
	assert.Equal("SUCC", dve.ResponseCode, "DeleteVoiceEnrollment() message: "+dve.Message)

	ret = myVoiceIt.GetAllVoiceEnrollments(userId1)
	var gve2 structs.GetAllVoiceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve2)
	assert.Equal(200, gve2.Status, "GetAllVoiceEnrollments() message: "+gve2.Message)
	assert.Equal("SUCC", gve2.ResponseCode, "GetAllVoiceEnrollments() message: "+gve2.Message)

	// Delete All Voice Enrollments
	ret = myVoiceIt.DeleteAllVoiceEnrollments(userId1)
	var dave structs.DeleteAllVoiceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &dave)
	assert.Equal(200, dave.Status, "DeleteAllVoiceEnrollments() message: "+dave.Message)
	assert.Equal("SUCC", dave.ResponseCode, "DeleteAllVoiceEnrollments() message: "+dave.Message)

	var gve3 structs.GetAllVoiceEnrollmentsReturn
	ret = myVoiceIt.GetAllVoiceEnrollments(userId1)
	json.Unmarshal([]byte(ret), &gve3)
	assert.Equal(200, gve3.Status, "GetAllVoiceEnrollments() message: "+gve3.Message)
	assert.Equal("SUCC", gve3.ResponseCode, "GetAllVoiceEnrollments() message: "+gve3.Message)
	assert.Equal(0, len(gve3.VoiceEnrollments), "GetAllVoiceEnrollments() message: "+gve3.Message)

	// Delete All Enrollments
	ret = myVoiceIt.DeleteAllEnrollments(userId2)
	var dae structs.DeleteAllEnrollmentsReturn
	json.Unmarshal([]byte(ret), &dae)
	assert.Equal(200, dae.Status, "DeleteAllEnrollments() message: "+dae.Message)
	assert.Equal("SUCC", dae.ResponseCode, "DeleteAllEnrollments() message: "+dae.Message)

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
	ret = myVoiceIt.CreateVoiceEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentA1.wav")
	var cvebu1 structs.CreateVoiceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu1)
	assert.Equal(201, cvebu1.Status, "CreateVoiceEnrollmentByUrl() message: "+cvebu1.Message)
	assert.Equal("SUCC", cvebu1.ResponseCode, "CreateVoiceEnrollmentByUrl() message: "+cvebu1.Message)

	ret = myVoiceIt.CreateVoiceEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentA2.wav")
	var cvebu2 structs.CreateVoiceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu2)
	assert.Equal(201, cvebu2.Status, "CreateVoiceEnrollmentByUrl() message: "+cvebu2.Message)
	assert.Equal("SUCC", cvebu2.ResponseCode, "CreateVoiceEnrollmentByUrl() message: "+cvebu2.Message)

	ret = myVoiceIt.CreateVoiceEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentA3.wav")
	var cvebu3 structs.CreateVoiceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu3)
	assert.Equal(201, cvebu3.Status, "CreateVoiceEnrollmentByUrl() message: "+cvebu3.Message)
	assert.Equal("SUCC", cvebu3.ResponseCode, "CreateVoiceEnrollmentByUrl() message: "+cvebu3.Message)

	ret = myVoiceIt.CreateVoiceEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentC1.wav")
	var cvebu4 structs.CreateVoiceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu4)
	assert.Equal(201, cvebu4.Status, "CreateVoiceEnrollmentByUrl() message: "+cvebu4.Message)
	assert.Equal("SUCC", cvebu4.ResponseCode, "CreateVoiceEnrollmentByUrl() message: "+cvebu4.Message)

	ret = myVoiceIt.CreateVoiceEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentC2.wav")
	var cvebu5 structs.CreateVoiceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu5)
	assert.Equal(201, cvebu5.Status, "CreateVoiceEnrollmentByUrl() message: "+cvebu5.Message)
	assert.Equal("SUCC", cvebu5.ResponseCode, "CreateVoiceEnrollmentByUrl() message: "+cvebu5.Message)

	ret = myVoiceIt.CreateVoiceEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentC3.wav")
	var cvebu6 structs.CreateVoiceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu6)
	assert.Equal(201, cvebu6.Status, "CreateVoiceEnrollmentByUrl() message: "+cvebu6.Message)
	assert.Equal("SUCC", cvebu6.ResponseCode, "CreateVoiceEnrollmentByUrl() message: "+cvebu6.Message)

	// Voice Verification
	ret = myVoiceIt.VoiceVerificationByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/verificationA1.wav")
	var vvbu structs.VoiceVerificationByUrlReturn
	json.Unmarshal([]byte(ret), &vvbu)
	assert.Equal(200, vvbu.Status, "VoiceVerificationByUrl() message: "+vvbu.Message)
	assert.Equal("SUCC", vvbu.ResponseCode, "VoiceVerificationByUrl() message: "+vvbu.Message)

	// Voice Identification
	ret = myVoiceIt.VoiceIdentificationByUrl(groupId, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/verificationA1.wav")
	var vibu structs.VoiceIdentificationByUrlReturn
	json.Unmarshal([]byte(ret), &vibu)
	assert.Equal(200, vibu.Status, "VoiceIdentificationByUrl() message: "+vibu.Message)
	assert.Equal("SUCC", vibu.ResponseCode, "VoiceIdentificationByUrl() message: "+vibu.Message)
	assert.Equal(userId1, vibu.UserId, "VoiceIdentificationByUrl() message: "+vibu.Message)

	myVoiceIt.DeleteAllEnrollments(userId1)
	myVoiceIt.DeleteAllEnrollments(userId2)
	myVoiceIt.DeleteUser(userId1)
	myVoiceIt.DeleteUser(userId2)
	myVoiceIt.DeleteGroup(groupId)

	os.Remove("./enrollmentA1.wav")
	os.Remove("./enrollmentA2.wav")
	os.Remove("./enrollmentA3.wav")
	os.Remove("./verificationA1.wav")
	os.Remove("./enrollmentC1.wav")
	os.Remove("./enrollmentC2.wav")
	os.Remove("./enrollmentC3.wav")
}

func TestFace(t *testing.T) {
	assert := assert.New(t)
	myVoiceIt := &VoiceIt2{ApiKey: os.Getenv("VIAPIKEY"), ApiToken: os.Getenv("VIAPITOKEN"), BaseUrl: "https://api.voiceit.io"}
	ret := myVoiceIt.CreateUser()
	userId1 := getUserId(ret)
	ret = myVoiceIt.CreateUser()
	userId2 := getUserId(ret)
	ret = myVoiceIt.CreateGroup("Sample Group Description")
	groupId := getGroupId(ret)
	myVoiceIt.AddUserToGroup(groupId, userId1)
	myVoiceIt.AddUserToGroup(groupId, userId2)

	// Face Enrollments
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceEnrollmentB1.mp4")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceEnrollmentB2.mp4")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceEnrollmentB3.mp4")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceVerificationB1.mp4")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentC1.mov")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentC2.mov")
	downloadFromUrl("https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentC3.mov")

	ret, err := myVoiceIt.CreateFaceEnrollment(userId1, "./faceEnrollmentB1.mp4")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cfe1 structs.CreateFaceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cfe1)
	assert.Equal(201, cfe1.Status, "CreateFaceEnrollment() message: "+cfe1.Message)
	assert.Equal("SUCC", cfe1.ResponseCode, "CreateFaceEnrollment() message: "+cfe1.Message)
	faceEnrollmentId := cfe1.FaceEnrollmentId

	ret, err = myVoiceIt.CreateFaceEnrollment(userId1, "./faceEnrollmentB2.mp4")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cfe2 structs.CreateFaceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cfe2)
	assert.Equal(201, cfe2.Status, "CreateFaceEnrollment() message: "+cfe2.Message)
	assert.Equal("SUCC", cfe2.ResponseCode, "CreateFaceEnrollment() message: "+cfe2.Message)

	ret, err = myVoiceIt.CreateFaceEnrollment(userId1, "./faceEnrollmentB3.mp4")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cfe3 structs.CreateFaceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cfe3)
	assert.Equal(201, cfe3.Status, "CreateFaceEnrollment() message: "+cfe3.Message)
	assert.Equal("SUCC", cfe3.ResponseCode, "CreateFaceEnrollment() message: "+cfe3.Message)

	ret, err = myVoiceIt.CreateFaceEnrollment(userId2, "./videoEnrollmentC1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cfe4 structs.CreateFaceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cfe4)
	assert.Equal(201, cfe4.Status, "CreateFaceEnrollment() message: "+cfe4.Message)
	assert.Equal("SUCC", cfe4.ResponseCode, "CreateFaceEnrollment() message: "+cfe4.Message)

	ret, err = myVoiceIt.CreateFaceEnrollment(userId2, "./videoEnrollmentC1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cfe5 structs.CreateFaceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cfe5)
	assert.Equal(201, cfe5.Status, "CreateFaceEnrollment() message: "+cfe5.Message)
	assert.Equal("SUCC", cfe5.ResponseCode, "CreateFaceEnrollment() message: "+cfe5.Message)

	ret, err = myVoiceIt.CreateFaceEnrollment(userId2, "./videoEnrollmentC1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cfe6 structs.CreateFaceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cfe6)
	assert.Equal(201, cfe6.Status, "CreateFaceEnrollment() message: "+cfe6.Message)
	assert.Equal("SUCC", cfe6.ResponseCode, "CreateFaceEnrollment() message: "+cfe6.Message)

	// Get Face Enrollment
	ret = myVoiceIt.GetAllFaceEnrollments(userId1)
	var fve1 structs.GetAllFaceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &fve1)
	assert.Equal(200, fve1.Status, "GetAllFaceEnrollments() message: "+ret)
	assert.Equal("SUCC", fve1.ResponseCode, "GetAllFaceEnrollments() message: "+ret)
	assert.Equal(3, len(fve1.FaceEnrollments), "GetAllFaceEnrollments() message: "+ret)

	// Face Verification
	ret, err = myVoiceIt.FaceVerification(userId1, "./faceVerificationB1.mp4")
	if err != nil {
		t.Fatal(err.Error())
	}
	var fv structs.FaceVerificationReturn
	json.Unmarshal([]byte(ret), &fv)
	assert.Equal(200, fv.Status, "FaceVerification() message: "+fv.Message)
	assert.Equal("SUCC", fv.ResponseCode, "FaceVerification() message: "+fv.Message)

	// Face Identification
	ret, err = myVoiceIt.FaceIdentification(groupId, "./faceVerificationB1.mp4")
	if err != nil {
		t.Fatal(err.Error())
	}
	var fi structs.FaceIdentificationReturn
	json.Unmarshal([]byte(ret), &fi)
	assert.Equal(200, fi.Status, "FaceIdentification() message: "+fi.Message)
	assert.Equal("SUCC", fi.ResponseCode, "FaceIdentification() message: "+fi.Message)
	assert.Equal(userId1, fi.UserId, "FaceIdentification() message: "+fi.Message)

	ret = myVoiceIt.GetAllFaceEnrollments(userId1)

	// Delete Face Enrollment
	ret = myVoiceIt.DeleteFaceEnrollment(userId1, faceEnrollmentId)
	var dfe structs.DeleteFaceEnrollmentReturn
	json.Unmarshal([]byte(ret), &dfe)
	assert.Equal(200, dfe.Status, "DeleteFaceEnrollment() message: "+dfe.Message)
	assert.Equal("SUCC", dfe.ResponseCode, "DeleteFaceEnrollment() message: "+dfe.Message)

	ret = myVoiceIt.GetAllFaceEnrollments(userId1)
	var fve2 structs.GetAllFaceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &fve2)
	assert.Equal(200, fve2.Status, "GetAllFaceEnrollments() message: "+fve2.Message)
	assert.Equal("SUCC", fve2.ResponseCode, "GetAllFaceEnrollments() message: "+fve2.Message)

	// Delete All Face Enrollments
	ret = myVoiceIt.DeleteAllFaceEnrollments(userId1)
	var dafe structs.DeleteAllFaceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &dafe)
	assert.Equal(200, dafe.Status, "DeleteAllFaceEnrollments() message: "+dafe.Message)
	assert.Equal("SUCC", dafe.ResponseCode, "DeleteAllFaceEnrollments() message: "+dafe.Message)

	var fve3 structs.GetAllFaceEnrollmentsReturn
	ret = myVoiceIt.GetAllFaceEnrollments(userId1)
	json.Unmarshal([]byte(ret), &fve3)
	assert.Equal(200, fve3.Status, "GetAllFaceEnrollments() message: "+fve3.Message)
	assert.Equal("SUCC", fve3.ResponseCode, "GetAllFaceEnrollments() message: "+fve3.Message)
	assert.Equal(0, len(fve3.FaceEnrollments), "GetAllFaceEnrollments() message: "+ret)

	// Delete All Enrollments
	ret = myVoiceIt.DeleteAllEnrollments(userId2)
	var dae structs.DeleteAllEnrollmentsReturn
	json.Unmarshal([]byte(ret), &dae)
	assert.Equal(200, dae.Status, "DeleteAllEnrollments() message: "+dae.Message)
	assert.Equal("SUCC", dae.ResponseCode, "DeleteAllEnrollments() message: "+dae.Message)

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
	ret = myVoiceIt.CreateFaceEnrollmentByUrl(userId1, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceEnrollmentB1.mp4")
	var cfebu1 structs.CreateFaceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cfebu1)
	assert.Equal(201, cfebu1.Status, "CreateFaceEnrollmentByUrl() message: "+cfebu1.Message)
	assert.Equal("SUCC", cfebu1.ResponseCode, "CreateFaceEnrollmentByUrl() message: "+cfebu1.Message)

	ret = myVoiceIt.CreateFaceEnrollmentByUrl(userId1, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceEnrollmentB2.mp4")
	var cfebu2 structs.CreateFaceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cfebu2)
	assert.Equal(201, cfebu2.Status, "CreateFaceEnrollmentByUrl() message: "+cfebu2.Message)
	assert.Equal("SUCC", cfebu2.ResponseCode, "CreateFaceEnrollmentByUrl() message: "+cfebu2.Message)

	ret = myVoiceIt.CreateFaceEnrollmentByUrl(userId1, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceEnrollmentB3.mp4")
	var cfebu3 structs.CreateFaceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cfebu3)
	assert.Equal(201, cfebu3.Status, "CreateFaceEnrollmentByUrl() message: "+cfebu3.Message)
	assert.Equal("SUCC", cfebu3.ResponseCode, "CreateFaceEnrollmentByUrl() message: "+cfebu3.Message)

	ret = myVoiceIt.CreateFaceEnrollmentByUrl(userId2, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentC1.mov")
	var cfebu4 structs.CreateFaceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cfebu4)
	assert.Equal(201, cfebu4.Status, "CreateFaceEnrollmentByUrl() message: "+cfebu4.Message)
	assert.Equal("SUCC", cfebu4.ResponseCode, "CreateFaceEnrollmentByUrl() message: "+cfebu4.Message)

	ret = myVoiceIt.CreateFaceEnrollmentByUrl(userId2, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentC2.mov")
	var cfebu5 structs.CreateFaceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cfebu5)
	assert.Equal(201, cfebu5.Status, "CreateFaceEnrollmentByUrl() message: "+cfebu5.Message)
	assert.Equal("SUCC", cfebu5.ResponseCode, "CreateFaceEnrollmentByUrl() message: "+cfebu5.Message)

	ret = myVoiceIt.CreateFaceEnrollmentByUrl(userId2, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentC3.mov")
	var cfebu6 structs.CreateFaceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cfebu6)
	assert.Equal(201, cfebu6.Status, "CreateFaceEnrollmentByUrl() message: "+cfebu6.Message)
	assert.Equal("SUCC", cfebu6.ResponseCode, "CreateFaceEnrollmentByUrl() message: "+cfebu6.Message)

	// Face Verification
	ret = myVoiceIt.FaceVerificationByUrl(userId1, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceVerificationB1.mp4")
	var fvbu structs.FaceVerificationByUrlReturn
	json.Unmarshal([]byte(ret), &fvbu)
	assert.Equal(200, fvbu.Status, "FaceVerificationByUrl() message: "+fvbu.Message)
	assert.Equal("SUCC", fvbu.ResponseCode, "FaceVerificationByUrl() message: "+fvbu.Message)

	// Face Identification
	ret = myVoiceIt.FaceIdentificationByUrl(groupId, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceVerificationB1.mp4")
	var fibu structs.FaceIdentificationByUrlReturn
	json.Unmarshal([]byte(ret), &fibu)
	assert.Equal(200, fibu.Status, "FaceIdentificationByUrl() message: "+fibu.Message)
	assert.Equal("SUCC", fibu.ResponseCode, "FaceIdentificationByUrl() message: "+fibu.Message)
	assert.Equal(userId1, fibu.UserId, "FaceIdentificationByUrl() message: "+fibu.Message)

	myVoiceIt.DeleteAllEnrollments(userId1)
	myVoiceIt.DeleteAllEnrollments(userId2)
	myVoiceIt.DeleteUser(userId1)
	myVoiceIt.DeleteUser(userId2)
	myVoiceIt.DeleteGroup(groupId)

	os.Remove("./faceEnrollmentB1.mp4")
	os.Remove("./faceEnrollmentB2.mp4")
	os.Remove("./faceEnrollmentB3.mp4")
	os.Remove("./faceVerificationB1.mp4")
	os.Remove("./videoEnrollmentC1.mov")
	os.Remove("./videoEnrollmentC2.mov")
	os.Remove("./videoEnrollmentC3.mov")
}
