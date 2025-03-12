/*
Copyright © 2024 weak_ptr <weak_ptr@163.com>
This file is apart of the project media-vault. All rights reserved.
*/
package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"

	"github.com/nnnewb/media-vault/internal/api"
	"github.com/nnnewb/media-vault/internal/bootstrap"
	"github.com/nnnewb/media-vault/internal/logging"
	"github.com/nnnewb/media-vault/internal/service"
)

var serveOptions struct {
	ListenAddr           string
	DatabaseDSN          string
	FFMPEGPath           string
	DataRoot             string
	AnimeOfflineDatabase string
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "启动服务",
	PreRun: func(cmd *cobra.Command, args []string) {
		err := logging.Setup(
			logging.WithZapEncoding("console"),
			logging.WithZapOutputPath("media-vault.log"),
			logging.WithZapEncoderConfigEncodeLevel(zapcore.CapitalColorLevelEncoder),
			logging.WithZapLevel(zapcore.DebugLevel),
			logging.WithZapDevelopment(true),
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

		err = bootstrap.BootstrapDataFolder(serveOptions.DataRoot)
		if err != nil {
			logging.GetLogger().Panic("bootstrap data folder failed", zap.Error(err))
			return
		}

		if serveOptions.AnimeOfflineDatabase != "" {
			logging.GetLogger().Info("import anime offline database ...", zap.String("path", serveOptions.AnimeOfflineDatabase))
			err = bootstrap.BootstrapAnimeOfflineDatabase(db, serveOptions.AnimeOfflineDatabase)
			if err != nil {
				logging.GetLogger().Panic("bootstrap anime offline database failed", zap.Error(err))
				return
			}
		}

		// setup data folder
		err = os.MkdirAll(serveOptions.DataRoot, 0o755)
		if err != nil {
			logging.GetLogger().Panic("setup data folder failed", zap.Error(err))
			return
		}

		// setup service
		taskService := service.NewTaskService(db)
		inferService := service.NewMediaInfer()
		ffmpegService := service.NewFFMPEGService(db, serveOptions.FFMPEGPath)
		mediaService := service.NewMediaService(db, serveOptions.DataRoot, inferService, ffmpegService, taskService)
		animeService := service.NewAnimeService(db)
		pathService := service.NewPathService()

		// setup controller
		mediaControllerV1 := api.NewMediaControllerV1(mediaService)
		pathControllerV1 := api.NewPathControllerV1(pathService)
		taskController := api.NewTaskControllerV1(taskService)
		animeController := api.NewAnimeControllerV1(animeService)

		// create gin app
		app := gin.New()

		// setup middlewares
		app.Use(api.AccessLog(logging.GetLogger()))

		// setup routes
		router := app.Group("/api")
		mediaControllerV1.RegisterRoutes(router)
		pathControllerV1.RegisterRoutes(router)
		taskController.RegisterRoutes(router)
		animeController.RegisterRoutes(router)

		// start serving
		logging.GetLogger().Info("start serving", zap.String("listen", serveOptions.ListenAddr))
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
	serveCmd.Flags().StringVarP(&serveOptions.ListenAddr, "listen", "l", "127.0.0.1:39876", "监听地址")
	serveCmd.Flags().StringVar(&serveOptions.FFMPEGPath, "ffmpeg", "ffmpeg", "ffmpeg 命令路径")
	serveCmd.Flags().StringVar(&serveOptions.DataRoot, "data-root", "./data", "数据根目录，保存封面、预览等数据")
	serveCmd.Flags().StringVar(&serveOptions.AnimeOfflineDatabase, "anime-offline-database", "", "导入动漫离线数据库 manami-project/anime-offline-database")
}
