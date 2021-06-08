package main

import (
	//"crypto/x509"
	"fmt"
	//"strings"
	//"time"
	"math"
)

/////////////////////////////////////////////////////////////////
func topla(değer1 int, değer2 int) int{
	return(değer1+değer2)
}
func çarp(değer1 int, değer2 int) int{
    return(değer1*değer2)
}
func böl(değer1 int, değer2 int) int{
	return(değer1/değer2)
}
func çıkart(değer1 int, değer2 int) int{
	return(değer1-değer2)
}
/////////////////////////////////////////////////////////////////
/*func hata() string{
	return("Bir hata oluştu, kod çalıştırılamıyor") //her hata kodunu tek tek yazmak yerine tek yere tanımlıyıp devam edicem
}*/
/////////////////////////////////////////////////////////////////
func success() {
  fmt.Println("main.go dosyasında herhangi bir kod hatası bulunmuyor ve dosya başarıyla çalıştırıldı.")
}
func main() {
	success()
	fmt.Println("İsmin?: ")
	var isim string
	fmt.Scanln(&isim)
	fmt.Println("Güzel, merhaba", isim)
	fmt.Println("Hmm, biraz eğlenmek istiyorum seninle işlem yapalım mı?")
	var cevap string
	fmt.Scanln(&cevap)
	if cevap == "evet"{
fmt.Println("Güzel Şimdi girdiğin sayının karekökünü bulacağım")
var sqrt int
fmt.Scanln(&sqrt)
cevab := math.Sqrt(float64(sqrt))
fmt.Println(sqrt, "sayısının karekökü: ", cevab)
	}
if cevap == "hayır"{
	fmt.Println("Oh, peki. Sanırım matematikle aran iyi değil. O zamaaan, görüşürüz :)")
}
}
/*if strings.Contains(islem, "*") == "true"{
	sonuc := çarp(strings.Replace(islem, "*", ",", -1))
}
if strings.Contains(islem, "+") == "true"{
	topla(strings.Replace(islem, "+", ",", -1))
}
if strings.Contains(islem, "-") == "true"{
	çıkart(strings.Replace(islem, "-", ",", -1))
}
if strings.Contains(islem, "/") == "true"{
	böl(strings.Replace(islem, "/", ",", -1))
}
if strings.Contains(islem, ".") == "true"{
	çarp(strings.Replace(islem, ".", ",", -1))
}
if strings.Contains(islem, "x") == "true"{
	çarp(strings.Replace(islem, "x", ",", -1))
}
if strings.Contains(islem, ":") == "true"{
	çarp(strings.Replace(islem, ":", ",", -1))
}*/