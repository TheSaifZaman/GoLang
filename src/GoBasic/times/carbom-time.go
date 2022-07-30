package carbonTime
 
import (
    "fmt"   
    "github.com/uniplaces/carbon"
)
 
func CarbonTime() {
    fmt.Printf("Right now is %s\n", carbon.Now().DateTimeString())
    today, _ := carbon.NowInLocation("Europe/London")
    fmt.Printf("Right now in London is %s\n", today)
     
    fmt.Printf("\n#######################################\n")   
    fmt.Printf("Tomorrow is %s\n", carbon.Now().AddDay())
    fmt.Printf("Yesterday is %s\n", carbon.Now().SubDay())
    fmt.Printf("After 5 Days %s\n", carbon.Now().AddDays(5))
    fmt.Printf("Before 5 Days %s\n", carbon.Now().SubDays(5))
     
    fmt.Printf("\n#######################################\n")
    fmt.Printf("Next Month is %s\n", carbon.Now().AddMonth())
    fmt.Printf("Last Month is %s\n", carbon.Now().SubMonth())
         
    fmt.Printf("\n#######################################\n")   
    fmt.Printf("Next week is %s\n", carbon.Now().AddWeek())
    fmt.Printf("Last week is %s\n", carbon.Now().SubWeek())
     
    fmt.Printf("\n#######################################\n")   
    fmt.Printf("Next Year %s\n", carbon.Now().AddYear())
    fmt.Printf("Last Year %s\n", carbon.Now().SubYear())
    fmt.Printf("After 5 Years %s\n", carbon.Now().AddYears(5))
    fmt.Printf("Before 5 Years %s\n", carbon.Now().SubYears(5))
     
    fmt.Printf("\n#######################################\n")   
    fmt.Printf("Next Hour %s\n", carbon.Now().AddHour())
    fmt.Printf("Last Hour %s\n", carbon.Now().SubHour())
    fmt.Printf("After 5 Mins %s\n", carbon.Now().AddMinutes(5))
    fmt.Printf("Before 5 Mins %s\n", carbon.Now().SubMinutes(5))
     
    fmt.Printf("\n#######################################\n")   
    fmt.Printf("Weekday? %t\n", carbon.Now().IsWeekday())
    fmt.Printf("Weekend? %t\n", carbon.Now().IsWeekend())
    fmt.Printf("LeapYear? %t\n", carbon.Now().IsLeapYear())
    fmt.Printf("Past? %t\n", carbon.Now().IsPast())
    fmt.Printf("Future? %t\n", carbon.Now().IsFuture())
     
    fmt.Printf("\n#######################################\n")
    fmt.Printf("Start of day:   %s\n", today.StartOfDay())
    fmt.Printf("End of day: %s\n", today.EndOfDay())
    fmt.Printf("Start of month: %s\n", today.StartOfMonth())    
    fmt.Printf("End of month:   %s\n", today.EndOfMonth())
    fmt.Printf("Start of year:  %s\n", today.StartOfYear())
    fmt.Printf("End of year:    %s\n", today.EndOfYear())
    fmt.Printf("Start of week:  %s\n", today.StartOfWeek())
    fmt.Printf("End of week:    %s\n", today.EndOfWeek())
     
}