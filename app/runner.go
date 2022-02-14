package app

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func Run(port int) {
	e := echo.New()
	Init(e.Group("api"))
	go func() {
		fmt.Println(e.Start(":" +
			strconv.Itoa(port)))
	}()
	GracefulShutdown(e)
}

func GracefulShutdown(e *echo.Echo) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = e.Shutdown(ctx)
}
