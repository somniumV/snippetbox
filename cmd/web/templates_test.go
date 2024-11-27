package main

import (
	"snippetbox.somniumV.net/internal/assert"
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2022, 2, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Feb 2022 at 10:15",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2022, 2, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Feb 2022 at 09:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, tt.want, hd)
		})
	}
}
