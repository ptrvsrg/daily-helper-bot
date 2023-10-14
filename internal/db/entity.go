package db

type UserEntity struct {
	ChatID       int64
	AccessToken  string
	RefreshToken string
	Scenario     string
}
