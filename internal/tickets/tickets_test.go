package tickets_test

import (
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
		expectError bool
	}{
		{
			name:        "nil tickets",
			country:     "Brazil",
			expect:      0,
			expectError: true,
		}, {
			name:        "no county",
			expect:      0,
			expectError: true,
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
			country:     "Brazil",
			expect:      1,
			expectError: false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result, err := tickets.GetTotalTickets(tC.tickets, tC.country)

			if tC.expectError {
				require.Error(t, err)
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
		expectError bool
	}{
		{
			name:        "nil tickets",
			period:      tickets.Morning,
			expect:      0,
			expectError: true,
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
			period:      tickets.EarlyMorning,
			expect:      2,
			expectError: false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result, err := tickets.GetCountByPeriod(tC.tickets, tC.period)

			if tC.expectError {
				require.Error(t, err)
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
		expectError bool
	}{
		{
			name:        "nil tickets",
			destination: "Brazil",
			expect:      0,
			expectError: true,
		}, {
			name:        "no county",
			expect:      0,
			expectError: true,
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
			expectError: false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result, err := tickets.DestinationPercentage(tC.tickets, tC.destination)

			if tC.expectError {
				require.Error(t, err)
			} else {
				require.Equal(t, tC.expect, result)
			}
		},
		)
	}
}
