package modloader

import (
	"github.com/tsaikd/gogstash/config"
	"github.com/tsaikd/gogstash/filter/addfield"
	"github.com/tsaikd/gogstash/filter/date"
	"github.com/tsaikd/gogstash/filter/geoip2"
	"github.com/tsaikd/gogstash/filter/gonx"
	"github.com/tsaikd/gogstash/filter/grok"
	"github.com/tsaikd/gogstash/filter/json"
	"github.com/tsaikd/gogstash/filter/ratelimit"
	"github.com/tsaikd/gogstash/filter/removefield"
	"github.com/tsaikd/gogstash/filter/typeconv"
	"github.com/tsaikd/gogstash/input/dockerlog"
	"github.com/tsaikd/gogstash/input/dockerstats"
	"github.com/tsaikd/gogstash/input/exec"
	"github.com/tsaikd/gogstash/input/file"
	"github.com/tsaikd/gogstash/input/http"
	"github.com/tsaikd/gogstash/input/httplisten"
	"github.com/tsaikd/gogstash/input/redis"
	"github.com/tsaikd/gogstash/input/socket"
	"github.com/tsaikd/gogstash/output/amqp"
	"github.com/tsaikd/gogstash/output/elastic"
	"github.com/tsaikd/gogstash/output/email"
	"github.com/tsaikd/gogstash/output/prometheus"
	"github.com/tsaikd/gogstash/output/redis"
	"github.com/tsaikd/gogstash/output/report"
	"github.com/tsaikd/gogstash/output/stdout"
)

func init() {
	config.RegistInputHandler(inputdockerlog.ModuleName, inputdockerlog.InitHandler)
	config.RegistInputHandler(inputdockerstats.ModuleName, inputdockerstats.InitHandler)
	config.RegistInputHandler(inputexec.ModuleName, inputexec.InitHandler)
	config.RegistInputHandler(inputfile.ModuleName, inputfile.InitHandler)
	config.RegistInputHandler(inputhttp.ModuleName, inputhttp.InitHandler)
	config.RegistInputHandler(inputhttplisten.ModuleName, inputhttplisten.InitHandler)
	config.RegistInputHandler(inputredis.ModuleName, inputredis.InitHandler)
	config.RegistInputHandler(inputsocket.ModuleName, inputsocket.InitHandler)

	config.RegistFilterHandler(filteraddfield.ModuleName, filteraddfield.InitHandler)
	config.RegistFilterHandler(filterdate.ModuleName, filterdate.InitHandler)
	config.RegistFilterHandler(filtergeoip2.ModuleName, filtergeoip2.InitHandler)
	config.RegistFilterHandler(filtergonx.ModuleName, filtergonx.InitHandler)
	config.RegistFilterHandler(filterjson.ModuleName, filterjson.InitHandler)
	config.RegistFilterHandler(filterratelimit.ModuleName, filterratelimit.InitHandler)
	config.RegistFilterHandler(filterremovefield.ModuleName, filterremovefield.InitHandler)
	config.RegistFilterHandler(filtertypeconv.ModuleName, filtertypeconv.InitHandler)
	config.RegistFilterHandler(filtertgrok.ModuleName, filtergrok.InitHandler)

	config.RegistOutputHandler(outputamqp.ModuleName, outputamqp.InitHandler)
	config.RegistOutputHandler(outputelastic.ModuleName, outputelastic.InitHandler)
	config.RegistOutputHandler(outputemail.ModuleName, outputemail.InitHandler)
	config.RegistOutputHandler(outputprometheus.ModuleName, outputprometheus.InitHandler)
	config.RegistOutputHandler(outputredis.ModuleName, outputredis.InitHandler)
	config.RegistOutputHandler(outputreport.ModuleName, outputreport.InitHandler)
	config.RegistOutputHandler(outputstdout.ModuleName, outputstdout.InitHandler)
}
