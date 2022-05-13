package mqtt

import "strings"

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
