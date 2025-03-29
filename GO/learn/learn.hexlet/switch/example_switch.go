package main

/*
x := 10

switch x {
    default: // default всегда выполняется последним независимо от расположения в конструкции
        fmt.Println("default case")
    case 10:
        fmt.Println("case 10")
}
Output:

case 10

//////////////////////////////////////////////////////

x := 10

switch { // выражение отсутствует. Для компилятора выглядит как: switch true
    default:
        fmt.Println("default case")
    case x == 10:
        fmt.Println("equal 10 case")
        fallthrough
    case x <= 10:
        fmt.Println("less or equal 10 case")
}
Output:

equal 10 case
less or equal 10 case
*/
