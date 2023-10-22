package main

import (
    "fmt"
    "net1_recording/lib"
    "os"
    "flag"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

func main() {
    addr := ":8080"

    serverFlagPtr := flag.Bool("server", true, "Flag for server")
    clientNameFlagPtr := flag.String("client", "", "The client user name")
    flag.Parse()

    serverFlag := *serverFlagPtr

    if serverFlag {

        server  , err := lib.NewServer(addr)

        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        server.Start()
        return

    } 

    if *clientNameFlagPtr == "" {
        fmt.Println("Client needs a client name.")
        os.Exit(4)
    }



    client, err := lib.NewClient(addr, *clientNameFlagPtr)


    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }

    //go client.Start()
    go client.Listen()
    defer client.Close()

    mainApp := app.New()
    mainWindow := mainApp.NewWindow(*clientNameFlagPtr)

    mainWindow.Resize(fyne.NewSize(300, 400))

    textField := widget.NewEntry()
    textField.PlaceHolder = "Enter Message"

    messageData := binding.NewUntypedList()

    messageListWidget := widget.NewListWithData(
        messageData,
        func () fyne.CanvasObject {
            return container.NewBorder(
                nil,
                nil,
                nil,
                nil,
                widget.NewLabel(""),
            )
        },
        func (di binding.DataItem, o fyne.CanvasObject) {
            ctr := o.(*fyne.Container)
            label := ctr.Objects[0].(*widget.Label)
            value, _ := di.(binding.Untyped).Get()
            label.SetText(value.(string))

        },
    )


    client.SetMessageData(messageData)

    mainWindow.SetContent(
        container.NewBorder(
            nil,
            container.NewBorder(
                nil,
                nil,
                nil,
                widget.NewButton("Send", func() {
                    fmt.Println("Send pressed")
                    message := textField.Text
                    client.Send(message)
                    textField.SetText("")

                }),
                textField,

            ),
            nil,
            nil,
            messageListWidget,

        ),
    )

    mainWindow.ShowAndRun()


}

