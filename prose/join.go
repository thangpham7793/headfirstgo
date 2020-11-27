package prose

import (
	"errors"
	"strings"
)

//JoinWithCommas join 2 or 3 more words
func JoinWithCommas(phrases ...string) (string, error) {
	switch len(phrases) {
	case 0, 1:
		return "", errors.New("need 2 words or more")
	case 2:
		return strings.Join(phrases, " and "), nil
	default:
		result := strings.Join(phrases[:len(phrases)-1], ", ")
		result += ", and "
		result += phrases[len(phrases)-1]
		return result, nil
	}
}
