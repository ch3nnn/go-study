/**
 * 测试阿里云文档翻译api
 *
 * @Author: chentong
 * @Date: 2022/04/05 19:05
 */

package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alimt"
	"time"
)

// 阿里云
var accessKeyId = ""
var accessKeySecret = ""

/*
 * getDocTranslateTask
 * @Description:
 * @param taskId
 * @return *alimt.GetDocTranslateTaskResponse
 */
func getDocTranslateTask(taskId string) *alimt.GetDocTranslateTaskResponse {

	client, err := alimt.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessKeySecret)
	request := alimt.CreateGetDocTranslateTaskRequest()
	request.Scheme = "https"
	request.TaskId = taskId

	response, err := client.GetDocTranslateTask(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)

	return response
}

//
//  createDocTranslateTask
//  @Description:  创建文档翻译任务
//  @param fileUrl 翻译文件url
//  @param sourceLanguage 原文语言
//  @param targetLanguage 目标语言
//  @return *alimt.CreateDocTranslateTaskResponse
//
func createDocTranslateTask(fileUrl, sourceLanguage, targetLanguage string) *alimt.CreateDocTranslateTaskResponse {
	client, err := alimt.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessKeySecret)

	request := alimt.CreateCreateDocTranslateTaskRequest()
	request.Scheme = "https"
	request.SourceLanguage = sourceLanguage
	request.TargetLanguage = targetLanguage
	request.FileUrl = fileUrl

	response, err := client.CreateDocTranslateTask(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)

	return response

}

func main() {
	fileUrl := ""

	var sourceLanguage = "ko"
	var targetLanguage = "zh"

	taskResponse := createDocTranslateTask(fileUrl, sourceLanguage, targetLanguage)
	fmt.Println("翻译id", taskResponse.TaskId)

	var translateTask *alimt.GetDocTranslateTaskResponse
	for {
		time.Sleep(time.Duration(10) * time.Second)
		translateTask = getDocTranslateTask(translateTask.TaskId)
		if translateTask.Status == "translated" {
			break
		}
	}
	fmt.Println("翻译完成")
	fmt.Println(translateTask.TranslateFileUrl)

}
