package Controller

import (
	"Sipanjul/Model"
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