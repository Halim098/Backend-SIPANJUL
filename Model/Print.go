package Model

import (
	"Sipanjul/Database"
	"errors"
	"sync"
	"time"
)

type Print struct {
	Komoditi      string `json:"komoditi"` //done
	Kemasan       string `json:"kemasan"` // done
	Harga         int    `json:"harga"` // done
	StockAwal     int    `json:"stock_awal"`
	StockTambahan int    `json:"stock_tambahan"` 
	Terjual       int    `json:"terjual"` 
	Sisa          int    `json:"sisa"`
	Hasil         int    `json:"hasil"`
}

type Printsales struct {
	Komoditi string `json:"komoditi"`
	StockAwal int `json:"stock_awal"`
	StockAkhir int `json:"stock_akhir"`
	Terjual int `json:"terjual"`
	Total int `json:"total"`
	Date time.Time `json:"date"`
}

type PrintProduct struct {
	Komoditi string `json:"komoditi"`
	Kemasan string `json:"kemasan"`
	Harga int `json:"harga"`
	Stock int `json:"stock"`
}

type PrintStock struct {
	Komoditi string `json:"komoditi"`
	Action string `json:"action"`
	Quantity int `json:"quantity"`
	Description string `json:"description"`
}

type PrintOut struct {
	Komoditi  string `json:"komoditi"`
	Deskripsi string
	Jumlah    int
}

func DataPrint(startdate, enddate string, oprid uint) ([]Print, []PrintOut, error) {
	var product []PrintProduct
	var errProduct error
	
	var sales []Printsales
	var errSales error

	var out []PrintStock
	var errOut error

	// menunggu gorutine
	var wg1 sync.WaitGroup

	// Goroutine pertama untuk menghitung CurrentValue dan OldValue
	wg1.Add(3)

	go func() {
		defer wg1.Done()
		err := Database.Database.Raw(`
		SELECT
		name as komoditi,
		packagesize as kemasan,
		price as harga,
		stock as stock

		FROM products WHERE opr_id = ?`, oprid).Scan(&product)
	
		if err.Error != nil {
			errProduct = err.Error
			return
		}

		if err.RowsAffected == 0 {
			errProduct = errors.New("data tidak ditemukan")
			return
		}
	}()

	go func() {
		defer wg1.Done()
		err := Database.Database.Raw(`
		SELECT
		p.name as komoditi,
		sd.stock_awal as stock_awal,
		sd.stock_akhir as stock_akhir,
		sd.quantity as terjual,
		sd.total as total,
		s.date as date

		FROM sales_details sd
		JOIN sales s ON sd.sales_id = s.id
		JOIN products p ON sd.prod_id = p.id
		WHERE s.date >= ? AND s.date < ? AND s.opr_id = ?`, startdate, enddate, oprid).Scan(&sales)
		
		if err.Error != nil {
			errSales = err.Error
			return
		}

		if err.RowsAffected == 0 {
			errSales = errors.New("data tidak ditemukan")
			return
		}
	}()

	go func() {
		defer wg1.Done()
		err := Database.Database.Raw(`
		SELECT
		p.name as komoditi,
		r.action as action,
		r.quantity as quantity,
		r.description as description

		FROM product_reports r
		JOIN products p ON r.prod_id = p.id
		WHERE r.date >= ? AND r.date < ? AND p.opr_id = ?`, startdate, enddate, oprid).Scan(&out)

		if err.Error != nil {
			errOut = err.Error
			return
		}

		if err.RowsAffected == 0 {
			errOut = errors.New("data tidak ditemukan")
			return
		}
	}()

	wg1.Wait()

	if errProduct != nil {
		return nil, nil, errProduct
	}

	if errSales != nil {
		return nil, nil, errSales
	}

	if errOut != nil {
		return nil, nil, errOut
	}

	var print []Print
	var printOut []PrintOut

	// Map untuk menyimpan data gabungan
	merged := map[string]Printsales{}

	for _, sale := range sales {
		// Cek apakah komoditi sudah ada di map
		if existing, found := merged[sale.Komoditi]; found {
			// Update data jika sudah ada
			existing.Terjual += sale.Terjual
			existing.Total += sale.Total
			existing.StockAkhir = sale.StockAkhir
			if sale.Date.Before(existing.Date) {
				existing.StockAwal = sale.StockAwal
				existing.Date = sale.Date
			}
			merged[sale.Komoditi] = existing
		} else {
			// Tambahkan data baru jika belum ada
			merged[sale.Komoditi] = sale
		}
	}

	// Convert map kembali ke slice
	combinedSales := make([]Printsales, 0, len(merged))
	for _, v := range merged {
		combinedSales = append(combinedSales, v)
	}

	for _, v := range product {
		var p Print
		p.Komoditi = v.Komoditi
		p.Kemasan = v.Kemasan
		p.Harga = v.Harga
		p.StockTambahan = 0
		p.Sisa = v.Stock

		for _, s := range combinedSales {
			if s.Komoditi == v.Komoditi {
				p.StockAwal = s.StockAwal
				p.Terjual = s.Terjual
				p.Hasil = s.Total
			}
		}

		for _, o := range out {
			if o.Komoditi == v.Komoditi {
				if o.Action == "penambahan" {
					p.StockTambahan += o.Quantity
				} 
			}

			if o.Action == "pengurangan" {
				printOut = append(printOut, PrintOut{
					Komoditi:  o.Komoditi,
					Deskripsi: o.Description,
					Jumlah:    o.Quantity,
				})
			}
		}

		print = append(print, p)
	}

	return print, printOut, nil
}
