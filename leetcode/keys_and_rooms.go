package keys_and_rooms

func canVisitAllRooms(rooms [][]int) bool {
    visited := make(map[int]bool)
    visit(rooms, visited, 0)
    return len(visited) == len(rooms)
}

func visit(rooms [][]int, visited map[int]bool, room int) {
    visited[room] = true
    for _, i := range rooms[room] {
        if visited[i] == false {
            visit(rooms, visited, i)
        }
    }
}
