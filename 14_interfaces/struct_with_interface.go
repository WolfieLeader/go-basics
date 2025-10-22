package main

import "fmt"

type NPC struct {
	Name string
}

type Player struct {
	Name  string
	Level int
}

type Character interface {
	Play()
}

// You can also use interfaces inside structs
type Group struct {
	Members []Character
}

func (npc NPC) Play() {
	fmt.Println("- Playing as NPC:", npc.Name)
}

func (player Player) Play() {
	fmt.Println("- Playing as Player:", player.Name, "at level", player.Level)
}

func structWithInterfaceExample() {
	group := Group{
		Members: []Character{NPC{Name: "Goblin"}, Player{Name: "Hero", Level: 10}},
	}

	for _, member := range group.Members {
		member.Play() // Calls the Play method of each member
	}
}
