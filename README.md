```bash
curl -L https://github.com/nats-io/nats-server/releases/download/v2.0.0/nats-server-v2.0.0-linux-amd64.zip -o nats-server.zip
```

```bash
docker run -p 4222:4222 -ti nats:latest
```

```bash
nats-server -p 4222 -cluster nats://localhost:4248 
nats-server -p 5222 -cluster nats://localhost:5248 -routes nats://localhost:4248 
nats-server -p 6222 -cluster nats://localhost:6248 -routes nats://localhost:4248 
```