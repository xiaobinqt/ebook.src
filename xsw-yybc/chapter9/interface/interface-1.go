package main

import "fmt"

type ISpeaker interface {
	Speak()
}

type SimpleSpeaker struct {
	Message string
}

func (speaker *SimpleSpeaker) Speak() {
	fmt.Println("i am speaking ? ", speaker.Message)
}

func main() {
	var speaker ISpeaker
	speaker = &SimpleSpeaker{
		Message: "Hello....",
	}

	speaker.Speak()

}
