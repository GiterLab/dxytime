package dxytime

import (
	"time"
	"errors"
	"os"
)

const time_lay_out = "2006-01-02 15:04:05"

var time_zone string

// 设置要获取的时间的时区
func SetTimeZone(loc string) {
	time_zone = loc
}

// UTC时间戳转换为用户设置的时区的时间字符串
func UTCToTimeZoneString(at int64) (str string, err error) {
	if time_zone == "" {
		return "", errors.New("time_zone must be set")
	}
	loc, err := time.LoadLocation(time_zone)
	if err != nil {
		os.Exit(0)
	}
	str = time.Unix(at, 0).In(loc).Format(time_lay_out)
	return str, nil
}

// 时间转换为用户设置的时区字符串
func TimeToTimeZoneString(time time.Time) (str string, err error) {
	str, err = UTCToTimeZoneString(time.Unix())
	if err != nil {
		return "", err
	}
	return str, nil
}

// 字符串转换为用户设置的时区的时间
func StringToTimeZoneTime(str, tz string) (t *time.Time, err error) {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		os.Exit(0)
	}
	t, err = time.ParseInLocation(time_lay_out, str, loc)
	if err != nil {
		return nil, errors.New("string to time fail")
	}
	return &t, nil
}

// 字符串转换为本地时区的时间
func StringToLocalTime(str string) (t *time.Time, err error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		os.Exit(0)
	}
	t, err = time.ParseInLocation(time_lay_out, str, loc)
	if err != nil {
		return nil, errors.New("string to time fail")
	}
	return &t, nil
}