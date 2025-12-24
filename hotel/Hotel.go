package Hotel

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

func NewHotel() Hotel {
	return Hotel{Rooms: make(map[string]Room)}
}

func (h *Hotel) AddRoom(room Room) {
	h.Rooms[room.RoomNumber] = room
}

func (h *Hotel) CheckIn(roomNumber string) {
	room, ok := h.Rooms[roomNumber]
	if !ok {
		fmt.Println("Room not found")
		return
	}
	if room.IsOccupied {
		fmt.Println("Room already occupied")
		return
	}
	room.IsOccupied = true
	h.Rooms[roomNumber] = room
	fmt.Println("Checked in successfully")
}

func (h *Hotel) CheckOut(roomNumber string) {
	room, ok := h.Rooms[roomNumber]
	if !ok {
		fmt.Println("Room not found")
		return
	}
	room.IsOccupied = false
	h.Rooms[roomNumber] = room
	fmt.Println("Checked out successfully")
}

func (h *Hotel) ListVacantRooms() {
	fmt.Println("Vacant Rooms:")
	for _, room := range h.Rooms {
		if !room.IsOccupied {
			fmt.Println(room.RoomNumber, "-", room.Type, "-", room.PricePerNight)
		}
	}
}

func RunHotelMenu() {
	hotel := NewHotel()

	hotel.AddRoom(Room{"101", "Single", 10000, false})
	hotel.AddRoom(Room{"102", "Double", 15000, false})

	for {
		fmt.Println("\n--- HOTEL MENU ---")
		fmt.Println("1. Check In")
		fmt.Println("2. Check Out")
		fmt.Println("3. List Vacant Rooms")
		fmt.Println("0. Back")

		var choice int
		fmt.Scanln(&choice)

		var room string
		switch choice {
		case 1:
			fmt.Print("Room number: ")
			fmt.Scanln(&room)
			hotel.CheckIn(room)
		case 2:
			fmt.Print("Room number: ")
			fmt.Scanln(&room)
			hotel.CheckOut(room)
		case 3:
			hotel.ListVacantRooms()
		case 0:
			return
		}
	}
}
