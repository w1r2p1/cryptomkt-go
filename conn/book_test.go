package conn

import (
	"math/rand"
	"testing"
	"time"

	"github.com/cryptomkt/cryptomkt-go/args"
)

var argus [8]args.Argument = [8]args.Argument{args.Market("ETHCLP"), args.Type("buy"), args.Type("sell"), args.Page(0), args.Limit(50), args.Start("2017-03-03"), args.End("2018-03-03"), args.Timeframe("60")}


func TestGetBookPage(t *testing.T) {
	client := NewClient("NoKey", "NoSecret")
	rand.Seed(time.Now().UnixNano())
	var optional [2]args.Argument = [2]args.Argument{argus[3], argus[4]}
	for i := 0; i < 100; i++ {
		var numArgs int = rand.Intn(3)
		switch numArgs {
		case 0:
			if _, err := client.GetBookPage(argus[0], argus[1]); err != nil {
				t.Errorf("Book with cero optional args failed: %s", err)
			}
		case 1:
			var random int = rand.Intn(2)
			if _, err := client.GetBookPage(argus[0], argus[1], optional[random]); err != nil {
				t.Errorf("Book with %v optional args failed: %s", 1, err)
			}
		case 2:
			if _, err := client.GetBookPage(argus[0], argus[1], optional[0], optional[1]); err != nil {
				t.Errorf("Book with 2 optional arguments failed because %s ", err)
			}
		}
		time.Sleep(3 * time.Second)
	}
}

func TestBookGetPrevious(t *testing.T) {
	client := NewClient("NoKey", "NoSecret")
	book, err := client.GetBookPage(args.Market("ETHCLP"), args.Type("buy"), args.Page(1))
	if err != nil {
		t.Errorf("Error getting the book: %s", err)
	}
	_, err = book.GetPrevious()
	if err != nil {
		t.Errorf("Error in previous book: %s", err)
	}
}

func TestBookGetNext(t *testing.T) {
	client := NewClient("NoKey", "NoSecret")
	book, err := client.GetBookPage(args.Market("ETHCLP"), args.Type("buy"), args.Page(0))
	if err != nil {
		t.Errorf("Error getting the book: %s", err)
	}
	_, err = book.GetNext()
	if err != nil {
		t.Errorf("Error in next book: %s", err)
	}
}

func TestGetBooks(t *testing.T) {
	client := NewClient("NoKey", "NoSecret")
	time.Sleep(3 * time.Second)
	if _, err := client.GetBook(args.Market("ETHCLP"),args.Type("buy")); err != nil {
		t.Errorf("failed to retrieve books, %s", err)
	}
}
