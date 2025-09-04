package main

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

func (npc NPC) Play() {
	println("Playing as NPC:", npc.Name)
}

func (player Player) Play() {
	println("Playing as Player:", player.Name, "at level", player.Level)
}

// You can also use interfaces inside structs
type Group struct {
	Members []Character
}

func structWithInterfaceExample() {
	println("\nInterface and Struct Example:")

	group := Group{
		Members: []Character{NPC{Name: "Goblin"}, Player{Name: "Hero", Level: 10}},
	}

	for _, member := range group.Members {
		member.Play() // Calls the Play method of each member
	}
}