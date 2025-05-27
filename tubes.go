package main

import "fmt"

// CONSTANTS
const NMAX int = 1000

// DATA STRUCTURES
type Users struct {
	Nama, Email, Password string
}

type AsetInvestasi struct {
	NamaAset, JenisAset              string
	JumlahDanaInvestasi, NilaiAset   float64
	HargaBeli, HargaJual, Keuntungan float64
	PersentaseKeuntungan             float64
}

// GLOBAL VARIABLES
type DataUsers = [NMAX]Users
type DataInvestasi = [NMAX]AsetInvestasi

var AllUsers DataUsers
var AllInvestasi DataInvestasi

var idxUser = 2
var idxInvestasi = 6

func main() {
	// memasukan dataDummy untuk user default ke array
	InitDummyUsers(&AllUsers)
	// memasukan dataDumy untuk Data Investasi default ke array
	InitDummyDataInvestasi(&AllInvestasi)

	var aplikasiAktif bool = true

	for aplikasiAktif {
		var accountOption int
		var isLoggIn bool
		WelcomeLogin()
		fmt.Scan(&accountOption)
		if accountOption == 1 {
			isLoggIn = InputLogin()

		} else if accountOption == 2 {
			RegisterUser(&AllUsers, idxUser)
			var input string
			fmt.Println("Ingin Login? (Y/N):")
			fmt.Scan(&input)
			if input == "Y" || input == "y" {
				isLoggIn = InputLogin()
			}
		} else if accountOption == 3 {
			GoodBye()
			aplikasiAktif = false
		} else {
			fmt.Println("Pilihan Tidak valid. Silahkan Pilih 1, 2, atau 3.")
		}

		var inDashboard bool = true
		for isLoggIn && inDashboard {
			var dashboradOption int
			var namaAset, jenisAset string
			var jumlahDana, nilaiAset, hargaBeli, hargaJual float64

			DashboardOption()
			fmt.Scan(&dashboradOption)

			switch dashboradOption {
			case 1:
				namaAset, jenisAset, jumlahDana, nilaiAset, hargaBeli, hargaJual = InputInvestasi()
				CreateDataInvestasi(namaAset, jenisAset, jumlahDana, nilaiAset, hargaBeli, hargaJual)
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
				fmt.Println("1. Nama Aset --Sequential Search")
				fmt.Println("2. Jenis Aset --Sequential Search")
				fmt.Println("3. Jumlah Dana --Binary Search")
				fmt.Println("---------------------------------------------------------")
				var input int
				fmt.Print("Masukan Piihan(1-3): ")
				fmt.Scan(&input)

				switch input {
				case 1:
					var namaAset string
					fmt.Print("Masukkan Nama Aset: ")
					fmt.Scan(&namaAset)
					found := FindDataByName(namaAset)
					if found == -1 {
						fmt.Println("\nData Tidak Ditemukan,Pastikan Nama Aset Yang di Masukan Benar")
					} else {
						fmt.Printf("\nTotal data dengan nama '%s': %d\n", namaAset, found)
					}
				case 2:
					var jenisAset string
					fmt.Print("Masukkan Jenis Aset: ")
					fmt.Scan(&jenisAset)
					found := FindDataByJenis(jenisAset)
					if found == -1 {
						fmt.Print("\nData Tidak Ditemukan, Pastikan Jenis Aset yang Dimasukkan Benar.\n")
					} else {
						fmt.Printf("\nTotal data dengan jenis '%s': %d\n", jenisAset, found)
					}
				case 3:
					var jumlahDana float64
					fmt.Print("Masukkan Jumlah Dana: ")
					fmt.Scan(&jumlahDana)
					InsertionSortAscendingJumlahDana(&AllInvestasi)
					found := FindByJumlahDana(jumlahDana)
					if found == -1 {
						fmt.Print("\nData Tidak Ditemukan,Pastikan Jumlah Dana Yang di Masukan Benar")
					}
				default:
					fmt.Println("Pilihan Tidak Valid, Silahkan Pilih 1,2, atau 3")
				}
			case 5:
				var inputBerdasarkan int
				var inputJenisSort int
				fmt.Println("========================================")
				fmt.Println("          PENGURUTAN DATA INVESTASI     ")
				fmt.Println("========================================")
				fmt.Println("Pilih Jenis Pengurutan:")
				fmt.Println("1. Ascending (Naik)")
				fmt.Println("2. Descending (Turun)")
				fmt.Print("Masukkan pilihan (1/2): ")
				fmt.Scan(&inputJenisSort)

				if inputJenisSort == 1 {
					fmt.Println("\nUrutkan data investasi berdasarkan:")
					fmt.Println("1. Nama Aset --SelectionSort")
					fmt.Println("2. Jenis Aset --SelectionSort")
					fmt.Println("3. Jumlah Dana --InsertionSort")
					fmt.Println("4. Keuntungan --InsertionSort")
					fmt.Println("5. Persetase Keuntungan --InsertionSort")
					fmt.Print("Masukkan pilihan (1-5): ")
					fmt.Scan(&inputBerdasarkan)
					switch inputBerdasarkan {
					case 1:
						fmt.Print("\nData Sebelum di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
						SelectionSortAscendingNamaAset(&AllInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
					case 2:
						fmt.Print("\nData Sebelum di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
						SelectionSortAscendingJenisAset(&AllInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
					case 3:
						fmt.Print("\nData Sebelum di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
						InsertionSortAscendingJumlahDana(&AllInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
					case 4:
						InitPersentaseDanKeuntungan(&AllInvestasi)
						fmt.Print("\nData Sebelum di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
						InsertionSortAscendingKeuntungan(&AllInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
					case 5:
						InitPersentaseDanKeuntungan(&AllInvestasi)
						fmt.Print("\nData Sebelum di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
						InsertionSortAscendingPersentaseKeuntungan(&AllInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
					}
				} else if inputJenisSort == 2 {
					fmt.Println("\nUrutkan data investasi berdasarkan:")
					fmt.Println("1. Nama Aset --SelectionSort")
					fmt.Println("2. Jenis Aset --SelectionSort")
					fmt.Println("3. Jumlah Dana --InsertionSort")
					fmt.Println("4. Keuntungan --InsertionSort")
					fmt.Println("5. Persetase Keuntungan --InsertionSort")
					fmt.Print("Masukkan pilihan (1-5): ")
					fmt.Scan(&inputBerdasarkan)
					switch inputBerdasarkan {
					case 1:
						fmt.Print("\nData Sebelum di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
						SelectionSortDescendingNamaAset(&AllInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
					case 2:
						fmt.Print("\nData Sebelum di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
						SelectionSortDescendingJenisAset(&AllInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
					case 3:
						fmt.Print("\nData Sebelum di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
						InsertionSortDescendingJumlahDana(&AllInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
					case 4:
						InitPersentaseDanKeuntungan(&AllInvestasi)
						fmt.Print("\nData Sebelum di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
						InsertionSortDescendingKeuntungan(&AllInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
					case 5:
						InitPersentaseDanKeuntungan(&AllInvestasi)
						fmt.Print("\nData Sebelum di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
						InsertionSortDescendingPersentaseKeuntungan(&AllInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&AllInvestasi)
					}
				} else {
					fmt.Println("Pilihan tidak valid! Silahkan pilih 1-2.")
				}
			case 6:
				CetakLaporanInvestasi(&AllInvestasi)
			case 7:
				fmt.Println("Berhasil Logout.")
				inDashboard = false
			case 8:
				GoodBye()
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

func InitDummyDataInvestasi(TAset *DataInvestasi) {
	// IS. Terdefinisi array TAset dalam keadaan kosong tanpa data aset investasi apapun.
	// FS. Mengisi array TAset dengan beberapa data dummy aset investasi yang sudah ditentukan
	//     pada indeks awal array untuk tujuan pengujian atau inisialisasi awal program.
	TAset[0] = AsetInvestasi{
		NamaAset:            "ADRO",
		JenisAset:           "Saham",
		NilaiAset:           3000000,
		JumlahDanaInvestasi: 90000000,
		HargaBeli:           3000000,
		HargaJual:           3200000,
	}
	TAset[1] = AsetInvestasi{
		NamaAset:            "BBRI",
		JenisAset:           "Saham",
		NilaiAset:           1500000,
		JumlahDanaInvestasi: 50000000,
		HargaBeli:           1500000,
		HargaJual:           1800000,
	}
	TAset[2] = AsetInvestasi{
		NamaAset:            "TLKM",
		JenisAset:           "Saham",
		NilaiAset:           2000000,
		JumlahDanaInvestasi: 75000000,
		HargaBeli:           2000000,
		HargaJual:           2400000,
	}
	TAset[3] = AsetInvestasi{
		NamaAset:            "Mandiri-Investa-Pasar-Uang",
		JenisAset:           "Reksadana",
		NilaiAset:           100000,
		JumlahDanaInvestasi: 30000000,
		HargaBeli:           100000,
		HargaJual:           103000,
	}
	TAset[4] = AsetInvestasi{
		NamaAset:            "Schroder-Dana-Likuid",
		JenisAset:           "Reksadana",
		NilaiAset:           125000,
		JumlahDanaInvestasi: 40000000,
		HargaBeli:           125000,
		HargaJual:           127500,
	}
	TAset[5] = AsetInvestasi{
		NamaAset:            "ORI022",
		JenisAset:           "Obligasi",
		NilaiAset:           1000000,
		JumlahDanaInvestasi: 100000000,
		HargaBeli:           1000000,
		HargaJual:           1050000,
	}
	TAset[6] = AsetInvestasi{
		NamaAset:            "SR018",
		JenisAset:           "Obligasi",
		NilaiAset:           1000000,
		JumlahDanaInvestasi: 85000000,
		HargaBeli:           1000000,
		HargaJual:           1080000,
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
func CreateDataInvestasi(namaAset, jenisAset string, nilaiAset, jumlahDana, hargaBeli, hargaJual float64) {
	// IS. Terdefinisi suatu array AllInvestasi yang berukuran maksimum NMAX dan telah terisi sebanyak idxInvestasi elemen.
	//      Diberikan input namaAset, jenisAset, nilaiAset, dan danaInvestasi yang akan dimasukkan ke dalam array.
	// FS. Menambahkan data aset investasi ke dalam array AllInvestasi apabila belum penuh (idxInvestasi < NMAX),
	//      atau menampilkan pesan "Data Investasi Sudah Penuh" apabila array telah mencapai kapasitas maksimum.

	dataInvestasi := AsetInvestasi{
		NamaAset:            namaAset,
		JenisAset:           jenisAset,
		NilaiAset:           nilaiAset,
		JumlahDanaInvestasi: jumlahDana,
		HargaBeli:           hargaBeli,
		HargaJual:           hargaJual,
	}
	if idxInvestasi < NMAX {
		AllInvestasi[idxInvestasi] = dataInvestasi
		fmt.Println("---------------------------------------------------------")
		fmt.Println("        Data Investasi Berhasil Ditambahkan       ")
		fmt.Println("---------------------------------------------------------")
		fmt.Println("Nama Aset          :", dataInvestasi.NamaAset)
		fmt.Println("Jenis Aset         :", dataInvestasi.JenisAset)
		fmt.Printf("Jumlah Dana        : %.2f\n", dataInvestasi.JumlahDanaInvestasi)
		fmt.Printf("Nilai Aset         : %.2f\n", dataInvestasi.NilaiAset)
		fmt.Printf("Harga Beli         : %.2f\n", dataInvestasi.HargaBeli)
		fmt.Printf("Harga Jual         : %.2f\n", dataInvestasi.HargaJual)
		fmt.Println("---------------------------------------------------------")
		idxInvestasi++
	} else {
		fmt.Println("\n=========================================================")
		fmt.Println("    Gagal! Data Investasi Sudah Penuh       ")
		fmt.Println("=========================================================")
	}
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
			fmt.Printf("Jumlah Dana: %.2f\n", AllInvestasi[i].JumlahDanaInvestasi)
			fmt.Printf("Nilai Aset: %.2f\n", AllInvestasi[i].NilaiAset)
			fmt.Printf("Harga Beli: %.2f\n", AllInvestasi[i].HargaBeli)
			fmt.Printf("Harga Jual: %.2f\n", AllInvestasi[i].HargaJual)
			fmt.Println("---------------------------------------------------------")

			// Ambil input data baru
			var namaAsetBaru, jenisAsetBaru string
			var jumlahDanaBaru, nilaiAsetBaru, hargaBeliBaru, hargaJualBaru float64

			fmt.Print("\nMasukkan Data Baru: \nJika Nama atau Aset lebih dari 1 kata gunakan '-' sebagai pemisah !!\n")
			fmt.Print("Nama Aset: ")
			fmt.Scan(&namaAsetBaru)
			fmt.Print("Jenis Aset: ")
			fmt.Scan(&jenisAsetBaru)
			fmt.Print("Jumlah Dana: ")
			fmt.Scan(&jumlahDanaBaru)
			fmt.Print("Nilai Aset: ")
			fmt.Scan(&nilaiAsetBaru)
			fmt.Print("Harga Beli : ")
			fmt.Scan(&hargaBeliBaru)
			fmt.Print("Harga Jual : ")
			fmt.Scan(&hargaJualBaru)

			// Update data
			AllInvestasi[i].NamaAset = namaAsetBaru
			AllInvestasi[i].JenisAset = jenisAsetBaru
			AllInvestasi[i].JumlahDanaInvestasi = jumlahDanaBaru
			AllInvestasi[i].NilaiAset = nilaiAsetBaru
			AllInvestasi[i].HargaBeli = hargaBeliBaru
			AllInvestasi[i].HargaJual = hargaJualBaru

			fmt.Println("---------------------------------------------------------")
			fmt.Println("Data investasi berhasil dimodifikasi.")
			fmt.Println("=========================================================")
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
		fmt.Printf("Harga Beli  : %.2f\n", dataAset.HargaBeli)
		fmt.Printf("harga Jual  : %.2f\n", dataAset.HargaJual)
		fmt.Println("---------------------------------------------------------")
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Data Aset tidak ditemukan.")
	}
}

// sequential search
func FindDataByName(nameAset string) int {
	// {diberikan array AllInvestasi yang berisi idxInvestasi data aset dan sebuah nama aset,
	//  untuk mengembalikan indeks data apabila nama aset ditemukan pada array,
	//  serta menampilkan detail data aset tersebut; jika tidak ditemukan, mengembalikan -1}
	found := 0
	for i := 0; i < idxInvestasi; i++ {
		if AllInvestasi[i].NamaAset == nameAset {
			fmt.Println("===========================")
			fmt.Printf("Data Ditemukan di Index: %d\n", i)
			fmt.Println("===========================")
			fmt.Printf("Nama Aset      : %s\n", AllInvestasi[i].NamaAset)
			fmt.Printf("Jenis Aset     : %s\n", AllInvestasi[i].JenisAset)
			fmt.Printf("Nilai Aset     : %.2f\n", AllInvestasi[i].NilaiAset)
			fmt.Printf("Jumlah Dana    : %.2f\n", AllInvestasi[i].JumlahDanaInvestasi)
			fmt.Printf("Harga Beli    : %.2f\n", AllInvestasi[i].HargaBeli)
			fmt.Printf("Harga Jual    : %.2f\n", AllInvestasi[i].HargaJual)
			fmt.Println("===========================\n")
			found++
		}
	}
	if found == 0 {
		return -1
	}
	return found
}

// sequential search
func FindDataByJenis(jenisAset string) int {
	// {diberikan array AllInvestasi yang berisi idxInvestasi data aset dan sebuah jenis aset,
	//  untuk mengembalikan indeks data apabila jenis aset ditemukan pada array,
	//  serta menampilkan detail data aset tersebut; jika tidak ditemukan, mengembalikan -1}
	found := 0
	for i := 0; i < idxInvestasi; i++ {
		if AllInvestasi[i].JenisAset == jenisAset {
			fmt.Println("===========================")
			fmt.Printf("Data Ditemukan di Index: %d\n", i)
			fmt.Println("===========================")
			fmt.Printf("Nama Aset      : %s\n", AllInvestasi[i].NamaAset)
			fmt.Printf("Jenis Aset     : %s\n", AllInvestasi[i].JenisAset)
			fmt.Printf("Nilai Aset     : %.2f\n", AllInvestasi[i].NilaiAset)
			fmt.Printf("Jumlah Dana    : %.2f\n", AllInvestasi[i].JumlahDanaInvestasi)
			fmt.Printf("Harga Beli    : %.2f\n", AllInvestasi[i].HargaBeli)
			fmt.Printf("Harga Jual    : %.2f\n", AllInvestasi[i].HargaJual)
			fmt.Println("===========================\n")
			found++
		}
	}
	if found == 0 {
		return -1
	}
	return found
}

// binary search
func FindByJumlahDana(jumlahDana float64) int {
	n := idxInvestasi
	left := 0
	right := n - 1
	for left <= right {
		mid := (left + right) / 2
		midValue := AllInvestasi[mid].JumlahDanaInvestasi

		if midValue == jumlahDana {
			// Data ditemukan
			fmt.Println("===========================")
			fmt.Printf("Data Ditemukan di Index: %d\n", mid)
			fmt.Println("===========================")
			fmt.Printf("Nama Aset      : %s\n", AllInvestasi[mid].NamaAset)
			fmt.Printf("Jenis Aset     : %s\n", AllInvestasi[mid].JenisAset)
			fmt.Printf("Nilai Aset     : %.2f\n", AllInvestasi[mid].NilaiAset)
			fmt.Printf("Jumlah Dana    : %.2f\n", AllInvestasi[mid].JumlahDanaInvestasi)
			fmt.Printf("Harga Beli    : %.2f\n", AllInvestasi[mid].HargaBeli)
			fmt.Printf("Harga Jual    : %.2f\n", AllInvestasi[mid].HargaJual)
			fmt.Println("===========================\n")
			return mid
		} else if midValue < jumlahDana {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// Data tidak ditemukan
	return -1
}

func SelectionSortAscendingNamaAset(TAset *DataInvestasi) {
	var i, idx, pass int
	pass = 1
	for pass < idxInvestasi {
		idx = pass - 1
		i = pass
		for i < idxInvestasi {
			if TAset[i].NamaAset < TAset[idx].NamaAset {
				idx = i
			}
			i++
		}
		temp := TAset[pass-1]
		TAset[pass-1] = TAset[idx]
		TAset[idx] = temp
		pass++
	}
}
func SelectionSortDescendingNamaAset(TAset *DataInvestasi) {
	var i, idx, pass int
	pass = 1
	for pass < idxInvestasi {
		idx = pass - 1
		i = pass
		for i < idxInvestasi {
			if TAset[i].NamaAset > TAset[idx].NamaAset {
				idx = i
			}
			i++
		}
		temp := TAset[pass-1]
		TAset[pass-1] = TAset[idx]
		TAset[idx] = temp
		pass++
	}
}
func SelectionSortAscendingJenisAset(TAset *DataInvestasi) {
	var i, idx, pass int
	pass = 1
	for pass < idxInvestasi {
		idx = pass - 1
		i = pass
		for i < idxInvestasi {
			if TAset[i].JenisAset < TAset[idx].JenisAset {
				idx = i
			}
			i++
		}
		temp := TAset[pass-1]
		TAset[pass-1] = TAset[idx]
		TAset[idx] = temp
		pass++
	}
}
func SelectionSortDescendingJenisAset(TAset *DataInvestasi) {
	var i, idx, pass int
	pass = 1
	for pass < idxInvestasi {
		idx = pass - 1
		i = pass
		for i < idxInvestasi {
			if TAset[i].JenisAset > TAset[idx].JenisAset {
				idx = i
			}
			i++
		}
		temp := TAset[pass-1]
		TAset[pass-1] = TAset[idx]
		TAset[idx] = temp
		pass++
	}
}
func InsertionSortAscendingJumlahDana(TAset *DataInvestasi) {

	for i := 1; i < idxInvestasi; i++ {
		key := TAset[i]
		j := i - 1

		// Pindahkan elemen yang lebih besar dari key
		for j >= 0 && TAset[j].JumlahDanaInvestasi > key.JumlahDanaInvestasi {
			TAset[j+1] = TAset[j]
			j--
		}
		TAset[j+1] = key
	}
}
func InsertionSortDescendingJumlahDana(TAset *DataInvestasi) {

	for i := 1; i < idxInvestasi; i++ {
		key := TAset[i]
		j := i - 1

		// Pindahkan elemen yang lebih besar dari key
		for j >= 0 && TAset[j].JumlahDanaInvestasi < key.JumlahDanaInvestasi {
			TAset[j+1] = TAset[j]
			j--
		}
		TAset[j+1] = key
	}
}
func HitungKeuntungan(Aset AsetInvestasi) float64 {
	jumlahUnit := Aset.JumlahDanaInvestasi / Aset.HargaBeli
	keuntungan := (Aset.HargaJual - Aset.HargaBeli) * jumlahUnit
	return keuntungan
}
func HitungPersentaseKeuntungan(Aset AsetInvestasi) float64 {
	if Aset.HargaBeli == 0 {
		return 0 // menghindari pembagian dengan nol
	}
	jumlahUnit := Aset.JumlahDanaInvestasi / Aset.HargaBeli
	keuntungan := (Aset.HargaJual - Aset.HargaBeli) * jumlahUnit
	return (keuntungan / Aset.JumlahDanaInvestasi) * 100
}
func InitPersentaseDanKeuntungan(TAset *DataInvestasi) {
	for i := 0; i < idxInvestasi; i++ {
		TAset[i].Keuntungan = HitungKeuntungan(TAset[i])
		TAset[i].PersentaseKeuntungan = HitungPersentaseKeuntungan(TAset[i])
	}
}
func InsertionSortAscendingKeuntungan(TAset *DataInvestasi) {
	for i := 1; i < idxInvestasi; i++ {
		key := TAset[i]
		j := i - 1
		// Pindahkan elemen yang lebih besar dari key
		for j >= 0 && TAset[j].Keuntungan > key.Keuntungan {
			TAset[j+1] = TAset[j]
			j--
		}
		TAset[j+1] = key
	}
}
func InsertionSortDescendingKeuntungan(TAset *DataInvestasi) {
	for i := 1; i < idxInvestasi; i++ {
		key := TAset[i]
		j := i - 1

		// Pindahkan elemen yang lebih besar dari key
		for j >= 0 && TAset[j].Keuntungan < key.Keuntungan {
			TAset[j+1] = TAset[j]
			j--
		}
		TAset[j+1] = key
	}
}
func InsertionSortAscendingPersentaseKeuntungan(TAset *DataInvestasi) {
	for i := 1; i < idxInvestasi; i++ {
		key := TAset[i]
		j := i - 1

		// Pindahkan elemen yang lebih besar dari key
		for j >= 0 && TAset[j].PersentaseKeuntungan > key.PersentaseKeuntungan {
			TAset[j+1] = TAset[j]
			j--
		}
		TAset[j+1] = key
	}
}
func InsertionSortDescendingPersentaseKeuntungan(TAset *DataInvestasi) {
	for i := 1; i < idxInvestasi; i++ {
		key := TAset[i]
		j := i - 1

		// Pindahkan elemen yang lebih besar dari key
		for j >= 0 && TAset[j].PersentaseKeuntungan < key.PersentaseKeuntungan {
			TAset[j+1] = TAset[j]
			j--
		}
		TAset[j+1] = key
	}
}

func CetakDataInvestasi(TAset *DataInvestasi) {
	fmt.Println("=================================================================================================================================")
	fmt.Printf("| %-3s | %-30s | %-10s | %-15s | %-12s | %-12s | %-12s | %-10s |\n",
		"No", "Nama Aset", "Jenis", "Jumlah Dana", "Harga Beli", "Harga Jual", "Keuntungan", "% Untung")
	fmt.Println("=================================================================================================================================")
	for i := 0; i < idxInvestasi; i++ {
		fmt.Printf("| %-3d | %-30s | %-10s | %-15.2f | %-12.2f | %-12.2f | %-12.2f | %-10.2f |\n",
			i+1,
			TAset[i].NamaAset,
			TAset[i].JenisAset,
			TAset[i].JumlahDanaInvestasi,
			TAset[i].HargaBeli,
			TAset[i].HargaJual,
			TAset[i].Keuntungan,
			TAset[i].PersentaseKeuntungan,
		)
	}
	fmt.Println("=================================================================================================================================")
}

func CetakLaporanInvestasi(TAset *DataInvestasi) {
	// IS: Terdefinisi array TAset yang berisi data aset investasi sebanyak idxInvestasi
	// FS: Menampilkan laporan portofolio investasi yang meliputi:
	//     - Total investasi, total keuntungan, dan rata-rata keuntungan
	//     - Aset paling untung dan paling rugi
	//     - Komposisi portofolio berdasarkan jenis aset

	if idxInvestasi == 0 {
		fmt.Println("=========================================================")
		fmt.Println("              TIDAK ADA DATA INVESTASI                  ")
		fmt.Println("=========================================================")
		return
	}

	// Inisialisasi perhitungan keuntungan dan persentase
	InitPersentaseDanKeuntungan(TAset)

	// Variabel untuk perhitungan
	var totalInvestasi float64 = 0
	var totalKeuntungan float64 = 0
	var asetPalingUntung, asetPalingRugi AsetInvestasi
	var maxKeuntungan float64 = TAset[0].Keuntungan
	var minKeuntungan float64 = TAset[0].Keuntungan

	// Variabel untuk komposisi portofolio
	var totalSaham, totalReksadana, totalObligasi float64 = 0, 0, 0

	// Perhitungan total dan pencarian aset terbaik/terburuk
	for i := 0; i < idxInvestasi; i++ {
		aset := TAset[i]

		// Total investasi dan keuntungan
		totalInvestasi += aset.JumlahDanaInvestasi
		totalKeuntungan += aset.Keuntungan

		// Cari aset paling untung dan paling rugi
		if aset.Keuntungan > maxKeuntungan {
			maxKeuntungan = aset.Keuntungan
			asetPalingUntung = aset
		}
		if aset.Keuntungan < minKeuntungan {
			minKeuntungan = aset.Keuntungan
			asetPalingRugi = aset
		}

		// Komposisi berdasarkan jenis aset
		switch aset.JenisAset {
		case "Saham":
			totalSaham += aset.JumlahDanaInvestasi
		case "Reksadana":
			totalReksadana += aset.JumlahDanaInvestasi
		case "Obligasi":
			totalObligasi += aset.JumlahDanaInvestasi
		}
	}

	// Hitung rata-rata keuntungan
	rataRataKeuntungan := (totalKeuntungan / totalInvestasi) * 100

	// Hitung persentase komposisi
	persentaseSaham := (totalSaham / totalInvestasi) * 100
	persentaseReksadana := (totalReksadana / totalInvestasi) * 100
	persentaseObligasi := (totalObligasi / totalInvestasi) * 100

	// Tampilkan laporan
	fmt.Println()
	fmt.Println("=========================================================")
	fmt.Println("              Rangkuman Portofolio Investasi             ")
	fmt.Println("=========================================================")
	fmt.Printf("Total Investasi       : Rp %.3f\n", totalInvestasi)
	fmt.Printf("Total Keuntungan      : Rp %.3f\n", totalKeuntungan)
	fmt.Printf("Rata-rata Keuntungan  : %.2f%%\n", rataRataKeuntungan)
	fmt.Println()

	fmt.Printf("Aset Paling Untung    : %s (Rp %.3f)\n", asetPalingUntung.NamaAset, maxKeuntungan)
	fmt.Printf("Aset Paling Rugi      : %s (Rp %.3f)\n", asetPalingRugi.NamaAset, minKeuntungan)
	fmt.Println()

	fmt.Println("Komposisi Portofolio:")
	fmt.Printf("- Saham      : %.3f (%.2f%%)\n", totalSaham, persentaseSaham)
	fmt.Printf("- Reksadana  : %.3f (%.2f%%)\n", totalReksadana, persentaseReksadana)
	fmt.Printf("- Obligasi   : %.3f (%.2f%%)\n", totalObligasi, persentaseObligasi)
	fmt.Println("=========================================================")
}

// UI
func WelcomeLogin() {
	fmt.Println("=========================================================")
	fmt.Println("        	SELAMAT DATANG DI APLIKASI       			  ")
	fmt.Println("        	 MANAGEMENT DATA INVESTASI\n\nUser Default:\nEmail: son.com | password: son\nEmail: keanu.com | password: nu                    ")
	fmt.Println("=========================================================")
	fmt.Println("\nPilihan Menu:")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("3. Keluar Aplikasi")
	fmt.Println("---------------------------------------------------------")
	fmt.Print("\nMasukkan Pilihan(1-3): ")
}
func DashboardOption() {
	fmt.Println("\n\n-----------------Welcome To Dashboard------------------")
	fmt.Println()
	fmt.Println("=========================================================")
	fmt.Println("1. Tambah Data Investasi")
	fmt.Println("2. Modifikasi Data Investasi")
	fmt.Println("3. Hapus Data Investasi")
	fmt.Println("4. Cari Data Investasi")
	fmt.Println("5. Urutkan Data Investasi")
	fmt.Println("6. Tampilkan Laporan Data Investasi")
	fmt.Println("7. Keluar Dashboard")
	fmt.Println("8. Keluar Aplikasi")
	fmt.Println("=========================================================")
	fmt.Print("Masukkan Pilihan(1-8): ")
}
func InputInvestasi() (string, string, float64, float64, float64, float64) {
	var namaAset, jenisAset string
	var jumlahDana, nilaiAset, hargaBeli, hargaJual float64

	fmt.Println("=========================================================")
	fmt.Println("               Tambah Data Investasi                     ")
	fmt.Println("Jika Nama Atau Jenis Aset Lebih Dari 1 Kata Gunakan '-' Untuk Spasi")
	fmt.Println("=========================================================")
	fmt.Print("Masukkan Nama Aset: ")
	fmt.Scan(&namaAset)
	fmt.Print("Masukkan Jenis Aset: ")
	fmt.Scan(&jenisAset)
	fmt.Print("Masukkan Jumlah Dana Investasi: ")
	fmt.Scan(&jumlahDana)
	fmt.Print("Masukkan Nilai  Aset: ")
	fmt.Scan(&nilaiAset)
	fmt.Print("Masukkan Harga Beli: ")
	fmt.Scan(&hargaBeli)
	fmt.Print("Masukkan Harga Jual: ")
	fmt.Scan(&hargaJual)
	fmt.Println("=========================================================")
	return namaAset, jenisAset, jumlahDana, nilaiAset, hargaBeli, hargaJual
}
func GoodBye() {
	fmt.Println("=========================================================")
	fmt.Println("               TERIMA KASIH TELAH MENGGUNAKAN            ")
	fmt.Println("                	APLIKASI INI  :)           ")
	fmt.Println("=========================================================")
	fmt.Println("                      Sampai Jumpa!                      ")
	fmt.Println("=========================================================")
}
