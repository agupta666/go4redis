package go4redis

import (
  "errors"
)

func toInteger (intf interface{}) (int, err) {
  length, ok := val.(int)
	if (ok == false) {
		return 0, errors.New("Cannot convert response to interger")
	} else {
    return length, nil
  }
}
