package main

import (
	"crypto/md5"
	"io/ioutil"
	"unsafe"
)

const (
	MAX_RANDOM_TASKS = 32
)

//This function is used for get max tasks with input string
//input string is formated : '${USERID}/YYYYMMDD'
//Consistent hash function : (md5(input) 8 bytes % MAX_RANDOM_TASKS) + 1
func GetMaxTasks(input string) uint {
	hash := md5.Sum([]byte(input))

	//hash is 16 bytes, uint golang is 8 bytes
	//so get the first 8 bytes in hash only
	hashData := hash[0:8]
	//get the first data pointer to convert to int
	uintHash := *(*uint)(unsafe.Pointer(&hashData[0]))

	maxTasks := uintHash % MAX_RANDOM_TASKS

	//don't want zero max tasks
	return maxTasks + 1
}

func CountFile(folder string) int {
	files, _ := ioutil.ReadDir(folder)
	return len(files)
}
