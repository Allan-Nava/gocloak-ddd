package authuser


type LoginRequest struct {
	Username           string `json:"username"`
	Password           string `json:"password"`
}
//
type LogoutRequest struct {
	RefreshToken           string `json:"refresh_token"`
}
//