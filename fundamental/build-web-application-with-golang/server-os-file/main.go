package main

import (
	"fmt"
	"os"
)

func main() {
	folder := "one/two/three"
	os.MkdirAll(folder, 0777)
	
	userFile := folder + "/test_file.txt"
	
	fout, err := os.Create(userFile)
    if err != nil {
        fmt.Println(userFile, err)
        return
    }	
	defer fout.Close()
	for i := 0; i < 10; i++ {
        fout.WriteString("Just a test!\r\n")
        fout.Write([]byte("Just a test!\r\n"))
    }
	
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
    for {
        n, _ := fl.Read(buf)
        if 0 == n {
            break
        }
        os.Stdout.Write(buf[:n])
	}
	os.Remove(userFile)
	os.RemoveAll("one")
}