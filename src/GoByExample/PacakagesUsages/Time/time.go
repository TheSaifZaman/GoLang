package timeUsage

import (
	"fmt"
	"time"
)

/*
Functions in Time Package:

func After(d Duration) <-chan Time
func Sleep(d Duration)
func Tick(d Duration) <-chan Time

type Duration
func ParseDuration(s string) (Duration, error)
func Since(t Time) Duration
func Until(t Time) Duration
func (d Duration) Hours() float64
func (d Duration) Microseconds() int64
func (d Duration) Milliseconds() int64
func (d Duration) Minutes() float64
func (d Duration) Nanoseconds() int64
func (d Duration) Round(m Duration) Duration
func (d Duration) Seconds() float64
func (d Duration) String() string
func (d Duration) Truncate(m Duration) Duration

type Location
func FixedZone(name string, offset int) *Location
func LoadLocation(name string) (*Location, error)
func LoadLocationFromTZData(name string, data []byte) (*Location, error)
func (l *Location) String() string

type Month
func (m Month) String() string

type ParseError
func (e *ParseError) Error() string

type Ticker
func NewTicker(d Duration) *Ticker
func (t *Ticker) Reset(d Duration)
func (t *Ticker) Stop()

type Time
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
func Now() Time
func Parse(layout, value string) (Time, error)
func ParseInLocation(layout, value string, loc *Location) (Time, error)
func Unix(sec int64, nsec int64) Time
func UnixMicro(usec int64) Time
func UnixMilli(msec int64) Time
func (t Time) Add(d Duration) Time
func (t Time) AddDate(years int, months int, days int) Time
func (t Time) After(u Time) bool
func (t Time) AppendFormat(b []byte, layout string) []byte
func (t Time) Before(u Time) bool
func (t Time) Clock() (hour, min, sec int)
func (t Time) Date() (year int, month Month, day int)
func (t Time) Day() int
func (t Time) Equal(u Time) bool
func (t Time) Format(layout string) string
func (t Time) GoString() string
func (t *Time) GobDecode(data []byte) error
func (t Time) GobEncode() ([]byte, error)
func (t Time) Hour() int
func (t Time) ISOWeek() (year, week int)
func (t Time) In(loc *Location) Time
func (t Time) IsDST() bool
func (t Time) IsZero() bool
func (t Time) Local() Time
func (t Time) Location() *Location
func (t Time) MarshalBinary() ([]byte, error)
func (t Time) MarshalJSON() ([]byte, error)
func (t Time) MarshalText() ([]byte, error)
func (t Time) Minute() int
func (t Time) Month() Month
func (t Time) Nanosecond() int
func (t Time) Round(d Duration) Time
func (t Time) Second() int
func (t Time) String() string
func (t Time) Sub(u Time) Duration
func (t Time) Truncate(d Duration) Time
func (t Time) UTC() Time
func (t Time) Unix() int64
func (t Time) UnixMicro() int64
func (t Time) UnixMilli() int64
func (t Time) UnixNano() int64
func (t *Time) UnmarshalBinary(data []byte) error
func (t *Time) UnmarshalJSON(data []byte) error
func (t *Time) UnmarshalText(data []byte) error
func (t Time) Weekday() Weekday
func (t Time) Year() int
func (t Time) YearDay() int
func (t Time) Zone() (name string, offset int)

type Timer
func AfterFunc(d Duration, f func()) *Timer
func NewTimer(d Duration) *Timer
func (t *Timer) Reset(d Duration) bool
func (t *Timer) Stop() bool

type Weekday
func (d Weekday) String() string
*/

/*
Constants:

const (
	Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)
These are predefined layouts for use in Time.Format and time.Parse. The reference time used in these layouts is the specific time stamp:

01/02 03:04:05PM '06 -0700
(January 2, 15:04:05, 2006, in time zone seven hours west of GMT). That value is recorded as the constant named Layout, listed below. As a Unix time, this is 1136239445. Since MST is GMT-0700, the reference would be printed by the Unix date command as:

Mon Jan 2 15:04:05 MST 2006
It is a regrettable historic error that the date uses the American convention of putting the numerical month before the day.

The example for Time.Format demonstrates the working of the layout string in detail and is a good reference.

Note that the RFC822, RFC850, and RFC1123 formats should be applied only to local times. Applying them to UTC times will use "UTC" as the time zone abbreviation, while strictly speaking those RFCs require the use of "GMT" in that case. In general RFC1123Z should be used instead of RFC1123 for servers that insist on that format, and RFC3339 should be preferred for new protocols. RFC3339, RFC822, RFC822Z, RFC1123, and RFC1123Z are useful for formatting; when used with time.Parse they do not accept all the time formats permitted by the RFCs and they do accept time formats not formally defined. The RFC3339Nano format removes trailing zeros from the seconds field and thus may not sort correctly once formatted.

Most programs can use one of the defined constants as the layout passed to Format or Parse. The rest of this comment can be ignored unless you are creating a custom layout string.

To define your own format, write down what the reference time would look like formatted your way; see the values of constants like ANSIC, StampMicro or Kitchen for examples. The model is to demonstrate what the reference time looks like so that the Format and Parse methods can apply the same transformation to a general time value.

Here is a summary of the components of a layout string. Each element shows by example the formatting of an element of the reference time. Only these values are recognized. Text in the layout string that is not recognized as part of the reference time is echoed verbatim during Format and expected to appear verbatim in the input to Parse.

Year: "2006" "06"
Month: "Jan" "January"
Textual day of the week: "Mon" "Monday"
Numeric day of the month: "2" "_2" "02"
Numeric day of the year: "__2" "002"
Hour: "15" "3" "03" (PM or AM)
Minute: "4" "04"
Second: "5" "05"
AM/PM mark: "PM"
Numeric time zone offsets format as follows:

"-0700"  ±hhmm
"-07:00" ±hh:mm
"-07"    ±hh
Replacing the sign in the format with a Z triggers the ISO 8601 behavior of printing Z instead of an offset for the UTC zone. Thus:

"Z0700"  Z or ±hhmm
"Z07:00" Z or ±hh:mm
"Z07"    Z or ±hh
Within the format string, the underscores in "_2" and "__2" represent spaces that may be replaced by digits if the following number has multiple digits, for compatibility with fixed-width Unix time formats. A leading zero represents a zero-padded value.

The formats __2 and 002 are space-padded and zero-padded three-character day of year; there is no unpadded day of year format.

A comma or decimal point followed by one or more zeros represents a fractional second, printed to the given number of decimal places. A comma or decimal point followed by one or more nines represents a fractional second, printed to the given number of decimal places, with trailing zeros removed. For example "15:04:05,000" or "15:04:05.000" formats or parses with millisecond precision.

Some valid layouts are invalid time values for time.Parse, due to formats such as _ for space padding and Z for zone information.

const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)
Common durations. There is no definition for units of Day or larger to avoid confusion across daylight savings time zone transitions.

To count the number of units in a Duration, divide:

second := time.Second
fmt.Print(int64(second/time.Millisecond)) // prints 1000
To convert an integer number of units to a Duration, multiply:

seconds := 10
fmt.Print(time.Duration(seconds)*time.Second) // prints 10s
*/

