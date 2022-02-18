# Issue Report

with golang at least 1.18.

Replicate by build and running files under `pkg/cmd`. You should get :

```
panic: interface conversion: either.Either[go.shape.string_0,go.shape.interface {}_1] is either.right[github.com/rguilmont/types-from-different-scopes-issue/pkg/pkga.Test,interface {}], not either.right[github.com/rguilmont/types-from-different-scopes-issue/pkg/pkga.Test,interface {}] (types from different scopes)

goroutine 1 [running]:
github.com/rguilmont/types-from-different-scopes-issue/pkg/either.Fold[...]({0x4b2c00, 0xc00009e210}, 0x49c850, 0x49c858)
        /home/romain/Documents/git/github/replicate-issue/types-from-different-scopes-issue/pkg/either/either.go:42 +0xc5
main.main()
        /home/romain/Documents/git/github/replicate-issue/types-from-different-scopes-issue/pkg/cmd/main.go:13 +0x65
```

Checkout the branch `bypassed` to see the repo changes, where the issue is bypassed by putting the type in another package, and you'll get the following which was expected : 

```
It works! test
```

