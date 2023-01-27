package main

// kegunaan interface
// - parsing / menggunakan beberapa struct yg berbeda-beda untuk dijadikan parameter

type BangunDatar interface {
// biasanya berupa method
HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

type Origin struct {
	Number int
}

func (origin Origin) HitungLuas() int {
	return 100
}

func main() {
	square := Persegi{5}
	luas := HowBroad(square)
	println(luas)

	block := PersegiPanjang{7, 5}
	luas = HowBroad(block)
	println(luas)

	asal := Origin{6}
	luas = HowBroad(asal)
	println(luas)
	// luas2 := block.HitungLuas()
}

func HowBroad(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}