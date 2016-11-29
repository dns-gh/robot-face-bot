// Space Rocks Bot is a bot watching
// asteroids coming too close to earth for the incoming days/week.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/dns-gh/freeze"
	robot "github.com/dns-gh/robohash-client/robohashclient"
	"github.com/dns-gh/twbot"

	"strconv"

	conf "github.com/dns-gh/flagsconfig"
)

// Twitter constants
const (
	projectURL               = "https://github.com/dns-gh/robot-face-bot"
	updateFlag               = "update"
	twitterFollowersPathFlag = "twitter-followers-path"
	twitterFriendsPathFlag   = "twitter-friends-path"
	twitterTweetsPathFlag    = "twitter-tweets-path"
	debugFlag                = "debug"
)

type timeWriter struct {
	writer io.Writer
}

func (w timeWriter) Write(p []byte) (int, error) {
	date := time.Now().Format("[2006-01-02 15:04:05] ")
	p = append([]byte(date), p...)
	return w.writer.Write(p)
}

func makeDateWriter(w io.Writer) io.Writer {
	return &timeWriter{w}
}

func makeLogger(path string) (string, *os.File, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", nil, err
	}
	err = os.MkdirAll(filepath.Dir(abs), os.ModePerm)
	if err != nil {
		return "", nil, err
	}
	f, err := os.OpenFile(abs, os.O_WRONLY+os.O_APPEND+os.O_CREATE, os.ModePerm)
	return abs, f, err
}

var (
	// original quotes from here: https://github.com/e1ven/Robohash/blob/master/robohash/webfront.py
	quotes = []string{
		"#But #but #but.. I still #love you!",
		"#Please don't leave this #tweet.. When no one's #here.. It gets #dark...",
		"#Script #fatal #error on line #42",
		"Don't #trust the other #robots. I'm the only #trustworthy one.",
		"My #fuel is the #misery of children. And #Rum. Mostly #Rum... Well... #Energy also that's true...",
		"When they said they'd give me a body transplant, I didn't think they meant this!",
		"Subject 14 has had it's communication subroutines deleted for attempting suicide.",
		"I am the cleverest robot on the whole page.",
		"#Oil can. #Can I? I #think I use #Oil. #Oil can. #Can I? I #think I use #Oil...",
		"I am #fleunt in over 6 #million #forms of #communishin.",
		"I see a little #silhouette of a #bot..",
		"I WANT MY #HANDS #BACK!",
		"#Please don't #reload, I'll #DIE!",
		"#Robots don't have #souls, you know. But they do #feel #pain.",
		"I #wonder what would #happen if all the #robots went #rogue.",
		"10: #KILL ALL #HUMANS. 20: GO 10",
		"I'm the #best #robot here.",
		"The #green #robot #thinks you're #cute.",
		"Any #robot you don't click on, they# dismantle.",
		"#Robot #tears #taste like #candy.",
		"01010010010011110100001001001111010101000101001100100001!",
		"Your #mouse #cursor #tickles.",
		"#Logic #dictates #placing me on your #tweet.",
		"I #think my #arm is on #backward.",
		"I'm #different!",
		"It was the #best of #times, it was ಠ_ಠ worth #living...",
		"#String is #Gnirts spelled #backward, you know",
		"We're no #strangers to #hashing.. You know the 3 rules, and so do I..",
		"Please. #Destroy. Me... ... well no no no no on second #thought...",
		"#Pick #Me! #Pick #Me!",
		"Tzzzzzz shhhh - df d- test... I'm still #configuring #myself as a #machine... - , dfsl sdjbl , please be #gentle.... I'm a #robot --- fdsi !",
	}
)

func main() {
	update := flag.Duration(updateFlag, 6*time.Hour, "[twitter] update frequency of the bot for tweets")
	twitterFollowersPath := flag.String(twitterFollowersPathFlag, "followers.json", "[twitter] data file path for followers")
	twitterFriendsPath := flag.String(twitterFriendsPathFlag, "friends.json", "[twitter] data file path for friends")
	twitterTweetsPath := flag.String(twitterTweetsPathFlag, "tweets.json", "[twitter] data file path for tweets")
	twitterConsumerKey := flag.String("TWITTER_CONSUMER_KEY", "", "[twitter] consumer key")
	twitterConsumerSecret := flag.String("TWITTER_CONSUMER_SECRET", "", "[twitter] consumer secret")
	twitterAccessToken := flag.String("TWITTER_ACCESS_TOKEN", "", "[twitter] access token")
	twitterAccessSecret := flag.String("TWITTER_ACCESS_SECRET", "", "[twitter] access secret")
	debug := flag.Bool(debugFlag, false, "[twitter] debug mode")
	_, err := conf.NewConfig("robot.config")
	log.SetFlags(0)
	logPath, f, err := makeLogger(filepath.Join(filepath.Dir(os.Args[0]), "Debug", "bot.log"))
	if err == nil {
		defer f.Close()
		log.SetOutput(makeDateWriter(io.MultiWriter(f, os.Stderr)))
	}
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("[info] logging to:", logPath)
	log.Println("[twitter] update:", *update)
	log.Println("[twitter] twitter-followers-path:", *twitterFollowersPath)
	log.Println("[twitter] twitter-friends-path:", *twitterFriendsPath)
	log.Println("[twitter] twitter-tweets-path:", *twitterTweetsPath)
	log.Println("[twitter] twitterConsumerKey:", *twitterConsumerKey)
	log.Println("[twitter] twitterConsumerSecret:", *twitterConsumerSecret)
	log.Println("[twitter] twitterAccessToken:", *twitterAccessToken)
	log.Println("[twitter] twitterAccessSecret:", *twitterAccessSecret)
	log.Println("[twitter] debug:", *debug)
	bot := twbot.MakeTwitterBotWithAccess(*twitterFollowersPath, *twitterFriendsPath, *twitterTweetsPath,
		*twitterConsumerKey,
		*twitterConsumerSecret,
		*twitterAccessToken,
		*twitterAccessSecret,
		*debug)
	defer bot.Close()
	roboClient := robot.MakeRobohashClient(200, 200, 1, 1)
	Fetch := func() (string, string, string, error) {
		timeStr := strconv.FormatInt(time.Now().UnixNano(), 10)
		img, err := roboClient.Fetch(timeStr)
		return freeze.GetRandomElement(quotes), img, "", err
	}
	bot.TweetImagePeriodicallyAsync(Fetch, 30*time.Minute)
	bot.TweetPeriodicallyAsync(func() (string, error) {
		return fmt.Sprintf("Hey, I'm a bot, check out my source code %s and help me improve ! 🤖", projectURL), nil
	}, *update)
	bot.Wait()
}
