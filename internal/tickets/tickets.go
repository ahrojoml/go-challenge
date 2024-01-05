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

type Period int

const (
	EarlyMorning Period = iota
	Morning
	Afternoon
	Evening
)

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

func GetCountByPeriod(tickets *[]Ticket, period Period) (int, error) {
	if tickets == nil {
		return 0, NewNullTicketsError("no ticket data was given")
	}

	var counter int = 0

	for _, ticket := range *tickets {
		periodFromTime := TimeToPeriod(ticket.Time)

		if periodFromTime == period {
			counter += 1
		}
	}

	return counter, nil
}

func TimeToPeriod(clockTime time.Time) Period {
	hour := clockTime.Hour()

	switch {
	case 0 <= hour && hour <= 6:
		return EarlyMorning
	case 7 <= hour && hour <= 12:
		return Morning
	case 13 <= hour && hour <= 19:
		return Afternoon
	}

	// since we are receiving a Time the value hour will always be between 0 and 23
	return Evening
}

func DestinationPercentage(tickets *[]Ticket, destination string) (float64, error) {
	if destination == "" {
		return 0, NewCountryNotFoundError("no country of destination was given")
	}

	if tickets == nil {
		return 0, NewNullTicketsError("no ticket data was given")
	}

	var counter float64

	for _, ticket := range *tickets {
		if ticket.Destination == destination {
			counter += 1
		}
	}

	return counter / float64(len(*tickets)), nil
}
