<!-- This page reflects every GET parameter value in the body -->
<html>
    <head>
        <title>Query XSS Test Server</title>
    </head>
    <body>
        <?php
        foreach ($_GET as $key => $value) {
            echo "<p>$value</p>";
        }
        ?>
    </body>
</html>