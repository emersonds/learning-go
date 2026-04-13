package main

import "testing"

func Example_main() {
	main()
	// Output:
	// Hello world
}

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	// Preparation phase, defines expected values
	var tests = map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello world",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Akkadian, not supported": {
			lang: "akk",
			want: `unsupported language: "akk"`,
		},
		"German": {
			lang: "de",
			want: "Hallo Welt",
		},
		"Greek": {
			lang: "el",
			want: "Χαίρετε Κόσμε",
		},
		"Empty": {
			lang: "",
			want: `unsupported language: ""`,
		},
	}

	// Execution phase, Range over each language
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)

			// Decision phase, check returned value
			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}
}
