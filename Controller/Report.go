package Controller

import (
	"Sipanjul/Model"
	"bytes"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func GetProductReport(startdate,enddate,divisi,detail string) ([]Model.ProductReportOpr, error) {
	data,err := Model.GetProductReport(startdate,enddate,divisi,detail)
	if err != nil {
		return nil, err
	}

	return data, err
}

func GetSelesReport(startdate,enddate,divisi string, id uint) ([]Model.Sales_Report, error) {
	data, err := Model.GetSalesDetail(id,startdate, enddate)
	if err != nil {
		return nil, err
	}

	var filter []Model.Sales_Report

	if divisi == "SCPP" {
		for _, v := range data {
			if v.Divisi == "SCPP"{
				filter = append(filter, v)
			}
		}

		return filter, nil
	}

	if divisi == "Komersil"{
		for _, v := range data {
			if v.Divisi == "Komersil"{
				filter = append(filter, v)
			}
		}
		return filter,nil
	}

	return data,nil
}

func GenerateDataExcel(salesReport []Model.Sales_Report, productReport []Model.ProductReportOpr) ([]Model.Print, []Model.PrintOut, error){
	data, dataOut, err := Model.DataPrint(salesReport, productReport)
	if err != nil {
		return nil, nil, err
	}

	return data, dataOut, nil
}

func GenerateExcelPenjualan(data []Model.Print, dataOut []Model.PrintOut ) ([]byte, error) {
	f := excelize.NewFile()
	sheet := "Data Penjualan Toko"

	// Ganti nama sheet
	f.SetSheetName("Sheet1", sheet)

	// Tambahkan tanggal di B1
	f.MergeCell(sheet, "A1", "B1")
	f.SetCellValue(sheet, "A1", "20 Desember 24")

	f.MergeCell(sheet, "D1", "G1")
	f.SetCellValue(sheet, "D1", "Data Penjualan Toko")

	// Tambahkan header utama
	f.MergeCell(sheet, "A2", "A4")
	f.SetCellValue(sheet, "A2", "No")
	f.MergeCell(sheet, "B2", "B4")
	f.SetCellValue(sheet, "B2", "Komoditi")
	f.MergeCell(sheet, "C2", "C4")
	f.SetCellValue(sheet, "C2", "Kemasan")
	f.MergeCell(sheet, "D2", "D4")
	f.SetCellValue(sheet, "D2", "Harga Jual (RP)")
	f.MergeCell(sheet, "E2", "H2")
	f.SetCellValue(sheet, "E2", "Dikeluarkan dari BM")
	f.MergeCell(sheet, "E3", "F3")
	f.SetCellValue(sheet, "E3", "Stok")
	f.SetCellValue(sheet, "E4", "Awal")
	f.SetCellValue(sheet, "F4", "Tambahan")
	f.MergeCell(sheet, "G3", "G4")
	f.SetCellValue(sheet, "G3", "Terjual")
	f.MergeCell(sheet, "H3", "H4")
	f.SetCellValue(sheet, "H3", "Sisa")
	f.MergeCell(sheet, "I2", "I4")
	f.SetCellValue(sheet, "I2", "Hasil Penjualan (RP)")
	f.MergeCell(sheet, "J2", "J4")
	f.SetCellValue(sheet, "J2", "Stok Akhir")

	// Tambahkan data produk
	for i, product := range data {
		no := 1
		row := i + 5
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), no)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), product.Komoditi)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), product.Kemasan)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), product.Harga)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), product.StockAwal)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), product.StockTambahan)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), product.Terjual)
		f.SetCellValue(sheet, fmt.Sprintf("H%d", row), product.Sisa)
		f.SetCellValue(sheet, fmt.Sprintf("I%d", row), product.Hasil)
		f.SetCellValue(sheet, fmt.Sprintf("J%d", row), product.Sisa)
		no++
	}

	for i, product := range dataOut {
		row := i + 32
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), product.Komoditi)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), product.Deskripsi)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), product.Jumlah)
	}

	// Tambahkan header untuk "Stok Keluar"
	f.SetCellValue(sheet, "C30", "Stok Keluar")
	f.SetCellValue(sheet, "B31", "Nama")
	f.SetCellValue(sheet, "C31", "Komoditi")
	f.SetCellValue(sheet, "D31", "Jumlah")

	// Tambahkan footer
	f.SetCellValue(sheet, "I30", "Total			:")
	f.SetCellValue(sheet, "I31", "Pengeluaran	:")
	f.SetCellValue(sheet, "I32", "Uang Fisik	:")
	f.SetCellValue(sheet, "I33", "Selisih		:")

    // Apply the style to cells
    styleID, _ := f.NewStyle(&excelize.Style{
        Border: []excelize.Border{
            {Type: "left", Color: "000000", Style: 1},
            {Type: "top", Color: "000000", Style: 1},
            {Type: "bottom", Color: "000000", Style: 1},
            {Type: "right", Color: "000000", Style: 1},
        },
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Font: &excelize.Font{
			Bold: true,
		},
    })

	styleFont , _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "left",
			Vertical:   "center",
		},
		Border: []excelize.Border{
            {Type: "left", Color: "000000", Style: 1},
            {Type: "top", Color: "000000", Style: 1},
            {Type: "bottom", Color: "000000", Style: 1},
            {Type: "right", Color: "000000", Style: 1},
        },
	})

	styleCenter , _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
            {Type: "left", Color: "000000", Style: 1},
            {Type: "top", Color: "000000", Style: 1},
            {Type: "bottom", Color: "000000", Style: 1},
            {Type: "right", Color: "000000", Style: 1},
        },
	})

	styleLeft , _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "left",
			Vertical:   "center",
		},
		Border: []excelize.Border{
            {Type: "left", Color: "000000", Style: 1},
            {Type: "top", Color: "000000", Style: 1},
            {Type: "bottom", Color: "000000", Style: 1},
            {Type: "right", Color: "000000", Style: 1},
        },
	})

	styleFontBold, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})

	styleBoldCenter, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})

	f.SetColWidth(sheet, "A", "A", 4.44)
	f.SetColWidth(sheet, "B", "B", 19.11)
	f.SetColWidth(sheet, "C", "C", 16.22)
	f.SetColWidth(sheet, "D", "D", 15.56)
	f.SetColWidth(sheet, "E", "E", 4.89)
	f.SetColWidth(sheet, "F", "F", 10.22)
	f.SetColWidth(sheet, "G", "G", 6.89)
	f.SetColWidth(sheet, "H", "H", 5.78)
	f.SetColWidth(sheet, "I", "I", 17.44)
	f.SetColWidth(sheet, "J", "J", 13.37)

	f.SetCellStyle(sheet, "A2", "J4", styleID)
	f.SetCellStyle(sheet, "D1", "G1", styleBoldCenter)
	f.SetCellStyle(sheet, "B31", "D31", styleID)
	
	f.SetCellStyle(sheet, "I30", "J33", styleFont)

	f.SetCellStyle(sheet, "A5", "J29", styleCenter)

	f.SetCellStyle(sheet, "B5", "J29", styleLeft)
	f.SetCellStyle(sheet, "B32", "D34", styleLeft)

	f.SetCellStyle(sheet, "A1", "B1", styleFontBold)
	f.SetCellStyle(sheet, "A30", "D30", styleBoldCenter)
	
    // Simpan file ke buffer
    var buf bytes.Buffer
    if err := f.Write(&buf); err != nil {
        return nil, err
    }

    return buf.Bytes(), nil
}