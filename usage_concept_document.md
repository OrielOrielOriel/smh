# snake-many-head usage concept doc

## Example usage

Commands
```
smh
    smh help
    smh brute
    smh spray
    smh load
```

Universal Useage Format
```
smh [protocol] --<options>/-<o> "value" Target(s)>
               --config/-c "config.toml"
               --blob/-b   "base64"
```

## Examples
Bruteforce an SSH login
`smh ssh -o "username:Nana" -o "password_file:passwords.txt" 1.2.3.4`

Bruteforce an HTTP Post form
`smh http -o "header:'Cookie: beans'" -o 'method:post' -o 'ignore_ssl:true' -o 'body:user=^USER^&pass=^PASS^' -o 'user:Nana' -o 'password_file:pass.txt' https://example.com/login.php`
`smh http --blob "request:base64 encoded burp request" --blob "config:base64 encoded config file" https://example.com/login.php`
