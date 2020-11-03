package structs

type Group struct {
	CreatedAt int      `json:"createdAt"`
	GroupId   string   `json:"groupId"`
	Users     []string `json:"users"`
	UserCount int      `json:"userCount"`
	APICallId string   `json:"apiCallId"`
}

type GetAllGroupsReturn struct {
	Message      string  `json:"message"`
	Count        int     `json:"count"`
	Status       int     `json:"status"`
	TimeTaken    string  `json:"timeTaken"`
	Groups       []Group `json:"groups"`
	ResponseCode string  `json:"responseCode"`
	APICallId    string  `json:"apiCallId"`
}

type GetGroupReturn struct {
	Message      string   `json:"message"`
	CreatedAt    int      `json:"createdAt"`
	Users        []string `json:"users"`
	UserCount    int      `json:"userCount"`
	Status       int      `json:"status"`
	TimeTaken    string   `json:"timeTaken"`
	ResponseCode string   `json:"responseCode"`
	APICallId    string   `json:"apiCallId"`
}

type CheckGroupExistsReturn struct {
	Message      string `json:"message"`
	Exists       bool   `json:"exists"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
	APICallId    string `json:"apiCallId"`
}

type CreateGroupReturn struct {
	Message      string `json:"message"`
	Description  string `json:"description"`
	GroupId      string `json:"groupId"`
	Status       int    `json:"status"`
	CreatedAt    int    `json:"createdAt"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
	APICallId    string `json:"apiCallId"`
}

type AddUserToGroupReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
	APICallId    string `json:"apiCallId"`
}

type RemoveUserFromGroupReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
	APICallId    string `json:"apiCallId"`
}

type DeleteGroupReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
	APICallId    string `json:"apiCallId"`
}
