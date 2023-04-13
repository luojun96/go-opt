package algorithms

// https://leetcode.cn/problems/design-twitter/description/?languageTags=golang
type Node struct {
	followee map[int]bool
	tweet    []int
}

func newNode() *Node {
	return &Node{
		followee: make(map[int]bool),
		tweet:    make([]int, 0),
	}
}

var recentMax, time int
var tweetTime map[int]int
var user map[int]*Node

type Twitter struct {
	time      int
	recentMax int
	tweetTime map[int]int
	user      map[int]*Node
}

func Constructor() Twitter {
	return Twitter{
		time:      0,
		recentMax: 10,
		tweetTime: map[int]int{},
		user:      map[int]*Node{},
	}
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := t.user[userId]; !ok {
		t.user[userId] = newNode()
	}
	if len(t.user[userId].tweet) == recentMax && len(t.user[userId].tweet) > 0 {
		t.user[userId].tweet = t.user[userId].tweet[1:]
	}
	t.user[userId].tweet = append(t.user[userId].tweet, tweetId)
	t.time++
	t.tweetTime[tweetId] = t.time
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	ans := []int{}
	var user *Node
	user, ok := t.user[userId]
	if ok {
		for i := range user.tweet {
			if len(ans) < t.recentMax {
				ans = append(ans, user.tweet[len(user.tweet)-i-1])
			}
		}
	} else {
		return []int{}
	}

	preTweets := ans
	for followeeId, _ := range user.followee {
		if followeeId == userId {
			continue
		}
		res := []int{}
		followeeTweets := t.user[followeeId].tweet
		for len(preTweets) > 0 && len(followeeTweets) > 0 {
			preTweet := preTweets[0]
			followeeTweet := followeeTweets[len(followeeTweets)-1]

			if t.tweetTime[preTweet] > t.tweetTime[followeeTweet] {
				res = append(res, preTweet)
				preTweets = preTweets[1:]
			} else {
				res = append(res, followeeTweet)
				followeeTweets = followeeTweets[:len(followeeTweets)-1]
			}

			if len(res) >= t.recentMax {
				break
			}
		}

		for len(preTweets) > 0 && len(res) < t.recentMax {
			preTweet := preTweets[0]
			res = append(res, preTweet)
			preTweets = preTweets[1:]
		}

		for len(followeeTweets) > 0 && len(res) < t.recentMax {
			followeeTweet := followeeTweets[len(followeeTweets)-1]
			res = append(res, followeeTweet)
			followeeTweets = followeeTweets[:len(followeeTweets)-1]
		}

		preTweets = res
		ans = res
	}

	return ans
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := t.user[followerId]; !ok {
		t.user[followerId] = newNode()
	}
	if _, ok := t.user[followeeId]; !ok {
		t.user[followeeId] = newNode()
	}
	t.user[followerId].followee[followeeId] = true
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if u, ok := t.user[followerId]; ok {
		delete(u.followee, followeeId)
	}
}
