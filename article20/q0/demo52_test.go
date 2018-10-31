package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	var name string
	greeting, err := hello(name)
	if err == nil {
		t.Errorf("The error is nil, but it should not be. (name=%q)",
			name)
	}
	if greeting != "" {
		t.Errorf("Nonempty greeting, but it should not be. (name=%q)",
			name)
	}
	name = "Robert"
	greeting, err = hello(name)
	if err != nil {
		t.Errorf("The error is not nil, but it should be. (name=%q)",
			name)
	}
	if greeting == "" {
		t.Errorf("Empty greeting, but it should not be. (name=%q)",
			name)
	}
	expected := fmt.Sprintf("Hello, %s!", name)
	if greeting != expected {
		t.Errorf("The actual greeting %q is not the expected. (name=%q)",
			greeting, name)
	}
	t.Logf("The expected greeting is %q.\n", expected)
}

func TestWorld(t *testing.T) {
	var name string
	greeting,err := hello(name)
	if err == nil {
		t.Errorf("The error is null,but it should not be .(name=%q)",name)
	}
	if greeting != "" {
		t.Errorf("Noempty greeting ,but it should not be.(name=%q)",name)
	}
	name = "Rebert"
	greeting, err = hello(name)
	if err != nil {
		t.Errorf("The error is not nil ,but it should be. (name=%q)",name)
	}
	if greeting == "" {
		t.Errorf("Empty greeting ,but it should bot be.(name=%q)",name)
	}
	expected := fmt.Sprintf("Hello ,%s",name)
	if greeting != expected {
		t.Errorf("The actual greeting %q is not the expected . (name=%q)",greeting,name)
	}
	t.Logf("The expected greeting is %q.\n",expected)
}

//不执行的原因是 不是以Test开头的函数名
func TesIntroduce(t *testing.T) {
	intro := introduce()
	expected := "Welcome to my Golang column."
	if intro != expected {
		t.Errorf("The actual introduce %q is not the expected.",
			intro)
	}
	t.Logf("The expected introduce is %q.\n", expected)
}

