package common

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

type idGenerator struct {
	node *snowflake.Node
}

func NewIdGenerator(now time.Time, machineId int64) (*idGenerator, error) {

	snowflake.Epoch = now.Unix()
	node, err := snowflake.NewNode(machineId)
	if err != nil {
		return nil, err
	}

	return &idGenerator{node: node}, nil
}

func (idg *idGenerator) GenID() int64 {
	return idg.node.Generate().Int64()
}
