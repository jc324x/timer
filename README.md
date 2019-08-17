timer
=====

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

When a `Moment` is added to the `Timer`, its `Start` and `Split` values are
calculated against the previous `Moment` and the Start `Moment`.

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

You can get information about the last `Moment` with `Time` and
`Split` methods. `Get` provides access to a specific `Moment`.

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
