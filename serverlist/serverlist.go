package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"time"
	"strconv"
	"math"
)

/*
mpServerConf["port"] = "" + get_config(20);
mpServerConf["username"] = "δ֪";
mpServerConf["use"] = "";
path = get_config(2);
mpServerConf["ligic_dir"] = path;
*/
type register_server_info struct{
	ip, port, username, use, desc, logic_dir string
	update_time time.Time
}

var mpAllServerInfo = map[string] register_server_info{
	"dummy:1001" : {"127.0.0.1", "2010", "mincao", "test", "detail", "/home/tx2/logic/test", time.Now()},
}

func get_server_list(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	for _, info := range mpAllServerInfo {
		diff := now.Sub(info.update_time)
		if 35 < diff.Seconds() {
			continue
		}
		arr_data := [] string {info.ip, info.port, info.username, info.use, info.desc, info.logic_dir, ""}
		data := strings.Join(arr_data, "|")
		fmt.Fprintf(w, data)
		fmt.Fprintf(w, "\n")
	}
}

func register_server(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	remote_info := strings.Split(r.RemoteAddr, ":")
	ip := remote_info[0]

	port      := r.Form["port"][0]
	username  := r.Form["username"][0]
	use       := r.Form["use"][0]
	desc      := r.Form["desc"][0]
	logic_dir := r.Form["logic_dir"][0]

	now := time.Now()
	elem, ok := r.Form["time"]
	if ok {
		server_now, err := strconv.ParseInt(elem[0], 10, 0)
		if err != nil {
			return
		}
		serverlist_time := now.Unix()
		if 30 < math.Abs(float64(server_now - serverlist_time)) {
			logic_dir = fmt.Sprintf("%s#rError Time : %dsec", logic_dir, int(math.Abs(float64(server_now - serverlist_time))))
		}
	}

	info := [] string {ip, port}
	server_key := strings.Join(info, ":")
	data := register_server_info{ip, port, username, use, desc, logic_dir, now}
	mpAllServerInfo[server_key] = data
	fmt.Fprintf(w, "register succ\n")
	fmt.Fprintf(w, server_key)
	fmt.Fprintf(w, now.Format("\n2006 7 2 15:04:05"))
}

func main() {
	http.HandleFunc("/get_server_list", get_server_list)
	http.HandleFunc("/register_server", register_server)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err);
	}
}
