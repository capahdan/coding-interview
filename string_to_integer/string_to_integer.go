package string_to_integer

import "strconv"

func MyAtoi(s string) int {
	stack := ""

	digitMap := map[byte]bool{
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
		6: true,
		7: true,
		8: true,
		9: true,
	}
	strLength := len(s)
	for i := range s {
		switch s[i] {
		case '-':
			// if i != 0 {
			//     if ok := digitMap[s[i-1]] ok {

			//     }
			// }
			stack += string(s[i])
		case '1':
			if i == strLength {
				if ok := digitMap[s[i+1]]; !ok {
					break
				}
			}

			stack += string(s[i])
		case '2':
			if i == strLength {
				if ok := digitMap[s[i+1]]; !ok {
					break
				}
			}
			stack += string(s[i])
		case '3':
			if i == strLength {
				if ok := digitMap[s[i+1]]; !ok {
					break
				}
			}
			stack += string(s[i])
		case '4':
			if i == strLength {
				if ok := digitMap[s[i+1]]; !ok {
					break
				}
			}
			stack += string(s[i])
		case '5':
			if i == strLength {
				if ok := digitMap[s[i+1]]; !ok {
					break
				}
			}
			stack += string(s[i])
		case '6':
			if i == strLength {
				if ok := digitMap[s[i+1]]; !ok {
					break
				}
			}
			stack += string(s[i])
		case '7':
			if i == strLength {
				if ok := digitMap[s[i+1]]; !ok {
					break
				}
			}
			stack += string(s[i])
		case '8':
			if i == strLength {
				if ok := digitMap[s[i+1]]; !ok {
					break
				}
			}
			stack += string(s[i])

		case '9':
			if i == strLength {
				if ok := digitMap[s[i+1]]; !ok {
					break
				}
			}
			stack += string(s[i])
		}
	}

	nilai, _ := strconv.Atoi(stack)
	return nilai

}
