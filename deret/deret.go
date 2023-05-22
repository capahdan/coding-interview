package deret

func Deret(nilai1, nilai2, deret int) []int {

	var data = []int{}
	data = append(data, nilai1)
	data = append(data, nilai2)
	for i := 3; i <= deret; i++ {
		deret := nilai1*(i-1) + nilai1*(i-2)
		data = append(data, deret)
	}
	return data
}
