<!-- This page reflects every GET parameter keys in the headers -->
<html>
    <head>
        <title>Query XSS Test Server</title>
    </head>
    <body>
        <?php
        foreach ($_GET as $key => $value) {
            header("X-$key: fixed-value");
        }
        ?>
    </body>
</html>