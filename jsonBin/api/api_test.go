package api

import (
	"jsonBin/config"
	"testing"
)

func setupConfig(t *testing.T) *config.Config {
	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatalf("Error initializing config: %v", err)
	}
	return cfg
}

func TestCreateBin(t *testing.T) {
	cfg := setupConfig(t)

	data := []byte(`{
		"accounts": [
			{
				"login": "test_user@example.com",
				"password": "securepassword123",
				"url": "http://example.com",
				"createdAt": "2024-10-16T12:00:00+00:00",
				"updatedAt": "2024-10-16T12:30:00+00:00"
			}
		]
	}`)

	binName := "test-bin"

	binID, err := CreateBin(cfg, data, binName)
	if err != nil {
		t.Fatalf("Failed to create bin: %v", err)
	}
	t.Logf("Bin created with ID: %s", binID)

	t.Cleanup(func() {
		err := DeleteBin(cfg, binID)
		if err != nil {
			t.Errorf("Failed to delete bin during cleanup: %v", err)
		}
	})
}

func TestUpdateBin(t *testing.T) {
	cfg := setupConfig(t)

	data := []byte(`{
		"accounts": [
			{
				"login": "test_user@example.com",
				"password": "securepassword123",
				"url": "http://example.com",
				"createdAt": "2024-10-16T12:00:00+00:00",
				"updatedAt": "2024-10-16T12:30:00+00:00"
			}
		]
	}`)

	binID, err := CreateBin(cfg, data, "test-bin")
	if err != nil {
		t.Fatalf("Failed to create bin: %v", err)
	}

	t.Cleanup(func() {
		err := DeleteBin(cfg, binID)
		if err != nil {
			t.Errorf("Failed to delete bin during cleanup: %v", err)
		}
	})

	updateData := []byte(`{
		"accounts": [
			{
				"login": "updated_user@example.com",
				"password": "updatedpassword456",
				"url": "http://updated.com",
				"createdAt": "2024-10-16T12:00:00+00:00",
				"updatedAt": "2024-10-16T12:30:00+00:00"
			}
		]
	}`)

	err = UpdateBin(cfg, binID, updateData)
	if err != nil {
		t.Fatalf("Failed to update bin: %v", err)
	}
	t.Logf("Bin with ID %s updated successfully", binID)
}

func TestGetBin(t *testing.T) {
	cfg := setupConfig(t)

	data := []byte(`{
		"accounts": [
			{
				"login": "test_user@example.com",
				"password": "securepassword123",
				"url": "http://example.com",
				"createdAt": "2024-10-16T12:00:00+00:00",
				"updatedAt": "2024-10-16T12:30:00+00:00"
			}
		]
	}`)

	binID, err := CreateBin(cfg, data, "test-bin")
	if err != nil {
		t.Fatalf("Failed to create bin: %v", err)
	}

	t.Cleanup(func() {
		err := DeleteBin(cfg, binID)
		if err != nil {
			t.Errorf("Failed to delete bin during cleanup: %v", err)
		}
	})

	binData, err := GetBin(cfg, binID)
	if err != nil {
		t.Fatalf("Failed to get bin: %v", err)
	}
	t.Logf("Bin data: %s", binData)
}

func TestDeleteBin(t *testing.T) {
	cfg := setupConfig(t)

	data := []byte(`{
		"accounts": [
			{
				"login": "test_user@example.com",
				"password": "securepassword123",
				"url": "http://example.com",
				"createdAt": "2024-10-16T12:00:00+00:00",
				"updatedAt": "2024-10-16T12:30:00+00:00"
			}
		]
	}`)

	binID, err := CreateBin(cfg, data, "test-bin")
	if err != nil {
		t.Fatalf("Failed to create bin: %v", err)
	}

	err = DeleteBin(cfg, binID)
	if err != nil {
		t.Fatalf("Failed to delete bin: %v", err)
	}
	t.Logf("Bin with ID %s deleted successfully", binID)
}
