<?php
$files = scandir(".");
array_shift($files);
array_shift($files);
$files = array_diff($files, array("index.php"));
$files = array_map(function ($file) {
    return "http://localhost:8080/$file";
}, $files);
foreach ($files as $file) {
    echo "$file ";
}
?>