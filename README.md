# whoamI

Tiny Go webserver that displays host IP information (for demo purpose)

Forked from emilevauge/whoamI.

Added: 
- register/unregister itself to/from Consul service catalog (for Traefik/Consul demo)
- cmd flags for Consul connection
- optional banner to be displayed in aasci art on the web page


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
