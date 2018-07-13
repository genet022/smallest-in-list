// If I was talking to you while I did this problem,
// I had a few questions I would ask:
// ----------------------------------------------------------------------------
// * What is the use case of this data?
// * Is there duplicate data in this circular list?
// * What types are the data? Is it just ints? Could there be strings?
// * Can there be multiple different types of data? 
//     - The flexible solution will handle mixed data of floats/str/ints
// ----------------------------------------------------------------------------

// Just an idea, if customers are often looking for the smallest number in this
// set of data, it may be worth keeping a persistent sorted list of the data. 
// (I'm thinking prices of basketball jerseys. Customers will likely often be 
// sorting a page based on price) Then everytime a getSmallest() is called,
// the mostly or completely sorted list (depending on if the list size has
// increased) will be sorted again using insertion or bubble sort and stored
// in the persistent variable. For the sake of simplicity, we'll keep the
// individual algorithms contained though.

// And of course, I added a bit of personality to my comments.
// If this were a professional project, I'd keep them more succinct :)

package main

import (
    "fmt"
    "sort"
    "container/ring"
    "strconv"
    "errors"
)

func main() {
    // Just some static sample data
    x := []int {
        42, 49, 86, 143,
        234, 334, 401, 435,
        2, 14, 21}

    // Create a new ring with the values from x, because Go is awesome and 
    // natively supports circular lists 
    r := ring.New(len(x))
    for i := 0; i < r.Len(); i++ {
        r.Value = x[i]
        r = r.Next()
    }

    // Print out readable answer
    if lowest, err := readableGetSmallest(r); err == nil {
        fmt.Println("smallest from extendable algorithm : ", lowest)
    } else {
        fmt.Println(err)
    }

    // Print out flexible answer
    if lowest, err := flexibleGetSmallest(r, 0); err == nil {
        fmt.Println("smallest from extendable algorithm : ", lowest)
    } else {
        fmt.Println(err)
    }

    // Print out performant answer
    fmt.Println("smallest from performant algorithm : ", perfGetSmallest(r))
}

// I had the cool idea of possibly returning some sort of iterator here
// It could add some flexibility in how you use this function,
// but I kept it fairly simple instead
func flexibleGetSmallest(r *ring.Ring, index int) (float64, error) {
    var arr []float64

    // Put each ring value into an array
    for i := 0; i < r.Len(); i++ {
        switch r.Value.(type) {
        case int:
            arr = append(arr, float64(r.Value.(int)))
            r = r.Next()
        case float32:
            arr = append(arr, float64(r.Value.(float32)))
            r = r.Next()
        case float64:
            arr = append(arr, r.Value.(float64))
            r = r.Next()
        case string:
            if val, err := strconv.ParseFloat(r.Value.(string), 64); err == nil {
                arr = append(arr, val)
                r = r.Next()
            } else {
                return -1, errors.New(err.Error())
            }
        default:
            return -1, errors.New("Data type not yet supported")
        }   
    }

    // Sort the array (smallest to largest)
    sort.Float64s(arr)

    return arr[index], nil
}

// I would definitely consider this "as readable as possible" :)
// A more readable version would've just been a copy/paste of flexibleGetSmallest()
// without the index parameter, hardcoding arr[0], or error/type handling.
// But duplicate code is bad so I'm going with this.
// I would also consider perfGetSmallest sufficiently readable.
func readableGetSmallest(r *ring.Ring) (float64, error) {
    // index 0 = smallest
    // index 1 = second smallest
    // index 2 = third smallest
    var index int = 0
    return flexibleGetSmallest(r, index)
}

// O(n)
// I'm just going to assume the data is all ints for this one
func perfGetSmallest(r *ring.Ring) int {
    // Set smallest to first value
    var smallest int = r.Value.(int)

    // Do a single loop through ring to determine smallest number
    r.Do(func(p interface{}) {
		if p.(int) < smallest {
            smallest = p.(int)
        }
    })
    
    return smallest
}