package main

import (
	"log"
	"net"
)

type LogEntry struct {
	Level   string  `json:"level"`
	Ts      float64 `json:"ts"`
	Logger  string  `json:"logger"`
	Msg     string  `json:"msg"`
	Request struct {
		Method     string              `json:"method"`
		URI        string              `json:"uri"`
		Proto      string              `json:"proto"`
		RemoteAddr string              `json:"remote_addr"`
		Host       string              `json:"host"`
		Headers    map[string][]string `json:"headers"`
		TLS        struct {
			Resumed     bool   `json:"resumed"`
			Version     int    `json:"version"`
			Ciphersuite int    `json:"ciphersuite"`
			Proto       string `json:"proto"`
			ProtoMutual bool   `json:"proto_mutual"`
			ServerName  string `json:"server_name"`
		} `json:"tls"`
	} `json:"request"`
	CommonLog string  `json:"common_log"`
	Duration  float64 `json:"duration"`
	Size      int     `json:"size"`
	Status    int     `json:"status"`
}

func (entry *LogEntry) UserAgent() string {
	return entry.getRequestHeader("User-Agent")
}

func (entry *LogEntry) RemoteAddr() string {
	host, _, err := net.SplitHostPort(entry.Request.RemoteAddr)
	if err != nil {
		log.Println(err)
	}
	return host
}

func (entry *LogEntry) getRequestHeader(key string) string {
	values, exists := entry.Request.Headers[key]
	if !exists {
		return ""
	}
	return values[0]
}
