<!-- This page reflects the xss key from the query -->
<html>

<head>
    <title>Query XSS Test Server</title>
</head>

<body>
    <?php
    header("X-XSS: {$_GET['xss']}")
    ?>
</body>