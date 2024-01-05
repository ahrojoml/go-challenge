package tickets

import "time"

type Ticket struct {
	Id          int
	Name        string
	Email       string
	Destination string
	Time        time.Time
	Price       float64
}

func GetTotalTickets(tickets *[]Ticket, destination string) (int, error) {
	if destination == "" {
		return 0, NewCountryNotFoundError("no country of destination was given")
	}
	if tickets == nil {
		return 0, NewNullTicketsError("no ticket data was given")
	}

	var counter int = 0

	for _, ticket := range *tickets {
		if ticket.Destination == destination {
			counter += 1
		}
	}

	return counter, nil
}

func GetMornings(tickets []Ticket, time string) (int, error) {
	return 0, nil
}

func AverageDestination(tickets []Ticket, destination string, total int) (int, error) {
	return 0, nil
}
