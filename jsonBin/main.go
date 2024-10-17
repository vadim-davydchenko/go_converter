package main

import (
	"flag"
	"fmt"
	"log"

	"jsonBin/api"
	"jsonBin/config"
	"jsonBin/files"
	"jsonBin/storage"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Printf("error for initiliazation configuration: %v", err)
	}
	fmt.Printf("Uploaded key: %s\n", cfg.Key)

	createFlag := flag.Bool("create", false, "Create new bin")
	updateFlag := flag.Bool("update", false, "Update current bin")
	deleteFlag := flag.Bool("delete", false, "Delete bin")
	getFlag := flag.Bool("get", false, "Get bin by ID")
	listFlag := flag.Bool("list", false, "List all bins from local storage")
	filePath := flag.String("file", "", "Path to JSON file")
	binName := flag.String("name", "", "Name of the bin")
	binID := flag.String("id", "", "ID bin (for update)")
	flag.Parse()

	var storageService storage.StorageService = &storage.FileStorage{}

	if *createFlag {
		if *filePath == "" || *binName == "" {
			fmt.Println("You need to pass the file and name to create the bin")
		}

		if !files.IsJSONFile(*filePath) {
			fmt.Println("File don't have extension .json")
		}

		data, err := files.ReadFile(*filePath)
		if err != nil {
			fmt.Printf("Error for reading file: %v", err)
		}

		binID, err := api.CreateBin(cfg, data, *binName)
		if err != nil {
			fmt.Printf("Error creating bin via API: %v", err)
		}
		fmt.Printf("Bin successfully created with ID: %s\n", binID)

		err = storageService.SaveBinInfo(binID, *binName)
		if err != nil {
			fmt.Printf("Error saving bin information locally: %v", err)
		}
		fmt.Println("Bin information saved locally")
		return
	}

	if *updateFlag {
		if *filePath == "" || *binID == "" {
			fmt.Println("Need pointer file and ID for update bin")
		}

		if !files.IsJSONFile(*filePath) {
			fmt.Println("File don't have extension .json")
		}

		data, err := files.ReadFile(*filePath)
		if err != nil {
			fmt.Printf("Error for reading file: %v", err)
		}

		err = api.UpdateBin(cfg, *binID, data)
		if err != nil {
			fmt.Printf("Error for update bin: %v", err)
		}
		fmt.Println("Bin successfully updated")

		err = storageService.UpdateLocalBinInfo(*binID, *filePath)
		if err != nil {
			log.Fatalf("Error updating bin info locally: %v", err)
		}
		fmt.Println("Local bin information updated")
		return
	}

	if *deleteFlag {
		if *binID == "" {
			fmt.Printf("You need to pass the bin ID for deleting")
		}

		err = api.DeleteBin(cfg, *binID)
		if err != nil {
			fmt.Printf("Error deleting bin: %v", err)
		}
		fmt.Println("Bin successfully deleted from API")

		err = storageService.DeleteLocalBinInfo(*binID)
		if err != nil {
			fmt.Printf("Error deleting bin info locally: %v", err)
		}
		fmt.Println("Local bin information deleted")
		return
	}

	if *getFlag {
		if *binID == "" {
			fmt.Printf("You need to pass the bin ID to get bin\n")
			return
		}

		binData, err := api.GetBin(cfg, *binID)
		if err != nil {
			fmt.Printf("Error getting bin: %v\n", err)
			return
		}

		fmt.Printf("Bin data: %s\n", binData)
		return
	}

	if *listFlag {
		bins, err := storageService.ListBins()
		if err != nil {
			fmt.Printf("Error reading bin list: %v\n", err)
			return
		}

		fmt.Println("Bin List:")
		for _, bin := range bins {
			fmt.Printf("ID: %s, Name: %s\n", bin.ID, bin.Name)
		}
		return
	}
	fmt.Println("Use --create, --update, --delete, --get, or --list to perform actions on a bin")
}
