/*
Say, you need to buy a sports kit,
a set of two different products:
	a pair of shoes
	a shirt.
You would want to buy a full sports kit of the same brand to match all the items.
If we try to turn this into the code, the abstract factory will help us create sets of products so that they would always match each other.
*/

package main

import "fmt"

func main() {
    adidasFactory, _ := getSportsFactory("adidas")
    nikeFactory, _ := getSportsFactory("nike")
    spunkFactory,_ := getSportsFactory("spunk")

    nikeShoe := nikeFactory.makeShoe()
    nikeShirt := nikeFactory.makeShirt()

    adidasShoe := adidasFactory.makeShoe()
    adidasShirt := adidasFactory.makeShirt()

    spunkShoe := spunkFactory.makeShoe()
    spunkShirt := spunkFactory.makeShirt()

    fmt.Println("\n----- SPUNK ------")
    printShoeDetails(spunkShoe)
    printShirtDetails(spunkShirt)

    fmt.Println("\n----- NIKE ------")
    printShoeDetails(nikeShoe)
    printShirtDetails(nikeShirt)

    fmt.Println("\n----- ADIDAS ------")
    printShoeDetails(adidasShoe)
    printShirtDetails(adidasShirt)
}

func printShoeDetails(s iShoe) {

    fmt.Println("\t --- SHOE ---")
    fmt.Printf("Logo: %s", s.getLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.getSize())
    fmt.Println()

}

func printShirtDetails(s iShirt) {
    fmt.Println("\t --- SHIRT ---")
    fmt.Printf("Logo: %s", s.getLogo())
    fmt.Println()
    fmt.Printf("Size: %d", s.getSize())
    fmt.Println()

}