package permission

type Permission struct {
	AccessCode string `json:"access_code" db:"access_code"`
	GrantCode  string `json:"grant_code" db:"grant_code"`
}
