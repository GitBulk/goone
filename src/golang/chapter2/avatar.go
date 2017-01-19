package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func getGravatarHash(email string) string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)
	hashMd5 := md5.New()
	io.WriteString(hashMd5, email)
	finalBytes := hashMd5.Sum(nil)
	finalString := hex.EncodeToString(finalBytes)
	return finalString
}

func main() {
	//var email string = "danchithancongab@gmail.com"
	var email string = os.Args[1]
	var gravatar = getGravatarHash(email)
	fmt.Println(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8" />
    <title>Gravatar in Go</title>
</head>
<body>
    <div>
        <img src="http://www.gravatar.com/avatar/` + gravatar + `" alt="user picture" />
	</div>
</body>
</html>`)
}
