package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// CONSTANTS
const NMAX int = 1000

// DATA STRUCTURES
type Users struct {
	Nama, Email, Password string
}

type AsetInvetasi struct {
	NamaAset, JenisAset     string
	JumlahDanaInvestasi, NilaiAset float64
}

// GLOBAL VARIABLES
type DataUsers = [NMAX]Users
type DataInvestasi = [NMAX]AsetInvetasi

var AllUsers DataUsers
var AllInvestasi DataInvestasi

var idxUser = 2
var idxInvestasi = 6



func main() {
	// memasukan dataDummy untuk user default ke array
	InitDummyUsers(&AllUsers)
	// memasukan dataDumy untuk Investasi default ke array
	InitDummyDataInvestasi(&AllInvestasi)

	var aplikasiAktif bool = true

	for aplikasiAktif {
		var accountOption int
		var isLoggIn bool
		WelcomeLogin()
		fmt.Scan(&accountOption)
		ClearScreen()
		if accountOption == 1 {
			isLoggIn = InputLogin()

		} else if accountOption == 2 {
			RegisterUser(&AllUsers,idxUser)
			var input string
			fmt.Println("Ingin Login? (Y/N):")
			fmt.Scan(&input)
			if input == "Y" || input == "y" {
				isLoggIn = InputLogin()
			}
		} else if accountOption == 3 {
			fmt.Println("Terimakasih Sudah Menggunakan Aplikasi. Sampai Jumpa")
			aplikasiAktif = false
		} else {
			fmt.Println("Pilihan Tidak valid. Silahkan Pilih 1, 2, atau 3.")
		}

		var inDashboard bool = true
		for isLoggIn && inDashboard {
			var dashboradOption int
			var namaAset, jenisAset string
			var jumlahDana, nilaiAwal float64

			DashboardOption()
			fmt.Scan(&dashboradOption)
			ClearScreen()
			
			switch dashboradOption {
			case 1:
    			namaAset, jenisAset, jumlahDana, nilaiAwal = InputInvestasi()
    			CreateDataInvestasi(namaAset, jenisAset, nilaiAwal, jumlahDana)
    			TampilkanDataTerakhir()
			case 2:
    			fmt.Print("Masukkan Nama Aset yang akan di ubah: ")
    			fmt.Scan(&namaAset)
    			ModifyDataInvestasi(namaAset)
			case 3:
    			fmt.Print("Masukkan Nama Aset yang akan di hapus: ")
    			fmt.Scan(&namaAset)
    			HapusDataInvestasi(namaAset)
			case 4:
				fmt.Println("---------------------------------------------------------")
				fmt.Println("            CARI DATA INVESTASI BERDASARKAN             ")
				fmt.Println("---------------------------------------------------------")
				fmt.Println("1. Nama Aset")
				fmt.Println("2. Jenis Aset")
				fmt.Println("3. Jumlah Dana")
				fmt.Println("---------------------------------------------------------")
				var input int
				fmt.Scan(&input)
				switch input{
				case 1:
					var namaAset string
					fmt.Print("Masukkan Nama Aset: ")
					fmt.Scan(&namaAset)
					FindDataByName(namaAset)
				case 2:
					var jenisAset string
					fmt.Print("Masukkan Jenis Aset: ")
					fmt.Scan(&jenisAset)
					FindDataByJenis(jenisAset)
				case 3:
					
					var jumlahDana float64
					fmt.Print("Masukkan Jumlah Dana: ")
					fmt.Scan(&jumlahDana)
					FindDataByJumlahDana(jumlahDana)
				}
			case 5:
				var inputBerdasarkan int
				var inputJenisSort int
				fmt.Println("Pilih Pengurutan Data:")
				fmt.Println("1. Ascending")
				fmt.Println("2. Descending")

				fmt.Scan(&inputJenisSort)			

				fmt.Println("Urutkan data investasi berdasarkan:")
				fmt.Println("1. Nama Aset")
				fmt.Println("2. Jenis Aset")
				fmt.Println("3. Jumlah Dana")
				fmt.Println("4. Total Keuntungan")
				
				fmt.Scan(&inputBerdasarkan)
				switch inputBerdasarkan{
				case 1:
					if inputJenisSort == 1{
						SortAscendingNamaAset()
					}else if inputJenisSort == 2 {
						SortAscendingJenisAset()
					}else if inputJenisSort == 3 {
						SortAscendingJumlahDana()
					}else if inputJenisSort == 4 {
						SortAscendingTotalKeuntungan()
					}
				case 2:
					if inputJenisSort == 1{
						SortDescendingNamaAset()
					}else if inputJenisSort == 2 {
						SortDescendingJenisAset()
					}else if inputJenisSort == 3 {
						SortDescendingJumlahDana()
					}else if inputJenisSort == 4 {
						SortDescendingTotalKeuntungan()
					}
				}
			case 6:
    			TampilkanLaporanInvestasi()
			case 7:
    			fmt.Println("Anda telah Logout.")
    			inDashboard = false
			case 8:
    			fmt.Println("Terima kasih telah menggunakan aplikasi.")
    			aplikasiAktif = false
    			inDashboard = false
			default:
    			fmt.Println("Pilihan tidak valid! Silahkan pilih 1-8.")
			}
		}
	}
}



