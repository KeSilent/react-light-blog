package utils

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

func GenID(nodeNumber int64) (int64, error) {
	// Create a new Node with a Node number of 1
	if nodeNumber == 0 {
		nodeNumber = 1
	}

	node, err := snowflake.NewNode(nodeNumber)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Generate a snowflake ID.
	id := node.Generate()
	return id.Int64(), nil
}
