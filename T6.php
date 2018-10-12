<?php 

// you can write to stdout for debugging purposes, e.g.
// print "this is a debug message\n";

function solution($X, $Y, $D) {
    // write your code in PHP7.0
    
    $jump = intval( ($Y - $X) / $D );
    
    if ( ($jump * $D) + $X < $Y) {
        $jump ++;
    }
    
    return $jump;
}