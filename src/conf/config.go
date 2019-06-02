package conf

import "time"

const Verbose = true
const DelayWorker = 100 * time.Millisecond
const DelayClient = 100 * time.Millisecond
const DelayMachine = 750 * time.Millisecond
const TimeoutWorker = 500 * time.Millisecond
const TaskListSize = 10
const ItemListSize = 10
const Workers = 5
const Clients = 5
const Machines = 10
const DelayCeoHi = 100 * time.Millisecond
const DelayCeoLo = 100 * time.Millisecond
const BreakingProb = 0.5
const Repairman = 2
const RepairTime = 300 * time.Millisecond
