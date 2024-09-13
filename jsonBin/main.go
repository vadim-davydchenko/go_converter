package main

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func createBin(id, name string, private bool) *Bin {
	newBin := &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
	return newBin
}

func createBinList(id, name string, private bool) *BinList {
	newBinList := &BinList{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
	return newBinList
}

func main() {
}
