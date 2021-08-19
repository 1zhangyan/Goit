package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Git"
	app.Usage = "Basic git Command Utils"
	app.Version = "0.0.1"
	app.Author = "Zhang Yan"
	app.Email = "zhangyan1dev@163.com"

	app.Commands = []cli.Command{
		{
			Name: "add",
			Usage: "add file to tmp",
			Action: func(c *cli.Context) error {
				fmt.Println("Add file to index filename is " ,c.Args().First())
				return nil
			},
		},
		{},
		{},
	}

	app.Run(os.Args)
}
