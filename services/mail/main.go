package main

import "github.com/kataras/iris"

type callback string

func (c callback) Callback(ctx iris.Context) {
    ctx.JSON(iris.Map{"message": "Hello MAIL"})
}

// exported as symbol named "CmfiveCallback"
var CmfiveCallback callback