package main

func base() Configuration {
	return Configuration{
		Esp8266Url:            "ws://esp8266-door-monitor:80",
		DisplayErrorCommand:   "",
		DoorOpenedCommand:     "",
		DoorClosedCommand:     "",
		OnStartOpenCommand:    "",
		OnStartClosedCommand:  "",
		WaitUntilCloseCommand: "",
	}
}

func windows() Configuration {
	c := base()

	/* Powershell error message
	Add-Type -AssemblyName System.Windows.Forms; $global:x = New-Object System.Windows.Forms.NotifyIcon; $x.BalloonTipText = 'My text'; $x.BalloonTipTitle = "My Title"; $x.Visible = $true; $x.ShowBalloonTip(20000)
	*/

	c.DoorOpenedCommand = "(New-Object -ComObject Shell.Application).minimizeall()"

	return c
}
