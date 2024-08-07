package main

import "fmt"

func main() {
	// var m otherMonkey
	// m = monkeyImpl{}
	var peo People = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

type monkey interface {
	eat()
	sleep()
	run()
}

type goldMonkey struct{}

func (goldMonkey goldMonkey) sleep() {
	// TODO implement me
	panic("implement me")
}

type otherMonkey interface {
	monkey
	walk()
}

type monkeyImpl struct{}

func (monkeyImpl) walk() {
	panic("implement walk")
}

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}
