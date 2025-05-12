// divmod.go
package piscine

func DivMod(a int, b int, div *int, mod *int) {
	if b != 0 {
		*div = a / b
		*mod = a % b
	} else {
		*div = 0
		*mod = 0
	}
}
