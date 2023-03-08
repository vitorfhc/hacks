<!-- This page reflects every GET parameter value in the headers -->
<html>
    <head>
        <title>Query XSS Test Server</title>
    </head>
    <body>
        <?php
        foreach ($_GET as $key => $value) {
            header("X-$key: $value");
        }
        ?>
    </body>
</html>