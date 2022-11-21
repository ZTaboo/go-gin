package main

import (
	"zm/config"
	"zm/lib/z"
	"zm/router"
)

func main() {
	r := router.Router()
	err := r.Run(config.Port)
	if err != nil {
		z.Log.Error(err)
	}
}
