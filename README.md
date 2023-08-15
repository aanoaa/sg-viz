# sgviz

local SecurityGroup data Visualize and management.

## Usage

``` bash
#
# import data from csv file.
#

$ cat <<EOL | sgviz import --host
foo-host01,192.168.0.1
foo-host02,192.168.0.2
bar-vserver,10.0.0.1
EOL

$ cat <<EOL | sgviz import --group
group,hostname
foo,foo-host01
foo,foo-host02
bar.example.com,bar-vserver
EOL

$ cat <<EOL | sgviz import --policy
src,dst,port,desc
foo,bar.example.com,8080,'foo to bar.example.com:8080'
EOL

#
# export readable csv format.
#
$ sgviz host
foo-host01,192.168.0.1,bar-vserver,10.0.0.1,8080,tcp
foo-host02,192.168.0.2,bar-vserver,10.0.0.1,8080,tcp

$ sgviz policy
foo,bar.example.com,8080,tcp

#
# visualize dot lang.
#
$ sgviz policy | sgviz graph
digraph {
  "foo" -> "bar.example.com"
}

$ sgviz policy | sgviz graph | dot -Tsvg > example.svg
```

## Raw Query

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
