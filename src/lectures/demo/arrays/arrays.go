package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "Living Room", cleaned: true},
		{name: "Kitchen", cleaned: false},
		{name: "Bedroom", cleaned: true},
		{name: "Bathroom", cleaned: false},
	}
	checkCleanliness(rooms)

	fmt.Println("Performing some cleaning...")
	rooms[1].cleaned = true
	rooms[2].cleaned = true

	checkCleanliness(rooms)
}
