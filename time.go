package common_tools

import (
	"strings"
	"time"
)

//time handle
const (
	Year      = "2006"
	YearMonth = "2006-01"
	Date      = "2006-01-02"
	DateTime  = "2006-01-02 15:04:05"

	Month    = "01"
	Day      = "02"
	MonthDay = "01-02"
	Hour     = "15"
	Minute   = "04"
	Second   = "05"
	Time     = "15:04:05"

	_year      = "year"
	_yearMonth = "yearMonth"
	_date      = "date"
	_dateTime  = "dateTime"

	_month    = "month"
	_day      = "day"
	_monthDay = "monthDay"
	_hour     = "hour"
	_minute   = "minute"
	_second   = "second"
	_time     = "time"

	defaultLoc = "Asia/Shanghai"
)

// ForeignTimeToLocTime
// str format is like 01/02/2021 7:42:48 GMT
func ForeignTimeToLocTime(str string, loc string, format string) (t string) {
	t = str
	arr0 := strings.Split(str, " ")
	if len(arr0) != 3 {
		return
	}
	s1 := arr0[0] //date
	s2 := arr0[1] //time
	s3 := arr0[2] //type
	if strings.Contains(s1, "/") {
		s1 = strings.Replace(s1, "/", "-", -1)
	}
	arr1 := strings.Split(s1, "-")
	if len(arr1) != 3 {
		return
	}
	if len(arr1[2]) == 4 {
		s1 = arr1[2] + "-" + arr1[1] + "-" + arr1[0]
	}
	tn := s1 + " " + s2 + " " + s3
	if format == "" {
		format = DateTime
	}
	{
		m := format + " " + s3
		tm, _ := time.Parse(m, tn)
		if loc == "" {
			loc = defaultLoc
		}
		timeLocation, _ := time.LoadLocation(loc)
		time.Local = timeLocation
		t = tm.In(time.Local).Format(format)
	}
	return
}

//ForeignTimeToLocTimeByFormat
// order refund format = "2006-01-02T15:04:05-07:00"     -- 2021-08-04T19:40:48+03:00
// order        format = "2006-01-02T15:04:05.999Z"      -- 2021-07-07T09:52:50.178Z
func ForeignTimeToLocTimeByFormat(str, format, toFormat, loc string) (t string) {
	ts, _ := time.Parse(format, str)
	if loc == "" {
		loc = defaultLoc
	}
	timeLocation, _ := time.LoadLocation(loc)
	time.Local = timeLocation
	if toFormat == "" {
		toFormat = DateTime
	}
	t = ts.In(time.Local).Format(toFormat)
	return
}

/*
TimeToSlice
	s is from query time
	t is year or month or day
*/
func TimeToSlice(s []string, t, loc string) (res []string, err error) {
	//t 是年 或者 月 或者 日
	st := s[0]
	se := s[1]
	if loc == "" {
		loc = defaultLoc
	}
	location, _ := time.LoadLocation(loc)
	stT, _ := time.ParseInLocation(DateTime, st, location)
	seT, _ := time.ParseInLocation(DateTime, se, location)

	format := DateTime
	if t == _year {
		format = Year
	} else if t == _yearMonth {
		format = YearMonth
	} else if t == _date {
		format = Date
	} else {
		panic("t is only support year or month or day")
	}
	for {
		if stT.Format(format) < seT.Format(format) {
			res = append(res, stT.Format(format))
		} else if stT.Format(format) == seT.Format(format) {
			res = append(res, stT.Format(format))
			break
		} else if stT.Format(format) > seT.Format(format) {
			break
		}
		if t == _year {
			stT = stT.AddDate(1, 0, 0)
		} else if t == _yearMonth {
			stT = stT.AddDate(0, 1, 0)
		} else if t == _date {
			stT = stT.AddDate(0, 0, 1)
		}
	}
	return
}

/*
TimeAdd
	s  is like "2021-12-13"
	t  is year or day  or month ...
    n  is sub durations
*/
func TimeAdd(s, t string, n int) (r string) {
	r = s
	var f = DateTime
	switch t {
	case _year:
		f = Year
	case _month, _yearMonth:
		f = YearMonth
	case _day, _monthDay, _date:
		f = Date
	case _dateTime:
		fallthrough
	case _hour:
		fallthrough
	case _minute:
		fallthrough
	case _second:
		fallthrough
	case _time:
		f = DateTime
	default:
		panic("t is not suitable!")
	}
	//fmt.Println(f)
	t0, _ := time.Parse(f, s)
	switch t {
	case _year:
		r = t0.AddDate(n, 0, 0).Format(Year)

	case _month:
		r = t0.AddDate(0, n, 0).Format(Month)
	case _yearMonth:
		r = t0.AddDate(0, n, 0).Format(YearMonth)

	case _day:
		r = t0.AddDate(0, 0, n).Format(Day)
	case _monthDay:
		r = t0.AddDate(0, 0, n).Format(MonthDay)
	case _date:
		r = t0.AddDate(0, 0, n).Format(Date)

	case _dateTime:
		r = t0.Add(time.Second * time.Duration(n)).Format(DateTime)
	case _hour:
		r = t0.Add(time.Hour * time.Duration(n)).Format(Hour)
	case _minute:
		r = t0.Add(time.Minute * time.Duration(n)).Format(Minute)
	case _second:
		r = t0.Add(time.Second * time.Duration(n)).Format(Second)
	case _time:
		r = t0.Add(time.Second * time.Duration(n)).Format(Time)
	}
	return
}
