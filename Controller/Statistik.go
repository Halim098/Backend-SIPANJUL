package Controller

import (
	"Sipanjul/Model"
	"time"
)

func Statistik(id uint) ([]Model.Statistik,[]Model.Statistik,[]Model.Statistik,[]Model.Statistik, error) {
	var harian []Model.Statistik 
	var mingguan []Model.Statistik 
	var bulanan []Model.Statistik 
	var tahunan []Model.Statistik 

	var statistik Model.Statistik

	now := time.Now()

	enddate := now.Format("2006-01-02")
	hariandate := now.AddDate(0, 0, -1).Format("2006-01-02")
	mingguandate := now.AddDate(0, 0, -7).Format("2006-01-02")
	bulanandate := now.AddDate(0, -1, 0).Format("2006-01-02")
	tahunandate := now.AddDate(-1, 0, 0).Format("2006-01-02")

	dataharian,err := Model.GetSalesDetail(id, hariandate, enddate)
	if err != nil {
		return harian, mingguan, bulanan, tahunan ,err
	}

	datamingguan,err := Model.GetSalesDetail(id, mingguandate, enddate)
	if err != nil {
		return harian, mingguan, bulanan, tahunan, err
	}

	databulanan,err := Model.GetSalesDetail(id, bulanandate, enddate)
	if err != nil {
		return harian, mingguan, bulanan, tahunan, err
	}

	datatahunan,err := Model.GetSalesDetail(id, tahunandate, enddate)
	if err != nil {
		return harian, mingguan, bulanan, tahunan, err
	}

	for _, v := range dataharian {
		statistik = Model.Statistik{
			Komoditi : v.Komoditi,
			Harga : v.Price,
		}
		harian = append(harian, statistik)
	}

	for _, v := range datamingguan {
		statistik = Model.Statistik{
			Komoditi : v.Komoditi,
			Harga : v.Price,
		}
		mingguan = append(mingguan, statistik)
	}

	for _, v := range databulanan {
		statistik = Model.Statistik{
			Komoditi : v.Komoditi,
			Harga : v.Price,
		}
		bulanan = append(bulanan, statistik)
	}

	for _, v := range datatahunan {
		statistik = Model.Statistik{
			Komoditi : v.Komoditi,
			Harga : v.Price,
		}
		tahunan = append(tahunan, statistik)
	}

	return harian, mingguan, bulanan, tahunan ,nil
}