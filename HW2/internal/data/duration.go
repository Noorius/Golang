package data

import (
	"fmt"
	"strconv"
)

type Duration int32

func (d Duration) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d days", d)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
