/*
 * @Author: Yang
 * @Date: 2025-03-04 16:05:13
 * @Description: 请填写简介
 */
package utils

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

func GenID(nodeNumber int64) int64 {
	// Create a new Node with a Node number of 1
	if nodeNumber == 0 {
		nodeNumber = 1
	}

	node, err := snowflake.NewNode(nodeNumber)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	// Generate a snowflake ID.
	id := node.Generate()
	return id.Int64()
}
