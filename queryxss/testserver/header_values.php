<!-- This page reflects every GET parameter values in the headers -->
<html>
    <head>
        <title>Query XSS Test Server</title>
    </head>
    <body>
        <?php
        foreach ($_GET as $key => $value) {
            header("X-Fixed-Value: $value");
        }
        ?>
    </body>
</html>