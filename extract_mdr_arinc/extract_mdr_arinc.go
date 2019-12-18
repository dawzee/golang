package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

func main() {

	// Create output file
	outFile, err := os.Create("output.csv")
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(outFile,"Timestamp,Channel,Raw")

	// Just count up from 1 and stop when we fail to open a file
	for num := 1; num < 1000; num++ {

		nextFilename := fmt.Sprintf("A429_%d.dat",num)

		// Open the file and create a reader for it
		nextFile, err := os.Open(nextFilename)

		if err != nil {
			break
		}

		fmt.Println("Processing ",nextFilename)

		bufr := bufio.NewReader(nextFile)

		// Read in the header and check it's a 1 as expected
		var header uint32
		err = binary.Read(bufr, binary.LittleEndian, &header)
		if err != nil {
			panic(err)
		}
		if header != 1 {
			panic("Header mismatch")
		}

		var data struct {
			Timestamp int64
			Channel uint8
			ArincWord uint32
		}

		for {
			// Read in the next packet
			err = binary.Read(bufr, binary.LittleEndian, &data)
			if err != nil {
				break
			}
			// Convert the timestamp to nanoseconds, then into a time struct
			data.Timestamp -= 116444736000000000
			data.Timestamp *= 100
			timeVal := time.Unix(0,data.Timestamp)

			fmt.Fprintf(outFile,"%s,%d,%X\n",timeVal,data.Channel,data.ArincWord)
		}
		nextFile.Close()
	}

	outFile.Close()
}
