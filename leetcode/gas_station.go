package go_station

func canCompleteCircuit(gas []int, cost []int) int {
    total_gas := 0
    total_cost := 0
    start := 0
    current_gas := 0

    for i := 0; i < len(gas); i++ {
        total_gas += gas[i]
        total_cost += cost[i]
        current_gas += gas[i] - cost[i]
        if current_gas < 0 {
            start = i + 1
            current_gas = 0
        }
    }

    is_possible := total_cost <= total_gas
    if is_possible {
        return start
    }
    return -1
}