/*
func After(d Duration) <-chan Time
After waits for the duration to elapse and then sends the current time on the returned channel. It is equivalent to NewTimer(d).C. The underlying Timer is not recovered by the garbage collector until the timer fires. If efficiency is a concern, use NewTimer instead and call Timer.Stop if the timer is no longer needed.
*/
var c chan int

func handle(int) {}

func AfterTest() {
	select {
	case m := <-c:
		handle(m)
	case <-time.After(10 * time.Second):
		fmt.Println("timed out")
	}
}

/*
func Sleep(d Duration)
Sleep pauses the current goroutine for at least the duration d. A negative or zero duration causes Sleep to return immediately.
*/
func SleepTest() {
	time.Sleep(100 * time.Millisecond)
}

/*
func Tick(d Duration) <-chan Time
Tick is a convenience wrapper for NewTicker providing access to the ticking channel only. While Tick is useful for clients that have no need to shut down the Ticker, be aware that without a way to shut it down the underlying Ticker cannot be recovered by the garbage collector; it "leaks". Unlike NewTicker, Tick will return nil if d <= 0.
*/
func statusUpdate() string { return "" }

func TickUpdate() {
	c := time.Tick(5 * time.Second)
	for next := range c {
		fmt.Printf("%v %s\n", next, statusUpdate())
	}
}

/*
type Duration int64
A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years.
*/
func expensiveCall() {}

func DurationTest() {
	t0 := time.Now()
	expensiveCall()
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}

/*
func ParseDuration(s string) (Duration, error)
ParseDuration parses a duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
*/
func ParseDurationTest() {
	hours, _ := time.ParseDuration("10h")
	complex, _ := time.ParseDuration("1h10m10s")
	micro, _ := time.ParseDuration("1µs")
	// The package also accepts the incorrect but common prefix u for micro.
	micro2, _ := time.ParseDuration("1us")

	fmt.Println(hours)
	fmt.Println(complex)
	fmt.Printf("There are %.0f seconds in %v.\n", complex.Seconds(), complex)
	fmt.Printf("There are %d nanoseconds in %v.\n", micro.Nanoseconds(), micro)
	fmt.Printf("There are %6.2e seconds in %v.\n", micro2.Seconds(), micro)
}

/*
func Since(t Time) Duration
Since returns the time elapsed since t. It is shorthand for time.Now().Sub(t).

added in go1.8
func Until(t Time) Duration
Until returns the duration until t. It is shorthand for t.Sub(time.Now()).
*/

/*
func (d Duration) Hours() float64
Hours returns the duration as a floating point number of hours.
*/
func HoursTest() {
	h, _ := time.ParseDuration("4h30m")
	fmt.Printf("I've got %.1f hours of work left.", h.Hours())
}

// func (d Duration) Microseconds() int64
// Microseconds returns the duration as an integer microsecond count.
func MicrosecondsTest() {
	u, _ := time.ParseDuration("1s")
	fmt.Printf("One second is %d microseconds.\n", u.Microseconds())
}

// func (d Duration) Milliseconds() int64
// Milliseconds returns the duration as an integer millisecond count.
func MillisecondsTest() {
	u, _ := time.ParseDuration("1s")
	fmt.Printf("One second is %d milliseconds.\n", u.Milliseconds())
}

// func (d Duration) Minutes() float64
// Minutes returns the duration as a floating point number of minutes.
func MinutesTest() {
	m, _ := time.ParseDuration("1h30m")
	fmt.Printf("The movie is %.0f minutes long.", m.Minutes())
}

// func (d Duration) Nanoseconds() int64
// Nanoseconds returns the duration as an integer nanosecond count.
func NanosecondsTest() {
	u, _ := time.ParseDuration("1µs")
	fmt.Printf("One microsecond is %d nanoseconds.\n", u.Nanoseconds())
}

