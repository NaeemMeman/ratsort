package ratsort

import (
	"sort"
	"regexp"
	"strconv"
)

//Public Methods

//Sort implementation
func Sort(s []string) {
    sort.Sort(stringSlice(s))
}

//Less implementation
func Less(i string, j string) bool {
    //chunkify the strings
	iChunks := chunkify(i)
	jChunks := chunkify(j)

    //get the number of chunks identified for the strings
	iChunksNum := len(iChunks)
	jChunksNum := len(jChunks)

	for idx := range iChunks {
	    //j ends before i
		if idx >= jChunksNum {
			return false
		}
    
		iFloat, iErr := strconv.ParseFloat(iChunks[idx], 64)
		jFloat, jErr := strconv.ParseFloat(jChunks[idx], 64)

		//both numeric
		if iErr == nil && jErr == nil {
		    //same number
			if iFloat == jFloat {
			    //last i chunk
				if idx == iChunksNum - 1 {
				    //j > i
					return true
			    //last j chunk
				} else if idx == jChunksNum - 1 {
				    //i > j
					return false
				}
				//look at next chunk
				continue
			}
            
            //different numbers
			return iFloat < jFloat
		}

		//same strings
		if iChunks[idx] == jChunks[idx] {
		    //last i chunk
			if idx == iChunksNum - 1 {
				//j > i
				return true
		    //last j chunk
			} else if idx == jChunksNum - 1 {
				//i > j
				return false
			}
            //look at next chunk
			continue
		}

        //different strings
		return iChunks[idx] < jChunks[idx]
	}
    
    //i could be an empty string while j contains something
	return true
}

//Private Methods

//Attaching sorting interface to stringSlice
type stringSlice []string

func (s stringSlice) Len() int {
	return len(s)
}

func (s stringSlice) Less(i, j int) bool {
	return Less(s[i], s[j])
}

func (s stringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//Helper methods for Less implementation
var chunkifyRegexp = regexp.MustCompile(`(\d+\.\d+|\d+|\D+)`)

func chunkify(s string) []string {
	return chunkifyRegexp.FindAllString(s, -1)
}

