package kata

import "strings"

func Solution(str string) []string {
	arr := make([]string, 0)
	strArr := strings.Split(str, "")
	for index, item := range strArr {

		if index%2 == 0 {
			if index+1 == len(strArr) {
				arr = append(arr, item+"_")
				break
			}
			arr = append(arr, item+strArr[index+1])
		}

	}
	return arr
}
