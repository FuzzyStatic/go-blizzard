package blizzard

import (
	"fmt"
	"testing"
)

func TestWoWCharacterMythicKeystoneProfile(t *testing.T) {
	dat, err := c.WoWCharacterMythicKeystoneProfile("illidan", "norilockxo")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWCharacterMythicKeystoneProfileSeason(t *testing.T) {
	dat, err := c.WoWCharacterMythicKeystoneProfileSeason("illidan", "norilockxo", 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}