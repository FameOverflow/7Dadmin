package global

import (
	"dst-admin-go/collect"
	"dst-admin-go/config"
	"net/http/httputil"
	"net/url"
)

type Route struct {
	Proxy *httputil.ReverseProxy
	Url   *url.URL
}

var RoutingTable = make(map[string]*Route)

var Config *config.Config

const ClusterToken = "pds-g^KU_qE7e8rv1^VVrVXd/01kBDicd7UO5LeL+uYZH1+geZlrutzItvOaw="

var CollectMap = collect.NewCollectMap()

var Collect *collect.Collect
