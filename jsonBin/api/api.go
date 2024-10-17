package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"jsonBin/config"
)

type CreateBinResponse struct {
	Metadata struct {
		ID string `json:"id"`
	} `json:"metadata"`
}

func CreateBin(cfg *config.Config, data []byte, name string) (string, error) {
	url := "https://api.jsonbin.io/v3/b"

	payload := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", cfg.MasterKey)
	req.Header.Set("X-Bin-Name", name)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error executing request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected response status: %v", resp.StatusCode)
	}

	var result CreateBinResponse
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&result)
	if err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	return result.Metadata.ID, nil
}

func UpdateBin(cfg *config.Config, binID string, data []byte) error {
	url := fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", binID)

	payload := bytes.NewBuffer(data)

	req, err := http.NewRequest("PUT", url, payload)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", cfg.MasterKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error executing request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %v", resp.Status)
	}

	return nil
}

func DeleteBin(cfg *config.Config, binID string) error {
	url := fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", binID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("X-Master-Key", cfg.MasterKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error executing request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %v", resp.Status)
	}

	return nil
}

func GetBin(cfg *config.Config, binID string) (string, error) {
	url := fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", binID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("X-Master-Key", cfg.MasterKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error executing request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected response status: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	return string(body), nil
}
