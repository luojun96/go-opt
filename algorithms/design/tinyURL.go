package main

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/encode-and-decode-tinyurl/

const k1, k2 = 1117, 1e9 + 7

type Codec struct {
	dataBase map[int]string
	urlToKey map[string]int
}

func NewCodec() Codec {
	return Codec{map[int]string{}, map[string]int{}}
}

func (c *Codec) encode(longUrl string) string {
	if key, ok := c.urlToKey[longUrl]; ok {
		return "http://tinyurl.com/" + strconv.Itoa(key)
	}

	key, base := 0, 1
	for _, c := range longUrl {
		key = (key + int(c)*base) % k2
		base = (base + k1) % k2
	}

	for c.dataBase[key] != "" {
		key = (key + 1) % k2
	}

	c.dataBase[key] = longUrl
	c.urlToKey[longUrl] = key
	return "http://tinyurl.com/" + strconv.Itoa(key)

}

func (c *Codec) decode(shortUrl string) string {
	i := strings.LastIndexByte(shortUrl, '/')
	key, _ := strconv.Atoi(shortUrl[i+1:])
	return c.dataBase[key]
}
