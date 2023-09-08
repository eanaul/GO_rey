package main

import("fmt")

func main() {
	var ujian = 90
	var absen = 78

	var lulusUjian = ujian > 80
	var lulusAbsensi = absen > 80

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println(lulus)

}
