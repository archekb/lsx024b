# lsx024b

Epsolar LSx024B controler monitor (Web and Home Assistant).

App will read data from device (every `{Device.Interval}` seconds) and push it to `MQTT` topic `{MQTT.Topic}/{Device.Name}` or (and) show `http(s)://{HTTP.Address}/api/state` you can also use Web client `http(s)://{HTTP.Address}`.

	{
		"connected": "online",
		"updated": "2022-05-13T18:23:51.997628484Z",
		"update_interval": 5,
		"device": {
			"rated": { ... },
			"real_time": { ... },
			"status": { ... },
			"statistical": { ... },
			"settings": { ... }
		}
	}


### Configuration ###
---

For configuration create `config.yml` (see `example/config.yml`) file and discribe Device and MQTT or Web or bouth services.


### Build ###
---

You have two ways for build this project:
* Install `Go` lang, `Nodejs 12 (14 or 16)`, `yarn` and build manually
* Use `make` and `Docker`

*or download binary from realese section*.

**Attention!** For update web client you should run web client building before server building, if you use `docker` and `make` new client will added automatically.


#### Server ####
---

You have golang on your computer, you can build it by default: 

	go mod download
	go build

or if you have `make` and `Docker`, you can build without `golang` instalation: (for your host platform) 

	make build

if you wanna build version for arm32v7 (Orabge PI, etc) you need qemu emulator you can install it directly `sudo apt install qemu-user-static` or add docker container with emulator by `make qemu`

	make qemu
	make build_arm32v7


#### Web client ####
---

You have `nodejs 12 / 14 / 16` and `yarn` on your computer, you can build it by default:

	yarn install
	yarn build
	rm -rf ../web/*
	cp -r dist/* ../web/

or if you have `make` and `Docker`, you can build without `node` instalation:

	make build


### Run as service ###
---

Add config file and service file (see `example/lsx024b.service`) to `/opt/lsx024b/` and create symlink to service:

	sudo ln -s /opt/lsx024b/lsx024b.service /etc/systemd/system/lsx024b.service
	sudo systemctl enable lsx024b
	sudo systemctl start lsx024b


#### for status: ####

	sudo systemctl status lsx024b


#### for stop ####

	sudo systemctl disable lsx024b


#### for delete service ####

	sudo systemctl disable lsx024b


### More info ###
---

Device docs and vendor client(for Windows 7/10) you can fount in `/docs`.
