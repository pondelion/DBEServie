# DBEServie

## Build

- android/arm64 target

```bash
$ GOOS=android GOARCH=arm64 GOARM=7 go build -o droid_dbe_service
```

- 

```bash
$ go build -o dbe_service
```

## Run API server on android

```bash
$ adb push droid_dbe_service /data/local/tmp
$ adb shell
$ su
# cd /data/local/tmp
# chmod 777 droid_dbe_service
# ./droid_dbe_service
```

## API call from client

- Start tcpdump

```bash
$ curl -X POST http://[ANDROID_HOST]:8080/network/start_tcpdump -d "save_dir=."
{"message":"successfully started tcpdump","pcap_filepath":"1640937116.pcap"}
```

- Stop tcpdump

```bash
$ curl -X POST http://[ANDROID_HOST]:8080/network/stop_tcpdump
{"message":"killed tcpdump process : 24208","pcap_filepath":"1640937116.pcap"}
```
