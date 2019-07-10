## Introduction
This story is about my journey in the Hackathon by Digital Ocean called CloudHack. In this hackathon, one is expected to pick a problem statement that he/she thinks can help the community and solve it by using several Digital Ocean services.

The problem I’m trying to address
These days there are several good static site generators, like Hugo, Nikola, etc, in which creating a static site is easy and the end result is a beautiful and fast static site, but a major problem still remains the deployment part, that would you deploy the site on cloud that you’ve just created. There are several services that can do that for you like AWS Amplify but they charge you for that, but what’s the fun in letting someone else handle your deployment! Instead, the real fun lies in managing the deployment in a cool way as possible. Plus, sometimes these deployment services abstract a lot of implementation details and hence makes the process of deployment less customizable. Moreover, deployment of the site can be done manually as well, but it’s really cumbersome and time-consuming to do it manually, also, when you can automate, why not automate :D

## The solution I’m building
Here my goal to create a tool that can easily deploy your static site to Digital Ocean Droplet in the most convenient way possible. You can either deploy your local contents or make your GitHub repo serve as a source for your static site using Hugo.

## Solution
1. Create a Digital Ocean droplet and assign it your public SSH key (generally stored in ~/.ssh/id_rsa.pb)  

2. Copy the public ssh key of this droplet to the GitHub account which has the contents for the site. (if the content is stored in local, it is advised to store it in Github as a repository)  

3. Install Ansible on your machine. On Mac, you can install ansible using its brew package manager- `brew install ansible`   

4. cd into cloudhack/scripts then add executable permission to deploy.sh by running chmod +x deploy.sh.  

5. Append export GOPATH="/root/development/go" in the ~/.profile file. Then source ~/.profile.  

6. Now go to the /scripts directory, add required data in the inventory file and WorkingDirectory information in the hugoserver.service file, then run `./deploy.sh`.  


## DigitalOcean products used in the hack
* Digital Ocean Droplet

## Conclusion
Your Hugo powered static site is now available at http://<ip-address>. To make it point your custom domain name, follow this link. Hope this automation helped you deploy your site easily to the cloud of your choice, preferably Digital Ocean!

* [link to the write-up](https://medium.com/@souvikhaldar/digital-ocean-cloudhack-deploy-static-site-to-do-droplet-easily-ec89e5136f8c)