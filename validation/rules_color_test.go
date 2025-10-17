package validation

import (
	"testing"
)

// ============================================================================
// HEXCOLOR VALIDATOR TESTS
// ============================================================================

func TestHexColor(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid hex colors (6 digits)
		{"valid 6-digit lowercase", "#ff5733", false},
		{"valid 6-digit uppercase", "#FF5733", false},
		{"valid 6-digit mixed case", "#Ff5733", false},
		{"valid 6-digit without #", "ff5733", false},
		{"valid black", "#000000", false},
		{"valid white", "#FFFFFF", false},

		// Valid hex colors (3 digits)
		{"valid 3-digit lowercase", "#f57", false},
		{"valid 3-digit uppercase", "#F57", false},
		{"valid 3-digit mixed case", "#F5a", false},
		{"valid 3-digit without #", "f57", false},
		{"valid black 3-digit", "#000", false},
		{"valid white 3-digit", "#FFF", false},

		// Invalid hex colors
		{"empty string", "", true},
		{"only #", "#", true},
		{"too short", "#ff", true},
		{"too long", "#ff57333", true},
		{"invalid char g", "#ff573g", true},
		{"invalid char z", "#ff573z", true},
		{"4 digits", "#ff57", true},
		{"5 digits", "#ff573", true},
		{"7 digits", "#ff57331", true},
		{"contains space", "#ff 733", true},
		{"not a string", 12345, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "color_field")
			v.HexColor()
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})
	}
}

// ============================================================================
// RGB VALIDATOR TESTS
// ============================================================================

func TestRGB(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid RGB colors
		{"valid RGB", "rgb(255, 87, 51)", false},
		{"valid RGB no spaces", "rgb(255,87,51)", false},
		{"valid RGB extra spaces", "rgb( 255 , 87 , 51 )", false},
		{"valid RGB min values", "rgb(0, 0, 0)", false},
		{"valid RGB max values", "rgb(255, 255, 255)", false},
		{"valid RGB mid values", "rgb(128, 128, 128)", false},

		// Invalid RGB colors
		{"empty string", "", true},
		{"missing rgb prefix", "(255, 87, 51)", true},
		{"wrong prefix rgba", "rgba(255, 87, 51)", true},
		{"only two values", "rgb(255, 87)", true},
		{"four values", "rgb(255, 87, 51, 0.5)", true},
		{"red out of range high", "rgb(256, 87, 51)", true},
		{"green out of range high", "rgb(255, 256, 51)", true},
		{"blue out of range high", "rgb(255, 87, 256)", true},
		{"negative red", "rgb(-1, 87, 51)", true},
		{"negative green", "rgb(255, -1, 51)", true},
		{"negative blue", "rgb(255, 87, -1)", true},
		{"decimal values", "rgb(255.5, 87.5, 51.5)", true},
		{"not a string", 12345, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "color_field")
			v.RGB()
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})
	}
}

// ============================================================================
// RGBA VALIDATOR TESTS
// ============================================================================

func TestRGBA(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid RGBA colors
		{"valid RGBA", "rgba(255, 87, 51, 0.8)", false},
		{"valid RGBA no spaces", "rgba(255,87,51,0.8)", false},
		{"valid RGBA extra spaces", "rgba( 255 , 87 , 51 , 0.8 )", false},
		{"valid RGBA alpha 0", "rgba(255, 87, 51, 0)", false},
		{"valid RGBA alpha 0.0", "rgba(255, 87, 51, 0.0)", false},
		{"valid RGBA alpha 1", "rgba(255, 87, 51, 1)", false},
		{"valid RGBA alpha 1.0", "rgba(255, 87, 51, 1.0)", false},
		{"valid RGBA alpha 0.5", "rgba(255, 87, 51, 0.5)", false},
		{"valid RGBA min RGB", "rgba(0, 0, 0, 0.5)", false},
		{"valid RGBA max RGB", "rgba(255, 255, 255, 0.5)", false},

		// Invalid RGBA colors
		{"empty string", "", true},
		{"missing rgba prefix", "(255, 87, 51, 0.8)", true},
		{"wrong prefix rgb", "rgb(255, 87, 51, 0.8)", true},
		{"only three values", "rgba(255, 87, 51)", true},
		{"five values", "rgba(255, 87, 51, 0.8, 0)", true},
		{"red out of range", "rgba(256, 87, 51, 0.8)", true},
		{"green out of range", "rgba(255, 256, 51, 0.8)", true},
		{"blue out of range", "rgba(255, 87, 256, 0.8)", true},
		{"alpha out of range high", "rgba(255, 87, 51, 1.1)", true},
		{"alpha out of range low", "rgba(255, 87, 51, -0.1)", true},
		{"negative red", "rgba(-1, 87, 51, 0.8)", true},
		{"negative green", "rgba(255, -1, 51, 0.8)", true},
		{"negative blue", "rgba(255, 87, -1, 0.8)", true},
		{"not a string", 12345, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "color_field")
			v.RGBA()
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})
	}
}

