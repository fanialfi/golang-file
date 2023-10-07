package lib

import (
	"fmt"
	"os"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func CreateFile() {
	path := "/tmp/tes.json"

	// deteksi, apakah file sudah ada
	_, err := os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		file, err := os.Create(path)

		if isError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("=====>\tfile berhasil dibuat", path)
}
