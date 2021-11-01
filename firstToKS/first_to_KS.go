package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type firstItem struct {
	pc     string
	pcNext string
	action string
	same   string
	samePC string
}

type state struct {
	pca string
	pcb string
	t   int
	seq int
}
type seq struct {
	parent int
	child  int
}

var statea state
var stateb state
var globalANum int
var globalBNum int
var global int
var varMap = make(map[string]int)
var itemList []firstItem
var handleList []state
var globalStateNum int
var seqMap []seq
var buff string
func getItem(str string) firstItem {
	s := strings.Split(str, "&")
	pca := strings.Split(s[0], "=")
	pcb := strings.Split(s[1], "=")
	w := strings.Split(s[4], " |")
	newFirstItem := firstItem{pca[1], pcb[1], s[2], s[3], w[0]}
	return newFirstItem
}

func matchItems(st state) []firstItem {
	var matchItems []firstItem
	for i := 0; i < len(itemList); i++ {
		if itemList[i].pc == st.pca || itemList[i].pc == st.pcb {
			matchItems = append(matchItems, itemList[i])
		}
	}
	return matchItems
}

func notExist(st state) int {
	for i := 0; i < len(handleList); i++ {
		if ((st.pca == handleList[i].pca) && (st.pcb == handleList[i].pcb)) && (st.t == handleList[i].t) {
			return handleList[i].seq
		}
	}
	return -1
}
func test() {
	li := []int{3,4,7,8,10,11}
	for i:=0;i<len(li);i++ {
		seqMap = append(seqMap, seq{li[i],li[i]})
	}
}
func deal(item state) {
	matchItems := matchItems(item)
	for i := 0; i < len(matchItems); i++ {
		if matchItems[i].samePC == "Same(PCa)" {
			if matchItems[i].same == "Same(V)" {
				//说明这里是判断表达式
				if matchItems[i].action == "true" {
					newState := state{matchItems[i].pcNext, item.pcb, item.t, global + 1}
					if reSeq:=notExist(newState);reSeq!=-1 {
						seqMap = append(seqMap, seq{item.seq, reSeq})
						return
					}
					handleList = append(handleList, newState)

					seqMap = append(seqMap, seq{item.seq, global + 1})
					global = global + 1
				} else {
					s := strings.Split(matchItems[i].action, "==")
					if (len(s) < 2) {
						return
					}
					num, _ := strconv.Atoi(s[1])
					if item.t == num {

						newState := state{matchItems[i].pcNext, item.pcb, item.t, global + 1}
						if reSeq:=notExist(newState);reSeq!=-1 {
							seqMap = append(seqMap, seq{item.seq, reSeq})

							return
						}

						handleList = append(handleList, newState)

						seqMap = append(seqMap, seq{item.seq, global + 1})
						global = global + 1
					}
				}
			}

		} else if matchItems[i].samePC == "Same(PCb)" {
			
			if matchItems[i].same == "Same(V)" {
				if matchItems[i].action == "true" {
					newState := state{item.pca, matchItems[i].pcNext, item.t, global + 1}
					if reSeq:=notExist(newState);reSeq!=-1 {
						seqMap = append(seqMap, seq{item.seq, reSeq})
						return
					}
					handleList = append(handleList, newState)
					seqMap = append(seqMap, seq{item.seq, global + 1})
					global = global + 1
				} else {
					s := strings.Split(matchItems[i].action, "==")
					if (len(s)<2) {
						return
					}
					num, _ := strconv.Atoi(s[1])
					if item.t == num {
						newState := state{item.pca, matchItems[i].pcNext, item.t, global + 1}
						if reSeq:=notExist(newState);reSeq!=-1 {
							seqMap = append(seqMap, seq{item.seq, reSeq})
							//globalANum = globalANum + 1
							return
						}
						handleList = append(handleList, newState)
						seqMap = append(seqMap, seq{item.seq, global + 1})
						global = global + 1
					}
				}
			}
		}
	}

}
func main() {
	file, err := os.Open("/root/analysisVerificationOfsystem/firstToKS/fake_first")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := bufio.NewScanner(file)
	var newItem firstItem
	for {
		if !buf.Scan() {
			break //文件读完了,退出for
		}
		line := buf.Text() //获取每一行
		newItem = getItem(line)
		itemList = append(itemList, newItem)
	}

	globalANum = 0
	globalBNum = 1
	global = 1
	globalStateNum = 0
	handleList = append(handleList, state{"La1", "Lb1", 0, 1})
	handleList = append(handleList, state{"La1", "Lb1", 1, 0})
	i := 0

	for {
		if len(handleList) == i {
			break
		}
		dealItem := handleList[i]
		deal(dealItem)
		i++
	}
	test()
	out_file, err := os.OpenFile("/root/analysisVerificationOfsystem/firstToKS/outputks", os.O_CREATE|os.O_WRONLY, 0)
    if err != nil {
        fmt.Println("打开文件失败")
        return
	}
	fmt.Fprintf(file, "abc123222")
	
	for i := 0; i < len(handleList); i++ {
		fmt.Fprintf(out_file, fmt.Sprintf("S%d[label=\"S%d\\n(%s,%s,%d)\"]\n",handleList[i].seq,handleList[i].seq,handleList[i].pca,handleList[i].pcb,handleList[i].t))
	}
	for _, v := range seqMap {
		fmt.Fprintf(out_file, fmt.Sprintf("S%d->S%d\n",v.parent,v.child))
	}
	out_file.Close()
}