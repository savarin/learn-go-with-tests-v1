package sum

func Sum(numbers []int) int {
    result := 0

    for _, item := range numbers {
        result += item
    }
    return result
}

func SumAll(numbers ...[]int) (result []int) {
    result = make([]int, 0)

    for _, item := range numbers {
        result = append(result, Sum(item))
    }
    return
}

func SumAllTails(numbers ...[]int) (result []int) {
    result = make([]int, 0)

    for _, item := range numbers {
        if len(item) == 0 {
            result = append(result, 0)
            continue
        }
        result = append(result, Sum(item[1:]))
    }
    return
}
