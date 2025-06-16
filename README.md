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
  ...
```

We would then write our `syntax` file as this.

Syntax:
```json
{
	"uml_type": "blockdiag",
	"language": "python",
	"diagram": [
	    {
			"label": "IMPORT",
            "keywords": [
		        {
			        "word": "from",
                    "extension": ".",
			        "extends": true,
                    "recursive": false,
                    "inherits": false
		        },
		        {
			        "word": "import",
			        "extension": ",",
			        "extends": true,
                    "recursive": true,
                    "inherits": false
		        }	
	        ]   
		},
        {
            "label": "CLASSES",
            "keywords": [
                {
                    "word": "class",
                    "extension": "",
                    "extends": false,
                    "recursive": false,
                    "inherits": true
                }
            ]
        },
        {
            "label": "METHODS",
            "keywords": [
                {
                    "word": "def",
                    "extension": "",
                    "extends": false,
                    "recursive": false,
                    "inherits": true 
                }
            ]
        }    
	]
}
```

Krokifier will then parse the above imports into ready to use blockdiag UML

Into

```blockdiag
blockdiag {

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

App -> NirnWeaver

NirnWeaver -> __init__
NirnWeaver -> on_mount
NirnWeaver -> compose
NirnWeaver -> action_re_stage

group {
    label="IMPORTS"
    color="#a4bde1"
    nirn_weaver; 
    ui; 
    ESPManager; 
    PAKManager; 
    OBSEManager;
    UE4SSManager;
    NirnPaths;
    textual; 
    app; 
    App;
    widgets;
    ComposeResult
    Header
    Footer
    TabbedContent
    TabPane
    Label
}

group {
    label="CLASSES"
    color="#c4adb1"
    NirnWeaver;
}

group {
    label="METHODS"
    color="#bcadb1"
    __init__;
    on_mount;
    compose;
    action_re_stage;
}
}
```

Generating the following BlockDiag.

![Example Block Diagram])(https://github.com/ScorpioGameKing/krokifier/blob/main/assets/diagram_example.png)
