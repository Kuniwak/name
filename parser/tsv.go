package parser

import (
	"bufio"
	"fmt"
	"github.com/Kuniwak/name/eval"
	"github.com/Kuniwak/name/filter"
	"github.com/Kuniwak/name/mora"
	"golang.org/x/text/unicode/norm"
	"io"
	"strconv"
	"strings"
)

func ParseRank(s string) (eval.Rank, error) {
	switch s {
	case "大大吉":
		return eval.DaiDaiKichi, nil
	case "大吉":
		return eval.DaiKichi, nil
	case "吉":
		return eval.Kichi, nil
	case "凶":
		return eval.Kyo, nil
	case "大凶":
		return eval.DaiKyo, nil
	default:
		return 0, fmt.Errorf("invalid rank: %q", s)
	}
}

func ParseByte(s string) (byte, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid byte: %w", err)
	}
	if i < 0 || i > 255 {
		return 0, fmt.Errorf("invalid byte: %d", i)
	}
	return byte(i), nil
}

func ParseTSV(r io.Reader, ch chan<- filter.Target) error {
	defer close(ch)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fields := strings.Split(line, "\t")
		if len(fields) != 9 {
			return fmt.Errorf("invalid number of fields: %q", line)
		}

		isHeader := fields[0] == "評点"
		if isHeader {
			continue
		}

		strokes, err := ParseByte(fields[1])
		if err != nil {
			return fmt.Errorf("invalid strokes: %w", err)
		}

		yomi := norm.NFC.String(fields[3])
		yomiRunes := []rune(yomi)

		m := mora.Count(yomiRunes)

		tenkaku, err := ParseRank(fields[4])
		if err != nil {
			return fmt.Errorf("invalid tenkaku: %w", err)
		}

		chikaku, err := ParseRank(fields[5])
		if err != nil {
			return fmt.Errorf("invalid chikaku: %w", err)
		}

		jinkaku, err := ParseRank(fields[6])
		if err != nil {
			return fmt.Errorf("invalid jinkaku: %w", err)
		}

		gaikaku, err := ParseRank(fields[7])
		if err != nil {
			return fmt.Errorf("invalid gaikaku: %w", err)
		}

		sokaku, err := ParseRank(fields[8])
		if err != nil {
			return fmt.Errorf("invalid sokaku: %w", err)
		}

		ch <- filter.Target{
			Kanji:      []rune(norm.NFC.String(fields[2])),
			Yomi:       yomiRunes,
			YomiString: yomi,
			Strokes:    strokes,
			Mora:       m,
			EvalResult: eval.Result{
				Tenkaku: tenkaku,
				Jinkaku: jinkaku,
				Chikaku: chikaku,
				Gaikaku: gaikaku,
				Sokaku:  sokaku,
			},
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}