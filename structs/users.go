package structs

type User struct {
	CreatedAt int    `json:"createdAt"`
	UserId    string `json:"userId"`
}

type GetAllUsersReturn struct {
	Message      string `json:"message"`
	Count        int    `json:"count"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	Users        []User `json:"users"`
	ResponseCode string `json:"responseCode"`
	APICallId    string `json:"apiCallId"`
}

type CreateUserReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	UserId       string `json:"userId"`
	ResponseCode string `json:"responseCode"`
	APICallId    string `json:"apiCallId"`
}

type CheckUserExistsReturn struct {
	Message      string `json:"message"`
	Exists       bool   `json:"exists"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
	APICallId    string `json:"apiCallId"`
}

type DeleteUserReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
	APICallId    string `json:"apiCallId"`
}

type GetGroupsForUserReturn struct {
	Message      string   `json:"message"`
	Groups       []string `json:"groups"`
	Count        int      `json:"count"`
	Status       int      `json:"status"`
	TimeTaken    string   `json:"timeTaken"`
	ResponseCode string   `json:"responseCode"`
	APICallId    string   `json:"apiCallId"`
}

type CreateUserTokenReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	UserToken    string `json:"userToken"`
	CreatedAt    int    `json:"createdAt"`
	ResponseCode string `json:"responseCode"`
	APICallId    string `json:"apiCallId"`
}

type ExpireUserTokensReturn struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	TimeTaken    string `json:"timeTaken"`
	ResponseCode string `json:"responseCode"`
	APICallId    string `json:"apiCallId"`
}
