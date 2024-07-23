package common

import "time"

type DefaultFieldTable struct {
	CreatedAt  string `json:"created_at" db:"created_at"`
	CreatedBy  string `json:"created_by" db:"created_by"`
	ModifiedAt string `json:"modified_at" db:"modified_at"`
	ModifiedBy string `json:"modified_by" db:"modified_by"`
}

func (d *DefaultFieldTable) SetDefaultField(createdAt, createdBy, modifiedAt, modifiedBy string) {
	d.CreatedBy = "SYSTEM"
	d.ModifiedBy = "SYSTEM"
	d.CreatedAt = time.Now().String()
	d.ModifiedAt = time.Now().String()

	if !IsEmptyField(createdAt) {
		d.CreatedAt = createdAt
	}

	if !IsEmptyField(createdBy) {
		d.CreatedBy = createdBy
	}

	if !IsEmptyField(modifiedAt) {
		d.ModifiedAt = modifiedAt
	}

	if !IsEmptyField(modifiedBy) {
		d.ModifiedBy = modifiedBy
	}
}
