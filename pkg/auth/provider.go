package auth

type Role string

const (
	Admin Role = "admin"
	User  Role = "user"
)

type TokenClaims struct {
	Role Role `json:"role"`
	Id   int  `json:"id"`
}

type Provider interface {
	NewAdminJWT(adminId int) (string, error)
	NewUserJWT(userId int) (string, error)
	Parse(accessToken string) (TokenClaims, error)
}
