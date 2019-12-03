package common

import (
	mapset "github.com/deckarep/golang-set"
)

const (
	//ReportFilePrefix reportfile 前缀
	ReportFilePrefix string = "report-"
	//ReportFileSuffix reportfile 后缀
	ReportFileSuffix string = ".txt"
)

var (
	//SupportFeatures 支持功能常量
	SupportFeatures = []string{"ping", "traceroute"}

	// UpdateCmd redis更新语句
	UpdateCmd mapset.Set = mapset.NewSet("DEL",
		"EXPIRE",
		"EXPIREAT",
		"MIGRATE",
		"MOVE",
		"PERSIST",
		"PEXPIREAT",
		"RENAME",
		"RENAMENX",
		"APPEND",
		"BITOP",
		"DECR",
		"DECRBY",
		"GETSET",
		"INCR",
		"INCRBY",
		"INCRBYFLOAT",
		"MSET",
		"MSETNX",
		"PSETEX",
		"SET",
		"SETBIT",
		"SETEX",
		"SETNX",
		"SETRANGE",
		"HDEL",
		"HINCRBY",
		"HINCRBYFLOAT",
		"HMSET",
		"HSETNX",
		"BLPOP",
		"BRPOP",
		"BRPOPLPUSH",
		"LINSERT",
		"LPOP",
		"LPUSH",
		"LPUSHX",
		"LREM",
		"LSET",
		"LTRIM",
		"RPOP",
		"RPOPLPUSH",
		"RPUSH",
		"RPUSHX",
		"SADD",
		"SDIFFSTORE",
		"SINTERSTORE",
		"SMOVE",
		"SPOP",
		"SREM",
		"SUNIONSTORE",
		"ZADD",
		"ZINCRBY",
		"ZREM",
		"ZREMRANGEBYRANK",
		"ZREMRANGEBYSCORE",
		"ZUNIONSTORE",
		"ZINTERSTORE",
		"DISCARD",
		"EXEC",
		"MULTI",
		"UNWATCH",
		"WATCH",
		"EVAL",
		"EVALSHA",
		"SCRIPT FLUSH",
		"SCRIPT KILL",
		"SCRIPT LOAD",
		"SELECT",
		"FLUSHALL",
		"FLUSHDB ")
)
