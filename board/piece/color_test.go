package piece

import "testing"

func TestValidateWhite(t *testing.T) {
	c := White
	res := c.validate()
	if res != nil {
		t.Errorf("failed in validate white, err : %v", res.Error())
	}
}

func TestValidateBlack(t *testing.T) {
	c := Black
	res := c.validate()
	if res != nil {
		t.Errorf("failed in validate black, err : %v", res.Error())
	}
}

func TestValidateNoColor(t *testing.T) {
	c := NoColor
	res := c.validate()
	if res != nil {
		t.Errorf("failed in validate no color, err : %v", res.Error())
	}
}

func TestValidateExpectSomeError(t *testing.T) {
	c := Color(0b111111)
	res := c.validate()
	if res == nil {
		t.Errorf("failed in validate no color, expected some error")
	}
}