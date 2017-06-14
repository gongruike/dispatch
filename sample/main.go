package main

import (
	"dispatch"
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
		for i := 1; i < 200; i++ {
			displayJob := &dispatch.DisplayJob{Title: "title" + fmt.Sprintln(i)}
			outputJob := dispatch.OutputJob{Name: "Name" + fmt.Sprintln(i)}
			//
			dispatcher.JobQueue <- displayJob
			dispatcher.JobQueue <- outputJob
		}
		ctx.WriteString("input success")
	})

	app.Run(iris.Addr(":8080"))
}
