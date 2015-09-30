package main

import (
	"github.com/codegangsta/cli"

	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "jv"
	app.Usage = "Query a JSON file with a path-like thing."
	app.Action = mainAction

	app.Run(os.Args)
}

func mainAction(c *cli.Context) {
	if len(c.Args()) < 2 {
		fmt.Printf("Query string and filename are required\n")
		os.Exit(1)
	}
	q := c.Args()[0]
	f := c.Args()[1]

	if err := query(q, f); err != nil {
		fmt.Printf("Failed query: %s\n", err)
		os.Exit(1)
	}
}

func query(q, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %s\n", filename, err)
		os.Exit(1)
	}
	defer f.Close()

	parts := strings.Split(q, "/")

	data, err := parse(f)
	if err != nil {
		fmt.Printf("Error parsing file %s: %s\n", filename, err)
		os.Exit(2)
	}

	return match(data, parts, 0)
}

func parse(r io.Reader) (interface{}, error) {
	dec := json.NewDecoder(r)
	var thingy interface{}
	return &thingy, dec.Decode(&thingy)
}

func puts(d interface{}) {
	k := reflect.Indirect(reflect.ValueOf(d)).Kind()
	switch k {
	case reflect.Struct, reflect.Map, reflect.Array, reflect.Slice:
		b, _ := json.Marshal(d)
		fmt.Printf("%s\n", b)
	default:
		fmt.Printf("%v\n", d)
	}
}

func match(jdata interface{}, q []string, i int) error {
	if len(q) == i {
		// Print the token
		//fmt.Printf("%v\n", jdata)
		puts(jdata)
		return nil
	}
	if jdata == nil {
		panic("Why is the token nil?")
	}

	target := q[i]
	if target == "" {
		i++
		return match(jdata, q, i)
	}

	// Basically, we just want to safely dereference any pointers.
	data := reflect.Indirect(reflect.ValueOf(jdata)).Interface()

	if val, ok := data.([]interface{}); ok {
		ii, err := strconv.Atoi(target)
		if err != nil {
			return fmt.Errorf("Expected an integer index on an array. Got %s.", data)
		}

		if len(val) <= ii {
			fmt.Printf("No value found at index %d\n", ii)
			return nil
		}
		i++
		return match(val[ii], q, i)
	} else if val, ok := data.(map[string]interface{}); ok {
		// Gotta love the fumpt.
		ss := fmt.Sprintf("%v", target)

		vv, ok := val[ss]
		if !ok {
			fmt.Printf("No value found for key %s\n", ss)
			return nil
		}
		i++
		return match(vv, q, i)
	} else {
		ss := fmt.Sprintf("%v", data)
		if target == ss {
			i++
			return match(data, q, i)
		}
	}

	return fmt.Errorf("no match found")
}
