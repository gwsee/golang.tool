package easy

import (
	"strings"
	"time"
)

func TimeSub(t1, t2 time.Time) int {
	t1 = t1.UTC().Truncate(24 * time.Hour)
	t2 = t2.UTC().Truncate(24 * time.Hour)
	return int(t1.Sub(t2).Hours() / 24)
}
func TimeToSlice(s []string, t string) (res []string, err error) {
	//t 是年 或者 月 或者 日
	st := s[0]
	se := s[1]
	loc, _ := time.LoadLocation("Asia/Shanghai")
	stT, _ := time.ParseInLocation("2006-01-02 15:04:05", st, loc)
	seT, _ := time.ParseInLocation("2006-01-02 15:04:05", se, loc)
	format := "2006-01-02 15:04:05"
	if t == "year" {
		format = "2006"
	} else if t == "month" {
		format = "2006-01"
	} else if t == "day" {
		format = "2006-01-02"
	}
	for {
		// fmt.Println(stT.Format(format),"--》---《--",seT.Format(format))
		if stT.Format(format) < seT.Format(format) {
			res = append(res, stT.Format(format))
		} else if stT.Format(format) == seT.Format(format) {
			res = append(res, stT.Format(format))
			break
		} else if stT.Format(format) > seT.Format(format) {
			break
		}
		if t == "year" {
			stT = stT.AddDate(1, 0, 0)
		} else if t == "month" {
			stT = stT.AddDate(0, 1, 0)
		} else if t == "day" {
			stT = stT.AddDate(0, 0, 1)
		}
	}
	return
}
func ForeignTimeToChinaTime(str string) (t string) {
	t = str
	arr0 := strings.Split(str, " ")
	if len(arr0) != 3 {
		return
	}
	s1 := arr0[0] //年月日
	s2 := arr0[1] //时间
	s3 := arr0[2] //类型
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
	format := "2006-01-02 15:04:05"
	{
		m := format + " " + s3
		tm, _ := time.Parse(m, tn)
		timeLocation, _ := time.LoadLocation("Asia/Shanghai")
		time.Local = timeLocation
		t = tm.In(time.Local).Format(format)
	}
	return
}

func ForeignTimeToChinaTimeByFormat(str string, format string) (t string) {
	//订单退货表 format = "2006-01-02T15:04:05-07:00"     -- 2021-08-04T19:40:48+03:00
	//订单表 format = "2006-01-02T15:04:05.999Z" "2021-07-07T09:52:50.178Z"
	ts, _ := time.Parse(format, str)
	timeLocation, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = timeLocation
	t = ts.In(time.Local).Format("2006-01-02 15:04:05")
	return
}
