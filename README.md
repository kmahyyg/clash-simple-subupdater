K2P 16M ROM+128M RAM

## Clash

Binary and Clash Web is from https://opt.cn2qq.com

All scripts except listed above are written by Patrick Young and licensed under AGPL v3.

Use clash-dashboard siince it's smaller.

## TroubleShooting

- You should at least have 5M available space.
- Modify Update-Config.sh and manually update clash config, service will auto restart.
- Too many open files:

```
root@K2P:/etc/sysctl.d# cat ./99-ulimit.conf
fs.file-max=65536
root@K2P:/etc/sysctl.d# sysctl -p /etc/sysctl.d/99-ulimit.conf
fs.file-max = 65536
```

- **Due to the lack of feature SO_MARK, this method CANNOT proxy gateway itself unless you use cgroup.**
