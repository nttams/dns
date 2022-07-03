# DNS server & client

DnsClient is independent
DnsServer depends on DnsClient for upstream query

## Todo
* network handling (UDP, TCP, timeout, retransmission, TCP fallback)
* parser and encoder
* EDNS (RFC6891)
* zone handler
* zone transer
* upstream query
* caching
* asynchronous