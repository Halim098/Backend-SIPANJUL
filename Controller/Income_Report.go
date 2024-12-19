package Controller

import (
	"Sipanjul/Model"
	"sync"
	"time"
)

func IncomeReport(id uint) (Model.IncomeReport, Model.IncomeReport, Model.IncomeReport, error, error, error) {
	harian := Model.IncomeReport{
		CurrentValue: 0,
		OldValue:     0,
		Percentage:   0,
		IsNegative:   false,
	}

	mingguan := Model.IncomeReport{
		CurrentValue: 0,
		OldValue:     0,
		Percentage:   0,
		IsNegative:   false,
	}

	bulanan := Model.IncomeReport{
		CurrentValue: 0,
		OldValue:     0,
		Percentage:   0,
		IsNegative:   false,
	}

	var errHarian error
	var errMingguan error
	var errBulanan error

	var persentase float64

	now := time.Now()
	enddate := now.Format("2006-01-02")
	hariandate := now.AddDate(0, 0, -1).Format("2006-01-02")
	mingguandate := now.AddDate(0, 0, -7).Format("2006-01-02")
	bulanandate := now.AddDate(0, -1, 0).Format("2006-01-02")

	oldhariandate := now.AddDate(0, 0, -1).Format("2006-01-02")
	oldmingguandate := now.AddDate(0, 0, -14).Format("2006-01-02")
	oldbulanandate := now.AddDate(0, -2, 0).Format("2006-01-02")

	// WaitGroup untuk goroutine pertama
	var wg1 sync.WaitGroup

	// Goroutine pertama untuk menghitung CurrentValue dan OldValue
	wg1.Add(6) // 6 operasi di goroutine pertama
	go func() {
		defer wg1.Done()
		dataharian, err := Model.GetSalesDetail(id, hariandate, enddate)
		if err != nil {
			errHarian = err
		}
		for _, v := range dataharian {
			harian.CurrentValue += v.Price
		}
	}()
	go func() {
		defer wg1.Done()
		datamingguan, err := Model.GetSalesDetail(id, mingguandate, enddate)
		if err != nil {
			errMingguan = err
		}
		for _, v := range datamingguan {
			mingguan.CurrentValue += v.Price
		}
	}()
	go func() {
		defer wg1.Done()
		databulanan, err := Model.GetSalesDetail(id, bulanandate, enddate)
		if err != nil {
			errBulanan = err
		}
		for _, v := range databulanan {
			bulanan.CurrentValue += v.Price
		}
	}()
	go func() {
		defer wg1.Done()
		dataharian, err := Model.GetSalesDetail(id, oldhariandate, hariandate)
		if err != nil {
			errHarian = err
		}
		for _, v := range dataharian {
			harian.OldValue += v.Price
		}
	}()
	go func() {
		defer wg1.Done()
		datamingguan, err := Model.GetSalesDetail(id, oldmingguandate, mingguandate)
		if err != nil {
			errMingguan = err
		}
		for _, v := range datamingguan {
			mingguan.OldValue += v.Price
		}
	}()
	go func() {
		defer wg1.Done()
		databulanan, err := Model.GetSalesDetail(id, oldbulanandate, bulanandate)
		if err != nil {
			errBulanan = err
		}
		for _, v := range databulanan {
			bulanan.OldValue += v.Price
		}
	}()

	// Tunggu semua goroutine pertama selesai
	wg1.Wait()

	// WaitGroup untuk goroutine kedua
	var wg2 sync.WaitGroup
	wg2.Add(3) // 3 operasi di goroutine kedua
	go func() {
		defer wg2.Done()
		if harian.CurrentValue > harian.OldValue {
			harian.IsNegative = false
		} else {
			harian.IsNegative = true
		}
		
		if harian.OldValue != 0 {
			persentase = float64(harian.CurrentValue-harian.OldValue) / float64(harian.OldValue) * 100
			if persentase < 0 {
				persentase = persentase * -1
			}
		} else {
			persentase = 100 // atau nilai default lain
		}
		
		harian.Percentage = persentase
	}()
	go func() {
		defer wg2.Done()
		if mingguan.CurrentValue > mingguan.OldValue {
			mingguan.IsNegative = false
		} else {
			mingguan.IsNegative = true
		}
		
		if mingguan.OldValue != 0 {
			persentase = float64(mingguan.CurrentValue-mingguan.OldValue) / float64(mingguan.OldValue) * 100
			if persentase < 0 {
				persentase = persentase * -1
			}
		} else {
			persentase = 100 // atau nilai default lain
		}
		
		mingguan.Percentage = persentase
	}()
	go func() {
		defer wg2.Done()
		if bulanan.CurrentValue > bulanan.OldValue {
			bulanan.IsNegative = false
		} else {
			bulanan.IsNegative = true
		}
		
		if bulanan.OldValue != 0 {
			persentase = float64(bulanan.CurrentValue-bulanan.OldValue) / float64(bulanan.OldValue) * 100
			if persentase < 0 {
				persentase = persentase * -1
			}
		} else {
			persentase = 100 // atau nilai default lain
		}
		
		bulanan.Percentage = persentase
	}()

	// Tunggu semua goroutine kedua selesai
	wg2.Wait()

	return harian, mingguan, bulanan, errHarian, errMingguan, errBulanan
}
