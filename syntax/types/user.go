package main

import (
	"fmt"
)

// NewUser 如何初始化
func NewUser() {
	// 初始化结构体
	u := User{}
	fmt.Printf("u %v \n", u)
	fmt.Printf("u %+v \n", u)

	// up 是一个指针，指向 User 的地址
	up := &User{}
	fmt.Printf("up %v \n", up)
	fmt.Printf("up %+v \n", up)
	up2 := new(User)
	println(up2.FirstName)
	fmt.Printf("up2 %+v \n", up2)

	u4 := User{Name: "Coconut", Age: 18} // 注意使用这个，更推荐
	u5 := User{"胖蛋", "4", 4}             // 不推荐
	fmt.Printf("u4 %+v \n", u4)
	fmt.Printf("u5 %+v \n", u5)

	u4.Name = "Judy"
	u5.Age = 18

	var up3 *User
	// 报错，在 nil 上访问字段，或者方法，"空指针"
	//println(up3.FirstName)
	println(up3)
}

type User struct {
	Name      string
	FirstName string
	Age       int
}

func (u User) ChangeName(name string) {
	fmt.Printf("ChangName u 的地址 %p \n", &u)
	u.Name = name
}

// ChangeNameV1  和 ChangeName，是等同的，会生成新的变量
func ChangeNameV1(u User, name string) {
	fmt.Printf("ChangName u 的地址 %p \n", &u)
	u.Name = name
}

// ChangeAge 遇事不决用指针
func (u *User) ChangeAge(age int) {
	fmt.Printf("ChangAge u 的地址 %p \n", u) // 这里不需要取地址了，因为本身就是指针，就是地址
	u.Age = age
}

// ChangeAgeV1 和 ChangeAge，是等同的，指针复制后，还是指向相同的地址
func ChangeAgeV1(u *User, age int) {
	fmt.Printf("ChangAge u 的地址 %p \n", u) // 这里不需要取地址了，因为本身就是指针，就是地址
	u.Age = age
}

func ChangeUser() {
	u1 := User{Name: "Judy", Age: 18}
	fmt.Printf("u1 的地址是 %p \n", &u1)
	// 这一步执行的时候，相当于复制了一个 u1，改的是复制体
	// u1 的 Name 原封不动
	u1.ChangeName("Bob")
	// 等价于 (&u1).ChangeAge(35)
	u1.ChangeAge(35)
	//u2 = User{Name: "Judy", Age: 18}
	fmt.Printf("%+v \n", u1)

	u2 := &User{}
	fmt.Printf("u2 的地址是 %p \n", &u2)
	u2.ChangeAge(35)
	// 等价于 (*u2).ChangeName("Rose")
	u2.ChangeName("Rose")
	fmt.Printf("%+v \n", u2)

	// u1 的地址是 0x1400006c180
	// ChangName u 的地址 0x1400006c1b0
	// ChangAge u 的地址 0x1400006c180
	// {Name:Judy FirstName: Age:35}
}
