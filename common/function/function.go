package function

import (
	"time"
)

func ConvertT(in int64) (out string) {
	tm := time.Unix(in, 0)
	out = tm.Format("2006-01-02 03:04:05 PM")
	return out
}
