package main

import "github.com/rs/xid"

func generatePrimaryKey() string {
	guid := xid.New()
	return guid.String()
}
