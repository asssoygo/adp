package gym

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

type Gym struct {
	Members map[uint64]Member
}

func (g *Gym) AddMember(id uint64, member Member) {
	g.Members[id] = member
}

func (g *Gym) ListMembers() {
	for id, m := range g.Members {
		fmt.Println(id, m.GetDetails())
	}
}
