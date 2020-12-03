mackerel-plugin-mysqlrouter
============================
[![Build Status](https://cloud.drone.io/api/badges/rluisr/mysqlrouter_exporter/status.svg)](https://cloud.drone.io/rluisr/mysqlrouter_exporter)

Supported version
-----------------
check [here](https://github.com/rluisr/mysqlrouter-go#supported-version)

Usage
-----
1. Enable REST API on MySQL Router [here](https://github.com/rluisr/mysqlrouter-go#supported-version)
2. Download binary from [release page](https://github.com/rluisr/mackerel-plugin-mysqlrouter/releases) and put to `/usr/local/bin`
3. Edit mackerel-agent.conf then restart.
```
[plugin.metrics.mysqlrouter]
command = ["/usr/local/bin/mackerel-plugin-mysqlrouter"]
env = { MYSQLROUTER_URL = "http://localhost:8080", MYSQLROUTER_USER = "luis", MYSQLROUTER_PASS = "luis", MYSQLROUTER_TLS_SKIP_VERIFY = "true" }
```

FYI
---
You should set alert of health per route like below.

![](https://raw.githubusercontent.com/rluisr/image-store/master/mackerel-plugin-mysqlrouter/mackerel-plugin-mysqlrouter01.png)

This `0.1` is no meaning. If route is no heath return `0`.

Todo
----
- [ ] Show graph per route connections
