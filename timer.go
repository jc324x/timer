// Package timer records Moments in time.
package timer

import (
	"errors"
	"time"
)

const (
	// M is a millisecond time.Duration
	M time.Duration = time.Millisecond
)

// Moment marks a moment in time.
type Moment struct {
	Name    string
	Time    time.Time
	Elapsed time.Duration // time since start
	Split   time.Duration // time since last moment
}

// Timer collects Moments.
type Timer struct {
	Moments []Moment
}

// Init initializes a *Timer with a Start Moment.
func Init() *Timer {
	ti := new(Timer)                              // new Timer
	st := Moment{Name: "Start", Time: time.Now()} // initialize starting Moment
	ti.Moments = append(ti.Moments, st)           // append the starting Moment
	return ti                                     // return *Timer
}

// Mark marks a moment in time as a Moment and appends t.Moments.
func (ti *Timer) Mark(s string) {
	sm := ti.Moments[0]                            // starting Moment
	lm := ti.Moments[len(ti.Moments)-1]            // last Moment
	m := Moment{Name: s, Time: time.Now()}         // initialize Moment with name and time
	m.Elapsed = time.Since(sm.Time).Truncate(1000) // total elapsed time as a Duration
	m.Split = m.Elapsed - lm.Elapsed               // time since last Moment as a Duration
	ti.Moments = append(ti.Moments, m)             // append Moment
}

// Elapsed returns the elapsed time at the last recorded moment.
func (ti *Timer) Elapsed() time.Duration {
	lm := ti.Moments[len(ti.Moments)-1] // last moment
	return lm.Elapsed
}

// Split returns the split time for the last recorded moment.
func (ti *Timer) Split() time.Duration {
	lm := ti.Moments[len(ti.Moments)-1] // last moment
	return lm.Split
}

// Get returns a Moment given its name.
func (ti *Timer) Get(s string) (Moment, error) {
	for _, m := range ti.Moments {
		if m.Name == s {
			return m, nil
		}
	}

	var em Moment // empty moment
	return em, errors.New("no moment found")
}
