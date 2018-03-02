package gigasecond

import "time"

const Gigasecond = 1e9 * time.Second

func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
