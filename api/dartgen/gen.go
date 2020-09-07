package dartgen

import (
	"errors"
	"strings"

	"github.com/gofaith/go-zero/core/logx"
	"github.com/gofaith/goctl/api/parser"
	"github.com/urfave/cli"
)

func DartCommand(c *cli.Context) error {
	apiFile := c.String("api")
	dir := c.String("dir")
	if len(apiFile) == 0 {
		return errors.New("missing -api")
	}
	if len(dir) == 0 {
		return errors.New("missing -dir")
	}

	p, err := parser.NewParser(apiFile)
	if err != nil {
		return err
	}
	api, err := p.Parse()
	if err != nil {
		return err
	}

	if !strings.HasSuffix(dir, "/") {
		dir = dir + "/"
	}
	api.Info.Title = strings.Replace(apiFile, ".api", "", -1)
	logx.Must(genData(dir+"data/", api))
	logx.Must(genApi(dir+"api/", api))
	logx.Must(genVars(dir + "vars/"))
	return nil
}
