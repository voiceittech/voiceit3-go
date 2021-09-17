package voiceit2

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/voiceittech/VoiceIt2-Go/v2/structs"
)

func getUserId(arg []byte) string {
	var dat map[string]interface{}
	json.Unmarshal(arg, &dat)
	return dat["userId"].(string)
}

func getGroupId(arg []byte) string {
	var dat map[string]interface{}
	json.Unmarshal(arg, &dat)
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
	myVoiceIt := VoiceIt2{APIKey: os.Getenv("VIAPIKEY"), APIToken: os.Getenv("VIAPITOKEN"), BaseUrl: "https://api.voiceit.io"}
	_, err := myVoiceIt.CreateVoiceEnrollment("", "", "", "not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to CreateVoiceEnrollmentFunction (should return real error)")
	_, err = myVoiceIt.CreateVideoEnrollment("", "", "", "not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to CreateVideoEnrollment (should return real error)")
	_, err = myVoiceIt.CreateSplitVideoEnrollment("", "", "", "not_a_real.file", "also_not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to CreateSplitVideoEnrollment (should return real error)")
	_, err = myVoiceIt.CreateFaceEnrollment("", "not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to CreateFaceEnrollment (should return real error)")
	_, err = myVoiceIt.CreateFaceEnrollment("", "not_a_real.file", false)
	assert.NotEqual(err, nil, "passing not existent filepath to CreateFaceEnrollment (should return real error)")
	_, err = myVoiceIt.CreateFaceEnrollment("", "not_a_real.file", true)
	assert.NotEqual(err, nil, "passing not existent filepath to CreateFaceEnrollment (should return real error)")
	_, err = myVoiceIt.VoiceVerification("", "", "", "not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to VoiceVerification(should return real error)")
	_, err = myVoiceIt.VideoVerification("", "", "", "not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to VideoVerification(should return real error)")
	_, err = myVoiceIt.SplitVideoVerification("", "", "", "not_a_real.file", "also_not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to SplitVideoVerification(should return real error)")
	_, err = myVoiceIt.FaceVerification("", "not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to VideoVerification(should return real error)")
	_, err = myVoiceIt.FaceVerification("", "not_a_real.file", false)
	assert.NotEqual(err, nil, "passing not existent filepath to VideoVerification(should return real error)")
	_, err = myVoiceIt.FaceVerification("", "not_a_real.file", true)
	assert.NotEqual(err, nil, "passing not existent filepath to VideoVerification(should return real error)")
	_, err = myVoiceIt.VoiceIdentification("", "en-US", "", "not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to VoiceIdentification(should return real error)")
	_, err = myVoiceIt.VideoIdentification("", "", "", "not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to VideoIdentification(should return real error)")
	_, err = myVoiceIt.SplitVideoIdentification("", "", "", "not_a_real.file", "also_not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to SplitVideoIdentification(should return real error)")
	_, err = myVoiceIt.FaceIdentification("", "not_a_real.file")
	assert.NotEqual(err, nil, "passing not existent filepath to FaceIdentification(should return real error)")
	_, err = myVoiceIt.FaceIdentification("", "not_a_real.file", false)
	assert.NotEqual(err, nil, "passing not existent filepath to FaceIdentification(should return real error)")
	_, err = myVoiceIt.FaceIdentification("", "not_a_real.file", true)
	assert.NotEqual(err, nil, "passing not existent filepath to FaceIdentification(should return real error)")
}

func TestBasics(t *testing.T) {
	assert := assert.New(t)
	myVoiceIt := VoiceIt2{APIKey: os.Getenv("VIAPIKEY"), APIToken: os.Getenv("VIAPITOKEN"), BaseUrl: "https://api.voiceit.io"}

	ret, err := myVoiceIt.CreateUser()
	assert.Equal(err, nil)
	var cu structs.CreateUserReturn
	json.Unmarshal(ret, &cu)
	assert.Equal(201, cu.Status, "CreateUser() message: "+cu.Message)
	assert.Equal("SUCC", cu.ResponseCode, "CreateUser() message: "+cu.Message)
	userId := getUserId(ret)

	ret, err = myVoiceIt.GetAllUsers()
	assert.Equal(err, nil)
	var gau structs.GetAllUsersReturn
	json.Unmarshal(ret, &gau)
	assert.Equal(200, gau.Status, "GetAllUsers() message: "+gau.Message)
	assert.Equal("SUCC", gau.ResponseCode, "GetAllUsers() message: "+gau.Message)
	assert.True(0 < len(gau.Users), "GetAllUsers() message: "+gau.Message)

	ret, err = myVoiceIt.CheckUserExists(userId)
	assert.Equal(err, nil)
	var cue structs.CheckUserExistsReturn
	json.Unmarshal(ret, &cue)
	assert.Equal(200, cue.Status, "CheckUserExists() message: "+cue.Message)
	assert.Equal("SUCC", cue.ResponseCode, "CheckUserExists() message: "+cue.Message)

	ret, err = myVoiceIt.CreateGroup("Sample Group Description")
	assert.Equal(err, nil)
	var cg structs.CreateGroupReturn
	json.Unmarshal(ret, &cg)
	assert.Equal(201, cg.Status, "CreateGroup() message: "+cg.Message)
	assert.Equal("SUCC", cg.ResponseCode, "CreateGroup() message: "+cg.Message)
	groupId := cg.GroupId

	ret, err = myVoiceIt.CheckGroupExists(groupId)
	assert.Equal(err, nil)
	var cge structs.CheckGroupExistsReturn
	json.Unmarshal(ret, &cge)
	assert.Equal(200, cge.Status, "CheckGroupExists() message: "+cge.Message)
	assert.Equal("SUCC", cge.ResponseCode, "CheckGroupExists() message: "+cge.Message)

	ret, err = myVoiceIt.AddUserToGroup(groupId, userId)
	assert.Equal(err, nil)
	var autg structs.AddUserToGroupReturn
	json.Unmarshal(ret, &autg)
	assert.Equal(200, autg.Status, "AddUserToGroup() message: "+autg.Message)
	assert.Equal("SUCC", autg.ResponseCode, "AddUserToGroup() message: "+autg.Message)

	ret, err = myVoiceIt.GetGroup(groupId)
	assert.Equal(err, nil)
	var gg structs.GetGroupsForUserReturn
	json.Unmarshal(ret, &gg)
	assert.Equal(200, gg.Status, "GetGroup() message: "+gg.Message)
	assert.Equal("SUCC", gg.ResponseCode, "GetGroup() message: "+gg.Message)

	ret, err = myVoiceIt.GetAllGroups()
	assert.Equal(err, nil)
	var gag structs.GetAllGroupsReturn
	json.Unmarshal(ret, &gag)
	assert.Equal(200, gag.Status, "GetAllGroups() message: "+gag.Message)
	assert.Equal("SUCC", gag.ResponseCode, "GetAllGroups() message: "+gag.Message)

	ret, err = myVoiceIt.GetGroupsForUser(userId)
	assert.Equal(err, nil)
	var ggfu structs.GetGroupsForUserReturn
	json.Unmarshal(ret, &ggfu)
	assert.Equal(200, ggfu.Status, "GetGroupsForUser() message: "+ggfu.Message)
	assert.Equal("SUCC", ggfu.ResponseCode, "GetGroupsForUser() message: "+ggfu.Message)
	assert.Equal(1, len(ggfu.Groups), "GetGroupsForUser() message: "+ggfu.Message)
	assert.Equal(200, ggfu.Status, "GetGroupsForUser() message: "+ggfu.Message)

	ret, err = myVoiceIt.RemoveUserFromGroup(groupId, userId)
	assert.Equal(err, nil)
	var rufg structs.RemoveUserFromGroupReturn
	json.Unmarshal(ret, &rufg)
	assert.Equal(200, rufg.Status, "RemoveUserFromGroup() message: "+rufg.Message)
	assert.Equal("SUCC", rufg.ResponseCode, "RemoveUserFromGroup() message: "+rufg.Message)

	ret, err = myVoiceIt.CreateUserToken(userId, 3*time.Second)
	assert.Equal(err, nil)
	var cut structs.CreateUserTokenReturn
	json.Unmarshal(ret, &cut)
	assert.Equal(201, cut.Status, "CreateUserToken() message: "+cut.Message)
	assert.Equal("SUCC", cut.ResponseCode, "CreateUserToken() message: "+cut.Message)

	ret, err = myVoiceIt.ExpireUserTokens(userId)
	assert.Equal(err, nil)
	var eutr structs.ExpireUserTokensReturn
	json.Unmarshal(ret, &eutr)
	assert.Equal(201, eutr.Status, "ExpireUserTokens() message: "+eutr.Message)
	assert.Equal("SUCC", eutr.ResponseCode, "ExpireUserTokens() message: "+eutr.Message)

	ret, err = myVoiceIt.DeleteUser(userId)
	assert.Equal(err, nil)
	var du structs.DeleteUserReturn
	json.Unmarshal(ret, &du)
	assert.Equal(200, du.Status, "DeleteUser() message: "+du.Message)
	assert.Equal("SUCC", du.ResponseCode, "DeleteUser() message: "+du.Message)

	ret, err = myVoiceIt.DeleteGroup(groupId)
	assert.Equal(err, nil)
	var dg structs.DeleteGroupReturn
	json.Unmarshal(ret, &dg)
	assert.Equal(200, dg.Status, "DeleteGroup() message: "+dg.Message)
	assert.Equal("SUCC", dg.ResponseCode, "DeleteGroup() message: "+dg.Message)

	ret, err = myVoiceIt.GetPhrases("en-US")
	assert.Equal(err, nil)
	var gp structs.GetPhrasesReturn
	json.Unmarshal(ret, &gp)
	assert.Equal(200, gp.Status, "GetPhrases() message: "+gp.Message)
	assert.Equal("SUCC", gp.ResponseCode, "GetPhrases() message: "+gp.Message)

	myVoiceIt.AddNotificationUrl("https://voiceit.io")
	assert.Equal("?notificationURL=https%3A%2F%2Fvoiceit.io", myVoiceIt.NotificationUrl, nil)
}

