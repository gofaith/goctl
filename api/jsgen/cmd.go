package jsgen

import (
	"errors"

	"github.com/gofaith/goctl/api/parser"

	"github.com/urfave/cli"
)

func JsCommand(c *cli.Context) error {
	apiFile := c.String("api")
	if apiFile == "" {
		return errors.New("missing -api")
	}
	dir := c.String("dir")
	if dir == "" {
		return errors.New("missing -dir")
	}

	return JsGen(apiFile, dir)
}

func JsGen(apiFile, dir string) error {
	p, e := parser.NewParser(apiFile)
	if e != nil {
		return e
	}

	api, e := p.Parse()
	if e != nil {
		return e
	}

	e = genBase(dir, api)
	if e != nil {
		return e
	}
	e = genApi(dir, api)
	if e != nil {
		return e
	}
	return nil
}
