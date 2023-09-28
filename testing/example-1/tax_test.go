package tax

import "testing"

// func TestCalculateTax(t *testing.T) {
// 	amount := 1000.0
// 	expected := 5.0

// 	result := CalculateTax(amount)

// 	if result != expected {
// 		t.Errorf("Expected %f but got %f", expected, result)
// 	}
// }

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct { amount, expected float64 }
	table := []calcTax{
		{ 0.0, 0.0 },
		{ 100.0, 5.0 },
		{ 999.9, 5.0 },
		{ 1000.1, 10.0 },
	}

	for _, item := range table {
		result := CalculateTax(item.amount)

		if result != item.expected {
			t.Errorf("Expected %f but got %f", item.expected, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.1)
	}
}