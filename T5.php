<?php 

// you can write to stdout for debugging purposes, e.g.
// print "this is a debug message\n";

function solution($A) {
    // write your code in PHP7.0
    
    asort($A);
    $miss = 0;
    
    $size = count($A);
    for ($i=1; $i <= $size; $i++) { 
        if (in_array($i, $A)) {
            
        } else {
            $miss = $i;
            break;
        }
    }
    
    return $miss;
}