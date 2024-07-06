package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Based on Twitter Snowflake
// ID generator
// |0|timestamp|datacenterID|machineID|sequenceNumber
// |1|41|5|5|12| bits
// maximum 4096 sequence

func IdGenerator(hashTable *map[int]int) string {
	currentTimestamp := int(time.Now().UTC().UnixMilli())
	if count, ok := (*hashTable)[currentTimestamp]; ok {
		count++

		(*hashTable)[currentTimestamp] = count
		return idGeneratorFunc(currentTimestamp, count)
	} else {
		// sequence starts from 0
		(*hashTable)[currentTimestamp] = 0
		return idGeneratorFunc(currentTimestamp, 0)
	}
}

// poc only, potentially would have conflict with the current millisecond
func ResetHashTable(hashTable *map[int]int) {
	for {
		// fmt.Println(*hashTable)
		currentTimestamp := int(time.Now().UTC().UnixMilli())
		// Remove those not in the current timestamp
		for key := range *hashTable {
			if key != currentTimestamp {
				delete(*hashTable, key)
			}
		}

		// log.Println("Hashtable cleaned up.")

		// Sleep for 1 second
		time.Sleep(10 * time.Second)
	}
}

// unique, sortable, distributed without single point of failure
func idGeneratorFunc(currentTimestamp, sequenceNumber int) string {
	initialBit := "0"
	dataCenterId := os.Getenv("DATA_CENTER_ID")
	dataCenterId_int, _ := strconv.Atoi(dataCenterId)
	machineId := os.Getenv("MACHINE_ID")
	machineId_int, _ := strconv.Atoi(machineId)
	// currentTimestamp := time.Now().UTC().UnixMilli()

	//uniqueId
	timestamp_bin := fmt.Sprintf("%041b", currentTimestamp)
	dataCenterId_bin := fmt.Sprintf("%05b", dataCenterId_int)
	machineId_bin := fmt.Sprintf("%05b", machineId_int)
	sequenceNumber_bin := fmt.Sprintf("%012b", sequenceNumber)

	id_bin := fmt.Sprintf("%s%s%s%s%s", initialBit, timestamp_bin, dataCenterId_bin, machineId_bin, sequenceNumber_bin)
	id_int, _ := strconv.ParseInt(id_bin, 2, 64)
	// fmt.Println(id_bin)
	// fmt.Println(id_int)
	return fmt.Sprint(id_int)
}

type IdStruct struct {
	Timestamp      int64  `json:"unix_timestamp"`
	Datetime       string `json:"datetime"`
	DatacenterId   int64  `json:"datacenter_id"`
	MachineId      int64  `json:"machine_id"`
	SequenceNumber int64  `json:"sequence_number"`
	BinaryForm     string `json:"binary_form"`
}

func parseId(id string, idStruct *IdStruct) error {
	if len(id) != 64 {
		return errors.New("invalid id length")
	}
	// get binary string
	timestamp := id[1 : 1+41]
	datacenterId := id[42 : 42+5]
	machineId := id[47 : 47+5]
	sequenceNumber := id[52 : 52+12]

	// convert to integer
	timestamp_int, _ := strconv.ParseInt(timestamp, 2, 64)
	datacenterId_int, _ := strconv.ParseInt(datacenterId, 2, 64)
	machineId_int, _ := strconv.ParseInt(machineId, 2, 64)
	sequenceNumber_int, _ := strconv.ParseInt(sequenceNumber, 2, 64)

	idStruct.Timestamp = timestamp_int // Redundant, just for POC
	idStruct.Datetime = time.UnixMilli(timestamp_int).UTC().String()
	idStruct.DatacenterId = datacenterId_int
	idStruct.MachineId = machineId_int
	idStruct.SequenceNumber = sequenceNumber_int
	idStruct.BinaryForm = id // Redundant, just for POC
	// fmt.Println(timestamp)
	// fmt.Println(datacenterId)
	// fmt.Println(machineId)
	// fmt.Println(sequenceNumber)

	return nil
}
