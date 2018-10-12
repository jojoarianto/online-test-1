<?php 

// you can write to stdout for debugging purposes, e.g.
// print "this is a debug message\n";

function solution($A) {
    // write your code in PHP7.0
    $temp = array();
    
    // get all positif number
    for ($i=0; $i < sizeof($A); $i++) { 
        if ($A[$i] > 0) {
            array_push($temp, $A[$i]);
        }
    }
    
    $ct = 1;
    $loop = true;
    
    // searching method
    while ($loop) {
        // searching value from array
        if (in_array($ct, $temp)) {
            //  deleting an element from array
            array_diff($temp, $ct);
            $ct++;
        } else {
            $loop = false;
            break;
        }
    }
    
    return $ct;

}
