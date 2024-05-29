# ASCII-ART-WEB

This project is created to print an art of ascii symbols.

```
              _____     _____    _____    _____                          _____     _______
    /\       / ____|   / ____|  |_   _|  |_   _|                /\      |  __ \   |__   __|
   /  \     | (___    | |         | |      | |     ______      /  \     | |__) |     | |
  / /\ \     \___ \   | |         | |      | |    |______|    / /\ \    |  _  /      | |
 / ____ \    ____) |  | |____    _| |_    _| |_              / ____ \   | | \ \      | |
/_/    \_\  |_____/    \_____|  |_____|  |_____|            /_/    \_\  |_|  \_\     |_|


```

# Developer:

Abay Aliyev (abaliyev),

# Instructions

1. Using terminal, run server using `go run app.go` command

```
 _____    _    _   _   _          _____   ______   _____   __      __  ______   _____
|  __ \  | |  | | | \ | |        / ____| |  ____| |  __ \  \ \    / / |  ____| |  __ \
| |__) | | |  | | |  \| |       | (___   | |__    | |__) |  \ \  / /  | |__    | |__) |
|  _  /  | |  | | | . ` |        \___ \  |  __|   |  _  /    \ \/ /   |  __|   |  _  /
| | \ \  | |__| | | |\  |        ____) | | |____  | | \ \     \  /    | |____  | | \ \
|_|  \_\  \____/  |_| \_|       |_____/  |______| |_|  \_\     \/     |______| |_|  \_\


```

2. Follow link that printed in terminal (`http://localhost:8080`)
3. You can write only one argument.
4. Only ascii symbols from 32 til 126 are available (from ascii table)
5. Choose any style and done! Your ascii-art printed in the bottom field.

```
 _____                               _
|  __ \                             | |
| |  | |    ___     _ __      ___   | |
| |  | |   / _ \   | '_ \    / _ \  | |
| |__| |  | (_) |  | | | |  |  __/  |_|
|_____/    \___/   |_| |_|   \___|  (_)


```

# Good luck using this program

```
  _____     ____      ____     _____            _         _    _     _____    _  __
 / ____|   / __ \    / __ \   |  __ \          | |       | |  | |   / ____|  | |/ /
| |  __   | |  | |  | |  | |  | |  | |         | |       | |  | |  | |       | ' /
| | |_ |  | |  | |  | |  | |  | |  | |         | |       | |  | |  | |       |  <
| |__| |  | |__| |  | |__| |  | |__| |         | |____   | |__| |  | |____   | . \
 \_____|   \____/    \____/   |_____/          |______|   \____/    \_____|  |_|\_\

```

go test -coverprofile=data_test.out
go tool cover -html=data_test.out
go test -short -v

-list
-count
-json
