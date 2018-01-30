package library

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/mohuishou/scu"

	"github.com/PuerkitoBio/goquery"

	"github.com/gocolly/colly"
)

// Library 图书馆操作类
type Library struct {
	URL string
	c   *colly.Collector
}

// LoanBook 借阅的书籍
type LoanBook struct {
	BookID      string  `json:"book_id"` // 书籍id 用于续借
	Author      string  `json:"author"`
	Title       string  `json:"title"`
	PublishYear int     `json:"publish_year"` // 出版年
	DueDate     string  `json:"due_date"`     // 到期日期
	ReturnDate  string  `json:"return_date"`  // 归还日期(借阅历史)
	ReturnTime  string  `json:"return_time"`  // 归还时间(借阅历史)
	Arrearage   float64 `json:"arrearage"`    // 欠费
	Address     string  `json:"address"`      // 分馆
	Number      string  `json:"number"`       // 索书号(当前借阅)
}

// NewLibrary 新建一个图书馆对象
func NewLibrary(studentID, password string) (*Library, error) {
	c, err := scu.NewCollector(studentID, password)
	if err != nil {
		return nil, err
	}

	urlstr := ""
	c.OnHTML("#header > a:nth-child(1)", func(e *colly.HTMLElement) {
		urlstr = e.Attr("href")

		if urlstr != "" {
			uri, err := url.Parse(urlstr)
			if err == nil {
				urlstr = "http://opac.scu.edu.cn:8080" + uri.EscapedPath()
			}
		}
	})

	c.Visit("http://opac.scu.edu.cn:8118/ice/login_ice.jsp?type=borinfo")

	lib := &Library{
		URL: urlstr,
		c:   c.Clone(),
	}
	return lib, nil
}

// GetLoan 获取当前借阅
func (lib *Library) GetLoan() []LoanBook {
	books := make([]LoanBook, 0)
	lib.c.OnHTML("body > center > center table", func(e *colly.HTMLElement) {
		e.DOM.Find("tr").Each(func(i int, s *goquery.Selection) {
			if s.Find("td:nth-child(6)").Text() == "" {
				return
			}
			book := LoanBook{}
			v := reflect.ValueOf(&book)
			elem := v.Elem()
			typeOfBook := elem.Type()
			eq := 1
			for k := 0; k < elem.NumField(); k++ {
				val := strings.TrimSpace(s.Find("td").Eq(eq).Text())
				switch typeOfBook.Field(k).Name {
				case "ReturnDate", "ReturnTime":
				case "BookID":
					elem.Field(k).SetString(strings.TrimSpace(s.Find("td").Eq(eq).Find("input").AttrOr("name", "")))
					eq++
				case "PublishYear":
					v, _ := strconv.Atoi(val)
					elem.Field(k).SetInt(int64(v))
					eq++
				case "Arrearage":
					v, _ := strconv.ParseFloat(val, 10)
					elem.Field(k).SetFloat(v)
					eq++
				default:
					elem.Field(k).SetString(val)
					eq++
				}
			}
			books = append(books, book)
		})
	})
	lib.c.Visit(lib.URL + "?func=bor-loan&adm_library=SCU50")
	return books
}

// GetLoanAll 获取历史借阅
func (lib *Library) GetLoanAll() []LoanBook {
	books := make([]LoanBook, 0)
	lib.c.OnHTML("body > center table", func(e *colly.HTMLElement) {
		e.DOM.Find("tr").Each(func(i int, s *goquery.Selection) {
			if s.Find("td:nth-child(6)").Text() == "" {
				return
			}
			book := LoanBook{}
			v := reflect.ValueOf(&book)
			elem := v.Elem()
			typeOfBook := elem.Type()
			eq := 1
			for k := 0; k < elem.NumField(); k++ {
				val := strings.TrimSpace(s.Find("td").Eq(eq).Text())
				switch typeOfBook.Field(k).Name {
				case "Number", "BookID":
				case "PublishYear":
					v, _ := strconv.Atoi(val)
					elem.Field(k).SetInt(int64(v))
					eq++
				case "Arrearage":
					v, _ := strconv.ParseFloat(val, 10)
					elem.Field(k).SetFloat(v)
					eq++
				default:
					elem.Field(k).SetString(val)
					eq++
				}
			}
			books = append(books, book)
		})
	})
	lib.c.Visit(lib.URL + "?func=bor-history-loan&adm_library=SCU50")
	return books
}

// Loan 续借
func (lib *Library) Loan(bookID string) bool {
	return lib.loan(bookID)
}

// LoanAll 续借全部
func (lib *Library) LoanAll() bool {
	return lib.loan("")
}

// bookID 等于空时为续借全部
func (lib *Library) loan(bookID string) bool {
	ok := false
	param := "func=bor-renew-all&adm_library=SCU50"
	if bookID != "" {
		param += fmt.Sprintf("&renew_selected=Y&%s=Y", bookID)
	}
	lib.c.OnHTML("body", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "续借不成功") {
			ok = false
		} else if strings.Contains(e.Text, "续借成功") {
			ok = true
		}
	})
	lib.c.Visit(lib.URL + "?" + param)
	return ok
}

// getURL 获取真实url
func getURL() string {
	c := colly.NewCollector()
	urlstr := ""
	c.OnHTML("#header > a:nth-child(1)", func(e *colly.HTMLElement) {
		urlstr = e.Attr("href")

		if urlstr != "" {
			uri, err := url.Parse(urlstr)
			if err == nil {
				urlstr = "http://opac.scu.edu.cn:8080" + uri.EscapedPath()
			}
		}
	})

	c.Visit("http://opac.scu.edu.cn:8080/F")
	return urlstr
}
