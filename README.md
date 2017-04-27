# gpio-example
Experimental example on reading data from DHT sensor(temperature, humidity), on Raspberry Pi 2.

# Hardware
## Raspberry Pi
Raspberry Pi 2 (model B+) BCM2709 Armv7l

## Sensor
Adafruit DHT22

# Build & Deploy & Run

## Directly
This is recommended way,

```sh
# On dev machine
./build.sh
scp bin/dht pi@rasphost:/tmp
# On Raspi
/tmp/dht
```

## Docker
If docker already installed on Raspi,

```sh
# On dev machine
./build.sh
docker build -t zyfdedh/gpio-dht:dev .
docker push zyfdedh/gpio-dht:dev
# On Raspi
docker logs -f $(docker run --privileged -v /dev/mem:/dev/mem zyfdedh/gpio-dht:dev)
```
> Note: replace zyfdedh with your own DockerHub ID.