/*
func (d Duration) Round(m Duration) Duration
Round returns the result of rounding d to the nearest multiple of m. The rounding behavior for halfway values is to round away from zero. If the result exceeds the maximum (or minimum) value that can be stored in a Duration, Round returns the maximum (or minimum) duration. If m <= 0, Round returns d unchanged.
*/
func RoundTest() {
	d, err := time.ParseDuration("5h15m30.918273645s")
	if err != nil {
		panic(err)
	}

	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		20 * time.Second,
		time.Minute,
		5 * time.Minute,
		10 * time.Minute,
		time.Hour,
		5 * time.Hour,
		10 * time.Hour,
	}

	for _, r := range round {
		fmt.Printf("d.Round(%9s) = %s\n", r, d.Round(r).String())
	}
}
/*
Output:
d.Round(      1ns) = 5h15m30.918273645s
d.Round(      1µs) = 5h15m30.918274s
d.Round(      1ms) = 5h15m30.918s
d.Round(       1s) = 5h15m31s
d.Round(       2s) = 5h15m30s
d.Round(      20s) = 5h15m40s
d.Round(     1m0s) = 5h16m0s
d.Round(     5m0s) = 5h15m0s
d.Round(    10m0s) = 5h20m0s
d.Round(   1h0m0s) = 5h0m0s
d.Round(   5h0m0s) = 5h0m0s
d.Round(  10h0m0s) = 10h0m0s
*/

// func (d Duration) Seconds() float64
// Seconds returns the duration as a floating point number of seconds.
func SecondsTest() {
	m, _ := time.ParseDuration("1m30s")
	fmt.Printf("Take off in t-%.0f seconds.", m.Seconds())
}

/*
func (d Duration) String() string
String returns a string representing the duration in the form "72h3m0.5s". Leading zero units are omitted. As a special case, durations less than one second format use a smaller unit (milli-, micro-, or nanoseconds) to ensure that the leading digit is non-zero. The zero duration formats as 0s.
*/
func StringTest() {
	fmt.Println(1*time.Hour + 2*time.Minute + 300*time.Millisecond)
	fmt.Println(300 * time.Millisecond)
}

// func (d Duration) Truncate(m Duration) Duration
// Truncate returns the result of rounding d toward zero to a multiple of m. If m <= 0, Truncate returns d unchanged.
func TruncateTest() {
	d, err := time.ParseDuration("5h15m30.918273645s")
	if err != nil {
		panic(err)
	}

	trunc := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		20 * time.Second,
		time.Minute,
		5 * time.Minute,
		10 * time.Minute,
		time.Hour,
		5 * time.Hour,
		10 * time.Hour,
	}

	for _, t := range trunc {
		fmt.Printf("d.Truncate(%9s) = %s\n", t, d.Truncate(t).String())
	}
}
/*
OutPut:
d.Truncate(      1ns) = 5h15m30.918273645s
d.Truncate(      1µs) = 5h15m30.918273s
d.Truncate(      1ms) = 5h15m30.918s
d.Truncate(       1s) = 5h15m30s
d.Truncate(       2s) = 5h15m30s
d.Truncate(      20s) = 5h15m20s
d.Truncate(     1m0s) = 5h15m0s
d.Truncate(     5m0s) = 5h15m0s
d.Truncate(    10m0s) = 5h10m0s
d.Truncate(   1h0m0s) = 5h0m0s
d.Truncate(   5h0m0s) = 5h0m0s
d.Truncate(  10h0m0s) = 0s
*/

/*
Type Location
type Location struct {
	// contains filtered or unexported fields
}
A Location maps time instants to the zone in use at that time. Typically, the Location represents the collection of time offsets in use in a geographical area. For many Locations the time offset varies depending on whether daylight savings time is in use at the time instant.
*/
func LocationTypeTest() {
	// China doesn't have daylight saving. It uses a fixed 8 hour offset from UTC.
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// If the system has a timezone database present, it's possible to load a location
	// from that, e.g.:
	//    newYork, err := time.LoadLocation("America/New_York")

	// Creating a time requires a location. Common locations are time.Local and time.UTC.
	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)

	// Although the UTC clock time is 1200 and the Beijing clock time is 2000, Beijing is
	// 8 hours ahead so the two dates actually represent the same instant.
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	fmt.Println(timesAreEqual)

}

/*
var Local *Location = &localLoc
Local represents the system's local time zone. On Unix systems, Local consults the TZ environment variable to find the time zone to use. No TZ means use the system default /etc/localtime. TZ="" means use UTC. TZ="foo" means use file foo in the system timezone directory.

var UTC *Location = &utcLoc
UTC represents Universal Coordinated Time (UTC).
*/

/*
func FixedZone(name string, offset int) *Location
FixedZone returns a Location that always uses the given zone name and offset (seconds east of UTC).
*/
func FixedZoneTest() {
	loc := time.FixedZone("UTC+6", +6*60*60)
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, loc)
	fmt.Println("The time is:", t.Format(time.RFC822))
}
//The time is: 10 Nov 09 23:00 UTC+6

/*
func LoadLocation(name string) (*Location, error)
LoadLocation returns the Location with the given name.

If the name is "" or "UTC", LoadLocation returns UTC. If the name is "Local", LoadLocation returns Local.

Otherwise, the name is taken to be a location name corresponding to a file in the IANA Time Zone database, such as "America/New_York".

LoadLocation looks for the IANA Time Zone database in the following locations in order:

- the directory or uncompressed zip file named by the ZONEINFO environment variable - on a Unix system, the system standard installation location - $GOROOT/lib/time/zoneinfo.zip - the time/tzdata package, if it was imported
*/
func LoadLocationTest() {
	location, err := time.LoadLocation("Asia/Dhaka")
	if err != nil {
		panic(err)
	}

	timeInUTC := time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println(timeInUTC.In(location))
}
//2018-08-30 18:00:00 +0600 +06

/*
func LoadLocationFromTZData(name string, data []byte) (*Location, error)
LoadLocationFromTZData returns a Location with the given name initialized from the IANA Time Zone database-formatted data. The data should be in the format of a standard IANA time zone file (for example, the content of /etc/localtime on Unix systems).

func (l *Location) String() string
String returns a descriptive name for the time zone information, corresponding to the name argument to LoadLocation or FixedZone.
*/

