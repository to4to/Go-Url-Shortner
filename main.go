package main

import "time"

type request struct {
	URL string
	CustomShort string
	Expiry time.Duration
}

type response struct {
	URL
	CustomShort
	Expiry
	XRateRemaining
	XRateLimitReset
}
