package main

import (
	"container/list"
	"fmt"
	"strings"
)

type Depend struct {
	num     int
	depType int // 0 - same, 1 - reverse
}

type Plate struct {
	pos     int
	depends []Depend
}

type Plates struct {
	plates []Plate
	answer string
}

type Move struct {
	i   int  // plate index
	dir byte // 'L' or 'R'
}

var count, min, max int = 0, 1, 7
var base = make(map[string]bool)

func main() {
	fmt.Print("Enter the count of plates: ")
	_, err := fmt.Scanf("%d", &count)
	if err != nil {
		fmt.Println("err_0:", err.Error())
		return
	}
	fmt.Scanln()
	fmt.Println()

	start := make([]Plate, count)
	fmt.Println("Enter the position of the pins on the plates, starting from the one closest to you")
	for i := 0; i < count; i++ {
		fmt.Printf("Enter the position of the %d pin: ", i+1)
		_, err := fmt.Scanf("%d", &start[i].pos)
		if err != nil {
			fmt.Println("err_1:", err)
		}
		fmt.Scanln()

		fmt.Print("Enter the count of dependent plates: ")
		var dependenceCount int
		_, err = fmt.Scanf("%d", &dependenceCount)
		if err != nil {
			fmt.Println("err_2:", err)
		}
		fmt.Scanln()
		for j := 0; j < dependenceCount; j++ {
			fmt.Print("Enter the number of dependent plate: ")
			var depPlate int
			_, err = fmt.Scanf("%d", &depPlate)
			if err != nil {
				fmt.Println("err_3:", err)
			}
			fmt.Scanln()

			fmt.Print("Enter the type of dependence (S/s - same, R/r - reverse): ")
			var typeOfDep int
			_, err = fmt.Scanf("%d", &typeOfDep)
			if err != nil {
				fmt.Println("err_4:", err)
			}
			fmt.Scanln()

			temp := Depend{num: depPlate - 1, depType: typeOfDep}
			start[i].depends = append(start[i].depends, temp)
		}
		fmt.Println()
	}

	startState := Plates{plates: start}
	startKey := makeKey(startState)

	queue := list.New()

	visited := make(map[string]bool)
	parent := make(map[string]string)
	parentMove := make(map[string]Move)

	visited[startKey] = true
	parent[startKey] = "" // корень
	queue.PushBack(startState)

	var goal string
	for queue.Len() > 0 && goal == "" {
		front := queue.Remove(queue.Front()).(Plates)
		k := makeKey(front)

		// Check the goal
		if gigaCheck(front.plates) {
			goal = k
			break
		}

		for i := 0; i < count; i++ {
			// Try R (+1) и L (-1)
			for _, dir := range []byte{'R', 'L'} {
				side := -1
				if dir == 'L' {
					side = 1
				}

				next := clonePlates(front)
				next.plates[i].pos += side
				if !check(next.plates[i]) {
					continue
				}

				ok := true
				for _, d := range next.plates[i].depends {
					next.plates[d.num] = minorShift(next.plates[d.num], d.depType, side)
					if !check(next.plates[d.num]) {
						ok = false
						break
					}
				}
				if !ok {
					continue
				}

				nk := makeKey(next)
				if visited[nk] {
					continue
				}
				visited[nk] = true
				parent[nk] = k
				parentMove[nk] = Move{i: i, dir: dir}
				queue.PushBack(next)
			}
		}
	}

	if goal == "" {
		fmt.Println("No answer.")
		return
	}

	// recovery
	steps := make([]Move, 0)
	for ck := goal; ck != startKey; ck = parent[ck] {
		steps = append(steps, parentMove[ck])
	}

	// reverse steps
	for l, r := 0, len(steps)-1; l < r; l, r = l+1, r-1 {
		steps[l], steps[r] = steps[r], steps[l]
	}

	parts := make([]string, 0, len(steps))
	for _, m := range steps {
		parts = append(parts, fmt.Sprintf("%d%c", m.i+1, m.dir))
	}
	fmt.Printf("%s\n", strings.Join(parts, " "))
}

func makeKey(p Plates) string {
	var sb strings.Builder
	for _, v := range p.plates {
		sb.WriteByte(byte('0' + v.pos))
		// switch v.pos {
		// case 1:
		// 	sb.WriteString("1")
		// case 2:
		// 	sb.WriteString("2")
		// case 3:
		// 	sb.WriteString("3")
		// case 4:
		// 	sb.WriteString("4")
		// case 5:
		// 	sb.WriteString("5")
		// case 6:
		// 	sb.WriteString("6")
		// case 7:
		// 	sb.WriteString("7")
		// }
	}
	return sb.String()
}

func check(plate Plate) bool {
	return plate.pos >= min && plate.pos <= max
}

func gigaCheck(plates []Plate) bool {
	for _, p := range plates {
		if p.pos != 4 {
			return false
		}
	}
	return true
}

func minorShift(plate Plate, typeOfShift, side int) Plate {
	temp := plate
	if typeOfShift == 0 {
		temp.pos += side
	} else {
		temp.pos -= side
	}
	return temp
}

func clonePlates(pl Plates) Plates {
	plates := make([]Plate, len(pl.plates))
	copy(plates, pl.plates)
	return Plates{plates: plates, answer: pl.answer}
}

// func majorShift(pl Plates, plate_num, side int) *Node {
// 	temp := clonePlates(pl)
// 	temp.plates[plate_num].pos += side
// 	if check(temp.plates[plate_num]) {
// 		for _, p := range temp.plates[plate_num].depends {
// 			temp.plates[p.num] = minorShift(temp.plates[p.num], p.depType, side)
// 			if !check(temp.plates[p.num]) {
// 				return nil
// 			}
// 		}

// 		if side == -1 {
// 			temp.answer += fmt.Sprintf("%dR ", plate_num+1)
// 		} else {
// 			temp.answer += fmt.Sprintf("%dL ", plate_num+1)
// 		}

// 		if ok := gigaCheck(temp.plates); ok {
// 			result = temp.answer
// 			return nil
// 		}

// 		key := makeKey(temp)
// 		if _, okey := base[key]; !okey {
// 			base[key] = true
// 			node := &Node{p: temp, node: nil}
// 			return node
// 		}
// 	}
// 	return nil
// }