/*
type Month int
A Month specifies a month of the year (January = 1, ...).
*/
func MonthTypeTest() {
	_, month, day := time.Now().Date()
	if month == time.November && day == 10 {
		fmt.Println("Happy Go day!")
	}
}

/*
const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)
func (m Month) String() string
String returns the English name of the month ("January", "February", ...).

type ParseError struct {
	Layout     string
	Value      string
	LayoutElem string
	ValueElem  string
	Message    string
}
ParseError describes a problem parsing a time string.

func (e *ParseError) Error() string
Error returns the string representation of a ParseError.

type Ticker struct {
	C <-chan Time // The channel on which the ticks are delivered.
	// contains filtered or unexported fields
}
A Ticker holds a channel that delivers “ticks” of a clock at intervals.
*/

/*
func NewTicker(d Duration) *Ticker
NewTicker returns a new Ticker containing a channel that will send the current time on the channel after each tick. The period of the ticks is specified by the duration argument. The ticker will adjust the time interval or drop ticks to make up for slow receivers. The duration d must be greater than zero; if not, NewTicker will panic. Stop the ticker to release associated resources.
*/
func NewTickerYest() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current Time: ", t)
		}
	}
}
//Output:
// Current time:  2009-11-10 23:00:10 +0000 UTC m=+10.000000001
// Done!

/*
func (t *Ticker) Reset(d Duration)
Reset stops a ticker and resets its period to the specified duration. The next tick will arrive after the new period elapses. The duration d must be greater than zero; if not, Reset will panic.

func (t *Ticker) Stop()
Stop turns off a ticker. After Stop, no more ticks will be sent. Stop does not close the channel, to prevent a concurrent goroutine reading from the channel from seeing an erroneous "tick".

type Time struct {
	// contains filtered or unexported fields
}
A Time represents an instant in time with nanosecond precision.

Programs using times should typically store and pass them as values, not pointers. That is, time variables and struct fields should be of type time.Time, not *time.Time.

A Time value can be used by multiple goroutines simultaneously except that the methods GobDecode, UnmarshalBinary, UnmarshalJSON and UnmarshalText are not concurrency-safe.

Time instants can be compared using the Before, After, and Equal methods. The Sub method subtracts two instants, producing a Duration. The Add method adds a Time and a Duration, producing a Time.

The zero value of type Time is January 1, year 1, 00:00:00.000000000 UTC. As this time is unlikely to come up in practice, the IsZero method gives a simple way of detecting a time that has not been initialized explicitly.

Each Time has associated with it a Location, consulted when computing the presentation form of the time, such as in the Format, Hour, and Year methods. The methods Local, UTC, and In return a Time with a specific location. Changing the location in this way changes only the presentation; it does not change the instant in time being denoted and therefore does not affect the computations described in earlier paragraphs.

Representations of a Time value saved by the GobEncode, MarshalBinary, MarshalJSON, and MarshalText methods store the Time.Location's offset, but not the location name. They therefore lose information about Daylight Saving Time.

In addition to the required “wall clock” reading, a Time may contain an optional reading of the current process's monotonic clock, to provide additional precision for comparison or subtraction. See the “Monotonic Clocks” section in the package documentation for details.

Note that the Go == operator compares not just the time instant but also the Location and the monotonic clock reading. Therefore, Time values should not be used as map or database keys without first guaranteeing that the identical Location has been set for all values, which can be achieved through use of the UTC or Local method, and that the monotonic clock reading has been stripped by setting t = t.Round(0). In general, prefer t.Equal(u) to t == u, since t.Equal uses the most accurate comparison available and correctly handles the case when only one of its arguments has a monotonic clock reading.
*/

/*
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
Date returns the Time corresponding to

yyyy-mm-dd hh:mm:ss + nsec nanoseconds
in the appropriate zone for that time in the given location.

The month, day, hour, min, sec, and nsec values may be outside their usual ranges and will be normalized during the conversion. For example, October 32 converts to November 1.

A daylight savings time transition skips or repeats times. For example, in the United States, March 13, 2011 2:15am never occurred, while November 6, 2011 1:15am occurred twice. In such cases, the choice of time zone, and therefore the time, is not well-defined. Date returns a time that is correct in one of the two zones involved in the transition, but it does not guarantee which.

Date panics if loc is nil.
*/
func DateTest() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
}
//Go launched at 2009-11-10 15:00:00 -0800 PST

// func Now() Time
// Now returns the current local time.

/*
func Parse(layout, value string) (Time, error)
Parse parses a formatted string and returns the time value it represents. See the documentation for the constant called Layout to see how to represent the format. The second argument must be parseable using the format string (layout) provided as the first argument.
*/
func ParseTest() {
	// See the example for Time.Format for a thorough description of how
	// to define the layout string to parse a time.Time value; Parse and
	// Format use the same model to describe their input and output.

	// longForm shows by example how the reference time would be represented in
	// the desired layout.
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t)

	// shortForm is another way the reference time would be represented
	// in the desired layout; it has no time zone present.
	// Note: without explicit zone, returns time in UTC.
	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)

	// Some valid layouts are invalid time values, due to format specifiers
	// such as _ for space padding and Z for zone information.
	// For example the RFC3339 layout 2006-01-02T15:04:05Z07:00
	// contains both Z and a time zone offset in order to handle both valid options:
	// 2006-01-02T15:04:05Z
	// 2006-01-02T15:04:05+07:00
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println(t)
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	fmt.Println(t)
	_, err := time.Parse(time.RFC3339, time.RFC3339)
	fmt.Println("error", err) // Returns an error as the layout is not a valid time value
}
/*
Output:

2013-02-03 19:54:00 -0800 PST
2013-02-03 00:00:00 +0000 UTC
2006-01-02 15:04:05 +0000 UTC
2006-01-02 15:04:05 +0700 +0700
error parsing time "2006-01-02T15:04:05Z07:00": extra text: "07:00"
*/

