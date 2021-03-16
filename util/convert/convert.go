package convert

import (
	"io/ioutil"

	yaml2 "k8s.io/apimachinery/pkg/util/yaml"
)

var (
	data []byte
	err  error
)

func Init(fileName string, dirPath string) []byte {
	filePath := dirPath + fileName
	// 读取yaml文件
	data, err = ioutil.ReadFile(filePath)
	if err != nil {
		panic(err.Error())
	}

	data, err = yaml2.ToJSON(data)
	if err != nil {
		panic(err.Error())
	}

	return data
}
