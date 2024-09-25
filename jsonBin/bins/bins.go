package bins

import "time"

type Bin struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

type BinList struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Bins      []Bin     `json:"bins"`
}

func createBin(id, name string, private bool) *Bin {
	return &Bin{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

func createBinList(id, name string, private bool) *BinList {
	return &BinList{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
		Bins:      []Bin{},
	}
}

func (bl *BinList) AddBin(bin Bin) {
	bl.Bins = append(bl.Bins, bin)
}
