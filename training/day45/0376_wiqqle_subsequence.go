package day45

func wiggleMaxLength(nums []int) int {
    if len(nums) == 1 {
        return 1
    }

    res := 1
    prediff := 0
    curdiff := 0

    for i := 0; i < len(nums)-1; i++ {
        curdiff = nums[i+1] - nums[i]
        if prediff <= 0 && curdiff > 0 || prediff >= 0 && curdiff < 0 {
            res++
            prediff = curdiff
        }
    }

    return res
}
