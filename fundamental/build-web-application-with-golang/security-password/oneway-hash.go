package main

import (
	"fmt"
	"io"
	"crypto/sha256"
	"crypto/sha1"
	"crypto/md5"
)

func main() {
	// 単方向ハッシュ
	// import "crypto/sha256"
	h256 := sha256.New()
	io.WriteString(h256, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("SHA256: %s x\n", h256.Sum(nil))
	// import "crypto/sha1"
	h1 := sha1.New()
	io.WriteString(h1, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("SHA1: %s x\n", h1.Sum(nil))
	// import "crypto/md5"
	hMd5 := md5.New()
	io.WriteString(hMd5, "暗号化する必要のあるパスワード")
	fmt.Printf("MD5: %x\n", hMd5.Sum(nil))

	// ソルトの利用
	h := md5.New()
	io.WriteString(h, "暗号化が必要なパスワード")
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
	salt1 := "@#$%"
	salt2 := "^&*()"

	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)
	last := fmt.Sprintf("%s", h.Sum(nil))
	fmt.Printf("PW-hash: %x\n", last)
}
