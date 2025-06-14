# krokifier

A configurable syntax based approach to parsing generic files into Kroki Compatable UML Diagrams

## The Idea Example

Take this as an example. We have a set of Imports we want to quickly diagram to visualize your 
dependecies. The file will be in python and we want a `blockdiag` output.

Imports:
```python
  from nirn_weaver.ui import ESPManager
  from nirn_weaver.ui import PAKManager
  from nirn_weaver.ui import OBSEManager
  from nirn_weaver.ui import UE4SSManager
  from nirn_weaver import NirnPaths
  from textual.app import App, ComposeResult
  from textual.widgets import Header, Footer, TabbedContent, TabPane, Label
```

We would then write our `syntax` file as this. Take note that parts of this syntax example such
as `keywords` and `extending` can be broken out into a seperate `language` file later in development

Syntax:
```json
{
    "uml_type": "blockdiag",
    "file_language": "python",
    "diagram": {
        "import_group": {
            "label": "IMPORTS",
            "keywords": [
                "from",
                "import"
            ],
            "extending": ","
        }
    }
}
```

Krokifier will then parse the above imports into ready to use blockdiag UML

Into
```blockdiag
// "Auto-Generated" with Krokifer V my.brain
blockdiag {
  group {
    label="IMPORTS" 
    nirn_weaver -> ui -> ESPManager 
    nirn_weaver -> ui -> PAKManager
    nirn_weaver -> ui -> OBSEManager
    nirn_weaver -> ui -> UE4SSManager
    nirn_weaver -> NirnPaths
    textual -> app -> App
    textual -> app -> ComposeResult
    textual -> widgets -> Header
    textual -> widgets -> Footer
    textual -> widgets -> TabbedContent
    textual -> widgets -> TabPane
    textual -> widgets -> Label
  }
}
```