func InitDummyUsers(TUsers *DataUsers) {
	// IS. Terdefinisi array TUsers dalam keadaan kosong tanpa data user apapun.
	// FS. Mengisi array TUsers dengan data dummy berupa beberapa user yang sudah ditentukan
	//     pada indeks awal array untuk keperluan pengujian atau inisialisasi awal program.
	TUsers[0] = Users{
		Nama:     "son",
		Email:    "son.com",
		Password: "son",
	}
	TUsers[1] = Users{
		Nama:     "Keanu",
		Email:    "keanu.com",
		Password: "nu",
	}
}

func InitDummyDataInvestasi(TAset  *DataInvestasi) {
	// IS. Terdefinisi array TAset dalam keadaan kosong tanpa data aset investasi apapun.
	// FS. Mengisi array TAset dengan beberapa data dummy aset investasi yang sudah ditentukan
	//     pada indeks awal array untuk tujuan pengujian atau inisialisasi awal program.
	TAset[0]= AsetInvetasi{
		NamaAset: "BBRI",
		JenisAset: "Saham",
		NilaiAset: 1000000,
		JumlahDanaInvestasi: 100000000,
	}
	TAset[1]= AsetInvetasi{
		NamaAset: "BBRI",
		JenisAset: "Saham",
		NilaiAset: 1000000,
		JumlahDanaInvestasi: 100000000,
	}	
	TAset[2]= AsetInvetasi{
		NamaAset: "BBRI",
		JenisAset: "Saham",
		NilaiAset: 1000000,
		JumlahDanaInvestasi: 100000000,
	}	
	TAset[3]= AsetInvetasi{
		NamaAset: "BBRI",
		JenisAset: "Saham",
		NilaiAset: 1000000,
		JumlahDanaInvestasi: 100000000,
	}
	TAset[4]= AsetInvetasi{
		NamaAset: "BBRI",
		JenisAset: "Saham",
		NilaiAset: 1000000,
		JumlahDanaInvestasi: 100000000,
	}	
	TAset[5]= AsetInvetasi{
		NamaAset: "BBRI",
		JenisAset: "Saham",
		NilaiAset: 1000000,
		JumlahDanaInvestasi: 100000000,
	}		
}

// UTILITY FUNCTIONS
func ClearScreen() {
	// IS: Layar terminal dalam kondisi apapun (bisa berisi teks hasil eksekusi sebelumnya)
	// FS: Layar terminal dibersihkan (kosong) sesuai perintah sistem operasi yang digunakan
	osClear := runtime.GOOS

	if osClear == "linux" || osClear == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if osClear == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Println("sistem operasi tidak mendukung.")
	}
}

