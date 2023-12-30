package luhn
import (
    "strings"
    "strconv"
)
func Valid(id string) bool {
    raw := strings.ReplaceAll(id, " ", "")
    
	if len(raw) <= 1 {
        return false
    }
    var judge int
    numList := strings.Split(raw, "")
    flag := true
    for i := len(numList) - 1; i >= 0; i-- {
        tmp, err := strconv.Atoi(numList[i])
        if err != nil {
            return false
        }
        if !flag {
            tmp *= 2
        	if tmp > 9 {
            	tmp -= 9
        	}
        }
        judge += tmp
        flag = !flag
    }
    return judge % 10 == 0
}