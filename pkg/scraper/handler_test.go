package scraper

import "testing"

// Url is the url we are testing, and the expected stands for the 'result we expect'
type urlEndTest struct {
	url, expected string
}

// Define the tests we want to run
var urlEndTests = []urlEndTest{
	{"https://www.formula1.com/en/results.html/2023/drivers.html", "drivers.html"},
	{"en/results.html/2023/races.html", "races.html"},
	{"https://www.formula1.com/en/results.html/1958/team/ferrari.html", "ferrari.html"},
	{"/en/results.html/1998/drivers/MICSCH01/michael-schumacher.html", "michael-schumacher.html"},
}

// Test the URL function in handler
func TestUrlEnd(t *testing.T) {
	for _, test := range urlEndTests {
		if output := getUrlEnd(test.url); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

// Benchmark the URL function
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getUrlEnd(urlEndTests[0].url)
	}
}
