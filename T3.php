<?php 

// you can write to stdout for debugging purposes, e.g.
// print "this is a debug message\n";

function solution($A, $K) {
    // write your code in PHP7.0
    $newArray = array();
    
    for ($i=0; $i < sizeOf($A); $i++) { 
        $newIndex = ($i + $K) % sizeof($A);
        $newArray[$newIndex] = $A[$i];
    }
    
    return $newArray;
}

