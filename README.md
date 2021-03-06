## CloudHack
Easily deploy static sites created using [Hugo](https://gohugo.io/) to any cloud instance of choice!  
(currently supports debian based servers only like Ubuntu and it's derivatives, Kali Linux, Raspbian,etc) 

## Steps
1. Create a cloud instance and assign it your system's public SSH key (generally stored in ~/.ssh/id_rsa.pb)  
2. Copy the public ssh key of this droplet to the GitHub account which has the contents for the site. (if the content is stored in local, it is advised to store it in Github as a repository).   
3. Install Ansible on your machine. On Mac, you can install ansible using its brew package manager- `brew install ansible`   
4. `cd` into the cloned repo i.e `cloudhack/` then add executable permission to `deploy.sh` and `conf<OS>` by running `chmod +x deploy.sh` and `chmod +x confLinux`.  
5. `./conf<OS> -ip <ip> -repo <repo-url>`  eg `confMac -ip <ip> -repo <repo-url>` on Mac.  
6. `./deploy.sh`  

P.S: Make sure port 80 of this server is open in the firewall setting and is not busy. 


## Result
Your Hugo powered static site is now available at `http://<ip-address>`  
I followed [this](https://www.howlthemes.com/point-domain-name-digitalocean-droplet/) article to assign a custom domain name (souvikhaldar.info in my case) to the given IP address. If you have already assigned the IP to your domain name, you can provide the domain name in the `ip` flag to `conf`. I've used Digital Ocean's droplet, but you can use any cloud service provider of choice. Hope this automation tool helped you deploy your site easily to the cloud of your choice.   
## Terminal recording
[![asciicast](https://asciinema.org/a/343311.svg)](https://asciinema.org/a/343311)

## Video demonstration 
Link to youtube video for easy demonstration
https://youtu.be/Op3M1PBSQL4

* [Link to my website deployed by above method](http://souvikhaldar.info)
