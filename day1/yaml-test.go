package day1

import (
	"fmt"
	yml "gopkg.in/yaml.v2"
	"os"
)

var confMap = make(map[string]interface{})

func ReadYaml() {
	dat, err := os.ReadFile("./day1/test.yaml")
	if err != nil {
		panic("文件不存在")
	}
	yml.Unmarshal(dat, confMap)
	//打印 server下面的port
	fmt.Println(confMap["server"].(map[interface{}]interface{})["port"])
}
