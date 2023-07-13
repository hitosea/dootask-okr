package command

import (
	"dootask-okr/app/core"
	"dootask-okr/app/utils/common"
	"dootask-okr/config"
	"dootask-okr/database"
	"dootask-okr/i18n"
	"dootask-okr/router"
	"dootask-okr/router/middleware"
	"dootask-okr/web"
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/gin-contrib/gzip"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "dootask",
	Short: "启动服务",
	PreRun: func(cmd *cobra.Command, args []string) {
		if config.CONF.System.Host == "" {
			config.CONF.System.Host = "0.0.0.0"
		}
		if config.CONF.System.Port == "" {
			config.CONF.System.Port = "5566"
		}
		if config.CONF.System.Cache == "" {
			config.CONF.System.Cache = common.RunDir("/.cache")
		}
		if config.CONF.System.Dsn == "" {
			config.CONF.System.Dsn = fmt.Sprintf("sqlite3://%s/%s", config.CONF.System.Cache, "database.db")
		}
		config.CONF.System.Start = time.Now().Format(common.YYYY_MM_DD_HH_MM_SS)
		//
		err := common.WriteFile(config.CONF.System.Cache+"/config.json", common.StructToJson(config.CONF.System))
		if err != nil {
			common.PrintError(fmt.Sprintf("配置文件写入失败: %s", err.Error()))
			os.Exit(1)
		}
		// 初始化db
		err = core.InitDB()
		if err != nil {
			common.PrintError(fmt.Sprintf("数据库加载失败: %s", err.Error()))
			os.Exit(1)
		}
		// 初始化数据库
		err = database.Init()
		if err != nil {
			common.PrintError(fmt.Sprintf("数据库初始化失败: %s", err.Error()))
			os.Exit(1)
		}
		//
		common.PrintSuccess("启动成功")
	},
	Run: func(cmd *cobra.Command, args []string) {
		t, err := template.New("index").Parse(string(web.IndexByte))
		if err != nil {
			common.PrintError(fmt.Sprintf("模板解析失败: %s", err.Error()))
			os.Exit(1)
		}
		if config.CONF.System.Mode == "debug" {
			gin.SetMode(gin.DebugMode)
		} else if config.CONF.System.Mode == "test" {
			gin.SetMode(gin.TestMode)
		} else {
			gin.SetMode(gin.ReleaseMode)
		}
		r := gin.Default()
		r.Use(middleware.OperationLog())
		r.Use(gzip.Gzip(gzip.DefaultCompression))
		r.Use(i18n.GinI18nLocalize())
		r.SetFuncMap(template.FuncMap{
			"Localize": ginI18n.GetMessage,
		})
		r.SetHTMLTemplate(t)
		r.Any("/*path", func(context *gin.Context) {
			if !router.InitMicroapp(context) {
				router.Init(context)
			}
		})
		_ = r.Run(fmt.Sprintf("%s:%s", config.CONF.System.Host, config.CONF.System.Port))
	},
}

func Execute() {
	godotenv.Load(".env")

	rootCommand.CompletionOptions.DisableDefaultCmd = true

	flags := rootCommand.Flags()
	system := &config.CONF.System
	jwt := &config.CONF.Jwt
	redis := &config.CONF.Redis

	flags.StringVar(&system.Host, "host", "", "主机名，默认：0.0.0.0")
	flags.StringVar(&system.Port, "port", "", "端口号，默认：5566")
	flags.StringVar(&system.Mode, "mode", "release", "运行模式，可选：debug|test|release")
	flags.StringVar(&system.Cache, "cache", "", "数据缓存目录，默认：{RunDir}/.cache")

	mysqlDsn := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DBNAME"),
	)
	if mysqlDsn != "mysql://:@tcp(:)/?charset=utf8mb4&parseTime=True&loc=Local" {
		flags.StringVar(&system.Dsn, "dsn", mysqlDsn, "数据来源名称，如：sqlite://{CacheDir}/database.db")
	} else {
		flags.StringVar(&system.Dsn, "dsn", "", "数据来源名称，如：sqlite://{CacheDir}/database.db")
	}

	flags.StringVar(&jwt.SecretKey, "secret_key", "base64:ONdadQs1W4pY3h3dzr1jUSPrqLdsJQ9tCBZnb7HIDtk=", "jwt密钥")
	flags.StringVar(&redis.RedisUrl, "redis_url", "redis://localhost:56379", "RedisUrl")

	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
