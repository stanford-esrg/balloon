balloon
=========

balloon inflates a delta-compressed list of IPv4 services. 

## Building

Set up `$GOPATH` (see https://golang.org/doc/code.html).
```
$ go install github.com/stanford-esrg/balloon@latest
$ cd $GOPATH/src/github.com/stanford-esrg/balloon
```

balloon inflates IPv4 services that are compressed using:

## Usage

To inflate a list of delta-compressed IPv4 called `theList.csv`:

```
./balloon theList.csv
```

## How to delta-compress IPv4 services

Delta compressing IPv4 services dramatically reduces the amount of memory needed to store large ``hitlists'' of services.
Run the following BiqQuery query on a table with a list of ips and ports to produce a delta-compressed list:
```
SELECT case when comp_port is not null then
  FORMAT("%t", TO_HEX(SUBSTR(NET.IPV4_FROM_INT64(comp_port),3,2))) else
  null end portb,
  case  when s3 is null and s2 is not null then 
  FORMAT("%t",TO_HEX(SUBSTR(NET.IP_FROM_STRING(ip),1,3)))
  when s2 is null and s1 is not null then 
  FORMAT("%t",TO_HEX(SUBSTR(NET.IP_FROM_STRING(ip),1,2) ))
  when s1 is null then
  FORMAT("%t",TO_HEX(SUBSTR(NET.IP_FROM_STRING(ip),1,1)))
  else  FORMAT("%t", TO_HEX(NET.IP_FROM_STRING(ip))) end ipb, FROM (
  SELECT  ip,port,r, case when port = prev_port then null else port end comp_port, s0,
          CASE when port = prev_port and s1 = prev_s1 and s2 = prev_s2 and s3 = prev_s3 then null else s1 end s1,
          CASE when port = prev_port and s2 = prev_s2  and s3 = prev_s3 then null else s2 end s2,
          CASE when port = prev_port and s3 = prev_s3 then null else s3 end s3, FROM (
            SELECT r, port, ip, LAG(port) over (order by r) prev_port,
                    CAST(SPLIT(ip,".")[OFFSET(3)] as int64) s3,
                    LAG(CAST(SPLIT(ip,".")[OFFSET(3)] as int64)) over (order by r) prev_s3,
                    CAST(SPLIT(ip,".")[OFFSET(2)] as int64) s2,
                    LAG(CAST(SPLIT(ip,".")[OFFSET(2)] as int64)) over (order by r) prev_s2,
                    CAST(SPLIT(ip,".")[OFFSET(1)] as int64) s1,
                    LAG(CAST(SPLIT(ip,".")[OFFSET(1)] as int64)) over (order by r) prev_s1,
                    CAST(SPLIT(ip,".")[OFFSET(0)] as int64) s0 FROM (
                      SELECT distinct ip, port,
                            ROW_NUMBER() over (ORDER BY port, CAST(SPLIT(ip,".")[OFFSET(3)] as int64), 
                            CAST(SPLIT(ip,".")[OFFSET(2)] as int64),
                            CAST(SPLIT(ip,".")[OFFSET(1)] as int64),
                            CAST(SPLIT(ip,".")[OFFSET(0)] as int64)) r FROM (
                              SELECT distinct ip, port FROM dataset.table
                            )
                    ORDER BY port, CAST(SPLIT(ip,".")[OFFSET(3)] as int64), 
                    CAST(SPLIT(ip,".")[OFFSET(2)] as int64),
                    CAST(SPLIT(ip,".")[OFFSET(1)] as int64),
                    CAST(SPLIT(ip,".")[OFFSET(0)] as int64)
          ) 
  ORDER BY r
  )
) ORDER BY r
```
