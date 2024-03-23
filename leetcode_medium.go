// spiral order

type Location struct {
    x int
    y int
}

func turnRight (c Location) Location {

    var o Location

    if c.x == 0 && c.y == -1 {
            o.x = 1
            o.y = 0
    } else if c.x == 1 && c.y == 0 {
            o.x = 0
            o.y = 1
    } else if c.x == 0 && c.y == 1 {
            o.x = -1
            o.y = 0
    } else if c.x == -1 && c.y == 0 {
            o.x = 0
            o.y = -1 
    }

    return o
}

func moveForward (cur_location Location, cur_dir Location) Location {

    var out_location Location

    out_location.x = cur_location.x + cur_dir.x
    out_location.y = cur_location.y + cur_dir.y

    return out_location
}

func spiralOrder(matrix [][]int) []int {

    m := matrix

    out := make([]int, 1)

    var cur_loc Location
    cur_loc.x = 0
    cur_loc.y = 0
  
    var facing Location 
    facing.x = 1
    facing.y = 0

    x_max := len(matrix[0]) - 1 
    y_max := len(matrix) - 1

    out[0] = m[cur_loc.y][cur_loc.x];

    for {
        m[cur_loc.y][cur_loc.x] = -999; // mark current location as -999

        // can we move forward
        new_loc := moveForward(cur_loc,facing)
        if new_loc.x >= 0 && new_loc.x <= x_max && new_loc.y >= 0 && new_loc.y <= y_max && m[new_loc.y][new_loc.x] != -999 {
            cur_loc = new_loc
        } else {
            facing = turnRight(facing)
            new_loc = moveForward(cur_loc,facing)
            if new_loc.x >= 0 && new_loc.x <= x_max && new_loc.y >= 0 && new_loc.y <= y_max && m[new_loc.y][new_loc.x] != -999 {
                cur_loc = new_loc
            } else { break}   
        }

     out = append(out, m[cur_loc.y][cur_loc.x]);    
     }

    return out
}
