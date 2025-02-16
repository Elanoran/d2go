package itemfilter

import (
	"github.com/Elanoran/d2go/pkg/data/stat"
	"strings"

	"github.com/Elanoran/d2go/pkg/data"
	"github.com/Elanoran/d2go/pkg/nip"
)

func Evaluate(i data.Item, rules []nip.Rule) bool {
	for _, r := range rules {
		if !evaluateGroups(i, r.Properties, checkProperty) {
			// Properties not matching, skipping
			continue
		}

		// We can not check stats, item is not identified, but properties matching
		if !i.Identified {
			return true
		}

		if evaluateGroups(i, r.Stats, checkStat) {
			return true
		}
	}

	return false
}

func evaluateGroups(i data.Item, groups []nip.Group, evalFunc func(i data.Item, prop nip.Comparable) bool) bool {
	groupChain := evaluationChain{}
	for _, group := range groups {
		propChain := evaluationChain{}
		for _, st := range group.Comparable {
			propChain.Add(evalFunc(i, st), st.Operand)
		}
		groupChain.Add(propChain.Evaluate(), group.Operand)
	}

	return groupChain.Evaluate()
}

func checkStat(i data.Item, cmp nip.Comparable) bool {
	st, found := stats[cmp.Keyword]
	if !found {
		// pass it, just in case...
		return true
	}

	itemStat, found := i.Stats[stat.ID(st[0])]
	if !found {
		return false
	}

	if !compare(itemStat.Value, cmp.ValueInt, cmp.Comparison) {
		return false
	}

	if len(st) == 1 {
		return true
	}

	return st[1] == itemStat.Layer
}

func checkProperty(i data.Item, prop nip.Comparable) bool {
	switch prop.Keyword {
	case nip.PropertyType:
		return strings.EqualFold(i.Type(), prop.ValueString)
	case nip.PropertyName:
		return strings.EqualFold(string(i.Name), prop.ValueString)
	case nip.PropertyClass:
		// TODO: Implement
	case nip.PropertyQuality:
		quality, found := qualities[prop.ValueString]
		if !found {
			return false
		}

		return compare(int(i.Quality), int(quality), prop.Comparison)
	case nip.PropertyFlag:
		if prop.Comparison == nip.OperandEqual && !i.Ethereal {
			return false
		}
		if prop.Comparison == nip.OperandNotEqualTo && i.Ethereal {
			return false
		}
	case nip.PropertyLevel:
		// TODO: Implement
	case nip.PropertyPrefix:
		// TODO: Implement
	case nip.PropertySuffix:
		// TODO: Implement
	}

	return true
}

func compare(val1, val2 int, operand nip.Operand) bool {
	switch operand {
	case nip.OperandEqual:
		return val1 == val2
	case nip.OperandGreaterThan:
		return val1 > val2
	case nip.OperandGreaterOrEqualTo:
		return val1 >= val2
	case nip.OperandLessThan:
		return val1 < val2
	case nip.OperandLessThanOrEqualTo:
		return val1 <= val2
	case nip.OperandNotEqualTo:
		return val1 != val2
	}

	return false
}

func EvaluateWithNote(i data.Item, rules []nip.Rule) (bool, string) {
	for _, r := range rules {
		if !evaluateGroups(i, r.Properties, checkProperty) {
			// Properties not matching, skipping
			continue
		}

		// We can not check stats, item is not identified, but properties matching
		if !i.Identified {
			return true, ""
		}

		if evaluateGroups(i, r.Stats, checkStat) {
			return true, extractNoteFromComment(r.Comment)
		}
	}

	return false, ""
}

func extractNoteFromComment(comment string) string {
	// Assuming the note is after the last occurrence of "//"
	index := strings.LastIndex(comment, "//")
	if index == -1 {
		return ""
	}

	// Trim any leading or trailing white spaces
	return strings.TrimSpace(comment[index+2:])
}

