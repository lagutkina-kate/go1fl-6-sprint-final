package service

import (
	"fmt"
	"go1fl-sprint6-final/pkg/morse"
)

var (
	point rune = '.'
	dash  rune = '-'
	space rune = ' '
)

func ConvertMorse(text string) (string, error) {
	if len(text) == 0 {
		return "", fmt.Errorf("text is empty")
	}

	isMorse := true
	for _, v := range text {
		if v == point || v == dash || v == space {
			continue
		} else {
			isMorse = false
			break
		}
	}

	if isMorse {
		convertText := morse.ToText(text)
		return convertText, nil
	} else {
		convertMorse := morse.ToMorse(text)
		return convertMorse, nil
	}
}