// ============================================================================
// HSL VALIDATOR TESTS
// ============================================================================

func TestHSL(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid HSL colors
		{"valid HSL", "hsl(9, 100%, 60%)", false},
		{"valid HSL no spaces", "hsl(9,100%,60%)", false},
		{"valid HSL extra spaces", "hsl( 9 , 100% , 60% )", false},
		{"valid HSL hue 0", "hsl(0, 100%, 60%)", false},
		{"valid HSL hue 360", "hsl(360, 100%, 60%)", false},
		{"valid HSL saturation 0", "hsl(180, 0%, 60%)", false},
		{"valid HSL saturation 100", "hsl(180, 100%, 60%)", false},
		{"valid HSL lightness 0", "hsl(180, 50%, 0%)", false},
		{"valid HSL lightness 100", "hsl(180, 50%, 100%)", false},
		{"valid HSL mid values", "hsl(180, 50%, 50%)", false},

		// Invalid HSL colors
		{"empty string", "", true},
		{"missing hsl prefix", "(180, 50%, 50%)", true},
		{"wrong prefix hsla", "hsla(180, 50%, 50%)", true},
		{"only two values", "hsl(180, 50%)", true},
		{"four values", "hsl(180, 50%, 50%, 0.5)", true},
		{"hue out of range high", "hsl(361, 50%, 50%)", true},
		{"hue out of range low", "hsl(-1, 50%, 50%)", true},
		{"saturation out of range high", "hsl(180, 101%, 50%)", true},
		{"saturation out of range low", "hsl(180, -1%, 50%)", true},
		{"lightness out of range high", "hsl(180, 50%, 101%)", true},
		{"lightness out of range low", "hsl(180, 50%, -1%)", true},
		{"missing % on saturation", "hsl(180, 50, 50%)", true},
		{"missing % on lightness", "hsl(180, 50%, 50)", true},
		{"decimal hue", "hsl(180.5, 50%, 50%)", true},
		{"decimal saturation", "hsl(180, 50.5%, 50%)", true},
		{"decimal lightness", "hsl(180, 50%, 50.5%)", true},
		{"not a string", 12345, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "color_field")
			v.HSL()
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})
	}
}

// ============================================================================
// MULTI-FIELD COLOR VALIDATION TESTS
// ============================================================================

func TestMultiFieldColorValidation(t *testing.T) {
	type ThemeColors struct {
		Primary   string
		Secondary string
		Accent    string
		Overlay   string
	}

	colors := ThemeColors{
		Primary:   "#FF5733",
		Secondary: "rgb(255, 87, 51)",
		Accent:    "hsl(9, 100%, 60%)",
		Overlay:   "rgba(255, 87, 51, 0.8)",
	}

	mv := NewValidator()
	mv.Field(colors.Primary, "primary").HexColor()
	mv.Field(colors.Secondary, "secondary").RGB()
	mv.Field(colors.Accent, "accent").HSL()
	mv.Field(colors.Overlay, "overlay").RGBA()

	err := mv.Validate()
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
}

func TestColorValidationChaining(t *testing.T) {
	// Test chaining with other validators
	color := "#FF5733"

	v := New(color, "brand_color")
	v.Required().HexColor().MinLength(4)

	err := v.Validate()
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
}

func TestColorValidationWithStopOnError(t *testing.T) {
	v := New("invalid", "color").StopOnError()
	v.HexColor().RGB()

	err := v.Validate()
	if err == nil {
		t.Error("expected error but got none")
	}

	// Should only have one error due to StopOnError
	if len(v.errors) > 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(v.errors))
	}
}
