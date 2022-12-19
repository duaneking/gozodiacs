package gozodiacs

import (
	"fmt"
	"testing"
	"time"
)

func TestGetWesternZodiacSignPicsesCusp(t *testing.T) {
	testTime, err := time.Parse(ZodiacDayDateFormat, "1982-Mar-20")

	if err != nil {
		t.Fatal(err)
	}

	signs := GetWesternZodiacsForDate(testTime)

	if signs[0] != Pisces || signs[1] != Aries {
		t.Fatal(signs)
	}
}

func TestGetWesternZodiacSignListCorrectSizeAndNoEmptyStrings(t *testing.T) {
	signs := GetAllWesternZodiacSigns()

	if len(signs) != 12 {
		t.Fatal(signs)
	}

	for _, sign := range signs {
		if sign.String() == "" || sign.UnicodeCharacterAsString() == "" {
			t.Fatalf("%v -> %s -> %s", sign, sign.String(), sign.UnicodeCharacterAsString())
		}
	}
}

func TestGetWesternZodiacStrings(t *testing.T) {
	var tests = []struct {
		expected string
		sign     WesternZodiacSign
	}{
		{"Aries", Aries},
		{"Taurus", Taurus},
		{"Gemini", Gemini},
		{"Cancer", Cancer},
		{"Leo", Leo},
		{"Virgo", Virgo},
		{"Libra", Libra},
		{"Scorpio", Scorpio},
		{"Sagittarius", Sagittarius},
		{"Capricorn", Capricorn},
		{"Aquarius", Aquarius},
		{"Pisces", Pisces},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%v", tt.expected, tt.sign)

		t.Run(
			testname,
			func(t *testing.T) {
				if tt.expected != tt.sign.String() {
					t.Fatalf("%v != %v", tt.expected, tt.sign.String())
				}
			},
		)
	}
}

func TestGetWesternZodiacsAsUnicodeCharacterAsStrings(t *testing.T) {
	var tests = []struct {
		expected string
		sign     WesternZodiacSign
	}{
		{"♈︎", Aries},
		{"♉︎", Taurus},
		{"♊︎", Gemini},
		{"♋︎", Cancer},
		{"♌︎", Leo},
		{"♍︎", Virgo},
		{"♎︎", Libra},
		{"♏︎", Scorpio},
		{"♐︎", Sagittarius},
		{"♑︎", Capricorn},
		{"♒︎", Aquarius},
		{"♓︎", Pisces},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%v", tt.expected, tt.sign)

		t.Run(
			testname,
			func(t *testing.T) {
				if tt.expected != tt.sign.UnicodeCharacterAsString() {
					t.Fatalf("%v != %v", tt.expected, tt.sign.UnicodeCharacterAsString())
				}
			},
		)
	}
}

func TestGetWesternZodiacSignsOverYearRange(t *testing.T) {
	min := 1900
	bound := 2000
	max := min + bound

	// First *waves sonic screwdriver* we start in the past..
	tardis := time.Date(min, time.January, 0, 0, 0, 0, 0, time.Local)

	// ...Then we count the years (as usual)
	for tardis.UTC().Year() < max {
		// There are rules.
		dateString := tardis.Format(ZodiacDayDateFormat)

		// LoL
		testTime, err := time.Parse(ZodiacDayDateFormat, dateString)

		// Oops.
		if err != nil {
			t.Fatal(err)
		}

		// ... and look at the planets spin around the sun ...
		signs := GetWesternZodiacsForDate(testTime)

		// Bad news. If this happens, you have to go back and try to save the past. Its all up to you.
		if len(signs) == 0 {
			panic("no western zodiac for given date of " + testTime.String() + ", and thats wierd.")
		}

		// *pulls big lever*
		tardis = tardis.AddDate(0, 0, 1)
	}
}

/* Solar drift says this should actually fail? How Accurate is the Western Zodiac, Cosmologically speaking? */
func BenchmarkGetWesternZodiacSignsOverYearRange(b *testing.B) {
	bound := 5000

	currentYear := time.Now().Year()

	for y := 1900; y < currentYear+bound; y++ {
		dateString := fmt.Sprintf("%d-Mar-20", y)

		testTime, err := time.Parse(ZodiacDayDateFormat, dateString)

		if err != nil {
			b.Fatal(err)
		}

		signs := GetWesternZodiacsForDate(testTime)

		if len(signs) != 2 || signs[0] != Pisces || signs[1] != Aries {
			b.Fatal(err)
		}
	}
}
