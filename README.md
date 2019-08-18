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
  sm := ti.Moments[0]                          // starting moment
  lm := ti.Moments[len(ti.Moments)-1]          // last moment
  m := Moment{Name: s, Time: time.Now()}       // name and time
  m.Start = time.Since(sm.Time).Truncate(1000) // duration since start
  m.Split = m.Start - lm.Start                 // duration since last moment
  ti.Moments = append(ti.Moments, m)           // append Moment
}
```

`Time` and `Split` return the elapsed time and last split times 
as `time.Duration`.

```go
// Time returns the elapsed time at the last recorded moment in *Timer.
func (ti *Timer) Time() time.Duration {
  lm := ti.Moments[len(ti.Moments)-1] // last moment
  return lm.Start
}

// Split returns the split time for the last recorded moment in *Timer.
func (ti *Timer) Split() time.Duration {
  lm := ti.Moments[len(ti.Moments)-1] // last moment
  return lm.Split
}
```

`Get` provides access to a specific `Moment` in a `Timer`.

``` go
// Get returns a Moment and an error value from a *Timer.
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
