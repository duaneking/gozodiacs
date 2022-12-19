package gozodiacs

// Only require system time module.
import "time"

// ChineseZodiacSign represents the chinese zodiac signs based on the year of a persons birth. - https://en.wikipedia.org/wiki/Chinese_zodiac
type ChineseZodiacSign int

// The "Chinese" zodiac sign as a simple name string.
func (s ChineseZodiacSign) String() string {
	switch s {
	case Monkey:
		return "Monkey"
	case Rooster:
		return "Rooster"
	case Dog:
		return "Dog"
	case Pig:
		return "Pig"
	case Rat:
		return "Rat"
	case Ox:
		return "Ox"
	case Tiger:
		return "Tiger"
	case Rabbit:
		return "Rabbit"
	case Dragon:
		return "Dragon"
	case Snake:
		return "Snake"
	case Horse:
		return "Horse"
	}

	// case Goat:
	return "Goat"
}

// Chinese Zodiac Signs as Constants
const (
	// Monkey chinese zodiac sign
	Monkey ChineseZodiacSign = iota

	// Rooster chinese zodiac sign
	Rooster

	// Dog chinese zodiac sign
	Dog

	// Pig chinese zodiac sign
	Pig

	// Rat chinese zodiac sign
	Rat

	// Ox chinese zodiac sign
	Ox

	// Tiger chinese zodiac sign
	Tiger

	// Rabbit chinese zodiac sign
	Rabbit

	// Dragon chinese zodiac sign
	Dragon

	// Snake chinese zodiac sign
	Snake

	// Horse chinese zodiac sign
	Horse

	// Goat chinese zodiac sign
	Goat
)

// GetAllChineseZodiacSigns - Return a new list of all possible chinese zodiac sign constants in sequence. Used for UI/etc.
func GetAllChineseZodiacSigns() []ChineseZodiacSign {
	return []ChineseZodiacSign{Monkey, Rooster, Dog, Pig, Rat, Ox, Tiger, Rabbit, Dragon, Snake, Horse, Goat}
}

// GetChineseZodiacSign returns the single matching chinese zodiac sign based on the year given.
func GetChineseZodiacSign(date time.Time) ChineseZodiacSign {

	chineseZodiac := int(date.Year() - ((date.Year() / 12) * 12))

	switch chineseZodiac {
	case 0:
		return Monkey

	case 1:
		return Rooster

	case 2:
		return Dog

	case 3:
		return Pig

	case 4:
		return Rat

	case 5:
		return Ox

	case 6:
		return Tiger

	case 7:
		return Rabbit

	case 8:
		return Dragon

	case 9:
		return Snake

	case 10:
		return Horse
	}

	return Goat
}
