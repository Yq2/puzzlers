package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var SignSecret string  = "code_rand"

func main() {
	for i := 0; i< 1000; i++ {
		time.Sleep(1 * time.Millisecond)
		code := Random(1000)
		str := strconv.Itoa(code)
		fmt.Printf("%s\n",CreateSign(str))
	}
}

func Random(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if max < 1 {
		return r.Int()
	} else {
		return r.Intn(max)
	}
}

func CreateSign(str string) string {
	SignSecret += fmt.Sprintf("_%d",Random(88))
	str = string(SignSecret)
	sign := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return sign
}