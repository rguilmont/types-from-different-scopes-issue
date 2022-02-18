package pkga

import (
	"github.com/rguilmont/types-from-different-scopes-issue/pkg/either"
	"github.com/rguilmont/types-from-different-scopes-issue/pkg/pkgb"
)

func ToEither(s string) either.Either[pkgb.Test, any] {
	return either.Right[pkgb.Test, any](pkgb.Test(s))
}
