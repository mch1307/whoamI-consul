# whoamI-consul

Tiny Go webserver that displays host IP information

Forked from emilevauge/whoamI.

Story on http://wp.me/p95aLs-1c

Added: 
 - register/unregister itself to/from Consul service catalog (for Traefik/Consul demo)
 - cmd flags for Consul connection
 - ascii art banner displayed on the web page comes from consul k/v

```Usage: whoamI
  -consul string
        Consul service catalog address
  -consulPort int
        Consul service catalog port (default 8500)
  -consulToken string
        Consul ACL token (optional)
  -kvPath string
        Consul KV path for banner (optional) (default "PUBLIC/whoamI")
  -port int
        Port number for HTTP listen (default 8080)
  -service string
        Service name that will be registered (fqdn better) (default "whoamI")

```
