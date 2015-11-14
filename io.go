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
