package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var p = fmt.Println
var imageName = "registry.cn-beijing.aliyuncs.com/yunli_ai/video-media-platform"

//var imageName = "elasticsearch"
var fullImageNamePrefix = "registry.cn-beijing.aliyuncs.com/yunli_ai/video-media-platform:"

//var fullImageNamePrefix = "yunli/media-server:"

func getContainerId() string {
	contailerLsCmd := exec.Command("bash", "-c", "docker ps")
	contailerLsOut, err := contailerLsCmd.Output()
	if err != nil {
		panic(err)
	}
	out := string(contailerLsOut)
	//p(out)
	splitOut := strings.Split(out, "\n")
	for _, line := range splitOut {
		//p("line: ",line)
		splitReg, err := regexp.Compile("\\s+")
		if err != nil {
			panic(err)
		}
		if len(line) == 0 {
			break
		}
		lineVals := splitReg.Split(line, -1)
		if len(lineVals) > 0 && strings.HasPrefix(lineVals[1], fullImageNamePrefix) {
			return lineVals[0]
		}
	}
	return ""

}

func getImageId() string {
	imageLsCmd := exec.Command("bash", "-c", "docker image ls")
	imageLsOut, err := imageLsCmd.Output()
	if err != nil {
		panic(err)
	}
	out := string(imageLsOut)
	//p(out)
	splitOut := strings.Split(out, "\n")
	for _, line := range splitOut {
		//p("line: ",line)
		splitReg, err := regexp.Compile("\\s+")
		if err != nil {
			panic(err)
		}
		if len(line) == 0 {
			break
		}
		lineVals := splitReg.Split(line, -1)
		if len(lineVals) > 0 && lineVals[0] == imageName {
			return lineVals[2]
		}
	}
	return ""
}

func main2() {
	//停止容器
	containerId := getContainerId()
	stopContainerCmd := exec.Command("bash", "-c", strings.Join([]string{"docker", "stop", containerId}, " "))
	p("stopContainer ... , ", (stopContainerCmd))
	stopContainerCmd.Start()
	stopContainerCmd.Wait()
	//stopCaOutput, err := stopContainerCmd.Output()
	//if err != nil {
	//	panic(err)
	//}
	//p(stopCaOutput)

	//rm容器
	rmContainerCmd := exec.Command("bash", "-c", strings.Join([]string{"docker rm", containerId}, " "))
	p("rmContainerCmd ... , ", (rmContainerCmd))
	rmContainerCmd.Start()
	rmContainerCmd.Wait()
	//rmCaOutput, err := rmContainerCmd.Output()
	//if err != nil {
	//	panic(err)
	//}
	//p(rmCaOutput)

	//删除镜像
	imageId := getImageId()
	p("containerId: ", containerId)
	p("imageId: ", imageId)
	rmImageCmd := exec.Command("bash", "-c", strings.Join([]string{"docker image rm", imageId}, " "))
	p("rmImageCmd ... , ", (rmImageCmd))
	rmImageCmd.Start()
	rmImageCmd.Wait()
	//rmImageOutput, err := rmImageCmd.Output()
	//if err != nil {
	//	panic(err)
	//}
	//p(rmImageOutput)

}
