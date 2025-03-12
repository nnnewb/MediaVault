/*
Copyright © 2024 weak_ptr <weak_ptr@163.com>
This file is apart of the project media-vault. All rights reserved.
*/
package cmd

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"

	"gitee.com/uniqptr/media-vault.git/internal/bootstrap"
	"gitee.com/uniqptr/media-vault.git/internal/logging"
	"gitee.com/uniqptr/media-vault.git/internal/routes"
)

var serveOptions struct {
	ListenAddr  string
	DatabaseDSN string
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "启动服务",
	PreRun: func(cmd *cobra.Command, args []string) {
		err := logging.Setup(
			logging.WithZapOutputPath("media-vault.log"),
			logging.WithZapEncoderConfigEncodeLevel(zapcore.CapitalColorLevelEncoder),
			logging.WithZapLevel(zapcore.DebugLevel),
		)
		if err != nil {
			log.Panicf("setup logging facility failed, error %+v", err)
			return
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// open database
		db, err := gorm.Open(sqlite.Open(serveOptions.DatabaseDSN), &gorm.Config{})
		if err != nil {
			logging.GetLogger().Panic("open database failed", zap.String("dsn", serveOptions.DatabaseDSN), zap.Error(err))
			return
		}

		// bootstrap
		err = bootstrap.BootstrapDatabase(db)
		if err != nil {
			logging.GetLogger().Panic("bootstrap database failed", zap.Error(err))
			return
		}

		// start serving
		app := gin.New()
		routes.RegisterAppRoutes(app)
		if err := http.ListenAndServe(serveOptions.ListenAddr, app); err != nil {
			logging.GetLogger().Error("listen and serve failed", zap.Error(err))
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serveCmd.Flags().StringVar(&serveOptions.DatabaseDSN, "database", "media-vault.db", "数据库文件路径")
	serveCmd.Flags().StringVarP(&serveOptions.ListenAddr, "listen", "l", ":39876", "监听地址")
}
