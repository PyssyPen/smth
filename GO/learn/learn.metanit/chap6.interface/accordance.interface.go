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

type File struct {
	text string
}

func (f *File) write(message string) {
	f.text = message
	fmt.Println("Запись в файл строки", message)
}
func writeToStream(writer Writer, text string) {
	writer.write(text)
}

func (f *File) read() {
	fmt.Println(f.text)
}
func readFromStream(reader Reader) {
	reader.read()
}


func main() {

	myFile := &File{}
	writeToStream(myFile, "hello world")
	readFromStream(myFile)
}

////////////////////////////////////////

type Stream interface {
	read() string
	write(string)
	close()
}

type File struct { // структура файл
	text string
}

type Folder struct{} // структура папка

func (f *File) read() string { // реализация методов для типа *File
	return f.text
}

func (f *File) write(message string) {
	f.text = message
	fmt.Println("Запись в файл строки", message)
}

func writeToStream(stream Stream, text string) {
	stream.write(text)
}

func (f *File) close() {
	fmt.Println("Файл закрыт")
}

func (f *Folder) close() { // релизация методов для типа *Folder
	fmt.Println("Папка закрыта")
}

func closeStream(stream Stream) { // хз зачем эта хуйня, можно обойтись без нее
	stream.close()
}

func main() {

	myFile := &File{}
	myFolder := &Folder{}

	writeToStream(myFile, "hello world")
	myFile.close() //closeStream(myFile)
	//closeStream(myFolder)     // Ошибка: тип *Folder не реализует интерфейс Stream
	myFolder.close() // Так можно
}
*/
