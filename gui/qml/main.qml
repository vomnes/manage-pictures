import QtQuick 2.7			//ApplicationWindow
import QtQuick.Controls 2.1	//Dialog
import QtQuick.Layouts 1.12

ApplicationWindow {
	id: window

	visible: true
	title: "Rename pictures(.jpg/.jpeg) folder"
	width: 400
	height: 50

	RowLayout {
		id: layout
    anchors.fill: parent
		anchors.leftMargin: 5
		anchors.rightMargin: 15
		spacing: 7.5

		Rectangle {
			Layout.fillWidth: true
			Layout.minimumHeight: 20
			Layout.maximumHeight: 25
			id: textArea

			Label {
				text: JSON.parse(qmlInitContext)["Label"].text
			}
		}

		Rectangle {
			Layout.minimumWidth: 100
			Layout.maximumWidth: 150

			Button {
				anchors.centerIn: parent
				text: "Browse folder"
				onClicked: qmlBridge.sendToGo(parent.objectName, "browseFolder", "")
			}
		}

		Rectangle {
			Layout.minimumWidth: 50
			Layout.maximumWidth: 50

			Button {
				visible: JSON.parse(qmlInitContext)["Save"].visible

				anchors.centerIn: parent
				text: "Run"
				onClicked: qmlBridge.sendToGo(parent.objectName, "renamePictures", "")
			}
		}
	}
}
