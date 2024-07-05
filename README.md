# go-learning
long-term project to learn the go programming language.

## Docs
Unit testing adalah salah satu praktik penting dalam pengembangan perangkat lunak untuk memastikan bahwa setiap unit kode (seperti fungsi atau metode) berfungsi sesuai dengan yang diharapkan. Di Go (Golang), unit testing dapat dilakukan menggunakan paket bawaan testing. Berikut adalah contoh sederhana penerapan unit testing di Golang.

### Instalasi stretchr/testify
Anda bisa menginstal stretchr/testify dengan perintah berikut:

```sh
go get https://github.com/stretchr/testify
```


Kita bisa menggunakan library seperti stretchr/testify untuk memperluas kapabilitas testing di Go. stretchr/testify menyediakan berbagai fitur yang memudahkan penulisan test case

Mari kita buat contoh sederhana dengan fungsi yang ingin diuji.

```
func Addition(a, b int64) int64 {
	return a + b
}
```
Berikut ini penjelasan lebih detail penggunaan library ini di project kali ini.

##### Package Import
```
import (
	"github.com/stretchr/testify/assert"
	"testing"
)
```

Kode ini mengimpor dua package:

* github.com/stretchr/testify/assert 
* testing

##### Struct Definition
```
type TestCases struct {
	Name     string
	Actual   interface{}
	Expected interface{}
	Data
}
```

Nama atau deskripsi test case.
Actual adalah nilai hasil yang sebenarnya dari eksekusi fungsi yang diuji.
Expected adalah nilai hasil yang diharapkan dari eksekusi fungsi yang diuji.
Data adalah eEbedded struct yang berisi data yang diperlukan untuk test case.

##### Test Function
```
func TestAddition(t *testing.T) {
	testCases := []TestCases{
		{
			Name: "Success - Normal Addition 1",
			Data: Data{
				Number1: 10,
				Number2: 2,
			},
			Expected: int64(12),
		},
		{
			Name: "Success - Normal Addition 2",
			Data: Data{
				Number1: 1,
				Number2: 2,
			},
			Expected: int64(3),
		},
		{
			Name: "Success - Normal Addition 3",
			Data: Data{
				Number1: -10,
				Number2: 2,
			},
			Expected: int64(-8),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Actual = Addition(tc.Number1, tc.Number2)

			assert.Equal(t, tc.Expected, tc.Actual)
		})
	}
}

```

##### Pendefinisian Test Cases
* testCases adalah slice dari TestCases yang berisi berbagai test case untuk fungsi Addition. Masing-masing test case memiliki nama, data (dua angka untuk ditambahkan), dan hasil yang diharapkan.

##### Looping Melalui Test Cases
* for _, tc := range testCases { ... } akan melakukan iterasi melalui setiap test case.
* t.Run(tc.Name, func(t *testing.T) { ... }) menjalankan setiap test case dalam subtest yang dinamai sesuai dengan tc.Name.

##### Eksekusi Fungsi
* tc.Actual = Addition(tc.Number1, tc.Number2) menjalankan fungsi Addition (yang diasumsikan sudah didefinisikan di tempat lain) dengan dua angka dari Data struct, dan menyimpan hasilnya dalam tc.Actual.
##### Assertions
* assert.Equal(t, tc.Expected, tc.Actual) memeriksa apakah nilai yang dihasilkan (tc.Actual) sesuai dengan nilai yang diharapkan (tc.Expected). Jika tidak, maka test case akan gagal dan menampilkan pesan kesalahan.

setelah semua siap digunakan, kita bisa uji coba unit test kita dengan menggunakan perintah
```
go test -v ./...
```

## License
Project ini belum memiliki lisensi resmi. Anda bisa melakukan fork untuk clone project ini

