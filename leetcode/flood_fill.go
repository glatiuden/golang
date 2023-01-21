package flood_fill

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    if image[sr][sc] == color {
        return image
    }

    old := image[sr][sc]
    flood(image, sr, sc, color, old)
    return image
}

func flood(image [][]int, sr int, sc int, color int, old int) {
    if sr < 0 || sc < 0 || sr >= len(image) || sc >= len(image[sr]) {
        return;
    }

    curr := image[sr][sc]
    if curr != old {
        return;
    }

    image[sr][sc] = color
    flood(image, sr + 1, sc, color, old);
    flood(image, sr - 1, sc, color, old);
    flood(image, sr, sc + 1, color, old);
    flood(image, sr, sc - 1, color, old);
}