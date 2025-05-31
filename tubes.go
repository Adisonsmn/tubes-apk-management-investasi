package main

import "fmt"

// Variabel global
const NMAX int = 1000

type Users struct {
	Nama, Email, Password string
}
type AsetInvestasi struct {
	NamaAset, JenisAset              string
	JumlahDanaInvestasi, NilaiAset   float64
	HargaBeli, HargaJual, Keuntungan float64
	PersentaseKeuntungan             float64
}
type tabUser = [NMAX]Users
type tabInvestasi = [NMAX]AsetInvestasi

func main() {
	var dataUsers tabUser
	var dataInvestasi tabInvestasi
	var idxUser = 2
	var idxInvestasi = 6

	// memasukan dataDummy untuk user default ke array
	InitDummyUsers(&dataUsers)
	// memasukan dataDummy untuk Data Investasi  ke array
	InitDummytabInvestasi(&dataInvestasi)
	InitPersentaseDanKeuntungan(&dataInvestasi, idxInvestasi)

	var aplikasiAktif bool = true

	for aplikasiAktif {
		var accountOption int
		var isLoggIn bool
		WelcomeLogin()
		fmt.Scan(&accountOption)
		if accountOption == 1 {
			isLoggIn = InputLogin(dataUsers, idxUser)

		} else if accountOption == 2 {
			RegisterUser(&dataUsers, &idxUser)
			var input string
			fmt.Println("Ingin Login? (Y/N):")
			fmt.Scan(&input)
			if input == "Y" || input == "y" {
				isLoggIn = InputLogin(dataUsers, idxUser)
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

			DashboardOption()
			fmt.Scan(&dashboradOption)

			switch dashboradOption {
			case 1:
				CreateDataInvesatasi(&dataInvestasi, &idxInvestasi)
				InitPersentaseDanKeuntungan(&dataInvestasi, idxInvestasi)
			case 2:
				fmt.Print("Masukkan Nama Aset yang akan di ubah: ")
				fmt.Scan(&namaAset)
				ModifyDataInvestasi(&dataInvestasi, &idxInvestasi, namaAset)
			case 3:
				fmt.Print("Masukkan Nama Aset yang akan di hapus: ")
				fmt.Scan(&namaAset)
				HapusDataInvestasi(&dataInvestasi, &idxInvestasi, namaAset)
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
					found := FindDataByName(dataInvestasi, idxInvestasi, namaAset)
					if found == -1 {
						fmt.Println("\nData Tidak Ditemukan,Pastikan Nama Aset Yang di Masukan Benar")
					} else {
						fmt.Printf("\nTotal data dengan nama '%s': %d\n", namaAset, found)
					}
				case 2:
					fmt.Print("Masukkan Jenis Aset: ")
					fmt.Scan(&jenisAset)
					found := FindDataByJenis(dataInvestasi, idxInvestasi, jenisAset)
					if found == -1 {
						fmt.Print("\nData Tidak Ditemukan, Pastikan Jenis Aset yang Dimasukkan Benar.\n")
					} else {
						fmt.Printf("\nTotal data dengan jenis '%s': %d\n", jenisAset, found)
					}
				case 3:
					var jumlahDana float64
					fmt.Print("Masukkan Jumlah Dana: ")
					fmt.Scan(&jumlahDana)
					InsertionSortAscendingJumlahDana(&dataInvestasi, idxInvestasi)
					found := FindByJumlahDana(dataInvestasi, idxInvestasi, jumlahDana)
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
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
						SelectionSortAscendingNamaAset(&dataInvestasi, idxInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
					case 2:
						fmt.Print("\nData Sebelum di Urutkan\n")
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
						SelectionSortAscendingJenisAset(&dataInvestasi, idxInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
					case 3:
						fmt.Print("\nData Sebelum di Urutkan\n")
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
						InsertionSortAscendingJumlahDana(&dataInvestasi, idxInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
					case 4:
						InitPersentaseDanKeuntungan(&dataInvestasi, idxInvestasi)
						fmt.Print("\nData Sebelum di Urutkan\n")
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
						InsertionSortAscendingKeuntungan(&dataInvestasi, idxInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
					case 5:
						InitPersentaseDanKeuntungan(&dataInvestasi, idxInvestasi)
						fmt.Print("\nData Sebelum di Urutkan\n")
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
						InsertionSortAscendingPersentaseKeuntungan(&dataInvestasi, idxInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
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
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
						SelectionSortDescendingNamaAset(&dataInvestasi, idxInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
					case 2:
						fmt.Print("\nData Sebelum di Urutkan\n")
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
						SelectionSortDescendingJenisAset(&dataInvestasi, idxInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
					case 3:
						fmt.Print("\nData Sebelum di Urutkan\n")
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
						InsertionSortDescendingJumlahDana(&dataInvestasi, idxInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
					case 4:
						InitPersentaseDanKeuntungan(&dataInvestasi, idxInvestasi)
						fmt.Print("\nData Sebelum di Urutkan\n")
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
						InsertionSortDescendingKeuntungan(&dataInvestasi, idxInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
					case 5:
						InitPersentaseDanKeuntungan(&dataInvestasi, idxInvestasi)
						fmt.Print("\nData Sebelum di Urutkan\n")
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
						InsertionSortDescendingPersentaseKeuntungan(&dataInvestasi, idxInvestasi)
						fmt.Print("\nData Setelah di Urutkan\n")
						CetakDataInvestasi(&dataInvestasi, idxInvestasi)
					}
				} else {
					fmt.Println("Pilihan tidak valid! Silahkan pilih 1-2.")
				}
			case 6:
				CetakLaporanInvestasi(&dataInvestasi, idxInvestasi)
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
func InitDummyUsers(TUsers *tabUser) {
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
func InitDummytabInvestasi(TInvestasi *tabInvestasi) {
	// IS. Terdefinisi array TInvestasi dalam keadaan kosong tanpa data aset investasi apapun.
	// FS. Mengisi array TInvestasi dengan beberapa data dummy aset investasi yang sudah ditentukan
	//     pada indeks awal array untuk tujuan pengujian atau inisialisasi awal program.
	TInvestasi[0] = AsetInvestasi{
		NamaAset:            "ADRO",
		JenisAset:           "Saham",
		NilaiAset:           3000000,
		JumlahDanaInvestasi: 90000000,
		HargaBeli:           3000000,
		HargaJual:           3200000,
	}
	TInvestasi[1] = AsetInvestasi{
		NamaAset:            "BBRI",
		JenisAset:           "Saham",
		NilaiAset:           1500000,
		JumlahDanaInvestasi: 50000000,
		HargaBeli:           1500000,
		HargaJual:           1800000,
	}
	TInvestasi[2] = AsetInvestasi{
		NamaAset:            "TLKM",
		JenisAset:           "Saham",
		NilaiAset:           2000000,
		JumlahDanaInvestasi: 75000000,
		HargaBeli:           2000000,
		HargaJual:           2400000,
	}
	TInvestasi[3] = AsetInvestasi{
		NamaAset:            "Mandiri-Investa-Pasar-Uang",
		JenisAset:           "Reksadana",
		NilaiAset:           100000,
		JumlahDanaInvestasi: 30000000,
		HargaBeli:           100000,
		HargaJual:           103000,
	}
	TInvestasi[4] = AsetInvestasi{
		NamaAset:            "Schroder-Dana-Likuid",
		JenisAset:           "Reksadana",
		NilaiAset:           125000,
		JumlahDanaInvestasi: 40000000,
		HargaBeli:           125000,
		HargaJual:           127500,
	}
	TInvestasi[5] = AsetInvestasi{
		NamaAset:            "ORI022",
		JenisAset:           "Obligasi",
		NilaiAset:           1000000,
		JumlahDanaInvestasi: 100000000,
		HargaBeli:           1000000,
		HargaJual:           1050000,
	}
	TInvestasi[6] = AsetInvestasi{
		NamaAset:            "SR018",
		JenisAset:           "Obligasi",
		NilaiAset:           1000000,
		JumlahDanaInvestasi: 85000000,
		HargaBeli:           1000000,
		HargaJual:           1080000,
	}
}

// funciton dan procedure untuk user management
func ValidasiUsers(TUser *tabUser, inputUser Users, n int) bool {
	// function yang mengembalikan true jika  email dan password yang diberikan  ke parameter inputUser ada di dalam array tabuser
	for i := 0; i < n; i++ {
		if TUser[i].Email == inputUser.Email && TUser[i].Password == inputUser.Password {
			return true
		}
	}
	return false
}

func RegisterUser(TUsers *tabUser, n *int) {
	/*IS. Terdefinisi array TUsers yang berukuran maksimum NMAX dan sudah terisi sebanyak idxUser.
	     Jika idxUser sudah mencapai atau melebihi NMAX, maka tidak ada user baru yang dapat didaftarkan.
	 FS. Array TUsers bertamabah, yang isinya data user ketika idxUser tidak melebihi NMAX dan
		memunculkan pesan gagal ketika idxUser sudah.*/
	if *n >= NMAX {
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
	TUsers[*n] = Users{
		Nama:     name,
		Email:    email,
		Password: password,
	}
	*n++
	fmt.Println("\n=========================================================")
	fmt.Println("    Akun Berhasil Dibuat           ")
	fmt.Println("    Silahkan Login!                ")
	fmt.Println("=========================================================")
}

func InputLogin(TUser tabUser, n int) bool {
	// function yang mengembalikan true jika user dengan email dan password yang benar
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

	// membuat struct userInput untuk menampung input dari user agar mudah dimasukan ke validasi dan struct dataUsers
	userInput := Users{
		Email:    email,
		Password: password,
	}
	// menyimpan hasil dari function ValidasiUsers
	valid := ValidasiUsers(&TUser, userInput, n)

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
			usersInput := Users{
				Email:    email,
				Password: password,
			}
			valid = ValidasiUsers(&TUser, usersInput, n)
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

// function dan procedure untuk management investasi
func CreateDataInvesatasi(TInvestasi *tabInvestasi, n *int) {
	// IS. Terdefinisi suatu array dataInvestasi yang berukuran maksimum NMAX dan telah terisi sebanyak n elemen.
	//      Diberikan input namaAset, jenisAset, nilaiAset, dan danaInvestasi yang akan dimasukkan ke dalam array.
	// FS. Menambahkan data aset investasi ke dalam array dataInvestasi apabila belum penuh (n < NMAX),
	//      atau menampilkan pesan "Data Investasi Sudah Penuh" apabila array telah mencapai kapasitas maksimum.

	if *n > NMAX {
		fmt.Println("\n=========================================================")
		fmt.Println("    Gagal! Data Investasi Sudah Penuh       ")
		fmt.Println("=========================================================")
	} else {
		var namaAset, jenisAset string
		var nilaiAset, jumlahDanaInvestasi float64
		var hargaBeli, hargaJual float64
		fmt.Println("\n=========================================================")
		fmt.Println("               Tambah Data Investasi")
		fmt.Println("=========================================================")
		fmt.Print("Masukkan Nama Aset: ")
		fmt.Scan(&namaAset)
		fmt.Print("Masukkan Jenis Aset: ")
		fmt.Scan(&jenisAset)
		fmt.Print("Masukkan Nilai Aset: ")
		fmt.Scan(&nilaiAset)
		fmt.Print("Masukkan Jumlah Dana Investasi: ")
		fmt.Scan(&jumlahDanaInvestasi)
		fmt.Print("Masukkan Harga Beli: ")
		fmt.Scan(&hargaBeli)
		fmt.Print("Masukkan Harga Jual: ")
		fmt.Scan(&hargaJual)
		TInvestasi[*n] = AsetInvestasi{
			NamaAset:            namaAset,
			JenisAset:           jenisAset,
			NilaiAset:           nilaiAset,
			JumlahDanaInvestasi: jumlahDanaInvestasi,
			HargaBeli:           hargaBeli,
			HargaJual:           hargaJual,
		}
		*n++
	}
}
func ModifyDataInvestasi(TInvestasi *tabInvestasi, n *int, namaAset string) {
	// IS. Terdefinisi sebuah array dataInvestasi yang berisi data aset investasi sebanyak n elemen.
	//     Parameter namaAset berisi string yang akan digunakan untuk mencari data aset yang ingin dimodifikasi.
	// FS. Jika data dengan NamaAset ditemukan dalam dataInvestasi, maka data tersebut dimodifikasi dengan nilai-nilai baru
	//     yang dimasukkan oleh pengguna (nama aset, jenis aset, jumlah dana, nilai awal, nilai update).
	//     Jika tidak ditemukan, maka ditampilkan pesan bahwa data tidak ditemukan.
	var dataFound bool = false
	for i := 0; i < *n; i++ {
		if TInvestasi[i].NamaAset == namaAset {
			dataFound = true
			fmt.Println("=========================================================")
			fmt.Println("               Modifikasi Data Investasi")
			fmt.Println("=========================================================")
			fmt.Println("Data Sebelumnya:")
			fmt.Printf("Nama Aset: %s\n", TInvestasi[i].NamaAset)
			fmt.Printf("Jenis Aset: %s\n", TInvestasi[i].JenisAset)
			fmt.Printf("Jumlah Dana: %.2f\n", TInvestasi[i].JumlahDanaInvestasi)
			fmt.Printf("Nilai Aset: %.2f\n", TInvestasi[i].NilaiAset)
			fmt.Printf("Harga Beli: %.2f\n", TInvestasi[i].HargaBeli)
			fmt.Printf("Harga Jual: %.2f\n", TInvestasi[i].HargaJual)
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
			TInvestasi[i].NamaAset = namaAsetBaru
			TInvestasi[i].JenisAset = jenisAsetBaru
			TInvestasi[i].JumlahDanaInvestasi = jumlahDanaBaru
			TInvestasi[i].NilaiAset = nilaiAsetBaru
			TInvestasi[i].HargaBeli = hargaBeliBaru
			TInvestasi[i].HargaJual = hargaJualBaru

			fmt.Println("---------------------------------------------------------")
			fmt.Println("Data investasi berhasil dimodifikasi.")
			fmt.Println("=========================================================")
		}
	}

	if !dataFound {
		fmt.Println("Data investasi tidak ditemukan, Pastikan Nama Aset benar.")
	}
}

func HapusDataInvestasi(TInvestasi *tabInvestasi, n *int, namaAset string) {
	// IS: Terdefinisi array dataInvestasi berisi data aset investasi sebanyak n.
	//     Parameter namaAset berisi nama aset yang ingin dihapus.
	// FS: Jika aset dengan nama yang sesuai ditemukan, maka data aset tersebut dihapus dari array dataInvestasi,
	//     elemen setelahnya digeser ke kiri, dan nilai n dikurangi 1.
	//     Jika tidak ditemukan, maka akan ditampilkan pesan bahwa aset tidak ditemukan.
	var foundIndex int = -1
	for i := 0; i < *n; i++ {
		if TInvestasi[i].NamaAset == namaAset && foundIndex == -1 {
			foundIndex = i
		}
	}
	dataAset := TInvestasi[foundIndex]
	if foundIndex != -1 {
		for j := foundIndex; j < *n-1; j++ {
			TInvestasi[j] = TInvestasi[j+1]
		}
		*n--
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
func FindDataByName(TInvestasi tabInvestasi, n int, nameAset string) int {
	// {diberikan array dataInvestasi yang berisi n data aset dan sebuah nama aset,
	//  untuk mengembalikan indeks data apabila nama aset ditemukan pada array,
	//  serta menampilkan detail data aset tersebut; jika tidak ditemukan, mengembalikan -1}
	found := 0
	for i := 0; i < n; i++ {
		if TInvestasi[i].NamaAset == nameAset {
			fmt.Println("===========================")
			fmt.Printf("Data Ditemukan di Index: %d\n", i)
			fmt.Println("===========================")
			fmt.Printf("Nama Aset      : %s\n", TInvestasi[i].NamaAset)
			fmt.Printf("Jenis Aset     : %s\n", TInvestasi[i].JenisAset)
			fmt.Printf("Nilai Aset     : %.2f\n", TInvestasi[i].NilaiAset)
			fmt.Printf("Jumlah Dana    : %.2f\n", TInvestasi[i].JumlahDanaInvestasi)
			fmt.Printf("Harga Beli    : %.2f\n", TInvestasi[i].HargaBeli)
			fmt.Printf("Harga Jual    : %.2f\n", TInvestasi[i].HargaJual)
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
func FindDataByJenis(TInvestasi tabInvestasi, n int, jenisAset string) int {
	// {diberikan array dataInvestasi yang berisi n data aset dan sebuah jenis aset,
	//  untuk mengembalikan indeks data apabila jenis aset ditemukan pada array,
	//  serta menampilkan detail data aset tersebut; jika tidak ditemukan, mengembalikan -1}
	found := 0
	for i := 0; i < n; i++ {
		if TInvestasi[i].JenisAset == jenisAset {
			fmt.Println("===========================")
			fmt.Printf("Data Ditemukan di Index: %d\n", i)
			fmt.Println("===========================")
			fmt.Printf("Nama Aset      : %s\n", TInvestasi[i].NamaAset)
			fmt.Printf("Jenis Aset     : %s\n", TInvestasi[i].JenisAset)
			fmt.Printf("Nilai Aset     : %.2f\n", TInvestasi[i].NilaiAset)
			fmt.Printf("Jumlah Dana    : %.2f\n", TInvestasi[i].JumlahDanaInvestasi)
			fmt.Printf("Harga Beli    : %.2f\n", TInvestasi[i].HargaBeli)
			fmt.Printf("Harga Jual    : %.2f\n", TInvestasi[i].HargaJual)
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
func FindByJumlahDana(TInvestasi tabInvestasi, n int, jumlahDana float64) int {
	left := 0
	right := n - 1
	for left <= right {
		mid := (left + right) / 2
		midValue := TInvestasi[mid].JumlahDanaInvestasi

		if midValue == jumlahDana {
			// Data ditemukan
			fmt.Println("===========================")
			fmt.Printf("Data Ditemukan di Index: %d\n", mid)
			fmt.Println("===========================")
			fmt.Printf("Nama Aset      : %s\n", TInvestasi[mid].NamaAset)
			fmt.Printf("Jenis Aset     : %s\n", TInvestasi[mid].JenisAset)
			fmt.Printf("Nilai Aset     : %.2f\n", TInvestasi[mid].NilaiAset)
			fmt.Printf("Jumlah Dana    : %.2f\n", TInvestasi[mid].JumlahDanaInvestasi)
			fmt.Printf("Harga Beli    : %.2f\n", TInvestasi[mid].HargaBeli)
			fmt.Printf("Harga Jual    : %.2f\n", TInvestasi[mid].HargaJual)
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

func SelectionSortAscendingNamaAset(TInvestasi *tabInvestasi, n int) {
	var i, idx, pass int
	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if TInvestasi[i].NamaAset < TInvestasi[idx].NamaAset {
				idx = i
			}
			i++
		}
		temp := TInvestasi[pass-1]
		TInvestasi[pass-1] = TInvestasi[idx]
		TInvestasi[idx] = temp
		pass++
	}
}
func SelectionSortDescendingNamaAset(TInvestasi *tabInvestasi, n int) {
	var i, idx, pass int
	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if TInvestasi[i].NamaAset > TInvestasi[idx].NamaAset {
				idx = i
			}
			i++
		}
		temp := TInvestasi[pass-1]
		TInvestasi[pass-1] = TInvestasi[idx]
		TInvestasi[idx] = temp
		pass++
	}
}
func SelectionSortAscendingJenisAset(TInvestasi *tabInvestasi, n int) {
	var i, idx, pass int
	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if TInvestasi[i].JenisAset < TInvestasi[idx].JenisAset {
				idx = i
			}
			i++
		}
		temp := TInvestasi[pass-1]
		TInvestasi[pass-1] = TInvestasi[idx]
		TInvestasi[idx] = temp
		pass++
	}
}
func SelectionSortDescendingJenisAset(TInvestasi *tabInvestasi, n int) {
	var i, idx, pass int
	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if TInvestasi[i].JenisAset > TInvestasi[idx].JenisAset {
				idx = i
			}
			i++
		}
		temp := TInvestasi[pass-1]
		TInvestasi[pass-1] = TInvestasi[idx]
		TInvestasi[idx] = temp
		pass++
	}
}
func InsertionSortAscendingJumlahDana(TInvestasi *tabInvestasi, n int) {
	var pass, i int
	var temp AsetInvestasi
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = TInvestasi[pass]

		for i > 0 && temp.JumlahDanaInvestasi < TInvestasi[i-1].JumlahDanaInvestasi {
			TInvestasi[i] = TInvestasi[i-1]
			i = i - 1
		}

		TInvestasi[i] = temp
		pass = pass + 1
	}
}
func InsertionSortDescendingJumlahDana(TInvestasi *tabInvestasi, n int) {
	var pass, i int
	var temp AsetInvestasi
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = TInvestasi[pass]

		for i > 0 && temp.JumlahDanaInvestasi > TInvestasi[i-1].JumlahDanaInvestasi {
			TInvestasi[i] = TInvestasi[i-1]
			i = i - 1
		}

		TInvestasi[i] = temp
		pass = pass + 1
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
func InitPersentaseDanKeuntungan(TInvestasi *tabInvestasi, n int) {
	for i := 0; i < n; i++ {
		TInvestasi[i].Keuntungan = HitungKeuntungan(TInvestasi[i])
		TInvestasi[i].PersentaseKeuntungan = HitungPersentaseKeuntungan(TInvestasi[i])
	}
}
func InsertionSortAscendingKeuntungan(TInvestasi *tabInvestasi, n int) {
	var pass, i int
	var temp AsetInvestasi
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = TInvestasi[pass]

		for i > 0 && temp.Keuntungan < TInvestasi[i-1].Keuntungan {
			TInvestasi[i] = TInvestasi[i-1]
			i = i - 1
		}

		TInvestasi[i] = temp
		pass = pass + 1
	}
}
func InsertionSortDescendingKeuntungan(TInvestasi *tabInvestasi, n int) {
	var pass, i int
	var temp AsetInvestasi
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = TInvestasi[pass]

		for i > 0 && temp.Keuntungan > TInvestasi[i-1].Keuntungan {
			TInvestasi[i] = TInvestasi[i-1]
			i = i - 1
		}

		TInvestasi[i] = temp
		pass = pass + 1
	}
}
func InsertionSortAscendingPersentaseKeuntungan(TInvestasi *tabInvestasi, n int) {
	var pass, i int
	var temp AsetInvestasi
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = TInvestasi[pass]

		for i > 0 && temp.PersentaseKeuntungan < TInvestasi[i-1].PersentaseKeuntungan {
			TInvestasi[i] = TInvestasi[i-1]
			i = i - 1
		}

		TInvestasi[i] = temp
		pass = pass + 1
	}
}
func InsertionSortDescendingPersentaseKeuntungan(TInvestasi *tabInvestasi, n int) {
	var pass, i int
	var temp AsetInvestasi
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = TInvestasi[pass]

		for i > 0 && temp.PersentaseKeuntungan > TInvestasi[i-1].PersentaseKeuntungan {
			TInvestasi[i] = TInvestasi[i-1]
			i = i - 1
		}

		TInvestasi[i] = temp
		pass = pass + 1
	}
}

func CetakDataInvestasi(TInvestasi *tabInvestasi, n int) {
	fmt.Println("=================================================================================================================================")
	fmt.Printf("| %-3s | %-30s | %-10s | %-15s | %-12s | %-12s | %-12s | %-10s |\n",
		"No", "Nama Aset", "Jenis", "Jumlah Dana", "Harga Beli", "Harga Jual", "Keuntungan", "% Untung")
	fmt.Println("=================================================================================================================================")
	for i := 0; i < n; i++ {
		fmt.Printf("| %-3d | %-30s | %-10s | %-15.2f | %-12.2f | %-12.2f | %-12.2f | %-10.2f |\n",
			i+1,
			TInvestasi[i].NamaAset,
			TInvestasi[i].JenisAset,
			TInvestasi[i].JumlahDanaInvestasi,
			TInvestasi[i].HargaBeli,
			TInvestasi[i].HargaJual,
			TInvestasi[i].Keuntungan,
			TInvestasi[i].PersentaseKeuntungan,
		)
	}
	fmt.Println("=================================================================================================================================")
}

func CetakLaporanInvestasi(TInvestasi *tabInvestasi, n int) {
	// IS: Terdefinisi array TInvestasi yang berisi data aset investasi sebanyak n
	// FS: Menampilkan laporan portofolio investasi yang meliputi:
	//     - Total investasi, total keuntungan, dan rata-rata keuntungan
	//     - Aset paling untung dan paling rugi
	//     - Komposisi portofolio berdasarkan jenis aset
	if n == 0 {
		fmt.Println("=========================================================")
		fmt.Println("              TIDAK ADA DATA INVESTASI                  ")
		fmt.Println("=========================================================")
	}
	// Inisialisasi perhitungan keuntungan dan persentase
	InitPersentaseDanKeuntungan(TInvestasi, n)
	// Variabel untuk perhitungan
	var totalInvestasi float64 = 0
	var totalKeuntungan float64 = 0
	var asetPalingUntung, asetPalingRugi AsetInvestasi
	var maxKeuntungan float64 = TInvestasi[0].Keuntungan
	var minKeuntungan float64 = TInvestasi[0].Keuntungan

	// Variabel untuk komposisi portofolio
	var totalSaham, totalReksadana, totalObligasi float64 = 0, 0, 0

	// Perhitungan total dan pencarian aset terbaik/terburuk
	for i := 0; i < n; i++ {
		aset := TInvestasi[i]
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

// user interfaces
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
func GoodBye() {
	fmt.Println("=========================================================")
	fmt.Println("               TERIMA KASIH TELAH MENGGUNAKAN            ")
	fmt.Println("                	APLIKASI INI  :)           ")
	fmt.Println("=========================================================")
	fmt.Println("                      Sampai Jumpa!                      ")
	fmt.Println("=========================================================")
}
