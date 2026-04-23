package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Grave struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Born     int    `json:"born"`
	Died     int    `json:"died"`
	Cause    string `json:"cause"`
	Epitaph  string `json:"epitaph"`
}

var allowedCategories = map[string]bool{
	"Gaming":  true,
	"Tech":    true,
	"Retail":  true,
	"Finance": true,
	"Sports":  true,
}

const (
	minYear       = 1000
	minCauseLen   = 2
	maxCauseLen   = 80
	minEpitaphLen = 20
	maxEpitaphLen = 500
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <graves.json>\n", os.Args[0])
		os.Exit(2)
	}
	path := os.Args[1]

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read %s: %v\n", path, err)
		os.Exit(1)
	}

	var graves []Grave
	if err := json.Unmarshal(data, &graves); err != nil {
		fmt.Fprintf(os.Stderr, "parse %s: %v\n", path, err)
		os.Exit(1)
	}

	violations := validate(graves, time.Now().Year())

	if len(violations) > 0 {
		for _, v := range violations {
			fmt.Fprintln(os.Stderr, v)
		}
		fmt.Fprintf(os.Stderr, "\n%d violation(s) in %s\n", len(violations), path)
		os.Exit(1)
	}

	fmt.Printf("OK: %d entries validated in %s\n", len(graves), path)
}

func validate(graves []Grave, currentYear int) []string {
	var out []string
	seen := map[string]int{}

	for i, g := range graves {
		prefix := fmt.Sprintf("entry %d (%q)", i, g.Name)

		if strings.TrimSpace(g.Name) == "" {
			out = append(out, fmt.Sprintf("entry %d: name is empty", i))
		}
		if g.Category == "" {
			out = append(out, fmt.Sprintf("%s: category is empty", prefix))
		} else if !allowedCategories[g.Category] {
			out = append(out, fmt.Sprintf("%s: category %q is not one of Gaming/Tech/Retail/Finance/Sports", prefix, g.Category))
		}
		if g.Born < minYear {
			out = append(out, fmt.Sprintf("%s: born=%d must be >= %d", prefix, g.Born, minYear))
		}
		if g.Died < minYear {
			out = append(out, fmt.Sprintf("%s: died=%d must be >= %d", prefix, g.Died, minYear))
		}
		if g.Born >= minYear && g.Died >= minYear && g.Died < g.Born {
			out = append(out, fmt.Sprintf("%s: died=%d is before born=%d", prefix, g.Died, g.Born))
		}
		if g.Died > currentYear {
			out = append(out, fmt.Sprintf("%s: died=%d is in the future (current year: %d)", prefix, g.Died, currentYear))
		}
		if n := len(g.Cause); n < minCauseLen || n > maxCauseLen {
			out = append(out, fmt.Sprintf("%s: cause length %d not in [%d, %d]", prefix, n, minCauseLen, maxCauseLen))
		}
		if n := len(g.Epitaph); n < minEpitaphLen || n > maxEpitaphLen {
			out = append(out, fmt.Sprintf("%s: epitaph length %d not in [%d, %d]", prefix, n, minEpitaphLen, maxEpitaphLen))
		}

		key := strings.ToLower(strings.TrimSpace(g.Name))
		if key != "" {
			if first, dup := seen[key]; dup {
				out = append(out, fmt.Sprintf("%s: duplicate name (first seen at entry %d)", prefix, first))
			} else {
				seen[key] = i
			}
		}
	}

	return out
}
