## nsc revocations delete_activation

Remove an account revocation from an export

```
nsc revocations delete_activation [flags]
```

### Options

```
  -a, --account string          account name
  -h, --help                    help for delete_activation
      --service                 service
  -s, --subject string          export subject
  -t, --target-account string   target-account
```

### Options inherited from parent commands

```
  -H, --all-dirs string       sets --config-dir, --data-dir, and --keystore-dir to the same value
      --config-dir string     nsc config directory
      --data-dir string       nsc data store directory
  -i, --interactive           ask questions for various settings
      --keystore-dir string   nsc keystore directory
  -K, --private-key string    Key used to sign. Can be specified as role (where applicable),
                              public key (private portion is retrieved)
                              or file path to a private key or private key 
```

### SEE ALSO

* [nsc revocations](nsc_revocations.md)	 - Manage revocation for users and activations from an account

###### Auto generated by spf13/cobra on 2-Jan-2025
