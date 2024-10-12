package main

import (
	"math"
	"testing"
)

func TestSimplexMethod1(t *testing.T) {
	C := []float64{3, 2}
	A := [][]float64{
		{-2, 1},
	}
	b := []float64{2}
	eps := 1e-6

	status, _, _ := SimplexMethod(C, A, b, eps)
	if status != "The method is not applicable!" {
		t.Errorf("Expected status 'The method is not applicable!', got '%s'", status)
	}
}


func TestSimplexMethod2(t *testing.T) {
	C := []float64{2, 3}
	A := [][]float64{
		{1, 1},
		{2, 1},
	}
	b := []float64{4, 5}
	eps := 1e-6

	status, solution, objective := SimplexMethod(C, A, b, eps)
	if status != "solved" {
		t.Errorf("Expected status 'solved', got '%s'", status)
	}

	expectedSolution := []float64{0, 4}
	expectedObjective := 12.0

	for i := range solution {
		if math.Abs(solution[i]-expectedSolution[i]) > eps {
			t.Errorf("Expected solution x%d = %.6f, got %.6f", i+1, expectedSolution[i], solution[i])
		}
	}

	if math.Abs(objective-expectedObjective) > eps {
		t.Errorf("Expected objective value %.6f, got %.6f", expectedObjective, objective)
	}
}

func TestSimplexMethod3(t *testing.T) {
	C := []float64{4, 3}
	A := [][]float64{
		{2, 1},
		{1, 1},
	}
	b := []float64{8, 6}
	eps := 1e-6

	status, solution, objective := SimplexMethod(C, A, b, eps)
	if status != "solved" {
		t.Errorf("Expected status 'solved', got '%s'", status)
	}

	expectedSolution := []float64{2, 4}
	expectedObjective := 20.0

	for i := range solution {
		if math.Abs(solution[i]-expectedSolution[i]) > eps {
			t.Errorf("Expected solution x%d = %.6f, got %.6f", i+1, expectedSolution[i], solution[i])
		}
	}

	if math.Abs(objective-expectedObjective) > eps {
		t.Errorf("Expected objective value %.6f, got %.6f", expectedObjective, objective)
	}
}

func TestSimplexMethod4(t *testing.T) {
	C := []float64{5, 7}
	A := [][]float64{
		{1, 1},
		{3, 2},
	}
	b := []float64{10, 24}
	eps := 1e-6

	status, solution, objective := SimplexMethod(C, A, b, eps)
	if status != "solved" {
		t.Errorf("Expected status 'solved', got '%s'", status)
	}

	expectedSolution := []float64{0, 10}
	expectedObjective := 70.0

	for i := range solution {
		if math.Abs(solution[i]-expectedSolution[i]) > eps {
			t.Errorf("Expected solution x%d = %.6f, got %.6f", i+1, expectedSolution[i], solution[i])
		}
	}

	if math.Abs(objective-expectedObjective) > eps {
		t.Errorf("Expected objective value %.6f, got %.6f", expectedObjective, objective)
	}
}

func TestSimplexMethod5(t *testing.T) {
	C := []float64{1, 1}
	A := [][]float64{
		{1, 1},
	}
	b := []float64{5}
	eps := 1e-6

	status, solution, objective := SimplexMethod(C, A, b, eps)
	if status != "solved" {
		t.Errorf("Expected status 'solved', got '%s'", status)
	}

	expectedObjective := 5.0
	if math.Abs(objective-expectedObjective) > eps {
		t.Errorf("Expected objective value %.6f, got %.6f", expectedObjective, objective)
	}

	if math.Abs(solution[0]+solution[1]-5) > eps {
		t.Errorf("Expected x1 + x2 = 5, got x1 + x2 = %.6f", solution[0]+solution[1])
	}

	if solution[0] < -eps || solution[1] < -eps {
		t.Errorf("Variables should be non-negative, got x1 = %.6f, x2 = %.6f", solution[0], solution[1])
	}
}
