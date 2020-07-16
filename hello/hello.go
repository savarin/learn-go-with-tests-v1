package main

import "fmt"

const prefixEnglish = "Hello, "
const prefixSpanish = "Hola, "
const prefixFrench = "Bonjour, "

func greeting(language string) (prefix string) {
    switch language {
    case "Spanish":
        prefix = prefixSpanish
    case "French":
        prefix = prefixFrench
    default:
        prefix = prefixEnglish
    }
    return
}

func Hello(name, language string) string {
    if name == "" {
        name = "World"
    }

    return greeting(language) + name
}

func main() {
    fmt.Println(Hello("", "English"))
}
