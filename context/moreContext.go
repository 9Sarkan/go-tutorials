package main

import (
	"context"
	"fmt"
)

type aKey string

func searchContext(ctx context.Context, key aKey) {
	v := ctx.Value(key)
	if v != nil {
		fmt.Println("found key: ", v)
	} else {
		fmt.Println("Can not found key!")
	}
}

func main() {
	myKey := aKey("mySecret")
	cxt := context.WithValue(context.Background(), myKey, "mySecret")

	searchContext(cxt, myKey)
	searchContext(cxt, aKey("notKey"))
	emptyContext := context.TODO()
	searchContext(emptyContext, aKey("notKey"))
}
