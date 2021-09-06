package main

import (
	"fmt"
	"strings"
)

func main(){
	// Risky Kurniawan - Universitas Adhirajasa Reswara Sanjaya
	var menu int
	fmt.Println("===========================================")
	fmt.Println("          APLIKASI PENUKARAN UANG")
	fmt.Println("===========================================")
	fmt.Println("Silahkan pilih menu penukaran uang")
	fmt.Println("1. Rupiah ke Dollar")
	fmt.Println("2. Rupiah ke Euro")
	fmt.Println("3. GBP ke Knut")
	fmt.Println("===========================================")
	fmt.Print("Ketikan nomor menu pilihan anda : ")
	fmt.Scan(&menu)
	fmt.Println("===========================================")

	if menu == 1 {
		IDR_USD()
	}else if menu == 2 {
		IDR_EUR()
	}else if menu == 3 {
		GBP_KNUT()
	}else{
		fmt.Println("Menu pilihan tidak tersedia")
	}

	fmt.Println("===========================================")
	fmt.Println("           -- TERIMAKASIH --")
	fmt.Println("    By Risky Kurniawan - ARS University")
	fmt.Println("===========================================")
}

func GBP_KNUT() {
	var GBP int
	fmt.Println("===========================================")
	fmt.Println("          GBP -> KNUT")
	fmt.Println("===========================================")
	fmt.Print("Masukan Jumlah GB Pounds : ")
	fmt.Scan(&GBP)
	var knut int = GBP*100
	fmt.Println("Jumlah Knut yang didapat = ", knut)
	var galleon int = (knut / 29) / 17
	var sisa_knut int = knut - ((galleon * 17) * 29)
	var sickle int = sisa_knut / 29
	sisa_knut = sisa_knut - (29 * sickle)
	fmt.Println("Hasil penukaran mendapat = ", galleon, " Galleon (s)")
	fmt.Println("Sisa ditukar menjadi     = ", sickle, " Sickle (s)")
	fmt.Println("Keping knut yang tersisa = ", sisa_knut, " Knut (s)")
	fmt.Println("===========================================")
}

func IDR_EUR() {
	var IDR int
	var kurs float64 = 16892.74
	fmt.Println("===========================================")
	fmt.Println("          RUPIAH -> EURO")
	fmt.Println("===========================================")
	fmt.Print("Masukan Nilai Rupiah : ")
	fmt.Scan(&IDR)
	fmt.Println("===========================================")
	var convert float64 = float64(IDR) / kurs
	fmt.Println("Nilai Rupiah    = ", currency(float64(IDR), "IDR"))
	fmt.Println("Nilai Euro      = ", currency(convert, "EUR"))
	fmt.Println("===========================================")
	fmt.Println("Kurs            = ", currency(1, "EUR")," -> ", currency(kurs, "IDR"))
	fmt.Println("===========================================")
}

func IDR_USD() {
	var IDR int
	var kurs float64 = 14231.75
	fmt.Println("===========================================")
	fmt.Println("          RUPIAH -> US DOLLAR")
	fmt.Println("===========================================")
	fmt.Print("Masukan Nilai Rupiah : ")
	fmt.Scan(&IDR)
	fmt.Println("===========================================")
	var convert float64 = float64(IDR) / kurs
	fmt.Println("Nilai Rupiah    = ", currency(float64(IDR), "IDR"))
	fmt.Println("Nilai US Dollar = ", currency(convert, "USD"))
	fmt.Println("===========================================")
	fmt.Println("Kurs            = ", currency(1, "USD")," -> ", currency(kurs, "IDR"))
	fmt.Println("===========================================")
}

func currency(number float64, format string) string{
	// Modifikasi Dari tugas sebelumnya
	var money, sim string
	if format == "IDR"{
		money = "Rp. "
		sim = "."
	}else if format == "USD"{
		money = "$ "
		sim = ","
	}else if format == "EUR"{
		money = "€ "
		sim = ","
	}else{
		return "Error Number Format"
	}

	valnum := strings.Split(fmt.Sprintf("%.2f", number), ".")[0]
	var reverse string = ""

	var index int = 1
	
	for i:=(strings.Count(valnum, "")-2); i>=0; i--{
		if (index%3) == 1 && index != 1{
			reverse += sim
		}
		reverse += string(valnum[i])
		
		index++
	}

	valnum = reverse

	for i:=(strings.Count(valnum, "")-2); i>=0; i--{
		money += string(valnum[i])
	}

	if format == "USD"{
		money += "." + strings.Split(fmt.Sprintf("%.2f", number), ".")[1] + " USD"
	}else if format == "IDR"{
		money += "," + strings.Split(fmt.Sprintf("%.2f", number), ".")[1]
	}else if format == "EUR"{
		money += "." + strings.Split(fmt.Sprintf("%.2f", number), ".")[1] + " EUR"
	}
	
	return money
}