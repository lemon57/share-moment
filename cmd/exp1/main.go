package main

import (
	"context"
	"fmt"
	"strings"
)

type ctxKey string

const (
	favoriteColorKey ctxKey = "favorite-color"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, favoriteColorKey, "blue")

	anyValue := ctx.Value(favoriteColorKey)
	stringValue, ok := anyValue.(string)
	if !ok {
		fmt.Println(anyValue, " is not a string")
	} else {
		fmt.Println(strings.HasPrefix(stringValue, "b"))
	}

	intValue, ok := anyValue.(int)
	if !ok {
		fmt.Println(anyValue, " is not an int")
	} else {
		fmt.Println(intValue + 4)
	}
}
