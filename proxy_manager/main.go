package main

import (
	"context"
	"github.com/antchfx/htmlquery"
	"real_estate_scraping/infrastructure/redis_db"
	"strings"
)

// Free proxies:
// https://free-proxy-list.net/

// https://www.freeproxylists.net/
type FreeProxyListNetProxy struct {
	IP        string
	Port      string
	Protocol  string
	Anonymity string
	Country   string
	Region    string
	City      string
	//Uptime    string
	//Response  string
	//Transfer string
}

func getProxyFromFreeProxyListsNet(data []byte) (proxies []FreeProxyListNetProxy) {
	finalString := string(data)
	doc, _ := htmlquery.Parse(strings.NewReader(finalString))
	xpath := "//table[contains(.//text(), 'IP')]//tr[position() > 1 and not(.//iframe)]"
	rows := htmlquery.Find(doc, xpath)
	proxies = make([]FreeProxyListNetProxy, len(rows))
	for i, row := range rows {
		proxy := FreeProxyListNetProxy{}
		cols := htmlquery.Find(row, "./td[not(.//span)]//text()")
		proxy.IP = cols[0].Data
		proxy.Port = cols[1].Data
		proxy.Protocol = cols[2].Data
		proxy.Anonymity = cols[3].Data
		proxy.Country = cols[4].Data
		proxy.Region = cols[5].Data
		proxy.City = cols[6].Data
		proxies[i] = proxy
	}
	return proxies
}

// https://geonode.com/free-proxy-list/
// https://hidemy.name/en/proxy-list/
// http://free-proxy.cz/ru/
// https://hidemy.name/ru/proxy-list/
func main() {
	ctx := context.TODO()
	conf := redis_db.RedisConfig{}
	client := redis_db.RedisConnect(ctx, conf)
	defer client.Close()

}
