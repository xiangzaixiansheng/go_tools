package holiday_util

import (
	"fmt"
	"testing"
	"time"
)

func TestHoliday(t *testing.T) {
	isholiday, _ := NewHolidayUtil().CheckTimeIsHoliday(time.Now())
	fmt.Println("isholiday", isholiday)
	isWeekend, days := NewHolidayUtil().GetWeekend(time.Now())

	fmt.Printf("isWeekend %v 距离days %d", isWeekend, days)

	holiday, days2 := NewHolidayUtil().GetNextHoliday()

	fmt.Printf("holiday %v 距离days %d", holiday, days2)

	a := []string{"a", "b", "c"}
	b := []string{}
	for _, item := range a {
		b = append(b, item)
	}
	fmt.Println(b)

}
