package bpolicy

import (
	"fmt"
	"time"
)

type Clock struct {
	Hour, Minute, Second int
}

func (clock *Clock) String() string {
	return fmt.Sprintf("Clock(Hour=%v, Minute=%v, Second=%v)", clock.Hour, clock.Minute, clock.Second)
}

func (clock *Clock) After(other *Clock) bool {
	check := false
	if clock.Hour == other.Hour {
		if clock.Minute == other.Minute {
			if clock.Second >= other.Second {
				check = true
			}

		} else if clock.Minute > other.Minute {
			check = true
		}
	} else if clock.Hour > other.Hour {
		check = true
	}
	return check
}

func (clock *Clock) InRange(start, end *Clock) bool {
	return clock.After(start) && end.After(clock)
}

func CurrentClock() *Clock {
	Hour, Minute, Second := time.Now().Local().Clock()
	return &Clock{Hour, Minute, Second}
}
