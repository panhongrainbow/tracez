package tracingLru

import (
	"github.com/panhongrainbow/tracez/model"
	"math/rand"
	"time"
)

/*
Use the 108 characters from Water Margin as SpanID for each test data.
1	 卢俊义	 Lu Junyi - Lu Zhishen
2	 花荣	 Hua Rong
3	 花豪	 Hua Xiong
4	 武勇	 Wu Yong
5	 林冲	 Lin Chong
6	 柴进	 Chai Jin
7	 秦明	 Qin Ming
8	 秦庆同	 Qin Qingtong
9	 关胜	 Guan Sheng
10	 包祥	 Bao Xu
11	 包拯	 Bao Zheng
12	 董平	 Dong Ping
13	 何雄	 He Xiong
14	 吕方	 Lü Fang
15	 彭盈予	 Peng Yingyu
16	 严庆	 Yan Qing
17	 阮小二	 Ruan Xiaoer
18	 阮小五	 Ruan Xiaowu
19	 胡三娘	 Hu Sanniang
20	 杨志	 Yang Zhi
21	 徐宁	 Xu Ning
22	 宋江	 Song Jiang
23	 杜兴	 Du Xing
24	 张顺	 Zhang Shun
25	 孙立	 Sun Li
26	 黄兴	 Huang Xin
27	 关勇长	 Guan Yunchang
28	 蒋敬	 Jiang Jing
29	 江平	 Jiangping
30	 李应	 Li Ying
31	 穆弘育	 Mu Hongyu
32	 豹子韦	 Leopard Wei
33	 雷横	 Lei Heng
34	 詹一翻	 Zhan Yifan
35	 詹一线	 Zhan Yixian
36	 卢俊义	 Lu Zhishen
37	 栾庭芸	 Luan Tingyu
38	 武松	 Wu Song
39	 戴宗	 Dai Zong
40	 朱通	 Zhu Tong
41	 朱彪	 Zhu Biao
42	 李逵	 Li Kui
43	 张青	 Zhang Qing
44	 张时标	 Zhang Shibiao
45	 顿时可	 Yan Shike
46	 颜时炼	 Yan Shilian
47	 王英	 Wang Ying
48	 解珍	 Xie Zhen
49	 解豹	 Xie Bao
50	 韩濤	 Han Tao
51	 裴宣	 Pei Xuan
52	 裴玉	 Pei Yu
53	 裴行	 Pei Xing
54	 柳堂	 Liu Tang
55	 刘挺书	 Liu Tingshu
56	 桑天影	 Sang Tianying
57	 石遨	 Shi Qian
58	 石秀	 Shi Xiu
59	 石宝	 Shi Bao
60	 杨林	 Yang Lin
61	 杨雄	 Yang Xiong
62	 卢大	 Lu Da
63	 卢方政	 Lu Fangzheng
64	 汤隆	 Tang Long
65	 汤琴	 Tang Qin
66	 概功	 Gai Gong
67	 张大	 Zhang Da
68	 孙新	 Sun Xin
69	 孙霸	 Sun Ba
70	 孙一枝	 Sun Yizhi
71	 孙猿	 Sun Xuan
72	 曹大	 Cu Da
73	 胡烈	 Hu Lie
74	 胡冲	 Hu Chong
75	 黄明山	 Huang Mingshan
76	 郝思文	 Hao Siwen
77	 俞保四	 Yu Baosi
78	 俞觥薰	 Yu Mixun
79	 俞大有	 Yu Dayou
80	 李决	 Li Jue
81	 李安	 Li An
82	 李良才	 Li Liangcai
83	 李应政	 Li Yingcheng
84	 李宏	 Li Hon
85	 李俊	 Li Jun
86	 王定	 Wang Ding
87	 王元霸	 Wang Yuanba
88	 王坚	 Wang Kang
89	 张烁	 Zhang Shuo
90	 张茂	 Zhang Mao
91	 周通	 Zhou Tong
92	 周起	 Zhou Qi
93	 苏青	 Su Qing
94	 袁雄	 Yuan Xiong
95	 韦洪	 Wei Hong
96	 丛宝	 Cong Bo
97	 陈达	 Chen Da
98	 陈世美	 Chen Shimei
99	 陈许山	 Chen Xushan
100	 陈庆之	 Chen Qingzhi
101	 莫然	 Mo Ran
102	 莫大谢	 Mo Daxie
103	 蒋敬旭	 Jiang Jingxu
104	 蒋敬	 Jiang Jing
105	 石秀	 Shi Xiu
106	 柳堂	 Liu Tang
107	 阮小五	 Ruan Xiaowu
108	 杨林	 Yang Lin
*/

