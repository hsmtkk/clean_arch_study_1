package memo

import "github.com/labstack/echo"

type MemoController interface {
	GetMemos(ctx echo.Context) error
}

type memoControllerImpl struct{}

func NewMemoController() MemoController {
	return &memoControllerImpl{}
}

func (memoCtrl *memoControllerImpl) GetMemos(ctx echo.Context) error {
	return nil
}
