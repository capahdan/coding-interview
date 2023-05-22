package pantul

import "fmt"

func Pantul(a int) []string {
	var data = []string{}
	data = append(data, fmt.Sprintf("%d", a))
	var b = float32(a)
	for b > 0.5 {
		b = b / 2

		data = append(data, fmt.Sprintf("%.2f", float32(b)))
	}
	return data
}
