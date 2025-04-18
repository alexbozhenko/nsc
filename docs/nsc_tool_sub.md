## nsc tool sub

Subscribe to a subject on a NATS account

```
nsc tool sub [flags]
```

### Examples

```
nsc tool sub <subject>
nsc tool --queue <name> subject
```

### Options

```
  -a, --account string   account name
  -E, --encrypt          encrypted payload
  -h, --help             help for sub
  -q, --queue string     subscription queue name
  -u, --user string      user name
```

### Options inherited from parent commands

```
  -H, --all-dirs string       sets --config-dir, --data-dir, and --keystore-dir to the same value
      --config-dir string     nsc config directory
      --data-dir string       nsc data store directory
  -i, --interactive           ask questions for various settings
      --keystore-dir string   nsc keystore directory
      --nats string           nats url, defaults to the operator's service URLs
  -K, --private-key string    Key used to sign. Can be specified as role (where applicable),
                              public key (private portion is retrieved)
                              or file path to a private key or private key 
```

### SEE ALSO

* [nsc tool](nsc_tool.md)	 - NATS tools: pub, sub, req, rep, rtt

###### Auto generated by spf13/cobra on 2-Jan-2025
