package main

import (
	"sync"
	"log"
	"time"
	"fmt"
	"unsafe"
	"io"
	"os"
)

func main() {
	var mailbox uint8
	var lock sync.Mutex
	sendCond := sync.NewCond(&lock)
	recvCond := sync.NewCond(&lock)

	send := func(id, index int) {
		lock.Lock()
		for mailbox == 1 {
			sendCond.Wait()
		}
		log.Printf("sender [%d:%d] the mailbox is empty.",id,index)
		mailbox = 1
		log.Printf("sender [%d-%d]: the letter has been sent",id,index)
		lock.Unlock()
		//唤醒一个因此等待的go
		recvCond.Broadcast()
	}

	recv := func(id ,index int) {
		lock.Lock()
		for mailbox == 0 {
			recvCond.Wait()
		}
		log.Printf("recver [%d-%d] the mailbox is full.",id,index)
		mailbox = 0
		log.Printf("recver [%d-%d] the letter has been received.",id,index)
		lock.Unlock()
		sendCond.Signal()
	}

	sign := make(chan struct{},3)
	max := 10
	go func(id,max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i :=0 ;i< max;i++ {
			send(1,i)
			time.Sleep(500 * time.Millisecond)
		}
	}(0 ,max)

	go func(id,max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i:=0;i<max;i++{
			recv(1,i)
			time.Sleep(500*time.Millisecond)
		}
	}(1,max/2)

	go func(id,max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i:=0;i<max;i++{
			recv(1,i)
			time.Sleep(500*time.Millisecond)
		}
	}(2,max/2)
	//发送者的数量要比接收者数量多才能保证不会出现deadlock，
	// 如果要保证所有信都被成功收到则发送者和接收者数量应该一样
	<- sign
	<- sign
	<- sign
	var x struct{
		a bool
		b int16
		c []int
	}
	pb := (*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb =42
	fmt.Println("x:%T",x)
	w :=(io.Writer)(os.Stdout)
	fmt.Println(w)
}