package main

import "fmt"

type Passenger struct {
	Name string
	TicketNumber int
	Boarded bool
}

type Bus struct {
	FrontSeat Passenger // Passenger is the type of the FrontSeat field
}

func main() {
	casey := Passenger{"casey", 1, false}
	fmt.Println(casey)

	var (
		bill = Passenger{Name: "bill", TicketNumber: 2}
		ella = Passenger{Name: "ella", TicketNumber: 3}
	)
	fmt.Println(bill, ella)

	var heidi Passenger
	heidi.Name = "heidi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	casey.Boarded = true
	bill.Boarded = true
	if bill.Boarded {
		fmt.Println("Bill has boarded")
	} else {
		fmt.Println("Bill has not boarded")
	}

	if casey.Boarded {
		fmt.Println("Casey has boarded")
	} else {
		fmt.Println("Casey has not boarded")
	}

	heidi.Boarded = true
	bus := Bus{FrontSeat: heidi}
	if heidi.Boarded {
		fmt.Println("Heidi has boarded")
	} else {
		fmt.Println("Heidi has not boarded")
	}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")
}

