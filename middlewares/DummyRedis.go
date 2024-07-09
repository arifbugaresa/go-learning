package middlewares

import "time"

var DummyRedis = make(map[string]UserLoginRedis)

type UserLoginRedis struct {
	UserId    int64
	Username  string
	Role      string
	LoginAt   time.Time
	ExpiredAt time.Time
}
