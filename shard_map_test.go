package shard_map_test

import (
	"testing"
	shard_map "github.com/mapfeeling/shard"
	_ "github.com/mapfeeling/shard"
)

func BenchmarkItems(b *testing.B) {
	var m := shard_map.New(string)
}
