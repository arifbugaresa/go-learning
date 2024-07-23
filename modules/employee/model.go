package employee

import "go-learning/utils/common"

type Employee struct {
	ID         uint   `db:"id"`
	FullName   string `db:"full_name"`
	Email      string `db:"email"`
	Age        int    `db:"age"`
	Division   string `db:"division"`
	CreatedAt  string `db:"created_at"`
	ModifiedAt string `db:"modified_at"`
	CreatedBy  string `db:"created_by"`
	ModifiedBy string `db:"modified_by"`
}

func (e Employee) ConvertModelToResponseForGetListEmployee() GetEmployeeResponse {
	return GetEmployeeResponse{
		ID:         e.ID,
		FullName:   e.FullName,
		Email:      e.Email,
		Age:        e.Age,
		Division:   e.Division,
		ModifiedAt: e.ModifiedAt,
		ModifiedBy: e.ModifiedBy,
		CreatedBy:  e.CreatedBy,
		CreatedAt:  e.CreatedAt,
	}
}

type GetEmployeeResponse struct {
	ID         uint   `json:"id"`
	FullName   string `json:"full_name"`
	Email      string `json:"email"`
	Age        int    `json:"age"`
	Division   string `json:"division"`
	CreatedBy  string `json:"created_by"`
	CreatedAt  string `json:"created_at"`
	ModifiedBy string `json:"modified_by"`
	ModifiedAt string `json:"modified_at"`
}

type GetEmployeeRequest struct {
	SearchBy SearchGetEmployeeRequest `json:"search"`
	common.ListRequest
}

type SearchGetEmployeeRequest struct {
	FullName *string `json:"full_name"`
}
