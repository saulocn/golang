package app

import (
	mcache "github.com/go-macaron/cache"
	"github.com/go-macaron/gzip"
	"github.com/go-macaron/i18n"
	"github.com/go-macaron/jade"
	"github.com/go-macaron/session"
	"github.com/go-macaron/toolbox"
	"github.com/saulocn/testemercurius/conf"
	"github.com/saulocn/testemercurius/handler"
	"github.com/saulocn/testemercurius/lib/cache"
	"github.com/saulocn/testemercurius/lib/contx"
	"github.com/saulocn/testemercurius/lib/template"
	"github.com/saulocn/testemercurius/lib/cors"
	"gopkg.in/macaron.v1"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//SetupMiddlewares configures the middlewares using in each web request
func SetupMiddlewares(app *macaron.Macaron) {
	app.Use(macaron.Logger())
	app.Use(macaron.Recovery())
	app.Use(gzip.Gziper())
	app.Use(toolbox.Toolboxer(app, toolbox.Options{
		HealthCheckers: []toolbox.HealthChecker{
			new(handler.AppChecker),
		},
	}))
	app.Use(macaron.Static("public"))
	app.Use(i18n.I18n(i18n.Options{
		Directory: "locale",
		Langs:     []string{"pt-BR", "en-US"},
		Names:     []string{"Português do Brasil", "American English"},
	}))
	app.Use(jade.Renderer(jade.Options{
		Directory: "public/templates",
		Funcs:     template.FuncMaps(),
	}))
	app.Use(macaron.Renderer(macaron.RenderOptions{
    	Directory: "public/templates",
    	Funcs:     template.FuncMaps(),
    }))
	//Cache in memory
	app.Use(mcache.Cacher(
		cache.Option(conf.Cfg.Section("").Key("cache_adapter").Value()),
	))
	/*
	Redis Cache
	Add this lib to import session: _ "github.com/go-macaron/cache/redis"
	Later replaces the cache in memory instructions for the lines below
	optCache := mcache.Options{
			Adapter:       conf.Cfg.Section("").Key("cache_adapter").Value(),
			AdapterConfig: conf.Cfg.Section("").Key("cache_adapter_config").Value(),
		}
	app.Use(mcache.Cacher(optCache))
	*/
	app.Use(session.Sessioner())
	app.Use(contx.Contexter())
	app.Use(cors.Cors())
}

//SetupRoutes defines the routes the Web Application will respond
func SetupRoutes(app *macaron.Macaron) {
	app.Get("", func() string {
		return "Mercurius Works!"
	})

	//HealthChecker
	app.Get("/health", handler.HealthCheck)

	//Prometheus metrics
	app.Get("/metrics", promhttp.Handler())
	
	/*
		//An example to test DB connection
		app.Get("", func() string {
			db, err := conf.GetDB()
			if err != nil {
				return err.Error()
			}
			err = db.Ping()
			if err != nil {
				return err.Error()
			}
			col, err := conf.GetMongoCollection("teste")
			if err != nil {
				return err.Error()
			}
			defer col.Database.Session.Close()
			teste := Teste{Status: "OK"}
			err = col.Insert(teste)
			return "Mercurius Works!"
		})

		//Include this struct after import session
		type Teste struct {
			Status string
		}
	*/
}
