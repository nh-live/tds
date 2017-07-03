package tds

import (
	"math"
	"fmt"
	"strconv"
	"regexp"
	"errors"
	"time"
)

type Date uint32

const MAX_DATE = math.MaxUint32

func (this Date) DayString() string {
	return fmt.Sprintf("%d", this)
}

func (this Date) MinuteString() string {
	dayValue := uint16(this & 0xFFFF)
	minuteValue := uint16((this >> 16) & 0xFFFF)

	year := (dayValue / 2048) + 2004
	month := (dayValue % 2048) / 100
	day := (dayValue % 2048) % 100

	hour := minuteValue / 60
	minute := minuteValue % 60

	return fmt.Sprintf("%04d%02d%02d %02d:%02d:00", year, month, day, hour, minute)
}

// Day of minute date
func (this Date) MinuteDay() uint32 {
	dayValue := uint32(this & 0xFFFF)
	year := (dayValue / 2048) + 2004
	month := (dayValue % 2048) / 100
	day := (dayValue % 2048) % 100

	return year * 10000 + month * 100 + day
}

func FromDayString(s string) (error, Date) {
	ret, err := strconv.ParseUint(s, 10, 64)
	return err, Date(ret)
}

func FromMinuteString(s string) (error, Date) {
	regExp, err := regexp.Compile("^([0-9]{4})([0-9]{2})([0-9]{2}) ([0-9]{2}):([0-9]{2}):([0-9]{2})$")
	if err != nil {
		return errors.New("bad minute string"), 0
	}

	result := regExp.FindSubmatch([]byte(s))
	if result != nil {
		return errors.New("bad minute string"), 0
	}

	year, _ := strconv.Atoi(string(result[1]))
	month, _ := strconv.Atoi(string(result[2]))
	day, _ := strconv.Atoi(string(result[3]))
	hour, _ := strconv.Atoi(string(result[4]))
	minute, _ := strconv.Atoi(string(result[5]))
	second, _ := strconv.Atoi(string(result[6]))

	if year < 2004 || year > 2004 + 511 {
		return errors.New("bad year"), 0
	}

	if month <= 0 || month > 12 {
		return errors.New("bad month"), 0
	}

	if day <= 0 || day > 31 {
		return errors.New("bad day"), 0
	}

	if hour <= 0 || year >= 24 {
		return errors.New("bad hour"), 0
	}

	if minute <= 0 || minute > 59 {
		return errors.New("bad minute"), 0
	}

	if second <= 0 || second > 59 {
		return errors.New("bad second"), 0
	}

	dayValue := year * 2048 + month * 100 + day
	minuteValue := hour * 60 + minute

	return nil, Date((minuteValue << 16) | dayValue)
}

func GetDateDay(period Period, date uint32) uint32 {
	if period.Unit() == PERIOD_UNIT_MINUTE {
		return Date(date).MinuteDay()
	}
	return date
}

func GetDateWeek(period Period, date uint32) uint32 {
	if period.Unit() == PERIOD_UNIT_MINUTE {
		date = Date(date).MinuteDay()
	}

	year := date / 10000
	month := (date % 10000) / 100
	day := date % 100

	d := time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)

	y, week := d.ISOWeek()

	return uint32(y * 100 + week)
}

func GetDateMonth(period Period, date uint32) uint32 {
	if period.Unit() == PERIOD_UNIT_MINUTE {
		date = Date(date).MinuteDay()
	}
	return date / 100
}

var monthQuarterMap = map[int]uint32 {
	1: 3, 2: 3, 3: 3,
	4: 6, 5: 6, 6: 6,
	7: 9, 8: 9, 9: 9,
	10: 12, 11: 12, 12: 12,
}
func GetDateQuarter(period Period, date uint32) uint32 {
	if period.Unit() == PERIOD_UNIT_MINUTE {
		date = Date(date).MinuteDay()
	}
	year := date / 10000
	month := (date % 10000) / 100
	return year * 100 + monthQuarterMap[int(month)]
}

func GetDateYear(period Period, date uint32) uint32 {
	if period.Unit() == PERIOD_UNIT_MINUTE {
		date = Date(date).MinuteDay()
	}
	return date / 10000
}
