package models

import "errors"

type Theme string

const (
	ThemeLight Theme = "light"
	ThemeDark  Theme = "dark"
	ThemeMerqury  Theme = "merqury"
)

func ParseTheme(str string) (Theme, error) {
	switch str {
	case string(ThemeLight):
		return ThemeLight, nil
	case string(ThemeDark):
		return ThemeDark, nil
	case string(ThemeMerqury):
		return ThemeMerqury, nil
	}
	return ThemeDark, errors.New("unknown theme " + str)
}
