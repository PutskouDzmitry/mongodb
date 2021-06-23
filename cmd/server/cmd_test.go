package main

import (
	"github.com/stretchr/testify/require"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func init() {
	time.Sleep(2 * time.Second)
}

func getMessage(command string, symbol string) string {
	cmd, _ := exec.
		Command("/bin/sh", "-c", command).CombinedOutput()
	actualStr := string(cmd)
	index := strings.LastIndex(actualStr, symbol)
	return actualStr[index:]
}

func Test_Server_Add_True(t *testing.T){
	require := require.New(t)
	command := `curl -i -X POST -H 'Content-Type: application/json' -d '{"BookId":"62c3503f9a9e8b7cdfa813d4","AuthorId":1,"BookVolume":2,"NameOfBook":"qwe","Number":1,"PublisherId":2,"YearOfPublication":"2020-12-2"}' http://localhost:8081/books`
	cmd, _ := exec.
		Command("/bin/sh", "-c", command).CombinedOutput()
	actual := string(cmd)
	require.True(strings.Contains(actual, "200 OK"))
}

func Test_Server_Add_False(t *testing.T){
	require := require.New(t)
	command := `curl -i -X POST -H 'Content-Type: application/json' -d '{"BookId":"62c3503f9a9e8b7cdfa813d4","AuthorId":1,"BookVolume":2,"NameOfBook":"qwe","Number":1,"PublisherId":2,"YearOfPublication":"2020-12-2"}' http://localhost:8080/books`
	cmd, _ := exec.
		Command("/bin/sh", "-c", command).CombinedOutput()
	actual := string(cmd)
	require.False(strings.Contains(actual, "200 OK"))
}

func Test_Server_Add_Contains(t *testing.T){
	require := require.New(t)
	command := `curl -i -X POST -H 'Content-Type: application/json' -d '{"BookId":"66c3503f9a9e8b7cdfa813d4","AuthorId":1,"BookVolume":2,"NameOfBook":"qwe","Number":1,"PublisherId":2,"YearOfPublication":"2020-12-2"}' http://localhost:8081/books`
	cmd, _ := exec.
		Command("/bin/sh", "-c", command).CombinedOutput()
	actual := string(cmd)
	require.Contains(actual, "200 OK")
}

func Test_Server_Delete_Contains(t *testing.T){
	require := require.New(t)
	command := "curl -i -X DELETE http://localhost:8081/books66c3503f9a9e8b7cdfa813d4"
	cmd, _ := exec.
		Command("/bin/sh", "-c", command).CombinedOutput()
	actual := string(cmd)
	require.Contains(actual, "200 OK")
}

func Test_Server_ReadAlL_NotEmpty(t *testing.T) {
	require := require.New(t)
	command := `curl http://localhost:8081/books -H "Accept: application/json"`
	finalMessage := strings.TrimSpace(getMessage(command, "[{"))
	require.NotEmpty(finalMessage)
}

func Test_Server_ReadAll(t *testing.T) {
	require := require.New(t)
	command := `curl http://localhost:8081/books -H "Accept: application/json"`
	expected := `[{"BookId":"62c3503f9a9e8b7cdfa813d4","AuthorId":1,"BookVolume":2,"NameOfBook":"qwe","Number":1,"PublisherId":2,"YearOfPublication":"2020-12-2"}]`
	finalMessage := strings.TrimSpace(getMessage(command, "[{"))
	require.Equal(expected, finalMessage)
}

func Test_Server_ReadAllNotEqual(t *testing.T) {
	require := require.New(t)
	command := `curl http://localhost:8081/books -H "Accept: application/json"`
	expected := `[ObjectID("68c3503f9a9e8b7cdfa813d4") 1 2 qwe 2020-12-2]`
	require.NotEqual(expected, getMessage(command, "[{"))
}


func Test_Server_Read(t *testing.T) {
	require := require.New(t)
	command := "curl http://localhost:8081/book62c3503f9a9e8b7cdfa813d4"
	expected := `{"BookId":"62c3503f9a9e8b7cdfa813d4","AuthorId":1,"BookVolume":2,"NameOfBook":"qwe","Number":1,"PublisherId":2,"YearOfPublication":"2020-12-2"}`
	finalMessage := strings.TrimSpace(getMessage(command, `{"`))
	require.Equal(expected, finalMessage)
}

func Test_Server_ReadNotEqual(t *testing.T) {
	require := require.New(t)
	command := "curl http://localhost:8081/book62c3503f9a9e8b7cdfa813d4"
	expected := `{"BookId":"62c3503f9a9e8b7cdfa813d4","AuthorId":2,"BookVolume":2,"NameOfBook":"qwe","Number":1,"PublisherId":2,"YearOfPublication":"2020-12-2"}`
	finalMessage := strings.TrimSpace(getMessage(command, `{"`))
	require.NotEqual(expected, finalMessage)
}


func Test_Server_Update_True(t *testing.T){
	require := require.New(t)
	command := "curl -i -X PUT http://localhost:8081/books62c3503f9a9e8b7cdfa813d4/213"
	cmd, _ := exec.
		Command("/bin/sh", "-c", command).CombinedOutput()
	actual := string(cmd)
	require.True(strings.Contains(actual, "200 OK"))
}

func Test_Server_Update_Contains(t *testing.T){
	require := require.New(t)
	command := "curl -i -X PUT http://localhost:8081/books62c3503f9a9e8b7cdfa813d4/213"
	cmd, _ := exec.
		Command("/bin/sh", "-c", command).CombinedOutput()
	actual := string(cmd)
	require.Contains(actual, "200 OK")
}

func Test_Server_Update_False(t *testing.T){
	require := require.New(t)
	command := "curl -i -X PUT http://localhost:8080/books64c3503f9a9e8b7cdfa813d4/213"
	cmd, _ := exec.
		Command("/bin/sh", "-c", command).CombinedOutput()
	actual := string(cmd)
	require.False(strings.Contains(actual, "200 OK"))
}


func Test_Server_Delete_False(t *testing.T){
	require := require.New(t)
	command := "curl -i -X DELETE http://localhost:8080/books62c3503f9a9e8b7cdfa813d4"
	cmd, _ := exec.
		Command("/bin/sh", "-c", command).CombinedOutput()
	actual := string(cmd)
	require.False(strings.Contains(actual, "200 OK"))
}


func Test_Server_Delete_True(t *testing.T){
	require := require.New(t)
	command := "curl -i -X DELETE http://localhost:8081/books62c3503f9a9e8b7cdfa813d4"
	cmd, _ := exec.
		Command("/bin/sh", "-c", command).CombinedOutput()
	actual := string(cmd)
	require.True(strings.Contains(actual, "200 OK"))
}