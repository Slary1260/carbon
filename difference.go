package carbon

import (
	"math"
	"strings"
)

// DiffInYears gets the difference in years.
// 相差多少年
func (c Carbon) DiffInYears(carbon ...Carbon) int64 {
	start, end := c, c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}

	dy, dm, dd := end.Year()-start.Year(), end.Month()-start.Month(), end.Day()-start.Day()
	if dd < 0 || dm < 0 {
		dy--
	}
	if dy < 0 && (dd != 0 || dm != 0) {
		dy++
	}
	return int64(dy)
}

// DiffAbsInYears gets the difference in years with absolute value.
// 相差多少年(绝对值)
func (c Carbon) DiffAbsInYears(carbon ...Carbon) int64 {
	return getAbsValue(c.DiffInYears(carbon...))
}

// DiffInMonths gets the difference in months.
// 相差多少月
func (c Carbon) DiffInMonths(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	startYear, startMonth, startDay := c.Date()
	endYear, endMonth, endDay := end.Date()

	diffYear, diffMonth, diffDay := endYear-startYear, endMonth-startMonth, endDay-startDay
	if diffDay < 0 {
		diffMonth = diffMonth - 1
	}
	if diffYear == 0 && diffMonth == 0 {
		return int64(0)
	}
	if diffYear == 0 && diffMonth != 0 && diffDay != 0 {
		if int(end.DiffAbsInHours(c)) < c.DaysInMonth()*HoursPerDay {
			return int64(0)
		}
		return int64(diffMonth)
	}
	return int64(diffYear*MonthsPerYear + diffMonth)
}

// DiffAbsInMonths gets the difference in months with absolute value.
// 相差多少月(绝对值)
func (c Carbon) DiffAbsInMonths(carbon ...Carbon) int64 {
	return getAbsValue(c.DiffInMonths(carbon...))
}

// DiffInWeeks gets the difference in weeks.
// 相差多少周
func (c Carbon) DiffInWeeks(carbon ...Carbon) int64 {
	start, end := c, c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return int64(math.Floor(float64((end.Timestamp() - start.Timestamp()) / (7 * 24 * 3600))))
}

// DiffAbsInWeeks gets the difference in weeks with absolute value.
// 相差多少周(绝对值)
func (c Carbon) DiffAbsInWeeks(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInWeeks(end))
}

// DiffInDays gets the difference in days.
// 相差多少天
func (c Carbon) DiffInDays(carbon ...Carbon) int64 {
	start, end := c, c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return int64(math.Floor(float64((end.Timestamp() - start.Timestamp()) / (24 * 3600))))
}

// DiffAbsInDays gets the difference in days with absolute value.
// 相差多少天(绝对值)
func (c Carbon) DiffAbsInDays(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInDays(end))
}

// DiffInHours gets the difference in hours.
// 相差多少小时
func (c Carbon) DiffInHours(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return c.DiffInSeconds(end) / SecondsPerHour
}

// DiffAbsInHours gets the difference in hours with absolute value.
// 相差多少小时(绝对值)
func (c Carbon) DiffAbsInHours(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInHours(end))
}

// DiffInMinutes gets the difference in minutes.
// 相差多少分钟
func (c Carbon) DiffInMinutes(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return c.DiffInSeconds(end) / SecondsPerMinute
}

// DiffAbsInMinutes gets the difference in minutes with absolute value.
// 相差多少分钟(绝对值)
func (c Carbon) DiffAbsInMinutes(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInMinutes(end))
}

// DiffInSeconds gets the difference in seconds.
// 相差多少秒
func (c Carbon) DiffInSeconds(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return end.Timestamp() - c.Timestamp()
}

// DiffAbsInSeconds gets the difference in seconds with absolute value.
// 相差多少秒(绝对值)
func (c Carbon) DiffAbsInSeconds(carbon ...Carbon) int64 {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	return getAbsValue(c.DiffInSeconds(end))
}

// DiffInString gets the difference in string, i18n is supported.
// 相差字符串，支持i18n
func (c Carbon) DiffInString(carbon ...Carbon) string {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	if c.Error != nil || end.Error != nil {
		return ""
	}
	unit, value := c.diff(end)
	return c.lang.translate(unit, value)
}

// DiffAbsInString gets the difference in string with absolute value, i18n is supported.
// 相差字符串，支持i18n(绝对值)
func (c Carbon) DiffAbsInString(carbon ...Carbon) string {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	if c.Error != nil || end.Error != nil {
		return ""
	}
	unit, value := c.diff(end)
	return c.lang.translate(unit, getAbsValue(value))
}

// DiffForHumans gets the difference in a human-readable format, i18n is supported.
// 获取对人类友好的可读格式时间差，支持i18n
func (c Carbon) DiffForHumans(carbon ...Carbon) string {
	end := c.Now()
	if len(carbon) > 0 {
		end = carbon[len(carbon)-1]
	}
	if c.Error != nil || end.Error != nil {
		return ""
	}
	unit, value := c.diff(end)
	translation := c.lang.translate(unit, getAbsValue(value))
	if unit == "now" {
		return translation
	}
	if c.Lt(end) && len(carbon) == 0 {
		return strings.Replace(c.lang.resources["ago"], "%s", translation, 1)
	}
	if c.Lt(end) && len(carbon) > 0 {
		return strings.Replace(c.lang.resources["before"], "%s", translation, 1)
	}
	if c.Gt(end) && len(carbon) == 0 {
		return strings.Replace(c.lang.resources["from_now"], "%s", translation, 1)
	}
	return strings.Replace(c.lang.resources["after"], "%s", translation, 1)
}

// gets the difference for unit and value.
// 获取相差单位和差值
func (c Carbon) diff(end Carbon) (unit string, value int64) {
	switch true {
	case c.DiffAbsInYears(end) > 0:
		unit = "year"
		value = c.DiffInYears(end)
		break
	case c.DiffAbsInMonths(end) > 0:
		unit = "month"
		value = c.DiffInMonths(end)
		break
	case c.DiffAbsInWeeks(end) > 0:
		unit = "week"
		value = c.DiffInWeeks(end)
		break
	case c.DiffAbsInDays(end) > 0:
		unit = "day"
		value = c.DiffInDays(end)
		break
	case c.DiffAbsInHours(end) > 0:
		unit = "hour"
		value = c.DiffInHours(end)
		break
	case c.DiffAbsInMinutes(end) > 0:
		unit = "minute"
		value = c.DiffInMinutes(end)
	case c.DiffAbsInSeconds(end) > 0:
		unit = "second"
		value = c.DiffInSeconds(end)
	case c.DiffAbsInSeconds(end) == 0:
		unit = "now"
		value = 0
	}
	return
}
