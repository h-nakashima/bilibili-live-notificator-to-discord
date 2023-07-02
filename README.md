# bilibili-live-notificator

This application detects the start of live streaming on Bilibili and sends notifications to Twitter.

## How to use.

- Download binary file from https://github.com/h-nakashima/bilibili-live-notificator/releases
- Run

```
./bilibili-live-notificator -i BILIBILI_ROOM_ID -k API_KEYS_FILE
```

## -i: Bilibili room ID

https://live.bilibili.com/{$room_id}

## -k: API keys file

- copy config file

```
cp ./api_keys.yml.sample ./api_keys.yml
```

- edit ./api_keys.yml
