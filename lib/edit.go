package lib

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

var path = "/tmp/tes.json"

func WriteFile() {
	// buka file dengan level akses read & write
	file, err := os.OpenFile(path, os.O_RDWR, 0744)
	// file, err := os.Open(path)
	if isError(err) {
		return
	}
	defer file.Close()

	// tulis data ke file
	_, err = file.WriteString("Hello My Name Is Fani\n")
	if isError(err) {
		return
	}

	_, err = file.WriteString("Mari Belajar golang\n")
	if isError(err) {
		return
	}

	_, err = file.WriteString("cek cek\n")
	if isError(err) {
		return
	}

	// simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("=====>\tfile berhasil di isi")
}

func ReadFile() {
	// buka file
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {
		return
	}
	printAlloc()
	defer file.Close()

	// baca file
	text := make([]byte, 1024)
	printAlloc()
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		printAlloc()
		if n == 0 {
			break
		}
	}

	if isError(err) {
		return
	}
	fmt.Println("=====>\tfile berhasil dibaca")
	fmt.Println(string(text))
	printAlloc()
}

func DeleteFile() {
	err := os.Remove(path)

	if isError(err) {
		return
	}

	fmt.Println("=====>\tfile berhasil di delete")
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
