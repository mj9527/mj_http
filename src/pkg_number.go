/**
 * @Author: mjzheng
 * @Description:
 * @File:  pkg_number.go
 * @Version: 1.0.0
 * @Date: 2020/6/1 下午5:19
 */

package main

import (
	"fmt"
	redis_cluster "github.com/gitstliu/go-redis-cluster"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

const (
	CMD_BITMAP = 1
	CMD_SET    = 2
	CMD_HASH   = 3
)

func SendPage(ls []string, key string, c redis.Conn, cmd int) {
	if cmd == CMD_SET {
		args := []interface{}{key}
		count := len(ls)
		for i := 0; i < count; i++ {
			args = append(args, ls[i])
		}
		c.Send("sadd", args...)
	} else if cmd == CMD_BITMAP {
		count := len(ls)
		for i := 0; i < count; i++ {
			c.Send("SETBIT", key, ls[i], 1)
		}
	} else if cmd == CMD_HASH {
		args := []interface{}{key}
		count := len(ls)
		for i := 0; i < count; i++ {
			args = append(args, ls[i])
			args = append(args, 1)
		}
		c.Send("HMSET", args...)
	}
}

func WritePkgWithPipeline(key string, cmd int) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379") // use redigo
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	start := time.Now()

	ls := ReadFileString()
	fmt.Println("file len", len(ls))

	pageSize := 50
	page := len(ls) / pageSize
	count := 0
	for i := 0; i < page; i++ {
		SendPage(ls[i*pageSize:(i+1)*pageSize], key, c, cmd)
		count++
	}

	if len(ls)%pageSize != 0 {
		SendPage(ls[page*pageSize:], key, c, cmd)
	}
	c.Flush()

	for i := 0; i < count; i++ {
		_, err := c.Receive()
		if err != nil {
			fmt.Println("recv error", err)
		}
	}

	elapsed := time.Since(start)
	fmt.Println(cmd, "cost ", elapsed)
}

// pkg_string
func WritePkgWithNoPipeline(key string, cmd int) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379") // use redigo
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	start := time.Now()
	ls := ReadFileString()
	fmt.Println("file len", len(ls))
	for _, v := range ls {
		if cmd == CMD_SET {
			_, err = c.Do("sadd", key, v)
		} else if cmd == CMD_BITMAP {
			_, err = c.Do("SETBIT", key, v, 1)
		} else if cmd == CMD_HASH {
			_, err = c.Do("HSET", key, v, 1)
		}

		if err != nil {
			fmt.Println("redis set failed:", err)
			break
		}
	}
	elapsed := time.Since(start)
	fmt.Println(cmd, "cost ", elapsed)
}

func ReadPkgWithPipeline(key string, cmd int) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379") // use redigo
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	start := time.Now()

	for i := 0; i < 100000; i++ {
		if cmd == CMD_SET {
			err := c.Send("SISMEMBER", key, i)
			if err != nil {
				fmt.Println("send error ", err)
			}
		} else if cmd == CMD_HASH {
			//err := c.Send("HGET", key, i)
			err := c.Send("HEXISTS", key, i)
			if err != nil {
				fmt.Println("send error ", err)
			}
		} else if cmd == CMD_BITMAP {
			err := c.Send("GETBIT", key, i)
			if err != nil {
				fmt.Println("send error ", err)
			}
		}
	}

	c.Flush()
	for i := 0; i < 100000; i++ {
		_, err := c.Receive()
		if err != nil {
			fmt.Println("recv error", err)
		}
	}

	elapsed := time.Since(start)
	fmt.Println(cmd, "pipeline cost ", elapsed)
}

func ReadWithNoPipeline() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379") // use redigo
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	start := time.Now()

	for i := 0; i < 100000; i++ {
		username, err := c.Do("SISMEMBER", "pkg_803", "1002334872")
		if err != nil || username == nil {
			fmt.Println("redis get failed:", err)
		} else {
			_, err := redis.Int64(username, err)
			if err != nil {
				fmt.Println("failed to query")
			}

		}
	}

	elapsed := time.Since(start)
	fmt.Println("no pipeline cost ", elapsed)
}

func ReadFromCluster() {
	cluster, _ := redis_cluster.NewCluster(
		&redis_cluster.Options{
			StartNodes:   []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003", "127.0.0.1:7004", "127.0.0.1:7005"},
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})

	start := time.Now()
	batch := cluster.NewBatch()
	batch.Put("get", "name")
	batch.Put("get", "hi")
	batch.Put("get", "sex")

	reply, err := cluster.RunBatch(batch)
	if err != nil {
		log.Fatalf("RunBatch error: %s", err.Error())
	}

	for i := 0; i < 3; i++ {
		var resp string
		reply, err = redis.Scan(reply, &resp)
		if err != nil {
			log.Fatalf("RunBatch error: %s", err.Error())
		}

		fmt.Printf("[%d] return: %s\n", i, resp)
	}
	elapsed := time.Since(start)
	fmt.Println("cluster cost", elapsed)
}


func GetStrKey() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379") // use redigo
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	start := time.Now()

	data, err := c.Do("get", "mj")
	fmt.Println("data", data, "err", err)
	if err != nil {
		fmt.Println("failed to get mj", err)
	}

	if data == nil {
		fmt.Println("empty data")
	}

	b, ok := data.([]byte)

	fmt.Println("tansfer", string(b), ok)

	//if data == "" {
	//	fmt.Println("data is nil");
	//}


	elapsed := time.Since(start)
	fmt.Println("no pipeline cost ", elapsed)
}