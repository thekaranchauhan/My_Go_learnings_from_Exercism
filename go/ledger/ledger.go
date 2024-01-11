package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type localeSettings struct {
	Date        string
	Description string
	Change      string
	TSep        string
	Layout      map[bool]string
	DateFunc    func(time.Time) string
}

var locales = map[string]localeSettings{
	"nl-NL": localeSettings{"Datum", "Omschrijving", "Verandering", ".", map[bool]string{true: "%s %s,%02d-", false: "%s %s,%02d "},
		func(t time.Time) string { return fmt.Sprintf("%02d-%02d-%04d", t.Day(), t.Month(), t.Year()) }},
	"en-US": localeSettings{"Date", "Description", "Change", ",", map[bool]string{true: "(%s%s.%02d)", false: "%s%s.%02d "},
		func(t time.Time) string { return fmt.Sprintf("%02d/%02d/%04d", t.Month(), t.Day(), t.Year()) }},
}

func thousandSeparated(n int, sep string) string {
	arr := []string{strconv.Itoa(n % 1000)}
	for n /= 1000; n > 0; n /= 1000 {
		arr = append([]string{strconv.Itoa(n % 1000)}, arr...)
	}
	return strings.Join(arr, sep)
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	currencies := map[string]string{"EUR": "€", "USD": "$"}
	symbol, ok := currencies[currency]
	if !ok {
		return "", errors.New("unsupported currency")
	}
	settings, ok := locales[locale]
	if !ok {
		return "", errors.New("unsupported locale")
	}
	output := make([]string, len(entries)+1)
	output[0] = fmt.Sprintf("%-10s | %-25s | %s", settings.Date, settings.Description, settings.Change)
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)

	sort.Slice(entriesCopy, func(i, j int) bool {
		if entriesCopy[i].Date != entriesCopy[j].Date {
			return entriesCopy[i].Date < entriesCopy[j].Date
		} else if entriesCopy[i].Description != entriesCopy[j].Description {
			return entriesCopy[i].Description < entriesCopy[j].Description
		}
		return entriesCopy[i].Change < entriesCopy[j].Change
	})

	for i, entry := range entriesCopy {
		if len(entry.Description) > 25 {
			entry.Description = entry.Description[:22] + "..."
		}
		tm, err := time.Parse("2006-01-02", entry.Date)
		if err != nil {
			return "", err
		}
		date := settings.DateFunc(tm)
		negative := entry.Change < 0
		if negative {
			entry.Change = -entry.Change
		}
		amount := fmt.Sprintf(settings.Layout[negative], symbol, thousandSeparated(entry.Change/100, settings.TSep), entry.Change%100)
		output[i+1] = fmt.Sprintf("%10s | %-25s | %13s", date, entry.Description, amount)
	}
	return strings.Join(output, "\n") + "\n", nil
}
