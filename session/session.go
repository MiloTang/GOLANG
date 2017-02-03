package session

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Start(w http.ResponseWriter, r *http.Request, maxAge int) {
	if c, err := r.Cookie("BeenoobCookies"); err != nil {
		if c == nil {
			crutime := time.Now().Unix()
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			randstr := string(r.Intn(99999999))
			uid := encode(strconv.FormatInt(crutime, 10) + randstr)
			filename := encode(uid + "BeenoobCookie")
			uid_cookie := &http.Cookie{
				Name:     "BeenoobCookies",
				Value:    uid,
				Path:     "/",
				HttpOnly: false,
				MaxAge:   maxAge,
			}
			http.SetCookie(w, uid_cookie)
			file, err := os.Create("sessionfile/" + filename)
			checkError(err)
			defer file.Close()
			txt := "cookiename:" + uid + "\n" + "time:" + strconv.FormatInt(crutime, 10) + "\n"
			_, err1 := file.WriteString(txt)
			checkError(err1)
		}
	}

}
func Setvalue(r *http.Request, key string, value string) bool {
	if c, err := r.Cookie("BeenoobCookies"); err != nil {
		fmt.Println(err)
	} else {
		if c != nil {
			filename := encode(c.Value + "BeenoobCookie")
			file, err := os.OpenFile("sessionfile/"+filename, os.O_APPEND, 0666)
			checkError(err)
			defer file.Close()
			if Getvalue(r, key) == "" {
				txt := key + ":" + value + "\n"
				_, err1 := file.WriteString(txt)
				if err1 == nil {
					return true
				}
			}
		}
	}
	return false
}
func Getvalue(r *http.Request, key string) string {
	if c, err := r.Cookie("BeenoobCookies"); err != nil {
		fmt.Println(err)
	} else {
		if c != nil {
			filename := encode(c.Value + "BeenoobCookie")
			file, err := os.Open("sessionfile/" + filename)
			checkError(err)
			defer file.Close()
			reader := bufio.NewReader(file)
			for {
				line, _, err := reader.ReadLine()
				if err != nil {
					fmt.Println(err)
					break
				}
				linestr := string(line)
				fmt.Println(linestr)
				strs := strings.Split(linestr, ":")
				if key == strs[0] {
					return strs[1]
				}
			}
		}
	}
	return ""
}
func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func encode(str string) string {
	md5init := md5.New()
	io.WriteString(md5init, str)
	md5value := fmt.Sprintf("%x", md5init.Sum(nil))
	return md5value
}
