<h1 align="center">rawdns 📡  </h1>

<h5 align="center">DNS messages (un)marshaller and UDP client</h5>

<br/>

### Overview

`rawdns` a small DNS client library that performs A records queries against a given nameserver. 

It relies only on `net` to perform the UDP queries, constructing and parsing all the messages in transit manually.

This is a project that is **not** intended to be used in production **at all**. It's not as efficient as it should and does not handle all the cases you'd expect from a production-ready library.

For real use, see [miekg/dns](https://github.com/miekg/dns).


### Usage

CLI:

```sh
rawdns example.com
msg sent labels=[example com]
&{ID:0 QR:1 Opcode:0 AA:0 TC:0 RD:1 RA:1 Z:0 RCODE:0 QDCOUNT:1 ANCOUNT:1 NSCOUNT:0 ARCOUNT:0}
ANSWER: [93 184 216 34]
```

Programatically:

```go
import "github.com/cirocosta/rawdns/lib"

client, err := lib.NewClient(lib.ClientConfig{
        Address: "8.8.8.8:53",
})
must(err)
defer client.Close()

ips, err := client.LookupAddr("example.com")
must(err)

for _, ip := range ips {
        fmt.Println(ip)
}
```


