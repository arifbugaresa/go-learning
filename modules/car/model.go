package car

type Car struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Price int    `json:"price"`
}

var (
	DummyCars []Car
)
