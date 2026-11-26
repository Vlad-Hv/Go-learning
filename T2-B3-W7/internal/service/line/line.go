package line

import (
	"time"
)

type Lines struct {
	Line     string
	Duration time.Duration
}

func Create() []Lines {
	return []Lines{
		{
			Line:     "/slow",
			Duration: 400 * time.Millisecond,
		},

		{
			Line:     "/ok",
			Duration: 40 * time.Millisecond,
		},

		{
			Line:     "/misSing",
			Duration: 70 * time.Millisecond,
		},

		{
			Line:     "/laTer",
			Duration: 150 * time.Millisecond,
		},

		{
			Line:     "/Extra",
			Duration: 30 * time.Millisecond,
		},
	}
}

func CreateSource() []string {
	return []string{
		"/slow",
		"/ok",
		"/extra",
		"/later",
	}
}
