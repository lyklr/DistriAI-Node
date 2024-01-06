package core_task

import (
	"DistriAI-Node/chain/distri"
	"DistriAI-Node/chain/distri/distri_ai"
	"DistriAI-Node/docker"
	"DistriAI-Node/pattern"
	logs "DistriAI-Node/utils/log_utils"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gagliardetto/solana-go"
)

func OrderComplete(distri *distri.WrapperDistri, metadata string, isGPU bool, containerID string) {
	logs.Normal("Order is complete")

	if err := docker.StopWorkspaceContainer(containerID); err != nil {
		logs.Error(fmt.Sprintf("Stopping Workspace container error: %v", err))
		return
	}

	var orderPlacedMetadata pattern.OrderPlacedMetadata

	err := json.Unmarshal([]byte(metadata), &orderPlacedMetadata)
	if err != nil {
		logs.Error(fmt.Sprintf("error unmarshaling order metadata: %v", err))
		return
	}

	orderPlacedMetadata.MachineAccounts = distri.ProgramDistriMachine.String()

	distri.OrderCompleted(orderPlacedMetadata, isGPU)
}

func OrderFailed(distri *distri.WrapperDistri, metadata string, buyer solana.PublicKey, containerID string) {
	logs.Normal("Order is failed")

	if err := docker.StopWorkspaceContainer(containerID); err != nil {
		logs.Error(fmt.Sprintf("Stopping Workspace container error: %v", err))
		return
	}

	var orderPlacedMetadata pattern.OrderPlacedMetadata

	err := json.Unmarshal([]byte(metadata), &orderPlacedMetadata)
	if err != nil {
		logs.Error(fmt.Sprintf("error unmarshaling order metadata: %v", err))
		return
	}

	orderPlacedMetadata.MachineAccounts = distri.ProgramDistriMachine.String()

	distri.OrderFailed(buyer, orderPlacedMetadata)
}

func CheckOrder(done chan bool, distri *distri.WrapperDistri, oldDuration time.Time) {
	newOrder, err := distri.GetOrder()
	if err != nil {
		logs.Error(fmt.Sprintf("Error: %v", err))
		done <- false
		return
	}

	newDuration := time.Unix(newOrder.OrderTime, 0).Add(time.Hour * time.Duration(newOrder.Duration))

	logs.Normal(fmt.Sprintf("CheckOrder newDuration: %v", newDuration))
	logs.Normal(fmt.Sprintf("CheckOrder oldDuration: %v", oldDuration))

	if newDuration.After(oldDuration) {
		logs.Normal("Restart timer")
		if !StartTimer(distri, newOrder) {
			done <- false
			return
		}
	}
	done <- true
}

func StartTimer(distri *distri.WrapperDistri, order distri_ai.Order) bool {
	done := make(chan bool)

	duration := time.Unix(order.OrderTime, 0).Add(time.Hour * time.Duration(order.Duration))
	logs.Normal(fmt.Sprintf("Order OrderTime: %v", time.Unix(order.OrderTime, 0)))
	logs.Normal(fmt.Sprintf("Order Add: %v", time.Hour*time.Duration(order.Duration)))
	logs.Normal(fmt.Sprintf("Order duration: %v", duration))
	time.AfterFunc(time.Until(duration), func() {
		CheckOrder(done, distri, duration)
	})
	return <-done
}