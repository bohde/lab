package file

import "io/ioutil"

type FileReader struct{}

func (f *FileReader) Read(file string) (string, error) {
	bytes, err := ioutil.ReadFile(file)
	return string(bytes), err
}
