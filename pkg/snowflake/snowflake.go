package snowflake

import (
	"errors"
	"sync"
	"time"
)

const (
	epoch int64 = 1704067200000

	timestampBits = 41
	nodeIDBits    = 10
	sequenceBits  = 12

	maxNodeID   = -1 ^ (-1 << nodeIDBits)   // 1023
	maxSequence = -1 ^ (-1 << sequenceBits) // 4095

	timestampShift = nodeIDBits + sequenceBits
	nodeIDShift    = sequenceBits
)

type snowflake struct {
	mu            sync.Mutex
	nodeID        int64
	lastTimestamp int64
	sequence      int64
}

var (
	snowflakeImpl     *snowflake
	snowflakeImplOnce sync.Once
)

func InitSnowflake(nodeID int64) error {
	if nodeID < 0 || nodeID > maxNodeID {
		panic(errors.New("node ID must be between 0 and maxNodeID"))
	}

	snowflakeImplOnce.Do(func() {
		snowflakeImpl = &snowflake{
			nodeID:        nodeID,
			lastTimestamp: 0,
			sequence:      0,
		}
	})

	return nil
}

func GenerateClientID() (int64, error) {
	if snowflakeImpl == nil {
		return 0, errors.New("snowflake not initialized")
	}
	return snowflakeImpl.generateID()
}

func (s *snowflake) generateID() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	timestamp := time.Now().UnixMilli()

	if timestamp < s.lastTimestamp {
		return 0, errors.New("clock moved backwards, refusing to generate ID")
	}

	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			for timestamp <= s.lastTimestamp {
				timestamp = time.Now().UnixMilli()
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastTimestamp = timestamp

	id := ((timestamp - epoch) << timestampShift) |
		(s.nodeID << nodeIDShift) |
		s.sequence

	return id, nil
}
