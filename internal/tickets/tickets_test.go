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
			fmt.Println(tC)
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
