package email

type EmailTemplate struct {
	Id         int64  `db:"id"`
	Code       string `db:"code"`
	Name       string `db:"name"`
	Template   string `db:"template"`
	CreatedAt  string `db:"created_at"`
	ModifiedAt string `db:"modified_at"`
	CreatedBy  string `db:"created_by"`
	ModifiedBy string `db:"modified_by"`
}
