package bpTree

func makeLinkList(bpDatas []*BpData) {
	for i := 0; i < len(bpDatas)-1; i++ {
		bpDatas[i].Next = bpDatas[i+1]
		bpDatas[(len(bpDatas)-1)-i].Previous = bpDatas[(len(bpDatas)-1)-i-1]
	}
	return
}

func createRootTree7and5and11to13() (idx *BpIndex) {
	bpDatas := []*BpData{
		{
			Items: []BpItem{ // 第 0 笔
				{Key: 1},
				{Key: 1},
			},
		},
		{
			Items: []BpItem{ // 第 1 笔
				{Key: 3},
				{Key: 4},
			},
		},
		{
			Items: []BpItem{ // 第 2 笔
				{Key: 4},
				{Key: 5},
			},
		},
		{
			Items: []BpItem{ // 第 3 笔
				{Key: 5},
			},
		},
		{
			Items: []BpItem{ // 第 4 笔
				{Key: 6},
			},
		},
		{
			Items: []BpItem{ // 第 5 笔
				{Key: 7},
				{Key: 9},
			},
		},
		{
			Items: []BpItem{ // 第 6 笔
				{Key: 10},
			},
		},
		{
			Items: []BpItem{ // 第 7 笔
				{Key: 11},
			},
		},
		{
			Items: []BpItem{ // 第 8 笔
				{Key: 12},
			},
		},
		{
			Items: []BpItem{ // 第 9 笔
				{Key: 13},
				{Key: 15},
			},
		},
		{
			Items: []BpItem{ // 第 10 笔
				{Key: 15},
				{Key: 15},
			},
		},
		{
			Items: []BpItem{ // 第 11 笔
				{Key: 15},
				{Key: 15},
			},
		},
	}

	makeLinkList(bpDatas)

	idx = &BpIndex{
		Index: []int64{7},
		IndexNodes: []*BpIndex{
			{
				Index: []int64{5},
				IndexNodes: []*BpIndex{
					{
						Index: []int64{3, 4},
						DataNodes: []*BpData{
							bpDatas[0],
							bpDatas[1],
							bpDatas[2],
						},
					},
					{
						Index: []int64{6},
						DataNodes: []*BpData{
							bpDatas[3],
							bpDatas[4],
						},
					},
				},
			},
			{
				Index: []int64{11, 13},
				IndexNodes: []*BpIndex{
					{
						Index: []int64{10},
						DataNodes: []*BpData{
							bpDatas[5],
							bpDatas[6],
						},
					},
					{
						Index: []int64{12},
						DataNodes: []*BpData{
							bpDatas[7],
							bpDatas[8],
						},
					},
					{
						Index: []int64{15, 15},
						DataNodes: []*BpData{
							bpDatas[9],
							bpDatas[10],
							bpDatas[11],
						},
					},
				},
			},
		},
	}

	return
}

func createRootTree15and15to15and15() (idx *BpIndex) {
	bpDatas := []*BpData{
		{
			Items: []BpItem{ // 第 0 笔
				{Key: 15},
				{Key: 15},
			},
		},
		{
			Items: []BpItem{ // 第 1 笔
				{Key: 15},
				{Key: 15},
			},
		},
		{
			Items: []BpItem{ // 第 2 笔
				{Key: 15},
				{Key: 15},
			},
		},
		{
			Items: []BpItem{ // 第 3 笔
				{Key: 15},
				{Key: 15},
			},
		},
		{
			Items: []BpItem{ // 第 4 笔
				{Key: 15},
				{Key: 15},
			},
		},
	}

	makeLinkList(bpDatas)

	idx = &BpIndex{
		Index: []int64{15},
		IndexNodes: []*BpIndex{
			{
				Index: []int64{15, 15},
				DataNodes: []*BpData{
					bpDatas[0],
					bpDatas[1],
					bpDatas[2],
				},
			},
			{
				Index: []int64{15},
				DataNodes: []*BpData{
					bpDatas[3],
					bpDatas[4],
				},
			},
		},
	}

	return
}
