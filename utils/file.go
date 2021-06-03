package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func CheckPath(path string) {
	_, err := os.Stat(path)
	if err != nil {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func GetAllFile(dir string, s []string) []string {
	rd, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := dir + "/" + fi.Name()
			s = GetAllFile(fullDir, s)
			if err != nil {
				fmt.Println("read dir fail:", err)
				return s
			}
		} else {
			fullName := path.Join(dir, fi.Name())
			s = append(s, fullName)
		}
	}
	return s
}
