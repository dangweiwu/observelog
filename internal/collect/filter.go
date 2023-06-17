package collect

import "encoding/json"

func IsJson(txt string) bool {
	return json.Valid([]byte(txt))
}
