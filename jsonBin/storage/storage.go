package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

type LocalBinInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type StorageService interface {
	SaveBinInfo(id, name string) error
	UpdateLocalBinInfo(id, filePath string) error
	DeleteLocalBinInfo(id string) error
	ListBins() ([]LocalBinInfo, error)
}

type FileStorage struct{}

func (fs *FileStorage) SaveBinInfo(id, name string) error {
	file, err := os.OpenFile("local_bins.json", os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var bins []LocalBinInfo
	err = json.NewDecoder(file).Decode(&bins)
	if err != nil && err.Error() != "EOF" {
		return fmt.Errorf("error decoding existing data: %v", err)
	}

	bins = append(bins, LocalBinInfo{ID: id, Name: name})

	file.Seek(0, 0)
	err = json.NewEncoder(file).Encode(bins)
	if err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}

	return nil
}

func (fs *FileStorage) UpdateLocalBinInfo(id, filePath string) error {
	file, err := os.OpenFile("local_bins.json", os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var bins []LocalBinInfo
	err = json.NewDecoder(file).Decode(&bins)
	if err != nil && err.Error() != "EOF" {
		return fmt.Errorf("error decoding existing data: %v", err)
	}

	for i, bin := range bins {
		if bin.ID == id {
			bins[i].Name = filePath
			break
		}
	}

	file.Seek(0, 0)
	err = json.NewEncoder(file).Encode(bins)
	if err != nil {
		return fmt.Errorf("error updating local file: %v", err)
	}

	return nil
}

func (fs *FileStorage) DeleteLocalBinInfo(id string) error {
	file, err := os.OpenFile("local_bins.json", os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var bins []LocalBinInfo
	err = json.NewDecoder(file).Decode(&bins)
	if err != nil && err.Error() != "EOF" {
		return fmt.Errorf("error decoding existing data: %v", err)
	}

	for i, bin := range bins {
		if bin.ID == id {
			bins = append(bins[:i], bins[i+1:]...)
			break
		}
	}

	file.Truncate(0)
	file.Seek(0, 0)
	err = json.NewEncoder(file).Encode(bins)
	if err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}

	return nil
}

func (fs *FileStorage) ListBins() ([]LocalBinInfo, error) {
	file, err := os.Open("local_bins.json")
	if err != nil {
		return nil, fmt.Errorf("error opening local bins file: %v", err)
	}
	defer file.Close()

	var bins []LocalBinInfo
	err = json.NewDecoder(file).Decode(&bins)
	if err != nil && err.Error() != "EOF" {
		return nil, fmt.Errorf("error decoding bin data: %v", err)
	}

	return bins, nil
}
