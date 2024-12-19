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

	stock,err := GetAllProduct()
	if err != nil {
		return nil, nil, err
	}

	for _, v := range salesReport{
		var temp Print
		temp.Komoditi = v.Komoditi
		temp.Terjual = v.Terjual
		temp.StockAwal = v.Stockawal
		temp.Sisa = v.Stockakhir
		temp.Hasil = v.Price
		data = append(data, temp)
	}

	for _, v := range stock {
		for i , val := range data {
			if v.Name == val.Komoditi {
				data[i].Kemasan = v.Packagesize
				data[i].Harga = v.Price
			}
		}
	}

	for _, v := range productReport {
		for i, val := range data {
			if v.Name == val.Komoditi {
				if v.Action == "penambahan" {
					data[i].StockTambahan = v.Quantity
				}
				if v.Action == "pengurangan" {
					var temp PrintOut
					temp.Komoditi = v.Name
					temp.Deskripsi = v.Description
					temp.Jumlah = v.Quantity
					dataOut = append(dataOut, temp)
				}
			}
		}
	}
	return data, dataOut, nil
}
