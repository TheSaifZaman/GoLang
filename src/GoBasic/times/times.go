package times

import (
         "fmt"
         "time"     
)
func Times() {
    t := time.Date(2015, 02, 21, 23, 10, 52, 211, time.UTC) 
    fmt.Println(t)
    fmt.Println("######################################")
     
    y := t.Year()
    mon := t.Month()
    d := t.Day()
    h := t.Hour()
    m := t.Minute()
    s := t.Second()
    n := t.Nanosecond()
     
    fmt.Println("Year   :",y)
    fmt.Println("Month  :",mon)
    fmt.Println("Day    :",d)
    fmt.Println("Hour   :",h)
    fmt.Println("Minute :",m)
    fmt.Println("Second :",s)
    fmt.Println("Nanosec:",n)
}