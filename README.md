Greeter
=======

* this is just my playground for grpc / grpc-web in combination with traefik and envoyproxy
* this was created to answer the question in https://stackoverflow.com/questions/59732958
* on the server side, traefik takes requests to two different domains (greeter.asdf.ooo & greeter-web.adsf.ooo) and terminates TLS
* traefik passes calls to greeter.asdf.ooo directly to the greeter-server
* traefik passes calls to greeter-web.asdf.ooo to envoy, which acts as a grpc-web-proxy before passing the request to the greeter-server

# run the server
* you need to have docker and docker-compose installed
* if you do not have traefik running directly on your server, add a traefik confiugration to the docker-compose.yml
* in the docker-compose.yml, edit the traefik labels and adjust the host rules (replace with your domains) and the certresolver
* run ```docker-compose up -d```

# run the golang client
```bash
# replace greeter.asdf.ooo with your domain
GREETER_TARGET=greeter.asdf.ooo:443 go run ./cmd/greeter-client/client.go
```

# run the web client
```bash
# browse to the web client
cd ./web-client
# install dev dependencies
npm install
# replace greeter-web.asdf.ooo with your domain
SERVER_URL="https://greeter-web.asdf.ooo" npm run dev
```
