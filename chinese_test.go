package gozodiacs

import (
	"testing"
	"time"

	"fmt"
)

func TestGetChineseZodiacSign(t *testing.T) {
	testTime, err := time.Parse(ZodiacDayDateFormat, "1982-Mar-20")

	if err != nil {
		t.Fatal(err)
	}

	sign := GetChineseZodiacSign(testTime)

	if sign != Dog {
		t.Fatalf("%v", sign)
	}
}

func TestGetChineseZodiacSignsAreNotEmptyStrings(t *testing.T) {

	signs := GetAllChineseZodiacSigns()

	if len(signs) != 12 {
		panic("wrong number of chinese zodiac signs")
	}

	for _, sign := range signs {
		if sign.String() == "" {
			t.Fatalf("%v -> %s", sign, sign.String())
		}
	}
}

/* Scan years and expect a sign. */
func TestGetChineseZodiacSignsOverYearRange(t *testing.T) {
	min := 1900
	bound := 2000
	max := min + bound

	tardis := time.Date(min, time.January, 0, 0, 0, 0, 0, time.Local)

	for tardis.UTC().Year() < max {
		dateString := tardis.Format(ZodiacDayDateFormat)

		testTime, err := time.Parse(ZodiacDayDateFormat, dateString)

		if err != nil {
			t.Fatal(err)
		}

		_ = GetChineseZodiacSign(testTime)

		tardis = tardis.AddDate(0, 0, 1)
	}
}

/* See README: NASA says solar drift should make this fail, eventually. */
func BenchmarkGetChineseZodiacSignsOverYearRange(b *testing.B) {
	bound := 5000

	currentYear := time.Now().Year()

	for y := 1900; y < currentYear+bound; y++ {
		dateString := fmt.Sprintf("%d-Mar-20", y)

		testTime, err := time.Parse(ZodiacDayDateFormat, dateString)

		if err != nil {
			b.Fatal(err)
		}

		sign := GetChineseZodiacSign(testTime)

		if sign != Dog {
			b.Fatal(sign)
		}
	}
}
