# whoamI

Forked from emilevauge/whoamI
Tiny Go webserver that displays host IP information
Added: Register/Deregister itself to Consul service catalog (for Traefik/Consul demo)

```Usage: whoamI
-consul string
        Consul service catalog address
  -consulPort string
        Consul service catalog port (default "8500")
  -consulToken string
        Consul ACL token (optional)
  -port int
        Port number for HTTP listen (default 80)
  -banner string
        Banner displayed on web page
```
