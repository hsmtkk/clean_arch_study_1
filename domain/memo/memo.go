package memo

type Memo interface {
	ID() int
	Text() string
}

type memoImpl struct {
	id   int
	text string
}

func (m memoImpl) ID() int {
	return m.id
}

func (m memoImpl) Text() string {
	return m.text
}

func NewMemo() Memo {
	return &memoImpl{}
}
