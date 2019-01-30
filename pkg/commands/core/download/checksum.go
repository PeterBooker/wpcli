package download

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
)

func getMD5(r io.Reader) string {
	h := md5.New()
	if _, err := io.Copy(h, r); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(h.Sum(nil))
}
