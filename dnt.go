// Package donottrack is a small helper package for reading the
// DoNotTrack flag out of the http.Request.
// More Information about the DoNotTrack Header can be found at
// https://en.wikipedia.org/wiki/Do_Not_Track
//
// If a user sets the DoNotTrack option inside the browser you
// could not render elements, which would track the user.
// For example you could not show ads or you could not use
// external fonts.
package donottrack

import "net/http"

type Option int

const (
	// OptIn is used, when the user allows tracking
	OptIn Option = iota
	// OptOut is used, when no tracking for the user is allowed
	OptOut
	// NotSet if there is no DNT Element inside the header
	NotSet
)

// Request takes a *http.Request and returns the DNT Option
func Request(r *http.Request) Option {
	dnt, ok := r.Header["Dnt"]
	if ok {
		return mapDNT(dnt[0])
	}
	dnt, ok = r.Header["DNT"]
	if ok {
		return mapDNT(dnt[0])
	}
	return NotSet
}

func mapDNT(i string) Option {
	switch i {
	case "0":
		return OptIn
	case "1":
		return OptOut
	default:
		return NotSet
	}
}
