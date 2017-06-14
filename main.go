package main

import (
	"dispatcher/dispatch"
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()

	dispatcher := dispatch.NewDispatcher(200)
	dispatcher.Run()

	app.OnErrorCode(iris.StatusNotFound, func(ctx context.Context) {
		ctx.WriteString("not found")
	})
	app.Get("/start", func(ctx context.Context) {
		for i := 1; i < 1000; i++ {
			job := dispatch.Job{Title: "title" + fmt.Sprintln(i), Description: "description" + fmt.Sprintln(i)}
			dispatch.JobQueue <- job
		}
		ctx.WriteString("input success")
	})

	app.Run(iris.Addr(":8080"))
}
