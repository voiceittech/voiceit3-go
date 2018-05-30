# VoiceIt2-Go
[![Go project version](https://badge.fury.io/go/github.com%2Fvoiceittech%2FVoiceIt2-Go.svg)](https://badge.fury.io/go/github.com%2Fvoiceittech%2FVoiceIt2-Go)

A Go wrapper for VoiceIt's API2.0 featuring Face + Voice Verification and Identification.

* [Getting Started](#getting-started)
* [Installation](#installation)
* [API Calls](#api-calls)
  * [Initialization](#initialization)
  * [User API Calls](#user-api-calls)
      * [Get All Users](#get-all-users)
      * [Create User](#create-user)
      * [Check User Exists](#check-user-exists)
      * [Get Groups for User](#get-groups-for-user)
      * [Delete User](#delete-user)
  * [Group API Calls](#group-api-calls)
      * [Get All Groups](#get-all-groups)
      * [Create Group](#create-group)
      * [Get Group](#get-group)
      * [Delete Group](#delete-group)
      * [Check Group Exists](#check-group-exists)
      * [Add User to Group](#add-user-to-group)
      * [Remove User from Group](#remove-user-from-group)      
  * [Enrollment API Calls](#enrollment-api-calls)
      * [Get All Enrollments for User](#get-all-enrollments-for-user)
      * [Get Face Enrollments for User](#get-face-enrollments-for-user)
      * [Delete All Enrollments for User](#delete-all-enrollments-for-user)
      * [Delete Enrollment for User](#delete-enrollment-for-user)
      * [Delete Face Enrollment](#delete-face-enrollment)
      * [Create Voice Enrollment](#create-voice-enrollment)
      * [Create Voice Enrollment By URL](#create-voice-enrollment-by-url)
      * [Create Video Enrollment](#create-video-enrollment)
      * [Create Video Enrollment By URL](#create-video-enrollment-by-url)
      * [Create Face Enrollment](#create-face-enrollment)
  * [Verification API Calls](#verification-api-calls)
      * [Voice Verification](#voice-verification)
      * [Voice Verification By URL](#voice-verification-by-url)
      * [Video Verification](#video-verification)
      * [Video Verification By URL](#video-verification-by-url)
      * [Face Verification](#face-verification)
  * [Identification API Calls](#identification-api-calls)
      * [Voice Identification](#voice-identification)
      * [Voice Identification By URL](#voice-identification-by-url)
      * [Video Identification](#video-identification)
      * [Video Identification By URL](#video-identification-by-url)

## Getting Started

Sign up for a free Developer Account at [voiceit.io](https://voiceit.io/signup) and activate API 2.0 from the settings page. Then you should be able view the API Key and Token. You can also review the HTTP Documentation at [api.voiceit.io](https://api.voiceit.io)

## Installation

In order to easily integrate VoiceIt API 2 into your Go project, please install the VoiceIt Go Package by running the following command in your Go Workspace.

```
go get github.com/voiceittech/VoiceIt2-Go/voiceit2
```

## API Calls

### Initialization

Make Sure to add this at the top of your project

```
import "github.com/voiceittech/VoiceIt2-Go/voiceit2"
```

First assign the API Credentials an initialize a VoiceIt2 struct.

```go
myVoiceIt := voiceit2.NewClient("API_KEY", "API_TOKEN")
```

### User API Calls

#### Get All Users

Get all  users associated with the apiKey
```go
myVoiceIt.GetAllUsers()
```

#### Create User

Create a new user
```go
myVoiceIt.CreateUser()
```

#### Check User Exists

Check whether a user exists for the given userId(begins with 'usr_')
```go
myVoiceIt.CheckUserExists("USER_ID_HERE").
```

#### Get Groups for User

Get a list of groups that the user with given userId(begins with 'usr_') is a part of
```go
myVoiceIt.GetGroupsForUser("USER_ID_HERE")
```

#### Delete User

Delete user with given userId(begins with 'usr_')
```go
myVoiceIt.DeleteUser("USER_ID_HERE")
```

### Group API Calls

#### Get All Groups

Get all the groups associated with the apiKey
```go
myVoiceIt.GetAllGroups()
```

#### Create Group

Create a new group with the given description
```go
myVoiceIt.CreateGroup("Sample Group Description")
```

#### Get Group

Returns a group for the given groupId(begins with 'grp_')
```go
myVoiceIt.GetGroup("GROUP_ID_HERE")
```

#### Delete Group

Delete group with given groupId(begins with 'grp_'), Note: This call does not delete any users, but simply deletes the group and disassociates the users from the group

```go
myVoiceIt.DeleteGroup("GROUP_ID_HERE")
```

#### Check Group Exists

Checks if group with given groupId(begins with 'grp_') exists
```go
myVoiceIt.CheckGroupExists("GROUP_ID_HERE")
```

#### Add User to Group

Adds user with given userId(begins with 'usr_') to group with given groupId(begins with 'grp_')
```go
myVoiceIt.AddUserToGroup("GROUP_ID_HERE", "USER_ID_HERE")
```

#### Remove User from Group

Removes user with given userId(begins with 'usr_') from group with given groupId(begins with 'grp_')

```go
myVoiceIt.RemoveUserFromGroup("GROUP_ID_HERE", "USER_ID_HERE")
```


### Enrollment API Calls

#### Get All Enrollments for User

Gets all enrollment for user with given userId(begins with 'usr_')

```go
myVoiceIt.GetAllEnrollmentsForuser("USER_ID_HERE")
```

#### Get Face Enrollments for User

Gets face enrollments for user with given userId(begins with 'usr_')

```go
myVoiceIt.GetFaceEnrollmentsForuser("USER_ID_HERE")
```

#### Delete All Enrollments for User

Delete all enrollments for user with the given userId(begins with 'usr_')

```go
myVoiceIt.DeleteAllEnrollmentForUser( "USER_ID_HERE")
```

#### Delete Enrollment for User

Delete enrollment for user with given userId(begins with 'usr_') and enrollmentId(integer)

```go
myVoiceIt.DeleteEnrollmentForUser( "USER_ID_HERE", "ENROLLMENT_ID_HERE")
```

#### Delete Face Enrollment

Delete face enrollment for user with given userId(begins with 'usr_') and faceEnrollmentId(integer)

```go
myVoiceIt.DeleteFaceEnrollment( "USER_ID_HERE", "FACE_ENROLLMENT_ID_HERE")
```

#### Create Voice Enrollment

Create voice enrollment for user with given userId(begins with 'usr_') and contentLanguage('en-US','es-ES', etc.). Note: File recording need to be no less than 1.2 seconds and no more than 5 seconds

```go
myVoiceIt.CreateVoiceEnrollment("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", filePath);
```

#### Create Voice Enrollment by URL

Create voice enrollment for user with given userId(begins with 'usr_') and contentLanguage('en-US','es-ES', etc.). Note: File recording need to be no less than 1.2 seconds and no more than 5 seconds

```go
myVoiceIt.CreateVoiceEnrollmentByUrl("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", "PUBLIC_URL_TO_AUDIO_FILE_HERE");
```

#### Create Video Enrollment

Create video enrollment for user with given userId(begins with 'usr_'), contentLanguage('en-US','es-ES', etc.) and optionally a boolean to disable blink detection. Note: File recording need to be no less than 1.2 seconds and no more than 5 seconds

```go
myVoiceIt.CreateVideoEnrollment("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", filePath);
```

or with blinkDetection enabled

```go
myVoiceIt.CreateVideoEnrollment("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", filePath, true);
```

#### Create Video Enrollment by URL

Create video enrollment for user with given userId(begins with 'usr_'), contentLanguage('en-US','es-ES', etc.) and optionally a boolean to disable blink detection. Note: File recording need to be no less than 1.2 seconds and no more than 5 seconds

```go
myVoiceIt.CreateVideoEnrollmentByUrl("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", "PUBLIC_URL_TO_VIDEO_FILE_HERE");
```

or with blinkDetection enabled

```go
myVoiceIt.CreateVideoEnrollmentByUrl("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", "PUBLIC_URL_TO_VIDEO_FILE_HERE", true);
```

#### Create Face Enrollment

Create face enrollment for user with given userId(begins with 'usr_') and optionally a boolean to disable blink detection. Note: It is recommended that you send a 2 second mp4 video

```go
myVoiceIt.CreateFaceEnrollment("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", filePath);
```

or with blinkDetection enabled

```go
myVoiceIt.CreateFaceEnrollment("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", filePath, true);
```

### Verification API Calls

#### Voice Verification

Verify user with the given userId(begins with 'usr_') and contentLanguage('en-US','es-ES', etc.). Note: File recording need to be no less than 1.2 seconds and no more than 5 seconds

```go
myVoiceIt.VoiceVerification("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", filePath)
```

#### Voice Verification by URL

Verify user with the given userId(begins with 'usr_') and contentLanguage('en-US','es-ES', etc.). Note: File recording need to be no less than 1.2 seconds and no more than 5 seconds

```go
myVoiceIt.VoiceVerificationByUrl("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", "PUBLIC_URL_TO_AUDIO_FILE_HERE")
```

#### Video Verification

Verify user with given userId(begins with 'usr_'), contentLanguage('en-US','es-ES', etc.) and optionally a boolean to disable blink detection. Note: File recording needs to be no less than 1.2 seconds and no more than 5 seconds

```go
myVoiceIt.VideoVerification("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", filePath)
```

or with blinkDetection enabled

```go
myVoiceIt.VideoVerification("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", filePath, true)
```

#### Video Verification by URL

Verify user with given userId(begins with 'usr_'), contentLanguage('en-US','es-ES', etc.) and optionally a boolean to disable blink detection. Note: File recording needs to be no less than 1.2 seconds and no more than 5 seconds

```go
myVoiceIt.VideoVerificationByUrl("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", "PUBLIC_URL_TO_VIDEO_FILE_HERE")
```

or with blinkDetection enabled

```go
myVoiceIt.VideoVerificationByUrl("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", "PUBLIC_URL_TO_VIDEO_FILE_HERE", true)
```

#### Face Verification

Verify user's face with given userId(begins with 'usr_') and optionally a boolean to disable blink detection. Note: Provide an about 2 seconds long video(mp4 codec is recommended) of the user's face

```go
myVoiceIt.FaceVerification("USER_ID_HERE", filePath)
```

or with blinkDetection enabled

```go
myVoiceIt.FaceVerification("USER_ID_HERE", filePath, true)
```

### Identification API Calls

#### Voice Identification

Identify user inside group with the given groupId(begins with 'grp_') and contentLanguage('en-US','es-ES', etc.). Note: File recording needs to be no less than 1.2 seconds and no more than 5 seconds

```go
myVoiceIt.VoiceIdentification("GROUP_ID_HERE", "CONTENT_LANGUAGE_HERE", filePath)
```

#### Voice Identification by URL

Identify user inside group with the given groupId(begins with 'grp_') and contentLanguage('en-US','es-ES', etc.). Note: File recording needs to be no less than 1.2 seconds and no more than 5 seconds

```go
myVoiceIt.VoiceIdentificationByUrl("GROUP_ID_HERE", "CONTENT_LANGUAGE_HERE", "PUBLIC_URL_TO_AUDIO_FILE_HERE")
```

#### Video Identification

Identify user inside group with the given groupId(begins with 'grp_'), contentLanguage('en-US','es-ES', etc.) and optionally a boolean to disable blink detection. Note: File recording needs to be no less than 1.2 seconds and no more than 5 seconds

```go
myVoiceIt.VideoIdentification("GROUP_ID_HERE", "CONTENT_LANGUAGE_HERE", filePath)
```
or with blinkDetection enabled
```go
myVoiceIt.VideoIdentification("GROUP_ID_HERE", "CONTENT_LANGUAGE_HERE", filePath, true)
```

#### Video Identification by URL

Identify user inside group with the given groupId(begins with 'grp_') , contentLanguage('en-US','es-ES', etc.) and optionally a boolean to disable blink detection. Note: File recording needs to be no less than 1.2 seconds and no more than 5 seconds

```go
myVoiceIt.VideoIdentificationByUrl("GROUP_ID_HERE", "CONTENT_LANGUAGE_HERE", "PUBLIC_URL_TO_VIDEO_FILE_HERE")
```

or with blinkDetection enabled

```go
myVoiceIt.VideoIdentificationByUrl("GROUP_ID_HERE", "CONTENT_LANGUAGE_HERE", "PUBLIC_URL_TO_VIDEO_FILE_HERE", true)
```

## Authors

Stephen Akers, stephen@voiceit.io
Armaan Bindra, armaan@voiceit.io

## License

VoiceIt2-Go is available under the MIT license. See the LICENSE file for more info.
