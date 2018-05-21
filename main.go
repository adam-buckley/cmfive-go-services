package main

import (
	"github.com/kataras/iris"
	"io/ioutil"
	"log"
	"plugin"
)

func main() {
    app := iris.New()

    // Look at services directory and load each plugin
    files, err := ioutil.ReadDir("./services")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
    	fmt.Println(f.Name())

    	mod := "./services/" + f.Name() + "/" + f.Name() ".so"
    	plug, err := plugin.Open(mod)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		symCallback, err := plug.Lookup("CmfiveCallback")

		// There is an object here
		fmt.Printf("%+v\n", symCallback)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var callback CmfiveCallback
		callback, ok := symCallback.(CmfiveCallback)

		// But nothing found here
		fmt.Printf("%+v\n", callback)
		if !ok {
			fmt.Println("unexpected type from module symbol")
			os.Exit(1)
		}

		app.Get('/' + f.Name() + "/action", callback.Callback())
    }


    // Method:    GET
    // Resource:  http://localhost:8080
    app.Get("/", func(ctx iris.Context) {
        ctx.JSON(iris.Map{"message": "Hello World"})
    })

    // Method:    GET
    // Resource:  http://localhost:8080/user/42
    //
    // Need to use a custom regexp instead?
    // Easy,
    // just mark the parameter's type to 'string'
    // which accepts anything and make use of
    // its `regexp` macro function, i.e:
    // app.Get("/user/{id:string regexp(^[0-9]+$)}")
    // app.Get("/user/{id:long}", func(ctx iris.Context) {
    //     userID, _ := ctx.Params().GetInt64("id")
    //     ctx.Writef("User ID: %d", userID)
    // })

    // Start the server using a network address.
    app.Run(iris.Addr(":8080"))
}