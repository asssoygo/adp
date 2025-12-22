package hotel

import "fmt"

type Room struct {
	RoomNumber    string
	Type          string
	PricePerNight float64
	IsOccupied    bool
}

type Hotel struct {
	Rooms map[string]Room
}

func (h *Hotel) AddRoom(room Room) {
	h.Rooms[room.RoomNumber] = room
}

func (h *Hotel) CheckIn(roomNumber string) {
	room := h.Rooms[roomNumber]
	room.IsOccupied = true
	h.Rooms[roomNumber] = room
}

func (h *Hotel) ListVacantRooms() {
	for _, room := range h.Rooms {
		if !room.IsOccupied {
			fmt.Println(room.RoomNumber, room.Type, room.PricePerNight)
		}
	}
}
