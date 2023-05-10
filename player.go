package main

type Team struct {
	club    string
	country string
}

type Player struct {
	firstname string
	lastname  string
	team      Team
}

func (playerPointer *Player) changeClub(newClub string) {
	(*playerPointer).team.club = newClub
}
