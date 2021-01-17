/*
Copyright © 2021 Ambor <saltbo@foxmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/saltbo/gopkg/ginutil"
	"github.com/saltbo/gopkg/jwtutil"
	"github.com/spf13/cobra"

	"github.com/zplat-core/apiserver/api"
)

// serverCmd represents the v1 command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1 called")
		serverRun()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

// 这里是主服务，包含注册、登录等；也包含用户信息，提供两种鉴权方式，通过后即可拿到用户信息
func serverRun() {
	ge := gin.Default()
	ginutil.SetupPing(ge)
	ginutil.SetupSwagger(ge)
	jwtutil.Init("123")

	//conf := config.ParseServer()
	//if viper.ConfigFileUsed() != "" {
	//	gormutil.Init(conf.Database, true)
	//	gormutil.AutoMigrate(model.Tables())
	//}
	//
	//viper.WatchConfig()
	//viper.OnConfigChange(func(in fsnotify.Event) {
	//	gormutil.Init(conf.Database, true)
	//	gormutil.AutoMigrate(model.Tables())
	//})

	api.SetupServerRoutes(ge)
	ginutil.Startup(ge, ":8218")
}
