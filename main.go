package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"
)

func main() {
	path := CreateFol()
	path = path + "/id.log"
	fmt.Println(path)
	time := time.Now()
	_, ee := os.Create(path)
	if ee != nil {
		log.Fatal(ee)
	}
	var file, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, fs.ModeAppend)
	e(err)
	defer file.Close()
	date := fmt.Sprintf("[%s]", time.Format("2006-01-02 15:04:05"))
	data := date + " MESSAGE: " + "tushar"
	_, err = file.WriteString(data + "\n")
	e(err)
}
func CreateFol() string {
	homeDir, err := os.UserHomeDir()
	e(err)
	path := homeDir + "/.config" + "/Tsunagu"
	err = os.Mkdir(path, 0755)
	if err != nil {
		fmt.Print()
	}
	return path
}
func e(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
