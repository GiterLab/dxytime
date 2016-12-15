package dxytime

import (
	"testing"
)

func Test_Convert(t *testing.T) {
	SetTimeZone("Asia/Shanghai")
	var at int64
	at = 1481763382
	str,err :=UTCToTimeZoneString(at)
	if err != nil {
		t.Error(err)
	}
	if str != "2016-12-15 08:56:22"{
		t.Error(str)
	}
	tm,err :=StringToTimeZoneTime(str,"UTC")
	if err != nil {
		t.Error(err)
	}
	if tm.Unix() != at {
		t.Error(tm.Unix())
	}
	str1,err :=TimeToTimeZoneString(*tm)
	if err != nil {
		t.Error(err)
	}
	if str1 != "2016-12-15 08:56:22"{
		t.Error(str1)
	}
	tm1,err :=StringToLocalTime(str1)
	if err != nil {
		t.Error(err)
	}
	if tm1.Unix() != at {
		t.Error(tm1.Unix())
	}
}