// USER RELATED FUNCTIONS
func ValidasiUsers(inputUser Users) bool {
	// function yang mengembalikan true jika user dengan email dan password yang diberikan ada di dalam array AllUsers
	for i := 0; i < idxUser; i++ {
		if AllUsers[i].Email == inputUser.Email && AllUsers[i].Password == inputUser.Password {
			return true
		}
	}
	return false
}

func RegisterUser(TUsers *DataUsers, n int) {
	// IS. Terdefinisi array TUsers yang berukuran maksimum NMAX dan sudah terisi sebanyak idxUser elemen user aktif.
	//     Jika idxUser sudah mencapai atau melebihi NMAX, maka tidak ada user baru yang dapat didaftarkan.
	// FS. Jika idxUser masih kurang dari NMAX, maka data user baru akan dimasukkan pada indeks ke-n dalam TUsers,
	//     dan nilai idxUser bertambah satu. Jika idxUser sudah penuh, maka tidak ada perubahan pada TUsers
	//     dan hanya menampilkan pesan bahwa pendaftaran gagal karena kapasitas penuh.
	if idxUser >= NMAX {
		fmt.Println("\n=========================================================")
		fmt.Println("    Gagal! User Sudah Penuh       ")
		fmt.Println("=========================================================")
	}
	var name, email, password string

	fmt.Println("\n=========================================================")
	fmt.Println("         Registrasi Pengguna       ")
	fmt.Println("=========================================================")
	fmt.Print("Masukkan Nama: ")
	fmt.Scan(&name)
	fmt.Print("Masukkan Email: ")
	fmt.Scan(&email)
	fmt.Print("Masukkan Password: ")
	fmt.Scan(&password)

	TUsers[n] = Users{
		Nama:     name,
		Email:    email,
		Password: password,
	}
	idxUser++
	fmt.Println("\n=========================================================")
	fmt.Println("    Akun Berhasil Dibuat           ")
	fmt.Println("    Silahkan Login!                ")
	fmt.Println("=========================================================")
}

func InputLogin() bool {
	// function yang mengembalikan true jika user dengan email dan password yang diberikan
	// berhasil login dalam maksimal 5 kali percobaan. Jika gagal, mengembalikan false.
	var email, password string
	kesempatan := 5
	fmt.Println("=========================================================")
	fmt.Println("          WELCOME TO LOGIN        ")
	fmt.Println("=========================================================")
	fmt.Println("Percobaan Login: ", kesempatan)
	fmt.Println("---------------------------------------------------------")
	fmt.Print("Masukkan email: ")
	fmt.Scan(&email)
	fmt.Print("Masukkan Password: ")
	fmt.Scan(&password)
	fmt.Println("---------------------------------------------------------")

	fmt.Println("Email yang dimasukkan: ", email)
	fmt.Println("Password yang dimasukkan: ", password)

	// membuat struct userInput untuk menampung input dari user agar mudah dimasukan ke validasi dan struct AllUsers
	usersInput := Users{
		Email:    email,
		Password: password,
	}
	// menyimpan hasil dari function ValidasiUsers
	valid := ValidasiUsers(usersInput)

	if valid {
		fmt.Println("\nLogin Berhasil! Selamat datang.")
		return true
	} else {
		for !valid && kesempatan != 1 {
			fmt.Println("\n=========================================================")
			fmt.Println("    Login Gagal! Coba Lagi...")
			fmt.Println("=========================================================")
			fmt.Println("Pastikan Email dan Password Benar.")
			fmt.Println("---------------------------------------------------------")
			kesempatan--
			fmt.Println("Percobaan Login Tersisa: ", kesempatan)
			fmt.Print("Masukkan email: ")
			fmt.Scan(&email)
			fmt.Print("Masukkan Password: ")
			fmt.Scan(&password)
			fmt.Println("---------------------------------------------------------")
			// update data input nya agar ketika input baru, datannya berubah
			usersInput = Users{
				Email:    email,
				Password: password,
			}

			valid = ValidasiUsers(usersInput)
		}

		if valid {
			fmt.Println("\nLogin Berhasil! Selamat datang.")
			return true
		}
		if kesempatan == 1 {
			fmt.Println("\nLogin Gagal! Kesempatan Habis.")
			fmt.Println("\nSilahkan Register Terlebih Dahulu.")
		}
		return false
	}
}

