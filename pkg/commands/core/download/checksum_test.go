package download

import (
	"bytes"
	"testing"
)

type md5Test struct {
	data []byte
	md5  string
}

var md5Tests = []md5Test{
	{[]byte(`WordPress started in 2003 when Mike Little and Matt Mullenweg created a fork of b2/cafelog.`), "fa79868d88ebcc02000166a0e99e206f"},
	{[]byte(`The need for an elegant, well-architected personal publishing system was clear even then.`), "c3cee3be2c69fc355e44533a90928ea3"},
	{[]byte(`Today, WordPress is built on PHP and MySQL, and licensed under the GPLv2.`), "d8fac2c60e8c744ac6e815e88d385994"},
	{[]byte(`It is also the platform of choice for over 33% of all sites across the web.`), "68cede8370ca04cff8ed018690c002a2"},
}

func TestGetMD5(t *testing.T) {
	for _, tt := range md5Tests {
		md5 := getMD5(bytes.NewReader(tt.data))
		if md5 != tt.md5 {
			t.Errorf("MD5 Hash does not match: expected %s, actual %s", tt.md5, md5)
		}
	}
}
