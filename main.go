package main

import (
	"fmt"
	"math"
)

// Function to implement Simplex Method
func SimplexMethod(C []float64, A [][]float64, b []float64, eps float64) (string, []float64, float64) {
	numVars := len(C)
	numConstraints := len(A)
	numColumns := numVars + numConstraints + 1 // variables + slack variables + RHS

	// Initialize tableau
	tableau := make([][]float64, numConstraints+1)
	for i := 0; i < numConstraints; i++ {
		tableau[i] = make([]float64, numColumns)
		// Coefficients of variables
		for j := 0; j < numVars; j++ {
			tableau[i][j] = A[i][j]
		}
		// Slack variables
		tableau[i][numVars+i] = 1.0
		// RHS
		tableau[i][numColumns-1] = b[i]
	}

	// Objective function row
	tableau[numConstraints] = make([]float64, numColumns)
	for j := 0; j < numVars; j++ {
		tableau[numConstraints][j] = -C[j]
	}

	// Initialize basic variables (slack variables)
	basicVars := make([]int, numConstraints)
	for i := 0; i < numConstraints; i++ {
		basicVars[i] = numVars + i // indices of slack variables
	}

	// Simplex iterations
	for {
		// Find entering variable (most negative coefficient in objective function row)
		enteringVar := -1
		minCoeff := -eps
		for j := 0; j < numColumns-1; j++ {
			coeff := tableau[numConstraints][j]
			if coeff < minCoeff {
				minCoeff = coeff
				enteringVar = j
			}
		}

		// If no negative coefficients, optimal solution found
		if enteringVar == -1 {
			break
		}

		// Find leaving variable (minimum ratio test)
		leavingVar := -1
		minRatio := math.MaxFloat64
		for i := 0; i < numConstraints; i++ {
			aij := tableau[i][enteringVar]
			if aij > eps {
				ratio := tableau[i][numColumns-1] / aij
				if ratio < minRatio {
					minRatio = ratio
					leavingVar = i
				}
			}
		}

		// If no leaving variable, unbounded solution
		if leavingVar == -1 {
			return "The method is not applicable!", nil, 0
		}

		// Pivot operation
		pivot := tableau[leavingVar][enteringVar]
		// Normalize the pivot row
		for j := 0; j < numColumns; j++ {
			tableau[leavingVar][j] /= pivot
		}

		// Update other rows
		for i := 0; i <= numConstraints; i++ {
			if i != leavingVar {
				factor := tableau[i][enteringVar]
				for j := 0; j < numColumns; j++ {
					tableau[i][j] -= factor * tableau[leavingVar][j]
				}
			}
		}

		// Update basic variables
		basicVars[leavingVar] = enteringVar
	}

	// Extract solution
	x := make([]float64, numVars)
	for i := 0; i < numConstraints; i++ {
		basicVar := basicVars[i]
		if basicVar < numVars {
			x[basicVar] = tableau[i][numColumns-1]
		}
	}

	// Objective function value
	z := tableau[numConstraints][numColumns-1]

	return "solved", x, z
}

func main() {
	var n, m int
	fmt.Print("Enter the number of variables: ")
	fmt.Scan(&n)
	fmt.Print("Enter the number of constraints: ")
	fmt.Scan(&m)

	C := make([]float64, n)
	fmt.Println("Enter the coefficients of the objective function (C):")
	for i := 0; i < n; i++ {
		fmt.Printf("C[%d]: ", i)
		fmt.Scan(&C[i])
	}

	A := make([][]float64, m)
	fmt.Println("Enter the coefficients of the constraint functions (A):")
	for i := 0; i < m; i++ {
		A[i] = make([]float64, n)
		fmt.Printf("Constraint %d:\n", i+1)
		for j := 0; j < n; j++ {
			fmt.Printf("A[%d][%d]: ", i, j)
			fmt.Scan(&A[i][j])
		}
	}

	b := make([]float64, m)
	fmt.Println("Enter the right-hand side numbers (b):")
	for i := 0; i < m; i++ {
		fmt.Printf("b[%d]: ", i)
		fmt.Scan(&b[i])
	}

	var eps float64
	fmt.Print("Enter the approximation accuracy (epsilon): ")
	fmt.Scan(&eps)

	status, solution, objective := SimplexMethod(C, A, b, eps)
	if status == "solved" {
		fmt.Println("Optimal solution found:")
		for i := 0; i < len(solution); i++ {
			fmt.Printf("x%d = %.6f\n", i+1, solution[i])
		}
		fmt.Printf("Objective value: %.6f\n", objective)
	} else {
		fmt.Println(status)
	}
}
