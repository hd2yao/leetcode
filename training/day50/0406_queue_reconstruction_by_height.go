package day50

import "sort"

func reconstructQueue(people [][]int) [][]int {
    sort.Slice(people, func(i, j int) bool {
        if people[i][0] == people[j][0] {
            return people[i][1] < people[j][1]
        }
        return people[i][0] > people[j][0]
    })

    queue := make([][]int, 0)

    for i := 0; i < len(people); i++ {
        index := people[i][1]
        queue = append(queue[:index], append([][]int{people[i]}, queue[index:]...)...)
    }

    //for i, person := range people {
    //	copy(people[person[1]+1:i+1], people[person[1]:i+1])
    //	people[person[1]] = person
    //}

    return queue
}
