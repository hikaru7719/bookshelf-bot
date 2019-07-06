package finder

import (
	"github.com/hikaru7719/bookshelf-bot/domain"
	"strings"
	"testing"
)

func TestCSVFinder_Find(t *testing.T) {
	cases := map[string]struct {
		testCSV    string
		word       string
		expectBook domain.Book
		expectErr  bool
	}{
		"true": {
			testCSV: `
9784526049095,1,営業部門のためのデータウェアハウス入門,中地中／著,日刊工業新聞社,,,\n
9784798142593,1,VMware徹底入門 第4版 VMware vSphere 6.0対応,ヴイエムウェア株式会社／著,翔泳社,,徹底入門,https://cover.openbd.jp/9784798142593.jpg\n
9784798122458,1,VMware徹底入門,ヴイエムウェア株式会社／著 ヴイエムウェア／著,翔泳社,,,\n`,
			word:       "営業部門のためのデータウェアハウス入門",
			expectBook: domain.Book{ISBN: "9784526049095", Title: "営業部門のためのデータウェアハウス入門", Author: "中地中／著", Publisher: "日刊工業新聞社"},
			expectErr:  false,
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			testReader := strings.NewReader(tc.testCSV)
			csvFinder := CSV{reader: testReader}
			actualBook, err := csvFinder.Find(tc.word)
			if actualBook[0] != tc.expectBook {
				t.Errorf("want %v,but actual %v\n", tc.expectBook, actualBook[0])
			}
			actualErr := bool(err != nil)
			if actualErr != tc.expectErr {
				t.Errorf("want %t,but actual %t\n", tc.expectErr, actualErr)
			}
		})
	}

}

func TestCSVFinder_Find_InCSV(t *testing.T) {
	csv, err := NewCSV()
	if err != nil {
		t.Fatal(err)
	}
	bookSlice, err := csv.Find("Amazon")
	if len(bookSlice) < 0 {
		t.Errorf("want book Slice length more than 1")
	}
}
