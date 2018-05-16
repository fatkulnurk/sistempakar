/******************************************************************
*
				TUGAS SISTEM PAKAR MINGGU 10

			Nama	: Fatkul Nur Koirudin
			Nrp		: 2110161019
			Matkul	: Kecerdasan Komputasional (AI)
*
*******************************************************************/

package main

import (
	"fmt"
)

// Indication
type Angka int

type Indikasi struct {
	key int
	angka []Angka
	value float32
}

// Struct Questions Declaration
type Question struct {
	number int
	label string
	answer	string
}

// Deskripsi penyakit
type Description struct {
	number int
	label string
}

// Var to save temp string
var tempString string

func main() {
	/*
	*	Pengendali Looping
	*/
	i := 0	//	jumlah array mencari penyakit
	j :=0
	z := 0	//	jumlah jawaban Y

	var q int
	var pfloat float32
	var threshold int


	// Indikasi Penyakit 1
	var indikasi1 = [] Indikasi{
		Indikasi{
			key:33,
			angka:[]Angka{
				20,21,22,23,29,
			},
		},
		Indikasi{
			key:34,
			angka:[]Angka{
				20,21,22,24,30,
			},
		},
		Indikasi{
			key:35,
			angka:[]Angka{
				20,21,22,25,26,29,
			},
		},
		Indikasi{
			key:36,
			angka:[]Angka{
				21,27,31,
			},
		},
		Indikasi{
			key:37,
			angka:[]Angka{
				28,22,25,32,
			},
		},
	}

	// Indikasi Penyakit 2
	var indikasi2 = [] Indikasi{
		Indikasi{
			key:20,
			angka:[]Angka{
				1,2,4,5,
			},
		},
		Indikasi{
			key:21,
			angka:[]Angka{
				4,5,6,
			},
		},
		Indikasi{
			key:22,
			angka:[]Angka{
				4,7,
			},
		},
		Indikasi{
			key:23,
			angka:[]Angka{
				4,8,9,
			},
		},
		Indikasi{
			key:24,
			angka:[]Angka{
				8,19,
			},
		},
		Indikasi{
			key:25,
			angka:[]Angka{
				4,5,9,11,
			},
		},
		Indikasi{
			key:26,
			angka:[]Angka{
				4,8,11,12,
			},
		},
		Indikasi{
			key:27,
			angka:[]Angka{
				4,13,
			},
		},
		Indikasi{
			key:28,
			angka:[]Angka{
				1,2,3,4,5,
			},
		},
		Indikasi{
			key:29,
			angka:[]Angka{
				14,15,
			},
		},
		Indikasi{
			key:30,
			angka:[]Angka{
				14,16,
			},
		},
		Indikasi{
			key:31,
			angka:[]Angka{
				14,17,
			},
		},
		Indikasi{
			key:32,
			angka:[]Angka{
				18,19,
			},
		},
	}

	// deklarasi jawaban yes dari pertanyaan
	var answerquestion [19]int

	// Keterangan gambar
	var keteranganGambar = [] Description{
		{number:1,label:"Buang air besar (lebih dari 2 kali)"},
		{number:2,label:"Berak Encer"},
		{number:3,label:"Berak Berdarah"},
		{number:4,label:"Lesu dan tidak bergairah"},
		{number:5,label:"Tidak selera makan"},
		{number:6,label:"Merasa mual dan sring muntah (lebih dari 1 kali)"},
		{number:7,label:"Merasa sakir di bagian perut"},
		{number:8,label:"Tekanan darah rendah"},
		{number:9,label:"Pusing"},
		{number:10,label:"Pingsan"},
		{number:11,label:"Suhu badan tinggi"},
		{number:12,label:"Luka di bagian tertentu"},
		{number:13,label:"Tidak dapat menggerakan anggota badan tertentu"},
		{number:14,label:"Memakan sesuatu"},
		{number:15,label:"Memakan daging"},
		{number:16,label:"Memakan jamur"},
		{number:17,label:"Memakan makanan kaleng"},
		{number:18,label:"Membeli susu"},
		{number:19,label:"Meminum susu"},
		{number:20,label:"Mencret"},
		{number:21,label:"Muntah"},
		{number:22,label:"Sakit perut"},
		{number:23,label:"Darah Rendah"},
		{number:24,label:"Koma"},
		{number:25,label:"Demam"},
		{number:26,label:"Septicaemia"},
		{number:27,label:"Lumpuh"},
		{number:28,label:"Mencret Berdarah"},
		{number:29,label:"Makan daging"},
		{number:30,label:"Makan jamur"},
		{number:31,label:"Makan makanan kaleng"},
		{number:32,label:"Minum susu"},
		{number:33,label:"Keracunan Staphylococcus aureus"},
		{number:34,label:"Keracunan jamur beracun"},
		{number:35,label:"Keracunan Salmonellae"},
		{number:36,label:"Keracunan Clostridium botulinum"},
		{number:37,label:"Keracunan Clampylobacter"},
	}

	fmt.Println(keteranganGambar)

	// Kombinasi Slice & Struct, membuat Slice berdasarkan struct
	var questions = []Question{
		{number:1,label:"Apakah anda sering mengalami buang air besar (lebih dari 2 kali)",answer:"N"},
		{number:2,label:"Apakah anda mengalami berak encer",answer:"N"},
		{number:3,label:"Apakah anda mengalami berak berdarah",answer:"N"},
		{number:4,label:"Apakah anda merasa lesu dan tidak bergairah",answer:"N"},
		{number:5,label:"Apakah anda tidak selera makan",answer:"N"},
		{number:6,label:"Apakah anda merasa mual dan sering muntah (lebih dari 1)",answer:"N"},
		{number:7,label:"Apakah anda merasa sakit di bagian perut",answer:"N"},
		{number:8,label:"Apakah tekanan darah anda rendah",answer:"N"},
		{number:9,label:"Apakah anda merasa pusing",answer:"N"},
		{number:10,label:"Apakah anda mengalami pusing",answer:"N"},
		{number:11,label:"Apakah suhu badan anda tinggi",answer:"N"},
		{number:12,label:"Apakah anda mengamali luka dibagian tertantu",answer:"N"},
		{number:13,label:"Apakah anda tidak dapat menggerakkan anggota badan tertentu",answer:"N"},
		{number:14,label:"Apakah anda pernah memakan sesuatu",answer:"N"},
		{number:15,label:"Apakah anda memakan daging",answer:"N"},
		{number:16,label:"Apakah anda memakan jamur",answer:"N"},
		{number:17,label:"Apakah anda memakan makanan kaleng",answer:"N"},
		{number:18,label:"Apakah anda membeli susu",answer:"N"},
		{number:19,label:"Apakah anda meminum susu",answer:"N"},
	}


	// Change value allQuestions.answer
	for i, question := range questions {
		fmt.Print(question.number, ". ", question.label," <Y/N>: ")
		fmt.Scanf("%s",&tempString)
		if tempString == "Y" || tempString == "y"{
			questions[i].answer = "Y"
		}
	}

	// jawaban dari pertanyaan
	fmt.Println("Result questions is :")


	for _, question := range questions {
		fmt.Println(question.number, ". ", question.label,". ")
		fmt.Println("Answer is : ",question.answer)
	}


	// mengambil index yang bernilai yes
	for _, question := range questions {
		if question.answer == "Y" || question.answer == "y"{
			answerquestion[z] = question.number 		// di isi nilai number
			z++;
		}
	}


	// mencari penyakit

	for ooo,a := range indikasi2{
		//fmt.Println("key is ", a.key)
		i=0
		for j=0;j< len(a.angka);j++{
			for _,answqst := range answerquestion{
				if Angka(answqst) == a.angka[j] {
					i++
				}
			}
		}
		// hasilnya adalah i
		q = len(a.angka)
		//fmt.Println("Panjang ",q)
		//fmt.Println("Angka adalah ",i)
		//a.value = float32((i * 100 ) / q)
		indikasi2[ooo].value = float32((i * 100 ) / q)
		//fmt.Println("value % ",indikasi2[ooo].value)
		//fmt.Println("--------------------")

		// assign to struct
	}


	// mencari indikasi1
	fmt.Println("---------------")
	fmt.Println("Indikasi atas")
	fmt.Println("---------------")

	for ooo,bdump := range indikasi1{
		//fmt.Println("Key -> ",bdump.key)
		//fmt.Println("Angka -> ",bdump.angka)
		//fmt.Println("Value -> ",bdump.value)
		q = len(bdump.angka)

		//r = 100/q // berapa persen

		j =0
		pfloat = 0

		for j < q{
			//fmt.Println(bdump.angka[j])

			//fmt.Println("-----------------")
			for _,adump := range indikasi2{
				if int(bdump.angka[j]) == adump.key{
					pfloat = pfloat + adump.value
					//fmt.Println("key ",adump.key," nilainya adalah ", adump.value)
				}
				// fmt.Println(adump.key," Value ",adump.value)
			}

			j++
		}
		//fmt.Println("nilai Q adalah ",q)
		//fmt.Println("Nilai pfloat = ",pfloat)
		//pfloat = (pfloat * 100)/ float32(q)
		pfloat = pfloat / float32(q)
		indikasi1[ooo].value = pfloat
		//fmt.Println("Nilai pfloat = ",pfloat)
		//fmt.Println("-----------------")
	}


	/*
	Pertanyaan Jumlah Threshold
	*/
	fmt.Print("Masukan Jumlah Threshold [INT] : ")
	fmt.Scanf("%d",&threshold)

	fmt.Println("--------------------------------------")

	// MENAMPILKAN HASIL
	for _,adump := range indikasi1 {
		fmt.Println(keteranganGambar[adump.key-1])
		fmt.Println(" Value ",adump.value)

		if adump.value > float32(threshold){
			fmt.Println("Nilai Melebihi Threshold")
		}else {
			fmt.Println("Nilai tidak Melebihi Threshold")
		}
		fmt.Println("--------------------------------------")
	}

	// penghapusan nilai variabel
	_=q
}

