# defintly - An app to help you define concepts well.

![Defintly logo](.assets/logo.png)

This is the backend application for **defintly**, re-developed as part of a study assignment for the Berlin School of Economics and Law (Hochschule für Wirtschaft und Recht Berlin).

General command line interface overview:
````
Flags:
  -h, --help                             Show context-sensitive help.
      --database-hostname="127.0.0.1"    hostname of PostgreSQL instance
      --database-port=5432               port of the PostgreSQL instance
      --database-user="defintly"         user of the PostgreSQL instance
      --database-password=STRING         password of the PostgreSQL instance
      --database-name="defintly"         name of the PostgreSQL instance database
      --database-ssl-mode="disable"      enable/disable SSL connection to the PostgreSQL instance (see PostgreSQL
                                         documentation of specific values to enable)

Commands:
  serve --database-password=STRING
    Start the webserver.

  import --database-password=STRING --excel-file=STRING
    Import data.
````

Special parameters for ``serve``:
````
      --webserver-hostname="127.0.0.1"    ip to bind the webserver to
      --webserver-port=4269               port to bind the webserver to
````

Special parameters for ``import``:
````
      --excel-file=STRING                Path to the excel file to import data from
````

Importing is available from a special formatted Excel file with sheets "Categories", "Collections", "Concepts" and "Criteria".

Checkout the [Defintly site](https://defintly.glideapp.io) for more information. Also visit [The AGI Sentinel Initiative](http://agisi.org) for more research about Artificial General Intelligence.