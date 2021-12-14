package complicatedwires

import (
	"errors"
	"sort"
	"strings"
)

func validColors() []string {
	return []string{"B", "BW", "BR", "R", "RW", "W"}
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func normalizeColor(color string) (string, error) {

	color = strings.ToUpper(color)
	color = sortString(color)

	if !contains(validColors(), color) {
		return "", errors.New("not a valid color")
	}

	return color, nil
}
