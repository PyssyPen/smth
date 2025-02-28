package main

/*

nums := []int{}
fmt.Println(nums)


var nums = []int{1,2,3}
nums := []int{1,2,3}

nums := []int{1,2,3}
nums[2] // 3
nums[0] = 10 // [10, 2, 3]
// с помощью оператора : можно получить элементы от нижней до верхней границы
nums[1:3] // [2, 3]
nums[:2] // [10, 2]
nums[2:] // [3]

words := []string{"hello"}
words = append(words, "world") // ["hello", "world"]

// len = 5. Массив сразу будет заполнен 5-ю нулевыми значениями
nums := make([]int, 5, 5) // [0, 0, 0, 0, 0]
// len = 0, но cap = 5. Массив будет пустым, однако заполнение слайса через append будет эффективным,
потому что в памяти уже выделен массив нужной длины
nums := make([]int, 0, 5) // []


import (
    "fmt"
)
func main() {
    nums := []int{1, 2, 3, 4, 5} // вместимость 5

    modifySlice(nums)

    fmt.Println(nums) // [1 2 10 4 5]
}
func modifySlice(nums []int) {
    nums[2] = 10 // элемент изменится и в исходном слайсе
    nums = append(nums, 6) // элемент не добавится в исходный слайс, так как превысили изначальную
	вместимость и nums теперь ссылается на новый массив
    nums[3] = 15  // элемент НЕ изменится в исходном слайсе
}




Для проверки наличия элемента по определенному ключу можно применять выражение if:

var people = map[string]int{
    "Tom": 1,
    "Bob": 2,
    "Sam": 4,
    "Alice": 8,
}
if val, ok := people["Tom"]; ok{
    fmt.Println(val)
}




Для добавления элементов достаточно просто выполнить установку значения по новому ключу и
элемент с этим ключом будет добавлен в коллекцию:

var people = map[string]int{ "Tom": 1, "Bob": 2}
people["Kate"] = 128
fmt.Println(people)     // map[Tom:1  Bob:2  Kate:128]


Для удаления применяется встроенная функция delete(map, key), первым параметром которой
является отображение, а вторым - ключ, по которому надо удалить элемент.

var people = map[string]int{ "Tom": 1, "Bob": 2, "Sam": 8}
delete(people, "Bob")
fmt.Println(people)     // map[Tom:1  Sam:8]
*/
