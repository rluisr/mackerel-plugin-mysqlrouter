mackerel-plugin-mysqlrouter
============================
[![Build Status](https://cloud.drone.io/api/badges/rluisr/mysqlrouter_exporter/status.svg)](https://cloud.drone.io/rluisr/mysqlrouter_exporter)

Usage
-----
1. Download binary from release page and put to `/usr/local/bin`
2. Edit mackerel-agent.conf then restart.
```
[plugin.metrics.mysqlrouter]
command = "MYSQLROUTER_URL=http://localhost:8080 MYSQLROUTER_USER=luis MYSQLROUTER_PASS=luis /usr/local/bin/mackerel-plugin-mysqlrouter"
```

FYI
---


Todo
----
- [ ] Show graph per route connections