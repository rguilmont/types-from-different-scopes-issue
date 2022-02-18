package either

type Either[A, B any] interface {
	IsLeft() bool
	IsRight() bool

	Left() (bool, *B)
	Right() (bool, *A)
}

type left[A, B any] struct {
	Value B
}

func (left[A, B]) IsLeft() bool        { return true }
func (left[A, B]) IsRight() bool       { return false }
func (l left[A, B]) Left() (bool, *B)  { return true, &l.Value }
func (l left[A, B]) Right() (bool, *A) { return false, nil }

type right[A, B any] struct {
	Value A
}

func (right[A, B]) IsLeft() bool        { return false }
func (right[A, B]) IsRight() bool       { return true }
func (r right[A, B]) Left() (bool, *B)  { return false, nil }
func (r right[A, B]) Right() (bool, *A) { return true, &r.Value }

func Left[A, B any](l B) Either[A, B] {
	return left[A, B]{l}
}

func Right[A, B any](r A) Either[A, B] {
	return right[A, B]{r}
}

func Fold[A, B, C any](e Either[A, B], fa func(A) C, fb func(B) C) C {
	if e.IsRight() {
		return fa(e.(right[A, B]).Value)
	}
	return fb(e.(left[A, B]).Value)
}
