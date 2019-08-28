package main

import (
	"flag"
	"os"

	mp "github.com/mackerelio/go-mackerel-plugin"
	"github.com/rluisr/mysqlrouter-go"
)

var (
	url      = os.Getenv("MYSQLROUTER_URL")
	user     = os.Getenv("MYSQLROUTER_USER")
	pass     = os.Getenv("MYSQLROUTER_PASS")
	graphdef map[string]mp.Graphs
)

// MRPlugin is the prefix of struct of graph
type MRPlugin struct {
	Prefix   string
	MRClient *mysqlrouter.Client
}

func (mr MRPlugin) MetricKeyPrefix() string {
	if mr.Prefix == "" {
		mr.Prefix = "mysqlrouter"
	}
	return mr.Prefix
}

func (mr MRPlugin) GraphDefinition() map[string]mp.Graphs {
	return graphdef
}

func (mr MRPlugin) FetchMetrics() (map[string]float64, error) {
	return map[string]float64{"active_connections": float64(1)}, nil
}

func (mr MRPlugin) Prepare() {
	metadata, err := mr.MRClient.GetAllMetadata()
	if err != nil {
		panic(err)
	}

	g := map[string]mp.Graphs{}

	for _, m := range metadata {
		g["mysqlrouter."+m.Name] = mp.Graphs{
			Label: "Connections of route",
			Unit:  "integer",
			Metrics: []mp.Metrics{
				{Name: "active_connections", Label: "active_connections", Diff: false},
				//{Name: "total_connections", Label: "total_connections", Diff: true},
			},
		}
	}

	graphdef = g
}

func main() {
	if url == "" || user == "" || pass == "" {
		panic("The environment missing.\n" +
			"MYSQLROUTER_URL, MYSQLROUTER_USER and MYSQLROUTER_PASS is required.")
	}

	mrr, err := mysqlrouter.New(url, user, pass)
	if err != nil {
		panic(err)
	}

	prefix := flag.String("metric-key-prefix", "", "Metric key prefix")
	temp := flag.String("tempfile", "", "Temp file name")
	flag.Parse()

	mr := MRPlugin{
		Prefix:   *prefix,
		MRClient: mrr,
	}

	mr.Prepare()

	plugin := mp.NewMackerelPlugin(mr)
	plugin.Tempfile = *temp
	plugin.Run()
}
