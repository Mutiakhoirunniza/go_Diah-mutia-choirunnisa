package calculate

import "testing"

func TestAddition(t *testing.T) {
	result := Addition(10, 5)
	if result != 15 {
		t.Errorf("Expected : 15, didapatkan %f", result)
	}
}

func TestSubtraction(t *testing.T) { // Perbaiki nama fungsi menjadi "TestSubtract"
	result := Subtraction(10, 5)
	if result != 5 {
		t.Errorf("Expected : 5, didapatkan %f", result)
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(10, 5)
	if result != 50 {
		t.Errorf("Expected : 50, didapatkan %f", result)
	}
}

func TestDivision(t *testing.T) {
	result, err := Division(10, 5)
	if err != nil || result != 2 {
		t.Errorf("Expected : 2 tanpa error, but got %.2f dengan error: %v", float64(result), err)
	}
	_, err = Division(10, 0)
	if err == nil {
		t.Errorf("error saat pembagian oleh bilangan  0, tetapi tidak ada error yang ditemukan")
	}
}
