package runtime

import "github.com/emicklei/go-restful"

var container = restful.NewContainer()

func Container() *restful.Container {
    return container
}