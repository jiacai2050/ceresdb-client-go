// Copyright 2022 CeresDB Project Authors. Licensed under Apache-2.0.

package types

type RouteMode int

const (
	Direct RouteMode = iota
	Proxy
)

type Route struct {
	Table    string
	Endpoint string
	Ext      []byte
}
