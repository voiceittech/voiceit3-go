# VoiceIt2-Go

A Go wrapper for VoiceIt's API2.0 featuring Face + Voice Verification and Identification.

* [Getting Started](#getting-started)
* [Installation](#installation)
* [API Calls](#api-calls)
  * [Initialization](#initialization)
  * [User API Calls](#user-api-calls)
      * [Get All Users](#get-all-users)
      * [Create User](#create-user)
      * [Get User](#check-if-user-exists)
      * [Get Groups for User](#get-groups-for-user)
      * [Delete User](#delete-user)
  * [Group API Calls](#group-api-calls)
      * [Get All Groups](#get-all-groups)
      * [Create Group](#create-group)
      * [Get Group](#get-group)
      * [Delete Group](#delete-group)
      * [Group exists](#check-if-group-exists)
      * [Add User to Group](#add-user-to-group)
      * [Remove User from Group](#remove-user-from-group)      
  * [Enrollment API Calls](#enrollment-api-calls)
      * [Get All Enrollments for User](#get-all-enrollments-for-user)
      * [Delete Enrollment for User](#delete-enrollment-for-user)
      * [Delete Face Enrollment](#delete-face-enrollment)
      * [Create Voice Enrollment](#create-voice-enrollment)
      * [Create Voice Enrollment By Url](#create-voice-enrollment-by-url)
      * [Create Video Enrollment](#create-video-enrollment)
      * [Create Video Enrollment By Url](#create-video-enrollment-by-url)
      * [Create Face Enrollment](#create-face-enrollment)
  * [Verification API Calls](#verification-api-calls)
      * [Voice Verification](#voice-verification)
      * [Voice Verification By Url](#voice-verification-by-url)
      * [Video Verification](#video-verification)
      * [Video Verification](#video-verification-by-url)
      * [Face Verification](#face-verification)
  * [Identification API Calls](#identification-api-calls)
      * [Voice Identification](#voice-identification)
      * [Voice Identification](#voice-identification-by-url)
      * [Video Identification](#video-identification)
      * [Video Identification](#video-identification-by-url)

## Getting Started

Sign up for a free Developer Account at [voiceit.io](https://voiceit.io/signup) and activate API 2.0 from the settings page. Then you should be able view the API Key and Token. You can also review the HTTP Documentation at [api.voiceit.io](https://api.voiceit.io)

## Installation

In order to easily integrate VoiceIt API 2 into your Go project, please install the VoiceIt Go Package by running the following command in your Go Workspace.

```
go get github.com/voiceittech/VoiceIt2-Go
```

## API Calls

### Initialization

// Make Sure to add this at the top of your project

```
import "github.com/voiceittech/VoiceIt2-Go/voiceit2"
```

First assign the API Credentials an initialize a VoiceIt2 struct.

```go
myVoiceIt := voiceit2.New("API_KEY", "API_TOK")
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

#### Check if User Exists

Check whether a user exists for the given userId(begins with 'usr_')
```go
myVoiceIt.GetUser("USER_ID_HERE").
```

#### Delete User

Delete user with given userId(begins with 'usr_')
```go
myVoiceIt.DeleteUser("USER_ID_HERE")
```

#### Get Groups for User

Get a list of groups that the user with given userId(begins with 'usr_') is a part of
```go
myVoiceIt.GetGroupsForUser("USER_ID_HERE")
```

### Group API Calls

#### Get All Groups

Get all the groups associated with the apiKey
```go
myVoiceIt.GetAllGroups()
```

#### Get Group

Returns a group for the given groupId(begins with 'grp_')
```go
myVoiceIt.GetGroup("GROUP_ID_HERE")
```

#### Check if Group Exists

Checks if group with given groupId(begins with 'grp_') exists
```go
myVoiceIt.GroupExists("GROUP_ID_HERE")
```

#### Create Group

Create a new group with the given description
```go
myVoiceIt.CreateGroup("Sample Group Description")
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

#### Delete Group

Delete group with given groupId(begins with 'grp_'), Note: This call does not delete any users, but simply deletes the group and disassociates the users from the group

```go
myVoiceIt.DeleteGroup("GROUP_ID_HERE")
```

### Enrollment API Calls

#### Get All Enrollments for User

Gets all enrollment for user with given userId(begins with 'usr_')

```go
myVoiceIt.GetAllEnrollmentsForuser("USER_ID_HERE")
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

Create audio enrollment for user with given userId(begins with 'usr_') and contentLanguage('en-US','es-ES', etc.). Note: File recording need to be no less than 1.2 seconds and no more than 5 seconds

```go
myVoiceIt.CreateVoiceEnrollment("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", filePath);
```

#### Create Voice Enrollment by URL

Create audio enrollment for user with given userId(begins with 'usr_') and contentLanguage('en-US','es-ES', etc.). Note: File recording need to be no less than 1.2 seconds and no more than 5 seconds

```go
myVoiceIt.CreateVoiceEnrollmentByUrl("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", "PUBLIC_URL_TO_AUDIO_FILE_HERE");
```

#### Create Video Enrollment

Create video enrollment for user with given userId(begins with 'usr_'), contentLanguage('en-US','es-ES', etc.) and optionally a boolean to disable blink detection. Note: File recording need to be no less than 1.2 seconds and no more than 5 seconds

```go
myVoiceIt.CreateVideoEnrollment("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", filePath);
```

or with blinkDetection disabled

```go
myVoiceIt.CreateVideoEnrollment("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", filePath, false);
```

#### Create Video Enrollment by URL

Create video enrollment for user with given userId(begins with 'usr_'), contentLanguage('en-US','es-ES', etc.) and optionally a boolean to disable blink detection. Note: File recording need to be no less than 1.2 seconds and no more than 5 seconds

```go
myVoiceIt.CreateVideoEnrollmentByUrl("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", "PUBLIC_URL_TO_VIDEO_FILE_HERE");
```

or with blinkDetection disabled

```go
myVoiceIt.CreateVideoEnrollmentByUrl("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", "PUBLIC_URL_TO_VIDEO_FILE_HERE", false);
```

#### Create Face Enrollment

Create face enrollment for user with given userId(begins with 'usr_') and optionally a boolean to disable blink detection. Note: It is recommended that you send a 2 second mp4 video

```go
myVoiceIt.CreateFaceEnrollment("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", filePath);
```

or with blinkDetection disabled

```go
myVoiceIt.CreateFaceEnrollment("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", filePath, false);
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

or with blinkDetection disabled

```go
myVoiceIt.VideoVerification("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", filePath, false)
```

#### Video Verification by URL

Verify user with given userId(begins with 'usr_'), contentLanguage('en-US','es-ES', etc.) and optionally a boolean to disable blink detection. Note: File recording needs to be no less than 1.2 seconds and no more than 5 seconds

```go
myVoiceIt.VideoVerificationByUrl("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", "PUBLIC_URL_TO_VIDEO_FILE_HERE")
```

or with blinkDetection disabled

```go
myVoiceIt.VideoVerificationByUrl("USER_ID_HERE", "CONTENT_LANGUAGE_HERE", "PUBLIC_URL_TO_VIDEO_FILE_HERE", false)
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
or with blinkDetection disabled

```go
myVoiceIt.VideoIdentification("GROUP_ID_HERE", "CONTENT_LANGUAGE_HERE", filePath, false)
```

#### Video Identification by URL

Identify user inside group with the given groupId(begins with 'grp_') , contentLanguage('en-US','es-ES', etc.) and optionally a boolean to disable blink detection. Note: File recording needs to be no less than 1.2 seconds and no more than 5 seconds

```go
myVoiceIt.VideoIdentificationByUrl("GROUP_ID_HERE", "CONTENT_LANGUAGE_HERE", "PUBLIC_URL_TO_VIDEO_FILE_HERE")
```

or with blinkDetection disabled

```go
myVoiceIt.VideoIdentificationByUrl("GROUP_ID_HERE", "CONTENT_LANGUAGE_HERE", "PUBLIC_URL_TO_VIDEO_FILE_HERE", false)
```

## Authors

Stephen Akers, stephen@voiceit.io
Armaan Bindra, armaan@voiceit.io

## License

VoiceIt2-Go is available under the MIT license. See the LICENSE file for more info.
