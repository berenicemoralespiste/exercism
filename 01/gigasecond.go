package gigasecond

import (
	"time"
)

func AddGigaSecond(t time.Time) time.Time {
	timeAfter := t.Add(time.Second * 1000000000)
	return timeAfter
}