/*
func ParseInLocation(layout, value string, loc *Location) (Time, error)
ParseInLocation is like Parse but differs in two important ways. First, in the absence of time zone information, Parse interprets a time as UTC; ParseInLocation interprets the time as in the given location. Second, when given a zone offset or abbreviation, Parse tries to match it against the Local location; ParseInLocation uses the given location.
*/
func ParseInLocationTest() {
	loc, _ := time.LoadLocation("Asia/Dhaka")

	// This will look for the name UTC+6 in the Asia/Dhaka time zone.
	const longForm = "Jan 2, 2006 at 3:04pm (UTC+6)"
	t, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (UTC+06)", loc)
	fmt.Println(t)

	// Note: without explicit zone, returns time in given location.
	const shortForm = "2006-Jan-02"
	t, _ = time.ParseInLocation(shortForm, "2012-Jul-09", loc)
	fmt.Println(t)

}
// 0001-01-01 00:00:00 +0000 UTC
// 2012-07-09 00:00:00 +0600 +06

/*
func Unix(sec int64, nsec int64) Time
Unix returns the local Time corresponding to the given Unix time, sec seconds and nsec nanoseconds since January 1, 1970 UTC. It is valid to pass nsec outside the range [0, 999999999]. Not all sec values have a corresponding time value. One such value is 1<<63-1 (the largest int64 value).
*/
func UnixTest() {
	unixTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(unixTime.Unix())
	t := time.Unix(unixTime.Unix(), 0).UTC()
	fmt.Println(t)
}
/*
Output:

1257894000
2009-11-10 23:00:00 +0000 UTC
*/

// func UnixMilli(msec int64) Time
// UnixMilli returns the local Time corresponding to the given Unix time, msec milliseconds since January 1, 1970 UTC.
func UnixMilliTest() {
	umt := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(umt.UnixMilli())
	t := time.UnixMilli(umt.UnixMilli()).UTC()
	fmt.Println(t)

}
/*
Output:
1257894000000
2009-11-10 23:00:00 +0000 UTC
*/

// func UnixMicro(usec int64) Time
// UnixMicro returns the local Time corresponding to the given Unix time, usec microseconds since January 1, 1970 UTC.
func UnixMicroTest() {
	umt := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(umt.UnixMicro())
	t := time.UnixMicro(umt.UnixMicro()).UTC()
	fmt.Println(t)

}
/*
Output:
1257894000000000
2009-11-10 23:00:00 +0000 UTC
*/

// func (t Time) Add(d Duration) Time
// Add returns the time t+d.
func AddTimeTest() {
	start := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	afterTenSeconds := start.Add(time.Second * 10)
	afterTenMinutes := start.Add(time.Minute * 10)
	afterTenHours := start.Add(time.Hour * 10)
	afterTenDays := start.Add(time.Hour * 24 * 10)

	fmt.Printf("start = %v\n", start)
	fmt.Printf("start.Add(time.Second * 10) = %v\n", afterTenSeconds)
	fmt.Printf("start.Add(time.Minute * 10) = %v\n", afterTenMinutes)
	fmt.Printf("start.Add(time.Hour * 10) = %v\n", afterTenHours)
	fmt.Printf("start.Add(time.Hour * 24 * 10) = %v\n", afterTenDays)
}
/*
Output:
start = 2009-01-01 12:00:00 +0000 UTC
start.Add(time.Second * 10) = 2009-01-01 12:00:10 +0000 UTC
start.Add(time.Minute * 10) = 2009-01-01 12:10:00 +0000 UTC
start.Add(time.Hour * 10) = 2009-01-01 22:00:00 +0000 UTC
start.Add(time.Hour * 24 * 10) = 2009-01-11 12:00:00 +0000 UTC
*/

/*
func (t Time) AddDate(years int, months int, days int) Time
AddDate returns the time corresponding to adding the given number of years, months, and days to t. For example, AddDate(-1, 2, 3) applied to January 1, 2011 returns March 4, 2010.

AddDate normalizes its result in the same way that Date does, so, for example, adding one month to October 31 yields December 1, the normalized form for November 31.
*/
func AddDate() {
	start := time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC)
	oneDayLater := start.AddDate(0, 0, 1)
	oneMonthLater := start.AddDate(0, 1, 0)
	oneYearLater := start.AddDate(1, 0, 0)

	fmt.Printf("oneDayLater: start.AddDate(0, 0, 1) = %v\n", oneDayLater)
	fmt.Printf("oneMonthLater: start.AddDate(0, 1, 0) = %v\n", oneMonthLater)
	fmt.Printf("oneYearLater: start.AddDate(1, 0, 0) = %v\n", oneYearLater)

}
/*
Output:

oneDayLater: start.AddDate(0, 0, 1) = 2009-01-02 00:00:00 +0000 UTC
oneMonthLater: start.AddDate(0, 1, 0) = 2009-02-01 00:00:00 +0000 UTC
oneYearLater: start.AddDate(1, 0, 0) = 2010-01-01 00:00:00 +0000 UTC
*/

/*
func (t Time) After(u Time) bool
After reports whether the time instant t is after u.
*/
func DateAgainTest() {
	year2000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	year3000 := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	isYear3000AfterYear2000 := year3000.After(year2000) // True
	isYear2000AfterYear3000 := year2000.After(year3000) // False

	fmt.Printf("year3000.After(year2000) = %v\n", isYear3000AfterYear2000)
	fmt.Printf("year2000.After(year3000) = %v\n", isYear2000AfterYear3000)

}
/*
Output:
year3000.After(year2000) = true
year2000.After(year3000) = false
*/

