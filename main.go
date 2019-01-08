package main

import (
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {
	app := cli.NewApp()
	app.Name = "split"
	app.Usage = "file split"
	app.Author = "Zhang Jian Xin"
	app.Version = "v0.1";

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Value: "",
			Usage: "filename",
		},
		cli.IntFlag{
			Name:  "line, l",
			Value: 0,
			Usage: "split lines",
		},
		cli.Int64Flag{
			Name:  "buff, b",
			Value: 0,
			Usage: "split buffer",
		},
	}

	app.Action = func(c *cli.Context) error {

		file := c.String("file")
		line := c.Int("line")
		bufer := c.Int64("buff")
		if file != ""&& line!=0 {
			SplitFileByLine(file, line)
			return nil
		}
		if file != ""&& bufer!=0 {
			const  int64 = 1 << 20 // 1mb
			SplitFileByBuffer(file, int64*bufer)
			return nil
		}
		err := cli.ShowAppHelp(c)  //这个是打印app的help界面
		if err !=nil{
			fmt.Println("")
		}
		return nil
	}
	app.Run(os.Args)

}
func countFileLine(name string) (count int64, err error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return
	}
	count = 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			count++
		}
	}
	return
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
