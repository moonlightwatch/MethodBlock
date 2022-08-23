# HTTP Request Method Block

config example

```yml
# Static configuration

experimental:
  plugins:
    methodBlock:
        moduleName: github.com/moonlightwatch/methodblock
        version: v0.1.1

```

```yml
# Dynamic configuration

http:
  routers:
    my-router:
      rule: host(`demo.localhost`)
      service: service-foo
      entryPoints:
        - web
      middlewares:
        - my-plugin

  services:
   service-foo:
      loadBalancer:
        servers:
          - url: http://127.0.0.1:5000
  
  middlewares:
    my-plugin:
      plugin:
        methodBlock:
          Message: "Method Not Allowed"
          Methods:
            - GET
            - POST
```