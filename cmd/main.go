package main

import (
	"bufio"
	"os"
	"parking_lot/internal/commands"
	"parking_lot/internal/model"
	"parking_lot/internal/pkg/log"
)

func main() {

	var parkingLot model.ParkingLot

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	outF, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer outF.Close()
	writer := bufio.NewWriter(outF)
	log.Init(writer)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		cmd, err := commands.NewCommandFactory(line)
		if err != nil {
			log.Write(err.Error() + "\n")
			continue
		}

		if err := cmd.Parse(); err != nil {
			log.Write(err.Error() + "\n")
			continue
		}

		parkingLot, err = cmd.Execute(parkingLot)
		if err != nil {
			log.Write(err.Error() + "\n")
			continue
		}
	}

	if err := writer.Flush(); err != nil {
		panic(err)
	}
}
