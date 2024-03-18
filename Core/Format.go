package Core

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

type TweetData struct {
	ID            string
	URL           string
	Text          string
	ReplyTo       string
	Username      string
	Fullname      string
	Timestamp     string
	Replies       string
	Retweets      string
	Quotes        string
	Likes         string
	QuoteFullname string
	QuoteUsername string
	QuoteDate     string
	QuoteID       string
	QuoteText     string
}

func FormatTweets(format string, tweets []Tweet) {
	if format == "json" {
		FormatTweetsJSON(tweets)
	} else {
		FormatTweetsCSV(tweets)
	}
}

func FormatTweetsCSV(tweets []Tweet) {
	//var b []byte
	//buf := bytes.NewBuffer(b)
	var tweetData []TweetData
	file, err := os.Create("records.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	w.Comma = ';'

	for _, tweet := range tweets {
		attachments := make([]string, len(tweet.Attachments))
		for i, att := range tweet.Attachments {
			attachments[i] = *att.URL
		}
		tweetData = append(tweetData, TweetData{
			ID:            tweet.ID,
			URL:           tweet.URL,
			Timestamp:     tweet.Timestamp,
			Username:      tweet.Username,
			Fullname:      tweet.Fullname,
			Text:          tweet.Text,
			ReplyTo:       tweet.ReplyTo,
			Replies:       fmt.Sprint(tweet.Stats.Replies),
			Retweets:      fmt.Sprint(tweet.Stats.Retweets),
			Quotes:        fmt.Sprint(tweet.Stats.Quotes),
			Likes:         fmt.Sprint(tweet.Stats.Likes),
			QuoteFullname: tweet.QuoteFullname,
			QuoteUsername: tweet.QuoteUsername,
			QuoteDate:     tweet.QuoteDate,
			QuoteText:     tweet.QuoteText,
			QuoteID:       tweet.QuoteID,
		})
		//if err := w.Write(row); err != nil {
		//	log.Fatalln("error writing row to csv:", err)
		//}
	}
	// write CSV header
	header := make([]string, 0)
	valueOf := reflect.ValueOf(tweetData[0])
	typeOf := valueOf.Type()
	for i := 0; i < valueOf.NumField(); i++ {
		header = append(header, typeOf.Field(i).Name)
	}
	if err := w.Write(header); err != nil {
		panic(err)
	}

	// write CSV body
	for _, armada := range tweetData {
		values := make([]string, 0)
		valueOf := reflect.ValueOf(armada)
		for i := 0; i < valueOf.NumField(); i++ {
			values = append(values, valueOf.Field(i).Interface().(string))
		}
		if err := w.Write(values); err != nil {
			panic(err)
		}
	}

	w.Flush()
	if err = w.Error(); err != nil {
		log.Fatal(err)
	}

	//fmt.Print(string(buf.Bytes()))
}

func FormatTweetsJSON(tweets []Tweet) {
	for _, tweet := range tweets {
		tweetJSON, _ := json.Marshal(tweet)
		fmt.Println(string(tweetJSON))
	}
}
