Simple CLI to control [L6470](https://www.st.com/en/motor-drivers/l6470.html) via [MCP2210](https://www.microchip.com/wwwproducts/en/MCP2210).

## Requirement

| Title | Content |
|---|---|
| OS | Windows 10 |
| Compiler | gcc |
| Boards | MCP2210 board, L6470 board and a stepping motor. |

## How to build

Download DLL files from [Microchip web site](https://www.microchip.com/wwwproducts/en/MCP2210) and save DLL file to root directory of this project.

And then, build exe file.

```
go build -o main.exe
```