/*
func (t Time) AppendFormat(b []byte, layout string) []byte
AppendFormat is like Format but appends the textual representation to b and returns the extended buffer.
*/
func AppendFormatTest() {
	t := time.Date(2017, time.November, 4, 11, 0, 0, 0, time.UTC)
	text := []byte("Time: ")

	text = t.AppendFormat(text, time.Kitchen)
	fmt.Println(string(text))

}
/*
Output:
Time: 11:00AM
*/

/*
func (t Time) Before(u Time) bool
Before reports whether the time instant t is before u.
*/
func BeforeTest() {
	year2000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	year3000 := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	isYear2000BeforeYear3000 := year2000.Before(year3000) // True
	isYear3000BeforeYear2000 := year3000.Before(year2000) // False

	fmt.Printf("year2000.Before(year3000) = %v\n", isYear2000BeforeYear3000)
	fmt.Printf("year3000.Before(year2000) = %v\n", isYear3000BeforeYear2000)
}
/*
year2000.Before(year3000) = true
year3000.Before(year2000) = false
*/

/*
func (t Time) Clock() (hour, min, sec int)
Clock returns the hour, minute, and second within the day specified by t.
*/

/*
func (t Time) Date() (year int, month Month, day int)
Date returns the year, month, and day in which t occurs.
*/
func TimeDateTest() {
	d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	year, month, day := d.Date()

	fmt.Printf("year = %v\n", year)
	fmt.Printf("month = %v\n", month)
	fmt.Printf("day = %v\n", day)
}
/*
Output:

year = 2000
month = February
day = 1
*/

// func (t Time) Day() int
// Day returns the day of the month specified by t.
func TimeDayTest() {
	d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	day := d.Day()

	fmt.Printf("day = %v\n", day)
} //day = 1

/*
func (t Time) Equal(u Time) bool
Equal reports whether t and u represent the same time instant. Two times can be equal even if they are in different locations. For example, 6:00 +0200 and 4:00 UTC are Equal. See the documentation on the Time type for the pitfalls of using == with Time values; most code should use Equal instead.
*/
func main() {
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// Unlike the equal operator, Equal is aware that d1 and d2 are the
	// same instant but in different time zones.
	d1 := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	d2 := time.Date(2000, 2, 1, 20, 30, 0, 0, beijing)

	datesEqualUsingEqualOperator := d1 == d2
	datesEqualUsingFunction := d1.Equal(d2)

	fmt.Printf("datesEqualUsingEqualOperator = %v\n", datesEqualUsingEqualOperator)
	fmt.Printf("datesEqualUsingFunction = %v\n", datesEqualUsingFunction)
}
/*
Output:

datesEqualUsingEqualOperator = false
datesEqualUsingFunction = true
*/


