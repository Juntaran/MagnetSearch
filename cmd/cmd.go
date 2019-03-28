/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/10/16 22:46
  */

package cmd

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/Juntaran/MagnetSearch/crawler"
)

var Search = cli.Command {
	Name:			"search",
	Usage: 			"MagnetSearch search sth",
	Description: 	"Search what you need :)",
	Action: 		crawler.Search,
	Flags:          []cli.Flag{
		stringFlag("keyword, k", "sakuramomo", "search keyword"),
		intFlag("number, n", 10, "most return numbers"),
		intFlag("sort, s", 1, "sort way: 1-time, 2-hot, 3-size"),
		boolFlag("detail, d", "show detail information"),
	},
}

func stringFlag(name, value, usage string) cli.StringFlag {
	return cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func boolFlag(name, usage string) cli.BoolFlag {
	return cli.BoolFlag{
		Name:  name,
		Usage: usage,
	}
}

func intFlag(name string, value int, usage string) cli.IntFlag {
	return cli.IntFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func Lang(a string)  {
	fmt.Println("lang", a)
}