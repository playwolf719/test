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

//func Struct2Map(obj interface{}) map[string]interface{} {
//	t := reflect.TypeOf(obj)
//	v := reflect.ValueOf(obj)
//
//	var data = make(map[string]interface{})
//	for i := 0; i < t.NumField(); i++ {
//		data[t.Field(i).Name] = v.Field(i).Interface()
//	}
//	return data
//}
