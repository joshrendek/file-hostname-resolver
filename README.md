# Usage

Two modes:

RDNS lookup for IPs, IP lookup for Hostnames

```
go install github.com/joshrendek/file-hostname-resolver
cat list-of-hostnames.txt | file-hostname-resolver -ip-lookup
cat list-of-ips.txt | file-hostname-resolver -rdns
```
