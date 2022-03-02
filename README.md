# plantuml

This piece of code generates sprites to use with plantuml.

The icons sources should be fetched from [Google's official site](https://cloud.google.com/icons)

## generate icons

- Download the set of icons from [Google's official site](https://cloud.google.com/icons)
- run the code `go run .`

## Use the icons in plantuml

example:

```plantuml
@startuml Sample
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

!define GCPPuml https://raw.githubusercontent.com/owulveryck/PlantUML-icons-GCP/master/official
!include GCPPuml/GCPCommon.puml
!include GCPPuml/vertexai/vertexai.puml
System_Boundary(system, "sample system","sample") {
 System_Boundary(application,"WebApp") {
  Container(vertex,"Vertex","Vertex","vertex", "vertexai") {
   Container(code,"Python Engine","Python", "awesome code", "python")
  }
 }
}

@enduml
```
