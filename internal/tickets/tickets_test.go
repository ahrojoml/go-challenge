package tickets_test

import (
	"fmt"
	"testing"
	"tickets/internal/tickets"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetTotalTickets(t *testing.T) {
	testCases := []struct {
		name        string
		tickets     *[]tickets.Ticket
		country     string
		expect      int
		expectError error
	}{
		{
			name:        "nil tickets",
			country:     "Brazil",
			expect:      0,
			expectError: &tickets.NullTicketsError{},
		}, {
			name:        "no county",
			expect:      0,
			expectError: &tickets.CountryNotFoundError{},
		}, {
			name: "success",
			tickets: &[]tickets.Ticket{
				{
					Id:          1,
					Name:        "test name",
					Email:       "test@example.com",
					Destination: "Brazil",
					Time:        time.Now(),
					Price:       12345.34,
				},
			},
			country: "Brazil",
			expect:  1,
		},
	}

	for idx, tC := range testCases {
		t.Run(fmt.Sprintf("%d. %s", idx, tC.name), func(t *testing.T) {
			result, err := tickets.GetTotalTickets(tC.tickets, tC.country)

			if tC.expectError != nil {
				require.ErrorAs(t, err, tC.expectError)
			} else {
				require.Equal(t, tC.expect, result)
				require.NoError(t, err)
			}
		},
		)
	}
}

func TestGetCountByPeriod(t *testing.T) {
	testCases := []struct {
		name        string
		tickets     *[]tickets.Ticket
		period      tickets.Period
		expect      int
		expectError error
	}{
		{
			name:        "nil tickets",
			period:      tickets.Morning,
			expect:      0,
			expectError: &tickets.NullTicketsError{},
		}, {
			name: "success",
			tickets: &[]tickets.Ticket{
				{
					Id:          1,
					Name:        "test name",
					Email:       "test@example.com",
					Destination: "Brazil",
					Time:        time.Date(0, 0, 0, 0, 23, 0, 0, time.UTC),
					Price:       12345.34,
				},
				{
					Id:          1,
					Name:        "test name",
					Email:       "test@example.com",
					Destination: "Brazil",
					Time:        time.Date(0, 0, 0, 2, 23, 0, 0, time.UTC),
					Price:       12345.34,
				},
				{
					Id:          1,
					Name:        "test name",
					Email:       "test@example.com",
					Destination: "Brazil",
					Time:        time.Date(0, 0, 0, 18, 23, 0, 0, time.UTC),
					Price:       12345.34,
				},
			},
			period: tickets.EarlyMorning,
			expect: 2,
		},
	}

	for idx, tC := range testCases {
		t.Run(fmt.Sprintf("%d. %s", idx, tC.name), func(t *testing.T) {
			result, err := tickets.GetCountByPeriod(tC.tickets, tC.period)

			if tC.expectError != nil {
				require.ErrorAs(t, err, tC.expectError)
			} else {
				require.Equal(t, tC.expect, result)
			}
		},
		)
	}
}

func TestDestinationPercentage(t *testing.T) {
	testCases := []struct {
		name        string
		tickets     *[]tickets.Ticket
		destination string
		expect      float64
		expectError error
	}{
		{
			name:        "nil tickets",
			destination: "Brazil",
			expect:      0,
			expectError: &tickets.NullTicketsError{},
		}, {
			name:        "no county",
			expect:      0,
			expectError: &tickets.CountryNotFoundError{},
		}, {
			name: "success",
			tickets: &[]tickets.Ticket{
				{
					Id:          1,
					Name:        "test name",
					Email:       "test@example.com",
					Destination: "Brazil",
					Time:        time.Date(0, 0, 0, 0, 23, 0, 0, time.UTC),
					Price:       12345.34,
				},
				{
					Id:          1,
					Name:        "test name",
					Email:       "test@example.com",
					Destination: "Madagascar",
					Time:        time.Date(0, 0, 0, 2, 23, 0, 0, time.UTC),
					Price:       12345.34,
				},
				{
					Id:          1,
					Name:        "test name",
					Email:       "test@example.com",
					Destination: "Brazil",
					Time:        time.Date(0, 0, 0, 18, 23, 0, 0, time.UTC),
					Price:       12345.34,
				},
			},
			destination: "Brazil",
			expect:      2.0 / 3.0,
		},
	}

	for idx, tC := range testCases {
		t.Run(fmt.Sprintf("%d. %s", idx, tC.name), func(t *testing.T) {
			result, err := tickets.DestinationPercentage(tC.tickets, tC.destination)

			if tC.expectError != nil {
				require.ErrorAs(t, err, tC.expectError)
			} else {
				require.Equal(t, tC.expect, result)
			}
		},
		)
	}
}

func TestTimeToPeriod(t *testing.T) {
	testCases := []struct {
		name   string
		date   time.Time
		expect tickets.Period
	}{
		{
			name:   "early morning",
			date:   time.Date(0, 0, 0, 1, 23, 0, 0, time.UTC),
			expect: tickets.EarlyMorning,
		}, {
			name:   "morning",
			date:   time.Date(0, 0, 0, 8, 23, 0, 0, time.UTC),
			expect: tickets.Morning,
		}, {
			name:   "afternoon",
			date:   time.Date(0, 0, 0, 15, 23, 0, 0, time.UTC),
			expect: tickets.Afternoon,
		}, {
			name:   "evening",
			date:   time.Date(0, 0, 0, 23, 23, 0, 0, time.UTC),
			expect: tickets.Evening,
		},
	}

	for idx, tC := range testCases {
		t.Run(fmt.Sprintf("%d. %s", idx, tC.name), func(t *testing.T) {
			result := tickets.TimeToPeriod(tC.date)

			require.Equal(t, tC.expect, result)
		},
		)
	}
}
