package main

import (
	"errors"
	"fmt"
	"time"
)

// SnowflakeId 雪花ID生成器
type SnowflakeId struct {
	initialTimestamp int64 // 初始时间
	sequenceBits     int64 // 序列号位数
	sequenceMask     int64 // 序列号掩码
	workerBits       int64 // 机器码位数
	workerShift      int64 // 机器码左移位数
	workerMax        int64 // 机器码最大值
	timestampBits    int64 // 时间戳位数
	timestampShift   int64 // 时间戳左移位数
	workerId         int64 // 机器码
	lastTimestamp    int64 // 时间戳
	sequenceNo       int64 // 序列号
}

// Generate 生成雪花ID
func (s *SnowflakeId) Generate() (int64, error) {
	if s.workerId > s.workerMax || s.workerId < 0 {
		return 0, errors.New(fmt.Sprintf("worker Id can't be greater than %d or less than 0", s.workerMax))
	}

	timestamp := s.msTimestamp()

	if timestamp < s.lastTimestamp {
		return 0, errors.New(fmt.Sprintf("Clock moved backwards. Refusing to generate id for %d milliseconds", s.lastTimestamp-timestamp))
	}

	if timestamp == s.lastTimestamp {
		s.sequenceNo = (s.sequenceNo + 1) & s.sequenceMask

		if s.sequenceNo == 0 {
			timestamp = s.waitMsTimestamp()
		}
	} else {
		s.sequenceNo = 0
	}

	s.lastTimestamp = timestamp

	return ((timestamp - s.initialTimestamp) << s.timestampShift) | (s.workerId << s.workerShift) | s.sequenceNo, nil
}

// Parse 解析生成的雪花ID
func (s *SnowflakeId) Parse(id int64) (int64, int64, int64) {
	sequenceNo := id & (-1 ^ (-1 << s.sequenceBits))

	workerId := (id & ((-1 << s.workerShift) ^ (-1 << s.timestampShift))) >> s.workerShift

	timestamp := (id & ((-1 << s.timestampShift) ^ (-1 << (s.timestampShift + s.timestampBits)))) >> s.timestampShift

	timestamp = timestamp + s.initialTimestamp

	return timestamp, workerId, sequenceNo
}

// msTimestamp 获取毫秒级的时间戳
func (s *SnowflakeId) msTimestamp() int64 {
	return time.Now().UnixMilli()
}

// waitMsTimestamp 阻塞并获取下一毫秒的时间戳
func (s *SnowflakeId) waitMsTimestamp() int64 {
	for {
		time.Sleep(time.Microsecond * 100)
		timestamp := s.msTimestamp()
		if timestamp > s.lastTimestamp {
			return timestamp
		}
	}
}

// NewSnowflakeId 获取雪花ID生成器实例
func NewSnowflakeId(workerId int64) *SnowflakeId {
	return &SnowflakeId{
		initialTimestamp: 1420041600000,
		sequenceBits:     12,
		sequenceMask:     -1 ^ (-1 << 12),
		workerBits:       10,
		workerShift:      12,
		workerMax:        -1 ^ (-1 << 10),
		timestampBits:    41,
		timestampShift:   22,
		workerId:         workerId,
		lastTimestamp:    0,
		sequenceNo:       0,
	}
}

func main() {

	// 调用方法示例
	s := NewSnowflakeId(100)

	for i := 0; i < 1000; i++ {
		if id, err := s.Generate(); err == nil {
			fmt.Println(id)
			fmt.Println(s.Parse(id))
			fmt.Println()
		}
	}

}
