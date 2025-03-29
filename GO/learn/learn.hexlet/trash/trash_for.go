package main

/*
You cannot convert a string to a function. You may try using some kind of mapping, e.g.:

checkers := map[string]func() {
    "CheckIPReachability": checkIPReachability,
}

checkers[name]()

If you need to pass the arguments, the function would be like:

func checkIPReachability(args []interface{}) {
    // convert args and check the IPs
}

checkers := map[string]func([]interface{}) {
    "CheckIPReachability": "checkIPReachability",
}

If you want to allow the users of your framework to add their own checker, you can export that as
a function RegisterChecker(string, func(interface{})) or something similar.

A simpler alternative is a switch-case statement:

switch funcname {
    case "CheckIPReachability":
        // convert arguments and call the function
    case "OtherCheck":
        // same here
}
*/
