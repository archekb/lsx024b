# lsx024b

Epever (Epsolar) Controller Monitor (Web and Home Assistant).

**Expected compatible controller series**: `VS-B`, `Tracer-B`, `Tracer-A`, `iTracer`, `eTracer`

**Confirmed compatible controller series**: `LS-B (LS1024B, LS2024B, LS3024B)`

App will read data from device (every `{Device.Interval}` seconds) and push it to `MQTT` topic `{MQTT.Topic}/{Device.Name}` or (and) show `http(s)://{HTTP.Address}/api/state` you can also use Web client `http(s)://{HTTP.Address}`.

```json
{
	"connected": true,
	"updated": "2022-05-13T18:23:51.997628484Z",
	"update_interval": 10,
	"model": "LS-B compatible",
	"device": {
		"rated": { ... },
		"real_time": { ... },
		"status": { ... },
		"statistical": { ... },
		"settings": { ... }
	}
}
```

### Configuration ###
---

For configuration create `config.yml` (see `example/config.yml`) file and discribe Device and MQTT or Web or bouth services.


### Build ###
---

You have two ways for build this project:
* Install `Go` lang, `Nodejs 12 (14 or 16)`, `yarn` and build manually
* Use `make` and `Docker`

*or download binary from realese section*.

NOTE: Set version for production building `go build -s -w -X main.version=1.0` or empty `go build` for Debug log level.

**ATTENTION!** For update web client you should run web client building before server building, if you use `docker` and `make` new client will added automatically.


#### Server ####
---

You have golang on your computer, you can build it by default: 

```shell script
go mod download
go build
```

or if you have `make` and `Docker`, you can build without `golang` instalation: (for your host platform) 

```shell script
make build
```

if you wanna build version for arm32v7 (Orabge PI, etc) you need qemu emulator you can install it directly `sudo apt install qemu-user-static` or add docker container with emulator by `make qemu`

```shell script
make qemu
make build_arm32v7
```

> **NOTE!** After reboot you need run `make qemu` again.

> **NOTE!** When you see in console `#0 0.580 exec /bin/sh: exec format error` that's mean you need qemu.


#### Web client ####
---

You have `nodejs 12 / 14 / 16` and `yarn` on your computer, you can build it by default:

```shell script
yarn install
yarn build
rm -rf ../web/*
cp -r dist/* ../web/
```

or if you have `make` and `Docker`, you can build without `node` instalation:

```shell script
make build
```

> **NOTE!** You can use development docker container `make dev_web_client`. 

For using JS helper Vetur for VSCode create `vetur.config.js` in root, and set path to project `'./web_src'`. (see reference: https://vuejs.github.io/vetur/reference/#example)


### Run as service ###
---

Add config file and service file (see `example/lsx024b.service`) to `/opt/lsx024b/` and create symlink to service:

```shell script
sudo ln -s /opt/lsx024b/lsx024b.service /etc/systemd/system/lsx024b.service
sudo systemctl enable lsx024b
sudo systemctl start lsx024b
```

#### for status: ####

```shell script
sudo systemctl status lsx024b
```

#### for stop ####

```shell script
sudo systemctl disable lsx024b
```


#### for delete service ####

```shell script
sudo systemctl disable lsx024b
```

### More info ###
---

Device docs you can fount in `/docs`.
