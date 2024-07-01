package main

/**
你可以假设除了数字 0 之外，这两个数字都不会以零开头。

实例 1：
输入：l1 = 7->2->4->3, l2 = 5->6->4
输出：7->8->0->7
示例2：

输入：l1 = 2->4->3, l2 = 5->6->4
输出：8->0->7



	3->4->2->7
	4->6->5

*/

func main() {
	//记录每个Fingerprint是根据ars_log哪些数据算出来的
	arsLogIdsMap := make(map[string][]int64)

	//保证 arsId 不重复
	arsLogList := []string{"1", "2"}
	for _, v := range arsLogList {

		arsLogIdsMap[v] = append(arsLogIdsMap[v], v.Id)
	}

}
