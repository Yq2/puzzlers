package main

import "fmt"

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

type Other interface {
	SetName(name string)
}

type Dog struct {
	name string // 名字。
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	// 示例1。
	dog := Dog{"little pig"}
	_, ok := interface{}(dog).(Pet)
	//Dog没有实现SetName方法，*Dog类型实现了SetName方法，但是指针类型方法不包含在结构体类型方法中
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	_, ok = interface{}(&dog).(Pet)
	//Dog类型实现了Name方法和Category方法，*Dog类型实现了SetName方法，又由于同一个结构体类型的指针方法会包含结构体方法
	//所以*Dog类型实现了Pet接口的所有方法
	fmt.Printf("*Dog implements interface Pet: %v\n", ok)
	//dog.SetName("Big pig") 等价于 &dog.SetName("Big pig")

	//_ := &dog.Category() 编译报错
	//&dog.SetName("jj")  编译报错
	fmt.Println()

	// 示例2。
	//var pet Pet = dog  编译报错，Dog类型只实现了Pet接口的Name方法和Category方法，并没有实现SetName方法
	var pet Pet = &dog

	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
	//var _ Other = dog
}
