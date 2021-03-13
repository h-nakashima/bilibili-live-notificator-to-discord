# bilibili-live-notificator

It detects starting the live streaming on Bilibili and notifies a Twitter.

## How to use.
- Download binary file from https://github.com/h-nakashima/bilibili-live-notificator/releases
- Run
```
./bilibili-live-notificator -i BILIBILI_ROOM_ID -k API_KEYS_FILE
```

## -i: Bilibili room ID

https://live.bilibili.com/＜room ID＞

## -k: API keys file
- copy config file
```
cp ./api_keys.yml.sample ./api_keys.yml
```
- edit ./api_keys.yml