// INVESTMENT RELATED FUNCTIONS
func CreateDataInvestasi(namaAset, jenisAset string, nilaiAset, danaInvestasi float64) {
// IS. Terdefinisi suatu array AllInvestasi yang berukuran maksimum NMAX dan telah terisi sebanyak idxInvestasi elemen.
//      Diberikan input namaAset, jenisAset, nilaiAset, dan danaInvestasi yang akan dimasukkan ke dalam array.
// FS. Menambahkan data aset investasi ke dalam array AllInvestasi apabila belum penuh (idxInvestasi < NMAX),
//      atau menampilkan pesan "Data Investasi Sudah Penuh" apabila array telah mencapai kapasitas maksimum.
	
	dataInvestasi := AsetInvetasi{
		NamaAset:            namaAset,
		JenisAset:           jenisAset,
		NilaiAset:           nilaiAset,
		JumlahDanaInvestasi: danaInvestasi,
	}

	if idxInvestasi < NMAX {
		AllInvestasi[idxInvestasi] = dataInvestasi
		idxInvestasi++
	}else {
		fmt.Println("\n=========================================================")
		fmt.Println("    Gagal! Data Investasi Sudah Penuh       ")
		fmt.Println("=========================================================")
	}
}

func TampilkanDataTerakhir() {
	// IS. Terdefinisi array AllInvestasi yang menyimpan data aset investasi, serta nilai idxInvestasi 
	//     yang menyatakan jumlah data investasi yang sudah dimasukkan ke dalam array.
	// FS. Jika belum ada data investasi (idxInvestasi = 0), maka akan ditampilkan pesan bahwa belum ada data.
	//     Jika terdapat data, maka informasi dari elemen terakhir yang dimasukkan ke dalam array AllInvestasi 
	//     akan ditampilkan ke layar sebagai umpan balik berhasilnya penambahan data.
	if idxInvestasi == 0 {
		fmt.Println("Belum ada data investasi.")
	}
	dataTerakhir := AllInvestasi[idxInvestasi-1]
	fmt.Println("---------------------------------------------------------")
	fmt.Println("        Data Investasi Berhasil Ditambahkan       ")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("Nama Aset          :", dataTerakhir.NamaAset)
	fmt.Println("Jenis Aset         :", dataTerakhir.JenisAset)
	fmt.Println("Jumlah Dana        :", dataTerakhir.JumlahDanaInvestasi)
	fmt.Println("Nilai Aset         :", dataTerakhir.NilaiAset)
	fmt.Println("---------------------------------------------------------")
}

