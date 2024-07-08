package car

import "errors"

type Repository interface {
	CreateCarRepository(car Car) (result []Car, err error)
	GetAllCarRepository() (result []Car, err error)
	GetCarRepository(car Car) (result Car, err error)
	DeleteCarRepository(car Car) (err error)
	UpdateCarRepository(car Car) (err error)
}

type carRepository struct{}

func NewRepository() Repository {
	return &carRepository{}
}

func (r *carRepository) CreateCarRepository(car Car) (result []Car, err error) {
	car.Id = len(DummyCars) + 1
	DummyCars = append(DummyCars, car)

	return DummyCars, nil
}

func (r *carRepository) GetAllCarRepository() (result []Car, err error) {
	return DummyCars, nil
}

func (r *carRepository) GetCarRepository(car Car) (result Car, err error) {
	var isCarExists bool
	for _, item := range DummyCars {
		if item.Id == car.Id {
			isCarExists = true
			return item, nil
		}
	}

	if !isCarExists {
		err = errors.New("data not found")
		return
	}

	return
}

func (r *carRepository) DeleteCarRepository(car Car) (err error) {
	var indexDummyItem int

	for index, item := range DummyCars {
		if item.Id == car.Id {
			indexDummyItem = index + 1
		}
	}

	if indexDummyItem == 0 {
		return errors.New("data not found")
	}

	// process delete
	copy(DummyCars[indexDummyItem-1:], DummyCars[indexDummyItem:])
	DummyCars[len(DummyCars)-1] = Car{}
	DummyCars = DummyCars[:len(DummyCars)-1]

	return
}

func (r *carRepository) UpdateCarRepository(car Car) (err error) {
	var indexDummyItem int

	for index, item := range DummyCars {
		if item.Id == car.Id {
			indexDummyItem = index + 1
		}
	}

	if indexDummyItem == 0 {
		return errors.New("data not found")
	}

	// process update
	DummyCars[indexDummyItem-1] = car

	return
}
