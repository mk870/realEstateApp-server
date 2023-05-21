package tokens

import "os"

func GetAccessTokenSecret() string {
	return os.Getenv("ACCESS_TOKEN_SECRET")
}
func GetRefreshTokenSecret() string {
	return os.Getenv("REFRESH_TOKEN_SECRET")
}
