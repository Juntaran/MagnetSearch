/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/10/16 23:00
  */

package main

import (
	"os"

	"github.com/urfave/cli"

	"magnetSearch/cmd"
)

func main() {
	app := cli.NewApp()
	app.Name = "magent searcher"
	app.Author = "Juntaran"
	app.Email = "jacinthmail@gmail.com"
	app.Version = "beta"
	app.Usage = "Download what you need :)"
	app.Commands = []cli.Command{
		cmd.Search,
	}

	app.Flags = append(app.Flags, cmd.Search.Flags...)

	app.Run(os.Args)
}
