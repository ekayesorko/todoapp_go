package serializer

type LoginResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Username     string `json:"username"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoggedInUser struct {
	UserID uint `json:"userID"`
}
