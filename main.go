package main

import (
	"IM/configs"
	"IM/routers"
	"fmt"
)

func main() {
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", configs.Conf.API.Port)); err != nil {
		panic(fmt.Sprintf("run server failed, err:%v\n", err))
	}
}
