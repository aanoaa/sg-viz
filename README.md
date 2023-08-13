# sg-viz

local SecurityGroup data Visualize and management.

``` sql
-- policies by host.
SELECT fh.hostname AS src, fh.ipaddr AS src_addr, th.hostname AS dst, th.ipaddr AS dst_addr, p.port, p.protocol
FROM policy p
JOIN sgroup f ON f.id = p.`from`
JOIN sgroup t ON t.id = p.`to`
JOIN host_sgroup fhg ON fhg.sgroup_id = f.id
JOIN host_sgroup thg ON thg.sgroup_id = t.id
JOIN host fh ON fh.id = fhg.host_id
JOIN host th ON th.id = thg.host_id;
```

``` sql
-- policies by group.
SELECT f.name AS src, t.name AS dst, p.port, p.protocol
FROM policy p
JOIN sgroup f ON f.id = p.`from`
JOIN sgroup t ON t.id = p.`to`
JOIN host_sgroup fhg ON fhg.sgroup_id = f.id
JOIN host_sgroup thg ON thg.sgroup_id = t.id
JOIN host fh ON fh.id = fhg.host_id
JOIN host th ON th.id = thg.host_id
GROUP BY f.name, t.name, p.port;
```
