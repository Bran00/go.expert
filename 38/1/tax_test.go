package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{2000.0, 10.0},
		{400.0, 5.0},
		{0.0, 0.0},
		{-100.0, 0.0},
	}

	for _, test := range table {
		result := CalculateTax(test.amount)
		if result != test.expected {
			t.Errorf("Expected %f, but got %f for amount %f", test.expected, result, test.amount)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, 2.5, 500.0, 1000.0, 2000.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Recevied %f, but expected 0 for amount %f", result, amount)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Recevied %f, but expected 20.0 for amount %f", result, amount)
		}
	})
}