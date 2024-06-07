# Reverse Proxy Lolancer
pratice reverse proxy and load balancer

# Required
- [Download docker on your computer](https://docs.docker.com/get-docker/)

# Stack
- GIT
- Docker
- NGINX
- GOLANG Framework Fiber

# Setting
```
    cd ~/Desktop 
    git clone https://github.com/PingTeeratorn789/reverse_proxy.git
    cat .env.example > .env
    docker compose up -d --scale app=5
```
# Reference
[การ Scale Docker Compose ด้วย Nginx Reverse Proxy พร้อมทำ Load Balancer](https://www.youtube.com/watch?v=ykwAI_8Pkvw)