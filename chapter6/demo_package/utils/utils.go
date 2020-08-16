package utils
import(
	"fmt"
)
func Cal(n1 float64, n2 float64, op byte) float64{
	switch op {
	case '+':
		return n1 + n2
	case '-':
		return n1 - n2
	case '*':
		return n1 * n2
	case '/':
		return n1 / n2
	default:
		fmt.Println("error : worong operator")
		return 0
	}
}