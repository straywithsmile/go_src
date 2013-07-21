package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"time"
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

var html_data string = `<html>
<head>
  <title>Sample</title>
  <script type="text/javascript" src="https://www.google.com/jsapi"> </script>
  <script type="text/javascript">
    var ge;
    google.load("earth", "1");

    function init() {
      google.earth.createInstance('map3d', initCB, failureCB);
    }

    function initCB(instance) {
      ge = instance;
      ge.getWindow().setVisibility(true);
    }

    function failureCB(errorCode) {
    }

    google.setOnLoadCallback(init);
  </script>

</head>
<body>
  <div id="map3d" style="height: 400px; width: 600px;"></div>
</body>
</html>`

func online(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, html_data)
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

	info := [] string {ip, port}
	server_key := strings.Join(info, ":")
	data := register_server_info{ip, port, username, use, desc, logic_dir, time.Now()}
	mpAllServerInfo[server_key] = data
	fmt.Fprintf(w, "register succ\n")
	fmt.Fprintf(w, server_key)
	now := time.Now()
	fmt.Fprintf(w, now.Format("\n2006 7 2 15:04:05"))
}

func main() {
	http.HandleFunc("/online", online)
	err := http.ListenAndServe(":9080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err);
	}
}
