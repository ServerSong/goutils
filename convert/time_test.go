package convert

import (
	"testing"
	"time"
)

func Test_time(t *testing.T) {
	t.Log(TimeToStr(time.Now()))
	t.Log(TimeToTimestamp(time.Now()))
}


func Test_timestamp(t *testing.T) {
	t.Log(TimestampToTime(1553161842))
	t.Log(TimestampToStr(1553161842))
}

func Test_str(t *testing.T) {
	t.Log(StrToTime("2019-03-21 17:51:00"))
	t.Log(StrToTimestamp("2019-03-21 17:51:00"))
	t.Log(ValueToTime(2019, 3, 21, 17, 51, 0, 0))
}