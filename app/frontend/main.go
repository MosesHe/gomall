// Code generated by hertz generator.

package main

import (
	"context"
	"os"
	"time"

	"github.com/MosesHe/gomall/app/frontend/biz/router"
	"github.com/MosesHe/gomall/app/frontend/conf"
	"github.com/MosesHe/gomall/app/frontend/infra/rpc"
	"github.com/MosesHe/gomall/app/frontend/middleware"
	frontendUtils "github.com/MosesHe/gomall/app/frontend/utils"
	"github.com/MosesHe/gomall/common/mtl"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	prometheus "github.com/hertz-contrib/monitor-prometheus"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName  = frontendUtils.ServiceName
	MetricsPort  = conf.GetConf().Hertz.MetricsPort
	RegistryAddr = conf.GetConf().Hertz.RegistryAddr
)

func main() {
	_ = godotenv.Load()
	consul, registryInfo := mtl.InitMetric(ServiceName, MetricsPort, RegistryAddr)
	defer consul.Deregister(registryInfo)
	// init dal
	// dal.Init()
	rpc.Init()
	address := conf.GetConf().Hertz.Address
	h := server.New(server.WithHostPorts(address),
		server.WithTracer(prometheus.NewServerTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))),
	)

	registerMiddleware(h)

	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	router.GeneratedRegister(h)
	h.LoadHTMLGlob("template/*")
	h.Static("/static", "./")

	h.GET("/signin", func(c context.Context, ctx *app.RequestContext) {
		data := utils.H{
			"title": "Sign In",
			"Next":  ctx.Query("next"),
		}
		ctx.HTML(consts.StatusOK, "signin", data)
	})

	h.GET("/signup", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "signup", utils.H{"title": "Sign Up"})
	})

	h.GET("/about", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "about", utils.H{"title": "About"})
	})

	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	store, _ := redis.NewStore(10, "tcp", conf.GetConf().Redis.Address, "", []byte(os.Getenv("SESSION_SECRET")))
	h.Use(sessions.New("gomall", store))

	// log
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	hlog.SetOutput(asyncWriter)
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		asyncWriter.Sync()
	})

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())

	middleware.Register(h)
}
