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

type iSportsFactory interface{
	makeShoe() iShoe
	makeShirt() iShirt
}

func getSportsFactory(brand string) (iSportsFactory, error) {
    if brand == "adidas" {
        return &adidas{}, nil
    }

    if brand == "nike" {
        return &nike{}, nil
    }
    if brand == "spunk" {
        return &spunk{}, nil
    }



    return nil, fmt.Errorf("wrong brand type passed")
}