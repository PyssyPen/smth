package main

/*
nums := make([]int, 0, 10)
// начиная с 0; пока i меньше 10; инкрементим i после каждого шага
for i := 0; i < 10; i++ {
    nums = append(nums, i)
}
fmt.Println(nums) // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]


i := 0
nums := make([]int, 0, 10)
for i < 10 {
    nums = append(nums, i)
    i++
}
fmt.Println(nums) // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]


i := 0
nums := make([]int, 0, 10)
for {
    if i == 10 {
        break
    }
    nums = append(nums, i)
    i++
}
fmt.Println(nums) // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]


nums := make([]int, 0, 10)
// добавляем только четные числа в слайс
for i := 0; i < 10; i++ {
    if i % 2 != 0 {
        continue
    }
    nums = append(nums, i)
}
fmt.Println(nums) // [0 2 4 6 8]


names := []string{"John", "Harold", "Vince"}
// i — это индекс, name — это значение на текущем шаге цикла
for i, name := range names {
    fmt.Println("Hello ", name, " at index ", i)
}
// => Hello  John  at index  0
// => Hello  Harold  at index  1
// => Hello  Vince  at index  2


for i := range names {
    fmt.Println("index = ", i)
}
// => index =  0
// => index =  1
// => index =  2


for _, name := range names {
    fmt.Println("Hello ", name)
}
// => Hello  John
// => Hello  Harold
// => Hello  Vince


for _,_ := range names {
    fmt.Println("Nothing")
}
# command-line-arguments
./main.go:21:14: no new variables on left side of :=
*/
