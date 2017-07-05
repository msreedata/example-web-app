Example web app
===============

A simple web app for demonstration purposes in other projects, like
[Soapbox](https://github.com/adhocteam/soapbox).

Getting started
---------------

**Requirements:**

 - Docker 17.05 or above (for [multi-stage builds](https://docs.docker.com/engine/userguide/eng-image/multistage-build/#use-multi-stage-builds))

### Build the app

``` shell
$ docker build -t example-web-app .
```

### Run it

``` shell
$ docker run -it --rm --name example-web-app-running -p 9090:8080 example-web-app &
$ curl http://localhost:9090/_healthcheck
OK
```
