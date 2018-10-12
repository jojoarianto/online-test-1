<?php 

// you can write to stdout for debugging purposes, e.g.
// print "this is a debug message\n";

function solution($A) {
    // write your code in PHP7.0
    
    $newArray = array();
    $qty = array();

    $size = count($A);
    for ($i=0; $i < $size; $i++) { 
        // searching method
        if (in_array($A[$i], $newArray) == false) {
            array_push( $newArray, $A[$i] );
        }
    }
    
    $sizeNewArray = count($A);
    for ($i=0; $i < $sizeNewArray; $i++) { 
        for ($j=0; $j < $size; $j++) { 
            if ($A[$j] == $newArray[$i]) {
                $qty[$i] = $qty[$i] + 1;
            }
        }
    }
    
    $idx = array_search(1, $qty);
    
    return $newArray[$idx];
}