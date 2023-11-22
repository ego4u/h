package h

import "github.com/mitchellh/mapstructure"

func Map2Struct(data, result any) error {
	return mapstructure.Decode(data, result)
}
