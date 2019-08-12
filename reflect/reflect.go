package reflect

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (u User) Print(tag bool){
	if tag {
		fmt.Printf("my name is %s, i am %d years old\n", u.Name, u.Age)
	} else {
		fmt.Println("boronya can not say")
	}
}

func Testreflect() {
	user := User{"cfh", 1}
	userType := reflect.TypeOf(user)
	fmt.Println(userType)
	userValue := reflect.ValueOf(user)
	fmt.Println(userValue.Interface().(User).Name)
	fmt.Println(userType.Field(0).Name)
	fmt.Println(userType.Field(0).Type)
	fmt.Println(reflect.ValueOf(&user).Elem().Field(0).CanSet())
	fmt.Println(reflect.ValueOf(&user).Elem().Field(0).Interface())
	reflect.ValueOf(&user).Elem().Field(0).SetString("yj")
	fmt.Println(user)
	bol := true
	args := []reflect.Value{reflect.ValueOf(bol)}
	reflect.ValueOf(&user).MethodByName("Print").Call(args)

	var a float64 = 6.4
	floatType := reflect.TypeOf(a)
	fmt.Println(floatType)
	fmt.Println(reflect.ValueOf(a).Interface())
}
