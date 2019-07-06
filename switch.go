package main

import (
  "fmt"
  "time"
)

func main() {
  i := 2

  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("Two")
  case 3:
    fmt.Println("Three")
  }

  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's weekend!")
  default:
    fmt.Println("It's Workday")
  }

  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("It's before noon")
  default:
    fmt.Println("It's after noon")
  }

  whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
      fmt.Println("I'm a bool")
    case int:
      fmt.Println("I'm a int")
    default:
      fmt.Println("Don't know type", t)
    }
  }

  whatAmI(true)
  whatAmI(1)
  whatAmI("hey")
}
