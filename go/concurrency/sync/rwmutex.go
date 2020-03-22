package explore

import "sync"

// 適用讀多寫少的情況
// 同時間只能存在 讀鎖 或 寫鎖 兩者互斥
// 同時間只有一個 goroutine 可以獲得寫鎖
// 同時間可有多個 goroutine 可以獲得讀鎖

// AccountRW .
type AccountRW struct {
	name   string
	amount uint32
	mutex  sync.RWMutex
}

// Deposit .
func (a *AccountRW) Deposit(amount uint32) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.amount = a.amount + amount
}

// Balance .
func (a *AccountRW) Balance() uint32 {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	return a.amount
}
