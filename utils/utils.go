package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

var final_res = make(map[string]string)

func visit(path string, f os.FileInfo, err error) error {
	final_res[path] = path
	return nil
}
func GetFileMap(path string) map[string]string {
	err := filepath.Walk(path, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
	return final_res
}

