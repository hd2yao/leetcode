package _027_Remove_Element

// fast 是用于寻找 nums 中的 val，也可以理解为寻找新数组中的元素值
// slow 是用于构建新数组
// fast 是萝卜 slow 是坑

// 快慢指针-独立完成
func removeElement(nums []int, val int) int {
    slow, fast := -1, 0
    flag := false
    for fast < len(nums) {
        if nums[fast] == val {
            fast++
            flag = true
            continue
        }
        slow++
        if flag {
            nums[slow] = nums[fast]
        }
        fast++
    }
    return slow + 1
}

// 快慢指针-看完讲解后
func removeElementMod(nums []int, val int) int {
    slow := 0
    for _, num := range nums {
        // num != val 即 num 为新数组中的元素
        // 因此将 num 加入到 slow 的新数组中
        if num != val {
            nums[slow] = num
            slow++
        }
    }
    return slow
}
