# robot-face-bot

[![Go Report Card](https://goreportcard.com/badge/github.com/dns-gh/robot-face-bot)](https://goreportcard.com/report/github.com/dns-gh/robot-face-bot)

The robot face bot is a Twitter Bot connected to the https://robohash.org/ API able to tweet robot messages with hashed computed robot faces.

See https://twitter.com/RoboFaceBot for an instance of it.

It uses the following client to talk to the robohash API: https://github.com/dns-gh/robohash-client

It has also a specified and user-defined bot behaviors thanks to https://github.com/dns-gh/twbot.

## Motivation

Simply for fun and practice.

## Installation

- It requires Go language of course. You can set it up by downloading it here: https://golang.org/dl/
- Install it here C:/Go.
- Set your GOPATH, GOROOT and PATH environment variables with:

```
export GOROOT=C:/Go
export GOPATH=WORKING_DIR
export PATH=C:/Go/bin:${PATH}
```

or:

```
@working_dir $ source build/go.sh
```

and then set up your API keys/tokens/secrets (you can find them here: https://apps.twitter.com/)

```
export TWITTER_CONSUMER_KEY="your_twitter_consumer_key"
export TWITTER_CONSUMER_SECRET="your_twitter_consumer_secret"
export TWITTER_ACCESS_TOKEN="your_twitter_access_token"
export TWITTER_ACCESS_SECRET="your_twitter_access_secret"
```

or put them in the configuration file generated after a first launch.

## Build and usage

```
@working_dir $ go install robotfacebot
@working_dir $ bin/robotfacebot.exe -help
  -TWITTER_ACCESS_SECRET string
        [twitter] access secret
  -TWITTER_ACCESS_TOKEN string
        [twitter] access token
  -TWITTER_CONSUMER_KEY string
        [twitter] consumer key
  -TWITTER_CONSUMER_SECRET string
        [twitter] consumer secret
  -config string
        configuration filename (default "robot.config")
  -debug
        [twitter] debug mode
  -twitter-followers-path string
        [twitter] data file path for followers (default "followers.json")
  -twitter-friends-path string
        [twitter] data file path for friends (default "friends.json")
  -twitter-tweets-path string
        [twitter] data file path for tweets (default "tweets.json")
  -update duration
        [twitter] update frequency of the bot for tweets (default 6h0m0s)
@working_dir $ bin/robotfacebot.exe -debug=false
[2016-11-29 17:18:38] [info] logging to: D:\WORK\robot-face-bot\bin\Debug\bot.log
[2016-11-29 17:18:38] [twitter] update: 6h0m0s
[2016-11-29 17:18:38] [twitter] twitter-followers-path: followers.json
[2016-11-29 17:18:38] [twitter] twitter-friends-path: friends.json
[2016-11-29 17:18:38] [twitter] twitter-tweets-path: tweets.json
[2016-11-29 17:18:38] [twitter] twitterConsumerKey: YOUR_TWITTER_CONSUMER_KEY_OR_EMPTY
[2016-11-29 17:18:38] [twitter] twitterConsumerSecret: YOUR_TWITTER_CONSUMER_SECRET_OR_EMPTY
[2016-11-29 17:18:38] [twitter] twitterAccessToken: YOUR_TWITTER_ACCESS_TOKEN_OR_EMPTY
[2016-11-29 17:18:38] [twitter] twitterAccessSecret: YOUR_TWITTER_ACCESS_SECRET_OR_EMPTY
[2016-11-29 17:18:38] [twitter] debug: false
[2016-11-29 17:18:38] [robohash] making robohash client
```

## License

See the included LICENSE file.