package _027_Remove_Element

// 使用快慢指针
// 时间复杂度 O(n)
// 快指针：遍历数组，判断元素是否要放入新数组中
// 慢指针：不断将快指针指向的元素放入新数组中
func removeElement(nums []int, val int) int {
    slow := 0
    for _, num := range nums {
        if num != val {
            nums[slow] = num
            slow++
        }
    }
    return slow
}
