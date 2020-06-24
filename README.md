# lightshot
## What's this ?
This is a custom server for the screenshot tool [Lightshot](https://prnt.sc/).

## Installation
Run :
```
go run lightshot.go
```

## Usage
### Windows
1. Open your `hosts` file. (`%windir%\system32\drivers\etc\hosts`) 
2. Add this line: `127.0.0.1 upload.prntscr.com` (replace '127.0.0.1' with your server's IP)
3. Restart Lightshot.

### Android
1. Download [Hosts Go](https://play.google.com/store/apps/details?id=dns.hosts.server.change).
2. Enable `Hosts change switch` option.
3. Click on `HOSTS EDITOR` button and then click on the `+` button.
4. Add a new host with `127.0.0.1` (replace '127.0.0.1' with your server's IP) as `IP address` and `upload.prntscr.com` as `Domain`.
5. Click on "START".

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
This project is licensed under the [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/) license.
