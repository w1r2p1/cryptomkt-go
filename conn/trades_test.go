package conn

import (
	"math/rand"
	"testing"
	"time"

	"github.com/cryptomkt/cryptomkt-go/args"
)

func TestTrades(t *testing.T) {
	client := NewClient("NoKey", "NoSecret")
	rand.Seed(time.Now().UnixNano())
	var optional [4]args.Argument = [4]args.Argument{args.Page(0), args.Limit(50), args.Start("2017-03-03"), args.Timeframe("60")}
	for i := 0; i < 100; i++ {
		var numArgs int = rand.Intn(5)
		switch numArgs {
		case 0:
			if _, err := client.GetTrades(args.Market("ETHCLP")); err != nil {
				t.Errorf("Trades with cero optional arguments failed because %s", err)
			}
		case 1:
			var randomIndex int = rand.Intn(4)
			if _, err := client.GetTrades(args.Market("ETHCLP"), argus[randomIndex]); err != nil {
				t.Errorf("Trades with one optional argument failed")
			}
		case 2:
			var randomIndexes []int = generateIndexes(2, 4)
			if _, err := client.GetTrades(args.Market("ETHCLP"), optional[randomIndexes[0]], optional[randomIndexes[1]]); err != nil {
				t.Errorf("Trades with 2 optional arguments failed, %s", err)
			}
		case 3:
			var randomIndexes []int = generateIndexes(3, 4)
			if _, err := client.GetTrades(args.Market("ETHCLP"), optional[randomIndexes[0]], optional[randomIndexes[1]], optional[randomIndexes[2]]); err != nil {
				t.Errorf("Trades with 3 optional arguments failed, %s", err)
			}
		case 4:
			if _, err := client.GetTrades(args.Market("ETHCLP"), optional[0], optional[1], optional[2], optional[3]); err != nil {
				t.Errorf("Trades with 4 optional args failed %s", err)
			}
		}
		time.Sleep(3 * time.Second)
	}
}

func TestTradesGetPrevious(t *testing.T) {
	client := NewClient("NoKey", "NoSecret")
	trades, err := client.GetTrades(args.Market("ETHCLP"), args.Start("2019-12-12"), args.End("2020-01-01"), args.Page(1))
	if err != nil {
		t.Errorf("Error Trades: %s", err)
	}
	_, err = trades.GetPrevious()
	if err != nil {
		t.Errorf("Error in previous trades: %s", err)
	}
}

func TestTradesGetNext(t *testing.T) {
	client := NewClient("NoKey", "NoSecret")
	t.Run("optionals=no", func(t *testing.T) {
		trades, err := client.GetTrades(
			args.Market("ETHCLP"))
		if err != nil {
			t.Errorf("Error Trades: %s", err)
		}
		_, err = trades.GetNext()
		if err != nil {
			t.Errorf("Error in next trades: %s", err)
		}
	})
	t.Run("optionals=yes", func(t *testing.T) {
		trades, err := client.GetTrades(
			args.Market("ETHCLP"),
			args.Start("2019-12-12"),
			args.End("2020-02-21"),
			args.Page(0))
		if err != nil {
			t.Errorf("Error Trades: %s", err)
		}
		_, err = trades.GetNext()
		if err != nil {
			t.Errorf("Error in next trades: %s", err)
		}
	})
}

func TestGetTradesAllPages(t *testing.T) {
	client := NewClient("NoKey", "NoSecret")
	time.Sleep(3 * time.Second)
	_, err := client.GetTradesAllPages(
		args.Market("ETHCLP"),
		args.Start("2019-02-12"),
		args.End("2020-02-21"))
	if err != nil {
		t.Errorf("TestGetTradesAllPages failed: %s", err)
	}
}
