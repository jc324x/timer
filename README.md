timer
=====

[![Go Report Card](https://goreportcard.com/badge/github.com/jychri/timer)](https://goreportcard.com/report/github.com/jychri/timer) [![GoDoc](https://godoc.org/github.com/jychri/timer?status.svg)](https://godoc.org/github.com/jychri/timer)

Package timer 

```go
// Moment marks moments in time.
type Moment struct {
  Name  string
  Time  time.Time
  Start time.Duration // duration since start
  Split time.Duration // duration since last moment
}

// Timer collects Moments.
type Timer struct {
  Moments []Moment
}
```

A `Timer` is initialized with a Start `Moment`.

```go
// Init initializes a *Timer with a Start Moment.
func Init() *Timer {
  ti := new(Timer)
  st := Moment{Name: "Start", Time: time.Now()} // start
  ti.Moments = append(ti.Moments, st)
  return ti
}
```

When a `Moment` is added to the `Timer` with `Mark`, its `Start` and `Split` values are
set relative to the preceding `Moment` and the Start `Moment`.

```go
// Mark marks a moment in time as a Moment and appends t.Moments.
func (ti *Timer) Mark(s string) {
  sm := ti.Moments[0]                            // starting Moment
  lm := ti.Moments[len(ti.Moments)-1]            // last Moment
  m := Moment{Name: s, Time: time.Now()}         // initialize with name and current time
  m.Elapsed = time.Since(sm.Time).Truncate(1000) // total elapsed time as a Duration
  m.Split = m.Elapsed - lm.Elapsed               // time since last Moment as a Duration
  ti.Moments = append(ti.Moments, m)             // append Moment
}
```

`Elapsed` and `Split` return `time.Duration` at the last recorded
moment.

```go
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
```

`Get` provides access to a specific `Moment` in a `Timer`.

``` go
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
```
