package memo

import "github.com/hsmtkk/clean_arch_study_1/domain/memo"

type MemoUseCase interface {
	GetMemos([]memo.Memo, error)
}

type memoUseCaseImpl struct{}

func NewMemoUseCase() MemoUseCase {
	return &memoUseCaseImpl{}
}

func (memoUseCase *memoUseCaseImpl) GetMemos([]memo.Memo, error) {
	return nil, nil
}
