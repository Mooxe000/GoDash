package godash

import (
	"io/ioutil"
	"os"
)

func FileStrFromPath(filePath string) string {
	file, _ := os.Open(filePath)
	defer file.Close()
	reader, _ := ioutil.ReadAll(file)
	return string(reader)
}
