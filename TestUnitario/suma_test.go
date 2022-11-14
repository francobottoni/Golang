package TestUnitario

import "testing"

/*func TestSuma(t *testing.T){
	total := Suma(3,3)

	if total != 6{
		t.Errorf("Suma incorrecta, resultado: %d",total)
	}
}*/

func TestSuma(t *testing.T){
	tabla := []struct{
		a int
		b int
		c int
	}{
		{1,2,3},
		{2,2,4},
		{25,25,50},
	}

	for _,i := range tabla{
		total := Suma(i.a,i.b)

		if total != i.c {
			t.Errorf("Suma incorrecta, resultado: %d y se esperaba: %d", total,i.c)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tabla := []struct{
		n int
		r int
	}{
		{1,1},
		{8,21},
		{50,12586269025},
	}

	for _,i := range tabla{
		total := Fibonacci(i.n)

		if total != i.r {
			t.Errorf("Fibonacci incorrecto, resultado: %d",i.r)
		}
	}
}