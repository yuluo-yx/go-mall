package testUtils

import (
	"go-mall/pkg/util"
	"strings"
	"testing"
)

func TestAes(t *testing.T) {

	str := "test"

	// 设置key
	util.Encrypt.SetKey("test123456789876")

	get := util.Encrypt.AesEncoding(str)
	want := "Rk8IUoo/E6RUmcDfSLdhMg=="

	if !strings.EqualFold(get, want) {
		t.Errorf("want: %v \t get: %v", want, get)
	}

	decryptAESStr := util.Encrypt.AesDecoding(want)
	if !strings.EqualFold(str, decryptAESStr) {
		t.Errorf("decryptAESStr: %v \t str: %v", decryptAESStr, str)
	}
}
