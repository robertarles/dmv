# DMV

A simple registry server and client in one. The purpose is to be able to easily stand up a single server which can record and report (JSON format) basic info about any system that registers itslef.

## Example Use

- system A, the DMV registry server:
`dmv serve`

- system B, a client:
`dmv register -n $HOST -s [IP/hostname of DMV registry server]`
    -OR-
in a browser, open `http://[IP/hostname of DMV registry server]/register/[HOSTNAME]` (You must supply [HOSTNAME] in this case)

- get a report of registered systems:
`dmv list -s [IP/hostname of DMV registry server]`
    -OR-
in a browser, open `http://[IP/hostname of DMV registry server]`
example result:

```json
{
    "system1":{"hostname":"system1","remoteAddr":"192.168.168.101","requestURI":"/register/system1"},
    "system2":{"hostname":"system2","remoteAddr":"192.168.168.102","requestURI":"/register/system2"},
    "system3":{"hostname":"system3","remoteAddr":"192.168.168.103","requestURI":"/register/system3"}
}
```


```bash
Usage:
  dmv [command]

Available Commands:
  delete      A brief description of your command
  help        Help about any command
  list        A brief description of your command
  register    A brief description of your command
  serve       stand up a DMV http server

Flags:
      --config string   config file (default is $HOME/.dmv.yaml)
  -h, --help            help for dmv
  -t, --toggle          Help message for toggle

Use "dmv [command] --help" for more information about a command.
```
