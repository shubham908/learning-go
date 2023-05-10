package main

import "fmt"

func main() {
	player := Player{
		firstname: "Victor",
		lastname:  "Oshimen",
		team: Team{
			club:    "Napoli",
			country: "Nigeria",
		},
	}

	playerPointer := &player
	playerPointer.changeClub("Chelsea")

	fmt.Printf("%+v", player)
	fmt.Println()
}