/*
func (t Time) Format(layout string) string
Format returns a textual representation of the time value formatted according to the layout defined by the argument. See the documentation for the constant called Layout to see how to represent the layout format.

The executable example for Time.Format demonstrates the working of the layout string in detail and is a good reference.
*/
func TimeFormatTest() {
	// Parse a time value from a string in the standard Unix format.
	t, err := time.Parse(time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")
	if err != nil { // Always check errors even if they should not happen.
		panic(err)
	}

	// time.Time's Stringer method is useful without any format.
	fmt.Println("default format:", t)

	// Predefined constants in the package implement common layouts.
	fmt.Println("Unix format:", t.Format(time.UnixDate))

	// The time zone attached to the time value affects its output.
	fmt.Println("Same, in UTC:", t.UTC().Format(time.UnixDate))

	// The rest of this function demonstrates the properties of the
	// layout string used in the format.

	// The layout string used by the Parse function and Format method
	// shows by example how the reference time should be represented.
	// We stress that one must show how the reference time is formatted,
	// not a time of the user's choosing. Thus each layout string is a
	// representation of the time stamp,
	//	Jan 2 15:04:05 2006 MST
	// An easy way to remember this value is that it holds, when presented
	// in this order, the values (lined up with the elements above):
	//	  1 2  3  4  5    6  -7
	// There are some wrinkles illustrated below.

	// Most uses of Format and Parse use constant layout strings such as
	// the ones defined in this package, but the interface is flexible,
	// as these examples show.

	// Define a helper function to make the examples' output look nice.
	do := func(name, layout, want string) {
		got := t.Format(layout)
		if want != got {
			fmt.Printf("error: for %q got %q; expected %q\n", layout, got, want)
			return
		}
		fmt.Printf("%-16s %q gives %q\n", name, layout, got)
	}

	// Print a header in our output.
	fmt.Printf("\nFormats:\n\n")

	// Simple starter examples.
	do("Basic full date", "Mon Jan 2 15:04:05 MST 2006", "Wed Feb 25 11:06:39 PST 2015")
	do("Basic short date", "2006/01/02", "2015/02/25")

	// The hour of the reference time is 15, or 3PM. The layout can express
	// it either way, and since our value is the morning we should see it as
	// an AM time. We show both in one format string. Lower case too.
	do("AM/PM", "3PM==3pm==15h", "11AM==11am==11h")

	// When parsing, if the seconds value is followed by a decimal point
	// and some digits, that is taken as a fraction of a second even if
	// the layout string does not represent the fractional second.
	// Here we add a fractional second to our time value used above.
	t, err = time.Parse(time.UnixDate, "Wed Feb 25 11:06:39.1234 PST 2015")
	if err != nil {
		panic(err)
	}
	// It does not appear in the output if the layout string does not contain
	// a representation of the fractional second.
	do("No fraction", time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")

	// Fractional seconds can be printed by adding a run of 0s or 9s after
	// a decimal point in the seconds value in the layout string.
	// If the layout digits are 0s, the fractional second is of the specified
	// width. Note that the output has a trailing zero.
	do("0s for fraction", "15:04:05.00000", "11:06:39.12340")

	// If the fraction in the layout is 9s, trailing zeros are dropped.
	do("9s for fraction", "15:04:05.99999999", "11:06:39.1234")

}
/*
Output:

default format: 2015-02-25 11:06:39 -0800 PST
Unix format: Wed Feb 25 11:06:39 PST 2015
Same, in UTC: Wed Feb 25 19:06:39 UTC 2015

Formats:

Basic full date  "Mon Jan 2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
Basic short date "2006/01/02" gives "2015/02/25"
AM/PM            "3PM==3pm==15h" gives "11AM==11am==11h"
No fraction      "Mon Jan _2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
0s for fraction  "15:04:05.00000" gives "11:06:39.12340"
9s for fraction  "15:04:05.99999999" gives "11:06:39.1234"
*/

func TimeForamtWithPadTest() {
	// Parse a time value from a string in the standard Unix format.
	t, err := time.Parse(time.UnixDate, "Sat Mar 7 11:06:39 PST 2015")
	if err != nil { // Always check errors even if they should not happen.
		panic(err)
	}

	// Define a helper function to make the examples' output look nice.
	do := func(name, layout, want string) {
		got := t.Format(layout)
		if want != got {
			fmt.Printf("error: for %q got %q; expected %q\n", layout, got, want)
			return
		}
		fmt.Printf("%-16s %q gives %q\n", name, layout, got)
	}

	// The predefined constant Unix uses an underscore to pad the day.
	do("Unix", time.UnixDate, "Sat Mar  7 11:06:39 PST 2015")

	// For fixed-width printing of values, such as the date, that may be one or
	// two characters (7 vs. 07), use an _ instead of a space in the layout string.
	// Here we print just the day, which is 2 in our layout string and 7 in our
	// value.
	do("No pad", "<2>", "<7>")

	// An underscore represents a space pad, if the date only has one digit.
	do("Spaces", "<_2>", "< 7>")

	// A "0" indicates zero padding for single-digit values.
	do("Zeros", "<02>", "<07>")

	// If the value is already the right width, padding is not used.
	// For instance, the second (05 in the reference time) in our value is 39,
	// so it doesn't need padding, but the minutes (04, 06) does.
	do("Suppressed pad", "04:05", "06:39")

}
/*
Output:

Unix             "Mon Jan _2 15:04:05 MST 2006" gives "Sat Mar  7 11:06:39 PST 2015"
No pad           "<2>" gives "<7>"
Spaces           "<_2>" gives "< 7>"
Zeros            "<02>" gives "<07>"
Suppressed pad   "04:05" gives "06:39"
*/

/*
func (t Time) GoString() string
GoString implements fmt.GoStringer and formats t to be printed in Go source code.
*/
func GoStringTest() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(t.GoString())
	t = t.Add(1 * time.Minute)
	fmt.Println(t.GoString())
	t = t.AddDate(0, 1, 0)
	fmt.Println(t.GoString())
	t, _ = time.Parse("Jan 2, 2006 at 3:04pm (MST)", "Feb 3, 2013 at 7:54pm (UTC)")
	fmt.Println(t.GoString())

}
/*
Output:

time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
time.Date(2009, time.November, 10, 23, 1, 0, 0, time.UTC)
time.Date(2009, time.December, 10, 23, 1, 0, 0, time.UTC)
time.Date(2013, time.February, 3, 19, 54, 0, 0, time.UTC)
*/

/*
func (t *Time) GobDecode(data []byte) error
GobDecode implements the gob.GobDecoder interface.

func (t Time) GobEncode() ([]byte, error)
GobEncode implements the gob.GobEncoder interface.

func (t Time) Hour() int
Hour returns the hour within the day specified by t, in the range [0, 23].

func (t Time) ISOWeek() (year, week int)
ISOWeek returns the ISO 8601 year and week number in which t occurs. Week ranges from 1 to 53. Jan 01 to Jan 03 of year n might belong to week 52 or 53 of year n-1, and Dec 29 to Dec 31 might belong to week 1 of year n+1.

func (t Time) In(loc *Location) Time
In returns a copy of t representing the same time instant, but with the copy's location information set to loc for display purposes.

In panics if loc is nil.

added in go1.17
func (t Time) IsDST() bool
IsDST reports whether the time in the configured location is in Daylight Savings Time.

func (t Time) IsZero() bool
IsZero reports whether t represents the zero time instant, January 1, year 1, 00:00:00 UTC.

func (t Time) Local() Time
Local returns t with the location set to local time.

func (t Time) Location() *Location
Location returns the time zone information associated with t.

added in go1.2
func (t Time) MarshalBinary() ([]byte, error)
MarshalBinary implements the encoding.BinaryMarshaler interface.

func (t Time) MarshalJSON() ([]byte, error)
MarshalJSON implements the json.Marshaler interface. The time is a quoted string in RFC 3339 format, with sub-second precision added if present.

added in go1.2
func (t Time) MarshalText() ([]byte, error)
MarshalText implements the encoding.TextMarshaler interface. The time is formatted in RFC 3339 format, with sub-second precision added if present.

func (t Time) Minute() int
Minute returns the minute offset within the hour specified by t, in the range [0, 59].

func (t Time) Month() Month
Month returns the month of the year specified by t.

func (t Time) Nanosecond() int
Nanosecond returns the nanosecond offset within the second specified by t, in the range [0, 999999999].

added in go1.1
func (t Time) Round(d Duration) Time
Round returns the result of rounding t to the nearest multiple of d (since the zero time). The rounding behavior for halfway values is to round up. If d <= 0, Round returns t stripped of any monotonic clock reading but otherwise unchanged.

Round operates on the time as an absolute duration since the zero time; it does not operate on the presentation form of the time. Thus, Round(Hour) may return a time with a non-zero minute, depending on the time's Location.

func (t Time) Second() int
Second returns the second offset within the minute specified by t, in the range [0, 59].

func (t Time) String() string
String returns the time formatted using the format string

"2006-01-02 15:04:05.999999999 -0700 MST"
If the time has a monotonic clock reading, the returned string includes a final field "m=±<value>", where value is the monotonic clock reading formatted as a decimal number of seconds.

The returned string is meant for debugging; for a stable serialized representation, use t.MarshalText, t.MarshalBinary, or t.Format with an explicit format string.
*/

// func (t Time) Sub(u Time) Duration
// Sub returns the duration t-u. If the result exceeds the maximum (or minimum) value that can be stored in a Duration, the maximum (or minimum) duration will be returned. To compute t-d for a duration d, use t.Add(-d).
func SubTest() {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)

	difference := end.Sub(start)
	fmt.Printf("difference = %v\n", difference)

}
/*
Output:

difference = 12h0m0s
*/

