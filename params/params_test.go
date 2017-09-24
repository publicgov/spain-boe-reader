package params

import "testing"

func TestParamURL(t *testing.T) {
	expected := "http://boe.es/diario_boe/xml.php?id=BOE-S-20141006"
	params := &Params{"BOE", "S", "20141006"}
	result := params.ToString()
	if result != expected {
		t.Error("Expected ", expected, " got ", result)
	}
}
