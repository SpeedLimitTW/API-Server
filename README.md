# Getting Start
Want to starting server, paste this line to command.

```sh
go run main.go
```

then server will be started listen in `localhost:8080`.
# REST API
### Beep
Request beep to finding oldman, there is three method to control beep function.

##### `/isBeep`
To check current beep state online.

return `Bool`, `true` or `false`.

##### `/setBeep`
To set beep state to `fasle`, this api is for MT7688 device, if they finish beep, they will call this API to reset state.

##### `/requestBeep`
To set beep state to `true`, after called, state will pendding to be call.