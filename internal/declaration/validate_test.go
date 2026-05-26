package declaration

import "testing"

func TestIsEnumNameValid(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"simple capitalized", "Color", true},
		{"capitalized with digits", "Color2", true},
		{"camel case", "ColorName", true},
		{"with underscore", "Color_Name", true},
		{"single letter", "A", true},
		{"all caps", "RGB", true},

		{"empty", "", false},
		{"lowercase start", "color", false},
		{"digit start", "1Color", false},
		{"underscore start", "_Color", false},
		{"contains dash", "Color-Name", false},
		{"contains space", "Color Name", false},
		{"contains dot", "Color.Name", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := isEnumNameValid(tc.input)
			if got != tc.want {
				t.Errorf("isEnumNameValid(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestIsPackageNameValid(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"simple lowercase", "color", true},
		{"with underscore", "color_pkg", true},
		{"single letter", "a", true},
		{"only underscores after first", "c__", true},

		{"empty", "", false},
		{"uppercase start", "Color", false},
		{"contains digit", "color1", false},
		{"digit start", "1color", false},
		{"underscore start", "_color", false},
		{"contains dash", "color-pkg", false},
		{"contains dot", "color.pkg", false},
		{"all caps", "COLOR", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := isPackageNameValid(tc.input)
			if got != tc.want {
				t.Errorf("isPackageNameValid(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
