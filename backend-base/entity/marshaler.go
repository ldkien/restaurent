package entity

import "github.com/golang/protobuf/jsonpb"

var Marshaler = jsonpb.Marshaler{
	EmitDefaults: true,
	EnumsAsInts: true,
	OrigName: true,
}

