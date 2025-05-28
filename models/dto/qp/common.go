package qp

var priorityMp = map[string]string{
	":P1": "S(影响播放、会员、收入)",
	":P2": "A(有直接影响的其他外部服务或重要基础服务)",
	":P3": "B(有间接影响的重要支持功能)",
	":P4": "C(有影响但风险可控的服务)",
	":P5": "D(测试使用)",
}

var priorityMpReverse = map[string]string{
	"S": ":P1",
	"A": ":P2",
	"B": ":P3",
	"C": ":P4",
	"D": ":P5",
}

func ReplacePrio(prio string) string {
	if p, ok := priorityMp[prio]; ok {
		return p
	}

	return ""
}

func ReplacePrioReverse(prio string) string {
	if p, ok := priorityMpReverse[prio]; ok {
		return p
	}

	return ""
}
