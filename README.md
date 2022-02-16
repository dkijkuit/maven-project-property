# maven-project-property

This little utility fetches basic project information from a Maven pom.xml without having Maven installed. Output is sent back to the console.

| **Parameter** | **Allowed values**           | **Description**                                                                                                     |
|---------------|------------------------------|---------------------------------------------------------------------------------------------------------------------|
| `-f`          | `string - default 'pom.xml'` | Specify the location and filename of the pom file.                                                                  |
| `-c`          | `char - default '-'`         | Delimiter to eliminate trailing data. For example 1.5.4-SNAPSHOT will return 1.5.4 using the default '-' delimiter. |
| `-p`          | `string - default 'version'` | Property to fetch from the xml file. Valid values are: `groupId, artifactId, name, version, description`.           |
|               |                              |                                                                                                                     |
