package main

import (
	"testing"
	"go-reloaded/pkg"
)

func TestGoReloaded(t *testing.T) {
	if pkg.GoReloaded("1E (hex) + 10 (bin) = 32 ,IS A (low, 2) correct a answer .") != "30 + 2 = 32, is a correct an answer." {
		t.Error("Not a correct answer!")
	}
}

func TestDigit(t *testing.T) {
	num, err := pkg.Digit("(up, 2)")
	if num != 2 && err != nil {
		t.Error("Not a correct answer!")
	}
}

func TestSpacialSign(t *testing.T) {
	if pkg.SpacialSign(',') != true {
		t.Error("Not a correct answer!")
	}
}

func TestVowelChar(t *testing.T) {
	if pkg.VowelChar("a") != true {
		t.Error("Not a correct answer!")
	}
}

func TestCheckText(t *testing.T) {
	if pkg.CheckText(".,!?:;") != true {
		t.Error("Not a correct answer!")
	}
}
