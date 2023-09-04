package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Connection struct {
	ip string
	port string
	connection Conn 
	queryCount int 
	errCount int
	reconnectCount int
	bool 
}

var (
	
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/view", func(c *gin.Context) {
		u := c.Query("u")
		k := c.Query("k")
		ip, cid, imp := getClientInfo(c.Request)

		fileName := "log/click_" + strconv.FormatInt(time.Now().Unix()/3600, 10)
		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("err")
		}
		defer file.Close()

		log.SetOutput(file)

		log.Println(ip, cid, imp, u, k)
		c.Redirect(301, "https://bbs.byr.cn"+u)
	})
	r.GET("/api/s", func(c *gin.Context) {
		ip, cid, imp := getClientInfo(c.Request)
		key := c.Query("q")
		iStr := c.Query("i")
		if iStr == "" {
			i := 0
		} else {
			i, err := strconv.Atoi(iStr)
			if err != nil {

			}
		}

		if key == "" {
			c.String(" ")
		}

		if len(key) > 50 {
			key = key[:50]
		}

		if {

		}
	})

	r.Run()
}

func getClientInfo(req *http.Request) (string, string, string) {
	var ip, cid, imp string

	xForwardedFor := req.Header.Get("X-FORWARDED-FOR")
	if xForwardedFor != "" {
		ips := strings.Split(xForwardedFor, ",")
		ip = strings.TrimSpace(ips[0])
	} else {
		ip = req.Header.Get("REMOTE_ADDR")
	}

	cidCookie, err := req.Cookie("cid")
	if err == nil {
		cid = cidCookie.Value
	} else {
		cid = ""
	}

	impCookie, err := req.Cookie("imp")
	if err == nil {
		imp = impCookie.Value
	} else {
		imp = ""
	}

	return ip, cid, imp
}

func connectionManager(c *gin.Context) {
	key := c.Query("key")
	if key == "WSNqRKs!lQgd5wOVb" {
		op := req.Header.Get("op")
		if op == "connect" {
			c.String(200, "Has been deprecated!")
		} else if op == "get_hot_query" {
			idx := c.Query("idx")
			hotQuery, err := getHotQuery(idx)
			if err != nil {
				jsonData, err := json.Marshal(hotQuery)
				c.String(http.StatusOK, string(jsonData))
			}
		} else if op == "load_index_html" {
			c.HTML(http.StatusOK, "index.html", gin.H{})	
		} else if op == "get_server_load" {

		} else if op == "connect2" {
			ip := c.Query("ip")
			port := c.Query("port")
			key := c.Query("ckey")

			conn, err := connect2Establish(ip, port, key)
			if err != nil {
				c.String(http.StatusOK, "error" + err.Error())
			}

			c.String(http.StatusOK, "success")

		} else if op == "reconnect" {

		} else if op == "get_connection" {

		} else if op == "connect_cache_server" {

		} else if op == "close" {

		} else if op == "shutdown" {

		} else if op == "get" {

		} else if op == "update_hot" {

		} else if op == "add_announcement" {

		}
	} else {

	}
}

func getHotQuery(idx int) ([][2]string, error){
	var hotQuery [][2]string

	filePath := fmt.Sprintf("log/query_%d.log", idx)

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	queryMap := make(map[string]int) 

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			fields := strings.Split(line, "\t")
			query := fields[4]
			queryMap[query]++ 
		}
	}

	for query, count := range queryMap {
		hotQuery = append(hotQuery, [2]string{query, fmt.Sprintf("%d", count)})
	}

	sort.Slice(hotQuery, func(i, j int) bool {
		return hotQuery[i][1] > hotQuery[j][1]
	})

	if len(hotQuery) > 10 {
		hotQuery = hotQuery[:10]
	}

	return hotQuery, nil
}

func query(key string, i int) {

}

func connect2Establish(ip, port, key string) (Conn error) {
	
	conn, err := net.Dial("tcp", ip + ":" + port)
	if err != nil {
		fmt.Println("Failed to establish connection: ", err)
		return
	}

	fmt.Println("connect_establish: ", ip, port, key)

	_, err = conn.Write([]byte(key + "\n"))
	if err != nil {
		fmt.Println("Failed to send key:", err)
		return nil, err
	}

	reader := bufio.NewReader(conn)
	line, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("Failed to read string:", err)
		return nil, err
	}
	if sting(line) == "ok" {
		return conn, nil
	}

	conn.Close()
	return nil, err
}