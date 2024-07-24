package middlewares

import "time"

var DummyRedis = make(map[string]string)

type RedisSession struct {
	UserId     int64             `json:"user_id"`
	Username   string            `json:"username"`
	RoleId     int64             `json:"role_id"`
	Permission []RedisPermission `json:"permission"`
	LoginAt    time.Time         `json:"login_at"`
	ExpiredAt  time.Time         `json:"expired_at"`
}

type RedisPermission struct {
	AccessCode  string `json:"access_code"`
	AccessGrant string `json:"access_grant"`
}
