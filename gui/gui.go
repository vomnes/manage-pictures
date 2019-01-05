package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kr/pretty"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/widgets"

	"../rename-pictures/rename"
)

var (
	qmlObjects = make(map[string]*core.QObject)

	qmlBridge *QmlBridge

	state = make(map[string]map[string]interface{})

	directoryPath = ""

	loading = rename.Loading{}
)

type QmlBridge struct {
	core.QObject

	_ func(source, action, data string) `signal:"sendToQml"`
	_ func(source, action, data string) `slot:"sendToGo"`

	_ func(object *core.QObject) `slot:"registerToGo"`
	_ func(objectName string)    `slot:"deregisterToGo"`
}

func trimStringBefore(s string, char string) string {
	if idx := strings.LastIndex(s, char); idx != -1 {
		return s[idx:]
	}
	return s
}

func updateContext(quickWidget *quick.QQuickWidget, data map[string]map[string]interface{}) {
	var b, err = json.Marshal(data)
	if err != nil {
		log.Println("initQmlContext", err)
	}
	quickWidget.RootContext().SetContextProperty2("qmlInitContext", core.NewQVariant14(string(b)))
}

func initQmlBridge(quickWidget *quick.QQuickWidget) {

	qmlBridge = NewQmlBridge(nil)
	quickWidget.RootContext().SetContextProperty("qmlBridge", qmlBridge)

	qmlBridge.ConnectSendToGo(func(source, action, data string) {
		if action == "browseFolder" {
			fd := widgets.NewQFileDialog(nil, 0)
			fd.SetFileMode(widgets.QFileDialog__Directory)
			fd.ConnectDirectoryEntered(func(directory string) {
				loading = rename.StatFolders(directory)
				directoryPath = directory
				state["Save"]["visible"] = true
				state["Label"]["text"] = fmt.Sprintf("\r...%s (%d files)", trimStringBefore(directory, "/"), loading.Total)
				updateContext(quickWidget, state)
			})
			fd.Exec()
		}
		if action == "renamePictures" {
			go rename.Run(directoryPath, &loading)
			percent, current := 0, 0
			for loading.Done != loading.Total {
				current = 100 * loading.Done / loading.Total
				if current > percent {
					pretty.Println(fmt.Sprintf("[%3d%%] - ...%s (%d files)", current, trimStringBefore(directoryPath, "/"), loading.Total))
				}
			}
			state["Label"]["text"] = fmt.Sprintf("[OK] - ...%s (%d files)", trimStringBefore(directoryPath, "/"), loading.Total)
			state["Save"]["visible"] = false
			updateContext(quickWidget, state)
		}
	})

	qmlBridge.ConnectRegisterToGo(func(object *core.QObject) {
		qmlObjects[object.ObjectName()] = object
	})

	qmlBridge.ConnectDeregisterToGo(func(objectName string) {
		qmlObjects[objectName] = nil
	})
}

func initQmlContext(quickWidget *quick.QQuickWidget) {

	state = map[string]map[string]interface{}{
		"Button": {
			"text": "Browse folder",
		},
		"Save": {
			"text":    "Run",
			"visible": false,
		},
		"Label": {
			"text": "Select your folder...",
		},
	}

	updateContext(quickWidget, state)
}

func newQmlWidget() *quick.QQuickWidget {
	var quickWidget = quick.NewQQuickWidget(nil)
	quickWidget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)

	initQmlBridge(quickWidget)
	initQmlContext(quickWidget)

	quickWidget.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))

	return quickWidget
}

func main() {

	// enable high dpi scaling
	// useful for devices with high pixel density displays
	// such as smartphones, retina displays, ...
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	// Material, Default, Fusion, Imagine, Universal
	quickcontrols2.QQuickStyle_SetStyle("Imagine")

	// needs to be called once before you can start using QML
	widgets.NewQApplication(len(os.Args), os.Args)

	var layout = widgets.NewQHBoxLayout()
	layout.AddWidget(newQmlWidget(), 0, 0)

	var window = widgets.NewQMainWindow(nil, 0)

	var centralWidget = widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(layout)
	window.SetCentralWidget(centralWidget)

	window.Show()

	widgets.QApplication_Exec()
}
