package main

import (
	"github.com/pterm/pterm"
	"github.com/valyala/fasthttp"
	"time"
)

func serveWorld(name string) error {

	pterm.Println()

	pterm.Success.Println("------ ✨ Servidor iniciado correctamente en puerto 1000 ✨ ------")

	time.Sleep(time.Second * 1)

	pterm.Println()

	pterm.Info.Println("... 📚 Para cerrar el servidor pulsa Ctrl + C ...")

	handler := func(ctx *fasthttp.RequestCtx) {

		method := string(ctx.Method())

		if method == "GET" {

			ctx.Response.Header.Set("Name", name)
			ctx.Response.Header.Set("Content-Type", "application/zip")

			ctx.SendFile(name)

		} else {

			ctx.Write([]byte("Method " + method + " not allowed"))

		}

	}

	gzipHandler := fasthttp.CompressHandlerLevel(handler, fasthttp.CompressBestCompression)

	error := fasthttp.ListenAndServe(":1000", gzipHandler)

	if error != nil {

		return error

	}

	return nil

}
