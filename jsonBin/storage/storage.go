package storage

import (
	"encoding/json"
	"errors"
	"jsonBin/bins"
	"os"
)

func SaveBinListFile(filePath string, binList *bins.BinList) error {
	data, err := json.Marshal(binList)
	if err != nil {
		return errors.New("error for marhsal JSON")
	}

	file, err := os.Create(filePath)
	if err != nil {
		return errors.New("error in the creating file")
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return errors.New("error in the writing file")
	}
	return nil
}

func LoadBinListFile(filePath string) (*bins.BinList, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("error in the opening file")
	}
	defer file.Close()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("error in the reading file")
	}

	var binList bins.BinList
	err = json.Unmarshal(data, &binList)
	if err != nil {
		return nil, errors.New("error in the unmarshal JSON")
	}
	return &binList, nil
}
