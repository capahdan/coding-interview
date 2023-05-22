package float

import (
	"fmt"
	"strings"
)

func Solution(N int, nilai float32) string {
	data := fmt.Sprintf("%.2f", nilai)
	data = strings.ReplaceAll(data, ".", "")
	if len(data) != N {
		data = fmt.Sprintf("%0"+fmt.Sprintf("%d", N)+"s", data)
	}
	return data
}
