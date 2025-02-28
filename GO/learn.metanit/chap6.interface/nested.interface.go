package main

/*
import (
	"fmt"
)

type Reader interface {
	read()
}

type Writer interface {
	write(string)
}

type ReaderWriter interface {
	Reader
	Writer
}

type File struct {
	text string
}

func (f *File) write(message string) {
	f.text = message
	fmt.Println("Запись в файл строки", message)
}
func writeToStream(writer ReaderWriter, text string) {
	writer.write(text)
}

func (f *File) read() {
	fmt.Println(f.text)
}
func readFromStream(reader ReaderWriter) {
	reader.read()
}

func main() {

	myFile := &File{}
	writeToStream(myFile, "hello world")
	readFromStream(myFile)
	writeToStream(myFile, "lolly bomb")
	readFromStream(myFile)
	fmt.Println(myFile)
}
*/
