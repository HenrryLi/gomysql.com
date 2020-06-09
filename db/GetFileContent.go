package db

import (
	"fmt"
	"io/ioutil"

	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/logger"
)

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

//读取文件内容为一个字符串数组
func GetFileContentAsStringLines(filePath string) []string {

	dbc, err := sql.Open("mysql", "test:test@(192.168.56.103:3306)/test")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("MySQL db is not connected")
	} else {
		fmt.Println("MySQL db is connected.")
	}

	defer dbc.Close()

	// make sure connection is available
	err = dbc.Ping()
	// fmt.Println(err)
	if err != nil {
		fmt.Println("MySQL db is not connected")
		fmt.Println(err.Error())
	}

	logger.Infof("get file content as lines: %v", filePath)
	s3 := []string{}
	// var result string
	buffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Errorf("read file: %v error: %v", filePath, err)
		// return result
	}

	s1 := string(buffer)

	s2 := strings.Split(s1, "\n\n")
	// fmt.Println(len(s2), s2)

	for _, lineStr := range s2 {
		s3 := strings.Split(lineStr, "\n")

		lenS3 := len(s3)
		s31 := s3[0:3]
		s32 := s3[3 : lenS3-3]
		s33 := s3[lenS3-3:]

		// fmt.Println(n, ":s31:", s31)
		// fmt.Println(n, ":s32:", s32)
		// fmt.Println(n, ":s33:", s33[0], s33[1], s33[2])

		// Prepare statement for inserting data
		stmtIns, err := dbc.Prepare("INSERT INTO aws_bill VALUES( ?,?,?, ? )") // ? = placeholder
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

		var rows Result

		for i := 0; i < len(s32)/2; i++ {
			// fmt.Println(i, ":", len(s32)/2, ":s32: ", s32[i*2+1], s32[i*2])
			_, err = stmtIns.Exec(nil, s31[0], s32[i*2+1], s32[i*2])
			if err != nil {
				fmt.Println("insert service cost err:", err)
			}

		}

		_, err = stmtIns.Exec(nil, s31[0], s33[0], s33[1]) // Insert the Taxes cost
		if err != nil {
			fmt.Println("insert taxes err:", err)
		}

		rows, err = stmtIns.Exec(nil, s31[0], "Total", s31[2]) // Insert month total cost
		fmt.Println("rows:", rows)
		if err != nil {
			fmt.Println("insert month err:", err)
		}

	}

	logger.Infof("get file content as lines: %v, size: %v", filePath, len(s3))
	return s3
}
