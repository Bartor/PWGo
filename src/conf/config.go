package conf

import "time"

const Verbose = true
const DelayWorker = 100 * time.Millisecond
const DelayClient = 5 * time.Second
const DelayMachine = 100 * time.Millisecond
const TimeoutWorker = 50 * time.Millisecond
const TaskListSize = 10
const ItemListSize = 10
const Workers = 10
const Clients = 3
const Machines = 4
const DelayCeoHi = 500 * time.Millisecond
const DelayCeoLo = 100 * time.Millisecond
