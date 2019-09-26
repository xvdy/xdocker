# cgroup

### 概念

* cgroup 进程分组的机制
* subsystem 对某种资源进行配额限制
* hierachy

## demo

```$xslt
lssubsys

cd /sys/fs/cgroup

apt install stress
stress --vm-bytes 200m --vm-keep -m 1
top 

cd /sys/fs/cgroup/memory
mkdir test-limit-memory && cd test-limit-memory
echo 100m > memory.limit_in_bytes
echo $$ > tasks

stress --vm-bytes 200m --vm-keep -m 1
top 
```
