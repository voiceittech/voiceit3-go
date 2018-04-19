package voiceit2

import (
  "net/http"
  "os"
  "bytes"
  "path/filepath"
  "mime/multipart"
  "io"
  "io/ioutil"
  "strconv"
)

type VoiceIt2 struct {
    ApiKey string
    ApiToken string
    BaseUrl string
}

func NewClient(key string, tok string) *VoiceIt2 {
    return &VoiceIt2{
        ApiKey: key,
        ApiToken: tok,
        BaseUrl : "https://api.voiceit.io",
    }
}

func (vi *VoiceIt2) GetAllUsers()string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", vi.BaseUrl + "/users", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) CreateUser()string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("POST", vi.BaseUrl + "/users", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) GetUser(userId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", vi.BaseUrl + "/users/" + userId, body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) DeleteUser(userId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("DELETE", vi.BaseUrl + "/users/" + userId, body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) GetGroupsForUser(userId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", vi.BaseUrl + "/users/" + userId + "/groups", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) GetAllEnrollmentsForUser(userId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", vi.BaseUrl + "/enrollments/" + userId, body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) DeleteAllEnrollmentsForUser(userId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("DELETE", vi.BaseUrl + "/enrollments/" + userId + "/all", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) GetFaceFaceEnrollmentsForUser(userId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", vi.BaseUrl + "/enrollments/face/" + userId, body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

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

  req, _ := http.NewRequest("POST", vi.BaseUrl + "/enrollments", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) CreateVoiceEnrollmentByUrl(userId string, contentLanguage string, fileUrl string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("userId", userId)
  _ = writer.WriteField("contentLanguage", contentLanguage)
  _ = writer.WriteField("fileUrl", fileUrl)
  writer.Close()

  req, _ := http.NewRequest("POST", vi.BaseUrl + "/enrollments/byUrl", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) CreateVideoEnrollment(userId string, contentLanguage string, filePath string, doBlinkDetection ... bool)string {

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

  req, _ := http.NewRequest("POST", vi.BaseUrl + "/enrollments/video", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) CreateVideoEnrollmentByUrl(userId string, contentLanguage string, fileUrl string, doBlinkDetection ... bool)string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("userId", userId)
  _ = writer.WriteField("contentLanguage", contentLanguage)
  _ = writer.WriteField("fileUrl", fileUrl)
  if len(doBlinkDetection) > 0 {
    _ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
  }
  writer.Close()

  req, _ := http.NewRequest("POST", vi.BaseUrl + "/enrollments/video/byUrl", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) CreateFaceEnrollment(userId string, contentLanguage string, filePath string, doBlinkDetection ... bool)string {

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

  req, _ := http.NewRequest("POST", vi.BaseUrl + "/enrollments/face", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) DeleteFaceEnrollment(userId string, faceEnrollmentId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("DELETE", vi.BaseUrl + "/enrollments/face/" + userId + "/" + faceEnrollmentId, body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) DeleteEnrollmentForUser(userId string, faceEnrollmentId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("DELETE", vi.BaseUrl + "/enrollments/" + userId + "/" + faceEnrollmentId, body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) GetAllGroups()string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", vi.BaseUrl + "/groups", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) GetGroup(groupId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", vi.BaseUrl + "/groups/" + groupId, body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) GroupExists(groupId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", vi.BaseUrl + "/groups/" + groupId + "/exists", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) CreateGroup(description string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("description", description)

  writer.Close()

  req, _ := http.NewRequest("POST", vi.BaseUrl + "/groups", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) AddUserToGroup(groupId string, userId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("groupId", groupId)
  _ = writer.WriteField("userId", userId)

  writer.Close()

  req, _ := http.NewRequest("PUT", vi.BaseUrl + "/groups/addUser", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) RemoveUserFromGroup(groupId string, userId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("groupId", groupId)
  _ = writer.WriteField("userId", userId)

  writer.Close()

  req, _ := http.NewRequest("PUT", vi.BaseUrl + "/groups/removeUser", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) DeleteGroup(groupId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("DELETE", vi.BaseUrl + "/groups/" + groupId, body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

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

  req, _ := http.NewRequest("POST", vi.BaseUrl + "/verification", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) VoiceVerificationByUrl(userId string, contentLanguage string, fileUrl string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("userId", userId)
  _ = writer.WriteField("contentLanguage", contentLanguage)
  _ = writer.WriteField("fileUrl", fileUrl)
  writer.Close()

  req, _ := http.NewRequest("POST", vi.BaseUrl + "/verification/byUrl", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) FaceVerification(userId string, filePath string, doBlinkDetection ... bool)string {

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

  req, _ := http.NewRequest("POST", vi.BaseUrl + "/verification/face", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) VideoVerification(userId string, contentLanguage string, filePath string, doBlinkDetection ... bool)string {

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

  req, _ := http.NewRequest("POST", vi.BaseUrl + "/verification/video", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) VideoVerificationByUrl(userId string, contentLanguage string, fileUrl string, doBlinkDetection ... bool)string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("userId", userId)
  _ = writer.WriteField("contentLanguage", contentLanguage)
  _ = writer.WriteField("fileUrl", fileUrl)
  if len(doBlinkDetection) > 0 {
    _ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
  }
  writer.Close()

  req, _ := http.NewRequest("POST", vi.BaseUrl + "/verification/video/byUrl", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

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

  req, _ := http.NewRequest("POST", vi.BaseUrl + "/identification", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) VoiceIdentificationByUrl(groupId string, contentLanguage string, fileUrl string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("fileUrl", fileUrl)
  _ = writer.WriteField("groupId", groupId)
  _ = writer.WriteField("contentLanguage", contentLanguage)
  writer.Close()

  req, _ := http.NewRequest("POST", vi.BaseUrl + "/identification/byUrl", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) VideoIdentification(groupId string, contentLanguage string, filePath string, doBlinkDetection ... bool)string {

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

  req, _ := http.NewRequest("POST", vi.BaseUrl + "/identification/video", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}

func (vi *VoiceIt2) VideoIdentificationByUrl(groupId string, contentLanguage string, fileUrl string, doBlinkDetection ... bool)string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("fileUrl", fileUrl)
  _ = writer.WriteField("groupId", groupId)
  _ = writer.WriteField("contentLanguage", contentLanguage)
  if len(doBlinkDetection) > 0 {
    _ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
  }
  writer.Close()

  req, _ := http.NewRequest("POST", vi.BaseUrl + "/identification/video/byUrl", body)
  req.SetBasicAuth(vi.ApiKey, vi.ApiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  return string(reply)
}
