package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	ticketslib "tickets/internal/tickets"
	"time"
)

var tickets = make([]ticketslib.Ticket, 0, 1000)

const timeFormat = "15:04"

func main() {
	file, err := os.Open("tickets.csv")
	if err != nil {
		panic("file not found")
	}

	reader := csv.NewReader(file)
	var fields []string
	var ticket ticketslib.Ticket
	var id int
	var price float64
	var date time.Time

	for {
		fields, err = reader.Read()
		if err != nil {
			if errors.Is(io.EOF, err) {
				break
			}
			fmt.Println("line could not be read")
			continue
		}

		id, err = strconv.Atoi(fields[0])
		if err != nil {
			fmt.Printf("could not read field id: %s, skipping\n", fields[0])
			continue
		}

		date, err = time.Parse(timeFormat, fields[4])
		if err != nil {
			fmt.Printf("could not read field time: %s, skipping\n", fields[4])
			continue
		}

		price, err = strconv.ParseFloat(fields[5], 64)
		if err != nil {
			fmt.Printf("could not read field price: %s, skipping\n", fields[5])
			continue
		}

		ticket = ticketslib.Ticket{
			Id:          id,
			Name:        fields[1],
			Email:       fields[2],
			Destination: fields[3],
			Time:        date,
			Price:       price,
		}

		tickets = append(tickets, ticket)
	}

	fmt.Printf("total tickets: %d\n", len(tickets))

	totalTickets, err := ticketslib.GetTotalTickets(&tickets, "Brazil")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("GetTotalTickets: %d\n", totalTickets)

	totalTickets, err = ticketslib.GetCountByPeriod(&tickets, ticketslib.EarlyMorning)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("GetCountByPeriod: %d\n", totalTickets)

	destinationPercentage, err := ticketslib.DestinationPercentage(&tickets, "Brazil")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("DestinationPercentage: %.3f\n", destinationPercentage)
}
