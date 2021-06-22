package main

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getOneStr(str string, lenStrInt int) string{
	strArray := strings.Split(str, "time=")
	lenStr := len(strArray) - 1
	return strArray[lenStr][lenStrInt:]
}

func Test_Server_Add(t *testing.T){
	assert := assert.New(t)
	expectedId := "ObjectID(\\\"68c3503f9a9e8b7cdfa813d4\\\")\"\n"
	cmd := exec.
		Command("/bin/sh", "-c", "curl -i -X POST -H 'Content-Type: application/json' -d '{\"BookId\":\"68c3503f9a9e8b7cdfa813d4\",\"AuthorId\":1,\"BookVolume\":2,\"NameOfBook\":\"qwe\",\"Number\":1,\"PublisherId\":2,\"YearOfPublication\":\"2020-12-2\"}' http://localhost:8081/books\n")
	cmd.Run()
	actualByte, _ := exec.Command("/bin/sh", "-c", "docker logs db-book").CombinedOutput()
	actualStr := string(actualByte)
	finalStr := getOneStr(actualStr, 39)
	assert.Equal(expectedId, finalStr)
}

func Test_Server_ReadAll(t *testing.T) {
	assert := assert.New(t)
	command := "curl http://localhost:8081/books"
	expected := "[ObjectID(\\\"60c3501f9a9e8b7cdfa813d7\\\") 1 2 qwe 2020-12-2 2 1\\n ObjectID(\\\"60c3501f9a9e8b7cdfa813d5\\\") 1 2 qwe 2020-12-2 2 1\\n ObjectID(\\\"60c3501f9a9e8b7cdfa813d4\\\") 1 2 qwe 2020-12-2 2 1\\n ObjectID(\\\"69c3501f9a9e8b7cdfa813d3\\\") 1 2 qwe 2020-12-2 2 1\\n ObjectID(\\\"60c3502f9a9e8b7cdfa813d4\\\") 1 2 qwe 2020-12-2 2 1\\n ObjectID(\\\"68c3502f9a9e8b7cdfa813d4\\\") 1 2 qwe 2020-12-2 2 1\\n]\"\n"
	cmd := exec.
		Command("/bin/sh", "-c", command)
	cmd.Run()
	actualByte, _ := exec.Command("/bin/sh", "-c", "docker logs db-book").CombinedOutput()
	actualStr := string(actualByte)
	finalStr := getOneStr(actualStr, 39)
	assert.Equal(expected, finalStr)
}

func Test_Server_ReadAllNotEqual(t *testing.T) {
	assert := assert.New(t)
	command := "curl http://localhost:8081/books"
	expected := "ObjectID(\\\"60c3501f9a9e8b7cdfa813d7\\\") 1 2 qwe 2020-12-2 2 1\\n ObjectID(\\\"60c3501f9a9e8b7cdfa813d5\\\") 1 2 qwe 2020-12-2 2 1\\n ObjectID(\\\"60c3501f9a9e8b7cdfa813d4\\\") 1 2 qwe 2020-12-2 2 1\\n]\"\n"
	cmd := exec.
		Command("/bin/sh", "-c", command)
	cmd.Run()
	actualByte, _ := exec.Command("/bin/sh", "-c", "docker logs db-book").CombinedOutput()
	actualStr := string(actualByte)
	finalStr := getOneStr(actualStr, 39)
	assert.NotEqual(expected, finalStr)
}


func Test_Server_Read(t *testing.T) {
	assert := assert.New(t)
	command := "curl http://localhost:8081/book68c3503f9a9e8b7cdfa813d4"
	expected := "ObjectID(\\\"68c3503f9a9e8b7cdfa813d4\\\") 1 2 qwe 2020-12-2 2 1\\n\"\n"
	cmd := exec.
		Command("/bin/sh", "-c", command)
	cmd.Run()
	actualByte, _ := exec.Command("/bin/sh", "-c", "docker logs db-book").CombinedOutput()
	actualStr := string(actualByte)
	finalStr := getOneStr(actualStr, 39)
	assert.Equal(expected, finalStr)
}

func Test_Server_ReadNotEqual(t *testing.T) {
	assert := assert.New(t)
	command := "curl http://localhost:8081/books"
	expected := "ObjectID(\\\"68c3501f9a9e8b7cdfa813d7\\\") 1 2 qwe 2020-12-2 2 1\\n ObjectID(\\\"60c3501f9a9e8b7cdfa813d5\\\") 1 2 qwe 2020-12-2 2 1\\n ObjectID(\\\"60c3501f9a9e8b7cdfa813d4\\\") 1 2 qwe 2020-12-2 2 1\\n]\"\n"
	cmd := exec.
		Command("/bin/sh", "-c", command)
	cmd.Run()
	actualByte, _ := exec.Command("/bin/sh", "-c", "docker logs db-book").CombinedOutput()
	actualStr := string(actualByte)
	finalStr := getOneStr(actualStr, 39)
	assert.NotEqual(expected, finalStr)
}


func Test_Server_Update(t *testing.T){
	assert := assert.New(t)
	expectedId := "61c3501f9a9e8b7cdfa813d3\n"
	cmd := exec.
		Command("/bin/sh", "-c", "curl -i -X PUT http://localhost:8081/books61c3501f9a9e8b7cdfa813d3/432\n")
	cmd.Run()
	actualByte, _ := exec.Command("/bin/sh", "-c", "docker logs db-book").CombinedOutput()
	actualStr := string(actualByte)
	finalStr := getOneStr(actualStr, 38)
	assert.Equal(expectedId, finalStr)
}

func Test_Serve_Delete(t *testing.T){
	assert := assert.New(t)
	expectedId := "68c3503f9a9e8b7cdfa813d4\n"
	cmd := exec.
		Command("/bin/sh", "-c", "curl -i -X DELETE http://localhost:8081/books68c3503f9a9e8b7cdfa813d4\n")
	cmd.Run()
	actualByte, _ := exec.Command("/bin/sh", "-c", "docker logs db-book").CombinedOutput()
	actualStr := string(actualByte)
	finalStr := getOneStr(actualStr, 38)
	assert.Equal(expectedId, finalStr)
}