func ModifyDataInvestasi(namaAset string) {
	// IS. Terdefinisi sebuah array AllInvestasi yang berisi data aset investasi sebanyak idxInvestasi elemen.
	//     Parameter namaAset berisi string yang akan digunakan untuk mencari data aset yang ingin dimodifikasi.
	// FS. Jika data dengan NamaAset ditemukan dalam AllInvestasi, maka data tersebut dimodifikasi dengan nilai-nilai baru
	//     yang dimasukkan oleh pengguna (nama aset, jenis aset, jumlah dana, nilai awal, nilai update).
	//     Jika tidak ditemukan, maka ditampilkan pesan bahwa data tidak ditemukan.
	var dataFound bool = false

	for i := 0; i < idxInvestasi; i++ {
		if AllInvestasi[i].NamaAset == namaAset {
			dataFound = true
			fmt.Println("=========================================================")
			fmt.Println("               Modifikasi Data Investasi")
			fmt.Println("=========================================================")
			fmt.Println("Data Sebelumnya:")
			fmt.Printf("Nama Aset: %s\n", AllInvestasi[i].NamaAset)
			fmt.Printf("Jenis Aset: %s\n", AllInvestasi[i].JenisAset)
			fmt.Printf("Jumlah Dana: %.5f\n", AllInvestasi[i].JumlahDanaInvestasi)
			fmt.Printf("Nilai Awal: %.5f\n", AllInvestasi[i].NilaiAset)
			fmt.Println("---------------------------------------------------------")

			// Ambil input data baru
			var namaAsetBaru, jenisAsetBaru string
			var jumlahDanaBaru, nilaiAwalBaru, nilaiUpdateBaru float64

			fmt.Print("\nMasukkan Data Baru: \nJika Nama atau Aset lebih dari 1 kata gunakan '_' sebagai pemisah !!\n")
			fmt.Print("Nama Aset: ")
			fmt.Scan(&namaAsetBaru)
			fmt.Print("Jenis Aset: ")
			fmt.Scan(&jenisAsetBaru)
			fmt.Print("Jumlah Dana: ")
			fmt.Scan(&jumlahDanaBaru)
			fmt.Print("Nilai Awal: ")
			fmt.Scan(&nilaiAwalBaru)
			fmt.Print("Nilai Update Baru: ")
			fmt.Scan(&nilaiUpdateBaru)

			// Update data
			AllInvestasi[i].NamaAset = namaAsetBaru
			AllInvestasi[i].JenisAset = jenisAsetBaru
			AllInvestasi[i].JumlahDanaInvestasi = jumlahDanaBaru
			AllInvestasi[i].NilaiAset = nilaiAwalBaru

			fmt.Println("---------------------------------------------------------")
			fmt.Println("Data investasi berhasil dimodifikasi.")
			fmt.Println("=========================================================")

			return
		}
	}

	if !dataFound {
		fmt.Println("Data investasi tidak ditemukan, Pastikan Nama Aset benar.")
	}
}

func HapusDataInvestasi(namaAset string) {
	// IS: Terdefinisi array AllInvestasi berisi data aset investasi sebanyak idxInvestasi.
	//     Parameter namaAset berisi nama aset yang ingin dihapus.
	// FS: Jika aset dengan nama yang sesuai ditemukan, maka data aset tersebut dihapus dari array AllInvestasi,
	//     elemen setelahnya digeser ke kiri, dan nilai idxInvestasi dikurangi 1.
	//     Jika tidak ditemukan, maka akan ditampilkan pesan bahwa aset tidak ditemukan.
	var foundIndex int = -1

	for i := 0; i < idxInvestasi; i++ {
		if AllInvestasi[i].NamaAset == namaAset && foundIndex == -1 {
			foundIndex = i
		}
	}

	dataAset := AllInvestasi[foundIndex]
	if foundIndex != -1 {
		for j := foundIndex; j < idxInvestasi-1; j++ {
			AllInvestasi[j] = AllInvestasi[j+1]
		}
		idxInvestasi--
		fmt.Println("Data Aset:")
		fmt.Println("---------------------------------------------------------")
		fmt.Printf("Nama Aset    : %s\n", dataAset.NamaAset)
		fmt.Printf("Jenis Aset   : %s\n", dataAset.JenisAset)
		fmt.Printf("Nilai Aset   : %.2f\n", dataAset.NilaiAset)
		fmt.Printf("Jumlah Dana  : %.2f\n", dataAset.JumlahDanaInvestasi)
		fmt.Println("---------------------------------------------------------")
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Aset tidak ditemukan.")
	}
}

func FindDataByName(nameAset string) bool {
	// {diberikan array AllInvestasi yang berisi idxInvestasi data aset dan sebuah nama aset,
	//  untuk mengembalikan true apabila nama aset ditemukan di dalam AllInvestasi,
	//  atau false apabila sebaliknya}
	for i := 0; i < idxInvestasi; i++ {
		if AllInvestasi[i].NamaAset == nameAset {
			return true
		}
	}
	return false
}

