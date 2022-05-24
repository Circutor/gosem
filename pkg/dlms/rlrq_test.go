package dlms

import (
	"bytes"
	"testing"
)

func TestEncodeRLRQ(t *testing.T) {
	settings, _ := NewSettingsWithoutAuthentication()

	out, err := EncodeRLRQ(&settings)
	if err != nil {
		t.Errorf("Encode Failed. Err: %v", err)
	}
	result := decodeHexString("6200")
	if !bytes.Equal(out, result) {
		t.Errorf("Failed. Get: %s, should: %s", encodeHexString(out), encodeHexString(result))
	}
}

func TestEncodeRLRWithUserInformation(t *testing.T) {
	ciphering := Ciphering{
		Security:          SecurityEncryption | SecurityAuthentication,
		SystemTitle:       decodeHexString("4349520000000001"),
		UnicastKey:        decodeHexString("00112233445566778899AABBCCDDEEFF"),
		AuthenticationKey: decodeHexString("00112233445566778899AABBCCDDEEFF"),
		DedicatedKey:      decodeHexString("E803739DBE338C3A790D8D1B12C63FE2"),
		InvocationCounter: 0x00000107,
	}

	settings, _ := NewSettingsWithLowAuthenticationAndCiphering([]byte("JuS66BCZ"), ciphering)

	out, err := EncodeRLRQ(&settings)
	if err != nil {
		t.Errorf("Encode Failed. Err: %v", err)
	}

	result := decodeHexString("6239800100BE340432213030000001078E6341442275404C816C6BED3E33AE809EC51E1D0E428BE8F5F643E26C3ED8295297AF055F2BC322DA3BD8")
	if !bytes.Equal(out, result) {
		t.Errorf("Failed. Get: %s, should: %s", encodeHexString(out), encodeHexString(result))
	}
}
