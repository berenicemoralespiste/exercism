package acronym

import (
	"strings"
)

func acronym(s string) string{

	var acronymResponse string
	arrayString := strings.Split(s, " ")
	for _ ,value := range arrayString{
		acronymResponse += string((value) [0])
	}

	return strings.ToUpper(acronymResponse)
}
