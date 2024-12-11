package Model

type Print struct {
	Komoditi string `json:"komoditi"`
	Kemasan string `json:"kemasan"`
	Harga int `json:"harga"`
	StockAwal int `json:"stock_awal"`
	StockTambahan int `json:"stock_tambahan"`
	Terjual int `json:"terjual"`
	Sisa int `json:"sisa"`
	Hasil int `json:"hasil"`
}

type PrintOut struct {
	Komoditi	string `json:"komoditi"`
	Deskripsi	string 
	Jumlah		int
}

func DataPrint (salesReport []Sales_Report, productReport []ProductReportOpr) ([]Print, []PrintOut, error) {
	var data []Print
	var dataOut []PrintOut

	for _, s := range salesReport {
		data = append(data, Print{
			Komoditi: s.Komoditi,
			StockAwal: s.Stockawal,
			Terjual: s.Terjual,
			Hasil: s.Price,
		})
	}

	allProduct,err := GetAllProduct()
	if err != nil {
		return nil, nil, err
	}

	for _, a := range allProduct {
		for i, d := range data {
			if a.Name == d.Komoditi {
				data[i].Kemasan = a.Packagesize
				data[i].Harga = a.Price
				data[i].Sisa = a.Stock
			}
		}
	}

	for i, d := range data {
		for _, p := range productReport {
			if d.Komoditi == p.Name {
				if p.Action == "Masuk" {
					data[i].StockTambahan =+ p.Quantity
				}
			}
		}
	}

	for _, p := range productReport {
		if p.Action == "Keluar" {
			dataOut = append(dataOut, PrintOut{
				Komoditi: p.Name,
				Deskripsi: p.Description,
				Jumlah: p.Quantity,
			})
		}
	}

	return data, dataOut, nil
}
