
## 虚拟网络设备
veth
bridge

## 路由表

## 两个ns互通

```shell
ip netns add ns1
ip netns add ns2
ip link add veth0 type veth peer name veth1
ip link set veth0 netns ns1
ip link set veth1 netns ns2
ip netns exec ns1 ip link
ip netns exec ns2 ip link

ip netns exec ns1 ifconfig veth0 172.18.0.2/24 up
ip netns exec ns2 ifconfig veth1 172.18.0.3/24 up

# 容器内加容器出流量的路由表
ip netns exec ns1 route add default dev veth0
ip netns exec ns2 route add default dev veth1

ip netns exec ns1 ping 172.18.0.3

# 在宿主机上加入流量的路由表
route add -net 172.18.0.0/24 dev brO
```

## masquerade and nat
```shell
sysctl -w net.ipv4.conf.all.forwarding=l
iptables -t nat -A POSTROUTING -s 172.18.0.0/24 -o ethO -j MASQUERADE
```

```shell
iptables -t nat -A PREROUTING -p tcp -m tcp --dport 80 -j DNAT --to- destination 172.18.0.2:80
```

## 主机 ns 互通，bridge

```shell
ip netns add ns3
ip link add veth3 type veth peer name veth4
ip link veth3 setns ns3

brctl addbr br0
brctl addif br0 enp0s8
brctl addif br0 veth4
ip netns exec ns3 ping mi.com
```