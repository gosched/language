package explore

import (
	"fmt"
	"time"
)

// Time .
func Time() {
	// fmt.Println(time.Now().Unix())     // seconds since 1970
	// fmt.Println(time.Now().UnixNano()) // nanoseconds since 1970

	// time.Now()
	// 2020-01-08 18:50:40.087872 +0800 CST m=+0.000100892

	// time.Now().UTC()
	// 2020-01-08 10:50:40.087872 +0000 UTC

	// t := time.Now()

	// fmt.Println(t.String())
	// fmt.Println(t.UTC().String())
	// fmt.Println(t.Local().String())

	// location := time.Now().Location()
	// name, offset := time.Now().Zone()
	// fmt.Println(location, name, offset)

	// taipei, _ := time.LoadLocation("Asia/Taipei")
	// shanghai, _ := time.LoadLocation("Asia/Shanghai")
	// singapore, _ := time.LoadLocation("Asia/Singapore")

	// fmt.Println(taipei, time.Now().In(taipei).String())
	// fmt.Println(shanghai, time.Now().In(shanghai).String())
	// fmt.Println(singapore, time.Now().In(singapore).String())

	// now := time.Now().Format("2006-01-02 15:04:05")
	// fmt.Println(now)

	// year, month, day := t.Date()
	// hour, min, sec := t.Clock()

	// fmt.Printf("%v %v %v %v %v %v %v\n", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
	// fmt.Printf("%v %v %v %v %v %v\n", year, int(month), day, hour, min, sec)

	// year, week := t.ISOWeek()
	// fmt.Printf("%v %v\n", year, week)

	// yearDay, weekday := t.YearDay(), t.Weekday()
	// fmt.Printf("%v %v %v\n", yearDay, int(weekday), weekday)

	// t0 := time.Now()
	// t1 := t0.Add(time.Duration(time.Hour))
	// t2 := t0.AddDate(1, 1, 1)
	// duration := t1.Sub(t0)

	// fmt.Printf("%v\n", t1.String())
	// fmt.Printf("%v\n", t2.String())
	// fmt.Printf("%T %v\n", duration, duration)

	// fmt.Printf("%v\n", t0.Round(time.Duration(time.Hour)).String())
	// fmt.Printf("%v\n", t0.Round(time.Duration(time.Minute)).String())
	// fmt.Printf("%v\n", t0.Round(time.Duration(time.Second)).String())

	// fmt.Printf("%v\n", t0.Truncate(time.Duration(time.Hour)*24).String())
	// fmt.Printf("%v\n", t0.Truncate(time.Duration(time.Minute)).String())
	// fmt.Printf("%v\n", t0.Truncate(time.Duration(time.Second)).String())

	// isZero := t0.IsZero()
	// equal := t0.Equal(t1)
	// after := t0.After(t1)
	// before := t0.Before(t1)
	// fmt.Println(isZero, equal, after, before)
}

// HelloDuration .
func HelloDuration() {
	// 	hours, _ := time.ParseDuration("10h")
	//	complex, _ := time.ParseDuration("1h10m10s")
	///	micro, _ := time.ParseDuration("1µs")
}

// HelloTimer .
func HelloTimer() {
	timer := time.NewTimer(time.Second * 5)
	// timer.Reset()
	// timer.Stop()
	select {
	case <-timer.C:
		fmt.Println("time up")
	}
}

// HelloTicker .
func HelloTicker() {
	ticker := time.NewTicker(time.Second)
	timer := time.After(time.Second * 10)

tag:
	for {
		select {
		case <-ticker.C:
			fmt.Println("happen")
		case <-timer:
			ticker.Stop()
			fmt.Println("stop")
			break tag
		}
	}

	fmt.Println("end")
}

// https://golang.org/pkg/time/

// func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
// func Parse(layout, value string) (Time, error)
// func ParseInLocation(layout, value string, loc *Location) (Time, error)

// func (t Time) AppendFormat(b []byte, layout string) []byte

// func Unix(sec int64, nsec int64) Time

// func (t Time) MarshalBinary() ([]byte, error)
// func (t Time) MarshalJSON() ([]byte, error)
// func (t Time) MarshalText() ([]byte, error)

// func (t *Time) UnmarshalBinary(data []byte) error
// func (t *Time) UnmarshalJSON(data []byte) error
// func (t *Time) UnmarshalText(data []byte) error

// func (t Time) GobEncode() ([]byte, error)
// func (t *Time) GobDecode(data []byte) error
