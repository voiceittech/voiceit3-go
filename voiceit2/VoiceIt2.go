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
    apiKey string
    apiToken string
    baseURL string
}

func New(key string, tok string) *VoiceIt2 {
    return &VoiceIt2{
        apiKey: key,
        apiToken: tok,
        baseURL : "https://api.voiceit.io",
    }
}

func (vi *VoiceIt2) GetAllUsers()string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", vi.baseURL + "/users", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) CreateUser()string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("POST", vi.baseURL + "/users", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) GetUser(userId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", vi.baseURL + "/users/" + userId, body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) DeleteUser(userId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("DELETE", vi.baseURL + "/users/" + userId, body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) GetGroupsForUser(userId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", vi.baseURL + "/users/" + userId + "/groups", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) GetAllEnrollmentsForUser(userId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", vi.baseURL + "/enrollments/" + userId, body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) DeleteAllEnrollmentsForUser(userId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("DELETE", vi.baseURL + "/enrollments/" + userId + "/all", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) GetFaceFaceEnrollmentsForUser(userId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", vi.baseURL + "/enrollments/face/" + userId, body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
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

  req, _ := http.NewRequest("POST", vi.baseURL + "/enrollments", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) CreateVoiceEnrollmentByUrl(userId string, contentLanguage string, fileUrl string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("userId", userId)
  _ = writer.WriteField("contentLanguage", contentLanguage)
  _ = writer.WriteField("fileUrl", fileUrl)
  writer.Close()

  req, _ := http.NewRequest("POST", vi.baseURL + "/enrollments/byUrl", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
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

  req, _ := http.NewRequest("POST", vi.baseURL + "/enrollments/video", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
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

  req, _ := http.NewRequest("POST", vi.baseURL + "/enrollments/video/byUrl", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
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

  req, _ := http.NewRequest("POST", vi.baseURL + "/enrollments/face", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) DeleteFaceEnrollment(userId string, faceEnrollmentId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("DELETE", vi.baseURL + "/enrollments/face/" + userId + "/" + faceEnrollmentId, body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) DeleteEnrollmentForUser(userId string, faceEnrollmentId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("DELETE", vi.baseURL + "/enrollments/" + userId + "/" + faceEnrollmentId, body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) GetAllGroups()string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", vi.baseURL + "/groups", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) GetGroup(groupId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", vi.baseURL + "/groups/" + groupId, body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) GroupExists(groupId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", vi.baseURL + "/groups/" + groupId + "/exists", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) CreateGroup(description string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("description", description)

  writer.Close()

  req, _ := http.NewRequest("POST", vi.baseURL + "/groups", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) AddUserToGroup(groupId string, userId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("groupId", groupId)
  _ = writer.WriteField("userId", userId)

  writer.Close()

  req, _ := http.NewRequest("PUT", vi.baseURL + "/groups/addUser", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) RemoveUserToGroup(groupId string, userId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("groupId", groupId)
  _ = writer.WriteField("userId", userId)

  writer.Close()

  req, _ := http.NewRequest("PUT", vi.baseURL + "/groups/removeUser", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) DeleteGroup(groupId string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("DELETE", vi.baseURL + "/groups/" + groupId, body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
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

  req, _ := http.NewRequest("POST", vi.baseURL + "/verification", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) VoiceVerificationByUrl(userId string, contentLanguage string, fileUrl string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("userId", userId)
  _ = writer.WriteField("contentLanguage", contentLanguage)
  _ = writer.WriteField("fileUrl", fileUrl)
  writer.Close()

  req, _ := http.NewRequest("POST", vi.baseURL + "/verification/byUrl", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
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

  req, _ := http.NewRequest("POST", vi.baseURL + "/verification/face", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
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

  req, _ := http.NewRequest("POST", vi.baseURL + "/verification/video", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
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

  req, _ := http.NewRequest("POST", vi.baseURL + "/verification/video/byUrl", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
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

  req, _ := http.NewRequest("POST", vi.baseURL + "/identification", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}

func (vi *VoiceIt2) VoiceIdentificationByUrl(groupId string, contentLanguage string, fileUrl string) string {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("fileUrl", fileUrl)
  _ = writer.WriteField("groupId", groupId)
  _ = writer.WriteField("contentLanguage", contentLanguage)
  writer.Close()

  req, _ := http.NewRequest("POST", vi.baseURL + "/identification/byUrl", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
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

  req, _ := http.NewRequest("POST", vi.baseURL + "/identification/video", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
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

  req, _ := http.NewRequest("POST", vi.baseURL + "/identification/video/byUrl", body)
  req.SetBasicAuth(vi.apiKey, vi.apiToken)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  reply, _ := ioutil.ReadAll(resp.Body)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)
  result := string(reply[:len(reply)])
  return result
}
