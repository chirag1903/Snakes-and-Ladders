package main

import (
	"Users/chirag.soni/druva/snakesandladders/initconfig"
	"Users/chirag.soni/druva/snakesandladders/rungame"
	"Users/chirag.soni/druva/snakesandladders/service"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func main() {

	// parse input from yml
	config, err := initconfig.PrepareConfigFromYml()
	if err != nil {
		errorLog := fmt.Sprint("invalid config file ", err)
		fmt.Println(errorLog)
		os.Exit(1)
	}

	err = config.ValidateInputParams()
	if err != nil {
		errorLog := fmt.Sprint("invalid input argument ", err)
		fmt.Println(errorLog)
		os.Exit(1)
	}

	processManual := false
	var manualInputInfos []*initconfig.ManualInputInfo
	if processManual {
		manualInputInfos, err = processTextFile()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

	svcs := service.InitService(config)

	// do simulation
	err = rungame.StartGame(config, svcs, manualInputInfos)
	if err != nil {
		errorLog := fmt.Sprint("game simulation failed", err)
		fmt.Println(errorLog)
		os.Exit(1)
	}

}

func processTextFile() ([]*initconfig.ManualInputInfo, error) {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, errors.Errorf("failed to open file: %s", err)
	}
	defer file.Close()

	var inputInfos []*initconfig.ManualInputInfo

	// Use a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line by space
		parts := strings.Fields(line)
		if len(parts) < 2 {
			return nil, errors.Errorf("invalid line format: %s", line)
		}

		// Extract player name
		name := parts[0]

		// Extract dice values
		var dice []int
		for _, part := range parts[1:] {
			diceValue, err := strconv.Atoi(part)
			if err != nil {
				return nil, errors.Errorf("invalid dice value: %s", part)
			}
			dice = append(dice, diceValue)
		}

		// Create a player and append to the slice
		inputInfo := &initconfig.ManualInputInfo{
			PlayerName: name,
			RollValues: dice,
		}
		inputInfos = append(inputInfos, inputInfo)
	}

	// Check for errors in scanning
	if err := scanner.Err(); err != nil {
		return nil, errors.Errorf("error reading file: %s", err)
	}

	return inputInfos, nil
}
