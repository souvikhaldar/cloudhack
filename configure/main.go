package main
import(
	"flag"
	"fmt"
	"io/ioutil"
)
var serviceContent = `
[Unit]
Description= hugo server

[Service]
WorkingDirectory= /root/StaticSite/
ExecStart= /usr/local/bin/hugo server --watch --baseUrl=%s --port=80 --appendPort=false --bind=0.0.0.0
Restart= always

[Install]
WantedBy= multi-user.target
`
var inventoryContent = `
[droplet]
%s


[all:vars]
remote_user = %s
repo = %s
`

func main(){
	ip := flag.String("ip","","IP address of the server on which to deploy")
	repo := flag.String("repo","","url of the repo which contains the website content")
	remoteUser := flag.String("user","root","name of the remote user")
	flag.Parse()
	fmt.Printf("The target server is: %s\n",*ip)
	fmt.Println("Source repository is: ",*repo)
	if err := ioutil.WriteFile("hugoserver.service",[]byte(fmt.Sprintf(serviceContent,*ip)),0644); err != nil{
		fmt.Println("Unable to write to the unit file: ",err)
		return
	}
	if err := ioutil.WriteFile("inventory",[]byte(fmt.Sprintf(inventoryContent,*ip,*remoteUser,*repo)),0644); err != nil{
		fmt.Println("Unable to write to the inventory file: ",err)
		return
	}

}