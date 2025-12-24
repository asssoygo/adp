package Gym

import "fmt"

type Member interface {
	GetDetails() string
}

type BasicMember struct {
	Name string
}

func (b BasicMember) GetDetails() string {
	return "Basic Member: " + b.Name
}

type PremiumMember struct {
	Name string
}

func (p PremiumMember) GetDetails() string {
	return "Premium Member: " + p.Name
}

type Gym struct {
	Members map[uint64]Member
}

func NewGym() Gym {
	return Gym{Members: make(map[uint64]Member)}
}

func (g *Gym) AddMember(id uint64, m Member) {
	g.Members[id] = m
}

func (g *Gym) ListMembers() {
	for id, member := range g.Members {
		fmt.Println(id, "->", member.GetDetails())
	}
}

func RunGymDemo() {
	gym := NewGym()

	gym.AddMember(1, BasicMember{"Ali"})
	gym.AddMember(2, PremiumMember{"Aruzhan"})

	gym.ListMembers()
}
