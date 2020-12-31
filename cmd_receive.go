package kuronekocat

import (
	"log"
	"sort"

	"github.com/spf13/cobra"
)

func newReceiveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "receive [id or explain]",
		Short: "get kuroneko infomations",

		Args: cobra.MinimumNArgs(1),

		RunE: execReceiveCmd,
	}
	return cmd
}

func execReceiveCmd(_ *cobra.Command, args []string) error {
	orders, err := readFromHomeJSON()
	if err != nil {
		return err
	}

	type remoeFlag struct {
		index  int
		remove bool
	}

	orderInfoMap := make(map[string]*remoeFlag)
	for i, order := range orders {
		rf := &remoeFlag{
			index:  i,
			remove: false,
		}
		orderInfoMap[order.Explain] = rf
		orderInfoMap[order.ID] = rf
	}

	for _, arg := range args {
		if rf, ok := orderInfoMap[arg]; ok {
			rf.remove = true
			continue
		} else {
			log.Printf("[error] not found queue at %s", arg)
		}
	}

	replaceIndexMap := make(map[int]bool)
	var removeIndexList []int

	for _, rf := range orderInfoMap {
		if rf.remove {
			continue
		}
		if _, ok := replaceIndexMap[rf.index]; !ok {
			removeIndexList = append(removeIndexList, rf.index)
			replaceIndexMap[rf.index] = true
		}
	}

	sort.Slice(removeIndexList, func(i, j int) bool { return removeIndexList[i] < removeIndexList[j] })
	var newOrders []order

	for _, i := range removeIndexList {
		newOrders = append(newOrders, orders[i])
	}

	return writeForHomeJSON(newOrders)
}
