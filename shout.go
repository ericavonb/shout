package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func getLetter(str string) []string {

	A := make([]string, 8)
	A[0] = `   __      `
	A[1] = `  /\ _\    `
	A[2] = ` / /   \   `
	A[3] = `|\/ /_\ \  `
	A[4] = `| | |_\| | `
	A[5] = `| |  __  | `
	A[6] = ` \|_| \|_| `
	A[7] = `            `

	B := make([]string, 8)
	B[0] = ` ______    `
	B[1] = `|\ _____\  `
	B[2] = `| |  __  | `
	B[3] = `| | |_\| / `
	B[4] = `| |  __  \ `
	B[5] = `| | |_\| | `
	B[6] = ` \|______/ `
	B[7] = `           `

	C := make([]string, 8)
	C[0] = ` ______    `
	C[1] = `|\ _____\  `
	C[2] = `| |  ____| `
	C[3] = `| | |      `
	C[4] = `| | |___   `
	C[5] = `| | |____\ `
	C[6] = ` \|______| `
	C[7] = `           `

	D := make([]string, 8)
	D[0] = ` ______    `
	D[1] = `|\ _____\  `
	D[2] = `| |  __  \ `
	D[3] = `| | || | | `
	D[4] = `| | || | | `
	D[5] = `| | |_\| | `
	D[6] = ` \|______/ `
	D[7] = `           `

	E := make([]string, 8)
	E[0] = ` ______    `
	E[1] = `|\ _____\  `
	E[2] = `| |  ____| `
	E[3] = `| | |___\  `
	E[4] = `| |  ____| `
	E[5] = `| | |___\  `
	E[6] = ` \|______| `
	E[7] = `           `

	F := make([]string, 8)
	F[0] = ` ______    `
	F[1] = `|\ _____\  `
	F[2] = `| |  ____| `
	F[3] = `| | |___\  `
	F[4] = `| |  ____| `
	F[5] = `| | |      `
	F[6] = ` \|_|      `
	F[7] = `           `

	G := make([]string, 8)
	G[0] = ` ______    `
	G[1] = `|\ _____\  `
	G[2] = `| |  ____| `
	G[3] = `| | |  __  `
	G[4] = `| | |_|\_\ `
	G[5] = `| | |__\| |`
	G[6] = ` \|______/ `
	G[7] = `           `

	H := make([]string, 8)
	H[0] = ` __    __  `
	H[1] = `|\ _\ |\ _\`
	H[2] = `| | |_| | |`
	H[3] = `| | |__\| |`
	H[4] = `| |  ___  |`
	H[5] = `| | |   | |`
	H[6] = ` \|_|  \|_|`
	H[7] = `           `

	I := make([]string, 8)
	I[0] = ` _______   `
	I[1] = `|\ ______\ `
	I[2] = ` \|__   __|`
	I[3] = `   | | |   `
	I[4] = ` __| | |_  `
	I[5] = `|\ _\| |_\ `
	I[6] = ` \|_______|`
	I[7] = `           `

	J := make([]string, 8)
	J[0] = `       __  `
	J[1] = `      |\_\ `
	J[2] = `      | | |`
	J[3] = ` _    | | |`
	J[4] = `|\_\ _| | |`
	J[5] = `| | |__\| |`
	J[6] = ` \|______/ `
	J[7] = `           `

	K := make([]string, 8)
	K[0] = ` __    __  `
	K[1] = `|\ \  /\__\`
	K[2] = `| | |//  / `
	K[3] = `| | |/  /  `
	K[4] = `| |  _  \  `
	K[5] = `| | |\\  \ `
	K[6] = ` \|_| \\__\`
	K[7] = `           `

	L := make([]string, 8)
	L[0] = ` __        `
	L[1] = `|\ _\      `
	L[2] = `| | |      `
	L[3] = `| | |      `
	L[4] = `| | |___   `
	L[5] = `| | |____\ `
	L[6] = ` \|______| `
	L[7] = `           `

	M := make([]string, 8)
	M[0] = ` __    __  `
	M[1] = `|\ _\/\ _\ `
	M[2] = `| |  \/   |`
	M[3] = `| | |\_/| |`
	M[4] = `| | | | | |`
	M[5] = `| | | | | |`
	M[6] = ` \|_|  \|_|`
	M[7] = `           `

	N := make([]string, 8)
	N[0] = ` __    __  `
	N[1] = `|\ _\ |\ \ `
	N[2] = `| |  \| | |`
	N[3] = `| |   \ | |`
	N[4] = `| | |\ \| |`
	N[5] = `| | | \   |`
	N[6] = ` \|_|  \__|`
	N[7] = `           `

	O := make([]string, 8)
	O[0] = `   ____    `
	O[1] = ` /\ ____\  `
	O[2] = `| /  ___  \`
	O[3] = `| | | | | |`
	O[4] = `| | |_| | |`
	O[5] = `\ | |__\| |`
	O[6] = `  \______/ `
	O[7] = `           `

	P := make([]string, 8)
	P[0] = ` ______    `
	P[1] = `|\ _____\  `
	P[2] = `| |  __  | `
	P[3] = `| | |_\| | `
	P[4] = `| |  ____| `
	P[5] = `| | |      `
	P[6] = ` \|_|      `
	P[7] = `           `

	Q := make([]string, 8)
	Q[0] = `   ____    `
	Q[1] = ` /\ ____\  `
	Q[2] = `| /  ___ \`
	Q[3] = `| | || |  |`
	Q[4] = `| | |/\|  |`
	Q[5] = `\ | |/   |`
	Q[6] = `  \___/\__|`
	Q[7] = `           `

	R := make([]string, 8)
	R[0] = ` ______    `
	R[1] = `|\ _____\  `
	R[2] = `| |  __  | `
	R[3] = `| | |_\| | `
	R[4] = `| |  _  _| `
	R[5] = `| | |\ \ \ `
	R[6] = ` \|_|  \\_|`
	R[7] = `           `

	S := make([]string, 8)
	S[0] = ` _______  `
	S[1] = `|\ ______\ `
	S[2] = `| |  _____|`
	S[3] = `| | |____\ `
	S[4] = ` \|_____  |`
	S[5] = `|\ ____ | |`
	S[6] = ` \|_______|`
	S[7] = `           `

	T := make([]string, 8)
	T[0] = ` _______   `
	T[1] = `|\ ______\ `
	T[2] = ` \|__   __|`
	T[3] = `   | | |   `
	T[4] = `   | | |   `
	T[5] = `   | | |   `
	T[6] = `    \|_|   `
	T[7] = `           `

	U := make([]string, 8)
	U[0] = ` __    __  `
	U[1] = `|\_\  |\_\ `
	U[2] = `| | | | | |`
	U[3] = `| | | | | |`
	U[4] = `| | |_| | |`
	U[5] = `| | |__\| |`
	U[6] = ` \|_______|`
	U[7] = `           `

	V := make([]string, 8)
	V[0] = ` __    __  `
	V[1] = `|\_\  |\_\ `
	V[2] = `||  | | | |`
	V[3] = `||  \/ /  |`
	V[4] = ` \\  \/  / `
	V[5] = `  \\    /  `
	V[6] = `   \\__/   `
	V[7] = `           `

	W := make([]string, 8)
	W[0] = ` __    __  `
	W[1] = `|\ _\ |\_\ `
	W[2] = `| | | | | |`
	W[3] = `| | |/\ | |`
	W[4] = `| | |/ \| |`
	W[5] = `| |       |`
	W[6] = ` \|__/ \__|`
	W[7] = `           `

	X := make([]string, 8)
	X[0] = ` __    __  `
	X[1] = `|\_\  /\_\ `
	X[2] = ` \\  \/  / `
	X[3] = `  \\    /  `
	X[4] = `  //    \  `
	X[5] = ` //  /\\ \ `
	X[6] = `\/_/   \\_\`
	X[7] = `           `

	Y := make([]string, 8)
	Y[0] = ` __    __  `
	Y[1] = `|\_\  /\_\ `
	Y[2] = ` \\  \/  / `
	Y[3] = `  \\    /  `
	Y[4] = `  | |  |   `
	Y[5] = `  | |  |   `
	Y[6] = `   \|__|   `
	Y[7] = `           `

	Z := make([]string, 8)
	Z[0] = ` ________  `
	Z[1] = `|\_______\ `
	Z[2] = ` \|___    |`
	Z[3] = `   / /  /  `
	Z[4] = `  / / /__  `
	Z[5] = `|\/  /___\ `
	Z[6] = ` \|_______|`
	Z[7] = `           `

	BANG := make([]string, 8)
	BANG[0] = ` _     `
	BANG[1] = `|\_\   `
	BANG[2] = `| | |  `
	BANG[3] = `| | |  `
	BANG[4] = ` \|_|  `
	BANG[5] = `|\_\   `
	BANG[6] = ` \|_|  `
	BANG[7] = `       `

	PERIOD := make([]string, 8)
	PERIOD[0] = `       `
	PERIOD[1] = `       `
	PERIOD[2] = `       `
	PERIOD[3] = `       `
	PERIOD[4] = ` _     `
	PERIOD[5] = `|\_\   `
	PERIOD[6] = ` \|_|  `
	PERIOD[7] = `       `

	DASH := make([]string, 8)
	DASH[0] = `           `
	DASH[1] = `           `
	DASH[2] = ` ______    `
	DASH[3] = `|\______\  `
	DASH[4] = ` \|______| `
	DASH[5] = `           `
	DASH[6] = `           `
	DASH[7] = `           `

	UNDERSCORE := make([]string, 8)
	UNDERSCORE[0] = `           `
	UNDERSCORE[1] = `           `
	UNDERSCORE[2] = `           `
	UNDERSCORE[3] = `           `
	UNDERSCORE[4] = ` ______    `
	UNDERSCORE[5] = `|\______\  `
	UNDERSCORE[6] = ` \|______| `
	UNDERSCORE[7] = `           `

	QUESTIONMARK := make([]string, 8)
	QUESTIONMARK[0] = ` ______  `
	QUESTIONMARK[1] = `|\______\`
	QUESTIONMARK[2] = `| |  __  |`
	QUESTIONMARK[3] = ` \|_| | |`
	QUESTIONMARK[4] = `    \/_/ `
	QUESTIONMARK[5] = `   |\_\  `
	QUESTIONMARK[6] = `    \|_| `
	QUESTIONMARK[7] = `         `

	SPACE := make([]string, 8)
	SPACE[0] = `        `
	SPACE[1] = `        `
	SPACE[2] = `        `
	SPACE[3] = `        `
	SPACE[4] = `        `
	SPACE[5] = `        `
	SPACE[6] = `        `
	SPACE[7] = `        `

	BLANK := make([]string, 8)
	BLANK[0] = `   `
	BLANK[1] = `   `
	BLANK[2] = `   `
	BLANK[3] = `   `
	BLANK[4] = `   `
	BLANK[5] = `   `
	BLANK[6] = `   `
	BLANK[7] = `   `

	letterMap := make(map[string][]string)

	letterMap["A"] = A
	letterMap["B"] = B
	letterMap["C"] = C
	letterMap["D"] = D
	letterMap["E"] = E
	letterMap["F"] = F
	letterMap["G"] = G
	letterMap["H"] = H
	letterMap["I"] = I
	letterMap["J"] = J
	letterMap["K"] = K
	letterMap["L"] = L
	letterMap["M"] = M
	letterMap["N"] = N
	letterMap["O"] = O
	letterMap["P"] = P
	letterMap["Q"] = Q
	letterMap["R"] = R
	letterMap["S"] = S
	letterMap["T"] = T
	letterMap["U"] = U
	letterMap["V"] = V
	letterMap["W"] = W
	letterMap["X"] = X
	letterMap["Y"] = Y
	letterMap["Z"] = Z

	letterMap["-"] = DASH
	letterMap["_"] = UNDERSCORE
	letterMap["!"] = BANG
	letterMap["."] = PERIOD
	letterMap[" "] = SPACE
	letterMap["?"] = QUESTIONMARK
	if letterMap[str] != nil {
		return letterMap[str]
	} else {
		ltr := BLANK
		ltr[6] = fmt.Sprintf(" %s ", str)
		return ltr
	}
}

func GetWidth() (int, error) {
	cmd := exec.Command("tput", "cols")
	w, _ := cmd.Output()
	buf := bytes.NewBuffer(w)
	return strconv.Atoi(strings.TrimSpace(buf.String()))
}

func Blockify(str string, w int) string {
	strs := strings.Split(strings.ToUpper(str), "")
	res := make([]string, 8)
	for _, s := range strs {
		var ltr []string
		if s == "\n" {
			r := strings.Join(res, "\n")
			res = make([]string, 8)
			res[0] = fmt.Sprintf("%s\n", r)
			continue
		}
		ltr = getLetter(s)
		if len(res[1])+len(ltr[1]) > w {
			r := strings.Join(res, "\n")
			res = make([]string, 8)
			res[0] = fmt.Sprintf("%s\n", r)
		}
		for i := 0; i < 8; i++ {
			res[i] = fmt.Sprintf("%s%s", res[i], ltr[i])
		}
	}
	return strings.Join(res, "\n")
}

func PrintBlocks(str string) {
	w, err := GetWidth()
	if err != nil {
		w = 80
	}
	fmt.Println(Blockify(str, w))
}

func MegaFail() {
	PrintBlocks("FAIL!")
}

func NoGo() {
	PrintBlocks("no go!")
}

func BadGo() {
	PrintBlocks("Bad go!")
}

func main() {
	PrintBlocks("no go!")
}
