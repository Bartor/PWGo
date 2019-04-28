package conf

import "time"

const Verbose = true
const DelayWorker = 100 * time.Millisecond
const DelayClient = 100 * time.Millisecond
const DelayMachine = 1000 * time.Millisecond
const TimeoutWorker = 5 * time.Millisecond
const TaskListSize = 10
const ItemListSize = 10
const Workers = 5
const Clients = 5
const Machines = 10
const DelayCeoHi = 100 * time.Millisecond
const DelayCeoLo = 100 * time.Millisecond
