package day50

func lemonadeChange(bills []int) bool {
    change := [2]int{0, 0}

    for i := 0; i < len(bills); i++ {
        if i == 0 && bills[i] != 5 {
            return false
        }
        switch bills[i] {
        case 5:
            change[0]++
        case 10:
            if change[0] <= 0 {
                return false
            }
            change[1]++
            change[0]--
        case 20:
            if change[1] > 0 && change[0] > 0 {
                change[1]--
                change[0]--
            } else if change[1] == 0 && change[0] >= 3 {
                change[0] -= 3
            } else {
                return false
            }
        }
    }
    return true
}
