/**
 * @Author: mjzheng
 * @Description:
 * @File:  file_utils.go
 * @Version: 1.0.0
 * @Date: 2020/6/1 下午5:20
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadFile() []int {
	f, err := os.Open("/Users/mjzheng/Documents/cgi_proj/mj_http/bin/pkg_800.txt")
	if err != nil {
		fmt.Println("failed to open file ", err)
		return nil
	}
	defer f.Close()

	var ls []int
	rd := bufio.NewReader(f)
	for {
		line, _, err := rd.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		num, e := strconv.Atoi(string(line))
		if e != nil {
			fmt.Println(" error ", e)
			continue
		}
		ls = append(ls, num)
		//fmt.Println(num)
	}
	return ls
}

func GetMax() int {
	max := 0
	ls := ReadFile()
	for _, v := range ls {
		if v > max {
			max = v
		}
	}
	fmt.Println("max", max)
	return max
}

func ReadFileString() []string {
	f, err := os.Open("/Users/mjzheng/Documents/cgi_proj/mj_http/bin/pkg_800.txt")
	if err != nil {
		fmt.Println("failed to open file ", err)
		return nil
	}
	defer f.Close()

	var ls []string
	rd := bufio.NewReader(f)
	for {
		line, _, err := rd.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		ls = append(ls, string(line))
	}
	return ls
}
