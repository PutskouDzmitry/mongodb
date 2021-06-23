package main

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getMessage(command string, symbol string) string {
	cmd, _ := exec.
		Command("/bin/sh", "-c", command).CombinedOutput()
	actualStr := string(cmd)
	index := strings.LastIndex(actualStr, symbol)
	return actualStr[index:]
}

func Test_Server_Add(t *testing.T){
	assert := assert.New(t)
	expected := true
	command := `curl -i -X POST -H 'Content-Type: application/json' -d '{"BookId":"62c3503f9a9e8b7cdfa813d4","AuthorId":1,"BookVolume":2,"NameOfBook":"qwe","Number":1,"PublisherId":2,"YearOfPublication":"2020-12-2"}' http://localhost:8081/books`
	cmd, _ := exec.
		Command("/bin/sh", "-c", command).CombinedOutput()
	actual := string(cmd)
	assert.Equal(expected, strings.Contains(actual, "200 OK"))
}

func Test_Server_ReadAll(t *testing.T) {
	assert := assert.New(t)
	command := `curl http://localhost:8081/books -H "Accept: application/json"`
	expected := `[{"BookId":"62c3503f9a9e8b7cdfa813d4","AuthorId":1,"BookVolume":2,"NameOfBook":"qwe","Number":1,"PublisherId":2,"YearOfPublication":"2020-12-2"}]`
	finalMessage := strings.TrimSpace(getMessage(command, "[{"))
	assert.Equal(expected, finalMessage)
}

func Test_Server_ReadAllNotEqual(t *testing.T) {
	assert := assert.New(t)
	command := `curl http://localhost:8081/books -H "Accept: application/json"`
	expected := `[ObjectID("68c3503f9a9e8b7cdfa813d4") 1 2 qwe 2020-12-2]`
	assert.NotEqual(expected, getMessage(command, "[{"))
}


func Test_Server_Read(t *testing.T) {
	assert := assert.New(t)
	command := "curl http://localhost:8081/book62c3503f9a9e8b7cdfa813d4"
	expected := `{"BookId":"62c3503f9a9e8b7cdfa813d4","AuthorId":1,"BookVolume":2,"NameOfBook":"qwe","Number":1,"PublisherId":2,"YearOfPublication":"2020-12-2"}`
	finalMessage := strings.TrimSpace(getMessage(command, `{"`))
	assert.Equal(expected, finalMessage)
}

func Test_Server_ReadNotEqual(t *testing.T) {
	assert := assert.New(t)
	command := "curl http://localhost:8081/book62c3503f9a9e8b7cdfa813d4"
	expected := `{"BookId":"62c3503f9a9e8b7cdfa813d4","AuthorId":2,"BookVolume":2,"NameOfBook":"qwe","Number":1,"PublisherId":2,"YearOfPublication":"2020-12-2"}`
	finalMessage := strings.TrimSpace(getMessage(command, `{"`))
	assert.NotEqual(expected, finalMessage)
}


func Test_Server_Update(t *testing.T){
	assert := assert.New(t)
	expected := true
	command := "curl -i -X PUT http://localhost:8081/books62c3503f9a9e8b7cdfa813d4/213"
	cmd, _ := exec.
		Command("/bin/sh", "-c", command).CombinedOutput()
	actual := string(cmd)
	assert.Equal(expected, strings.Contains(actual, "200 OK"))
}

func Test_Serve_Delete(t *testing.T){
	assert := assert.New(t)
	command := "curl -i -X DELETE http://localhost:8081/books62c3503f9a9e8b7cdfa813d4"
	expected := true
	cmd, _ := exec.
		Command("/bin/sh", "-c", command).CombinedOutput()
	actual := string(cmd)
	assert.Equal(expected, strings.Contains(actual, "200 OK"))
}
