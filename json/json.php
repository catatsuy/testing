<?php

ini_set('memory_limit', -1);

$gz = file_get_contents("./testdata/code.json.gz");

$start = microtime(true);
for ($i = 0; $i < 1000; $i++) {
    $data = gzdecode($gz);
    $s = json_decode($data);
    json_encode($s);
}
$time_elapsed_secs = microtime(true) - $start;
var_dump($time_elapsed_secs);