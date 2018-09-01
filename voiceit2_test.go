package voiceit2

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"./structs"
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
	var cu structs.CreateUserReturn
	json.Unmarshal([]byte(ret), &cu)
	assert.Equal(201, cu.Status, "message: "+cu.Message)
	assert.Equal("SUCC", cu.ResponseCode, "message: "+cu.Message)
	userId := getUserId(ret)

	ret = myVoiceIt.GetAllUsers()
	var gau structs.GetAllUsersReturn
	json.Unmarshal([]byte(ret), &gau)
	assert.Equal(200, gau.Status, "message: "+ret)
	assert.Equal("SUCC", gau.ResponseCode, "message: "+ret)
	assert.True(0 < len(gau.Users), "message: "+ret)

	ret = myVoiceIt.CheckUserExists(userId)
	var cue structs.CheckUserExistsReturn
	json.Unmarshal([]byte(ret), &cue)
	assert.Equal(200, cue.Status, "message: "+cue.Message)
	assert.Equal("SUCC", cue.ResponseCode, "message: "+cue.Message)

	ret = myVoiceIt.CreateGroup("Sample Group Description")
	var cg structs.CreateGroupReturn
	json.Unmarshal([]byte(ret), &cg)
	assert.Equal(201, cg.Status, "message: "+cg.Message)
	assert.Equal("SUCC", cg.ResponseCode, "message: "+cg.Message)
	groupId := cg.GroupId

	ret = myVoiceIt.CheckGroupExists(groupId)
	var cge structs.CheckGroupExistsReturn
	json.Unmarshal([]byte(ret), &cge)
	assert.Equal(200, cge.Status, "message: "+cge.Message)
	assert.Equal("SUCC", cge.ResponseCode, "message: "+cge.Message)

	ret = myVoiceIt.GetGroup(groupId)
	var gg structs.GetGroupsForUserReturn
	json.Unmarshal([]byte(ret), &gg)
	assert.Equal(200, gg.Status, "message: "+gg.Message)
	assert.Equal("SUCC", gg.ResponseCode, "message: "+gg.Message)

	ret = myVoiceIt.AddUserToGroup(groupId, userId)
	var autg structs.AddUserToGroupReturn
	json.Unmarshal([]byte(ret), &autg)
	assert.Equal(200, autg.Status, "message: "+autg.Message)
	assert.Equal("SUCC", autg.ResponseCode, "message: "+autg.Message)

	ret = myVoiceIt.GetGroupsForUser(userId)
	var ggfu structs.GetGroupsForUserReturn
	json.Unmarshal([]byte(ret), &ggfu)
	assert.Equal(200, ggfu.Status, "message: "+ggfu.Message)
	assert.Equal("SUCC", ggfu.ResponseCode, "message: "+ggfu.Message)
	assert.Equal(1, len(ggfu.Groups), "message: "+ggfu.Message)
	assert.Equal(200, ggfu.Status, "message: "+ggfu.Message)

	ret = myVoiceIt.RemoveUserFromGroup(groupId, userId)
	var rufg structs.RemoveUserFromGroupReturn
	json.Unmarshal([]byte(ret), &rufg)
	assert.Equal(200, rufg.Status, "message: "+rufg.Message)
	assert.Equal("SUCC", rufg.ResponseCode, "message: "+rufg.Message)

	ret = myVoiceIt.DeleteUser(userId)
	var du structs.DeleteUserReturn
	json.Unmarshal([]byte(ret), &du)
	assert.Equal(200, du.Status, "message: "+du.Message)
	assert.Equal("SUCC", du.ResponseCode, "message: "+du.Message)

	ret = myVoiceIt.DeleteGroup(groupId)
	var dg structs.DeleteGroupReturn
	json.Unmarshal([]byte(ret), &dg)
	assert.Equal(200, dg.Status, "message: "+dg.Message)
	assert.Equal("SUCC", dg.ResponseCode, "message: "+dg.Message)

	ret = myVoiceIt.GetPhrases("en-US")
	var gp structs.GetPhrasesReturn
	json.Unmarshal([]byte(ret), &gp)
	assert.Equal(200, gp.Status, "message: "+gp.Message)
	assert.Equal("SUCC", gp.ResponseCode, "message: "+gp.Message)
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
	var cve1 structs.CreateVideoEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve1)
	assert.Equal(201, cve1.Status, "message: "+cve1.Message)
	assert.Equal("SUCC", cve1.ResponseCode, "message: "+cve1.Message)
	enrollmentId := cve1.Id

	ret, err = myVoiceIt.CreateVideoEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentArmaan2.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve2 structs.CreateVideoEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve2)
	assert.Equal(201, cve2.Status, "message: "+cve2.Message)
	assert.Equal("SUCC", cve2.ResponseCode, "message: "+cve2.Message)

	ret, err = myVoiceIt.CreateVideoEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentArmaan3.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve3 structs.CreateVideoEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve3)
	assert.Equal(201, cve3.Status, "message: "+cve3.Message)
	assert.Equal("SUCC", cve3.ResponseCode, "message: "+cve3.Message)

	ret, err = myVoiceIt.CreateVideoEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentStephen1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve4 structs.CreateVideoEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve4)
	assert.Equal(201, cve4.Status, "message: "+cve4.Message)
	assert.Equal("SUCC", cve4.ResponseCode, "message: "+cve4.Message)

	ret, err = myVoiceIt.CreateVideoEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentStephen2.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve5 structs.CreateVideoEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve5)
	assert.Equal(201, cve5.Status, "message: "+cve5.Message)
	assert.Equal("SUCC", cve5.ResponseCode, "message: "+cve5.Message)

	ret, err = myVoiceIt.CreateVideoEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentStephen3.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve6 structs.CreateVideoEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve6)
	assert.Equal(201, cve6.Status, "message: "+cve6.Message)
	assert.Equal("SUCC", cve6.ResponseCode, "message: "+cve6.Message)

	// Get Video Enrollment
	ret = myVoiceIt.GetAllVideoEnrollments(userId1)
	var gve1 structs.GetAllVideoEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve1)
	assert.Equal(200, gve1.Status, "message: "+ret)
	assert.Equal("SUCC", gve1.ResponseCode, "message: "+ret)
	assert.Equal(3, len(gve1.VideoEnrollments), "message: "+ret)

	// Video Verification
	ret, err = myVoiceIt.VideoVerification(userId1, "en-US", "never forget tomorrow is a new day", "./videoVerificationArmaan1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var vv1 structs.VideoVerificationReturn
	json.Unmarshal([]byte(ret), &vv1)
	assert.Equal(200, vv1.Status, "message: "+vv1.Message)
	assert.Equal("SUCC", vv1.ResponseCode, "message: "+vv1.Message)

	// Video Identification
	ret, err = myVoiceIt.VideoIdentification(groupId, "en-US", "never forget tomorrow is a new day", "./videoVerificationArmaan1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var vi1 structs.VideoIdentificationReturn
	json.Unmarshal([]byte(ret), &vi1)
	assert.Equal(200, vi1.Status, "message: "+vi1.Message)
	assert.Equal("SUCC", vi1.ResponseCode, "message: "+vi1.Message)
	assert.Equal(userId1, vi1.UserId, "message: "+vi1.Message)

	// Delete Video Enrollment
	ret = myVoiceIt.DeleteVideoEnrollment(userId1, enrollmentId)
	var dve structs.DeleteVideoEnrollmentReturn
	json.Unmarshal([]byte(ret), &dve)
	assert.Equal(200, dve.Status, "message: "+dve.Message)
	assert.Equal("SUCC", dve.ResponseCode, "message: "+dve.Message)

	ret = myVoiceIt.GetAllVideoEnrollments(userId1)
	var gve2 structs.GetAllVideoEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve2)
	for _, videoEnrollment := range gve2.VideoEnrollments {
		assert.NotEqual(enrollmentId, videoEnrollment.VideoEnrollmentId, "message: "+gve2.Message)
	}

	// Delete All Video Enrollments
	ret = myVoiceIt.DeleteAllVideoEnrollments(userId1)
	var dave structs.DeleteAllVideoEnrollmentsReturn
	json.Unmarshal([]byte(ret), &dave)
	assert.Equal(200, dave.Status, "message: "+dave.Message)
	assert.Equal("SUCC", dave.ResponseCode, "message: "+dave.Message)

	var gve3 structs.GetAllVideoEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve3)
	assert.Equal(200, gve3.Status, "message: "+dave.Message)
	assert.Equal("SUCC", gve3.ResponseCode, "message: "+dave.Message)
	assert.Equal(0, len(gve3.VideoEnrollments), "message: "+dave.Message)

	// Delete All Enrollments
	ret = myVoiceIt.DeleteAllEnrollments(userId2)
	var dae structs.DeleteAllEnrollmentsReturn
	json.Unmarshal([]byte(ret), &dae)
	assert.Equal(200, dae.Status, "message: "+dae.Message)
	assert.Equal("SUCC", dae.ResponseCode, "message: "+dae.Message)

	ret = myVoiceIt.GetAllVideoEnrollments(userId2)
	var gve4 structs.GetAllVideoEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve4)
	assert.Equal(200, gve4.Status, "message: "+dae.Message)
	assert.Equal("SUCC", gve4.ResponseCode, "message: "+dae.Message)
	assert.Equal(0, len(gve4.VideoEnrollments), "message: "+dae.Message)

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
	var cvebu1 structs.CreateVideoEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu1)
	assert.Equal(201, cvebu1.Status, "message: "+cvebu1.Message)
	assert.Equal("SUCC", cvebu1.ResponseCode, "message: "+cvebu1.Message)

	ret = myVoiceIt.CreateVideoEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentArmaan2.mov")
	var cvebu2 structs.CreateVideoEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu2)
	assert.Equal(201, cvebu2.Status, "message: "+cvebu2.Message)
	assert.Equal("SUCC", cvebu2.ResponseCode, "message: "+cvebu2.Message)

	ret = myVoiceIt.CreateVideoEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentArmaan3.mov")
	var cvebu3 structs.CreateVideoEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu3)
	assert.Equal(201, cvebu3.Status, "message: "+cvebu3.Message)
	assert.Equal("SUCC", cvebu3.ResponseCode, "message: "+cvebu3.Message)

	ret = myVoiceIt.CreateVideoEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentStephen1.mov")
	var cvebu4 structs.CreateVideoEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu4)
	assert.Equal(201, cvebu4.Status, "message: "+cvebu4.Message)
	assert.Equal("SUCC", cvebu4.ResponseCode, "message: "+cvebu4.Message)

	ret = myVoiceIt.CreateVideoEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentStephen2.mov")
	var cvebu5 structs.CreateVideoEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu5)
	assert.Equal(201, cvebu5.Status, "message: "+cvebu5.Message)
	assert.Equal("SUCC", cvebu5.ResponseCode, "message: "+cvebu5.Message)

	ret = myVoiceIt.CreateVideoEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentStephen3.mov")
	var cvebu6 structs.CreateVideoEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu6)
	assert.Equal(201, cvebu6.Status, "message: "+cvebu6.Message)
	assert.Equal("SUCC", cvebu6.ResponseCode, "message: "+cvebu6.Message)

	// Video Verification
	ret = myVoiceIt.VideoVerificationByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoVerificationArmaan1.mov")
	var vvbu structs.VideoVerificationByUrlReturn
	json.Unmarshal([]byte(ret), &vvbu)
	assert.Equal(200, vvbu.Status, "message: "+vvbu.Message)
	assert.Equal("SUCC", vvbu.ResponseCode, "message: "+vvbu.Message)

	// Video Identification
	ret = myVoiceIt.VideoIdentificationByUrl(groupId, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoVerificationArmaan1.mov")
	var vibu structs.VideoIdentificationByUrlReturn
	json.Unmarshal([]byte(ret), &vibu)
	assert.Equal(200, vibu.Status, "message: "+vibu.Message)
	assert.Equal("SUCC", vibu.ResponseCode, "message: "+vibu.Message)
	assert.Equal(userId1, vibu.UserId, "message: "+vibu.Message)

	myVoiceIt.DeleteAllEnrollments(userId1)
	myVoiceIt.DeleteAllEnrollments(userId2)
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
	var cve1 structs.CreateVoiceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve1)
	assert.Equal(201, cve1.Status, "message: "+cve1.Message)
	assert.Equal("SUCC", cve1.ResponseCode, "message: "+cve1.Message)
	enrollmentId := cve1.Id

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./enrollmentArmaan2.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve2 structs.CreateVoiceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve2)
	assert.Equal(201, cve2.Status, "message: "+cve2.Message)
	assert.Equal("SUCC", cve2.ResponseCode, "message: "+cve2.Message)

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./enrollmentArmaan3.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve3 structs.CreateVoiceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve3)
	assert.Equal(201, cve3.Status, "message: "+cve3.Message)
	assert.Equal("SUCC", cve3.ResponseCode, "message: "+cve3.Message)

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./enrollmentStephen1.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve4 structs.CreateVoiceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve4)
	assert.Equal(201, cve4.Status, "message: "+cve4.Message)
	assert.Equal("SUCC", cve4.ResponseCode, "message: "+cve4.Message)

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./enrollmentStephen2.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve5 structs.CreateVoiceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve5)
	assert.Equal(201, cve5.Status, "message: "+cve5.Message)
	assert.Equal("SUCC", cve5.ResponseCode, "message: "+cve5.Message)

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./enrollmentStephen3.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve6 structs.CreateVoiceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cve6)
	assert.Equal(201, cve6.Status, "message: "+cve6.Message)
	assert.Equal("SUCC", cve6.ResponseCode, "message: "+cve6.Message)

	// Get Voice Enrollment
	ret = myVoiceIt.GetAllVoiceEnrollments(userId1)
	var gve1 structs.GetAllVoiceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve1)
	assert.Equal(200, gve1.Status, "message: "+gve1.Message)
	assert.Equal("SUCC", gve1.ResponseCode, "message: "+gve1.Message)
	assert.Equal(3, len(gve1.VoiceEnrollments), "message: "+gve1.Message)

	// Voice Verification
	ret, err = myVoiceIt.VoiceVerification(userId1, "en-US", "never forget tomorrow is a new day", "./verificationArmaan1.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var vv structs.VoiceVerificationReturn
	json.Unmarshal([]byte(ret), &vv)
	assert.Equal(200, vv.Status, "message: "+vv.Message)
	assert.Equal("SUCC", vv.ResponseCode, "message: "+vv.Message)

	// Voice Identification
	ret, err = myVoiceIt.VoiceIdentification(groupId, "en-US", "never forget tomorrow is a new day", "./verificationArmaan1.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var vi structs.VoiceIdentificationReturn
	json.Unmarshal([]byte(ret), &vi)
	assert.Equal(200, vi.Status, "message: "+vi.Message)
	assert.Equal("SUCC", vi.ResponseCode, "message: "+vi.Message)
	assert.Equal(userId1, vi.UserId, "message: "+vi.Message)

	// Delete Voice Enrollment
	ret = myVoiceIt.DeleteVoiceEnrollment(userId1, enrollmentId)
	var dve structs.DeleteVoiceEnrollmentReturn
	json.Unmarshal([]byte(ret), &dve)
	assert.Equal(200, dve.Status, "message: "+dve.Message)
	assert.Equal("SUCC", dve.ResponseCode, "message: "+dve.Message)

	ret = myVoiceIt.GetAllVoiceEnrollments(userId1)
	var gve2 structs.GetAllVoiceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve2)
	assert.Equal(200, gve2.Status, "message: "+gve2.Message)
	assert.Equal("SUCC", gve2.ResponseCode, "message: "+gve2.Message)
	for _, videoEnrollment := range gve2.VoiceEnrollments {
		assert.NotEqual(enrollmentId, videoEnrollment.VoiceEnrollmentId, "message: "+ret)
	}

	// Delete All Voice Enrollments
	ret = myVoiceIt.DeleteAllVoiceEnrollments(userId1)
	var dave structs.DeleteAllVoiceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &dave)
	assert.Equal(200, dave.Status, "message: "+dave.Message)
	assert.Equal("SUCC", dave.ResponseCode, "message: "+dave.Message)

	var gve3 structs.GetAllVoiceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &gve3)
	assert.Equal(200, gve3.Status, "message: "+gve3.Message)
	assert.Equal("SUCC", gve3.ResponseCode, "message: "+gve3.Message)
	assert.Equal(0, len(gve3.VoiceEnrollments), "message: "+gve3.Message)

	// Delete All Enrollments
	ret = myVoiceIt.DeleteAllEnrollments(userId2)
	var dae structs.DeleteAllEnrollmentsReturn
	json.Unmarshal([]byte(ret), &dae)
	assert.Equal(200, dae.Status, "message: "+dae.Message)
	assert.Equal("SUCC", dae.ResponseCode, "message: "+dae.Message)

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
	var cvebu1 structs.CreateVoiceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu1)
	assert.Equal(201, cvebu1.Status, "message: "+cvebu1.Message)
	assert.Equal("SUCC", cvebu1.ResponseCode, "message: "+cvebu1.Message)

	ret = myVoiceIt.CreateVoiceEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentArmaan2.wav")
	var cvebu2 structs.CreateVoiceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu2)
	assert.Equal(201, cvebu2.Status, "message: "+cvebu2.Message)
	assert.Equal("SUCC", cvebu2.ResponseCode, "message: "+cvebu2.Message)

	ret = myVoiceIt.CreateVoiceEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentArmaan3.wav")
	var cvebu3 structs.CreateVoiceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu3)
	assert.Equal(201, cvebu3.Status, "message: "+cvebu3.Message)
	assert.Equal("SUCC", cvebu3.ResponseCode, "message: "+cvebu3.Message)

	ret = myVoiceIt.CreateVoiceEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentStephen1.wav")
	var cvebu4 structs.CreateVoiceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu4)
	assert.Equal(201, cvebu4.Status, "message: "+cvebu4.Message)
	assert.Equal("SUCC", cvebu4.ResponseCode, "message: "+cvebu4.Message)

	ret = myVoiceIt.CreateVoiceEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentStephen2.wav")
	var cvebu5 structs.CreateVoiceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu5)
	assert.Equal(201, cvebu5.Status, "message: "+cvebu5.Message)
	assert.Equal("SUCC", cvebu5.ResponseCode, "message: "+cvebu5.Message)

	ret = myVoiceIt.CreateVoiceEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/enrollmentStephen3.wav")
	var cvebu6 structs.CreateVoiceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cvebu6)
	assert.Equal(201, cvebu6.Status, "message: "+cvebu6.Message)
	assert.Equal("SUCC", cvebu6.ResponseCode, "message: "+cvebu6.Message)

	// Voice Verification
	ret = myVoiceIt.VoiceVerificationByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/verificationArmaan1.wav")
	var vvbu structs.VoiceVerificationByUrlReturn
	json.Unmarshal([]byte(ret), &vvbu)
	assert.Equal(200, vvbu.Status, "message: "+vvbu.Message)
	assert.Equal("SUCC", vvbu.ResponseCode, "message: "+vvbu.Message)

	// Voice Identification
	ret = myVoiceIt.VoiceIdentificationByUrl(groupId, "en-US", "never forget tomorrow is a new day", "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/verificationArmaan1.wav")
	var vibu structs.VoiceIdentificationByUrlReturn
	json.Unmarshal([]byte(ret), &vibu)
	assert.Equal(200, vibu.Status, "message: "+vibu.Message)
	assert.Equal("SUCC", vibu.ResponseCode, "message: "+vibu.Message)
	assert.Equal(userId1, vibu.UserId, "message: "+vibu.Message)

	myVoiceIt.DeleteAllEnrollments(userId1)
	myVoiceIt.DeleteAllEnrollments(userId2)
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
	var cfe1 structs.CreateFaceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cfe1)
	assert.Equal(201, cfe1.Status, "message: "+cfe1.Message)
	assert.Equal("SUCC", cfe1.ResponseCode, "message: "+cfe1.Message)
	faceEnrollmentId := cfe1.FaceEnrollmentId

	ret, err = myVoiceIt.CreateFaceEnrollment(userId1, "./faceEnrollmentArmaan2.mp4")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cfe2 structs.CreateFaceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cfe2)
	assert.Equal(201, cfe2.Status, "message: "+cfe2.Message)
	assert.Equal("SUCC", cfe2.ResponseCode, "message: "+cfe2.Message)

	ret, err = myVoiceIt.CreateFaceEnrollment(userId1, "./faceEnrollmentArmaan3.mp4")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cfe3 structs.CreateFaceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cfe3)
	assert.Equal(201, cfe3.Status, "message: "+cfe3.Message)
	assert.Equal("SUCC", cfe3.ResponseCode, "message: "+cfe3.Message)

	ret, err = myVoiceIt.CreateFaceEnrollment(userId2, "./videoEnrollmentStephen1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cfe4 structs.CreateFaceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cfe4)
	assert.Equal(201, cfe4.Status, "message: "+cfe4.Message)
	assert.Equal("SUCC", cfe4.ResponseCode, "message: "+cfe4.Message)

	ret, err = myVoiceIt.CreateFaceEnrollment(userId2, "./videoEnrollmentStephen1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cfe5 structs.CreateFaceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cfe5)
	assert.Equal(201, cfe5.Status, "message: "+cfe5.Message)
	assert.Equal("SUCC", cfe5.ResponseCode, "message: "+cfe5.Message)

	ret, err = myVoiceIt.CreateFaceEnrollment(userId2, "./videoEnrollmentStephen1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cfe6 structs.CreateFaceEnrollmentReturn
	json.Unmarshal([]byte(ret), &cfe6)
	assert.Equal(201, cfe6.Status, "message: "+cfe6.Message)
	assert.Equal("SUCC", cfe6.ResponseCode, "message: "+cfe6.Message)

	// Get Face Enrollment
	ret = myVoiceIt.GetAllFaceEnrollments(userId1)
	var fve1 structs.GetAllFaceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &fve1)
	assert.Equal(200, fve1.Status, "message: "+ret)
	assert.Equal("SUCC", fve1.ResponseCode, "message: "+ret)
	assert.Equal(3, len(fve1.FaceEnrollments), "message: "+ret)

	// Face Verification
	ret, err = myVoiceIt.FaceVerification(userId1, "./faceVerificationArmaan1.mp4")
	if err != nil {
		t.Fatal(err.Error())
	}
	var fv structs.FaceVerificationReturn
	json.Unmarshal([]byte(ret), &fv)
	assert.Equal(200, fv.Status, "message: "+fv.Message)
	assert.Equal("SUCC", fv.ResponseCode, "message: "+fv.Message)

	// Face Identification
	ret, err = myVoiceIt.FaceIdentification(groupId, "./faceVerificationArmaan1.mp4")
	if err != nil {
		t.Fatal(err.Error())
	}
	var fi structs.FaceIdentificationReturn
	json.Unmarshal([]byte(ret), &fi)
	assert.Equal(200, fi.Status, "message: "+fi.Message)
	assert.Equal("SUCC", fi.ResponseCode, "message: "+fi.Message)
	assert.Equal(userId1, fi.UserId, "message: "+fi.Message)

	ret = myVoiceIt.GetAllFaceEnrollments(userId1)

	// Delete Face Enrollment
	ret = myVoiceIt.DeleteFaceEnrollment(userId1, faceEnrollmentId)
	var dfe structs.DeleteFaceEnrollmentReturn
	json.Unmarshal([]byte(ret), &dfe)
	assert.Equal(200, dfe.Status, "message: "+dfe.Message)
	assert.Equal("SUCC", dfe.ResponseCode, "message: "+dfe.Message)

	ret = myVoiceIt.GetAllFaceEnrollments(userId1)
	var fve2 structs.GetAllFaceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &fve2)
	assert.Equal(200, fve2.Status, "message: "+fve2.Message)
	assert.Equal("SUCC", fve2.ResponseCode, "message: "+fve2.Message)
	for _, faceEnrollment := range fve2.FaceEnrollments {
		assert.NotEqual(faceEnrollmentId, faceEnrollment.FaceEnrollmentId, "message: "+ret)
	}

	// Delete All Face Enrollments
	ret = myVoiceIt.DeleteAllFaceEnrollments(userId1)
	var dafe structs.DeleteAllFaceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &dafe)
	assert.Equal(200, dafe.Status, "message: "+dafe.Message)
	assert.Equal("SUCC", dafe.ResponseCode, "message: "+dafe.Message)

	var fve3 structs.GetAllFaceEnrollmentsReturn
	json.Unmarshal([]byte(ret), &fve3)
	assert.Equal(200, fve3.Status, "message: "+fve3.Message)
	assert.Equal("SUCC", fve3.ResponseCode, "message: "+fve3.Message)
	assert.Equal(0, len(fve3.FaceEnrollments), "message: "+ret)

	// Delete All Enrollments
	ret = myVoiceIt.DeleteAllEnrollments(userId2)
	var dae structs.DeleteAllEnrollmentsReturn
	json.Unmarshal([]byte(ret), &dae)
	assert.Equal(200, dae.Status, "message: "+dae.Message)
	assert.Equal("SUCC", dae.ResponseCode, "message: "+dae.Message)

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
	var cfebu1 structs.CreateFaceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cfebu1)
	assert.Equal(201, cfebu1.Status, "message: "+cfebu1.Message)
	assert.Equal("SUCC", cfebu1.ResponseCode, "message: "+cfebu1.Message)

	ret = myVoiceIt.CreateFaceEnrollmentByUrl(userId1, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceEnrollmentArmaan2.mp4")
	var cfebu2 structs.CreateFaceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cfebu2)
	assert.Equal(201, cfebu2.Status, "message: "+cfebu2.Message)
	assert.Equal("SUCC", cfebu2.ResponseCode, "message: "+cfebu2.Message)

	ret = myVoiceIt.CreateFaceEnrollmentByUrl(userId1, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceEnrollmentArmaan3.mp4")
	var cfebu3 structs.CreateFaceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cfebu3)
	assert.Equal(201, cfebu3.Status, "message: "+cfebu3.Message)
	assert.Equal("SUCC", cfebu3.ResponseCode, "message: "+cfebu3.Message)

	ret = myVoiceIt.CreateFaceEnrollmentByUrl(userId2, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentStephen1.mov")
	var cfebu4 structs.CreateFaceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cfebu4)
	assert.Equal(201, cfebu4.Status, "message: "+cfebu4.Message)
	assert.Equal("SUCC", cfebu4.ResponseCode, "message: "+cfebu4.Message)

	ret = myVoiceIt.CreateFaceEnrollmentByUrl(userId2, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentStephen2.mov")
	var cfebu5 structs.CreateFaceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cfebu5)
	assert.Equal(201, cfebu5.Status, "message: "+cfebu5.Message)
	assert.Equal("SUCC", cfebu5.ResponseCode, "message: "+cfebu5.Message)

	ret = myVoiceIt.CreateFaceEnrollmentByUrl(userId2, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/videoEnrollmentStephen3.mov")
	var cfebu6 structs.CreateFaceEnrollmentByUrlReturn
	json.Unmarshal([]byte(ret), &cfebu6)
	assert.Equal(201, cfebu6.Status, "message: "+cfebu6.Message)
	assert.Equal("SUCC", cfebu6.ResponseCode, "message: "+cfebu6.Message)

	// Face Verification
	ret = myVoiceIt.FaceVerificationByUrl(userId1, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceVerificationArmaan1.mp4")
	var fvbu structs.FaceVerificationByUrlReturn
	json.Unmarshal([]byte(ret), &fvbu)
	assert.Equal(200, fvbu.Status, "message: "+fvbu.Message)
	assert.Equal("SUCC", fvbu.ResponseCode, "message: "+fvbu.Message)

	// Face Identification
	ret = myVoiceIt.FaceIdentificationByUrl(groupId, "https://s3.amazonaws.com/voiceit-api2-testing-files/test-data/faceVerificationArmaan1.mp4")
	var fibu structs.FaceIdentificationByUrlReturn
	json.Unmarshal([]byte(ret), &fibu)
	assert.Equal(200, fibu.Status, "message: "+fibu.Message)
	assert.Equal("SUCC", fibu.ResponseCode, "message: "+fibu.Message)
	assert.Equal(userId1, fibu.UserId, "message: "+fibu.Message)

	myVoiceIt.DeleteAllEnrollments(userId1)
	myVoiceIt.DeleteAllEnrollments(userId2)
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
