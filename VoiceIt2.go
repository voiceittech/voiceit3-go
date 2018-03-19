package VoiceIt2

import (
  "net/http"
  "os"
  "bytes"
  "path/filepath"
  "mime/multipart"
  "io"
  "strconv"
)

var api_key string
var api_token string
var BASE_URL string = "https://api.voiceit.io"

func SetAPIKeyAndToken(key string, tok string) {
  api_key = key
  api_token = tok
}

func GetAllUsers() *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", BASE_URL + "/users", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func CreateUser() *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("POST", BASE_URL + "/users", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func GetUser(userId string) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", BASE_URL + "/users/" + userId, body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func DeleteUser(userId string) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("DELETE", BASE_URL + "/users/" + userId, body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func GetGroupsForUser(userId string) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", BASE_URL + "/users/" + userId + "/groups", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func GetAllEnrollmentsForUser(userId string) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", BASE_URL + "/enrollments/" + userId, body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func DeleteAllEnrollmentsForUser(userId string) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("DELETE", BASE_URL + "/enrollments/" + userId + "/all", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func GetFaceFaceEnrollmentsForUser(userId string) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", BASE_URL + "/enrollments/face/" + userId, body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func CreateVoiceEnrollment(userId string, contentLanguage string, filePath string) *bytes.Buffer {

  file, _ := os.Open(filePath)
  defer file.Close()

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  part, _ := writer.CreateFormFile("recording", filepath.Base(file.Name()))
  io.Copy(part, file)
  _ = writer.WriteField("userId", userId)
  _ = writer.WriteField("contentLanguage", contentLanguage)
  writer.Close()

  req, _ := http.NewRequest("POST", BASE_URL + "/enrollments", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func CreateVoiceEnrollmentByUrl(userId string, contentLanguage string, fileUrl string) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("userId", userId)
  _ = writer.WriteField("contentLanguage", contentLanguage)
  _ = writer.WriteField("fileUrl", fileUrl)
  writer.Close()

  req, _ := http.NewRequest("POST", BASE_URL + "/enrollments/byUrl", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func CreateFaceEnrollment(userId string, contentLanguage string, filePath string, doBlinkDetection ... bool) *bytes.Buffer {

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

  req, _ := http.NewRequest("POST", BASE_URL + "/enrollments/face", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func CreateVideoEnrollment(userId string, contentLanguage string, filePath string, doBlinkDetection ... bool) *bytes.Buffer {

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

  req, _ := http.NewRequest("POST", BASE_URL + "/enrollments/video", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func CreateVideoEnrollmentByUrl(userId string, contentLanguage string, fileUrl string, doBlinkDetection ... bool) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("userId", userId)
  _ = writer.WriteField("contentLanguage", contentLanguage)
  _ = writer.WriteField("fileUrl", fileUrl)
  if len(doBlinkDetection) > 0 {
    _ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
  }
  writer.Close()

  req, _ := http.NewRequest("POST", BASE_URL + "/enrollments/video/byUrl", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func DeleteFaceEnrollment(userId string, faceEnrollmentId string) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("DELETE", BASE_URL + "/enrollments/face/" + userId + "/" + faceEnrollmentId, body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func DeleteEnrollmentForUser(userId string, faceEnrollmentId string) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("DELETE", BASE_URL + "/enrollments/" + userId + "/" + faceEnrollmentId, body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func GetAllGroups() *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", BASE_URL + "/groups", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func GetGroup(groupId string) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", BASE_URL + "/groups/" + groupId, body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func GroupExists(groupId string) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("GET", BASE_URL + "/groups/" + groupId + "/exists", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func CreateGroup(description string) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("description", description)

  writer.Close()

  req, _ := http.NewRequest("POST", BASE_URL + "/groups", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func AddUserToGroup(groupId string, userId string) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("groupId", groupId)
  _ = writer.WriteField("userId", userId)

  writer.Close()

  req, _ := http.NewRequest("PUT", BASE_URL + "/groups/addUser", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func RemoveUserToGroup(groupId string, userId string) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("groupId", groupId)
  _ = writer.WriteField("userId", userId)

  writer.Close()

  req, _ := http.NewRequest("PUT", BASE_URL + "/groups/removeUser", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func DeleteGroup(groupId string) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  writer.Close()

  req, _ := http.NewRequest("DELETE", BASE_URL + "/groups/" + groupId, body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func VoiceVerification(userId string, contentLanguage string, filePath string) *bytes.Buffer {

  file, _ := os.Open(filePath)
  defer file.Close()

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  part, _ := writer.CreateFormFile("recording", filepath.Base(file.Name()))
  io.Copy(part, file)
  _ = writer.WriteField("userId", userId)
  _ = writer.WriteField("contentLanguage", contentLanguage)
  writer.Close()

  req, _ := http.NewRequest("POST", BASE_URL + "/verification", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func VoiceVerificationByUrl(userId string, contentLanguage string, fileUrl string) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("userId", userId)
  _ = writer.WriteField("contentLanguage", contentLanguage)
  _ = writer.WriteField("fileUrl", fileUrl)
  writer.Close()

  req, _ := http.NewRequest("POST", BASE_URL + "/verification/byUrl", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func FaceVerification(userId string, filePath string, doBlinkDetection ... bool) *bytes.Buffer {

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

  req, _ := http.NewRequest("POST", BASE_URL + "/verification/face", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func VideoVerification(userId string, contentLanguage string, filePath string, doBlinkDetection ... bool) *bytes.Buffer {

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

  req, _ := http.NewRequest("POST", BASE_URL + "/verification/video", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func VideoVerificationByUrl(userId string, contentLanguage string, fileUrl string, doBlinkDetection ... bool) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("userId", userId)
  _ = writer.WriteField("contentLanguage", contentLanguage)
  _ = writer.WriteField("fileUrl", fileUrl)
  if len(doBlinkDetection) > 0 {
    _ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
  }
  writer.Close()

  req, _ := http.NewRequest("POST", BASE_URL + "/verification/video/byUrl", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func VoiceIdentification(groupId string, contentLanguage string, filePath string) *bytes.Buffer {

  file, _ := os.Open(filePath)
  defer file.Close()

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  part, _ := writer.CreateFormFile("recording", filepath.Base(file.Name()))
  io.Copy(part, file)
  _ = writer.WriteField("groupId", groupId)
  _ = writer.WriteField("contentLanguage", contentLanguage)
  writer.Close()

  req, _ := http.NewRequest("POST", BASE_URL + "/identification", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func VoiceIdentificationByUrl(groupId string, contentLanguage string, fileUrl string) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("fileUrl", fileUrl)
  _ = writer.WriteField("groupId", groupId)
  _ = writer.WriteField("contentLanguage", contentLanguage)
  writer.Close()

  req, _ := http.NewRequest("POST", BASE_URL + "/identification/byUrl", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func VideoIdentification(groupId string, contentLanguage string, filePath string, doBlinkDetection ... bool) *bytes.Buffer {

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

  req, _ := http.NewRequest("POST", BASE_URL + "/identification/video", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}

func VideoIdentificationByUrl(groupId string, contentLanguage string, fileUrl string, doBlinkDetection ... bool) *bytes.Buffer {

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  _ = writer.WriteField("fileUrl", fileUrl)
  _ = writer.WriteField("groupId", groupId)
  _ = writer.WriteField("contentLanguage", contentLanguage)
  if len(doBlinkDetection) > 0 {
    _ = writer.WriteField("doBlinkDetection", strconv.FormatBool(doBlinkDetection[0]))
  }
  writer.Close()

  req, _ := http.NewRequest("POST", BASE_URL + "/identification/video/byUrl", body)
  req.SetBasicAuth(api_key, api_token)
  req.Header.Add("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, _ := client.Do(req)
  respBody := &bytes.Buffer{}
  respBody.ReadFrom(resp.Body)

  return respBody
}