// Use 108-character names as SpanID or TraceID
var rawSpanIDs = []string{
	"Lu Junyi", "Hua Rong", "Hua Xiong", "Wu Yong", "Lin Chong", "Chai Jin", "Qin Ming", "Qin Qingtong", "Guan Sheng", "Bao Xu",
	"Bao Zheng", "Dong Ping", "He Xiong", "Lu Fang", "Peng Yingyu", "Yan Qing", "Ruan Xiaoer", "Ruan Xiaowu", "Hu Sanniang", "Yang Zhi",
	"Xu Ning", "Song Jiang", "Du Xing", "Zhang Shun", "Sun Li", "Huang Xin", "Guan Yunchang", "Jiang Jing", "Jiangping", "Li Ying",
	"Mu Hongyu", "Leopard Wei", "Lei Heng", "Zhan Yifan", "Zhan Yixian", "Lu Zhishen", "Luan Tingyu", "Wu Song", "Dai Zong", "Zhu Tong",
	"Zhu Biao", "Li Kui", "Zhang Qing", "Zhang Shibiao", "Yan Shike", "Yan Shilian", "Wang Ying", "Xie Zhen", "Xie Bao", "Han Tao",
	"Pei Xuan", "Pei Yu", "Pei Xing", "Liu Tang", "Liu Tingshu", "Sang Tianying", "Shi Qian", "Shi Xiu", "Shi Bao", "Yang Lin",
	"Yang Xiong", "Lu Da", "Lu Fangzheng", "Tang Long", "Tang Qin", "Gai Gong", "Zhang Da", "Sun Xin", "Sun Ba", "Sun Yizhi",
	"Sun Xuan", "Cu Da", "Hu Lie", "Hu Chong", "Huang Mingshan", "Hao Siwen", "Yu Baosi", "Yu Mixun", "Yu Dayou", "Li Jue",
	"Li An", "Li Liangcai", "Li Yingcheng", "Li Hon", "Li Jun", "Wang Ding", "Wang Yuanba", "Wang Kang", "Zhang Shuo", "Zhang Mao",
	"Zhou Tong", "Zhou Qi", "Su Qing", "Yuan Xiong", "Wei Hong", "Cong Bo", "Chen Da", "Chen Shimei", "Chen Xushan", "Chen Qingzhi",
	"Mo Ran", "Mo Daxie", "Jiang Jingxu", "Jiang Jing", "Shi Xiu", "Liu Tang", "Ruan Xiaowu", "Yang Lin",
}

// MockStandardRelationshipIDs function establishes relationships.
// Make the relationships clear.
// The sequence number of A is 1, the sequence number of B is 2, and the sequence number of C is 3.
// Then the sequence number of C's parent node is before 3.
// If the sequence number refers to the order of appearance, then the order of the parent node should be less than that of the child node.
// This is reasonable.
func MockStandardRelationshipIDs(rawSpanIDs []string) (relationshipIDs []string) {
	relationshipIDs = make([]string, len(rawSpanIDs), len(rawSpanIDs))
	for childID := 0; childID < len(relationshipIDs); childID++ {
		rand.NewSource(time.Now().UnixNano())
		var parentID int
		if childID != 0 {
			// ParentID is between 0 and ChildID, can be equal to 0, but must be less than ChildID
			parentID = rand.Intn(childID) // [0, childID)
		}
		relationshipIDs[childID] = rawSpanIDs[parentID]
	}
	return
}

// There are 108 nodes, but a proportion of nodes will create new TraceIDs.
// the percentage of being a new trace is 30%.
var percentageNewTrace = 30

// MockStandardRawData function
func MockStandardRawData(rawSpanIDs, relationshipIDs []string) (raw []model.TracingData) {
	// Create an empty TracingData slice with the same length as rawSpanIDs
	raw = make([]model.TracingData, len(rawSpanIDs), len(rawSpanIDs))

	// Iterate rawSpanIDs
	for i := 0; i < len(rawSpanIDs); i++ {
		// Fill in the tracing test data by using the previous relationship configure.
		raw[i].Parent.SpanID = relationshipIDs[i]
		raw[i].SpanContext.SpanID = rawSpanIDs[i]

		// Create the source span ID
		if raw[i].Parent.SpanID == raw[i].SpanContext.SpanID {
			raw[i].Parent.TraceID = "00000000000000000000000000000000"
			raw[i].Parent.SpanID = "0000000000000000"
		}

		// If within the probability, generate a new TraceID, otherwise inherit the TraceID of the Parent node.
		// (如果在概率内,就产生新的 TraceID,否则就继承 Parent 节点的 TraceID)
		rand.NewSource(time.Now().UnixNano())
		percentage := rand.Intn(100) + 1
		if percentage <= percentageNewTrace || raw[i].Parent.SpanID == "root" || raw[i].Parent.SpanID == "0000000000000000" {
			raw[i].SpanContext.TraceID = raw[i].SpanContext.SpanID
		}
	}

	// Damn !
	for i := 0; i < len(raw); i++ {
		if raw[i].SpanContext.TraceID != "" {
			for j := 0; j < len(raw); j++ {
				if raw[j].Parent.SpanID == raw[i].SpanContext.SpanID && raw[j].SpanContext.TraceID == "" {
					raw[j].SpanContext.TraceID = raw[i].SpanContext.TraceID
				}
			}
		}

	}

	return
}

func NewInfo(raw []model.TracingData) (info *TracingInfo) {

	info = new(TracingInfo)

	info.spanIDs = make(map[string]*Node, len(raw))
	info.traceIDs = make(map[string]*Node)

	info.root = new(Node)
	info.root.SpanID = "root"

	for i := 0; i < len(raw); i++ {
		currentID := raw[i].SpanContext.SpanID
		currentParentID := raw[i].Parent.SpanID

		info.spanIDs[currentID] = &Node{
			SpanID: currentID,
			Data:   raw[i],
		}

		if currentParentID != "0000000000000000" {
			info.spanIDs[currentID].Parent = info.spanIDs[currentParentID]
		} else {
			info.spanIDs[currentID].Parent = info.root
		}

		info.spanIDs[currentID].Parent.Children = append(info.spanIDs[currentID].Parent.Children, info.spanIDs[currentID])
	}

	for key, value := range info.spanIDs {
		// fmt.Println(key, value)
		if value.Parent.Data.SpanContext.TraceID != value.Data.SpanContext.TraceID {
			info.traceIDs[key] = value
		}
	}

	return
}
