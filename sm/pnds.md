## Docker 网络与 SRV

Docker 网络是 Docker 容器在网络层面的抽象，用于管理和连接容器之间的通信。Docker 网络允许容器之间进行相互通信和访问，以便构建分布式应用程序或服务。

SRV 记录是一种 DNS 记录类型，用于指定特定服务的域名、端口和优先级等信息。它用于在 DNS 中提供服务发现的功能，让客户端能够通过域名查找到提供特定服务的目标主机和端口。

SRV 记录的格式如下：

```
priority weight port target
```

+ `priority`: 优先级，用于指定备用服务器的顺序。值越小，优先级越高。
+ `weight`: 权重，用于在具有相同优先级的记录之间进行负载均衡。值越大，权重越高。
+ `port`: 服务的端口号。
+ `target`: 提供该服务的目标主机的域名。

## /etc/resolv.conf 配置文件

/etc/resolv.conf 配置文件，用于指定系统的 DNS 解析配置。它存储了用于解析域名的 DNS 服务器的相关信息。下面是关于
/etc/resolv.conf 文件的说明：

文件位置：/etc/resolv.conf 文件通常位于 Linux 系统中的 /etc 目录下。

文件格式：/etc/resolv.conf 是一个文本文件，每行包含一个配置项。配置项由关键字和对应的值组成，以空格或制表符分隔。常见的配置项包括：

+ nameserver：指定 DNS 服务器的 IP 地址。可以有多个 nameserver 行，按照优先级从上到下进行解析。
+ search：指定默认的搜索域名列表。当使用不完全限定的域名时，系统会自动尝试附加这些域名来进行解析。
+ domain：指定系统的默认域名。当使用不完全限定的主机名时，系统会自动尝试附加默认域名来进行解析。
+ options：指定其他的解析选项，如超时时间、转发等。
+

示例 /etc/resolv.conf 文件内容：

```shell
nameserver 8.8.8.8
nameserver 8.8.4.4
search example.com

```

可以看下 csphere 环境中的某个配置

![](https://cdn.xiaobinqt.cn/xiaobinqt.io/20230523/d83c10ae82e64b5f9722592503b351b2.png?imageView2/0/q/75|watermark/2/text/eGlhb2JpbnF0/font/dmlqYXlh/fontsize/1000/fill/IzVDNUI1Qg==/dissolve/52/gravity/SouthEast/dx/15/dy/15)

```shell
search xae-zcbus-20230523.28e5df10.csphere.local csphere.local
```

这表示系统默认的搜索域名列表包括两个域名：`xae-zcbus-20230523.28e5df10.csphere.local` 和 `csphere.local`。

当进行 DNS 解析时，如果使用一个不完全限定的域名（没有包含点号），系统会按照 search 配置中的顺序尝试逐个附加这些搜索域名，以尝试解析域名。例如，如果要解析的域名是
example，系统会依次尝试解析 `example.xae-zcbus-20230523.28e5df10.csphere.local` 和 `example.csphere.local`。

这样配置搜索域名列表可以简化 DNS 查询过程，特别是在本地网络内部使用内部域名时。通过设置合适的搜索域名，你可以直接使用不完全限定的主机名进行解析，而无需每次都输入完整的域名。

---

ndots 参数控制着系统在进行 DNS 解析时是否自动追加搜索域名。它指定了一个域名中至少要包含的点号（.）的数量。点号在域名中的数量表示域名的层级结构，例如
example.com 有一个点号，而 www.example.com 有两个点号。当一个域名的层级结构中点号的数量达到或超过 ndots 参数指定的值时，系统将不再追加搜索域名。

在 options ndots:2 的配置中，表示当进行 DNS 解析时，如果域名中包含至少两个点号（层级结构至少为三级），系统将不会自动追加搜索域名。如果域名的层级结构不足三级（点号少于两个），系统会自动尝试使用 /etc/resolv.conf 中指定的搜索域名列表进行解析。

---
`nameserver 172.17.0.1`用于指定 DNS 解析时要使用的 DNS 服务器的 IP 地址。这里的 172.17.0.1 是 bridge 网络下虚拟网桥 docker0 的 IP 地址。

因为容器中没有 dns 解析服务，不管是 dnsmasq 还是 pdns 都是装在宿主机上的，所以最后肯定是到宿主机上来解析的域名。但是这里如果直接写宿主机 IP 会影响效率「TODO 待补充」。

## /etc/hosts 配置文件

/etc/hosts 是本地的主机名解析文件，用于将主机名映射到对应的 IP 地址。

/etc/hosts 文件包含了主机名和对应 IP 地址的映射关系。每行的格式为 `<IP 地址> <主机名>` 或者 `<IP 地址> <主机别名> <主机名>`。

当系统需要解析主机名时，会首先查看 /etc/hosts 文件，如果找到匹配的主机名，则直接使用对应的 IP 地址进行通信，无需进行 DNS 查询。

可以看下 csphere 某个容器的 /etc/hosts

![](https://cdn.xiaobinqt.cn/xiaobinqt.io/20230523/8d0749739af8424ca863e0eecf3bcfca.png?imageView2/0/q/75|watermark/2/text/eGlhb2JpbnF0/font/dmlqYXlh/fontsize/1000/fill/IzVDNUI1Qg==/dissolve/52/gravity/SouthEast/dx/15/dy/15)

## 通过服务名称通信

![](https://cdn.xiaobinqt.cn/xiaobinqt.io/20230523/d0be2f67a7f742dea01c215ef99f02ad.png?imageView2/0/q/75|watermark/2/text/eGlhb2JpbnF0/font/dmlqYXlh/fontsize/1000/fill/IzVDNUI1Qg==/dissolve/52/gravity/SouthEast/dx/15/dy/15)

这个服务的服务名是`kingbase-x86`，当在容器中执行 `ping kingbase-x86`时是，正常会先走 hosts 但是这里的 hosts 都没有对应的域名，所以 hosts 失效，再去 nameserver 指定的 DNS 服务器解析域名，由于配置了 ndots:2 ，系统会自动追加搜索域名，所以会去 172.17.0.1 DNS 服务器搜索域名 `kingbase-x86.xae-zcbus-20230523.28e5df10.csphere.local`和 `kingbase-x86.csphere.local`。

由于是 bridge 网络，所以会走到虚拟网桥 docker0 的 172.0.0.1，通过 NAT 走到宿主机，在宿主机进行 DNS 解析。

每个 agent 上都有 dnsmasq，通过配置文件指定了上游的 DNS 服务器，`server=/域名/IP地址`：

![](https://cdn.xiaobinqt.cn/xiaobinqt.io/20230523/56eeb5b8e1cb495eb665eebebef1874e.png?imageView2/0/q/75|watermark/2/text/eGlhb2JpbnF0/font/dmlqYXlh/fontsize/1000/fill/IzVDNUI1Qg==/dissolve/52/gravity/SouthEast/dx/15/dy/15)

比如 agent11 指向了 14.5 的上游  DNS  服务。

在容器中执行 ping 时，比如 ping kingbase-x86，优先使用 A 记录来解析主机名为 IP 地址。













