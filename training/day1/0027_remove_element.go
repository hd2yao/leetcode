package day1

// 双指针法
func removeElementTwoPointers(nums []int, val int) int {
    if len(nums) == 0 {
        return 0
    }

    front, rear := 0, len(nums)-1
    for front < rear {
        if nums[front] != val {
            front++
            continue
        }
        if nums[rear] == val {
            rear--
            continue
        }
        if nums[front] == val && nums[rear] != val {
            nums[front], nums[rear] = nums[rear], nums[front]
            front++
            rear--
        }
    }

    if nums[front] == 2 {
        return front
    } else {
        return front + 1
    }
}
