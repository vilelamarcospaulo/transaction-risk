package level

type RiskLevel int64

const (
	Low RiskLevel = iota
	Medium
	High
)

func (r RiskLevel) String() string {
	switch r {
	case 0:
		return "Low"
	case 1:
		return "Medium"
	case 2:
		return "High"
	}
	return ""
}
