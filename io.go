package godash

import (
	"io"
	"io/ioutil"
	"os"
)

func FileFromFilePath(filePath string) *os.File {
	file, _ := os.Open(filePath)
	return file
}

func ReaderFromFilePath(filePath string) io.Reader {
	file := FileFromFilePath(filePath)

	var ir io.Reader
	ir = file

	return ir
}

func bufferFromFilePath(filePath string) []byte {
	reader := ReaderFromFilePath(filePath)

	var file = *(reader.(*os.File))
	defer file.Close()

	buffer, _ := ioutil.ReadAll(reader)
	return buffer
}

func StrFromFilePath(filePath string) string {
	buffer := bufferFromFilePath(filePath)

	return string(buffer)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func BytesToString(bs []byte) string {
	return string(bs[:])
}
func StringToBytes(s string) []byte {
	return []byte(s)
}

func BytesToFilePath(filePath string, bs []byte) {
	err := ioutil.WriteFile(filePath, bs, 0644)
	check(err)
}

func StringToFilePath(filePath string, s string) {
	bs := StringToBytes(s)
	BytesToFilePath(filePath, bs)
}
