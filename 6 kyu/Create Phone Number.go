package kata

import "fmt"

func CreatePhoneNumber(n [10]uint) string {
	arg := make([]interface{}, len(n))
	for i, v := range n {
		arg[i] = v
	}
	return fmt.Sprintf(`(%d%d%d) %d%d%d-%d%d%d%d`, arg...)
}
