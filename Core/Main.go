package Core

import (
	"net/url"
)

func Main(Query *string, Instance *string, Format *string) {
	(*Query) = url.QueryEscape(*Query)
	Scrape(Request(Query, Instance), Instance, Format)
}