/*
func (Time) UnixNano ¶
func (t Time) UnixNano() int64
UnixNano returns t as a Unix time, the number of nanoseconds elapsed since January 1, 1970 UTC. The result is undefined if the Unix time in nanoseconds cannot be represented by an int64 (a date before the year 1678 or after 2262). Note that this means the result of calling UnixNano on the zero Time is undefined. The result does not depend on the location associated with t.

func (*Time) UnmarshalBinary ¶
added in go1.2
func (t *Time) UnmarshalBinary(data []byte) error
UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.

func (*Time) UnmarshalJSON ¶
func (t *Time) UnmarshalJSON(data []byte) error
UnmarshalJSON implements the json.Unmarshaler interface. The time is expected to be a quoted string in RFC 3339 format.

func (*Time) UnmarshalText ¶
added in go1.2
func (t *Time) UnmarshalText(data []byte) error
UnmarshalText implements the encoding.TextUnmarshaler interface. The time is expected to be in RFC 3339 format.

func (Time) Weekday ¶
func (t Time) Weekday() Weekday
Weekday returns the day of the week specified by t.

func (Time) Year ¶
func (t Time) Year() int
Year returns the year in which t occurs.

func (Time) YearDay ¶
added in go1.1
func (t Time) YearDay() int
YearDay returns the day of the year specified by t, in the range [1,365] for non-leap years, and [1,366] in leap years.

func (Time) Zone ¶
func (t Time) Zone() (name string, offset int)
Zone computes the time zone in effect at time t, returning the abbreviated name of the zone (such as "CET") and its offset in seconds east of UTC.

type Timer ¶
type Timer struct {
	C <-chan Time
	// contains filtered or unexported fields
}
The Timer type represents a single event. When the Timer expires, the current time will be sent on C, unless the Timer was created by AfterFunc. A Timer must be created with NewTimer or AfterFunc.

func AfterFunc ¶
func AfterFunc(d Duration, f func()) *Timer
AfterFunc waits for the duration to elapse and then calls f in its own goroutine. It returns a Timer that can be used to cancel the call using its Stop method.

func NewTimer ¶
func NewTimer(d Duration) *Timer
NewTimer creates a new Timer that will send the current time on its channel after at least duration d.

func (*Timer) Reset ¶
added in go1.1
func (t *Timer) Reset(d Duration) bool
Reset changes the timer to expire after duration d. It returns true if the timer had been active, false if the timer had expired or been stopped.

For a Timer created with NewTimer, Reset should be invoked only on stopped or expired timers with drained channels.

If a program has already received a value from t.C, the timer is known to have expired and the channel drained, so t.Reset can be used directly. If a program has not yet received a value from t.C, however, the timer must be stopped and—if Stop reports that the timer expired before being stopped—the channel explicitly drained:

if !t.Stop() {
	<-t.C
}
t.Reset(d)
This should not be done concurrent to other receives from the Timer's channel.

Note that it is not possible to use Reset's return value correctly, as there is a race condition between draining the channel and the new timer expiring. Reset should always be invoked on stopped or expired channels, as described above. The return value exists to preserve compatibility with existing programs.

For a Timer created with AfterFunc(d, f), Reset either reschedules when f will run, in which case Reset returns true, or schedules f to run again, in which case it returns false. When Reset returns false, Reset neither waits for the prior f to complete before returning nor does it guarantee that the subsequent goroutine running f does not run concurrently with the prior one. If the caller needs to know whether the prior execution of f is completed, it must coordinate with f explicitly.

func (*Timer) Stop ¶
func (t *Timer) Stop() bool
Stop prevents the Timer from firing. It returns true if the call stops the timer, false if the timer has already expired or been stopped. Stop does not close the channel, to prevent a read from the channel succeeding incorrectly.

To ensure the channel is empty after a call to Stop, check the return value and drain the channel. For example, assuming the program has not received from t.C already:

if !t.Stop() {
	<-t.C
}
This cannot be done concurrent to other receives from the Timer's channel or other calls to the Timer's Stop method.

For a timer created with AfterFunc(d, f), if t.Stop returns false, then the timer has already expired and the function f has been started in its own goroutine; Stop does not wait for f to complete before returning. If the caller needs to know whether f is completed, it must coordinate with f explicitly.

type Weekday ¶
type Weekday int
A Weekday specifies a day of the week (Sunday = 0, ...).

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
func (Weekday) String ¶
func (d Weekday) String() string
String returns the English name of the day ("Sunday", "Monday", ...).
*/