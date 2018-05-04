package main

import (
	"fmt"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/asset"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/gl"
	"io/ioutil"
)

func main() {
	app.Main(func(a app.App) {
		var glctx gl.Context
		sz := size.Event{}
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				glctx, _ = e.DrawContext.(gl.Context)
			case size.Event:
				sz = e
			case paint.Event:
				if glctx == nil {
					continue
				}
				onDraw(glctx, sz)
				a.Publish()
			}
		}
	})
}
func onDraw(glctx gl.Context, sz size.Event) {
	glctx.ClearColor(0, 1, 1, 1)
	go pm()
	glctx.Clear(gl.COLOR_BUFFER_BIT)
}
func pm() {
	f, e := asset.Open("a.txt")
	if e != nil {
		fmt.Println(e)
		return
	}
	rp, e := ioutil.ReadAll(f)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Print(string(rp))
}
