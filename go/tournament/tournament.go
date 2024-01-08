package tournament

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

const format = "%-30v | %2v | %2v | %2v | %2v | %2v\n"

type team struct {
	name             string
	won, lost, drawn int
}

func (t *team) points() int { return 3*t.won + t.drawn }
func (t *team) played() int { return t.won + t.lost + t.drawn }

func getTeam(t *[]team, name string) *team {
	for i := range *t {
		if (*t)[i].name == name {
			return &(*t)[i]
		}
	}
	*t = append(*t, team{name: name})
	return &(*t)[len(*t)-1]
}

func Tally(r io.Reader, w io.Writer) error {

	// Read and parse input
	teams := make([]team, 0, 4)
	data, _ := io.ReadAll(r)
	for _, line := range strings.Split(string(data), "\n") {

		// Skip blank lines and comments
		if line == "" || line[0] == '#' {
			continue
		}

		// Extract fields and validate that there are 3
		fields := strings.Split(line, ";")
		if len(fields) != 3 {
			return fmt.Errorf("wrong field count")
		}

		// Get teams from slice (with lazy initialisation)
		t1, t2 := getTeam(&teams, fields[0]), getTeam(&teams, fields[1])

		// Update teams
		if fields[2] == "win" {
			t1.won, t2.lost = t1.won+1, t2.lost+1
		} else if fields[2] == "loss" {
			t2.won, t1.lost = t2.won+1, t1.lost+1
		} else if fields[2] == "draw" {
			t1.drawn, t2.drawn = t1.drawn+1, t2.drawn+1
		} else {
			return fmt.Errorf("invalid outcome")
		}
	}

	// Sort data
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points() == teams[j].points() {
			return teams[i].name < teams[j].name
		}
		return teams[i].points() > teams[j].points()
	})

	// Write report
	fmt.Fprintf(w, format, "Team", "MP", "W", "D", "L", "P")
	for _, t := range teams {
		fmt.Fprintf(w, format, t.name, t.played(), t.won, t.drawn, t.lost, t.points())
	}

	return nil
}
