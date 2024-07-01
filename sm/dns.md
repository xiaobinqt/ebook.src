
创建网络

```shell
docker network create mynetwork
```


```shell
docker run -d --network=mynetwork --name=wb-redis -p 60000:6379 redis

docker run -d --network=mynetwork --name=wb-mysql57 -p 60001:3306 mysql:5.7
```

DNS 解析工具 dnsmasq，配置文件在 /etc/dnsmasq.conf 下



