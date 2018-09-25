package test

import (
	"my_blog/utils"
	"testing"
)

func TestRedis(t *testing.T) {
	utils.SetRedisValue("liping", "110", "50")
}
