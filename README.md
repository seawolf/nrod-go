# NROD Train Movements Processor

Sample application to process [Train Movements](https://wiki.openraildata.com/index.php?title=Train_Movements) messages from the Network Rail Open Data feed.

General information about the feed data can be found on the excellent [Open Rail Data wiki](https://wiki.openraildata.com/index.php?title=About_the_Network_Rail_feeds). The connection is made with [the STOMP protocol](https://wiki.openraildata.com/index.php?title=Connecting_with_Stomp#Network_Rail).

#### Example Output
```
2022/01/29 11:56:16 Connecting: datafeeds.networkrail.co.uk:61618 ...
2022/01/29 11:56:16 Connected: ID:opendata-prod[...]
2022/01/29 11:56:16 Subscribed: 1

2022/01/29 11:56:16 Waiting for messages...

2022/01/29 11:56:22 train 1P27 by South Western Railway has arrived at HAVANT (HAV) platform 2 travelling DOWN
2022/01/29 11:56:22 train 1L40 by South Western Railway has departed from TEMPLECOMBE (TMC) platform 1 travelling UP towards GILLINGHAM (DORSET) (GIL)
2022/01/29 11:56:22 train 1L40 by South Western Railway has arrived at GILLINGHAM (DORSET) (GIL) platform 1 travelling UP
2022/01/29 11:56:22 train 2U32 by South Western Railway has arrived at DATCHET (DAT) platform 1 travelling UP
2022/01/29 11:56:22 train 1D25 by South Western Railway has departed from STONELEIGH (SNL) platform 2 travelling DOWN towards EWELL WEST (EWW)
2022/01/29 11:56:22 train 2D27 by South Western Railway has arrived at VAUXHALL (VXH) platform 6 travelling DOWN
2022/01/29 11:56:22 train 2G30 by South Western Railway has arrived at CLAPHAM JUNCTION (CLJ) / CLAPHAM JN. C.S. (XCP) platform 10 travelling UP
2022/01/29 11:56:22 train 1W58 by South Western Railway has arrived at HOLTON HEATH (HOL)
2022/01/29 11:56:22 train 1L44 by South Western Railway has departed from HONITON (HON) platform 1 travelling UP
2022/01/29 11:56:22 train 1P38 by South Western Railway has departed from FRATTON (FTN) platform 1 on line "M" travelling UP
2022/01/29 11:56:22 train 2L38 by South Western Railway has departed from BASINGSTOKE (BSK) platform 1 travelling UP towards HOOK (HOK)
2022/01/29 11:56:22 Waiting for messages...

^C

2022/01/29 11:56:25 Cleaning-up...
2022/01/29 11:56:25 Unsubscribing from subscription...
2022/01/29 11:56:25 Successfully unsubscribed from subscription.
2022/01/29 11:56:25 Disconnecting from connection...
2022/01/29 11:56:25 Sucessfully disconnected from connection.
2022/01/29 11:56:25 Cleaned-up; exiting.
```

## Usage

#### Prerequisites
1. an account at [the Network Rail Data Feeds website](http://datafeeds.networkrail.co.uk/)
2. a subscription to the Train Movements feed of **one** TOC - note that this software is currently limited to one connection/subscription/TOC

#### Invocation

1. copy the `.env.example` file to `.env`
2. fill in the username (`STOMP_USERNAME`) and password (`STOMP_PASSWORD`) with your log-in credentials of [the Network Rail Data Feeds website](http://datafeeds.networkrail.co.uk/).
3. fill in the TOC code for the feed name `TRAIN_MVT_??_TOC`, e.g. a subscription to `Stagecoach Sth Western Trains Ltd (HY)` shows a TOC code of `HY` to create the feed name `TRAIN_MVT_HY_TOC`.
4. build the container with `docker build -t nrod-go:latest .`
5. run the container with `docker run --rm --env-file ./.env nrod-go:latest`