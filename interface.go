package main

import "fmt"

type Animal interface{
    speak() string
}

type Dog struct{
}

func (d *Dog) speak() string{
    return "woof"
}

type Cat struct{
}

func (c *Cat) speak() string{
    return "meaow"
}

func noise(a Animal){
    fmt.Println(a.speak())
}

func main(){
    d := &Dog{}
    animals := []Animal{d, new(Dog), new(Cat)}
    for _,a := range animals{
        noise(a)
    }
}
