package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Sale struct {
	ProductID int
	Amount    float64
}

func main() {
	fmt.Println("Analise de Dados de Vendas para Empresas!!")

	sales, err := readSalesFromCSV("datas.csv") //o arquivo q contem as datas de vendas estão aq
	if err != nil {
		fmt.Println("Erro ao ler esse tipo de dados, analise se tem algo de errado:", err)
		return
	}

	totalSales := calculateTotalSales(sales)
	averageSale := calculateAverageSale(sales)

	fmt.Printf("Total de Vendas: R$%.2f\n", totalSales)
	fmt.Printf("Venda Média por Produto: R$%.2f\n", averageSale)
}

func readSalesFromCSV(filename string) ([]Sale, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var sales []Sale
	for _, line := range lines {
		productID, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, err
		}
		amount, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			return nil, err
		}

		sale := Sale{
			ProductID: productID,
			Amount:    amount,
		}
		sales = append(sales, sale)
	}

	return sales, nil
}

func calculateTotalSales(sales []Sale) float64 {
	total := 0.0
	for _, sale := range sales {
		total += sale.Amount
	}
	return total
}

func calculateAverageSale(sales []Sale) float64 {
	if len(sales) == 0 {
		return 0.0
	}

	totalSales := calculateTotalSales(sales)
	return totalSales / float64(len(sales))
}