// Helper function to download files to disk
func downloadFromUrl(url string) error {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]

	log.Println("Downloading " + url + "...")

	response, err := http.Get(url)
	if err != nil {
		return errors.New(`http.Get("` + url + `") Exception: ` + err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return errors.New(`HTTP GET to "` + url + `" gave a non 200's HTTP Response Code of ` + response.Status)
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.New(`ioutil.ReadAll(response.Body) Exception: ` + err.Error())
	}

	if err := ioutil.WriteFile("./"+fileName, bytes, 0644); err != nil {
		return errors.New(`ioutil.WriteFile("./` + fileName + `", bytes, 0644) Exception: ` + err.Error())
	}

	return nil

}

func TestVideo(t *testing.T) {
	assert := assert.New(t)
	myVoiceIt := VoiceIt2{APIKey: os.Getenv("VIAPIKEY"), APIToken: os.Getenv("VIAPITOKEN"), BaseUrl: "https://api.voiceit.io"}
	ret, err := myVoiceIt.CreateUser()
	assert.Equal(err, nil)
	userId1 := getUserId(ret)
	ret, err = myVoiceIt.CreateUser()
	assert.Equal(err, nil)
	userId2 := getUserId(ret)
	ret, err = myVoiceIt.CreateGroup("Sample Group Description")
	assert.Equal(err, nil)
	groupId := getGroupId(ret)
	myVoiceIt.AddUserToGroup(groupId, userId1)
	myVoiceIt.AddUserToGroup(groupId, userId2)

	// Video Enrollments
	if err := downloadFromUrl("https://drive.voiceit.io/files/videoEnrollmentA1.mov"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./videoEnrollmentA1.mov")
	}

	if err := downloadFromUrl("https://drive.voiceit.io/files/videoEnrollmentA2.mov"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./videoEnrollmentA2.mov")
	}

	if err := downloadFromUrl("https://drive.voiceit.io/files/videoEnrollmentA3.mov"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./videoEnrollmentA3.mov")
	}

	if err := downloadFromUrl("https://drive.voiceit.io/files/videoVerificationA1.mov"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./videoVerificationA1.mov")
	}
	if err := downloadFromUrl("https://drive.voiceit.io/files/videoEnrollmentD1.mov"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./videoEnrollmentD1.mov")
	}
	if err := downloadFromUrl("https://drive.voiceit.io/files/videoEnrollmentD2.mov"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./videoEnrollmentD2.mov")
	}
	if err := downloadFromUrl("https://drive.voiceit.io/files/videoEnrollmentD3.mov"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./videoEnrollmentD3.mov")
	}
	if err := downloadFromUrl("https://drive.voiceit.io/files/faceA1.png"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./faceA1.png")
	}
	if err := downloadFromUrl("https://drive.voiceit.io/files/faceA2.png"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./faceA2.png")
	}
	if err := downloadFromUrl("https://drive.voiceit.io/files/faceA3.png"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./faceA3.png")
	}
	if err := downloadFromUrl("https://drive.voiceit.io/files/enrollmentA1.wav"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./enrollmentA1.wav")
	}
	if err := downloadFromUrl("https://drive.voiceit.io/files/enrollmentA2.wav"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./enrollmentA2.wav")
	}
	if err := downloadFromUrl("https://drive.voiceit.io/files/enrollmentA3.wav"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./enrollmentA3.wav")
	}

	ret, err = myVoiceIt.CreateVideoEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentA1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve1 structs.CreateVideoEnrollmentReturn
	json.Unmarshal(ret, &cve1)
	assert.Equal(201, cve1.Status, "CreateVideoEnrollment() message: "+cve1.Message)
	assert.Equal("SUCC", cve1.ResponseCode, "CreateVideoEnrollment() message: "+cve1.Message)

	ret, err = myVoiceIt.CreateVideoEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentA2.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve2 structs.CreateVideoEnrollmentReturn
	json.Unmarshal(ret, &cve2)
	assert.Equal(201, cve2.Status, "CreateVideoEnrollment() message: "+cve2.Message)
	assert.Equal("SUCC", cve2.ResponseCode, "CreateVideoEnrollment() message: "+cve2.Message)

	ret, err = myVoiceIt.CreateVideoEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentA3.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve3 structs.CreateVideoEnrollmentReturn
	json.Unmarshal(ret, &cve3)
	assert.Equal(201, cve3.Status, "CreateVideoEnrollment() message: "+cve3.Message)
	assert.Equal("SUCC", cve3.ResponseCode, "CreateVideoEnrollment() message: "+cve3.Message)

	ret, err = myVoiceIt.CreateVideoEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentD1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve4 structs.CreateVideoEnrollmentReturn
	json.Unmarshal(ret, &cve4)
	assert.Equal(201, cve4.Status, "CreateVideoEnrollment() message: "+cve4.Message)
	assert.Equal("SUCC", cve4.ResponseCode, "CreateVideoEnrollment() message: "+cve4.Message)

	ret, err = myVoiceIt.CreateVideoEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentD2.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve5 structs.CreateVideoEnrollmentReturn
	json.Unmarshal(ret, &cve5)
	assert.Equal(201, cve5.Status, "CreateVideoEnrollment() message: "+cve5.Message)
	assert.Equal("SUCC", cve5.ResponseCode, "CreateVideoEnrollment() message: "+cve5.Message)

	ret, err = myVoiceIt.CreateVideoEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./videoEnrollmentD3.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve6 structs.CreateVideoEnrollmentReturn
	json.Unmarshal(ret, &cve6)
	assert.Equal(201, cve6.Status, "CreateVideoEnrollment() message: "+cve6.Message)
	assert.Equal("SUCC", cve6.ResponseCode, "CreateVideoEnrollment() message: "+cve6.Message)

	ret, err = myVoiceIt.CreateSplitVideoEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./enrollmentA1.wav", "./faceA1.png")
	if err != nil {
		t.Fatal(err.Error())
	}

	var cve7 structs.CreateVideoEnrollmentReturn
	json.Unmarshal(ret, &cve7)
	assert.Equal(201, cve7.Status, "CreateSplitVideoEnrollment() message: "+cve7.Message)
	assert.Equal("SUCC", cve7.ResponseCode, "CreateSplitVideoEnrollment() message: "+cve7.Message)

	// Get Video Enrollment
	ret, err = myVoiceIt.GetAllVideoEnrollments(userId1)
	assert.Equal(err, nil)
	var gve1 structs.GetAllVideoEnrollmentsReturn
	json.Unmarshal(ret, &gve1)
	assert.Equal(200, gve1.Status, "GetAllVideoEnrollments() message: "+gve1.Message)
	assert.Equal("SUCC", gve1.ResponseCode, "GetAllVideoEnrollments() message: "+gve1.Message)
	assert.Equal(4, len(gve1.VideoEnrollments), "GetAllVideoEnrollments() message: "+gve1.Message)

	// Video Verification
	ret, err = myVoiceIt.VideoVerification(userId1, "en-US", "never forget tomorrow is a new day", "./videoVerificationA1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var vv1 structs.VideoVerificationReturn
	json.Unmarshal(ret, &vv1)
	assert.Equal(200, vv1.Status, "VideoVerification() message: "+vv1.Message)
	assert.Equal("SUCC", vv1.ResponseCode, "VideoVerification() message: "+vv1.Message)

	// Split Video Verification
	ret, err = myVoiceIt.SplitVideoVerification(userId1, "en-US", "never forget tomorrow is a new day", "./enrollmentA2.wav", "./faceA2.png")
	if err != nil {
		t.Fatal(err.Error())
	}
	var vv2 structs.VideoVerificationReturn
	json.Unmarshal(ret, &vv2)
	assert.Equal(200, vv2.Status, "SplitVideoVerification() message: "+vv2.Message)
	assert.Equal("SUCC", vv2.ResponseCode, "SplitVideoVerification() message: "+vv2.Message)

	// Video Identification
	ret, err = myVoiceIt.VideoIdentification(groupId, "en-US", "never forget tomorrow is a new day", "./videoVerificationA1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var vi1 structs.VideoIdentificationReturn
	json.Unmarshal(ret, &vi1)
	assert.Equal(200, vi1.Status, "VideoIdentification() message: "+vi1.Message)
	assert.Equal("SUCC", vi1.ResponseCode, "VideoIdentification() message: "+vi1.Message)
	assert.Equal(userId1, vi1.UserId, "VideoIdentification() message: "+vi1.Message)

	// Split Video Identification
	ret, err = myVoiceIt.SplitVideoIdentification(groupId, "en-US", "never forget tomorrow is a new day", "./enrollmentA3.wav", "./faceA3.png")
	if err != nil {
		t.Fatal(err.Error())
	}
	var vi2 structs.VideoIdentificationReturn
	json.Unmarshal(ret, &vi2)
	assert.Equal(200, vi2.Status, "SplitVideoIdentification() message: "+vi2.Message)
	assert.Equal("SUCC", vi2.ResponseCode, "SplitVideoIdentification() message: "+vi2.Message)
	assert.Equal(userId1, vi2.UserId, "SplitVideoIdentification() message: "+vi2.Message)

	// Delete All Enrollments
	ret, err = myVoiceIt.DeleteAllEnrollments(userId2)
	assert.Equal(err, nil)
	var dae structs.DeleteAllEnrollmentsReturn
	json.Unmarshal(ret, &dae)
	assert.Equal(200, dae.Status, "DeleteAllEnrollments() message: "+dae.Message)
	assert.Equal("SUCC", dae.ResponseCode, "DeleteAllEnrollments() message: "+dae.Message)

	ret, err = myVoiceIt.GetAllVideoEnrollments(userId2)
	assert.Equal(err, nil)
	var gve4 structs.GetAllVideoEnrollmentsReturn
	json.Unmarshal(ret, &gve4)
	assert.Equal(200, gve4.Status, "GetAllVideoEnrollments() message: "+gve4.Message)
	assert.Equal("SUCC", gve4.ResponseCode, "GetAllVideoEnrollments() message: "+gve4.Message)
	assert.Equal(0, len(gve4.VideoEnrollments), "GetAllVideoEnrollments() message: "+gve4.Message)

	// Reset for ...ByUrl calls
	myVoiceIt.DeleteUser(userId1)
	myVoiceIt.DeleteUser(userId2)
	myVoiceIt.DeleteGroup(groupId)
	ret, err = myVoiceIt.CreateUser()
	assert.Equal(err, nil)
	userId1 = getUserId(ret)
	ret, err = myVoiceIt.CreateUser()
	assert.Equal(err, nil)
	userId2 = getUserId(ret)
	ret, err = myVoiceIt.CreateGroup("Sample Group Description")
	assert.Equal(err, nil)
	groupId = getGroupId(ret)
	myVoiceIt.AddUserToGroup(groupId, userId1)
	myVoiceIt.AddUserToGroup(groupId, userId2)

	// Video Enrollments By Url
	ret, err = myVoiceIt.CreateVideoEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://drive.voiceit.io/files/videoEnrollmentA1.mov")
	assert.Equal(err, nil)
	var cvebu1 structs.CreateVideoEnrollmentByUrlReturn
	json.Unmarshal(ret, &cvebu1)
	assert.Equal(201, cvebu1.Status, "CreateVideoEnrollmentByUrl() message: "+cvebu1.Message)
	assert.Equal("SUCC", cvebu1.ResponseCode, "CreateVideoEnrollmentByUrl() message: "+cvebu1.Message)

	ret, err = myVoiceIt.CreateVideoEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://drive.voiceit.io/files/videoEnrollmentA2.mov")
	assert.Equal(err, nil)
	var cvebu2 structs.CreateVideoEnrollmentByUrlReturn
	json.Unmarshal(ret, &cvebu2)
	assert.Equal(201, cvebu2.Status, "CreateVideoEnrollmentByUrl() message: "+cvebu2.Message)
	assert.Equal("SUCC", cvebu2.ResponseCode, "CreateVideoEnrollmentByUrl() message: "+cvebu2.Message)

	ret, err = myVoiceIt.CreateVideoEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://drive.voiceit.io/files/videoEnrollmentA3.mov")
	assert.Equal(err, nil)
	var cvebu3 structs.CreateVideoEnrollmentByUrlReturn
	json.Unmarshal(ret, &cvebu3)
	assert.Equal(201, cvebu3.Status, "CreateVideoEnrollmentByUrl() message: "+cvebu3.Message)
	assert.Equal("SUCC", cvebu3.ResponseCode, "CreateVideoEnrollmentByUrl() message: "+cvebu3.Message)

	ret, err = myVoiceIt.CreateVideoEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://drive.voiceit.io/files/videoEnrollmentD1.mov")
	assert.Equal(err, nil)
	var cvebu4 structs.CreateVideoEnrollmentByUrlReturn
	json.Unmarshal(ret, &cvebu4)
	assert.Equal(201, cvebu4.Status, "CreateVideoEnrollmentByUrl() message: "+cvebu4.Message)
	assert.Equal("SUCC", cvebu4.ResponseCode, "CreateVideoEnrollmentByUrl() message: "+cvebu4.Message)

	ret, err = myVoiceIt.CreateVideoEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://drive.voiceit.io/files/videoEnrollmentD2.mov")
	assert.Equal(err, nil)
	var cvebu5 structs.CreateVideoEnrollmentByUrlReturn
	json.Unmarshal(ret, &cvebu5)
	assert.Equal(201, cvebu5.Status, "CreateVideoEnrollmentByUrl() message: "+cvebu5.Message)
	assert.Equal("SUCC", cvebu5.ResponseCode, "CreateVideoEnrollmentByUrl() message: "+cvebu5.Message)

	ret, err = myVoiceIt.CreateVideoEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://drive.voiceit.io/files/videoEnrollmentD3.mov")
	assert.Equal(err, nil)
	var cvebu6 structs.CreateVideoEnrollmentByUrlReturn
	json.Unmarshal(ret, &cvebu6)
	assert.Equal(201, cvebu6.Status, "CreateVideoEnrollmentByUrl() message: "+cvebu6.Message)
	assert.Equal("SUCC", cvebu6.ResponseCode, "CreateVideoEnrollmentByUrl() message: "+cvebu6.Message)

	// Video Verification
	ret, err = myVoiceIt.VideoVerificationByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://drive.voiceit.io/files/videoVerificationA1.mov")
	assert.Equal(err, nil)
	var vvbu structs.VideoVerificationByUrlReturn
	json.Unmarshal(ret, &vvbu)
	assert.Equal(200, vvbu.Status, "VideoVerificationByUrl() message: "+vvbu.Message)
	assert.Equal("SUCC", vvbu.ResponseCode, "VideoVerificationByUrl() message: "+vvbu.Message)

	// Video Identification
	ret, err = myVoiceIt.VideoIdentificationByUrl(groupId, "en-US", "never forget tomorrow is a new day", "https://drive.voiceit.io/files/videoVerificationA1.mov")
	assert.Equal(err, nil)
	var vibu structs.VideoIdentificationByUrlReturn
	json.Unmarshal(ret, &vibu)
	assert.Equal(200, vibu.Status, "VideoIdentificationByUrl() message: "+vibu.Message)
	assert.Equal("SUCC", vibu.ResponseCode, "VideoIdentificationByUrl() message: "+vibu.Message)
	assert.Equal(userId1, vibu.UserId, "VideoIdentificationByUrl() message: "+vibu.Message)

	myVoiceIt.DeleteAllEnrollments(userId1)
	myVoiceIt.DeleteAllEnrollments(userId2)
	myVoiceIt.DeleteUser(userId1)
	myVoiceIt.DeleteUser(userId2)
	myVoiceIt.DeleteGroup(groupId)

}

func TestVoice(t *testing.T) {
	assert := assert.New(t)
	myVoiceIt := VoiceIt2{APIKey: os.Getenv("VIAPIKEY"), APIToken: os.Getenv("VIAPITOKEN"), BaseUrl: "https://api.voiceit.io"}
	ret, err := myVoiceIt.CreateUser()
	assert.Equal(err, nil)
	userId1 := getUserId(ret)
	ret, err = myVoiceIt.CreateUser()
	assert.Equal(err, nil)
	userId2 := getUserId(ret)
	ret, err = myVoiceIt.CreateGroup("Sample Group Description")
	assert.Equal(err, nil)
	groupId := getGroupId(ret)
	myVoiceIt.AddUserToGroup(groupId, userId1)
	myVoiceIt.AddUserToGroup(groupId, userId2)

	// Voice Enrollments
	if err := downloadFromUrl("https://drive.voiceit.io/files/enrollmentA1.wav"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./enrollmentA1.wav")
	}

	if err := downloadFromUrl("https://drive.voiceit.io/files/enrollmentA2.wav"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./enrollmentA2.wav")
	}

	if err := downloadFromUrl("https://drive.voiceit.io/files/enrollmentA3.wav"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./enrollmentA3.wav")
	}

	if err := downloadFromUrl("https://drive.voiceit.io/files/verificationA1.wav"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./verificationA1.wav")
	}

	if err := downloadFromUrl("https://drive.voiceit.io/files/enrollmentD1.m4a"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./enrollmentD1.m4a")
	}

	if err := downloadFromUrl("https://drive.voiceit.io/files/enrollmentD2.m4a"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./enrollmentD2.m4a")
	}

	if err := downloadFromUrl("https://drive.voiceit.io/files/enrollmentD3.m4a"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./enrollmentD3.m4a")
	}

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./enrollmentA1.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve1 structs.CreateVoiceEnrollmentReturn
	json.Unmarshal(ret, &cve1)
	assert.Equal(201, cve1.Status, "CreateVoiceEnrollment() message: "+cve1.Message)
	assert.Equal("SUCC", cve1.ResponseCode, "CreateVoiceEnrollment() message: "+cve1.Message)

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./enrollmentA2.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve2 structs.CreateVoiceEnrollmentReturn
	json.Unmarshal(ret, &cve2)
	assert.Equal(201, cve2.Status, "CreateVoiceEnrollment() message: "+cve2.Message)
	assert.Equal("SUCC", cve2.ResponseCode, "CreateVoiceEnrollment() message: "+cve2.Message)

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId1, "en-US", "never forget tomorrow is a new day", "./enrollmentA3.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve3 structs.CreateVoiceEnrollmentReturn
	json.Unmarshal(ret, &cve3)
	assert.Equal(201, cve3.Status, "CreateVoiceEnrollment() message: "+cve3.Message)
	assert.Equal("SUCC", cve3.ResponseCode, "CreateVoiceEnrollment() message: "+cve3.Message)

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./enrollmentD1.m4a")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve4 structs.CreateVoiceEnrollmentReturn
	json.Unmarshal(ret, &cve4)
	assert.Equal(201, cve4.Status, "CreateVoiceEnrollment() message: "+cve4.Message)
	assert.Equal("SUCC", cve4.ResponseCode, "CreateVoiceEnrollment() message: "+cve4.Message)

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./enrollmentD2.m4a")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve5 structs.CreateVoiceEnrollmentReturn
	json.Unmarshal(ret, &cve5)
	assert.Equal(201, cve5.Status, "CreateVoiceEnrollment() message: "+cve5.Message)
	assert.Equal("SUCC", cve5.ResponseCode, "CreateVoiceEnrollment() message: "+cve5.Message)

	ret, err = myVoiceIt.CreateVoiceEnrollment(userId2, "en-US", "never forget tomorrow is a new day", "./enrollmentD3.m4a")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cve6 structs.CreateVoiceEnrollmentReturn
	json.Unmarshal(ret, &cve6)
	assert.Equal(201, cve6.Status, "CreateVoiceEnrollment() message: "+cve6.Message)
	assert.Equal("SUCC", cve6.ResponseCode, "CreateVoiceEnrollment() message: "+cve6.Message)

	// Get Voice Enrollment
	ret, err = myVoiceIt.GetAllVoiceEnrollments(userId1)
	assert.Equal(err, nil)
	var gve1 structs.GetAllVoiceEnrollmentsReturn
	json.Unmarshal(ret, &gve1)
	assert.Equal(200, gve1.Status, "GetAllVoiceEnrollments() message: "+gve1.Message)
	assert.Equal("SUCC", gve1.ResponseCode, "GetAllVoiceEnrollments() message: "+gve1.Message)
	assert.Equal(3, len(gve1.VoiceEnrollments), "GetAllVoiceEnrollments() message: "+gve1.Message)

	// Voice Verification
	ret, err = myVoiceIt.VoiceVerification(userId1, "en-US", "never forget tomorrow is a new day", "./verificationA1.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var vv structs.VoiceVerificationReturn
	json.Unmarshal(ret, &vv)
	assert.Equal(200, vv.Status, "VoiceVerification() message: "+vv.Message)
	assert.Equal("SUCC", vv.ResponseCode, "VoiceVerification() message: "+vv.Message)

	// Voice Identification
	ret, err = myVoiceIt.VoiceIdentification(groupId, "en-US", "never forget tomorrow is a new day", "./verificationA1.wav")
	if err != nil {
		t.Fatal(err.Error())
	}
	var vi structs.VoiceIdentificationReturn
	json.Unmarshal(ret, &vi)
	assert.Equal(200, vi.Status, "VoiceIdentification() message: "+vi.Message)
	assert.Equal("SUCC", vi.ResponseCode, "VoiceIdentification() message: "+vi.Message)
	assert.Equal(userId1, vi.UserId, "VoiceIdentification() message: "+vi.Message)

	// Delete All Enrollments
	ret, err = myVoiceIt.DeleteAllEnrollments(userId2)
	assert.Equal(err, nil)
	var dae structs.DeleteAllEnrollmentsReturn
	json.Unmarshal(ret, &dae)
	assert.Equal(200, dae.Status, "DeleteAllEnrollments() message: "+dae.Message)
	assert.Equal("SUCC", dae.ResponseCode, "DeleteAllEnrollments() message: "+dae.Message)

	// Reset for ...ByUrl calls
	myVoiceIt.DeleteUser(userId1)
	myVoiceIt.DeleteUser(userId2)
	myVoiceIt.DeleteGroup(groupId)
	ret, err = myVoiceIt.CreateUser()
	assert.Equal(err, nil)
	userId1 = getUserId(ret)
	ret, err = myVoiceIt.CreateUser()
	assert.Equal(err, nil)
	userId2 = getUserId(ret)
	ret, err = myVoiceIt.CreateGroup("Sample Group Description")
	assert.Equal(err, nil)
	groupId = getGroupId(ret)
	myVoiceIt.AddUserToGroup(groupId, userId1)
	myVoiceIt.AddUserToGroup(groupId, userId2)

	// Voice Enrollments By Url
	ret, err = myVoiceIt.CreateVoiceEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://drive.voiceit.io/files/enrollmentA1.wav")
	assert.Equal(err, nil)
	var cvebu1 structs.CreateVoiceEnrollmentByUrlReturn
	json.Unmarshal(ret, &cvebu1)
	assert.Equal(201, cvebu1.Status, "CreateVoiceEnrollmentByUrl() message: "+cvebu1.Message)
	assert.Equal("SUCC", cvebu1.ResponseCode, "CreateVoiceEnrollmentByUrl() message: "+cvebu1.Message)

	ret, err = myVoiceIt.CreateVoiceEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://drive.voiceit.io/files/enrollmentA2.wav")
	assert.Equal(err, nil)
	var cvebu2 structs.CreateVoiceEnrollmentByUrlReturn
	json.Unmarshal(ret, &cvebu2)
	assert.Equal(201, cvebu2.Status, "CreateVoiceEnrollmentByUrl() message: "+cvebu2.Message)
	assert.Equal("SUCC", cvebu2.ResponseCode, "CreateVoiceEnrollmentByUrl() message: "+cvebu2.Message)

	ret, err = myVoiceIt.CreateVoiceEnrollmentByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://drive.voiceit.io/files/enrollmentA3.wav")
	assert.Equal(err, nil)
	var cvebu3 structs.CreateVoiceEnrollmentByUrlReturn
	json.Unmarshal(ret, &cvebu3)
	assert.Equal(201, cvebu3.Status, "CreateVoiceEnrollmentByUrl() message: "+cvebu3.Message)
	assert.Equal("SUCC", cvebu3.ResponseCode, "CreateVoiceEnrollmentByUrl() message: "+cvebu3.Message)

	ret, err = myVoiceIt.CreateVoiceEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://drive.voiceit.io/files/enrollmentD1.m4a")
	assert.Equal(err, nil)
	var cvebu4 structs.CreateVoiceEnrollmentByUrlReturn
	json.Unmarshal(ret, &cvebu4)
	assert.Equal(201, cvebu4.Status, "CreateVoiceEnrollmentByUrl() message: "+cvebu4.Message)
	assert.Equal("SUCC", cvebu4.ResponseCode, "CreateVoiceEnrollmentByUrl() message: "+cvebu4.Message)

	ret, err = myVoiceIt.CreateVoiceEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://drive.voiceit.io/files/enrollmentD2.m4a")
	assert.Equal(err, nil)
	var cvebu5 structs.CreateVoiceEnrollmentByUrlReturn
	json.Unmarshal(ret, &cvebu5)
	assert.Equal(201, cvebu5.Status, "CreateVoiceEnrollmentByUrl() message: "+cvebu5.Message)
	assert.Equal("SUCC", cvebu5.ResponseCode, "CreateVoiceEnrollmentByUrl() message: "+cvebu5.Message)

	ret, err = myVoiceIt.CreateVoiceEnrollmentByUrl(userId2, "en-US", "never forget tomorrow is a new day", "https://drive.voiceit.io/files/enrollmentD3.m4a")
	assert.Equal(err, nil)
	var cvebu6 structs.CreateVoiceEnrollmentByUrlReturn
	json.Unmarshal(ret, &cvebu6)
	assert.Equal(201, cvebu6.Status, "CreateVoiceEnrollmentByUrl() message: "+cvebu6.Message)
	assert.Equal("SUCC", cvebu6.ResponseCode, "CreateVoiceEnrollmentByUrl() message: "+cvebu6.Message)

	// Voice Verification
	ret, err = myVoiceIt.VoiceVerificationByUrl(userId1, "en-US", "never forget tomorrow is a new day", "https://drive.voiceit.io/files/verificationA1.wav")
	assert.Equal(err, nil)
	var vvbu structs.VoiceVerificationByUrlReturn
	json.Unmarshal(ret, &vvbu)
	assert.Equal(200, vvbu.Status, "VoiceVerificationByUrl() message: "+vvbu.Message)
	assert.Equal("SUCC", vvbu.ResponseCode, "VoiceVerificationByUrl() message: "+vvbu.Message)

	// Voice Identification
	ret, err = myVoiceIt.VoiceIdentificationByUrl(groupId, "en-US", "never forget tomorrow is a new day", "https://drive.voiceit.io/files/verificationA1.wav")
	assert.Equal(err, nil)
	var vibu structs.VoiceIdentificationByUrlReturn
	json.Unmarshal(ret, &vibu)
	assert.Equal(200, vibu.Status, "VoiceIdentificationByUrl() message: "+vibu.Message)
	assert.Equal("SUCC", vibu.ResponseCode, "VoiceIdentificationByUrl() message: "+vibu.Message)
	assert.Equal(userId1, vibu.UserId, "VoiceIdentificationByUrl() message: "+vibu.Message)

	myVoiceIt.DeleteAllEnrollments(userId1)
	myVoiceIt.DeleteAllEnrollments(userId2)
	myVoiceIt.DeleteUser(userId1)
	myVoiceIt.DeleteUser(userId2)
	myVoiceIt.DeleteGroup(groupId)
}

func TestFace(t *testing.T) {
	assert := assert.New(t)
	myVoiceIt := VoiceIt2{APIKey: os.Getenv("VIAPIKEY"), APIToken: os.Getenv("VIAPITOKEN"), BaseUrl: "https://api.voiceit.io"}
	ret, err := myVoiceIt.CreateUser()
	assert.Equal(err, nil)
	userId1 := getUserId(ret)
	ret, err = myVoiceIt.CreateUser()
	assert.Equal(err, nil)
	userId2 := getUserId(ret)
	ret, err = myVoiceIt.CreateGroup("Sample Group Description")
	assert.Equal(err, nil)
	groupId := getGroupId(ret)
	myVoiceIt.AddUserToGroup(groupId, userId1)
	myVoiceIt.AddUserToGroup(groupId, userId2)

	// Face Enrollments
	if err := downloadFromUrl("https://drive.voiceit.io/files/faceEnrollmentA1.mov"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./faceEnrollmentA1.mov")
	}

	if err := downloadFromUrl("https://drive.voiceit.io/files/faceA2.png"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./faceA2.png")
	}

	if err := downloadFromUrl("https://drive.voiceit.io/files/faceEnrollmentA3.mov"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./faceEnrollmentA3.mov")
	}

	if err := downloadFromUrl("https://drive.voiceit.io/files/videoVerificationA1.mov"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./videoVerificationA1.mov")
	}

	if err := downloadFromUrl("https://drive.voiceit.io/files/videoEnrollmentD1.mov"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./videoEnrollmentD1.mov")
	}

	if err := downloadFromUrl("https://drive.voiceit.io/files/videoEnrollmentD2.mov"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./videoEnrollmentD2.mov")
	}

	if err := downloadFromUrl("https://drive.voiceit.io/files/videoEnrollmentD3.mov"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./videoEnrollmentD3.mov")
	}

	if err := downloadFromUrl("https://drive.voiceit.io/files/faceA1.png"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./faceA1.png")
	}

	if err := downloadFromUrl("https://drive.voiceit.io/files/faceA3.png"); err != nil {
		t.Fatal(err)
	} else {
		defer os.Remove("./faceA3.png")
	}

	ret, err = myVoiceIt.CreateFaceEnrollment(userId1, "./faceEnrollmentA1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cfe1 structs.CreateFaceEnrollmentReturn
	json.Unmarshal(ret, &cfe1)
	assert.Equal(201, cfe1.Status, "CreateFaceEnrollment() message: "+cfe1.Message)
	assert.Equal("SUCC", cfe1.ResponseCode, "CreateFaceEnrollment() message: "+cfe1.Message)

	ret, err = myVoiceIt.CreateFaceEnrollment(userId1, "./faceA2.png", true)
	if err != nil {
		t.Fatal(err.Error())
	}
	var cfe2 structs.CreateFaceEnrollmentReturn
	json.Unmarshal(ret, &cfe2)
	assert.Equal(201, cfe2.Status, "CreateFaceEnrollment() message: "+cfe2.Message)
	assert.Equal("SUCC", cfe2.ResponseCode, "CreateFaceEnrollment() message: "+cfe2.Message)

	ret, err = myVoiceIt.CreateFaceEnrollment(userId1, "./faceEnrollmentA3.mov", false)
	if err != nil {
		t.Fatal(err.Error())
	}
	var cfe3 structs.CreateFaceEnrollmentReturn
	json.Unmarshal(ret, &cfe3)
	assert.Equal(201, cfe3.Status, "CreateFaceEnrollment() message: "+cfe3.Message)
	assert.Equal("SUCC", cfe3.ResponseCode, "CreateFaceEnrollment() message: "+cfe3.Message)

	ret, err = myVoiceIt.CreateFaceEnrollment(userId2, "./videoEnrollmentD1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cfe4 structs.CreateFaceEnrollmentReturn
	json.Unmarshal(ret, &cfe4)
	assert.Equal(201, cfe4.Status, "CreateFaceEnrollment() message: "+cfe4.Message)
	assert.Equal("SUCC", cfe4.ResponseCode, "CreateFaceEnrollment() message: "+cfe4.Message)

	ret, err = myVoiceIt.CreateFaceEnrollment(userId2, "./videoEnrollmentD1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cfe5 structs.CreateFaceEnrollmentReturn
	json.Unmarshal(ret, &cfe5)
	assert.Equal(201, cfe5.Status, "CreateFaceEnrollment() message: "+cfe5.Message)
	assert.Equal("SUCC", cfe5.ResponseCode, "CreateFaceEnrollment() message: "+cfe5.Message)

	ret, err = myVoiceIt.CreateFaceEnrollment(userId2, "./videoEnrollmentD1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var cfe6 structs.CreateFaceEnrollmentReturn
	json.Unmarshal(ret, &cfe6)
	assert.Equal(201, cfe6.Status, "CreateFaceEnrollment() message: "+cfe6.Message)
	assert.Equal("SUCC", cfe6.ResponseCode, "CreateFaceEnrollment() message: "+cfe6.Message)

	// Get Face Enrollment
	ret, err = myVoiceIt.GetAllFaceEnrollments(userId1)
	assert.Equal(err, nil)
	var fve1 structs.GetAllFaceEnrollmentsReturn
	json.Unmarshal(ret, &fve1)
	assert.Equal(200, fve1.Status, "GetAllFaceEnrollments() message: "+fve1.Message)
	assert.Equal("SUCC", fve1.ResponseCode, "GetAllFaceEnrollments() message: "+fve1.Message)
	assert.Equal(3, len(fve1.FaceEnrollments), "GetAllFaceEnrollments() message: "+fve1.Message)

	// Face Verification
	ret, err = myVoiceIt.FaceVerification(userId1, "./videoVerificationA1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var fv1 structs.FaceVerificationReturn
	json.Unmarshal(ret, &fv1)
	assert.Equal(200, fv1.Status, "FaceVerification() message: "+fv1.Message)
	assert.Equal("SUCC", fv1.ResponseCode, "FaceVerification() message: "+fv1.Message)

	// Split Face Verification
	ret, err = myVoiceIt.FaceVerification(userId1, "./faceA1.png", true)
	if err != nil {
		t.Fatal(err.Error())
	}
	var fv2 structs.FaceVerificationReturn
	json.Unmarshal(ret, &fv2)
	assert.Equal(200, fv2.Status, "FaceVerification() message: "+fv2.Message)
	assert.Equal("SUCC", fv2.ResponseCode, "FaceVerification() message: "+fv2.Message)

	// Face Identification
	ret, err = myVoiceIt.FaceIdentification(groupId, "./videoVerificationA1.mov")
	if err != nil {
		t.Fatal(err.Error())
	}
	var fi1 structs.FaceIdentificationReturn
	json.Unmarshal(ret, &fi1)
	assert.Equal(200, fi1.Status, "FaceIdentification() message: "+fi1.Message)
	assert.Equal("SUCC", fi1.ResponseCode, "FaceIdentification() message: "+fi1.Message)
	assert.Equal(userId1, fi1.UserId, "FaceIdentification() message: "+fi1.Message)

	// Split Face Identification
	ret, err = myVoiceIt.FaceIdentification(groupId, "./faceA3.png", true)
	if err != nil {
		t.Fatal(err.Error())
	}
	var fi2 structs.FaceIdentificationReturn
	json.Unmarshal(ret, &fi2)
	assert.Equal(200, fi2.Status, "FaceIdentification() message: "+fi2.Message)
	assert.Equal("SUCC", fi2.ResponseCode, "FaceIdentification() message: "+fi2.Message)
	assert.Equal(userId1, fi2.UserId, "FaceIdentification() message: "+fi2.Message)

	_, err = myVoiceIt.GetAllFaceEnrollments(userId1)
	assert.Equal(err, nil)

	// Delete All Enrollments
	ret, err = myVoiceIt.DeleteAllEnrollments(userId2)
	assert.Equal(err, nil)
	var dae structs.DeleteAllEnrollmentsReturn
	json.Unmarshal(ret, &dae)
	assert.Equal(200, dae.Status, "DeleteAllEnrollments() message: "+dae.Message)
	assert.Equal("SUCC", dae.ResponseCode, "DeleteAllEnrollments() message: "+dae.Message)

	// Reset for ...ByUrl calls
	myVoiceIt.DeleteUser(userId1)
	myVoiceIt.DeleteUser(userId2)
	myVoiceIt.DeleteGroup(groupId)
	ret, err = myVoiceIt.CreateUser()
	assert.Equal(err, nil)
	userId1 = getUserId(ret)
	ret, err = myVoiceIt.CreateUser()
	assert.Equal(err, nil)
	userId2 = getUserId(ret)
	ret, err = myVoiceIt.CreateGroup("Sample Group Description")
	assert.Equal(err, nil)
	groupId = getGroupId(ret)
	myVoiceIt.AddUserToGroup(groupId, userId1)
	myVoiceIt.AddUserToGroup(groupId, userId2)

	// Face Enrollments By Url
	ret, err = myVoiceIt.CreateFaceEnrollmentByUrl(userId1, "https://drive.voiceit.io/files/faceEnrollmentA1.mov")
	assert.Equal(err, nil)
	var cfebu1 structs.CreateFaceEnrollmentByUrlReturn
	json.Unmarshal(ret, &cfebu1)
	assert.Equal(201, cfebu1.Status, "CreateFaceEnrollmentByUrl() message: "+cfebu1.Message)
	assert.Equal("SUCC", cfebu1.ResponseCode, "CreateFaceEnrollmentByUrl() message: "+cfebu1.Message)

	ret, err = myVoiceIt.CreateFaceEnrollmentByUrl(userId1, "https://drive.voiceit.io/files/faceEnrollmentA2.mov")
	assert.Equal(err, nil)
	var cfebu2 structs.CreateFaceEnrollmentByUrlReturn
	json.Unmarshal(ret, &cfebu2)
	assert.Equal(201, cfebu2.Status, "CreateFaceEnrollmentByUrl() message: "+cfebu2.Message)
	assert.Equal("SUCC", cfebu2.ResponseCode, "CreateFaceEnrollmentByUrl() message: "+cfebu2.Message)

	ret, err = myVoiceIt.CreateFaceEnrollmentByUrl(userId1, "https://drive.voiceit.io/files/faceEnrollmentA3.mov")
	assert.Equal(err, nil)
	var cfebu3 structs.CreateFaceEnrollmentByUrlReturn
	json.Unmarshal(ret, &cfebu3)
	assert.Equal(201, cfebu3.Status, "CreateFaceEnrollmentByUrl() message: "+cfebu3.Message)
	assert.Equal("SUCC", cfebu3.ResponseCode, "CreateFaceEnrollmentByUrl() message: "+cfebu3.Message)

	ret, err = myVoiceIt.CreateFaceEnrollmentByUrl(userId2, "https://drive.voiceit.io/files/videoEnrollmentD1.mov")
	assert.Equal(err, nil)
	var cfebu4 structs.CreateFaceEnrollmentByUrlReturn
	json.Unmarshal(ret, &cfebu4)
	assert.Equal(201, cfebu4.Status, "CreateFaceEnrollmentByUrl() message: "+cfebu4.Message)
	assert.Equal("SUCC", cfebu4.ResponseCode, "CreateFaceEnrollmentByUrl() message: "+cfebu4.Message)

	ret, err = myVoiceIt.CreateFaceEnrollmentByUrl(userId2, "https://drive.voiceit.io/files/videoEnrollmentD2.mov")
	assert.Equal(err, nil)
	var cfebu5 structs.CreateFaceEnrollmentByUrlReturn
	json.Unmarshal(ret, &cfebu5)
	assert.Equal(201, cfebu5.Status, "CreateFaceEnrollmentByUrl() message: "+cfebu5.Message)
	assert.Equal("SUCC", cfebu5.ResponseCode, "CreateFaceEnrollmentByUrl() message: "+cfebu5.Message)

	ret, err = myVoiceIt.CreateFaceEnrollmentByUrl(userId2, "https://drive.voiceit.io/files/videoEnrollmentD3.mov")
	assert.Equal(err, nil)
	var cfebu6 structs.CreateFaceEnrollmentByUrlReturn
	json.Unmarshal(ret, &cfebu6)
	assert.Equal(201, cfebu6.Status, "CreateFaceEnrollmentByUrl() message: "+cfebu6.Message)
	assert.Equal("SUCC", cfebu6.ResponseCode, "CreateFaceEnrollmentByUrl() message: "+cfebu6.Message)

	// Face Verification
	ret, err = myVoiceIt.FaceVerificationByUrl(userId1, "https://drive.voiceit.io/files/videoVerificationA1.mov")
	assert.Equal(err, nil)
	var fvbu structs.FaceVerificationByUrlReturn
	json.Unmarshal(ret, &fvbu)
	assert.Equal(200, fvbu.Status, "FaceVerificationByUrl() message: "+fvbu.Message)
	assert.Equal("SUCC", fvbu.ResponseCode, "FaceVerificationByUrl() message: "+fvbu.Message)

	// Face Identification
	ret, err = myVoiceIt.FaceIdentificationByUrl(groupId, "https://drive.voiceit.io/files/videoVerificationA1.mov")
	assert.Equal(err, nil)
	var fibu structs.FaceIdentificationByUrlReturn
	json.Unmarshal(ret, &fibu)
	assert.Equal(200, fibu.Status, "FaceIdentificationByUrl() message: "+fibu.Message)
	assert.Equal("SUCC", fibu.ResponseCode, "FaceIdentificationByUrl() message: "+fibu.Message)
	assert.Equal(userId1, fibu.UserId, "FaceIdentificationByUrl() message: "+fibu.Message)

	myVoiceIt.DeleteAllEnrollments(userId1)
	myVoiceIt.DeleteAllEnrollments(userId2)
	myVoiceIt.DeleteUser(userId1)
	myVoiceIt.DeleteUser(userId2)
	myVoiceIt.DeleteGroup(groupId)
}

func TestSubAccounts(t *testing.T) {
	assert := assert.New(t)
	myVoiceIt := VoiceIt2{APIKey: os.Getenv("VIAPIKEY"), APIToken: os.Getenv("VIAPITOKEN"), BaseUrl: "https://api.voiceit.io"}

	// Managed

	ret, err := myVoiceIt.CreateManagedSubAccount(structs.CreateSubAccountRequest{})
	assert.Equal(err, nil)
	var csa structs.CreateSubAccountReturn
	json.Unmarshal(ret, &csa)
	assert.Equal("SUCC", csa.ResponseCode)
	assert.Equal(201, csa.Status)

	managed := VoiceIt2{APIKey: csa.APIKey, APIToken: csa.APIToken, BaseUrl: "https://api.voiceit.io"}
	ret, err = managed.CreateUser()
	assert.Equal(err, nil)
	var cu structs.CreateUserReturn
	json.Unmarshal(ret, &cu)
	assert.Equal("SUCC", cu.ResponseCode)
	assert.Equal(201, cu.Status)

	ret, err = myVoiceIt.RegenerateSubAccountAPIToken(managed.APIKey)
	assert.Equal(err, nil)
	var reg structs.RegenerateSubAccountAPITokenReturn
	json.Unmarshal(ret, &reg)
	assert.Equal("SUCC", reg.ResponseCode)
	assert.Equal(200, reg.Status)

	managed.APIToken = reg.APIToken

	ret, err = managed.CreateUser()
	assert.Equal(err, nil)
	json.Unmarshal(ret, &cu)
	assert.Equal("SUCC", cu.ResponseCode)
	assert.Equal(201, cu.Status)

	ret, err = myVoiceIt.SwitchSubAccountType(managed.APIKey)
	assert.Equal(err, nil)
	var csat structs.SwitchSubAccountTypeReturn
	json.Unmarshal(ret, &csat)
	assert.Equal("SUCC", csat.ResponseCode)
	assert.Equal(200, csat.Status)
	assert.Equal("unmanaged", csat.Type)

	ret, err = myVoiceIt.DeleteSubAccount(managed.APIKey)
	assert.Equal(err, nil)
	var dsa structs.DeleteSubAccountReturn
	json.Unmarshal(ret, &dsa)
	assert.Equal("SUCC", dsa.ResponseCode)
	assert.Equal(200, dsa.Status)

	// Unmanaged

	ret, err = myVoiceIt.CreateUnmanagedSubAccount(structs.CreateSubAccountRequest{})
	assert.Equal(err, nil)
	json.Unmarshal(ret, &csa)
	assert.Equal("SUCC", csa.ResponseCode)
	assert.Equal(201, csa.Status)

	unmanaged := VoiceIt2{APIKey: csa.APIKey, APIToken: csa.APIToken, BaseUrl: "https://api.voiceit.io"}
	ret, err = unmanaged.CreateUser()
	assert.Equal(err, nil)
	json.Unmarshal(ret, &cu)
	assert.Equal("SUCC", cu.ResponseCode)
	assert.Equal(201, cu.Status)

	ret, err = myVoiceIt.RegenerateSubAccountAPIToken(unmanaged.APIKey)
	assert.Equal(err, nil)
	json.Unmarshal(ret, &reg)
	assert.Equal("SUCC", reg.ResponseCode)
	assert.Equal(200, reg.Status)

	unmanaged.APIToken = reg.APIToken

	ret, err = unmanaged.CreateUser()
	assert.Equal(err, nil)
	json.Unmarshal(ret, &cu)
	assert.Equal("SUCC", cu.ResponseCode)
	assert.Equal(201, cu.Status)

	ret, err = myVoiceIt.SwitchSubAccountType(unmanaged.APIKey)
	assert.Equal(err, nil)
	json.Unmarshal(ret, &csat)
	assert.Equal("SUCC", csat.ResponseCode)
	assert.Equal(200, csat.Status)
	assert.Equal("managed", csat.Type)

	ret, err = myVoiceIt.DeleteSubAccount(unmanaged.APIKey)
	assert.Equal(err, nil)
	json.Unmarshal(ret, &dsa)
	assert.Equal("SUCC", dsa.ResponseCode)
	assert.Equal(200, dsa.Status)

}
