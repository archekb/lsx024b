package mqtt

import (
	"math/rand"
	"strings"
	"time"
)

func TopicPrepare(topic string) []string {
	topicSplited := strings.Split(topic, "/")
	newTopicSplited := []string{}

	for i, v := range topicSplited {
		if topicSplited[i] == "" {
			continue
		}
		newTopicSplited = append(newTopicSplited, v)
	}

	return newTopicSplited
}

func randSeq(n int) string {
	var str = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = str[rand.Intn(len(str))]
	}
	return string(b)
}
