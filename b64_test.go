package b64

import (
	"bytes"
	"encoding/base64"
	"testing"
)

func TestEncode(t *testing.T) {
	inputs := []string{"abcdef", "abcdefg", "abcdefgh"}
	for _, input := range inputs {
		t.Run("", func(t *testing.T) {
			byteInput := []byte(input)
			got := Encode(byteInput)
			want := base64.StdEncoding.EncodeToString(byteInput)
			if got != want {
				t.Errorf("Encode(%v) = %v, want %v", byteInput, got, want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	inputs := []string{"abcdef", "abcdefg", "abcdefgh"}
	for _, input := range inputs {
		t.Run("", func(t *testing.T) {
			byteInput := []byte(input)
			decodeInput := Encode(byteInput)
			got := Decode(decodeInput)
			if !bytes.Equal(got, byteInput) {
				t.Errorf("Decode(%v) = %v, want %v", decodeInput, got, byteInput)
			}
		})
	}
}
