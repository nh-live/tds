package date

import (
	"time"
	"github.com/stephenlyu/tds"
)

const DAY_MILLISECONDS = 24 * 60 * 60 * 1000
const SECOND_FORMAT = "20060102 15:04:05"
const DAY_FORMAT = "20060102"

func GetDayTimestamp(date uint64) uint64 {
	t := time.Unix(int64(date) / 1000, (int64(date) % 1000) * int64(time.Millisecond)).In(tds.Local)
	d := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, tds.Local)

	return uint64(d.UnixNano() / int64(time.Millisecond))
}

func GetDateDay(date uint64) int {
	d := time.Unix(int64(date) / 1000, (int64(date) % 1000) * int64(time.Millisecond)).In(tds.Local)
	return d.Year() * 10000 + int(d.Month()) * 100 + d.Day()
}

func GetDateWeek(date uint64) int {
	d := time.Unix(int64(date) / 1000, (int64(date) % 1000) * int64(time.Millisecond)).In(tds.Local)

	y, week := d.ISOWeek()

	return y * 100 + week
}

func GetDateWeekDay(date uint64) int {
	d := time.Unix(int64(date) / 1000, (int64(date) % 1000) * int64(time.Millisecond)).In(tds.Local)
	return int(d.Weekday())
}

func GetDateMonth(date uint64) int {
	d := time.Unix(int64(date) / 1000, (int64(date) % 1000) * int64(time.Millisecond)).In(tds.Local)
	return d.Year() * 100 + int(d.Month())
}

var monthQuarterMap = map[int]int {
	1: 3, 2: 3, 3: 3,
	4: 6, 5: 6, 6: 6,
	7: 9, 8: 9, 9: 9,
	10: 12, 11: 12, 12: 12,
}
func GetDateQuarter(date uint64) int {
	d := time.Unix(int64(date) / 1000, (int64(date) % 1000) * int64(time.Millisecond)).In(tds.Local)
	return d.Year() * 100 + monthQuarterMap[int(d.Month())]
}

func GetDateYear(date uint64) int {
	d := time.Unix(int64(date) / 1000, (int64(date) % 1000) * int64(time.Millisecond)).In(tds.Local)
	return d.Year()
}

func Timestamp2DayString(ts uint64) string {
	d := time.Unix(int64(ts) / 1000, (int64(ts) % 1000) * int64(time.Millisecond)).In(tds.Local)
	return d.Format(DAY_FORMAT)
}

func Timestamp2SecondString(ts uint64) string {
	d := time.Unix(int64(ts) / 1000, (int64(ts) % 1000) * int64(time.Millisecond)).In(tds.Local)
	return d.Format(SECOND_FORMAT)
}

func SecondString2Timestamp(date string) (uint64, error) {
	t, err := time.ParseInLocation(SECOND_FORMAT, date, tds.Local)
	if err != nil {
		return 0, err
	}
	return uint64(t.UnixNano() / int64(time.Millisecond)), nil
}

func DayString2Timestamp(date string) (uint64, error) {
	t, err := time.ParseInLocation(DAY_FORMAT, date, tds.Local)
	if err != nil {
		return 0, err
	}
	return uint64(t.UnixNano() / int64(time.Millisecond)), nil
}

func Tick() int64 {
	t := time.Now()
	return t.UnixNano() / int64(time.Millisecond)
}

func GetNowString() string {
	t := time.Now()
	t = t.In(tds.Local)
	return t.Format(SECOND_FORMAT)
}

func GetTodayString() string {
	return ToDayString(time.Now())
}

func ToDayString(t time.Time) string {
	t = t.In(tds.Local)
	return t.Format(DAY_FORMAT)
}

func GetTodayInt() int {
	return ToDayInt(time.Now())
}

func ToDayInt(t time.Time) int {
	t = t.In(tds.Local)
	return t.Year() * 10000 + int(t.Month()) * 100 + t.Day()
}
