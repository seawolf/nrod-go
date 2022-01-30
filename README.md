# NROD Train Movements Processor

Sample application to process [Train Movements](https://wiki.openraildata.com/index.php?title=Train_Movements) messages from the Network Rail Open Data feed.

General information about the feed data can be found on the excellent [Open Rail Data wiki](https://wiki.openraildata.com/index.php?title=About_the_Network_Rail_feeds). The connection is made with [the STOMP protocol](https://wiki.openraildata.com/index.php?title=Connecting_with_Stomp#Network_Rail).

#### Example Output
```
2022/01/29 22:30:20 Connecting: datafeeds.networkrail.co.uk:61618 ...
2022/01/29 22:30:20 Connected: ID:opendata-prod[...]
2022/01/29 22:30:20 Subscribed: TRAIN_MVT_ET_TOC (1)
2022/01/29 22:30:20 Subscribed: TRAIN_MVT_HY_TOC (2)

2022/01/29 22:30:22 [2] train 2K67 by South Western Railway has arrived at HAMPTON WICK (HMW) platform 2 travelling DOWN
2022/01/29 22:30:22 [2] train 1L73 by South Western Railway has arrived at VIRGINIA WATER (VIR)
2022/01/29 22:30:22 [2] activation of train 1A75 for South Western Railway
2022/01/29 22:30:22 [2] train 2U74 by South Western Railway has departed from SUNNYMEADS (SNY) platform 1 travelling UP towards WRAYSBURY (WRY)
2022/01/29 22:30:22 [2] train 2U74 by South Western Railway has arrived at WRAYSBURY (WRY) platform 1 travelling UP
2022/01/29 22:30:32 [1] train 1I66 by Govia Thameslink Railway (Thameslink / Southern) has arrived at HACKBRIDGE (HCB) platform 2 travelling DOWN
2022/01/29 22:30:32 [1] train 1F65 by Govia Thameslink Railway (Thameslink / Southern) has arrived at NORWOOD JUNCTION (NWD) platform 3 travelling UP
2022/01/29 22:30:32 [1] train 9O79 by Govia Thameslink Railway (Thameslink / Southern) has arrived at FARRINGDON (ZFD) platform 3 travelling DOWN
2022/01/29 22:30:32 [2] train 2S71 by South Western Railway has arrived at ST DENYS (SDN) platform 2 travelling DOWN
2022/01/29 22:30:32 [2] train 2E72 by South Western Railway has departed from BITTERNE (BTE) platform 1 travelling UP towards ST DENYS (SDN)
2022/01/29 22:30:32 [2] train 2H69 by South Western Railway has arrived at RAYNES PARK (RAY) platform 3 travelling DOWN
2022/01/29 22:30:32 [2] train 2E72 by South Western Railway has arrived at ST DENYS (SDN) platform 4 travelling UP
2022/01/29 22:30:32 [2] activation of train 2T75 for South Western Railway

2022/01/30 16:37:07 [1] Train 2H43 by South Western Railway has arrived at MORTLAKE (MTL) platform 2 travelling DOWN
2022/01/30 16:37:07 [1] Activation of train 2L59 for South Western Railway originating at LONDON WATERLOO (WAT)
2022/01/30 16:37:07 [1] Cancellation of train 2L59 for South Western Railway effective from LONDON WATERLOO (WAT) due to System generated cancellation
2022/01/30 16:37:07 [1] Train 2S45 by South Western Railway has departed from DEAN (DEN) platform 2 travelling DOWN

^C

2022/01/29 22:30:34 Cleaning-up...
2022/01/29 22:30:34 Unsubscribing from subscription...
2022/01/29 22:30:34 Successfully unsubscribed from subscription.
2022/01/29 22:30:34 Disconnecting from connection...
2022/01/29 22:30:34 Sucessfully disconnected from connection.
2022/01/29 22:30:34 Cleaned-up; exiting.
```

## Usage

#### Prerequisites
1. an account at [the Network Rail Data Feeds website](http://datafeeds.networkrail.co.uk/)
2. a subscription to the Train Movements feed of one or more TOCs

#### Invocation

1. copy the `.env.example` file to `.env`
2. fill in the username (`STOMP_USERNAME`) and password (`STOMP_PASSWORD`) with your log-in credentials of [the Network Rail Data Feeds website](http://datafeeds.networkrail.co.uk/).
3. fill in the TOC code for the feed name `FEEDS_TRAIN_MOVEMENTS`, e.g. a subscription to `Stagecoach Sth Western Trains Ltd (HY)` shows a TOC code of `HY` to create the feed name `TRAIN_MVT_HY_TOC`. Multiple feeds are comma-separated, e.g. `FEEDS_TRAIN_MOVEMENTS=TRAIN_MVT_HY_TOC,TRAIN_MVT_ET_TOC`
4. build the container with `docker build -t nrod-go:latest .`
5. run the container with `docker run --rm --env-file ./.env nrod-go:latest`