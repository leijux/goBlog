go build -o ./build/app.exe
Copy-Item ./config/config.json ./build/config/config.json
Copy-Item ./log/gin_log ./build/log –recurse