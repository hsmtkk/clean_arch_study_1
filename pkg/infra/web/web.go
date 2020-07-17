package web

import (
	"github.com/hsmtkk/clean_arch_study_1/pkg/interface/memo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	ctrl := memo.NewMemoController()
	e.GET("/", ctrl.GetMemos)
	e.Logger.Fatal(e.Start(":8000"))
}
