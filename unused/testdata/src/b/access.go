package iam

import "time"

type Access struct {
	Permissions []string
	ExpiryTime  time.Time
	Revoked     bool
}

func (a Access) Expired(clock Clock) bool {
	now := clock.Now()
	return now.After(a.ExpiryTime)
}
