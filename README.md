# fake-smtp

- [FakeSMTP](https://github.com/junlapong/FakeSMTP) java, desktop

```
java -jar jar/fakeSMTP-2.1-SNAPSHOT.jar -p 2525 -s
```

- [fake-smtp-server](https://github.com/gessnerfl/fake-smtp-server) java, web

```
java -jar -Dfakesmtp.port=2525 -Dserver.port=1080 jar/fake-smtp-server-1.5.0.jar
```

- [fake-smtp-server](https://github.com/ReachFive/fake-smtp-server) javascript, web

```
# npm install -g fake-smtp-server
fake-smtp-server -s 2525 -h 1080
```

## Web

[http://localhost:1080](http://localhost:1080/)

## SMTP

localhost:2525

## Test Send Mail

```
# cd test-send-mail
# go mod tidy

$ go run main.go

Try sending mail...
Mail sent without error
```