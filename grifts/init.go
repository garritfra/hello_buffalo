package grifts

import (
	"github.com/garritfra/hello_buffalo/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
