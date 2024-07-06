package modules

import "fmt"

func LearningInterface() {
	InterfaceAsBluePrint()
	InterfaceAsDataType()
}

func InterfaceAsDataType() {
	var (
		varInterface interface{}
	)

	varInterface = 1
	varInterface = "1"
	varInterface = true

	fmt.Println(varInterface)
}

func InterfaceAsBluePrint() {
	var (
		persegiPanjang BangunDatar
		persegi        BangunDatar
	)

	persegiPanjang = PersegiPanjang{
		Panjang: 10,
		Lebar:   5,
	}

	persegi = Persegi{
		Sisi: 10,
	}

	luasPersegiPanjang := persegiPanjang.Luas()
	luasPersegi := persegi.Luas()

	fmt.Println(luasPersegi)
	fmt.Println(luasPersegiPanjang)
}

type BangunDatar interface {
	Luas() int64
}

type PersegiPanjang struct {
	Panjang int64
	Lebar   int64
}

func (p PersegiPanjang) Luas() int64 {
	return p.Panjang * p.Lebar
}

type Persegi struct {
	Sisi int64
}

func (p Persegi) Luas() int64 {
	return p.Sisi * p.Sisi
}
