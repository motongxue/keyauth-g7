package utils_test

import (
	"github.com/motongxue/keyauth-g7/common/utils"
	"testing"
)

func TestToken(t *testing.T) {
	v := utils.MakeBearer(24)
	t.Log(v)
}