func FindDataByJenis(jenisAset string) bool {
	// {diberikan array AllInvestasi yang berisi idxInvestasi data aset dan sebuah jenis aset,
	//  untuk mengembalikan true apabila jenis aset tersebut ditemukan di dalam AllInvestasi,
	//  atau false apabila jenis aset tersebut tidak ditemukan}
	for i := 0; i < idxInvestasi; i++ {
		if AllInvestasi[i].JenisAset == jenisAset {
			return true
		}
	}
	return false
}

func FindByJumlahDana(jumlahDana float64) bool {
	n := idxInvestasi
	left := 0
	right := n - 1

	for left <= right {
		mid := (left + right) / 2
		midValue := AllInvestasi[mid].JumlahDanaInvestasi

		if midValue == jumlahDana {
			// Data ditemukan
			return true
		} else if midValue < jumlahDana {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// Data tidak ditemukan
	return false
}

func SortAscendingNamaAset(){

}

func SortDescendingNamaAset(){

}


func SortAscendingJumlahDana(TAset *DataInvestasi) {
	var i, idx, pass int
	pass = 1
	for pass < idxInvestasi {
		idx = pass - 1
		i = pass
		for i < idxInvestasi {
			if TAset[i].JumlahDanaInvestasi < TAset[idx].JumlahDanaInvestasi {
				idx = i
			}
			i++
		}	
		temp := TAset[pass - 1]
		TAset[pass - 1] = TAset[idx]
		TAset[idx] = temp
		pass++
	}
}

func SortDescendingJumlahDana(TAset *DataInvestasi) {	
	var i, idx, pass int
	pass = 1
	for pass < idxInvestasi {
		idx = pass - 1
		i = pass
		for i < idxInvestasi {
			if TAset[i].JumlahDanaInvestasi > TAset[idx].JumlahDanaInvestasi {
				idx = i
			}
			i++
		}	
		temp := TAset[pass - 1]
		TAset[pass - 1] = TAset[idx]
		TAset[idx] = temp
		pass++
	}
}



// UI FUNCTIONS
func WelcomeLogin() {
	fmt.Println("=========================================================")
	fmt.Println("        SELAMAT DATANG APLIKASI MANAGEMENT DATA      ")
	fmt.Println("                   INVESTASI                            ")
	fmt.Println("=========================================================")
	fmt.Println("\nPilihan Menu:")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("3. Keluar Aplikasi")
	fmt.Println("---------------------------------------------------------")
	fmt.Print("\nMasukkan Pilihan: ")
}

func DashboardOption() {
	fmt.Println("\n\n-----------------Welcome To Dashboard------------------")
	fmt.Println()
	fmt.Println("=========================================================")
	fmt.Println("1. Tambah Data Investasi")
	fmt.Println("2. Modfikasi Data Investasi")
	fmt.Println("3. Hapus Data Investasi")
	fmt.Println("4. Cari Data Investasi")
	fmt.Println("5. Urutkan Data Investasi")
	fmt.Println("6. Tampilkan Laporan Data Investasi")
	fmt.Println("7. LogOut")
	fmt.Println("8. Keluar Aplikasi")
	fmt.Println("=========================================================")
	fmt.Print("Masukkan Pilihan: ")
}

func InputInvestasi() (string, string, float64, float64) {
	var namaAset, jenisAset string
	var jumlahDana, nilaiAset float64

	fmt.Println("=========================================================")
	fmt.Println("               Tambah Data Investasi                     ")
	fmt.Println("=========================================================")
	fmt.Print("Masukkan Nama Aset: ")
	fmt.Scan(&namaAset)
	fmt.Print("Masukkan Jenis Aset: ")
	fmt.Scan(&jenisAset)
	fmt.Print("Masukkan Jumlah Dana Investasi: ")
	fmt.Scan(&jumlahDana)
	fmt.Print("Masukkan Nilai  Aset: ")
	fmt.Scan(&nilaiAset)
	fmt.Println("=========================================================")
	return namaAset, jenisAset, jumlahDana, nilaiAset
}

// Implementation of search function
func CariDataInvestasi() {
	fmt.Println("=========================================================")
	fmt.Println("               Cari Data Investasi                     ")
	fmt.Println("=========================================================")
	fmt.Println("1. Cari berdasarkan Nama Aset")
	fmt.Println("2. Cari berdasarkan Jenis Aset")
	fmt.Println("3. Cari berdasarkan Jumlah Dana")
	fmt.Println("=========================================================")
	fmt.Print("Masukkan Pilihan: ")
	
	var pilihan int
	fmt.Scan(&pilihan)
	
	if pilihan == 1 {
		var namaAset string
		fmt.Print("Masukkan Nama Aset yang dicari: ")
		fmt.Scan(&namaAset)
		
		if FindDataByName(namaAset) {
			fmt.Println("Data aset dengan nama", namaAset, "ditemukan!")
			// Tampilkan detail aset
			for i := 0; i < idxInvestasi; i++ {
				if AllInvestasi[i].NamaAset == namaAset {
					fmt.Println("---------------------------------------------------------")
					fmt.Printf("Nama Aset: %s\n", AllInvestasi[i].NamaAset)
					fmt.Printf("Jenis Aset: %s\n", AllInvestasi[i].JenisAset)
					fmt.Printf("Jumlah Dana: %.2f\n", AllInvestasi[i].JumlahDanaInvestasi)
					fmt.Printf("Nilai Aset: %.2f\n", AllInvestasi[i].NilaiAset)
					fmt.Println("---------------------------------------------------------")
				}
			}
		} else {
			fmt.Println("Data aset dengan nama", namaAset, "tidak ditemukan!")
		}
	} else if pilihan == 2 {
		var jenisAset string
		fmt.Print("Masukkan Jenis Aset yang dicari: ")
		fmt.Scan(&jenisAset)
		
		if FindDataByJenis(jenisAset) {
			fmt.Println("Data aset dengan jenis", jenisAset, "ditemukan!")
			// Tampilkan semua aset dengan jenis tersebut
			fmt.Println("---------------------------------------------------------")
			fmt.Println("Daftar aset dengan jenis", jenisAset, ":")
			for i := 0; i < idxInvestasi; i++ {
				if AllInvestasi[i].JenisAset == jenisAset {
					fmt.Printf("Nama Aset: %s\n", AllInvestasi[i].NamaAset)
					fmt.Printf("Jumlah Dana: %.2f\n", AllInvestasi[i].JumlahDanaInvestasi)
					fmt.Printf("Nilai Aset: %.2f\n", AllInvestasi[i].NilaiAset)
					fmt.Println("---------------------------------------------------------")
				}
			}
		} else {
			fmt.Println("Data aset dengan jenis", jenisAset, "tidak ditemukan!")
		}
	} else if pilihan == 3 {
		var jumlahDana float64
		fmt.Print("Masukkan Jumlah Dana yang dicari: ")
		fmt.Scan(&jumlahDana)
		
		// Urutkan data terlebih dahulu
		UrutDataInvestasi(1) // Urutkan berdasarkan jumlah dana
		
		if FindByJumlahDanaAsc(jumlahDana) {
			fmt.Println("Data aset dengan jumlah dana", jumlahDana, "ditemukan!")
			// Tampilkan semua aset dengan jumlah dana tersebut
			fmt.Println("---------------------------------------------------------")
			for i := 0; i < idxInvestasi; i++ {
				if AllInvestasi[i].JumlahDanaInvestasi == jumlahDana {
					fmt.Printf("Nama Aset: %s\n", AllInvestasi[i].NamaAset)
					fmt.Printf("Jenis Aset: %s\n", AllInvestasi[i].JenisAset)
					fmt.Printf("Jumlah Dana: %.2f\n", AllInvestasi[i].JumlahDanaInvestasi)
					fmt.Printf("Nilai Aset: %.2f\n", AllInvestasi[i].NilaiAset)
					fmt.Println("---------------------------------------------------------")
				}
			}
		} else {
			fmt.Println("Data aset dengan jumlah dana", jumlahDana, "tidak ditemukan!")
		}
	} else {
		fmt.Println("Pilihan tidak valid!")
	}
}

// Implementation of sorting function
func UrutDataInvestasi(tipe int) {
	// Implementasi pengurutan dengan selection sort
	for i := 0; i < idxInvestasi-1; i++ {
		minIdx := i
		for j := i + 1; j < idxInvestasi; j++ {
			// Urutkan berdasarkan jumlah dana (ascending)
			if tipe == 1 && AllInvestasi[j].JumlahDanaInvestasi < AllInvestasi[minIdx].JumlahDanaInvestasi {
				minIdx = j
			}
			// Urutkan berdasarkan nilai aset (ascending)
			if tipe == 2 && AllInvestasi[j].NilaiAset < AllInvestasi[minIdx].NilaiAset {
				minIdx = j
			}
			// Urutkan berdasarkan nama aset (ascending/alphabetical)
			if tipe == 3 && AllInvestasi[j].NamaAset < AllInvestasi[minIdx].NamaAset {
				minIdx = j
			}
		}
		// Swap
		temp := AllInvestasi[minIdx]
		AllInvestasi[minIdx] = AllInvestasi[i]
		AllInvestasi[i] = temp
	}
}

// Function to display all investment data
func TampilkanLaporanInvestasi() {
	fmt.Println("=========================================================")
	fmt.Println("               Laporan Data Investasi                   ")
	fmt.Println("=========================================================")
	
	if idxInvestasi == 0 {
		fmt.Println("Belum ada data investasi.")
		return
	}
	
	var totalDana, totalNilai float64
	
	fmt.Println("No.\tNama Aset\t\tJenis\t\tJumlah Dana\t\tNilai Aset")
	fmt.Println("---------------------------------------------------------")
	
	for i := 0; i < idxInvestasi; i++ {
		fmt.Printf("%d\t%-20s\t%-10s\t%.2f\t\t%.2f\n", 
			i+1, 
			AllInvestasi[i].NamaAset, 
			AllInvestasi[i].JenisAset, 
			AllInvestasi[i].JumlahDanaInvestasi, 
			AllInvestasi[i].NilaiAset)
			
		totalDana += AllInvestasi[i].JumlahDanaInvestasi
		totalNilai += AllInvestasi[i].NilaiAset
	}
	
	fmt.Println("---------------------------------------------------------")
	fmt.Printf("Total Dana Investasi: %.2f\n", totalDana)
	fmt.Printf("Total Nilai Aset: %.2f\n", totalNilai)
	
	// Hitung profit/loss
	profitLoss := totalNilai - totalDana
	if profitLoss >= 0 {
		fmt.Printf("Profit: %.2f (%.2f%%)\n", profitLoss, (profitLoss/totalDana)*100)
	} else {
		fmt.Printf("Loss: %.2f (%.2f%%)\n", profitLoss, (profitLoss/totalDana)*100)
	}
	fmt.Println("=========================================================")
}

// Implementation for sorting menu
func MenuUrutInvestasi() {
	fmt.Println("=========================================================")
	fmt.Println("            Urutkan Data Investasi                     ")
	fmt.Println("=========================================================")
	fmt.Println("1. Urutkan berdasarkan Jumlah Dana")
	fmt.Println("2. Urutkan berdasarkan Nilai Aset")
	fmt.Println("3. Urutkan berdasarkan Nama Aset")
	fmt.Println("=========================================================")
	fmt.Print("Masukkan Pilihan: ")
	
	var pilihan int
	fmt.Scan(&pilihan)
	
	if pilihan >= 1 && pilihan <= 3 {
		UrutDataInvestasi(pilihan)
		fmt.Println("Data berhasil diurutkan!")
		// Tampilkan data yang sudah diurutkan
		TampilkanLaporanInvestasi()
	} else {
		fmt.Println("Pilihan tidak valid!")
	}
}
