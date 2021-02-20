# bilibili-live-notificator

It detects starting the live streaming on Bilibili and notifies a Twitter.

## How to use.

```
git clone 
cd bilibili-live-notificator
./bilibili-live-notificator -r BILIBILI_ROOM_ID -t YOUR_TWITTER_ACCESS_TOKEN -w
```

## -i: Bilibili room ID

https://live.bilibili.com/<room ID>

## -s TWITTER_API_KEY


## -w watching

Start the process and check it every few minutes.
Once it detects starting live stream, it will not notify you until the live stream is finished.