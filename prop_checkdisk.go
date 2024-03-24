package main

import (
	"bytes"
	"fmt"
	"os"
	"syscall"
	"time"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

func checkDisk(diskPath string) {
	fmt.Printf("Checking disk under %s", diskPath)

	var stat syscall.Statfs_t
	info, err := os.Stat(diskPath)
	checkError(err)

	if !info.IsDir() {
		checkError(fmt.Errorf("Input '%s' is not a directory.\n", diskPath))
	}

	// Determine available disk space
	syscall.Statfs(diskPath, &stat)
	blockSize := uint64(stat.Bsize) // Size of each data block
	availableSpace := stat.Bavail * blockSize
	targetSize := availableSpace - GB*uint64(Opts.Spare)
	chunkSize := targetSize / blockSize * blockSize / uint64(Opts.DataPoints) // dataPoint Interval, Size of each data point
	debug(fmt.Sprintf("Target file size: %d bytes\n", targetSize), 1)
	fmt.Printf(" using %dMB with %d speed-measurement data points...\n", targetSize/MB, Opts.DataPoints)

	// prepare space/data once, before measuring
	speedData := make([]int, Opts.DataPoints)
	buffer := make([]byte, stat.Bsize)
	for i := range buffer {
		buffer[i] = (byte)(i % 512)
	}

	fileName := diskPath + "/test.dat"
	defer os.Remove(fileName)

	// Write pass
	measureWriteSpeed(speedData, fileName, buffer, blockSize, chunkSize)
	printSpeeds("Write speeds", speedData)

	wait := time.Duration(Opts.Wait) * time.Second
	fmt.Printf("\nWait after write before read for %v for sytem to take a breath...\n", wait)
	time.Sleep(wait)

	// Read pass
	fmt.Printf("\nRead back & verify...\n")
	measureReadSpeed(speedData, fileName, buffer, blockSize, chunkSize)
	printSpeeds("Read speeds", speedData)

}

func measureWriteSpeed(speedData []int, fileName string, buffer []byte, blockSize, chunkSize uint64) bool {

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return false
	}
	defer file.Close()

	// data points loop for each speed measurement
	for p := 0; p < Opts.DataPoints; p++ {
		// chunkSize loop to fill up each data point
		startTime := time.Now()
		for i := 0; i < int(chunkSize/blockSize); i++ {
			_, err := file.Write(buffer)
			if err != nil {
				fmt.Println("Error writing data:", err)
				return false
			}
		}
		file.Sync() // Flush to disk
		elapsed := time.Since(startTime)

		if Opts.KbSpeed {
			speedData[p] = (int(chunkSize)*10/KB*1000/int(elapsed.Milliseconds()) + 5) / 10
		} else {
			speedData[p] = (int(chunkSize)*10/MB*1000/int(elapsed.Milliseconds()) + 5) / 10
		}
		print(".")
	}
	print("\n")

	return true
}

func measureReadSpeed(speedData []int, fileName string, buffer []byte, blockSize, chunkSize uint64) bool {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return false
	}
	defer file.Close()

	// data points loop for each speed measurement
	data := make([]byte, blockSize)
	for p := 0; p < Opts.DataPoints; p++ {
		// chunkSize loop to fill up each data point
		startTime := time.Now()
		for i := 0; i < int(chunkSize/blockSize); i++ {
			n, err := file.Read(data)
			if err != nil {
				fmt.Println("Error writing data:", err)
				return false
			}

			if n != int(blockSize) || !bytes.Equal(data, buffer) {
				fmt.Printf("\nData mismatch at data point %d offset %d\n", p, i)
				return false
			}
		}
		elapsed := time.Since(startTime)

		if Opts.KbSpeed {
			speedData[p] = (int(chunkSize)*10/KB*1000/int(elapsed.Milliseconds()) + 5) / 10
		} else {
			speedData[p] = (int(chunkSize)*10/MB*1000/int(elapsed.Milliseconds()) + 5) / 10
		}
		print(".")
	}
	print("\n")

	return true
}

func printSpeeds(measureRType string, speedData []int) {
	fmt.Print(measureRType)
	if Opts.KbSpeed {
		fmt.Println(" (KB/s):")
	} else {
		fmt.Println(" (MB/s):")
	}
	fmt.Println(speedData)
}
