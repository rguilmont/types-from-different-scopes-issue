package main

import (
	"fmt"

	"github.com/rguilmont/types-from-different-scopes-issue/pkg/either"
	"github.com/rguilmont/types-from-different-scopes-issue/pkg/pkga"
	"github.com/rguilmont/types-from-different-scopes-issue/pkg/pkgb"
)

func main() {
	a := pkga.ToEither("test")

	either.Fold(a,

		func(right pkgb.Test) bool {
			fmt.Println("It works!", right)
			return true
		},
		func(_ any) bool {
			fmt.Println("It Fails !")
			return false
		},
	)
}
