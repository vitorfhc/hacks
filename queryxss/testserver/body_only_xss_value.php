<!-- This page reflects the xss key from the query -->
<html>

<head>
    <title>Query XSS Test Server</title>
</head>

<body>
    <?php
    if (isset($_GET['xss'])) {
        echo "<p>xss: {$_GET['xss']}</p>";
    }
    ?>
</body>