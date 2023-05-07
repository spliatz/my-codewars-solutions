package kata

import (
	"fmt"
	"strconv"
)

func Solution(list []int) string {

	str := ""

	sub := make([]int, 0)
	isSub := false

	for i := 0; i < len(list); i++ {

		elem := list[i]

		if i == len(list)-1 {
			if isSub && len(sub) >= 2 {
				sub = append(sub, elem)
				str += fmt.Sprintf("%d-%d,", sub[0], sub[len(sub)-1])
				continue
			} else {
				if !isSub && len(sub) >= 3 {
					str += fmt.Sprintf("%d-%d,", sub[0], sub[len(sub)-1])
					str += strconv.Itoa(elem)
					continue
				}
				if len(sub) < 3 {
					sub = append(sub, elem)
					for _, item := range sub {
						str += strconv.Itoa(item) + ","
					}
					continue
				}
			}
		}

		next := list[i+1]

		if isSub && elem+1 != next {
			isSub = false
			sub = append(sub, elem)

			if len(sub) < 3 {
				for _, item := range sub {
					str += strconv.Itoa(item) + ","
				}
			} else {
				str += fmt.Sprintf("%d-%d,", sub[0], sub[len(sub)-1])
			}

			sub = make([]int, 0)
			continue
		}

		if isSub {
			sub = append(sub, elem)
			continue
		}

		if elem+1 == next {
			isSub = true
			sub = append(sub, elem)
			continue
		}

		if !isSub {
			str += strconv.Itoa(elem) + ","
			continue
		}

	}

	if string(str[len(str)-1]) == "," {
		return str[:len(str)-1]
	}

	return str
}
