package main

import (
	"fmt"
)

type Animal interface {
	// ScientificName 用于获取动物的学名。
	ScientificName() string
	// Category 用于获取动物的基本分类。
	Category() string
}

type Named interface {
	// Name 用于获取名字。
	Name() string
	String() string
}

type Pet interface {
	Animal
	Named
}

type Other interface {
	String() string
}

type PetTag struct {
	name  string
	owner string
}

func (pt PetTag) String() string {
	return fmt.Sprintf("String")
}

func (pt PetTag) Name() string {
	return pt.name
}

func (pt PetTag) Owner() string {
	return pt.owner
}

type Dog struct {
	PetTag
	scientificName string
}

func (dog Dog) ScientificName() string {
	return dog.scientificName
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	petTag := PetTag{name: "little pig"}
	_, ok := interface{}(petTag).(Named)
	fmt.Printf("PetTag implements interface Named: %v\n", ok)
	dog := Dog{
		PetTag:         petTag,
		scientificName: "Labrador Retriever",
	}
	_, ok = interface{}(dog).(Animal)
	//Dog直接实现了 Animal的 ScientificName 和 Category 这两个接口（全部接口）
	fmt.Printf("Dog implements interface Animal: %v\n", ok)
	_, ok = interface{}(dog).(Named)
	//Dog虽然没有直接实现Named接口，但是其中嵌入PetTag实现了Named接口的Name（全部接口）
	fmt.Printf("Dog implements interface Named: %v\n", ok)
	_, ok = interface{}(dog).(Pet)
	//Dog直接实现了Animal的全部接口，内嵌PetTag实现了Named全部接口
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	fmt.Println()
	_, ok = interface{}(petTag).(Other)
	fmt.Printf("PetTag implements interface Other :%v\n", ok)
	fmt.Printf("petTag.other:%s", petTag)
}
