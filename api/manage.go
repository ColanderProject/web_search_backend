package api

import (
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (server *Server) connectionManager(ctx *gin.Context) {
	key := ctx.Query("key")
	if key == "" {
		ctx.String(http.StatusBadRequest, "Key is not provided!")
		return
	} else if key == "WSNqRKs!lQgd5wOVb" {
		operation := ctx.Query("op")
		switch {
		case operation == "connect":
			ctx.String(http.StatusOK, "Has been deprecated")
		case operation == "get_hot_query":
			idx := ctx.Query("idx")
			idx_val, err := strconv.Atoi(idx)
			if err != nil {
				log.Fatal(err)
			}
			getHotQuery(idx_val)
		}
	} else {
		ip, cid, imp := getRequestInfo(ctx)
		sendAlert("connection_manage_bad_key"+ip+" "+cid+" "+imp, "connection_manage_bad_key"+key)
		ctx.String(http.StatusBadRequest, "Key is not right!")
	}

}

func sendAlert(title, content string) error {
	conn, err := net.Dial("tcp", "1.ty.es2q.com:12345")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	titleLen := len(title)
	contentLen := len(content)

	_, err = conn.Write([]byte("lQ5xlEkdBN95o5D9\n"))
	if err != nil {
		log.Fatal("Failed to write", err)
	}

	_, err = conn.Write([]byte(strconv.Itoa(titleLen) + "\n"))
	if err != nil {
		log.Fatal("Failed to write titleLen", err)
	}

	_, err = conn.Write([]byte(strconv.Itoa(contentLen) + "\n"))
	if err != nil {
		log.Fatal("Failed to write contentLen", err)
	}

	_, err = conn.Write([]byte(title))
	if err != nil {
		log.Fatal("Failed to write title", err)
	}

	_, err = conn.Write([]byte(content))
	if err != nil {
		log.Fatal("Failed to write content", err)
	}

}

// Get the in the request
//
// @return ip, cid, imp as three seperate string on success
func getRequestInfo(ctx *gin.Context) (string, string, string) {
	var ip string

	xForwardedFor := ctx.GetHeader("X-Forwarded-For")
	if xForwardedFor == "" {
		ips := strings.Split(xForwardedFor, ",")
		ip = ips[len(ips)-1]
	} else {
		ip = ctx.Request.RemoteAddr
	}

	cid, err := ctx.Cookie("cid")
	if err != nil {
		log.Fatal("cid does not exist in cookie", err)
	}
	imp, err := ctx.Cookie("imp")
	if err != nil {
		log.Fatal("imp does not exist in cookie", err)
	}
	return ip, cid, imp
}

func getHotQuery(idx int) {
	
} 