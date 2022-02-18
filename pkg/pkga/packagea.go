package pkga

import "github.com/rguilmont/types-from-different-scopes-issue/pkg/either"

type Test string

func ToEither(s string) either.Either[Test, any] {
	return either.Right[Test, any](Test(s))
}
