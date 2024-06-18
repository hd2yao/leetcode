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
        // 由于上面的 if 条件，可以删去此处 '&& nums[rear] != val'
        if nums[front] == val {
            nums[front], nums[rear] = nums[rear], nums[front]
            front++
            rear--
        }
    }

    if nums[front] == val {
        return front
    } else {
        return front + 1
    }
}

// 快慢指针
func removeElementFastAndSlowPointers(nums []int, val int) int {
    slow, fast := 0, 0

    for fast < len(nums) {
        if nums[fast] != val {
            nums[slow], nums[fast] = nums[fast], nums[slow]
            fast++
            slow++
        } else {
            fast++
        }
    }

    return slow
}

// 快慢指针-改进
func removeElementFastAndSlowPointersOptimize(nums []int, val int) int {
    slow := 0

    for _, num := range nums {
        if num != val {
            nums[slow] = num
            slow++
        }
    }

    return slow
}
