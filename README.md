# sg-viz

local SecurityGroup data Visualize and management.

``` sql
-- policies by host.
SELECT sh.hostname AS src, sh.ipaddr AS src_addr, dh.hostname AS dst, dh.ipaddr AS dst_addr, p.port, p.protocol
FROM policy p
JOIN sgroup src ON src.id = p.src
JOIN sgroup dst ON dst.id = p.dst
JOIN host_sgroup shg ON shg.sgroup_id = src.id
JOIN host_sgroup dhg ON dhg.sgroup_id = dst.id
JOIN host sh ON sh.id = shg.host_id
JOIN host dh ON dh.id = dhg.host_id;
```

``` sql
-- policies by group.
SELECT src.name AS src, dst.name AS dst, p.port, p.protocol
FROM policy p
JOIN sgroup src ON src.id = p.src
JOIN sgroup dst ON dst.id = p.dst
JOIN host_sgroup shg ON shg.sgroup_id = src.id
JOIN host_sgroup dhg ON dhg.sgroup_id = dst.id
JOIN host sh ON sh.id = shg.host_id
JOIN host dh ON dh.id = dhg.host_id
GROUP BY src.name, dst.name, p.port;
```
