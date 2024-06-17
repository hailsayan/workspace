package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "this is golrang shampoo tutorial go go"

	fmt.Println("EqualFold", strings.Contains(myString, "go1"))
	fmt.Println("EqualFold", strings.ContainsAny(myString, "go1"))
	fmt.Println("EqualFold", strings.Count(myString, "go"))
	fmt.Println(strings.Cut(myString, "go1"))

	myStringArray := strings.Split(myString, " ")

	println("Word Count: ", len(myStringArray))

	for _, item := range myStringArray {
		println(item)
	}

	myStringArray2 := strings.Join(myStringArray, ",")

	println("Join", myStringArray2)

	println("Repeat", strings.Repeat("paris ", 10))

	println("Replace", strings.Replace(myString, "golang", "go", 3))
	println("ReplaceAll", strings.ReplaceAll(myString, "go", "golang"))

	println("Compare", strings.Compare("golang", "golang"))
	println("Compare", strings.Compare("Golang", "golang"))
	println("Compare", strings.Compare("Golang", "GOLANG"))

	println("EqualFold", strings.EqualFold("Golang", "GOLANG"))
	println("EqualFold", strings.EqualFold("Golang", "golang"))


	println("HasPrefix", strings.HasPrefix("Paris", "Pa"))
	println("HasPrefix", strings.HasPrefix("Paris", "PA"))

	println("HasSuffix", strings.HasSuffix("Paris", "is"))
	println("HasSuffix", strings.HasSuffix("Paris", "s"))

	println("Index an", strings.Index("Paris", "is"))
	println("Index r", strings.Index("Paris", "a"))


	println("ToLower", strings.ToLower("PaRis"))
	println("ToUpper", strings.ToUpper("PaRis"))
	println("Title", strings.Title("Paris is a country"))

	println("Trim", strings.Trim("  Paris is a country   ", " "))
	println("TrimLeft", strings.TrimLeft("!!Paris is a country!!", "!"))
	println("TrimRight", strings.TrimRight("!!Paris is a country!!", "!"))
	println("Trim", strings.Trim("!!Paris is a country!!", "!"))





}
