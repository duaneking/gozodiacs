package gozodiacs

// Only require system time module.
import "time"

// The "Western" zodiac sign. - https://en.wikipedia.org/wiki/Western_astrology
type WesternZodiacSign int

// The "Western" zodiac sign as a simple name string.
func (s WesternZodiacSign) String() string {
	switch s {
	case Aries:
		return "Aries"
	case Taurus:
		return "Taurus"
	case Gemini:
		return "Gemini"
	case Cancer:
		return "Cancer"
	case Leo:
		return "Leo"
	case Virgo:
		return "Virgo"
	case Libra:
		return "Libra"
	case Scorpio:
		return "Scorpio"
	case Sagittarius:
		return "Sagittarius"
	case Capricorn:
		return "Capricorn"
	case Aquarius:
		return "Aquarius"
	}

	// case Pisces:
	return "Pisces"
}

// The "Western" zodiac sign as a simple utf-8 emoji rune string.
func (s WesternZodiacSign) UnicodeCharacterAsString() string {
	switch s {
	case Aries:
		return "♈︎"
	case Taurus:
		return "♉︎"
	case Gemini:
		return "♊︎"
	case Cancer:
		return "♋︎"
	case Leo:
		return "♌︎"
	case Virgo:
		return "♍︎"
	case Libra:
		return "♎︎"
	case Scorpio:
		return "♏︎"
	case Sagittarius:
		return "♐︎"
	case Capricorn:
		return "♑︎"
	case Aquarius:
		return "♒︎"
	}

	// case Pisces:
	return "♓︎"
}

// Western Zodiac Signs as Constants
const (
	// Aries western zodiac sign
	Aries WesternZodiacSign = iota

	// Taurus western zodiac sign
	Taurus

	// Gemini western zodiac sign
	Gemini

	// Cancer western zodiac sign
	Cancer

	// Leo western zodiac sign
	Leo

	// Virgo western zodiac sign
	Virgo

	// Libra western zodiac sign
	Libra

	// Scorpio western zodiac sign
	Scorpio

	// Sagittarius western zodiac sign
	Sagittarius

	// Capricorn western zodiac sign
	Capricorn

	// Aquarius western zodiac sign
	Aquarius

	// Pisces western zodiac sign
	Pisces
)

// GetAllWesternZodiacSigns - Return a new list of all possible western zodiac sign constants in sequence. Used for UI/etc.
func GetAllWesternZodiacSigns() []WesternZodiacSign {
	return []WesternZodiacSign{Aries, Taurus, Gemini, Cancer, Leo, Virgo, Libra, Scorpio, Sagittarius, Capricorn, Aquarius, Pisces}
}

// GetWesternZodiacsForDate - Returns the matching zodiac sign(s) for a date. May include an additional cusp zodiac. https://en.wikipedia.org/wiki/Cusp_(astrology)
func GetWesternZodiacsForDate(date time.Time) []WesternZodiacSign {
	switch date.Month() {
	case time.January:
		if date.Day() < 20 {
			return []WesternZodiacSign{Capricorn}
		}

		if date.Day() == 20 {
			return []WesternZodiacSign{Capricorn, Aquarius}
		}

		return []WesternZodiacSign{Aquarius}
	case time.February:
		if date.Day() < 18 {
			return []WesternZodiacSign{Aquarius}
		}

		if date.Day() == 18 {
			return []WesternZodiacSign{Aquarius, Pisces}
		}

		return []WesternZodiacSign{Pisces}
	case time.March:
		if date.Day() < 20 {
			return []WesternZodiacSign{Pisces}
		}

		if date.Day() == 20 {
			return []WesternZodiacSign{Pisces, Aries} // A Cusp is when a date gives 2 signs.
		}

		return []WesternZodiacSign{Aries}
	case time.April:
		if date.Day() < 19 {
			return []WesternZodiacSign{Aries}
		}

		if date.Day() == 19 {
			return []WesternZodiacSign{Aries, Taurus}
		}

		return []WesternZodiacSign{Taurus}
	case time.May:
		if date.Day() < 20 {
			return []WesternZodiacSign{Taurus}
		}

		if date.Day() == 20 {
			return []WesternZodiacSign{Taurus, Gemini}
		}

		return []WesternZodiacSign{Gemini}
	case time.June:
		if date.Day() < 21 {
			return []WesternZodiacSign{Gemini}
		}

		if date.Day() == 21 {
			return []WesternZodiacSign{Gemini, Cancer}
		}

		return []WesternZodiacSign{Cancer}
	case time.July:
		if date.Day() < 22 {
			return []WesternZodiacSign{Cancer}
		}

		if date.Day() == 22 {
			return []WesternZodiacSign{Cancer, Leo}
		}

		return []WesternZodiacSign{Leo}
	case time.August:
		if date.Day() < 22 {
			return []WesternZodiacSign{Leo}
		}

		if date.Day() == 22 {
			return []WesternZodiacSign{Leo, Virgo}
		}

		return []WesternZodiacSign{Virgo}
	case time.September:

		if date.Day() < 22 {
			return []WesternZodiacSign{Virgo}
		}

		if date.Day() == 22 {
			return []WesternZodiacSign{Virgo, Libra}
		}

		return []WesternZodiacSign{Libra}
	case time.October:
		if date.Day() < 23 {
			return []WesternZodiacSign{Libra}
		}

		if date.Day() == 23 {
			return []WesternZodiacSign{Libra, Scorpio}
		}

		return []WesternZodiacSign{Scorpio}
	case time.November:
		if date.Day() < 22 {
			return []WesternZodiacSign{Scorpio}
		}

		if date.Day() == 22 {
			return []WesternZodiacSign{Scorpio, Sagittarius}
		}

		return []WesternZodiacSign{Sagittarius}
	}

	// case time.December:
	if date.Day() < 21 {
		return []WesternZodiacSign{Sagittarius}
	}

	if date.Day() == 21 {
		return []WesternZodiacSign{Sagittarius, Capricorn}
	}

	return []WesternZodiacSign{Capricorn}
}
