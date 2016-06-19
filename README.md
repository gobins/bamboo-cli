# bamboo-cli
Command Line Interface for Atlassian Bamboo

## Support
-----------------------------
The CLI has been tested with the following Atlassian Bamboo versions
* 5.7.2

## sub-commands
get-projects    Returns all Projects on Bamboo
get-plans       Returns all Plans in a Project
get-branches    Returns all Branches in a Project

## Examples
```
go build -o bamboocli

bamboocli --username admin --password password --url "http://192.168.99.100:8085" get-projects
+----------------+-----+-------------------------------------------------------+
| Name           | Key | Link                                                  |
+----------------+-----+-------------------------------------------------------+
| First Project  | FP  | http://192.168.99.100:8085/rest/api/latest/project/FP |
| Second Project | SP  | http://192.168.99.100:8085/rest/api/latest/project/SP |
+----------------+-----+-------------------------------------------------------+
```

## Contributing
-----------------------------
Please contribute by sending a pull request.
If there is an issue due to API changes, please raise an issue.