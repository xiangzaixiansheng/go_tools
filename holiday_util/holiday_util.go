package holiday_util

import (
	"fmt"
	"time"

	"github.com/6tail/lunar-go/HolidayUtil"
	"github.com/6tail/lunar-go/calendar"
)

var WEEK = []string{"日", "一", "二", "三", "四", "五", "六"}

//holiday工具类🔧
type holidayUtils struct{}

func NewHolidayUtil() *holidayUtils {
	return &holidayUtils{}
}

//判断今天是是否是假期, 返回是否是假期，和假期是什么
func (h holidayUtils) CheckTimeIsHoliday(t time.Time) (bool, string) {
	lunar := calendar.NewSolarFromYmd(t.Year(), int(t.Month()), t.Day())
	ymd := lunar.ToYmd() //年-月-日
	d := HolidayUtil.GetHoliday(ymd)
	if d == nil {
		week := lunar.GetWeekInChinese()
		if week == "日" || week == "六" {
			return true, fmt.Sprintf("星期%s", week)
		}
		return false, ""
	}
	return true, d.GetName()
}

//查询是否是周末，如果不是返回距离周末的时间
func (h holidayUtils) GetWeekend(t time.Time) (bool, int) {
	lunar := calendar.NewSolarFromYmd(t.Year(), int(t.Month()), t.Day())
	week := lunar.GetWeek()
	if week == 0 || week == 6 {
		return true, 0
	}
	return false, 6 - week - 1
}

//查询距离最近的法定假期多少天
func (h holidayUtils) GetNextHoliday() (string, int) {
	t := time.Now()
	thisYearEndDay := time.Date(t.Year(), 12, 31, 0, 0, 0, 0, time.Local)
	endDay := int(thisYearEndDay.Sub(t).Hours() / 24)
	fmt.Println("endDay", endDay)

	name := ""
	i := 0
	lunar := calendar.NewSolarFromYmd(t.Year(), int(t.Month()), t.Day())
	for {
		if i >= endDay {
			i = -1
			break
		}
		h := HolidayUtil.GetHoliday(lunar.ToYmd())
		if h == nil {
			lunar = lunar.Next(1, false)
			i++
			continue
		}
		if !h.IsWork() {
			name = h.GetName()
			break
		}
		lunar = lunar.Next(1, false)
		i++
	}
	return name, i

}
